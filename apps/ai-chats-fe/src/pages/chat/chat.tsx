import { useParams } from 'react-router-dom';

import { Header, Main, PageLayout } from 'components/layout';
import {
  ChatHeader,
  InputMessage,
  MessagesList,
  ModelResponseMessage,
} from 'features/chat/components';
import { ChatMenuProvider } from 'features/chat/contexts';
import { useChatLogic } from 'features/chat/hooks';

import styles from './chat.module.css';

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
      <Header>
        <ChatMenuProvider>
          <ChatHeader chatId={chatId} />
        </ChatMenuProvider>
      </Header>
      <Main>
        <section className={styles.scrollable}>
          <div className={styles.messagesList}>
            <MessagesList messages={messages.data ?? []} />
            <ModelResponseMessage />
          </div>
        </section>
        <section className={styles.inputMessage}>
          <InputMessage onSendMessage={handleSendMessage} />
        </section>
      </Main>
    </PageLayout>
  );
};
