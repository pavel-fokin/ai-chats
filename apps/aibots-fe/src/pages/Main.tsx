import { useState } from 'react';
import { Outlet } from 'react-router-dom';

import Hamburger from 'hamburger-react';
import { Heading, IconButton, DropdownMenu, Flex, Box, Button } from '@radix-ui/themes';
import { IconMessagePlus, IconPencilPlus, IconTrash } from '@tabler/icons-react';

import { Navbar } from 'components';

import styles from './Main.module.css';

export function Main() {
    const [isOpen, setIsOpen] = useState(false);

    let asideStyles = styles.Aside;
    if (isOpen) {
        asideStyles += ` ${styles.AsideOpen}`;
    }

    return (
        <div className={styles.Root}>
            <header className={styles.Header}>

                <Flex direction="row" align="center" gap="2" justify="between">
                    <div className={styles.MobileOnly}>
                        <Hamburger toggled={isOpen} toggle={() => setIsOpen(!isOpen)} />
                    </div>
                    <DropdownMenu.Root>
                        <DropdownMenu.Trigger>
                            <Button variant="ghost" size="4" highContrast style={{overflow: "hidden", textOverflow: "ellipsis"}}>
                                Chat
                                <DropdownMenu.TriggerIcon />
                            </Button>
                        </DropdownMenu.Trigger>
                        <DropdownMenu.Content>
                            <DropdownMenu.Item shortcut="">Settings</DropdownMenu.Item>
                            <DropdownMenu.Separator />
                            <DropdownMenu.Item color="tomato">Delete <IconTrash size={16} /></DropdownMenu.Item>
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