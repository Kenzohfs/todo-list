export type Workspace = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  name: string;
}

export type WorkspaceModel = {
  name: string;
}

export enum TaskStatus {
  NotStarted = "NotStarted",
  InProgress = "InProgress",
  Finished = "Done",
}

export type Task = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  desc: string;
  status: TaskStatus;
}

export type TaskModel = {
  desc: string;
  status: TaskStatus
}
