import { Box, Flex, Heading, IconButton, TextField } from '@radix-ui/themes';

import { NewChatIconButton } from 'features/chat';
import { OpenSidebarButton } from 'features/sidebar';
import { DownloadIcon } from 'components/icons';
import { Header, PageLayout } from 'components/layout';
import { usePullOllamaModel } from 'hooks';

import { OllamaStatus, OllamaModelsList } from './components';

export const OllamaSettings: React.FC = () => {
  const pullModel = usePullOllamaModel();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const inputElement = e.currentTarget.elements[0] as HTMLInputElement;
    pullModel.mutate(inputElement.value);
  };

  return (
    <PageLayout>
      <Header>
        <OpenSidebarButton />
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
          <OllamaStatus />
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
                  placeholder="Enter model name"
                />
              </Box>
              <IconButton size="3" highContrast loading={pullModel.isPending}>
                <DownloadIcon size={16} />
              </IconButton>
            </Flex>
          </form>
          <OllamaModelsList />
        </Flex>
      </Box>
    </PageLayout>
  );
};
