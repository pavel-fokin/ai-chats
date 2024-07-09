import { zodResolver } from '@hookform/resolvers/zod';
import { SubmitHandler, useForm } from 'react-hook-form';

import { userCredentialsSchema, UserCredentialsSchema } from 'schemas';

import { Button, Flex, Text, TextField } from '@radix-ui/themes';

interface Props {
  onSubmit: SubmitHandler<UserCredentialsSchema>;
  isLoading: boolean;
  submitButtonText: string;
}

export const UserCredentialsForm: React.FC<Props> = ({
  onSubmit,
  isLoading,
  submitButtonText,
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UserCredentialsSchema>({
    resolver: zodResolver(userCredentialsSchema),
  });

  return (
    <form role="form" onSubmit={handleSubmit(onSubmit)}>
      <Flex direction="column" gap="4">
        {errors.username && (
          <Text color="tomato">{errors.username.message?.toString()}</Text>
        )}
        <TextField.Root
          autoComplete="off"
          size="3"
          placeholder="Your username"
          {...register('username')}
        />
        {errors.password && (
          <Text color="tomato">{errors.password.message?.toString()}</Text>
        )}
        <TextField.Root
          size="3"
          type="password"
          placeholder="Your password"
          {...register('password')}
        />
        <Button loading={isLoading} size="4" highContrast>
          {submitButtonText}
        </Button>
      </Flex>
    </form>
  );
};
