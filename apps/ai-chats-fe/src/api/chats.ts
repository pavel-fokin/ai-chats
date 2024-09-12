import { client } from './baseAxios';
import { PostChatsRequest } from './requests';
import {
  DeleteChatsResponse,
  GetChatByIdResponse,
  GetChatsResponse,
  GetMessagesResponse,
  PostChatsResponse,
} from './responses';

export const getChatById = async (
  chatId: string,
): Promise<GetChatByIdResponse> => {
  const resp = await client.get<GetChatByIdResponse>(`/chats/${chatId}`);
  return resp.data;
};

export const postChats = async ({
  defaultModel,
  message,
}: PostChatsRequest): Promise<PostChatsResponse> => {
  const resp = await client.post<PostChatsResponse>('/chats', {
    defaultModel,
    message,
  });
  return resp.data;
};

export const deleteChats = async (
  chatId: string,
): Promise<DeleteChatsResponse> => {
  const resp = await client.delete<DeleteChatsResponse>(`/chats/${chatId}`);
  return resp.data;
};

export const getChats = async (): Promise<GetChatsResponse> => {
  const resp = await client.get<GetChatsResponse>('/chats');
  return resp.data;
};

export const getMessages = async (
  chatId: string,
): Promise<GetMessagesResponse> => {
  const resp = await client.get<GetMessagesResponse>(
    `/chats/${chatId}/messages`,
  );
  return resp.data;
};

export const postMessages = async (chatId: string, text: string) => {
  await client.post(`/chats/${chatId}/messages`, { text: text });
};

export const postGenerateChatTitle = async (chatId: string) => {
  await client.post(`/chats/${chatId}/generate-title`);
};
