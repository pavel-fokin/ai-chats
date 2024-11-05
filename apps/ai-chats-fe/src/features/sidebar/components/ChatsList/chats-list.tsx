import * as NavigationMenu from '@radix-ui/react-navigation-menu';

import { useChats } from 'hooks';

import { Link } from '../Link';

// import styles from './chats-list.module.css';

export const ChatsList = () => {
  const chats = useChats();

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
