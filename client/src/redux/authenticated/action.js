import authenticationService from '../../services/authentication';

export const authActionType = {
  REGISTER_USER: 'REGISTER_USER',
  LOGIN_USER:'LOGIN_USER',
}

export const registerAction = (userData) => async ( dispatch) => {
  try{
    let formData = userData
    formData.dateofbirth = new Date(userData.dateofbirth).toISOString();
    await authenticationService.register(formData);
    return dispatch({ type: authActionType.REGISTER_USER })
  }catch (error) {
    throw error
  }
}

export const loginAction = (userData) => async (dispatch) => {
  try {
    const authService = await authenticationService.login((userData));
    return dispatch({type: authActionType.LOGIN_USER}) 
  } catch (error) {
    throw error
  }
 
}
