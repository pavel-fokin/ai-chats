import { useEffect, useRef } from 'react';

import { useInvalidateMessages } from 'hooks';
import { EventTypes } from 'types';

import { useChatContext } from './use-chat-context';

type EventHandler = (event: MessageEvent) => void;
const eventHandlers = new Map<string, EventHandler>();

export function useChatEvents(chatId: string) {
  const eventSourceRef = useRef<EventSource | null>(null);
  const { setMessageChunk } = useChatContext();
  const invalidateMessages = useInvalidateMessages();

  const accessToken = localStorage.getItem('accessToken') || '';

  eventHandlers.set(EventTypes.MESSAGE_ADDED, (event) => {
    const messageAdded = JSON.parse(event.data);
    setMessageChunk(null);
    invalidateMessages(messageAdded.chatId);
  });

  eventHandlers.set(EventTypes.CHAT_MESSAGE, (event) => {
    const message = JSON.parse(event.data);
    setMessageChunk(message);
  });

  useEffect(() => {
    const eventSource = new EventSource(
      `/api/chats/${chatId}/events?accessToken=${accessToken}`,
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
      console.log('Closing connection to server.');
      eventSource.close();
    };
  }, [chatId, accessToken]);
}
