import { Chat, Message } from 'types';
import { fetchData, postData, deleteData } from './base';

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
  return await fetchData<GetChatByIdResponse>(`/api/chats/${chatId}`);
};

export const postChats = async (): Promise<PostChatsResponse> => {
  return await postData('/api/chats', {});
};

export const deleteChats = async (
  chatId: string,
): Promise<DeleteChatsResponse> => {
  return await deleteData(`/api/chats/${chatId}`);
};

export const fetchChats = async (): Promise<GetChatsResponse> => {
  return await fetchData<GetChatsResponse>('/api/chats');
};

export const fetchMessages = async (
  chatId: string,
): Promise<GetMessagesResponse> => {
  return await fetchData(`/api/chats/${chatId}/messages`);
};

export const postMessages = async (
  chatId: string,
  message: Message,
): Promise<PostMessagesResponse> => {
  return await postData(`/api/chats/${chatId}/messages`, message);
};
