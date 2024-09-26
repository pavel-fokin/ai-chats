import { Message } from 'features/chat';
import * as types from 'types/types';

interface MessagesListProps {
  messages: types.Message[];
}

export const MessagesList: React.FC<MessagesListProps> = ({ messages }) => {
  return (
    <>
      {messages.map((message) => (
        <Message key={message.id} sender={message.sender} text={message.text} />
      ))}
    </>
  );
};
