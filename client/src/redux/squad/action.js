
import squadService from '../../services/squad';

export const SQUAD_DATA_ACTION_TYPES = {
  GET_ALL_SQUAD_DATA: 'GET_ALL_SQUAD_DATA',
  ADD_USER_TO_SQUAD: 'ADD_USER_TO_SQUAD',
  ADD_COACH_TO_SQUAD: 'ADD_COACH_TO_SQUAD',
}

export const getSquadsData = () => async (dispatch) => {
  try{ 
    const response = await squadService.getSquadsData();
    return dispatch({ type: SQUAD_DATA_ACTION_TYPES.GET_ALL_SQUAD_DATA, payload: response})
  }catch (error) {
    throw error
  }
}

// export const getAllUsersSwimmingData = () => async ( dispatch) => {
//     try{
//       const usersSwimmingDatas = await swimmingServices.getAllUsersSwimmingData()
//       return dispatch({ type: SWIMMING_DATA_ACTION_TYPES.GET_ALL_USER_SWIMMING_DATA, payload: usersSwimmingDatas })
//     }catch (error) {
//       throw error
//     }
//   }
