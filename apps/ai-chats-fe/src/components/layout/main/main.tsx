import styles from './main.module.css';

interface MainProps {
  children: React.ReactNode;
}

// Main layout component.
export const Main = ({ children }: MainProps): JSX.Element => {
  return <main className={styles.Main}>{children}</main>;
};
