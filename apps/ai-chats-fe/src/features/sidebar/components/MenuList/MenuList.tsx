import * as NavigationMenu from '@radix-ui/react-navigation-menu';

import styles from './MenuList.module.css';

interface MenuListProps {
  ariaLabel: string;
  children: React.ReactNode;
}

export const MenuList = ({ ariaLabel, children }: MenuListProps) => {
  return (
    <NavigationMenu.List aria-label={ariaLabel} className={styles.menuList}>
      {children}
    </NavigationMenu.List>
  );
};
