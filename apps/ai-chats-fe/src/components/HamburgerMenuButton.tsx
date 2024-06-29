import { useContext } from 'react';

import { IconButton } from '@radix-ui/themes';

import { HamburgerMenuIcon, CloseIcon } from 'components';
import { SidebarContext } from 'contexts';

import 'styles/styles.css';

export const HamburgerMenuButton: React.FC = () => {
  const { isOpen, toggleSidebar } = useContext(SidebarContext);

  return (
    <IconButton
      className="mobile-only"
      variant="ghost"
      size="3"
      m="2"
      highContrast
      onClick={toggleSidebar}
    >
      {isOpen ? (
        <CloseIcon size="28" weight="light" />
      ) : (
        <HamburgerMenuIcon size="28" weight="light" />
      )}
    </IconButton>
  );
};