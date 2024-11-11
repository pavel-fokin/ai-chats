import { DropdownMenu } from '@radix-ui/themes';

import { Button } from '@/components/ui';
import { OpenSidebarButton } from '@/features/sidebar/components/OpenSidebarButton';

import { NewChatIconButton } from '../new-chat-button';
import { ChatMenu } from '../chat-menu';

import styles from './chat-header-mobile.module.css';

interface ChatHeaderMobileProps {
  chatId: string;
  title: string;
}

interface ChatTitleAsMenuButtonProps {
  title: string;
}

export const ChatHeaderMobile = ({ chatId, title }: ChatHeaderMobileProps) => {
  return (
    <div className={styles.chatHeaderMobile__container}>
      <OpenSidebarButton />
      <ChatMenu
        chatId={chatId}
        trigger={<ChatTitleAsMenuButton title={title} />}
      />
      <NewChatIconButton />
    </div>
  );
};

const ChatTitleAsMenuButton = ({ title }: ChatTitleAsMenuButtonProps) => {
  return (
    <Button aria-label="Open chat menu" variant="ghost" size="3" highContrast>
      <h1 className={styles.chatHeaderMobile__title}>{title}</h1>
      <DropdownMenu.TriggerIcon />
    </Button>
  );
};
