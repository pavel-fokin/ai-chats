import {
  skipToken,
  useMutation,
  useQuery,
  useQueryClient,
} from '@tanstack/react-query';

import {
  deleteChats,
  getChatById,
  getChats,
  postChats,
  postGenerateChatTitle,
} from 'api';
import { PostChatsRequest } from 'api/requests';

export const useGetChats = () => {
  return useQuery({
    queryKey: ['chats'],
    queryFn: getChats,
    select: (data) => data.data.chats,
  });
};

export const useGetChat = (chatId: string | undefined) => {
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
};

export const useInvalidateChats = () => {
  const queryClient = useQueryClient();

  return () => {
    queryClient.invalidateQueries({
      queryKey: ['chats'],
    });
  };
};

export const useInvalidateChat = () => {
  const queryClient = useQueryClient();

  return (chatId: string) => {
    queryClient.invalidateQueries({
      queryKey: ['chat', chatId],
    });
  };
};
