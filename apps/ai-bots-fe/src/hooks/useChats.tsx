import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

import { fetchChats, postChats } from 'api';

export function useChats() {
    const queryClient = useQueryClient();

    const { data: chats = [] } = useQuery({
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

    return {
        chats,
        createChat: mutation,
    }
}