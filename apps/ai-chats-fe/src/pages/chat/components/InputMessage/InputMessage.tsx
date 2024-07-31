import { useState } from 'react';

import { Box, Flex, IconButton } from '@radix-ui/themes';

import { TextArea, Tooltip } from 'components';
import { SendIcon } from 'components/icons';

interface InputMessageProps {
  handleSend: (text: string) => void;
}

export const InputMessage: React.FC<InputMessageProps> = ({ handleSend }) => {
  const [inputText, setInputText] = useState<string>('');

  const onInputChangeArea = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInputText(event.target.value);
  };

  const onSendClick = async () => {
    if (inputText.trim() !== '') {
      handleSend(inputText);
      setInputText('');
    }
  };

  return (
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
          onChange={onInputChangeArea}
          onEnterPress={onSendClick}
          placeholder="Type a message"
          value={inputText}
        />
      </Box>
      <Tooltip content="Send a message" side="top">
        <IconButton size="3" onClick={onSendClick} highContrast>
          <SendIcon size={16} />
        </IconButton>
      </Tooltip>
    </Flex>
  );
};
