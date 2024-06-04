import { useQuery, skipToken } from '@tanstack/react-query';

import { fetchChatById } from 'api';

const useChat = (chatId: string | undefined) => {
    return useQuery({
        queryKey: ['chat', chatId],
        queryFn: chatId ? () => fetchChatById(chatId) : skipToken,
        select: (data) => data.data.chat,
    });
}

export { useChat };