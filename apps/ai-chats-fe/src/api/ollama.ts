import { OllamaModel } from 'types';

import { client } from './baseAxios';

type GetOllamaResponse = {
  data: {
    models: OllamaModel[];
  };
};

export const fetchOllamaModels = async (): Promise<GetOllamaResponse> => {
  const resp = await client.get<GetOllamaResponse>('/ollama-models');
  return resp.data;
};

export const postOllamaModels = async (modelName: string) => {
  await client.post('/ollama-models', { model: modelName });
};

export const deleteOllamaModels = async (modelName: string) => {
  await client.delete(`/ollama-models/${modelName}`);
};
