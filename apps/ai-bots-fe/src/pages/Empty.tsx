import { Flex, Box, Heading } from "@radix-ui/themes";


export function Empty() {
    return (
            <Flex direction="column" align="center" justify="center" flexGrow="1">
                <Box>
                    <Heading as="h2">AI Bots</Heading>
                </Box>
            </Flex>
    );
}