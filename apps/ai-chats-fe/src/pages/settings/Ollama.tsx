import { useEffect } from 'react';

import { Box, Flex, Heading, IconButton, TextField } from '@radix-ui/themes';

import { HamburgerMenuButton, NewChatIconButton } from 'components';
import { Header, PageLayout } from 'components/layout';
import { DownloadIcon } from 'components/ui/icons';
import { useModelsLibrary, useOllamaModels, usePullOllamaModel } from 'hooks';

import { ModelCard, OllamaModel } from './components';

export const OllamaSettings: React.FC = () => {
  const modelsLibrary = useModelsLibrary();
  const ollamaModels = useOllamaModels();
  const pullModel = usePullOllamaModel();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const inputElement = e.currentTarget.elements[0] as HTMLInputElement;
    pullModel.mutate(inputElement.value);
  };

  return (
    <PageLayout>
      <Header>
        <HamburgerMenuButton />
        <Heading as="h2" size="3" weight="regular">
          Ollama Settings
        </Heading>
        <NewChatIconButton />
      </Header>
      <Box style={{ maxWidth: '688px', width: '100%', margin: '0 auto' }}>
        <Flex
          direction="column"
          align="center"
          justify="start"
          gap="3"
          flexGrow="1"
          mt="9"
          px="4"
        >
          <form onSubmit={handleSubmit} style={{ width: '100%' }}>
            <Flex
              gap="3"
              align="center"
              pb={{
                initial: '2',
                sm: '4',
              }}
            >
              <Box flexGrow="1">
                <TextField.Root
                  id="model"
                  size="3"
                  placeholder="Pull a model..."
                />
              </Box>
              <IconButton size="3" highContrast loading={pullModel.isPending}>
                <DownloadIcon size={16} />
              </IconButton>
            </Flex>
          </form>
          <Box width="100%" pb="8">
            <Box pb="4">
              <Heading as="h3" size="3" align="left" color="gray">
                Downloaded Models
              </Heading>
            </Box>
            {ollamaModels.data?.map((model) => (
              <OllamaModel key={model.model} model={model} />
            ))}
          </Box>
          <Box width="100%">
            <Box pb="4">
              <Heading as="h3" size="3" align="left" color="gray">
                Models Library
              </Heading>
            </Box>
            {modelsLibrary.data?.map((model) => (
              <ModelCard key={model.model} model={model} />
            ))}
          </Box>
        </Flex>
      </Box>
    </PageLayout>
  );
};
