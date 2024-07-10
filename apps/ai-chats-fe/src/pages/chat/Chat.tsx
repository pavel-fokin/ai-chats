import { useParams } from 'react-router-dom';

import { Box, Flex } from '@radix-ui/themes';

import {
  ChatMenu,
  HamburgerMenuButton,
  InputMessage,
  Message,
  NewChatIconButton,
} from 'components';
import { Header, PageLayout } from 'components/layout';
import { useChatEvents, useMessages, useSendMessage } from 'hooks';
import * as types from 'types';

export const Chat = () => {
  const { chatId } = useParams<{ chatId: string }>();

  const messages = useMessages(chatId);
  const sendMessage = useSendMessage(chatId!);
  const { messageChunk } = useChatEvents(chatId!);

  const handleSend = async (msg: types.Message) => {
    sendMessage.mutate({ sender: 'human', text: msg.text });
  };

  return (
    <PageLayout>
      <Header>
        <HamburgerMenuButton />
        <ChatMenu chatId={chatId!} />
        <NewChatIconButton />
      </Header>
      <Flex direction="column" height="100%" width="100%">
        <Box flexGrow="1" style={{ overflow: 'scroll' }}>
          <Box height="100%" style={{ maxWidth: '688px', margin: '0 auto' }}>
            <Flex flexGrow="1" justify="end" direction="column" gap="2">
              {messages.data?.length !== 0 && (
                <Box style={{ height: '64px' }}></Box>
              )}
              {messages.data?.map((message, index) => (
                <Message
                  key={index}
                  sender={message.sender}
                  text={message.text}
                />
              ))}
              {messageChunk.text && (
                <Message
                  sender={messageChunk.sender}
                  text={messageChunk.text}
                />
              )}
            </Flex>
          </Box>
        </Box>
        <Box style={{ maxWidth: '688px', width: '100%', margin: '0 auto' }}>
          <InputMessage handleSend={handleSend} />
        </Box>
      </Flex>
    </PageLayout>
  );
};