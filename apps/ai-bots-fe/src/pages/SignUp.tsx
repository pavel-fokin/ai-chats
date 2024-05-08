import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";

import { Button, Container, Flex, Heading, Link, Text, TextField } from "@radix-ui/themes";

import { AuthContext } from "contexts";
import { useAuth } from "hooks";

export const SignUp = () => {
    const { setIsAuthenticated } = useContext(AuthContext);
    const { signUp } = useAuth();

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const navigate = useNavigate();

    const onSignUp = async () => {
        if (await signUp(username, password)) {
            setIsAuthenticated(true);
            navigate('/app');
        }
    }

    return (
        <Container size="1" m="2">
            <Flex direction="column" gap="4">
                <Heading as="h2" size="8">Sign Up</Heading>
                <TextField.Root name="username" autoComplete="off" size="3" placeholder="Your username" onChange={e => { setUsername(e.target.value) }} />
                <TextField.Root name="password" size="3" type="password" placeholder="Your password" onChange={e => { setPassword(e.target.value) }} />
                <Button size="4" onClick={onSignUp} highContrast>Create an account</Button>
                <Text align="center">
                    Already have an account?  <Link href="/app/signin">Sign In</Link>
                </Text>
            </Flex>
        </Container>
    );
};

export default SignUp;