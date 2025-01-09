"use server"

import { Task, TaskModel } from "@/types";
import { api } from ".";
import { revalidatePath } from "next/cache";

export const CreateTask = async (workspaceId: number, model: TaskModel) => {
  await api.post(`/workspaces/${workspaceId}/tasks`, model)
  revalidatePath("/[workspaceId]/tasks", "page")
}

export const UpdateTask = async (workspaceId: number, taskId: number, task: Partial<Task>) => {
  await api.put(`/workspaces/${workspaceId}/tasks/${taskId}`, task)
}

export const RemoveTask = async (workspaceId: number, taskId: number) => {
  await api.delete(`/workspaces/${workspaceId}/tasks/${taskId}`)
  revalidatePath("/[workspaceId]/tasks", "page")
}
