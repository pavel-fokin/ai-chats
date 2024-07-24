import styles from './Main.module.css';

type MainProps = {
  children: React.ReactNode;
};

export const Main: React.FC<MainProps> = ({ children }) => {
  return <main className={styles.Main}>{children}</main>;
};
