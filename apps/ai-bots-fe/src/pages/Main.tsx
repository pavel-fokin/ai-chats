import { useState } from 'react';
import { Outlet } from 'react-router-dom';

import { Box, Container, Flex } from '@radix-ui/themes';
import Hamburger from 'hamburger-react'

import { Navbar } from 'components';

import styles from './Main.module.css';

export function Main() {
    const [isOpen, setOpen] = useState(false);

    let asideStyles = styles.Aside;

    if (isOpen) {
        asideStyles += ` ${styles.AsideOpen}`;
    }

    return (
        <div className={styles.Root}>
            <header className={styles.Header}>
                <Hamburger onToggle={() => setOpen(!isOpen)} />
            </header>
            <aside className={asideStyles}>
                <Navbar />
            </aside>
            <main className={styles.Main}>
                <Outlet />
            </main>
        </div>
    )
}