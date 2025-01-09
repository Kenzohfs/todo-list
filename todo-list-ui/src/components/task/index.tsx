"use client"

import { RemoveTask, UpdateTask } from '@/api/task'
import { TaskStatus, Task as TaskType } from '@/types'
import React, { useEffect, useState } from 'react'
import { IoIosRemoveCircleOutline } from 'react-icons/io'

interface TaskProps {
  task: TaskType
  workspaceId: number
}

enum Mode {
  READ_ONLY = "readonly",
  EDIT = "edit"
}

const DebouceTime = 500

const Task: React.FC<TaskProps> = ({ task, workspaceId }) => {
  const [desc, setDesc] = useState(task.desc)
  const [status, setStatus] = useState(task.status === TaskStatus.Finished ? TaskStatus.Finished : TaskStatus.NotStarted)
  const [mode, setMode] = useState(Mode.READ_ONLY)
  const [debDesc, setDebDesc] = useState(task.desc)

  const onStatusChange = (e: any) => {
    setStatus(e.target.checked ? TaskStatus.Finished : TaskStatus.NotStarted)
  }

  const onModeEdit = (_: any) => {
    setMode(Mode.EDIT)
  }

  const onRemoveClick = async () => {
    await RemoveTask(workspaceId, task.ID)
  }

  useEffect(() => {
    const id = setTimeout(() => {
      setDebDesc(desc)
    }, DebouceTime)

    return () => clearTimeout(id)
  }, [desc])

  useEffect(() => {
    UpdateTask(workspaceId, task.ID, { desc, status })
  }, [debDesc, status])

  return (
    <div className='flex items-center gap-4 p-3 bg-gray-700 rounded shadow-md'>
      <div>
        <input
          type="checkbox"
          defaultChecked={status === TaskStatus.Finished}
          onChange={onStatusChange}
        />
      </div>
      <div
        className={`flex w-full gap-4`}
        onClick={onModeEdit}
      >
        {mode === Mode.READ_ONLY ?
          (
            <p className={`${status === TaskStatus.Finished ? 'line-through' : ''}`}>
              {desc}
            </p>
          ) :
          (
            <input className="text-lg font-semibold" type="text" value={desc} onChange={(e) => setDesc(e.target.value)} onBlur={() => setMode(Mode.READ_ONLY)} autoFocus />
          )
        }
      </div>
      <div className='flex items-center gap-4'>
        <p>{status == TaskStatus.Finished ? "Completed" : "Pending"}</p>
        <IoIosRemoveCircleOutline onClick={onRemoveClick} color='red' />
      </div>
    </div >
  )
}

export default Task
