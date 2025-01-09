import { api } from "@/api";
import WorkspaceForm from "@/components/workspace-form";
import { Workspace, WorkspaceModel } from "@/types"
import { revalidatePath } from "next/cache";
import Link from "next/link";

async function fetchWorkspaces(): Promise<Workspace[]> {
  try {
    const queries = "page=0&perPage=1000&sortBy=name&sortDirection=asc"
    const res = await api.get(`/workspaces?${queries}`)
    return res.data
  } catch (error) {
    console.error("error fetching workspaces: ", error);
    return []
  }
}

async function createWorkspace(workspace: WorkspaceModel) {
  "use server"
  await api.post("/workspaces", workspace)
  revalidatePath("(workspaces)")
}

export default async function WorkspacesPage() {
  const workspaces = await fetchWorkspaces();

  return (
    <div className="min-h-screen bg-gray-900 p-4">
      <div className="w-full flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Your Workspaces ({workspaces?.length})</h1>
        <WorkspaceForm createWorkspaceAction={createWorkspace} />
      </div>

      <div className="space-y-4 flex flex-col gap-1">
        {workspaces.map((workspace) => (
          <Link key={workspace.ID} href={`/${workspace.ID}/tasks`}>
            <div className="p-4 bg-gray-400 rounded shadow-md flex justify-between items-center hover:bg-gray-100 text-black">
              <span className="font-bold">{workspace.name}</span>
            </div>
          </Link>
        ))}
      </div>
    </div>
  )
}
