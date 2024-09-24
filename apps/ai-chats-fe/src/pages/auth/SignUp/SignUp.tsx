import { useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';

import { Container, Flex, Heading, Link, Text } from '@radix-ui/themes';

import {
  useAuthContext,
  UserCredentialsForm,
  UserCredentials,
} from 'features/auth';
import { useSignUp } from 'hooks';

export const SignUp = () => {
  const navigate = useNavigate();
  const signup = useSignUp();
  const { setIsAuthenticated } = useAuthContext();
  const [error, setError] = useState<string | null>(null);

  const onSubmit: SubmitHandler<UserCredentials> = ({ username, password }) => {
    signup.mutate(
      { username, password },
      {
        onSuccess: () => {
          setIsAuthenticated(true);
          navigate('/app');
        },
        onError: (error: any) => {
          if (error.response.data.errors) {
            setError(error.response.data.errors[0].message);
          }
        },
      },
    );
  };

  return (
    <Container size="1" m="2">
      <Flex direction="column" gap="4">
        <Heading as="h2" size="8">
          Sign up
        </Heading>
        {signup.isError && <Text color="tomato">{error}</Text>}
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
