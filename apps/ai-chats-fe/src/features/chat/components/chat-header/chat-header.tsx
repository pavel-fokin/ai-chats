import { OnlyDesktop, OnlyMobile } from '@/components/layout';
import { useGetChat } from '@/hooks';

import { ChatHeaderMobile } from './chat-header-mobile';
import { ChatHeaderDesktop } from './chat-header-desktop';

interface ChatHeaderProps {
  chatId: string;
}

// Chat header component.
export const ChatHeader = ({ chatId }: ChatHeaderProps) => {
  const chat = useGetChat(chatId);

  const title = chat.data?.title || 'Chat';

  return (
    <>
      <OnlyMobile>
        <ChatHeaderMobile chatId={chatId} title={title} />
      </OnlyMobile>
      <OnlyDesktop>
        <ChatHeaderDesktop chatId={chatId} title={title} />
      </OnlyDesktop>
    </>
  );
};
