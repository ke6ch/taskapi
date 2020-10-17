import { Dispatch, Action } from 'redux'
import FetchWrapper from '../utils/FetchWrapper'

enum ActionTypes {
  GET_MAX_ID = 'GET_MAX_ID',
}

export interface Id {
  id: number
}

interface InisitalState {
  payload: number
}

interface IdAction extends Action {
  type: ActionTypes.GET_MAX_ID
  payload: number
}

// action creator
export function getMaxId() {
  return (dispatch: Dispatch) => {
    FetchWrapper(process.env.BASE_URL + '/id').then((id) =>
      dispatch({ type: ActionTypes.GET_MAX_ID, payload: id })
    )
  }
}

const initialState: InisitalState = {
  payload: 0,
}

export default function reducer(state = initialState, action: IdAction) {
  switch (action.type) {
    case ActionTypes.GET_MAX_ID:
      return {
        ...state,
        payload: action.payload,
      }
    default:
      return state
  }
}
