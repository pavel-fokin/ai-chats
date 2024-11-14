export type ModelResponse = {
  text: string;
  sender: string;
};

export enum EventTypes {
  MESSAGE_ADDED = 'messageAdded',
  CHAT_MESSAGE = 'chatMessage',
  MODEL_STREAM_MESSAGE = 'ModelStreamMessageNotification',
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

export type OllamaModelCard = {
  model: string;
  description: string;
  tags: string[];
};

export enum OllamaModelStatus {
  ANY = 'any',
  AVAILABLE = 'available',
  PULLING = 'pulling',
}
