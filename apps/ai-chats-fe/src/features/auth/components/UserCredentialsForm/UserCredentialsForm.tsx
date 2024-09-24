import { SubmitHandler, useForm } from 'react-hook-form';

import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';

import { Flex, Text, TextField } from '@radix-ui/themes';

import { Button } from 'shared/components';

const UserCredentialsSchema = z.object({
  username: z.string().min(1, 'Username is required'),
  password: z.string().min(6, 'Password must be at least 6 characters'),
});

export type UserCredentials = z.infer<typeof UserCredentialsSchema>;

interface UserCredentialsFormProps {
  onSubmit: SubmitHandler<UserCredentials>;
  isLoading: boolean;
  submitButtonText: string;
}

export const UserCredentialsForm: React.FC<UserCredentialsFormProps> = ({
  onSubmit,
  isLoading,
  submitButtonText,
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UserCredentials>({
    resolver: zodResolver(UserCredentialsSchema),
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
