import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

import { deleteOllamaModels, getOllamaModels, getOllamaModelsLibrary, postOllamaModels } from 'api';
import { OllamaModelStatus } from 'types';

export const useGetOllamaModels = (status?: OllamaModelStatus) => {
  return useQuery({
    queryKey: ['ollama-models', status],
    queryFn: () => getOllamaModels(status ?? OllamaModelStatus.ANY),
    select: (data) => data.data.models,
  });
};

export const useGetOllamaModelsLibrary = () => {
  return useQuery({
    queryKey: ['ollama-models-library'],
    queryFn: () => getOllamaModelsLibrary(),
    select: (data) => data.data.modelCards,
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
