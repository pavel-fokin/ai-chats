import { useSidebarContext } from 'features/sidebar';
import { IconButton } from 'shared/components';
import { CloseIcon } from 'shared/components/icons';

import styles from './CloseSidebarButton.module.css';

export const CloseSidebarButton: React.FC = () => {
  const { closeSidebar } = useSidebarContext();

  return (
    <IconButton
      className={styles.CloseSidebarButton}
      variant="ghost"
      onClick={closeSidebar}
    >
      <CloseIcon size={24} />
    </IconButton>
  );
};
