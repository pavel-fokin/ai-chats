import axios from 'axios';

export const clientUnauthed = axios.create({
  baseURL: '/api',
});

export const client = axios.create({
  baseURL: '/api',
});

client.interceptors.request.use((config) => {
  const token = localStorage.getItem('accessToken');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
});

// interceptor for handling errors
client.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response.status === 401) {
      localStorage.removeItem('accessToken');
      window.location.href = '/app/login';
    }
    return Promise.reject(error);
  },
);
