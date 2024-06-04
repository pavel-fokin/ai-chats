import { Chat, Message } from 'types';
import { fetchData, postData } from './base';

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

const fetchChatById = async (chatId: string): Promise<GetChatByIdResponse> => {
  return await fetchData<GetChatByIdResponse>(`/api/chats/${chatId}`);
};

const postChats = async (): Promise<PostChatsResponse> => {
  return await postData('/api/chats', {});
};

const fetchChats = async (): Promise<GetChatsResponse> => {
  return await fetchData<GetChatsResponse>('/api/chats');
};

const fetchMessages = async (chatId: string): Promise<GetMessagesResponse> => {
  return await fetchData(`/api/chats/${chatId}/messages`);
};

const postMessages = async (
  chatId: string,
  message: Message,
): Promise<PostMessagesResponse> => {
  return await postData(`/api/chats/${chatId}/messages`, message);
};

export { fetchChatById, fetchChats, fetchMessages, postChats, postMessages };
