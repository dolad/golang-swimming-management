import { setCookie, setLocalStorage } from 'src/utils/authHelper';
import { TOKEN, STORAGE_USER } from 'src/utils/constants';
import requestServices from './requestServices';

const login = async(userData) => {
  
  const response = await requestServices.post('http://localhost:5006/api/auth/login', userData);
  console.log(response.data);
  return response.data;
}

const register = async(userData) => {
  const response = await requestServices.post('http://localhost:5006/api/auth/signup', userData);
  return response.data;
}

const authorizeUser = async(data) => {
  setLocalStorage(STORAGE_USER, data);
  setCookie(TOKEN, data.accessToken);
}


export default  {
  login,
  register,
  authorizeUser,

}
