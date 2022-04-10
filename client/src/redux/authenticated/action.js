import authenticationService from '../../services/authentication';

export const authActionType = {
  REGISTER_USER: 'REGISTER_USER',
  LOGIN_USER:'LOGIN_USER',
}

export const registerAction = (userData) => async ( dispatch) => {
  //call authServices
  const authService = await authenticationService.register(userData);
  console.log(authService);
  return dispatch({ type: authActionType.REGISTER_USER })
}

export const loginAction = (userData) => async (dispatch) => {
  const authService = await authenticationService.login((userData));
  console.log(authService);
  return dispatch({type: authActionType.LOGIN_USER})
}
