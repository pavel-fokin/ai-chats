import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";

import { Button, Container, Flex, Heading, Link, Text, TextField } from "@radix-ui/themes";

import { AuthContext } from "contexts";

export const SignUp = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const navigate = useNavigate();

    const { signup, isLoading } = useContext(AuthContext);

    const onSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        const signedUp = await signup(username, password);
        if (signedUp) {
            navigate('/app');
        }
    }

    return (
        <Container size="1" m="2">
            <form role="form" onSubmit={onSubmit}>
                <Flex direction="column" gap="4">
                    <Heading as="h2" size="8">Sign up</Heading>
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
                    <Button loading={isLoading} size="4" highContrast>Create an account</Button>
                    <Text align="center">
                        Already have an account?  <Link href="/app/login">Log in</Link>
                    </Text>
                </Flex>
            </form>
        </Container>
    );
};

export default SignUp;