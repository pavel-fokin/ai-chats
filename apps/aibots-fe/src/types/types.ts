export type Chat = {
  id: string;
  title: string;
  createdAt: string;
};

export type Message = {
  sender: string;
  text: string;
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
