import { createStore, applyMiddleware, combineReducers } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'
import thunkMiddleware from 'redux-thunk'
import task from './modules/task'
import id from './modules/id'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const bindMiddleware = (middleware: any) => {
  if (process.env.NODE_ENV !== 'production') {
    return composeWithDevTools(applyMiddleware(...middleware))
  }
  return applyMiddleware(...middleware)
}

// root reducer
const rootReducer = combineReducers({
  task,
  id,
})

// state type
export type RootState = ReturnType<typeof rootReducer>

// store
export const store = () => {
  return createStore(rootReducer, bindMiddleware([thunkMiddleware]))
}
