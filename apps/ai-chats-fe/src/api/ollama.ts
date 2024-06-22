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

export const postOllamaModels = async (
  modelId: string,
): Promise<OllamaModel> => {
  return await doPost<OllamaModel>('/api/ollama-models', { modelId });
};

export const deleteOllamaModels = async (): Promise<OllamaModel> => {
  return await doDelete<OllamaModel>('/api/ollama-models');
};
