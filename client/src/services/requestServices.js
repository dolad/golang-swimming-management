import axios from 'axios';
import { getCookie, getFromLocalStorage } from '../utils/authHelper';
const TOKEN = 'token';

const requestServices = () => {
  const baseUrl = 'http://localhost:5006';

  const configOption = {
    baseUrl,
    headers: {
      "Content-Type": 'application/json',
      "Accept": "application/json"
    }
  };

  let axiosInstance = axios.create(configOption);
  
  axiosInstance.interceptors.request.use(async config =>{
  //  extract token from the localStorage and attached it to headers
    const token = getCookie(TOKEN);
    config.headers.Authorization = token ? `Bearer ${token}` : '';
    return config;
  })

  return axiosInstance;
}


export default requestServices()
