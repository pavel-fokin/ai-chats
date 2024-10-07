import { Flex, Heading, Progress, Text } from '@radix-ui/themes';

import { useOllamaModelPullingEvents } from 'features/ollama';
import { OllamaModel } from 'types';

interface OllamaModelPullingProps {
  model: OllamaModel;
}

export const OllamaModelPulling = ({ model }: OllamaModelPullingProps) => {
  const { progress } = useOllamaModelPullingEvents(model);

  return (
    <Flex direction="column" gap="2" key={`${model.model}`} width="100%">
      <Heading as="h2" size="3">{`${model.model}`}</Heading>
      <Text>{model.description}</Text>
      <Flex direction="column" gap="0">
        {progress && progress.total > 0 && (
          <>
            <Text size="1">{progress.status}</Text>
            <Progress
              mt="4"
              mb="8"
              value={Math.ceil((progress.completed / progress.total) * 100)}
              highContrast
            />
          </>
        )}
      </Flex>
    </Flex>
  );
};
