'use-strict'

import axios, { AxiosInstance } from 'axios'

const api = axios.create({
    baseURL: 'http://localhost:8080',
    withCredentials: true,
});

api.interceptors.response.use(
response => response,
error => {
if (error.response && error.response.status === 422) {
    // Perform permanent action, such as redirecting to a login page
    window.location.replace('/login');
}
return Promise.reject(error);
}
);

export interface GlobalVariables {
    api: AxiosInstance;
}

export const globalVariables: GlobalVariables = {
    api: api,
  };