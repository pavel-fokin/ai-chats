import { OllamaModel } from 'types';
import { fetchData } from './base';

type GetOllamaResponse = {
  data: {
    models: OllamaModel[];
  };
};

export const fetchOllamaModels = async (): Promise<GetOllamaResponse> => {
  return await fetchData<GetOllamaResponse>('/api/ollama-models');
};
