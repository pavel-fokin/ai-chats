import React, { useCallback, useState } from 'react';

import { Button, Select } from '@/components/ui';
import { usePullOllamaModel } from '@/hooks';

import classes from './ollama-model-card.module.css';

interface OllamaModelCardProps {
  modelName: string;
  description: string;
  tags: string[];
}

interface ModelPullButtonProps {
  onPull: () => void;
}

interface ModelTagSelectorProps {
  tags: string[];
  onSelect: (tag: string) => void;
}

// Renders a card displaying an Ollama model with options to select tags and pull the model.
const OllamaModelCard = ({
  modelName,
  description,
  tags,
}: OllamaModelCardProps) => {
  const pullModel = usePullOllamaModel();
  const [selectedTag, setSelectedTag] = useState(tags[0] || '');

  const handlePullModel = useCallback(() => {
    pullModel.mutate(`${modelName}:${selectedTag}`);
  }, [pullModel, modelName, selectedTag]);

  return (
    <article className={classes.ollamaModelCard__container}>
      <h2 className={classes.ollamaModelCard__title}>{modelName}</h2>
      <p>{description}</p>
      <div className={classes.ollamaModelCard__actions}>
        <ModelTagSelector tags={tags} onSelect={setSelectedTag} />
        <ModelPullButton onPull={handlePullModel} />
      </div>
    </article>
  );
};

const ModelTagSelector = React.memo(
  ({ tags, onSelect }: ModelTagSelectorProps) => {
    return (
      <Select.Root defaultValue={tags[0]} onValueChange={onSelect}>
        <Select.Trigger aria-label="Select a tag" />
        <Select.Content highContrast>
          {tags.map((tag, index) => (
            <Select.Item key={`${tag}-${index}`} value={tag}>
              {tag}
            </Select.Item>
          ))}
        </Select.Content>
      </Select.Root>
    );
  }
);

const ModelPullButton = React.memo(({ onPull }: ModelPullButtonProps) => {
  return (
    <Button
      aria-label="Pull model"
      variant="soft"
      size="2"
      highContrast
      onClick={onPull}
    >
      Pull
    </Button>
  );
});

export { OllamaModelCard };
