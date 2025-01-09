"use client"

import { WorkspaceModel } from "@/types"
import React, { useState } from "react"

interface WorkspaceFormProps {
  createWorkspaceAction: (form: WorkspaceModel) => Promise<void>
}

export const WorkspaceForm: React.FC<WorkspaceFormProps> = ({ createWorkspaceAction: createWorkspace }) => {
  const [open, setOpen] = useState(false)
  const [name, setName] = useState("")

  const onCreateClick = () => {
    setOpen(true)
  }

  const onSubmit = async () => {
    await createWorkspace({ name })

    setName("")
    setOpen(false)
  }

  return (
    <>
      <button onClick={onCreateClick} className="p-3 rounded bg-blue-800 hover:bg-blue-900 font-semibold text-xl">Create</button>
      {open && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-gray-700 p-6 rounded shadow-lg w-96">
            <h2 className="text-xl font-semibold mb-4">Create Workspace</h2>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="w-full px-4 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Enter workspace name"
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

export default WorkspaceForm;
