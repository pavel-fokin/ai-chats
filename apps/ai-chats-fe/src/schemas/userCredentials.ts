import { z } from 'zod';

export const userCredentialsSchema = z.object({
  username: z.string().min(1, 'Username is required'),
  password: z.string().min(6, 'Password must be at least 6 characters'),
});

export type UserCredentialsSchema = z.infer<typeof userCredentialsSchema>;
