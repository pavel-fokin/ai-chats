import { useState } from 'react';

import { Box, Flex, IconButton } from '@radix-ui/themes';

import { SendIcon } from 'components/ui/icons';
import { TextArea } from 'components';

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
        <IconButton size="3" onClick={onSendClick} highContrast>
          <SendIcon size={16} />
        </IconButton>
      </Flex>
    </form>
  );
}

export { InputMessage };
