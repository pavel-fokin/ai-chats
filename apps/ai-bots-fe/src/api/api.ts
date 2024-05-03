import { Chat, Message } from 'types';

type GetChatsResponse = {
  data: {
    chats: Chat[];
  };
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

const CreateChat = async () => {
  const resp = await fetch('/api/chats', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({}),
  });
  return await resp.json();
}

const GetChats = async () => {
  const resp = await fetch('/api/chats');
  const payload: GetChatsResponse = await resp.json();
  return payload.data.chats || [];
}

const fetchMessages = async (chatID: string): Promise<GetMessagesResponse> => {
  const resp = await fetch(`/api/chats/${chatID}/messages`);
  if (!resp.ok) {
    throw new Error('Failed to fetch messages');
  }
  return await resp.json();
}

const SendMessage = async (chatID: string, msg: Message) => {
  const resp = await fetch(`/api/chats/${chatID}/messages`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text: msg.text }),
  });
  return await resp.json();
}

export { GetChats, CreateChat, SendMessage, fetchMessages };