import { useChatContext } from '../../hooks';

import { Message } from '../message';

/**
 * Model response message component.
 * @returns {JSX.Element | null} - The model response message component.
 */
export const ModelResponseMessage = (): JSX.Element | null => {
  const { modelResponse } = useChatContext();

  if (!modelResponse) {
    return null;
  }

  return <Message sender={modelResponse.sender} text={modelResponse.text} />;
};
