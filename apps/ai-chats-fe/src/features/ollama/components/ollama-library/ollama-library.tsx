import * as types from '@/types';

import { OllamaModelCard } from './ollama-model-card';

interface OllamaLibraryProps {
  modelCards: types.OllamaModelCard[];
}

export const OllamaLibrary = ({
  modelCards,
}: OllamaLibraryProps): JSX.Element => {
  return (
    <>
      {modelCards.map((modelCard) => (
        <OllamaModelCard key={modelCard.modelName} {...modelCard} />
      ))}
    </>
  );
};
