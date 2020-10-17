import React, { useState, useEffect, useRef } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { Task, addTask } from '../modules/task'
import { getMaxId } from '../modules/id'
import { RootState } from '../store'
import moment from 'moment'
import FetchWrapper from '../utils/FetchWrapper'

interface Props {
  changeDisplay: (status: string) => void
}

export default function Input({ changeDisplay }: Props) {
  const [task, setTask] = useState<string>('')
  const inputEl = useRef<HTMLInputElement>(null)
  const dispatch = useDispatch()
  const tasks = useSelector((state: RootState) => state.task).payload
  const maxId = useSelector((state: RootState) => state.id).payload

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()

    const order = tasks
      .map((task: Task) => {
        return task.order
      })
      .reduce((a: number, b: number) => Math.max(a, b))

    const postData: Task = {
      id: maxId + 1,
      name: task,
      status: true,
      order: order + 1,
      timestamp: moment().format('YYYY-MM-DD'),
    }

    dispatch(addTask(postData))
    setTask('')
    changeDisplay('invisible')

    await FetchWrapper(process.env.BASE_URL + '/tasks', {
      method: 'POST',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json; charset=utf-8',
      },
      body: JSON.stringify(postData),
    })

    dispatch(getMaxId())
  }

  useEffect(() => {
    if (inputEl && inputEl.current) {
      inputEl.current.focus()
    }
  })

  return (
    <>
      <form onSubmit={handleSubmit}>
        <input
          className={`w-full bg-gray-500`}
          type="text"
          ref={inputEl}
          value={task}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
            setTask(e.target.value)
          }}
          onBlur={() => {
            setTask('')
            changeDisplay('invisible')
          }}
        ></input>
      </form>
    </>
  )
}
