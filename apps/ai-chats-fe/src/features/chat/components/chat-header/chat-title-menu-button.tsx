import { DropdownMenu } from '@radix-ui/themes';

import { Button } from '@/components/ui';

import { useChatMenu } from '../../hooks';

interface ChatTitleMenuButtonProps {
  title: string;
}

export const ChatTitleMenuButton = ({ title }: ChatTitleMenuButtonProps) => {
  const { isOpen, setIsOpen } = useChatMenu();

  return (
    <Button
      aria-label="Open chat menu"
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
        {title}
      </span>
      <DropdownMenu.TriggerIcon />
    </Button>
  );
};
