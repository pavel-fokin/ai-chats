export type Chat = {
  id: string;
  title: string;
  createdAt: string;
};

export type Message = {
  id: string;
  sender: string;
  text: string;
  createdAt: string;
};

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
