import { Flex, Box, Heading } from "@radix-ui/themes";


export function Empty() {
    return (
        <Flex>
            <Flex>
                <Box flexGrow="1">
                    <Heading>Empty</Heading>
                </Box>
            </Flex>
        </Flex>

    );
}