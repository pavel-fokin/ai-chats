import { useEffect, useRef } from 'react';

import { useInvalidateMessages } from 'hooks';
import { EventTypes } from 'types';

import { useChatContext } from './use-chat-context';

type EventHandler = (event: MessageEvent) => void;
const eventHandlers = new Map<string, EventHandler>();

export function useChatEvents(chatId: string) {
  const eventSourceRef = useRef<EventSource | null>(null);
  const { setModelResponse } = useChatContext();
  const invalidateMessages = useInvalidateMessages();

  const accessToken = localStorage.getItem('accessToken') || '';

  eventHandlers.set(EventTypes.MESSAGE_ADDED, (event) => {
    const messageAdded = JSON.parse(event.data);
    setModelResponse(null);
    invalidateMessages(messageAdded.chatId);
  });

  eventHandlers.set(EventTypes.MODEL_STREAM_MESSAGE, (event) => {
    const message = JSON.parse(event.data);
    setModelResponse(message);
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
