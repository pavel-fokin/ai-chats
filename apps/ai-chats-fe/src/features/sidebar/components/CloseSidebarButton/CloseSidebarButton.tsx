import { IconButton } from 'components';
import { CloseIcon } from 'components/icons';
import { useSidebarContext } from 'features/sidebar';

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
