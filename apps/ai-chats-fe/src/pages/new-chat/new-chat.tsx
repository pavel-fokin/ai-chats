import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { Heading } from '@radix-ui/themes';

import { Header, Main } from '@/components/layout';
import { InputMessage, NewChatIconButton } from '@/features/chat';
import { OllamaModelsSelect } from '@/features/ollama/components';
import { OpenSidebarButton } from '@/features/sidebar';
import { useCreateChat, useOllamaModels } from '@/hooks';
import { OllamaModel } from '@/types';

import styles from './new-chat.module.css';

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
      }
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
          <section className={styles.newChat__modelSelect}>
            <Heading as="h2" size="6" weight="bold" mb="4">
              Choose a model 🤖
            </Heading>
            <OllamaModelsSelect
              models={ollamaModels.data || []}
              selectedModel={selectedModel}
              setSelectedModel={setSelectedModel}
            />
          </section>
          <section className={styles.newChat__inputMessage}>
            <InputMessage onSendMessage={handleSendMessage} />
          </section>
      </Main>
    </>
  );
};
