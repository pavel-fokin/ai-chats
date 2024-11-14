import * as NavigationMenu from '@radix-ui/react-navigation-menu';

import { useGetChats } from 'hooks';

import { Link } from '../Link';

export const ChatsList = () => {
  const chats = useGetChats();

  return (
    <div>
      {chats.data?.map((chat) => (
        <NavigationMenu.Item key={chat.id}>
          <Link to={`/app/chats/${chat.id}`}>{chat.title}</Link>
        </NavigationMenu.Item>
      ))}
    </div>
  );
};
