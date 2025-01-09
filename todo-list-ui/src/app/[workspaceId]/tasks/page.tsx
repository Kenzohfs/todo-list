import { api } from "@/api";
import Task from "@/components/task";
import TaskForm from "@/components/task-form";
import { Task as TaskType } from "@/types";
import Link from "next/link";
import { IoIosArrowBack } from "react-icons/io";

async function fetchTasks(workspaceId: string | string[] | undefined): Promise<TaskType[]> {
  if (workspaceId == null) return []

  try {
    const res = await api.get(`/workspaces/${workspaceId}/tasks?page=0&perPage=1000&sortBy=created_at&sortDirection=asc`);
    return res.data;
  } catch (err) {
    console.error("error fetching tasks: ", err);
    return [];
  }
}

export default async function WorksapceTasksPage({ params }: any) {
  const { workspaceId } = await params;
  const tasks = await fetchTasks(workspaceId);
  console.log(tasks)

  return (
    <div className="min-h-screen bg-gray-900 p-4">
      <div className="flex gap-4 items-center justify-between w-full mb-6">
        <h1 className="flex items-center gap-2 text-3xl font-bold ">
          <Link href={`/`}>
            <IoIosArrowBack />
          </Link>
          Tasks in Workspace {workspaceId}</h1>
        <TaskForm workspaceId={parseInt(workspaceId)} />
      </div>

      <div className="space-y-4">
        {tasks.map((task) => (
          <Task key={task.ID} workspaceId={parseInt(workspaceId)} task={task} />
        ))}
      </div>
    </div>
  );
}
