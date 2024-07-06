import { useContext } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';

import { Container, Flex, Heading, Link, Text } from '@radix-ui/themes';

import { AuthContext } from 'contexts';
import { useSignUp } from 'hooks';
import { UserCredentialsSchema } from 'schemas';

import { UserCredentialsForm } from '../components/UserCredentialsForm';

export const SignUp = () => {
  const navigate = useNavigate();
  const signup = useSignUp();
  const { setIsAuthenticated } = useContext(AuthContext);

  const onSubmit: SubmitHandler<UserCredentialsSchema> = ({
    username,
    password,
  }) => {
    signup.mutate({ username, password }, {
      onSuccess: () => {
        setIsAuthenticated(true);
        navigate('/app');
      },
    });
  };

  return (
    <Container size="1" m="2">
      <Flex direction="column" gap="4">
        <Heading as="h2" size="8">
          Sign up
        </Heading>
        <UserCredentialsForm
          onSubmit={onSubmit}
          isLoading={signup.isPending}
          submitButtonText="Create an account"
        />
        <Text align="center">
          Already have an account? <Link href="/app/login">Log in</Link>
        </Text>
      </Flex>
    </Container>
  );
};
