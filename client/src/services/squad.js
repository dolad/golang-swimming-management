import requestServices from './requestServices';

const getSquadsData = async() => {
    try {
        const response = await requestServices.get('http://localhost:5006/api/squad/');
        return response.data;
    } catch (error) {
        console.log(error);
    }
 
}

// const getAllUsersSwimmingData = async() => {
//     const response = await requestServices.get('http://localhost:5006/api/swimming-data');
//     return response.data;
// }


export default  {
    getSquadsData,
    // updateSwimmingData

}
