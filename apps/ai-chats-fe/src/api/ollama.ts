import { OllamaModel } from 'types';

import { doGet, doPost, doDelete } from './base';

type GetOllamaResponse = {
  data: {
    models: OllamaModel[];
  };
};

export const fetchOllamaModels = async (): Promise<GetOllamaResponse> => {
  return await doGet<GetOllamaResponse>('/api/ollama-models');
};

export const postOllamaModels = async (modelName: string) => {
  return await doPost<undefined>('/api/ollama-models', { model: modelName });
};

export const deleteOllamaModels = async (modelName: string) => {
  return await doDelete<undefined>(`/api/ollama-models/${modelName}`);
};
