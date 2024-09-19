import { IconButton } from 'components';
import { HamburgerMenuIcon } from 'components/icons';
import { useSidebarContext } from 'features/sidebar';

import styles from './OpenSidebarButton.module.css';

export const OpenSidebarButton: React.FC = () => {
  const { toggleSidebar } = useSidebarContext();

  return (
    <IconButton
      className={styles.OpenSidebarButton}
      variant="ghost"
      size="3"
      m="2"
      highContrast
      onClick={toggleSidebar}
    >
      <HamburgerMenuIcon size="24" weight="light" />
    </IconButton>
  );
};
OpenSidebarButton.displayName = 'OpenSidebarButton';
