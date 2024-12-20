import { useState } from 'react';

import { SendIcon } from '@/components/icons';
import { IconButton, Tooltip, TextArea } from '@/components/ui';

import styles from './input-message.module.css';

interface InputMessageProps {
  onSendMessage: (message: string) => void;
}

export const InputMessage = ({
  onSendMessage,
}: InputMessageProps): JSX.Element => {
  const [message, setMessage] = useState<string>('');

  const handleInputChange = ({
    target,
  }: React.ChangeEvent<HTMLTextAreaElement>) => {
    setMessage(target.value);
  };

  const handleSendMessage = () => {
    if (message.trim() === '') {
      return;
    }

    onSendMessage(message);
    setMessage('');
  };

  return (
    <div className={styles.inputMessageContainer}>
      <TextArea
        aria-label="Type a message here"
        onChange={handleInputChange}
        onEnterPress={handleSendMessage}
        placeholder="Type a message"
        value={message}
      />
      <Tooltip content="Send a message" side="top">
        <IconButton
          aria-label="Send a message"
          size="3"
          onClick={handleSendMessage}
          highContrast
        >
          <SendIcon size={16} />
        </IconButton>
      </Tooltip>
    </div>
  );
};
