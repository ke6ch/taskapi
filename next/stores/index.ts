import { Action, combineReducers, Store } from 'redux'
import { configureStore } from '@reduxjs/toolkit'
import { MakeStore } from 'next-redux-wrapper'
import { ThunkAction } from 'redux-thunk'
import idReducer from './slices/id'
import taskReducer from './slices/task'

const rootReducer = combineReducers({
  id: idReducer,
  task: taskReducer,
})

export type RootState = ReturnType<typeof rootReducer>
export type AppThunk = ThunkAction<void, RootState, null, Action<any>>

const makeStore: MakeStore = (initialState): Store => {
  const store: Store = configureStore({
    reducer: rootReducer,
  })
  return store
}

export default makeStore
