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
  addedAt: string;
  updateAt: string;
  deletedAt: string;
  status: OllamaModelStatus;
};

export enum OllamaModelStatus {
  ADDDED = 'added',
  PULLING = 'pulling',
  AVAILABLE = 'available',
  ERROR = 'error',
}

export type UserCredentials = {
  username: string;
  password: string;
};
