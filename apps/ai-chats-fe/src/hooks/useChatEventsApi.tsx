import { useEffect, useState, useRef } from 'react';

import * as types from 'types';
import { useInvalidateMessages } from 'hooks/useMessagesApi';

export function useChatEvents(chatId: string) {
  const eventSourceRef = useRef<EventSource | null>(null);
  const [messageChunk, setMessageChunk] = useState<types.MessageChunk>(
    {} as types.MessageChunk,
  );
  const invalidateMessages = useInvalidateMessages();

  const accessToken = localStorage.getItem('accessToken') || '';

  useEffect(() => {
    const eventSource = new EventSource(
      `/api/chats/${chatId}/events?accessToken=${accessToken}`,
    );
    eventSourceRef.current = eventSource;

    eventSource.onopen = () => {
      console.log('Connection to server opened.');
    };

    eventSource.onmessage = (event) => {
      const message = JSON.parse(event.data);
      switch (message.type) {
        case types.EventTypes.MESSAGE_ADDED:
          invalidateMessages(chatId);
          break;
        case types.EventTypes.MESSAGE_CHUNK_RECEIVED:
          if (message.done) {
            setMessageChunk({} as types.MessageChunk);
            break;
          }
          setMessageChunk(message);
          break;
        default:
          console.error('Unknown message type:', message.type);
      }
    };

    eventSource.onerror = (error) => {
      console.error('EventSource failed:', error);
    };

    return () => {
      console.log('Closing connection to server.');
      eventSource.close();
    };
  }, [chatId]);

  return { messageChunk };
}
