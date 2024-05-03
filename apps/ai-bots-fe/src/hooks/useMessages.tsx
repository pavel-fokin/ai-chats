import { useQuery } from '@tanstack/react-query';

import { fetchMessages } from 'api';
import { Message } from 'types';

export function useMessages(chatID: string): Message[] {
    const { data: payload = [] } = useQuery({
        queryKey: ['messages', chatID],
        queryFn: () => fetchMessages(chatID),
    });

    return Array.isArray(payload) ? payload : payload.data.messages || [];
}