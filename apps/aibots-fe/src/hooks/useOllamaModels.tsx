import { useQuery } from '@tanstack/react-query';

import { fetchOllamaModels } from 'api';

export const useOllamaModels = () => {
  return useQuery({
    queryKey: ['ollama-models'],
    queryFn: fetchOllamaModels,
    select: (data) => data.data.models,
  });
};
