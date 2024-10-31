import { useSidebarContext } from 'features/sidebar';

import styles from './aside.module.css';

interface AsideProps {
  children: React.ReactNode;
}

export const Aside = ({ children }: AsideProps): JSX.Element => {
  const { isOpen } = useSidebarContext();

  let asideStyles = styles.Aside;
  if (isOpen) {
    asideStyles += ` ${styles.AsideOpen}`;
  }

  return <aside className={asideStyles}>{children}</aside>;
};
