import { SQUAD_DATA_ACTION_TYPES } from './action'

const squadData = {
  allSquadData: [],
  singleSquadData: [],
}

export default function reducer(state = squadData, action) {
  switch (action.type) {
    case SQUAD_DATA_ACTION_TYPES.GET_ALL_SQUAD_DATA:
      return Object.assign({}, state, {
        allSquadData: action.payload
      })
    default:
      return state
  }
}
