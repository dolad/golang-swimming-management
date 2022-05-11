
import usersServices from '../../services/usersServices';

export const USER_ACTION_TYPES = {
  GET_USERS: 'GET_USERS',
  SWIMMERS_PROFILE: 'SWIMMERS_PROFILE',
  UPDATE_SWIMMER: 'UPDATE_SWIMMER'
}

export const getUserActions = () => async ( dispatch) => {
  try{
   const usersList = await usersServices.getUsers()
    return dispatch({ type: USER_ACTION_TYPES.GET_USERS, payload: usersList })
  }catch (error) {
    throw error
  }
}

export const getSwimmerProfileAction = () => async ( dispatch) => {
    try{
      const userProfile = await usersServices.getSwimmerProfile()
      return dispatch({ type: USER_ACTION_TYPES.SWIMMERS_PROFILE, payload: userProfile })
    }catch (error) {
      throw error
    }
  }

  export const updateSwimmerAction = (userData) => async ( dispatch) => {
    try{
      const userProfile = await usersServices.updateSwimmerProfile(userData)
      return dispatch({ type: USER_ACTION_TYPES.UPDATE_SWIMMER, payload: userProfile })
    }catch (error) {
      throw error
    }
  }

