import styles from './page.module.css';

interface PageProps {
  children: React.ReactNode;
}

// Page layout component.
export const Page = ({ children }: PageProps): JSX.Element => {
  return <div className={styles.Root}>{children}</div>;
};
