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

export type MessageChunk = {
  id: string;
  sender: string;
  text: string;
  done: boolean;
};

export enum EventTypes {
  MESSAGE_ADDED = 'message_added',
  MESSAGE_CHUNK_RECEIVED = 'message_chunk_received',
}

export type OllamaModel = {
  name: string;
  tag: string;
};

export type UserCredentials = {
  username: string;
  password: string;
};
