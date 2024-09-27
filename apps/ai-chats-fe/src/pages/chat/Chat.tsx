import { useParams } from 'react-router-dom';

import { InputMessage } from 'features/chat';
import { PageLayout, Main } from 'shared/components/layout';

import { ChatHeader, MessageChunk, MessagesList } from './components';
import { useChatLogic } from './hooks/useChatLogic';

import styles from './Chat.module.css';

export const Chat = () => {
  const { chatId } = useParams();
  const { messages, sendMessage } = useChatLogic(chatId ?? '');

  if (!chatId) {
    return null;
  }

  const handleSendMessage = (text: string) => {
    sendMessage.mutate(text);
  };

  if (messages.isLoading) {
    return <div>Loading...</div>;
  }

  if (messages.isError) {
    return <div>Error: {messages.error.message}</div>;
  }

  return (
    <PageLayout>
      <ChatHeader chatId={chatId} />
      <Main>
        <div className={styles.Chat__scrollable}>
          <section className={styles.Chat__messagesList}>
            <MessagesList messages={messages.data ?? []} />
            <MessageChunk />
          </section>
        </div>
        <section className={styles.Chat__inputMessage}>
          <InputMessage handleSend={handleSendMessage} />
        </section>
      </Main>
    </PageLayout>
  );
};
