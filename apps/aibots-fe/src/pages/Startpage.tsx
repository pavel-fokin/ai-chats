import { Box, Flex, Heading, Text } from '@radix-ui/themes';

import { HamburgerMenuButton, NewChatIconButton, InputMessage } from 'components';
import { Header, PageLayout } from 'components/layout';

export const Startpage: React.FC = () => {
  const handleSend = async (msg: { text: string }) => {
    console.log(msg);
  }

  return (
    <PageLayout>
      <Header>
        <HamburgerMenuButton />
        <Heading as="h2" size="3" weight="regular">

        </Heading>
        <NewChatIconButton />
      </Header>
      <Flex direction="column" height="100%" width="100%">
        <Flex direction="column" align="center" justify="center" flexGrow="1">
          <Box>
            <Text as="p" size="6" weight="bold">What are you up to? 🤖 </Text>
          </Box>
        </Flex>
        <Box style={{ maxWidth: '688px', width: '100%', margin: '0 auto' }}>
          <InputMessage handleSend={handleSend} />
        </Box>
      </Flex>
    </PageLayout>
  );
};
