import { useState } from 'react';

import { Box, Flex, IconButton } from '@radix-ui/themes';

import { TextArea, Tooltip } from 'components';
import { SendIcon } from 'components/ui/icons';

type InputMessageProps = {
  handleSend: (text: string) => void;
};

function InputMessage({ handleSend }: InputMessageProps) {
  const [inputText, setInputText] = useState<string>('');

  const onInputChangeArea = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInputText(event.target.value);
  };

  const onSendClick = async (e: React.FormEvent) => {
    e.preventDefault();
    if (inputText.trim() !== '') {
      handleSend(inputText);
      setInputText('');
    }
  };

  return (
    <form role="form" onSubmit={onSendClick}>
      <Flex
        gap="2"
        justify="center"
        p={{
          initial: '2',
          sm: '4',
        }}
      >
        <Box flexGrow="1">
          <TextArea
            value={inputText}
            onChange={onInputChangeArea}
            placeholder="Type a message"
          />
        </Box>
        <Tooltip content="Send a message" side="top">
          <IconButton size="3" onClick={onSendClick} highContrast>
            <SendIcon size={16} />
          </IconButton>
        </Tooltip>
      </Flex>
    </form>
  );
}

export { InputMessage };
