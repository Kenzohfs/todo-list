"use client"

import { api } from "@/api";
import { Task, TaskStatus } from "@/types";
import axios from "axios";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

async function fetchTasks(workspaceId: string): Promise<Task[]> {
  try {
    const res = await api.get(`http://fedora:8080/api/v1/workspaces/${workspaceId}/tasks?page=0&perPage=1000&sortBy=name&sortDirection=asc`);
    return res.data;
  } catch (err) {
    console.error("error fetching tasks: ", err);
    return [];
  }
}

export default function WorksapceTasksPage() {
  const { workspaceId } = useParams();

  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!workspaceId) return;

    const getTasks = async () => {
      const tasks = await fetchTasks(workspaceId as string);
      setTasks(tasks)
      setLoading(false)
    }

    getTasks()
  }, [workspaceId])

  if (loading) {
    return <div>Loading tasks...</div>;
  }

  return (
    <div className="min-h-screen bg-gray-900 p-4">
      <h1 className="text-3xl font-bold mb-6">Tasks in Workspace {workspaceId}</h1>

      <div className="space-y-4">
        {tasks.map((task) => (
          <div
            key={task.ID}
            className={`p-4 bg-white rounded shadow-md ${task.status == TaskStatus.Finished ? "bg-green-100" : "bg-red-100"}`}
          >
            <h3 className="text-lg font-semibold">{task.desc}</h3>
            <p>{task.status == TaskStatus.Finished ? "Completed" : "Pending"}</p>
          </div>
        ))}
      </div>
    </div>
  );

}
