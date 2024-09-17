import { useEffect, useState, useRef } from 'react';

import * as types from 'types';
import { useInvalidateMessages } from 'hooks/useMessagesApi';

type EventHandler = (event: MessageEvent) => void;
const eventHandlers = new Map<string, EventHandler>();

export function useChatEvents(chatId: string) {
  const eventSourceRef = useRef<EventSource | null>(null);
  const [messageChunk, setMessageChunk] = useState<types.MessageChunk>(
    {} as types.MessageChunk,
  );
  const invalidateMessages = useInvalidateMessages();

  const accessToken = localStorage.getItem('accessToken') || '';

  eventHandlers.set(types.EventTypes.MESSAGE_ADDED, (event) => {
    const messageAdded = JSON.parse(event.data);
    invalidateMessages(messageAdded.chatId);
  });

  eventHandlers.set(types.EventTypes.MESSAGE_CHUNK_RECEIVED, (event) => {
    const messageChunk = JSON.parse(event.data);
    if (messageChunk.done) {
      setMessageChunk({} as types.MessageChunk);
    }
    setMessageChunk(messageChunk);
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

  return { messageChunk };
}
