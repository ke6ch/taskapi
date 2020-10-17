import React, { useState } from 'react'
import { Task, updateStatus, deleteTask } from '../modules/task'
import { getMaxId } from '../modules/id'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../store'
import FetchWrapper from '../utils/FetchWrapper'

interface Props {
  task: Task
}

interface Flick {
  startX: number
  endX: number
  flicked: boolean
}

export default function TaskList({ task }: Props) {
  const [flick, setFlick] = useState<Flick>({
    startX: 0,
    endX: 0,
    flicked: false,
  })
  const minDistance = 50
  const dispatch = useDispatch()
  const tasks: Task[] = useSelector((state: RootState) => state.task).payload

  function handleTouchStart(e: React.TouchEvent) {
    setFlick({
      startX: e.touches[0].pageX,
      endX: 0,
      flicked: false,
    })
  }

  function handleTouchMove(e: React.TouchEvent) {
    if (e.changedTouches && e.changedTouches.length) {
      setFlick({
        startX: flick.startX,
        endX: e.changedTouches[0].pageX,
        flicked: true,
      })
    }
  }

  async function handleTouchEnd(e: React.TouchEvent<HTMLLIElement>) {
    const absX = Math.abs(flick.endX - flick.startX)
    if (flick.flicked && absX > minDistance) {
      const signX = Math.sign(flick.endX - flick.startX)
      const id =
        e.currentTarget.dataset.id === undefined
          ? 0
          : e.currentTarget.dataset.id
      const task = tasks.filter((task) => task.id.toString() === id)[0]

      if (signX > 0) {
        const patchData: Task = {
          id: task.id,
          name: task.name,
          status: false,
          order: task.order,
          timestamp: task.timestamp,
        }
        dispatch(updateStatus(patchData))

        await FetchWrapper(
          process.env.BASE_URL +
            '/tasks/' +
            patchData.id +
            '?status=' +
            patchData.status,
          {
            method: 'PATCH',
            mode: 'cors',
            headers: {
              'Content-Type': 'application/json; charset=utf-8',
            },
          }
        )
      } else {
        dispatch(deleteTask(task.id))

        await FetchWrapper(process.env.BASE_URL + '/tasks/' + task.id, {
          method: 'DELETE',
          mode: 'cors',
          headers: {
            'Content-Type': 'application/json; charset=utf-8',
          },
        })

        dispatch(getMaxId())
      }
    }
  }

  return (
    <>
      <li
        className={
          task.status === true
            ? `bg-orange-${task.order * 100}`
            : `line-through bg-gray-400`
        }
        data-id={task.id}
        onTouchStart={handleTouchStart}
        onTouchMove={handleTouchMove}
        onTouchEnd={handleTouchEnd}
      >
        {task.name}
      </li>
    </>
  )
}
