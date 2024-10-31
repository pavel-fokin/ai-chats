import { Box, Flex, Heading, IconButton, TextField } from '@radix-ui/themes';

import { DownloadIcon } from 'components/icons';
import { Aside,Header, Page } from 'components/layout';
import { NewChatIconButton } from 'features/chat';
import { OllamaModelsList, OllamaStatus } from 'features/ollama/components';
import { OpenSidebarButton, Sidebar } from 'features/sidebar';
import { usePullOllamaModel } from 'hooks';

export const OllamaSettings: React.FC = () => {
  const pullModel = usePullOllamaModel();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const inputElement = e.currentTarget.elements[0] as HTMLInputElement;
    pullModel.mutate(inputElement.value);
  };

  return (
    <Page>
      <Aside>
        <Sidebar />
      </Aside>
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
    </Page>
  );
};
