import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

import { deleteOllamaModels, getOllamaModels, postOllamaModels } from 'api';

export const useOllamaModels = (onlyPulling?: boolean) => {
  return useQuery({
    queryKey: ['ollama-models', onlyPulling],
    queryFn: () => getOllamaModels(onlyPulling ?? false),
    select: (data) => data.data.models,
  });
};

export const usePullOllamaModel = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (model: string) => postOllamaModels(model),
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
    mutationFn: (model: string) => deleteOllamaModels(model),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['ollama-models'],
      });
    },
  });
};

export const useInvalidateOllamaModels = () => {
  const queryClient = useQueryClient();

  return () => {
    queryClient.invalidateQueries({
      queryKey: ['ollama-models'],
    });
  };
};
