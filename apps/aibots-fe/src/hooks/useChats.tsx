import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

import { fetchChats, postChats, fetchChatById } from 'api';

export function useChats() {
    const queryClient = useQueryClient();

    const { data: chats } = useQuery({
        queryKey: ['chats'],
        queryFn: fetchChats,
    });

    const mutation = useMutation({
        mutationFn: postChats,
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: ['chats'],
            });
        }
    });

    const getChatById = async (chatId: string) => {
        const { data: chat } = await queryClient.ensureQueryData({
            queryKey: ['chat', chatId],
            queryFn: () => fetchChatById(chatId),
        });
        return chat;
    }

    return {
        chats,
        createChat: mutation,
        getChatById,
    }
}