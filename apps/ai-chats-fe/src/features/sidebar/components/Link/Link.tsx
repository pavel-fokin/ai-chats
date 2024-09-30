import { NavLink, useLocation } from 'react-router-dom';

import * as NavigationMenu from '@radix-ui/react-navigation-menu';

import { useSidebarContext } from '../../hooks/useSidebarContext';

import styles from './Link.module.css';

interface LinkProps {
  to: string;
  children: React.ReactNode;
  onClick?: () => void;
}

export const Link = ({ to, children, ...props }: LinkProps): JSX.Element => {
  const { pathname } = useLocation();
  const isActive = to === pathname;

  const { toggleSidebar } = useSidebarContext();

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
