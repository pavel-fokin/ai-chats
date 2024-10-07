import styles from './Main.module.css';

interface MainProps {
  children: React.ReactNode;
}

export const Main: React.FC<MainProps> = ({ children }) => {
  return <main className={styles.Main}>{children}</main>;
};
