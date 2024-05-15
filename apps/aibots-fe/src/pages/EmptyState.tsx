import { Flex, Box, Heading } from "@radix-ui/themes";

export function EmptyState() {
    return (
        <Flex direction="column" align="center" justify="center" flexGrow="1">
            <Box>
                <Heading as="h2">There is nothing here 📭 😶</Heading>
            </Box>
        </Flex>
    );
}