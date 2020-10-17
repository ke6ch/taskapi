import { Dispatch, Action } from 'redux'
import _ from 'lodash'
import FetchWrapper from '../utils/FetchWrapper'

enum ActionTypes {
  SET_TASKS = 'SET_TASKS',
  ADD_TASK = 'ADD_TASK',
  UPDATE_STATUS = 'UPDATE_STATUS',
  DELETE_TASK = 'DELETE_TASK',
  DELETE_TASKS = 'DELETE_TASKS',
}

export interface Task {
  id: number
  name: string
  status: boolean
  order: number
  timestamp: string
}

interface SetTasksAction extends Action {
  type: ActionTypes.SET_TASKS
  payload: Task[]
}

interface AddTaskAction extends Action {
  type: ActionTypes.ADD_TASK
  payload: Task
}

interface UpdateTaskAction extends Action {
  type: ActionTypes.UPDATE_STATUS
  payload: {
    id: number
    status: boolean
  }
}

interface DeleteTaskAction extends Action {
  type: ActionTypes.DELETE_TASK
  payload: {
    id: number
  }
}

interface DeleteTasksAction extends Action {
  type: ActionTypes.DELETE_TASKS
}

type TaskAction =
  | SetTasksAction
  | AddTaskAction
  | UpdateTaskAction
  | DeleteTaskAction
  | DeleteTasksAction

interface InisitalState {
  payload: Task[]
}

// action creator
export function setTasks() {
  return (dispatch: Dispatch) => {
    FetchWrapper(process.env.BASE_URL + '/tasks').then((tasks) =>
      dispatch({ type: ActionTypes.SET_TASKS, payload: tasks })
    )
  }
}

export function addTask(task: Task): AddTaskAction {
  return {
    type: ActionTypes.ADD_TASK,
    payload: task,
  }
}

export function updateStatus(task: Task): UpdateTaskAction {
  return {
    type: ActionTypes.UPDATE_STATUS,
    payload: {
      id: task.id,
      // name: task.name,
      status: task.status,
      // order: task.order,
      // timestamp: task.timestamp,
    },
  }
}

export function deleteTask(id: number): DeleteTaskAction {
  return {
    type: ActionTypes.DELETE_TASK,
    payload: {
      id: id,
    },
  }
}

export function deleteTasks(): DeleteTasksAction {
  return {
    type: ActionTypes.DELETE_TASKS,
  }
}

const initialState: InisitalState = {
  payload: [],
}

export default function reducer(state = initialState, action: TaskAction) {
  switch (action.type) {
    case ActionTypes.SET_TASKS:
      return {
        ...state,
        payload: state.payload
          .concat(action.payload)
          .map((task: Task, idx: number) => {
            return {
              id: task.id,
              name: task.name,
              status: task.status,
              order: idx + 1,
              timestamp: task.timestamp,
            }
          }),
      }
    case ActionTypes.ADD_TASK:
      return {
        ...state,
        payload: _.orderBy(
          [...state.payload, action.payload],
          ['status', 'order'],
          ['desc', 'asc']
        ).map((task: Task, idx: number) => {
          return {
            id: task.id,
            name: task.name,
            status: task.status,
            order: idx + 1,
            timestamp: task.timestamp,
          }
        }),
      }
    case ActionTypes.UPDATE_STATUS:
      return {
        payload: _.orderBy(
          state.payload.map((task: Task) => {
            if (task.id !== action.payload.id) {
              return task
            }
            return {
              ...task,
              ...action.payload,
            }
          }),
          ['status', 'order'],
          ['desc', 'asc']
        ).map((task: Task, idx: number) => {
          return {
            id: task.id,
            name: task.name,
            status: task.status,
            order: idx + 1,
            timestamp: task.timestamp,
          }
        }),
      }
    case ActionTypes.DELETE_TASK:
      return {
        payload: state.payload
          .filter((task) => task.id !== action.payload.id)
          .map((task: Task, idx: number) => {
            return {
              id: task.id,
              name: task.name,
              status: task.status,
              order: idx + 1,
              timestamp: task.timestamp,
            }
          }),
      }
    case ActionTypes.DELETE_TASKS:
      return {
        payload: state.payload.filter((task) => task.status === true),
      }
    default:
      return state
  }
}
