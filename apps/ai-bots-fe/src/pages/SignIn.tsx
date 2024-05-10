import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";

import { Button, Container, Flex, Heading, Link, Text, TextField } from "@radix-ui/themes";

import { AuthContext } from "contexts";
import { useAuth } from "hooks";

export const SignIn = () => {
    const navigate = useNavigate();

    const { signIn } = useAuth();
    const { setIsAuthenticated } = useContext(AuthContext);

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const onSignIn = async () => {
        const response = await signIn(username, password);
        if (response) {
            setIsAuthenticated(true);
            navigate('/app');
        }
    }

    return (
        <Container size="1" m="2">
            <Flex direction="column" gap="4">
                <Heading as="h2" size="8">Sign In</Heading>
                <TextField.Root
                    name="username"
                    autoComplete="off"
                    size="3"
                    placeholder="Your username"
                    onChange={e => { setUsername(e.target.value) }}
                />
                <TextField.Root
                    name="password"
                    size="3"
                    type="password"
                    placeholder="Your password"
                    onChange={e => { setPassword(e.target.value) }}
                />
                <Button size="4" onClick={onSignIn} highContrast>Sign In</Button>
                <Text align="center">
                    Don't have an account?  <Link href="/app/signup">Sign Up</Link>
                </Text>
            </Flex>
        </Container>
    );
};

export default SignIn;