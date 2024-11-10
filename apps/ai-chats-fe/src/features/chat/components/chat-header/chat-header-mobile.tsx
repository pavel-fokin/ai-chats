import { DropdownMenu } from '@radix-ui/themes';

import { Button } from '@/components/ui';
import { OpenSidebarButton } from '@/features/sidebar/components/OpenSidebarButton';

import { ChatMenu } from '../chat-menu/chat-menu';
import { ChatMenuProvider } from '../chat-menu/chat-menu-provider';
import { useChatMenu } from '../chat-menu/use-chat-menu';
import { NewChatIconButton } from '../new-chat-button';

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
      <ChatMenuProvider>
        <ChatMenu
          chatId={chatId}
          trigger={<ChatTitleAsMenuButton title={title} />}
        />
      </ChatMenuProvider>
      <NewChatIconButton />
    </div>
  );
};

const ChatTitleAsMenuButton = ({ title }: ChatTitleAsMenuButtonProps) => {
  const { isOpen, setIsOpen } = useChatMenu();

  return (
    <Button
      aria-label="Open chat menu"
      variant="ghost"
      size="3"
      highContrast
      onClick={() => setIsOpen(!isOpen)}
    >
      <h1 className={styles.chatHeaderMobile__title}>{title}</h1>
      <DropdownMenu.TriggerIcon />
    </Button>
  );
};
