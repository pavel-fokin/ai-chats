import { useNavigate } from 'react-router-dom';

import { IconButton } from '@/components/ui';
import { ChatIcon } from '@/components/icons';
import { useSidebarContext } from '@/features/sidebar';

import styles from './new-chat-button.module.css';

export const NewChatIconButton = (): JSX.Element => {
  const navigate = useNavigate();
  const { closeSidebar } = useSidebarContext();

  const handleClick = () => {
    navigate('/app/new-chat');
    closeSidebar();
  };

  return (
    <IconButton
      aria-label="New chat"
      className={styles.NewChatIconButton}
      variant="ghost"
      size="3"
      m="2"
      highContrast
      onClick={handleClick}
    >
      <ChatIcon size="24" weight="light" />
    </IconButton>
  );
};
