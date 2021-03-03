import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import FetchWrapper from '../../utils/FetchWrapper'

export interface Task {
  id: number
  name: string
  status: boolean
  order: number
  timestamp: string
}

interface InitialState {
  payload: Task[]
}

export const setTasks = createAsyncThunk(
  'task/setTasks',
  async (args, thunkAPI) => {
    const response: Task[] = await FetchWrapper(process.env.BASE_URL + '/tasks')
    return response
  }
)

const initialState: InitialState = {
  payload: [],
}

const taskSlice = createSlice({
  name: 'task',
  initialState: initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(setTasks.fulfilled, (state, action) => {
      return {
        ...state,
        payload: action.payload.map((task: Task, index: number) => {
          return {
            ...task,
            order: index + 1,
          }
        }),
      }
    })
  },
})

export default taskSlice.reducer
