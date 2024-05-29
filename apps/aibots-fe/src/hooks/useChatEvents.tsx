import { useEffect, useState } from 'react';

import * as types from 'types';
import { useAuth, useMessages } from 'hooks';


export function useChatEvents(chatId: string) {
    const [messageChunk, setMessageChunk] = useState<types.MessageChunk>({} as types.MessageChunk)

    const { accessToken } = useAuth();
    const { invalidateMessages } = useMessages(chatId);

    useEffect(() => {
        const eventSource = new EventSource(`/api/chats/${chatId}/events?accessToken=${accessToken}`);

        eventSource.onopen = () => {
            console.log('Connection to server opened.');
        };

        eventSource.onmessage = (event) => {
            const message = JSON.parse(event.data);
            switch (message.type) {
                case types.EventTypes.MESSAGE_ADDDED:
                    invalidateMessages();
                    break;
                case types.EventTypes.MESSAGE_CHUNK_RECEIVED:
                    if (message.done) {
                        setMessageChunk({} as types.MessageChunk);
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
        }
    }, [chatId]);

    return { messageChunk };
}