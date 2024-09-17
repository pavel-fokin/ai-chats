import { useSidebarContext } from 'features/sidebar';

import styles from './Aside.module.css';

type AsideProps = {
  children: React.ReactNode;
};

export const Aside: React.FC<AsideProps> = ({ children }) => {
  const { isOpen } = useSidebarContext();

  let asideStyles = styles.Aside;
  if (isOpen) {
    asideStyles += ` ${styles.AsideOpen}`;
  }

  return <aside className={asideStyles}>{children}</aside>;
};
