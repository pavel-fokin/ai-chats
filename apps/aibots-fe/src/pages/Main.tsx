import { useContext, useEffect, useState } from 'react';
import { Outlet } from 'react-router-dom';

import { Button, DropdownMenu, Flex, IconButton } from '@radix-ui/themes';
import { IconMessagePlus, IconTrash } from '@tabler/icons-react';
import Hamburger from 'hamburger-react';

import { Navbar } from 'components';
import { ChatContext } from 'contexts';
import { useChats } from 'hooks';
import { Chat } from 'types';

import styles from './Main.module.css';

export function Main() {
    const [isOpen, setIsOpen] = useState(false);
    const [chat, setChat] = useState<Chat>({} as Chat);

    const { chatId } = useContext(ChatContext);
    const { getChatById } = useChats();

    let asideStyles = styles.Aside;
    if (isOpen) {
        asideStyles += ` ${styles.AsideOpen}`;
    }

    const handleDelete = () => {
        console.log('Delete chat', chatId);
    }

    useEffect(() => {
        if (chatId) {
            getChatById(chatId).then((chat) => {
                console.log('Chat:', chat);
                setChat(chat);
            }
            ).catch((error) => {
                console.error('Failed to get chat', error);
            });
        }
    }, [chatId]);

    return (
        <div className={styles.Root}>
            <header className={styles.Header}>
                <Flex direction="row" align="center" gap="2" justify="between">
                    <div className={styles.MobileOnly}>
                        <Hamburger toggled={isOpen} toggle={() => setIsOpen(!isOpen)} />
                    </div>
                    <DropdownMenu.Root>
                        <DropdownMenu.Trigger>
                            <Button
                                variant="ghost"
                                size="4"
                                highContrast
                                style={{ overflow: "hidden", textOverflow: "ellipsis" }}
                            >
                                {chat.title || 'Chat'}
                                <DropdownMenu.TriggerIcon />
                            </Button>
                        </DropdownMenu.Trigger>
                        <DropdownMenu.Content>
                            <DropdownMenu.Item shortcut="">Settings</DropdownMenu.Item>
                            <DropdownMenu.Separator />
                            <DropdownMenu.Item color="tomato" onClick={handleDelete}>Delete <IconTrash size={16} /></DropdownMenu.Item>
                        </DropdownMenu.Content>
                    </DropdownMenu.Root>
                    <IconButton className={styles.MobileOnly} variant="ghost" size="3" m="3" highContrast>
                        <IconMessagePlus size={30} />
                    </IconButton>
                </Flex>
            </header>
            <aside className={asideStyles}>
                <Navbar open={setIsOpen} />
            </aside>
            <main className={styles.Main}>
                <Outlet />
            </main>
        </div>
    )
}