import { useContext } from 'react';
import { useNavigate } from 'react-router-dom';

import { IconButton } from '@radix-ui/themes';

import { SidebarContext } from 'contexts';
import { ChatIcon } from 'components/ui/icons';

import 'styles/styles.css';

export const NewChatIconButton: React.FC = () => {
  const navigate = useNavigate();
  const { closeSidebar } = useContext(SidebarContext);

  const handleClick = () => {
    navigate('/app');
    closeSidebar();
  };

  return (
    <IconButton
      aria-label="New chat"
      className="mobile-only"
      variant="ghost"
      size="3"
      m="2"
      highContrast
      onClick={handleClick}
    >
      <ChatIcon size="28" weight="light" />
    </IconButton>
  );
};
