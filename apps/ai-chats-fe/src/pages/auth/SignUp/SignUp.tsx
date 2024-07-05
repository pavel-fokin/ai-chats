import { useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import { SubmitHandler } from 'react-hook-form';

import { Container, Flex, Heading, Link, Text } from '@radix-ui/themes';

import { AuthContext } from 'contexts';
import { UserCredentialsSchema } from 'schemas';

import { UserCredentialsForm } from '../components/UserCredentialsForm';

export const SignUp = () => {
  const navigate = useNavigate();
  const { signup, isLoading } = useContext(AuthContext);

  const onSubmit: SubmitHandler<UserCredentialsSchema> = async ({
    username,
    password,
  }) => {
    const signedUp = await signup(username, password);
    if (signedUp) {
      navigate('/app');
    }
  };

  return (
    <Container size="1" m="2">
      <Flex direction="column" gap="4">
        <Heading as="h2" size="8">
          Sign up
        </Heading>
        <UserCredentialsForm
          onSubmit={onSubmit}
          isLoading={isLoading}
          submitButtonText="Create account"
        />
        <Text align="center">
          Already have an account? <Link href="/app/login">Log in</Link>
        </Text>
      </Flex>
    </Container>
  );
};
