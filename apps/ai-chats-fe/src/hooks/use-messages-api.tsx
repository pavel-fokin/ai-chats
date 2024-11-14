import {
  skipToken,
  useMutation,
  useQuery,
  useQueryClient,
} from '@tanstack/react-query';

import { getMessages, postMessages } from 'api';

export const useGetMessages = (chatId: string | undefined) => {
  return useQuery({
    queryKey: ['messages', chatId],
    queryFn: chatId ? () => getMessages(chatId) : skipToken,
    select: (data) => data.data.messages,
  });
};

export const useSendMessage = (chatId: string) => {
  return useMutation({
    mutationFn: (text: string) => {
      return postMessages(chatId, text);
    },
  });
};

export const useInvalidateMessages = () => {
  const queryClient = useQueryClient();

  return (chatId: string) => {
    queryClient.invalidateQueries({
      queryKey: ['messages', chatId],
    });
  };
};
