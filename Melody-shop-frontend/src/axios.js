
import axios from 'axios';
const rootPath="http://localhost:3000/api"
const axiosInstance = axios.create({
    baseURL: rootPath,
});
axiosInstance.interceptors.request.use(
    (config) => new Promise((resolve) => setTimeout(() =>  resolve(config), 2000 ))
)
export default axiosInstance