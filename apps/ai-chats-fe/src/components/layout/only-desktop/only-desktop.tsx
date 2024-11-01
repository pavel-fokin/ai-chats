import styles from './only-desktop.module.css';

interface OnlyDesktopProps {
  children: React.ReactNode;
}

/**
 * OnlyDesktop component
 *
 * This component is used to render its children only on desktop devices.
 * It applies a CSS class that ensures the content is displayed only
 * when the viewport width is 768 pixels or greater.
 */
export const OnlyDesktop = ({ children }: OnlyDesktopProps) => {
  return <div className={styles.onlyDesktop}>{children}</div>;
};
