import { OllamaModel, OllamaModelStatus } from 'types';

import { client } from './baseAxios';

type GetOllamaResponse = {
  data: {
    models: OllamaModel[];
  };
};

export const getOllamaModels = async (
  status: OllamaModelStatus,
): Promise<GetOllamaResponse> => {
  const params = status === OllamaModelStatus.ANY ? {} : { status };
  const resp = await client.get<GetOllamaResponse>('/ollama/models', {
    params,
  });
  return resp.data;
};

export const postOllamaModels = async (model: string) => {
  await client.post('/ollama/models', { model: model });
};

export const deleteOllamaModels = async (model: string) => {
  await client.delete(`/ollama/models/${model}`);
};
