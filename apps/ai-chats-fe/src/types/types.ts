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
  model: string;
  description: string;
};

export type UserCredentials = {
  username: string;
  password: string;
};
