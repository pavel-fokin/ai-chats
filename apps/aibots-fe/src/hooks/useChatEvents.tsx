import { useEffect } from 'react';

import { useAuth, useMessages } from 'hooks';


export function useChatEvents(chatId: string) {
    const { accessToken } = useAuth();
    const { invalidateMessages } = useMessages(chatId);

    useEffect(() => {
        const eventSource = new EventSource(`/api/chats/${chatId}/events?accessToken=${accessToken}`);

        eventSource.onopen = () => {
            console.log('Connection to server opened.');
        };

        eventSource.onmessage = () => {
            console.log('Received event');
            invalidateMessages();
        };

        eventSource.onerror = (error) => {
            console.error('EventSource failed:', error);
        };

        return () => {
            console.log('Closing connection to server.');
            eventSource.close();
        }
    }, [chatId]);
}