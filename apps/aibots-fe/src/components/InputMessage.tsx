import { useState } from 'react';

import { Box, Flex, IconButton } from '@radix-ui/themes';
import { PaperPlaneTilt } from '@phosphor-icons/react';

import { TextArea } from 'components/TextArea';
import * as types from 'types';

type InputMessageProps = {
  handleSend: (msg: types.Message) => void;
};

function InputMessage({ handleSend }: InputMessageProps) {
  const [inputMessage, setInputMessage] = useState<types.Message>({
    sender: '',
    text: '',
  });
  const onInputChangeArea = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInputMessage({ ...inputMessage, text: event.target.value });
  };

  const onSendClick = async (e: React.FormEvent) => {
    e.preventDefault();
    if (inputMessage.text.trim() !== '') {
      handleSend({ sender: 'User', text: inputMessage.text });
      setInputMessage({ sender: '', text: '' });
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
            value={inputMessage.text}
            onChange={onInputChangeArea}
            placeholder="Type a message"
          />
        </Box>
        <IconButton size="3" onClick={onSendClick} highContrast>
          <PaperPlaneTilt size={16} weight="light" />
        </IconButton>
      </Flex>
    </form>
  );
}

export { InputMessage };
