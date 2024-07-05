import { useContext } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';

import { Container, Flex, Heading, Link, Text } from '@radix-ui/themes';

import { AuthContext } from 'contexts';
import { UserCredentialsSchema } from 'schemas';

import { UserCredentialsForm } from '../components/UserCredentialsForm';

export const LogIn = () => {
  const navigate = useNavigate();
  const { login, isLoading } = useContext(AuthContext);

  const onSubmit: SubmitHandler<UserCredentialsSchema> = async ({
    username,
    password,
  }) => {
    const isLoggedIn = await login(username, password);
    if (isLoggedIn) {
      navigate('/app');
    }
  };

  return (
    <Container size="1" m="2">
      <Flex direction="column" gap="4">
        <Heading as="h2" size="8">
          Log in
        </Heading>
        <UserCredentialsForm
          onSubmit={onSubmit}
          isLoading={isLoading}
          submitButtonText="Log in"
        />
        <Text align="center">
          Don't have an account? <Link href="/app/signup">Create one</Link>
        </Text>
      </Flex>
    </Container>
  );
};
