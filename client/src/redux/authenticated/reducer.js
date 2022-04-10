import { authActionType } from './action';

const authenticatedState = {
  isSuccessfull: false,
  userData: {}
}

export default function reducer(state = authenticatedState, action) {
  switch (action.type) {
    case authActionType.REGISTER_USER:
      return Object.assign({}, state, {
        isSuccessfull: !state.isSuccessfull
      })
    case authActionType.LOGIN_USER:
      return Object.assign({}, state, {
        userData: action.payload
      })
    default:
      return state
  }
}
