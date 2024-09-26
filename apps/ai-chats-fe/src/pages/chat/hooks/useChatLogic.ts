import { useChatEvents } from 'features/chat';
import { useMessages, useSendMessage } from 'shared/hooks';

export const useChatLogic = (chatId: string) => {
  useChatEvents(chatId);
  const messages = useMessages(chatId);
  const sendMessage = useSendMessage(chatId);

  return {
    messages,
    sendMessage,
  };
};
