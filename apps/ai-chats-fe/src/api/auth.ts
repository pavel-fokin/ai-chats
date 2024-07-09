import { client } from './baseAxios';

interface Response<T> {
  data?: T;
  errors?: string[];
}

interface AccessToken {
  accessToken: string;
}

export const postLogIn = async (
  username: string,
  password: string
): Promise<Response<AccessToken>> => {
  const resp = await client.post<Response<AccessToken>>('/auth/login', {
    username,
    password,
  });
  return resp.data;
};

export const postSignUp = async (
  username: string,
  password: string
): Promise<Response<AccessToken>> => {
  const resp = await client.post<Response<AccessToken>>('/auth/signup', {
    username,
    password,
  });
  return resp.data;
};
