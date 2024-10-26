export type ModelResponse = Message;

export enum EventTypes {
  MESSAGE_ADDED = 'messageAdded',
  MESSAGE_CHUNK_RECEIVED = 'messageChunkReceived',
  CHAT_MESSAGE = 'chatMessage',
  OLLAMA_MODEL_PULL_PROGRESS = 'ollamaModelPullProgress',
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
  status: OllamaModelStatus;
};

export enum OllamaModelStatus {
  ANY = 'any',
  AVAILABLE = 'available',
  PULLING = 'pulling',
}
