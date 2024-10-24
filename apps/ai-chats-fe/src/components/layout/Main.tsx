import styles from './Main.module.css';

interface MainProps {
  children: React.ReactNode;
}

/**
 * Main layout component.
 * @param {React.ReactNode} children - The children to be rendered inside the main layout
 * @returns {React.ReactElement} - The main layout component
 */
export const Main: React.FC<MainProps> = ({ children }) => {
  return <main className={styles.Main}>{children}</main>;
};
