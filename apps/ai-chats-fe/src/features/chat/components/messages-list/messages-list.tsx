import * as types from 'types/types';

import { Message } from '../message';

interface MessagesListProps {
  messages: types.Message[];
}

// Messages list component.
export const MessagesList = ({ messages }: MessagesListProps): JSX.Element => {
  return (
    <>
      {messages.map((message) => (
        <Message key={message.id} sender={message.sender} text={message.text} />
      ))}
    </>
  );
};
