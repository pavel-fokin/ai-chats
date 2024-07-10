import { skipToken, useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

import { fetchMessages, postMessages } from 'api';
import { Message } from 'types';

export const useMessages = (chatId: string | undefined) => {
  return useQuery({
    queryKey: ['messages', chatId],
    queryFn: chatId ? () => fetchMessages(chatId) : skipToken,
    select: (data) => data.data.messages,
  });
};

export const useSendMessage = (chatId: string) => {
  return useMutation({
    mutationFn: (msg: Message) => {
      return postMessages(chatId, msg);
    },
  });
};

export const useInvalidateMessages = (chatId: string) => {
  const queryClient = useQueryClient();

  return () => {
    queryClient.invalidateQueries({
      queryKey: ['messages', chatId],
    });
  }
}