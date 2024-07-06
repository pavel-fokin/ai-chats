import { useMutation } from '@tanstack/react-query';

import { postLogIn, postSignUp } from 'api';
import { UserCredentials } from 'types';

export const useLogIn = () => {
  return useMutation({
    mutationFn: ({ username, password }: UserCredentials) =>
      postLogIn(username, password),
    onSuccess: (data) => {
      localStorage.setItem('accessToken', data.accessToken);
    },
  });
};

export const useSignUp = () => {
  return useMutation({
    mutationFn: ({ username, password }: UserCredentials) =>
      postSignUp(username, password),
    onSuccess: (data) => {
      localStorage.setItem('accessToken', data.accessToken);
    },
  });
};