type SignInResponse = {
  accessToken: string;
}

type SignUpResponse = {
  accessToken: string;
}


export const SignIn = async (username: string, password: string): Promise<SignInResponse> => {
  const resp = await fetch('/api/auth/signin', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, password }),
  });
  if (!resp.ok) {
    throw new Error('Failed to sign in');
  }

  const payload = await resp.json();
  return payload.data;
}

export const SignUp = async (username: string, password: string): Promise<SignUpResponse> => {
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
}

export default { SignIn, SignUp }