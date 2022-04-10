import requestServices from './requestServices';

const login = async(userData) => {
  const response = await requestServices.post('auth/login', userData);
  console.log(response.data);
  return response.data;
}

const register = async(userData) => {
  const response = await requestServices.post('auth/signup', userData);
  console.log(response.data);
  return response.data;
}

export default  {
  login,
  register
}
