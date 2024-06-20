import { useQuery, skipToken, useMutation } from '@tanstack/react-query';

import { fetchChatById, deleteChats } from 'api';

export const useChat = (chatId: string | undefined) => {
  return useQuery({
    queryKey: ['chat', chatId],
    queryFn: chatId ? () => fetchChatById(chatId) : skipToken,
    select: (data) => data.data.chat,
  });
};

export const useDeleteChat = (chatId: string) => {
  return useMutation({
    mutationFn: () => deleteChats(chatId),
  });
};
