"use client"

import { CreateTask } from "@/api/task"
import { TaskStatus } from "@/types"
import React, { useState } from "react"

interface TaskFormProps {
  workspaceId: number
}

export const TaskForm: React.FC<TaskFormProps> = ({ workspaceId }) => {
  const [open, setOpen] = useState(false)
  const [desc, setDesc] = useState("")

  const onCreateClick = () => {
    setOpen(true)
  }

  const onSubmit = async () => {
    await CreateTask(workspaceId, { desc, status: TaskStatus.NotStarted })

    setDesc("")
    setOpen(false)
  }

  return (
    <>
      <button onClick={onCreateClick} className="p-3 rounded bg-blue-800 hover:bg-blue-900 font-semibold text-xl">Create</button>
      {open && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-gray-700 p-6 rounded shadow-lg w-96">
            <h2 className="text-xl font-semibold mb-4">Create Task</h2>
            <input
              type="text"
              value={desc}
              onChange={(e) => setDesc(e.target.value)}
              className="w-full px-4 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Enter task desc"
              autoFocus
              required
            />
            <div className="flex justify-end mt-4">
              <button
                type="button"
                onClick={() => setOpen(false)}
                className="px-4 py-2 mr-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300"
              >
                Cancel
              </button>
              <button
                onClick={onSubmit}
                className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
              >
                Create
              </button>
            </div>
          </div>
        </div>
      )}
    </>
  )
}

export default TaskForm;
