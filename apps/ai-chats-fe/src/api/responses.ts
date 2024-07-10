import { Chat, Message } from 'types';

export type Response<T> = {
  data?: T;
  errors?: string[];
};

export type AccessToken = {
  accessToken: string;
};

export type Error = {
  message: string;
};

export type PostChatsResponse = {
  data: {
    chat: Chat;
  };
};

export type GetChatsResponse = {
  data: {
    chats: Chat[];
  };
};

export type GetChatByIdResponse = {
  data: {
    chat: Chat;
  };
};

export type DeleteChatsResponse = {
  errors?: Error[];
};

export type GetMessagesResponse = {
  data: {
    messages: Message[];
  };
};

export type PostMessagesResponse = {
  data: {
    message: Message;
  };
};
