import { SWIMMING_DATA_ACTION_TYPES } from './action'

const swimmerData = {
  allUserSwimmingData: []
}

export default function reducer(state = swimmerData, action) {
  switch (action.type) {
    case SWIMMING_DATA_ACTION_TYPES.UPDATE_AUTH_USER_SEIMMING_DATA:
      return Object.assign({}, state)
    case SWIMMING_DATA_ACTION_TYPES.GET_ALL_USER_SWIMMING_DATA:
        return Object.assign({}, state, {
            allUserSwimmingData: action.payload
        })
    default:
      return state
  }
}
