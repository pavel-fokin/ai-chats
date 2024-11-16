import { Box, Flex, Heading, TextField } from '@radix-ui/themes';

import { SearchIcon } from '@/components/icons';
import { Header, Main } from '@/components/layout';
import { NewChatIconButton } from '@/features/chat';
import { OllamaLibrary, OllamaModelsList } from '@/features/ollama/components';
import { OpenSidebarButton } from '@/features/sidebar';
import { useGetOllamaModelsLibrary } from '@/hooks';

export const OllamaLibraryPage = () => {
  const modelCards = useGetOllamaModelsLibrary();

  return (
    <>
      <Header>
        <OpenSidebarButton />
        <Heading as="h1" size="3" weight="regular">
          Ollama Library
        </Heading>
        <NewChatIconButton />
      </Header>
      <Main>
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
            <Flex
              gap="3"
              align="center"
              pb={{
                initial: '2',
                sm: '4',
              }}
              width="100%"
            >
              <Box flexGrow="1">
                <TextField.Root
                  id="model"
                  size="3"
                  placeholder="Search model"
                >
                <TextField.Slot>
                  <SearchIcon />
                  </TextField.Slot>
                </TextField.Root>
              </Box>
            </Flex>
            <OllamaModelsList />
            <OllamaLibrary modelCards={modelCards.data ?? []} />
          </Flex>
        </Box>
      </Main>
    </>
  );
};
