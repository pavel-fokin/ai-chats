import { Box, Flex, Heading } from '@radix-ui/themes';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import {
  HamburgerMenuButton,
  InputMessage,
  NewChatIconButton,
} from 'components';
import { Header, PageLayout } from 'components/layout';
import { useCreateChat, useOllamaModels } from 'hooks';
import { OllamaModel } from 'types';

import { ModelsList } from './components/ModelsList';

export const Startpage: React.FC = () => {
  const navigate = useNavigate();
  const [selectedModel, setSelectedModel] = useState<OllamaModel | null>(null);
  const createChat = useCreateChat();
  const ollamaModels = useOllamaModels();

  const handleSend = async (msg: { text: string }) => {
    createChat.mutate(msg.text, {
      onSuccess: ({ data }) => {
        navigate(`/app/chats/${data.chat.id}`);
      },
    });
  };

  return (
    <PageLayout>
      <Header>
        <HamburgerMenuButton />
        <Heading as="h2" size="3" weight="regular" />
        <NewChatIconButton />
      </Header>
      <Flex direction="column" height="100%" width="100%">
        <Flex direction="column" align="center" justify="center" flexGrow="1">
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
