import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { Box, Code, Flex, Heading } from '@radix-ui/themes';

import { Header, Main } from '@/components/layout';
import { InputMessage, NewChatIconButton } from '@/features/chat';
import { OllamaModelsSelect } from '@/features/ollama/components';
import { OpenSidebarButton } from '@/features/sidebar';
import { useCreateChat, useOllamaModels } from '@/hooks';
import { OllamaModel } from '@/types';

export const NewChat = () => {
  const [selectedModel, setSelectedModel] = useState<OllamaModel | null>(null);
  const createChat = useCreateChat();
  const navigate = useNavigate();
  const ollamaModels = useOllamaModels();

  const handleSendMessage = async (text: string) => {
    if (!selectedModel) return;

    createChat.mutate(
      {
        defaultModel: `${selectedModel.model}`,
        message: text,
      },
      {
        onSuccess: ({ data }) => {
          navigate(`/app/chats/${data.chat.id}`);
        },
      },
    );
  };

  return (
    <>
      <Header>
        <OpenSidebarButton />
        <Heading as="h2" size="3" weight="regular">
          Start a new chat
        </Heading>
        <NewChatIconButton />
      </Header>
      <Main>
        <Flex direction="column" height="100%" width="100%">
          <Flex direction="column" align="center" justify="center" flexGrow="1">
            <Box mb="4">
              <Heading as="h2" size="6" weight="bold">
                Choose a model 🤖
              </Heading>
            </Box>
            <OllamaModelsSelect
              models={ollamaModels.data || []}
              selectedModel={selectedModel}
              setSelectedModel={setSelectedModel}
            />
          </Flex>
          <Box style={{ maxWidth: '688px', width: '100%', margin: '0 auto' }}>
            <InputMessage onSendMessage={handleSendMessage} />
          </Box>
        </Flex>
      </Main>
    </>
  );
};
