import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

import { fetchMessages, postMessages } from 'api';
import { Message } from 'types';


export function useMessages(chatId: string) {
    const queryClient = useQueryClient();

    const { data: payload = [] } = useQuery({
        queryKey: ['messages', chatId],
        queryFn: () => fetchMessages(chatId),
    });

    const mutation = useMutation({
        mutationFn: (msg: Message) => {
            return postMessages(chatId, msg);
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: ['messages', chatId],
            });
        }
    });

    return {
        messages: Array.isArray(payload) ? payload : payload.data.messages || [],
        sendMessage: mutation,
    }
}