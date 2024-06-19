import { OllamaModel } from 'types';

import { fetchData, postData, deleteData } from './base';

type GetOllamaResponse = {
  data: {
    models: OllamaModel[];
  };
};

export const fetchOllamaModels = async (): Promise<GetOllamaResponse> => {
  return await fetchData<GetOllamaResponse>('/api/ollama-models');
};

export const postOllamaModels = async (
  modelId: string,
): Promise<OllamaModel> => {
  return await postData<OllamaModel>('/api/ollama-models', { modelId });
};

export const deleteOllamaModels = async (
  modelId: string,
): Promise<OllamaModel> => {
  return await deleteData<OllamaModel>('/api/ollama-models');
};
