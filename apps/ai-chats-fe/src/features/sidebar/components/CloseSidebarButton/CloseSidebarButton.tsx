import { useContext } from 'react';

import { IconButton } from 'components/IconButton';
import { CloseIcon } from 'components/icons';
import { SidebarContext } from 'contexts';

import styles from './CloseSidebarButton.module.css';

export const CloseSidebarButton: React.FC = () => {
  const { closeSidebar } = useContext(SidebarContext);

  return (
    <IconButton className={styles.CloseSidebarButton} variant="ghost" onClick={closeSidebar}>
      <CloseIcon size={24} />
    </IconButton>
  );
};
