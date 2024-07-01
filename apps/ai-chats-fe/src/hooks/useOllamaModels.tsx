import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

import { fetchOllamaModels, postOllamaModels, deleteOllamaModels } from 'api';

export const useOllamaModels = () => {
  return useQuery({
    queryKey: ['ollama-models'],
    queryFn: fetchOllamaModels,
    select: (data) => data.data.models,
  });
};

export const usePullOllamaModel = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (modelName: string) => postOllamaModels(modelName),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['ollama-models'],
      });
    },
  });
};

export const useDeleteOllamaModel = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (modelName: string) => deleteOllamaModels(modelName),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['ollama-models'],
      });
    },
  });
};
