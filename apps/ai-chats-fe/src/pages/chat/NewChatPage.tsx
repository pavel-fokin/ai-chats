import { Box, Code, Flex, Heading } from '@radix-ui/themes';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { HamburgerMenuButton, NewChatIconButton } from 'components';
import { useCreateChat, useOllamaModels } from 'hooks';
import { Header, PageLayout } from 'layout';
import { OllamaModel } from 'types';

import { InputMessage, ModelsList } from './components';

export const NewChatPage: React.FC = () => {
  const navigate = useNavigate();
  const [selectedModel, setSelectedModel] = useState<OllamaModel | null>(null);
  const createChat = useCreateChat();
  const ollamaModels = useOllamaModels();

  const handleSend = async (text: string) => {
    createChat.mutate(
      {
        defaultModel: `${selectedModel?.model}`,
        message: text,
      },
      {
        onSuccess: ({ data }) => {
          navigate(`/app/chats/${data.chat.id}`);
        },
      }
    );
  };

  return (
    <PageLayout>
      <Header>
        <HamburgerMenuButton />
        <Heading as="h2" size="3" weight="regular">
          Start a new chat
        </Heading>
        <NewChatIconButton />
      </Header>
      <Flex direction="column" height="100%" width="100%">
        <Flex direction="column" align="center" justify="center" flexGrow="1">
          <Box mb="4">
            <Heading as="h2" size="6" weight="bold">
              Choose a model <Code variant="ghost">[*_*]</Code>
            </Heading>
          </Box>
          <ModelsList
            models={ollamaModels.data || []}
            selectedModel={selectedModel}
            setSelectedModel={setSelectedModel}
          />
        </Flex>
        <Box style={{ maxWidth: '688px', width: '100%', margin: '0 auto' }}>
          <InputMessage handleSend={handleSend} />
        </Box>
      </Flex>
    </PageLayout>
  );
};
