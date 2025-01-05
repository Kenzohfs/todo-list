export type Workspace = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  name: string;
}

export enum TaskStatus {
  NotStarted,
  InProgress,
  Finished,
}

export type Task = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  desc: string;
  status: TaskStatus;
}
