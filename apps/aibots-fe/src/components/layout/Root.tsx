import styles from './Root.module.css';

type RootProps = {
    children: React.ReactNode;
}

export const Root: React.FC<RootProps> = ({children}) => {
    return (
        <div className={styles.Root}>
            {children}
        </div>
    );
}