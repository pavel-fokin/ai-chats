import { useEffect, useRef } from 'react';

import { useInvalidateChat, useInvalidateChats } from 'shared/hooks';

type EventHandler = (event: MessageEvent) => void;
const eventHandlers = new Map<string, EventHandler>();

export function useAppEvents() {
  const eventSourceRef = useRef<EventSource | null>(null);
  const accessToken = localStorage.getItem('accessToken') || '';

  const invalidateChat = useInvalidateChat();
  const invalidateChats = useInvalidateChats();

  eventHandlers.set('chatTitleUpdated', (event) => {
    const chatTitleUpdated = JSON.parse(event.data);

    invalidateChat(chatTitleUpdated.chatId);
    invalidateChats();
  });

  useEffect(() => {
    const eventSource = new EventSource(
      `/api/events/app?accessToken=${accessToken}`,
    );
    eventSourceRef.current = eventSource;

    eventSource.onopen = () => {
      console.log('Connection to server opened.');
    };

    for (const [eventType, eventHandler] of eventHandlers) {
      eventSource.addEventListener(eventType, eventHandler);
    }

    eventSource.onerror = (error) => {
      console.error('EventSource failed:', error);
    };

    return () => {
      eventSource.close();
    };
  });
}
