import { useContext, useEffect, useState } from 'react';
import { Outlet } from 'react-router-dom';

import { Button, DropdownMenu, Flex, IconButton } from '@radix-ui/themes';
import {
    ChatText as ChatIcon,
    List as HamburgerMenuIcon2,
    X as CloseIcon,
    Trash as DeleteIcon,
    SlidersHorizontal as SettingsIcon
} from "@phosphor-icons/react";

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
                        <IconButton variant="ghost" size="3" m="2" highContrast onClick={() => setIsOpen(!isOpen)}>
                            {isOpen ? <CloseIcon size="28" weight="light" /> : <HamburgerMenuIcon2 size="28" weight="light" />}
                        </IconButton>
                    </div>
                    <DropdownMenu.Root >
                        <DropdownMenu.Trigger>
                            <Button
                                variant="ghost"
                                size="3"
                                highContrast
                                style={{ overflow: "hidden", textOverflow: "ellipsis" }}
                            >
                                {chat.title || 'Chat'}
                                <DropdownMenu.TriggerIcon />
                            </Button>
                        </DropdownMenu.Trigger>
                        <DropdownMenu.Content style={{ minWidth: "128px" }} >
                            <DropdownMenu.Item shortcut="">
                                <Flex direction="row" align="center" justify="between" width="100%">
                                    Configure <SettingsIcon size="16" />
                                </Flex>
                            </DropdownMenu.Item>
                            <DropdownMenu.Separator />
                            <DropdownMenu.Item color="tomato" onClick={handleDelete}>
                                <Flex direction="row" align="center" justify="between" width="100%">
                                    Delete <DeleteIcon size="16" />
                                </Flex>
                            </DropdownMenu.Item>
                        </DropdownMenu.Content>
                    </DropdownMenu.Root>
                    <IconButton className={styles.MobileOnly} variant="ghost" size="3" m="2" highContrast>
                        <ChatIcon size="28" weight="light" />
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