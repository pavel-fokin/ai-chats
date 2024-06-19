import { Box, Flex, Heading } from '@radix-ui/themes';

import {
  HamburgerMenuButton, NewChatIconButton
} from 'components';
import { Header, PageLayout } from 'components/layout';

export const Home: React.FC = () => {
  return (
    <PageLayout>
      <Header>
        <HamburgerMenuButton />
        <Heading as="h2" size="3" weight="regular">
          AI Chats
        </Heading>
        <NewChatIconButton />
      </Header>
      <Flex direction="column" align="center" justify="center" flexGrow="1">
        <Box>
          <Heading as="h2">There is nothing here 📭 😶</Heading>
        </Box>
      </Flex>
    </PageLayout>
  );
};
