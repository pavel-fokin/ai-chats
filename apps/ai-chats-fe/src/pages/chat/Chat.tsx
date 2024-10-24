import { useParams } from 'react-router-dom';

import { Main, PageLayout } from 'components/layout';
import {
  ChatHeader,
  InputMessage,
  MessagesList,
  ModelResponseMessage,
} from 'features/chat/components';
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
      <ChatHeader chatId={chatId} />
      <Main>
        <div className={styles.scrollable}>
          <section className={styles.messagesList}>
            <MessagesList messages={messages.data ?? []} />
            <ModelResponseMessage />
          </section>
        </div>
        <section className={styles.inputMessage}>
          <InputMessage onSendMessage={handleSendMessage} />
        </section>
      </Main>
    </PageLayout>
  );
};
