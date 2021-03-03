import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import FetchWrapper from '../../utils/FetchWrapper'

interface Id {
  id: number
}

export const fetchMaxId = createAsyncThunk(
  'id/fetchMaxId',
  async (args, thunkAPI) => {
    const response: Id = await FetchWrapper(process.env.BASE_URL + '/id')
    return response
  }
)

const idSlice = createSlice({
  name: 'id',
  initialState: {
    id: 0,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchMaxId.fulfilled, (state: Id, action) => {
      return { ...state, id: action.payload.id }
    })
  },
})

export default idSlice.reducer
