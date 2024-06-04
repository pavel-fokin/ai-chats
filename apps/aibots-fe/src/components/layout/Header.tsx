import styles from './Header.module.css';

type HeaderProps = {
    children?: React.ReactNode;
}

export const Header: React.FC<HeaderProps> = ({ children }) => {
    return (
        <header className={styles.Header}>
            {children}
        </header>
    )
}