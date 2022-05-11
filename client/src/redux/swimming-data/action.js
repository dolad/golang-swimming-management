
import swimmingServices from '../../services/swimmingServices';

export const SWIMMING_DATA_ACTION_TYPES = {
  GET_ALL_USER_SWIMMING_DATA: 'GET_ALL_USER_SWIMMING_DATA',
  UPDATE_AUTH_USER_SEIMMING_DATA: 'UPDATE_AUTH_USER_SEIMMING_DATA'
}

export const updateUserSwimmingData = (formData) => async (dispatch) => {
  try{
    
    const response = await swimmingServices.updateSwimmingData(formData)
    return dispatch({ type: SWIMMING_DATA_ACTION_TYPES.UPDATE_AUTH_USER_SEIMMING_DATA})
  }catch (error) {
    throw error
  }
}

export const getAllUsersSwimmingData = () => async ( dispatch) => {
    try{
      const usersSwimmingDatas = await swimmingServices.getAllUsersSwimmingData()
      return dispatch({ type: SWIMMING_DATA_ACTION_TYPES.GET_ALL_USER_SWIMMING_DATA, payload: usersSwimmingDatas })
    }catch (error) {
      throw error
    }
  }
