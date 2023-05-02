import axios from 'axios';

const VITE_API_ROOT = import.meta.env.VITE_API_ROOT;

const api = axios.create({
    baseURL: VITE_API_ROOT,
    withCredentials: true,
});

export default api;