import {
  skipToken,
  useMutation,
  useQuery,
  useQueryClient,
} from '@tanstack/react-query';

import { deleteChats, getChatById, getChats, postChats, postGenerateChatTitle } from 'api';
import { PostChatsRequest } from 'api/requests';

export const useChats = () => {
  return useQuery({
    queryKey: ['chats'],
    queryFn: getChats,
    select: (data) => data.data.chats,
  });
};

export const useChat = (chatId: string | undefined) => {
  return useQuery({
    queryKey: ['chat', chatId],
    queryFn: chatId ? () => getChatById(chatId) : skipToken,
    select: (data) => data.data.chat,
  });
};

export const useCreateChat = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (req: PostChatsRequest) => postChats(req),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['chats'],
      });
    },
  });
};

export const useDeleteChat = () => {
  return useMutation({
    mutationFn: (chatId: string) => deleteChats(chatId),
  });
};

export const useGenerateChatTitle = () => {
  return useMutation({
    mutationFn: (chatId: string) => postGenerateChatTitle(chatId),
  });
}