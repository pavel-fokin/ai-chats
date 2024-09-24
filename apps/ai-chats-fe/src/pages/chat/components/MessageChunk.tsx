import { useChatContext, Message } from 'features/chat';

export const MessageChunk = () => {
  const { messageChunk } = useChatContext();

  if (!messageChunk) {
    return null;
  }

  return <Message sender={messageChunk.sender} text={messageChunk.text} />;
};
