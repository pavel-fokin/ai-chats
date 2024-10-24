import { useMessages, useSendMessage } from 'hooks';

import { useChatEvents } from './use-chat-events';

export const useChatLogic = (chatId: string) => {
  useChatEvents(chatId);
  const messages = useMessages(chatId);
  const sendMessage = useSendMessage(chatId);

  return {
    messages,
    sendMessage,
  };
};
