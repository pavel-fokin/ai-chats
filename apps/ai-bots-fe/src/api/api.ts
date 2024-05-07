import { Chat, Message } from 'types';

type PostChatsResponse = {
  data: {
    id: string;
  };
};

type GetChatsResponse = {
  data: {
    chats: Chat[];
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

const postChats = async (): Promise<PostChatsResponse> => {
  const resp = await fetch('/api/chats', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({}),
  });
  if (!resp.ok) {
    throw new Error('Failed to create chat');
  }
  return await resp.json();
}

const fetchChats = async () => {
  const resp = await fetch('/api/chats');
  if (!resp.ok) {
    throw new Error('Failed to fetch chats');
  }
  const payload: GetChatsResponse = await resp.json();
  return payload.data.chats || [];
}

const fetchMessages = async (chatId: string): Promise<GetMessagesResponse> => {
  const resp = await fetch(`/api/chats/${chatId}/messages`);
  if (!resp.ok) {
    throw new Error('Failed to fetch messages');
  }
  return await resp.json();
}

const postMessages = async (chatId: string, msg: Message): Promise<PostMessagesResponse> => {
  const resp = await fetch(`/api/chats/${chatId}/messages`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text: msg.text }),
  });
  if (!resp.ok) {
    throw new Error('Failed to send message');
  }
  return await resp.json();
}

export { fetchChats, postChats, postMessages, fetchMessages };