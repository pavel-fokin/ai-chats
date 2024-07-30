import { Box, Button, Flex, Heading, Text } from '@radix-ui/themes';

export const LandingPage: React.FC = () => {
  return (
    <Flex direction="column" height="100vh" gap="2">
      <Flex p="4" direction="column" align="end">
        <header>
          <nav>
            <Flex gap="6" align="baseline">
              <Button asChild variant="ghost" size="3" highContrast>
                <a href="/app/login">Log in</a>
              </Button>
              <Button asChild size="3" highContrast>
                <a href="/app/signup">Sign up</a>
              </Button>
            </Flex>
          </nav>
        </header>
      </Flex>
      <Flex direction="column" align="center" justify="center" flexGrow="1">
        <main>
          <Text align="center">
            <Heading
              size={{
                initial: '7',
                lg: '8',
              }}
            >
              Create and Manage Your AI Chats [*_*]
            </Heading>
            <Heading
              as="h2"
              size={{
                initial: '5',
                lg: '7',
              }}
            >
              Boost your daily tasks with our easy-to-use platform!
            </Heading>
          </Text>
        </main>
      </Flex>
      <Box p="4">
        <footer>
          <Text as="p" align="center">
            {new Date().getFullYear()} AI Chats
          </Text>
        </footer>
      </Box>
    </Flex>
  );
};
