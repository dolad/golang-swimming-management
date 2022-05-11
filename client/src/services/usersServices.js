
import requestServices from './requestServices';

const getUsers = async() => {
  const response = await requestServices.get('http://localhost:5006/api/users');
  return response.data;
}

const getSwimmerProfile = async() => {
    const response = await requestServices.get('http://localhost:5006/api/users/profile');
    return response.data;
}

const updateSwimmerProfile = async(userData) => {
    try {
        const response = await requestServices.put('http://localhost:5006/api/users/update-profile', userData);
        console.log(response.data);
        return response.data;
    } catch (error) {
    console.log(error);
      throw error   
    }
}


export default  {
  getUsers,
  getSwimmerProfile,
  updateSwimmerProfile,
}
