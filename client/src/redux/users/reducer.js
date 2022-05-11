import { USER_ACTION_TYPES } from './actions';

const usersState = {
  usersList: [],
  swimmers: [],
  coaches:[],
  parents:[],
  swimmerProfile: {}

}

export default function reducer(state = usersState, action) {
  switch (action.type) {
    case USER_ACTION_TYPES.GET_USERS:
      return Object.assign({}, state, {
        usersList: action.payload,
        swimmers:action.payload.filter(swimmers => swimmers?.Role?.name === "swimmer"),
        coaches:action.payload.filter(swimmers => swimmers?.Role?.name === "coaches"),
        parents:action.payload.filter(swimmers => swimmers?.Role?.name === "role_parent"),
      })
    case USER_ACTION_TYPES.SWIMMERS_PROFILE:
        return Object.assign({}, state, {
         swimmerProfile: action.payload
        } )
    case USER_ACTION_TYPES.UPDATE_SWIMMER:
        return Object.assign({}, state)
    default:
      return state
  }
}
