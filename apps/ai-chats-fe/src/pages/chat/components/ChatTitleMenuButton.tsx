import { DropdownMenu } from '@radix-ui/themes';

import { useChatMenu } from 'features/chat';

import { Button } from 'shared/components';

interface ChatTitleButtonProps {
  title: string;
}

export const ChatTitleMenuButton = ({ title }: ChatTitleButtonProps) => {
  const { isOpen, setIsOpen } = useChatMenu();

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
        {title || 'Chat'}
      </span>
      <DropdownMenu.TriggerIcon />
    </Button>
  );
};
