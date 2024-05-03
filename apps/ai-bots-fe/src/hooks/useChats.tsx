import { useQuery } from '@tanstack/react-query';

import { GetChats } from 'api';
import { Chat } from 'types';

export function useChats(): Chat[]{
    const { data: chats = [] } = useQuery({
        queryKey: ['chats'],
        queryFn: GetChats,
    });
    return chats;
}