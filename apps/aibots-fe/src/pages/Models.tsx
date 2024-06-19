import {
  Box,
  Button,
  Flex,
  Heading,
  IconButton,
  Text,
  TextField,
} from '@radix-ui/themes';

import { HamburgerMenuButton, NewChatIconButton } from 'components';
import { Header, PageLayout } from 'components/layout';
import { DownloadIcon } from 'components/ui/icons';
import { useOllamaModels } from 'hooks';

export const Settings: React.FC = () => {
  const { data: models } = useOllamaModels();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log('submit');
  };

  const handleDelete = () => {
    console.log('delete');
  };

  return (
    <PageLayout>
      <Header>
        <HamburgerMenuButton />
        <Heading as="h2" size="3" weight="regular">
          Models
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
                <TextField.Root size="3" placeholder="Pull a model..." />
              </Box>
              <IconButton size="3" highContrast>
                <DownloadIcon size={16} />
              </IconButton>
            </Flex>
          </form>

          {models?.map((model) => (
            <Flex direction="column" gap="2" key={model.id} width="100%">
              <Heading as="h2">
                {model.name}:{model.tag}
              </Heading>
              <Text>
                Meta Llama 3: The most capable openly available LLM to date 8B.
              </Text>
              <Flex align="center" justify="end" flexGrow="1" mt="4" gap="4">
                <Button size="2" variant="soft" onClick={handleDelete}>
                  Delete
                </Button>
              </Flex>
            </Flex>
          ))}
        </Flex>
      </Box>
    </PageLayout>
  );
};
