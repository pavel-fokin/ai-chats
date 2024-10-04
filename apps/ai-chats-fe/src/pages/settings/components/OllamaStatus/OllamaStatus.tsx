import { Badge } from '@radix-ui/themes'

import styles from './OllamaStatus.module.css';

export const OllamaStatus = (): JSX.Element => {
    return (
        <div className={styles.ollamaStatus__container}>
            <h3>Ollama Server</h3>
            <Badge color="jade">Online</Badge>
        </div>
    );
};