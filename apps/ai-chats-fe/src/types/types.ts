export type MessageChunk = {
  id: string;
  sender: string;
  text: string;
  done: boolean;
};

export enum EventTypes {
  MESSAGE_ADDED = 'messageAdded',
  MESSAGE_CHUNK_RECEIVED = 'messageChunkReceived',
}

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

export type OllamaModel = {
  model: string;
  description: string;
  isPulling: boolean;
};

export type UserCredentials = {
  username: string;
  password: string;
};
