import styles from './header.module.css';

interface HeaderProps {
  children: React.ReactNode;
}

export const Header = ({ children }: HeaderProps): JSX.Element => {
  return <header className={styles.Header}>{children}</header>;
};
