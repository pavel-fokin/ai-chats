import { Chat, Message } from 'types';
import { doGet, doPost, doDelete } from './base';

type Error = {
  message: string;
};

type PostChatsResponse = {
  data: {
    chat: Chat;
  };
};

type GetChatsResponse = {
  data: {
    chats: Chat[];
  };
};

type GetChatByIdResponse = {
  data: {
    chat: Chat;
  };
};

type DeleteChatsResponse = {
  errors?: Error[];
};

type GetMessagesResponse = {
  data: {
    messages: Message[];
  };
};

type PostMessagesResponse = {
  data: {
    message: Message;
  };
};

export const fetchChatById = async (
  chatId: string,
): Promise<GetChatByIdResponse> => {
  return await doGet<GetChatByIdResponse>(`/api/chats/${chatId}`);
};

export const postChats = async (message: string): Promise<PostChatsResponse> => {
  return await doPost('/api/chats', { message });
};

export const deleteChats = async (
  chatId: string,
): Promise<DeleteChatsResponse> => {
  return await doDelete(`/api/chats/${chatId}`);
};

export const fetchChats = async (): Promise<GetChatsResponse> => {
  return await doGet<GetChatsResponse>('/api/chats');
};

export const fetchMessages = async (
  chatId: string,
): Promise<GetMessagesResponse> => {
  return await doGet(`/api/chats/${chatId}/messages`);
};

export const postMessages = async (
  chatId: string,
  message: Message,
): Promise<PostMessagesResponse> => {
  return await doPost(`/api/chats/${chatId}/messages`, message);
};
