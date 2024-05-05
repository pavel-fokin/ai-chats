import { useState } from 'react';
import { Outlet } from 'react-router-dom';

import { Box, Container, Flex } from '@radix-ui/themes';
import Hamburger from 'hamburger-react'

import { Navbar } from 'components';

export function Main() {
    const [isOpen, setOpen] = useState(false);

    return (
        <Flex direction={{ initial: "column", md: "row" }} minHeight="100vh">
            <Box asChild display={{ initial: "block", sm: "none" }} style={{ zIndex: "2" }}>
                <header>
                    <Hamburger onToggle={() => setOpen(!isOpen)} />
                </header>
            </Box>
            <Box
                display={{ initial: isOpen ? "block" : "none", sm: "block" }}
                height={{
                    initial: "100vh",
                    sm: "auto",
                }}
                width={{
                    initial: "100vw",
                    sm: "300px",
                }}
                position={{
                    initial: "fixed",
                    sm: "static",
                }}
                style={{
                    padding: "48px 0",
                    backgroundColor: "var(--color-background)",
                    zIndex: "1",
                }}

            >
                <Navbar />
            </Box>
            <Flex
                asChild
                direction={{
                    initial: "column",
                    sm: "row",
                }}
                width="100%"
                flexGrow={{
                    initial: "1",
                    sm: "0",
                }}
            >
                <main>
                    <Container asChild size="2" mx="2" my={{ initial: "4", sm: "6" }}>
                        <Flex direction="column-reverse">
                            <Outlet />
                        </Flex>
                    </Container>
                </main>
            </Flex>

        </Flex>
    )
}