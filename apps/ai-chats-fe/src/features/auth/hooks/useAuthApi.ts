import { useMutation } from '@tanstack/react-query';

import { postLogIn, postSignUp } from 'api';
import { UserCredentials } from 'features/auth';

export const useLogIn = () => {
  return useMutation({
    mutationFn: ({ username, password }: UserCredentials) =>
      postLogIn(username, password),
    onSuccess: (response) => {
      const accessToken = response.data?.accessToken;
      if (accessToken) {
        localStorage.setItem('accessToken', accessToken);
      }
    },
  });
};

export const useSignUp = () => {
  return useMutation({
    mutationFn: ({ username, password }: UserCredentials) =>
      postSignUp(username, password),
    onSuccess: (response) => {
      const accessToken = response.data?.accessToken;
      if (accessToken) {
        localStorage.setItem('accessToken', accessToken);
      }
    },
  });
};
