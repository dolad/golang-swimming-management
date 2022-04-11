import requestServices from './requestServices';

const login = async(userData) => {
  console.log(requestServices)
  const response = await requestServices.post('http://localhost:5006/api/auth/login', userData);
  console.log(response.data);
  return response.data;
}

const register = async(userData) => {

  const response = await requestServices.post('http://localhost:5006/api/auth/signup', userData);
  return response.data;
}

export default  {
  login,
  register
}
