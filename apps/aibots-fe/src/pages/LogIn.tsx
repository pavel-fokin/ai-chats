import { zodResolver } from '@hookform/resolvers/zod';
import { useContext } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';

import {
  Button,
  Container,
  Flex,
  Heading,
  Link,
  Text,
  TextField,
} from '@radix-ui/themes';

import { AuthContext } from 'contexts';
import { userCredentialsSchema, UserCredentialsSchema } from 'schemas';

export const LogIn = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UserCredentialsSchema>({
    resolver: zodResolver(userCredentialsSchema),
  });

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
      <form role="form" onSubmit={handleSubmit(onSubmit)}>
        <Flex direction="column" gap="4">
          <Heading as="h2" size="8">
            Log in
          </Heading>
          <TextField.Root
            autoComplete="off"
            size="3"
            placeholder="Your username"
            {...register('username')}
          />
          {errors.username && (
            <Text color="tomato">{errors.username.message?.toString()}</Text>
          )}
          <TextField.Root
            size="3"
            type="password"
            placeholder="Your password"
            {...register('password')}
          />
          {errors.password && (
            <Text color="tomato">{errors.password.message?.toString()}</Text>
          )}
          <Button loading={isLoading} size="4" highContrast type="submit">
            Log in
          </Button>
          <Text align="center">
            Don't have an account? <Link href="/app/signup">Create one</Link>
          </Text>
        </Flex>
      </form>
    </Container>
  );
};

export default LogIn;
