import { useGetMessages, useSendMessage } from 'hooks';

import { useChatEvents } from './use-chat-events';

export const useChatLogic = (chatId: string) => {
  useChatEvents(chatId);
  const messages = useGetMessages(chatId);
  const sendMessage = useSendMessage(chatId);

  return {
    messages,
    sendMessage,
  };
};
