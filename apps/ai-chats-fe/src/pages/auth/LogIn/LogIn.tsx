import { useContext, useState } from 'react';
import { SubmitHandler } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';

import { Container, Flex, Heading, Link, Text } from '@radix-ui/themes';

import { AuthContext } from 'contexts';
import { useLogIn } from 'hooks';
import { UserCredentialsSchema } from 'schemas';
import { UserCredentialsForm } from '../components/UserCredentialsForm';

export const LogIn = () => {
  const navigate = useNavigate();
  const { setIsAuthenticated } = useContext(AuthContext);
  const logIn = useLogIn();
  const [error, setError] = useState<string | null>(null);

  const onSubmit: SubmitHandler<UserCredentialsSchema> = async ({
    username,
    password,
  }) => {
    logIn.mutate(
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
          Log in
        </Heading>
        {logIn.isError && <Text color="tomato">{error}</Text>}
        <UserCredentialsForm
          onSubmit={onSubmit}
          isLoading={logIn.isPending}
          submitButtonText="Log in"
        />
        <Text align="center">
          Don't have an account? <Link href="/app/signup">Create one</Link>
        </Text>
      </Flex>
    </Container>
  );
};
