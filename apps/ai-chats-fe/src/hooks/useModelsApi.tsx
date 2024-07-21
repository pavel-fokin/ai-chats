import { useQuery } from '@tanstack/react-query';

import { getModelsLibrary } from 'api/models';

export const useModelsLibrary = () =>
  useQuery({
    queryKey: ['models-library'],
    queryFn: getModelsLibrary,
    select: (data) => data.data.modelCards,
  });
