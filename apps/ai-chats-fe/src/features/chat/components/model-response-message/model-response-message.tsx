import { useChatContext } from '../../hooks';

import { Message } from '../message';

/**
 * Model response message component.
 * @returns {JSX.Element | null} - The model response message component.
 */
export const ModelResponseMessage = (): JSX.Element | null => {
  const { messageChunk } = useChatContext();

  if (!messageChunk) {
    return null;
  }

  return <Message sender={messageChunk.sender} text={messageChunk.text} />;
};
