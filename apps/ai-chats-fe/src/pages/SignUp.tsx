import { useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import { SubmitHandler, useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';

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

export const SignUp = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UserCredentialsSchema>({
    resolver: zodResolver(userCredentialsSchema),
  });

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
      <form role="form" onSubmit={handleSubmit(onSubmit)}>
        <Flex direction="column" gap="4">
          <Heading as="h2" size="8">
            Sign up
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
          <Button loading={isLoading} size="4" highContrast>
            Create an account
          </Button>
          <Text align="center">
            Already have an account? <Link href="/app/login">Log in</Link>
          </Text>
        </Flex>
      </form>
    </Container>
  );
};

export default SignUp;
