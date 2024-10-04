import { Dialog } from '@radix-ui/themes';

import { Button, TextInput } from 'shared/components';
import { usePullOllamaModel } from 'shared/hooks';

import styles from './PullOllamaModelDialog.module.css';

export const PullOllamaModelDialog = (): JSX.Element => {
  const pullModel = usePullOllamaModel();

  const handlePullModel = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const inputElement = e.currentTarget.elements[0] as HTMLInputElement;
    pullModel.mutate(inputElement.value);
  };

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button highContrast>Pull model</Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth="450px">
        <Dialog.Title>Pull Ollama model</Dialog.Title>
        <form onSubmit={handlePullModel} style={{ width: '100%' }}>
          <div className={styles.container}>
            <Dialog.Description>
              This will pull an Ollama model from a registry.
            </Dialog.Description>

            <TextInput id="model" size="3" placeholder="Enter model name" />
            <div className={styles.buttons}>
              <Dialog.Close>
                <Button aria-label="Close dialog" variant="ghost">
                  Cancel
                </Button>
              </Dialog.Close>
              <Dialog.Close>
                <Button
                  type="submit"
                  aria-label="Pull ollama model"
                  variant="solid"
                  // onClick={handlePullModel}
                  highContrast
                >
                  Pull model
                </Button>
              </Dialog.Close>
            </div>
          </div>
        </form>
      </Dialog.Content>
    </Dialog.Root>
  );
};
