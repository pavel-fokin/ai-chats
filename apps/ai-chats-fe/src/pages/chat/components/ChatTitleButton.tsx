import { DropdownMenu } from '@radix-ui/themes';

import { useChatMenu } from 'features/chat';

import { useChat } from 'shared/hooks';
import { Button } from 'shared/components';

interface ChatTitleButtonProps {
  chatId: string;
}

export const ChatTitleButton = ({ chatId }: ChatTitleButtonProps) => {
  const chat = useChat(chatId);
  const {isOpen, setIsOpen} = useChatMenu();

  return (
    <Button
      variant="ghost"
      size="3"
      highContrast
      onClick={() => setIsOpen(!isOpen)}
    >
      <span
        style={{
          overflow: 'hidden',
          textOverflow: 'ellipsis',
          whiteSpace: 'nowrap',
          maxWidth: '192px',
        }}
      >
        {chat.data?.title || 'Chat'}
      </span>
      <DropdownMenu.TriggerIcon />
    </Button>
  );
};
