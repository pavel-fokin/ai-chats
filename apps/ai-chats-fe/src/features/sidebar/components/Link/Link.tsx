import { useContext } from 'react';
import { NavLink, useLocation } from 'react-router-dom';

import * as NavigationMenu from '@radix-ui/react-navigation-menu';

import { SidebarContext } from 'contexts';

import styles from './Link.module.css';

interface LinkProps {
  to: string;
  children: React.ReactNode;
}

export const Link: React.FC<LinkProps> = ({ to, children, ...props }) => {
  const { pathname } = useLocation();
  const isActive = to === pathname;

  const { toggleSidebar } = useContext(SidebarContext);

  const classNames = isActive
    ? `${styles.NavigationMenuLink} ${styles.NavigationMenuLinkActive}`
    : styles.NavigationMenuLink;

  return (
    <NavigationMenu.Link asChild active={isActive}>
      <NavLink
        to={to}
        className={classNames}
        onClick={toggleSidebar}
        {...props}
      >
        {children}
      </NavLink>
    </NavigationMenu.Link>
  );
};
