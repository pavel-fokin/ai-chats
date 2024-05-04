
import { Outlet } from 'react-router-dom';

import { Box, Container, Flex } from '@radix-ui/themes';

import { NavBar } from 'components';

export function Main() {
    return (
        <Flex direction="row" minHeight="100vh">
            <Box display={{ initial: "none", sm: "inline" }}>
                <NavBar />
            </Box>
            <main style={{ display: "flex", width: "100%" }}>
                <Container asChild size="2" mx="2" my={{ initial: "4", sm: "6" }}>
                    <Flex direction="column-reverse">
                        <Outlet />
                    </Flex>
                </Container>
            </main>
        </Flex>
    )
}