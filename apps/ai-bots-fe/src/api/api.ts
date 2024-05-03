import { Chat } from 'types';

export type Message = {
  sender: string;
  text: string;
};

const CreateChat = async () => {
  const resp = await fetch('/api/chats', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({}),
  });
  return await resp.json();
}

type GetChatsResponse = {
  data: {
    chats: Chat[];
  };
};

const GetChats = async () => {
  const resp = await fetch('/api/chats');
  const payload: GetChatsResponse = await resp.json();
  return payload.data.chats || [];
}

const SendMessage = async (chatID: string, msg: Message) => {
  const resp = await fetch(`/api/chats/${chatID}/messages`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text: msg.text }),
  });
  return await resp.json();
}

export { GetChats, CreateChat, SendMessage };