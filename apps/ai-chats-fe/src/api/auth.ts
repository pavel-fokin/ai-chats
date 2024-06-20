type SignInResponse = {
  accessToken: string;
};

type SignUpResponse = {
  accessToken: string;
};

export const postLogIn = async (
  username: string,
  password: string,
): Promise<SignInResponse> => {
  const resp = await fetch('/api/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, password }),
  });
  if (!resp.ok) {
    throw new Error('Failed to sign in');
  }

  const payload = await resp.json();
  return payload.data;
};

export const postSignUp = async (
  username: string,
  password: string,
): Promise<SignUpResponse> => {
  const resp = await fetch('/api/auth/signup', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, password }),
  });
  if (!resp.ok) {
    throw new Error('Failed to sign up');
  }

  const payload = await resp.json();
  return payload.data;
};

export default { postLogIn, postSignUp };
