import React, { useState, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { Task, setTasks } from '../modules/task'
import { getMaxId } from '../modules/id'
import { RootState } from '../store'
import Layout from '../components/Layout'
import Input from '../components/Input'
import TaskList from '../components/TaskList'
import FetchWrapper from '../utils/FetchWrapper'

interface Swipe {
  startY: number
  endY: number
  swiping: boolean
  swiped: boolean
}

const initialState: Swipe = {
  startY: 0,
  endY: 0,
  swiping: false,
  swiped: false,
}

export default function Index() {
  const [display, setDisplay] = useState<string | null>('hidden')
  const [swipe, setSwipe] = useState<Swipe>(initialState)
  const minDistance = 50
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(getMaxId())
    dispatch(setTasks())
  }, [])
  const tasks = useSelector((state: RootState) => state.task).payload

  // Input Componentsの表示有無を制御する
  function changeDisplay(status: string | null = null) {
    if (status == 'visible') {
      setDisplay(null)
    } else {
      setDisplay('hidden')
    }
  }

  function handleTouchStart(e: React.TouchEvent) {
    setSwipe({
      startY: e.touches[0].pageY,
      endY: 0,
      swiping: false,
      swiped: false,
    })
  }

  function handleTouchMove(e: React.TouchEvent) {
    if (e.changedTouches && e.changedTouches.length) {
      setSwipe({
        startY: swipe.startY,
        endY: e.changedTouches[0].pageY,
        swiping: true,
        swiped: false,
      })
    }
  }

  async function handleTouchEnd() {
    const absY = Math.abs(swipe.endY - swipe.startY)
    if (swipe.swiping && absY > minDistance) {
      const signY = Math.sign(swipe.endY - swipe.startY)
      if (signY > 0) {
        changeDisplay('visible')
      } else {
        // dispatch(deleteTasks())

        const result = await FetchWrapper(process.env.BASE_URL + '/tasks/', {
          method: 'DELETE',
          mode: 'cors',
          headers: {
            'Content-Type': 'application/json; charset=utf-8',
          },
        })
      }
    }
    setSwipe(initialState)
  }

  return (
    <Layout>
      <div
        className="container mx-auto"
        onTouchStart={handleTouchStart}
        onTouchMove={handleTouchMove}
        onTouchEnd={handleTouchEnd}
      >
        <div className="flex flex-col">
          <ul>
            <li key="0">
              <div className={`${display}`}>
                <Input changeDisplay={changeDisplay} />
              </div>
            </li>
            {tasks.map((task: Task) => (
              <TaskList task={task} key={task.id.toString()} />
            ))}
          </ul>
        </div>
      </div>
    </Layout>
  )
}
