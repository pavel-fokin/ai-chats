import {
  skipToken,
  useMutation,
  useQuery,
  useQueryClient,
} from '@tanstack/react-query';

import { deleteChats, fetchChatById, postChats, fetchChats } from 'api';

export const useChats = () => {
  return useQuery({
    queryKey: ['chats'],
    queryFn: fetchChats,
    select: (data) => data.data.chats,
  });
};

export const useChat = (chatId: string | undefined) => {
  return useQuery({
    queryKey: ['chat', chatId],
    queryFn: chatId ? () => fetchChatById(chatId) : skipToken,
    select: (data) => data.data.chat,
  });
};

export const useCreateChat = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (message: string) => postChats(message),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['chats'],
      });
    },
  });
};

export const useDeleteChat = (chatId: string) => {
  return useMutation({
    mutationFn: () => deleteChats(chatId),
  });
};
