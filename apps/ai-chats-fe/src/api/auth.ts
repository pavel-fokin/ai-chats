import { clientUnauthed as client } from './baseAxios';

import { AccessToken, Response } from './responses';

export const postLogIn = async (
  username: string,
  password: string,
): Promise<Response<AccessToken>> => {
  const resp = await client.post<Response<AccessToken>>('/auth/login', {
    username,
    password,
  });
  return resp.data;
};

export const postSignUp = async (
  username: string,
  password: string,
): Promise<Response<AccessToken>> => {
  const resp = await client.post<Response<AccessToken>>('/auth/signup', {
    username,
    password,
  });
  return resp.data;
};
