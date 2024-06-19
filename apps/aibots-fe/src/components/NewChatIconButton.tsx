import { IconButton } from '@radix-ui/themes';
import { ChatIcon } from 'components/ui/icons';

import 'styles/styles.css'

export const NewChatIconButton: React.FC = () => {
  return (
    <IconButton
      className="mobile-only"
      variant="ghost"
      size="3"
      m="2"
      highContrast
    >
      <ChatIcon size="28" weight="light" />
    </IconButton>
  );
};
