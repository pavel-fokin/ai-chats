import { Box, Flex, Heading, Text } from '@radix-ui/themes';
import { useNavigate } from 'react-router-dom';

import {
  HamburgerMenuButton,
  InputMessage,
  NewChatIconButton,
} from 'components';
import { Header, PageLayout } from 'components/layout';
import { useCreateChat } from 'hooks';

export const Startpage: React.FC = () => {
  const navigate = useNavigate();
  const createChat = useCreateChat();

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
          <Box>
            <Text as="p" size="6" weight="bold">
              What are you up to? 🤖
            </Text>
          </Box>
        </Flex>
        <Box style={{ maxWidth: '688px', width: '100%', margin: '0 auto' }}>
          <InputMessage handleSend={handleSend} />
        </Box>
      </Flex>
    </PageLayout>
  );
};
