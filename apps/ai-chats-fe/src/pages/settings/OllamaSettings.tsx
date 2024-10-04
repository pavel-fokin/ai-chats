import { Heading } from '@radix-ui/themes';

import { NewChatIconButton } from 'features/chat';
import { OpenSidebarButton } from 'features/sidebar';
import { Header, PageLayout } from 'shared/components/layout';
import { useOllamaModels } from 'shared/hooks';

import { OllamaModel, OllamaStatus, PullOllamaModelDialog } from './components';

import styles from './OllamaSettings.module.css';

export const OllamaSettings = (): JSX.Element => {
  const ollamaModels = useOllamaModels();

  return (
    <PageLayout>
      <Header>
        <OpenSidebarButton />
        <Heading as="h2" size="3" weight="regular">
          Ollama Settings
        </Heading>
        <NewChatIconButton />
      </Header>
      <div className={styles.ollamaSettings__container}>
        <div className={styles.ollamaSettings__bar}>
          <OllamaStatus />
          <PullOllamaModelDialog />
        </div>
        <div className={styles.ollamaSettings__modelsList}>
          {ollamaModels.data?.map((model) => (
            <OllamaModel key={model.model} model={model} />
          ))}
        </div>
      </div>
    </PageLayout>
  );
};
