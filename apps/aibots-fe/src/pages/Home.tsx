import { Box, Flex, Heading } from "@radix-ui/themes";

import { PageLayout } from "components/layout";

export const Home: React.FC = () => {
    return (
        <PageLayout>
            <Flex direction="column" align="center" justify="center" flexGrow="1">
                <Box>
                    <Heading as="h2">There is nothing here 📭 😶</Heading>
                </Box>
            </Flex>
        </PageLayout>
    );
}