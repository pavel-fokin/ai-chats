import * as NavigationMenu from '@radix-ui/react-navigation-menu';

import styles from './MenuList.module.css';

interface MenuListProps {
  children: React.ReactNode;
}

export const MenuList = ({ children }: MenuListProps) => {
  return (<NavigationMenu.List className={styles.menuList}>
    {children}
  </NavigationMenu.List>);
};
