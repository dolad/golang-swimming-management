import requestServices from './requestServices';

const updateSwimmingData = async(data) => {
  const response = await requestServices.post('http://localhost:5006/api/swimming-data', data);
  return response.data;
}

const getAllUsersSwimmingData = async() => {
    const response = await requestServices.get('http://localhost:5006/api/swimming-data');
    return response.data;
}



export default  {
    getAllUsersSwimmingData,
    updateSwimmingData

}
