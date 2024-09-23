import { Heading } from '@radix-ui/themes';

import { NewChatIconButton } from 'components';
import { NewChat } from 'features/chat';
import { OpenSidebarButton } from 'features/sidebar';
import { Header, PageLayout } from 'layout';

export const NewChatPage: React.FC = () => {
  return (
    <PageLayout>
      <Header>
        <OpenSidebarButton />
        <Heading as="h2" size="3" weight="regular">
          Start a new chat
        </Heading>
        <NewChatIconButton />
      </Header>
      <NewChat />
    </PageLayout>
  );
};
