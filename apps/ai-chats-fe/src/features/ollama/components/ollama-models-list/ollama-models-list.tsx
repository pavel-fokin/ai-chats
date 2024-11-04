import { useOllamaModels } from '@/hooks';
import { OllamaModelStatus } from '@/types';

import { OllamaModelAvailable } from './ollama-model-available';
import { OllamaModelPulling } from './ollama-model-pulling';

import classes from './ollama-models-list.module.css';

export const OllamaModelsList = () => {
  const ollamaModelsAvailable = useOllamaModels(OllamaModelStatus.AVAILABLE);
  const ollamaModelsPulling = useOllamaModels(OllamaModelStatus.PULLING);

  return (
    <div className={classes.OllamaModelsList}>
      {ollamaModelsPulling.data?.map((model) => (
        <OllamaModelPulling key={model.model} model={model} />
      ))}
      {ollamaModelsAvailable.data?.map((model) => (
        <OllamaModelAvailable key={model.model} model={model} />
      ))}
    </div>
  );
};
