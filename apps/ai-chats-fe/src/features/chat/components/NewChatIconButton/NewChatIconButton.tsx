import { useNavigate } from 'react-router-dom';

import { IconButton } from 'components';
import { ChatIcon } from 'components/icons';
import { useSidebarContext } from 'features/sidebar';

import styles from './NewChatIconButton.module.css';

export const NewChatIconButton: React.FC = () => {
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
