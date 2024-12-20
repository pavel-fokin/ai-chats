import styles from './only-mobile.module.css';

interface OnlyMobileProps {
  children: React.ReactNode;
}

/**
 * OnlyMobile component
 *
 * This component is used to render its children only on mobile devices.
 * It applies a CSS class that ensures the content is displayed only
 * when the viewport width is 768 pixels or less.
 */
export const OnlyMobile = ({ children }: OnlyMobileProps): JSX.Element => {
  return <div className={styles.onlyMobile}>{children}</div>;
};
