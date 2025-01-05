import { api } from "@/api";
import { Workspace } from "@/types"
import Link from "next/link";

async function fetchWorkspaces(): Promise<Workspace[]> {
  try {
    const res = await api.get("/workspaces?page=0&perPage=1000&sortBy=name&sortDirection=asc")
    console.log(res.data)
    return res.data
  } catch (error) {
    console.error("error fetching workspaces: ", error);
    return []
  }
}

export default async function WorkspacesPage() {
  const workspaces = await fetchWorkspaces();

  return (
    <div className="min-h-screen bg-gray-900 p-4">
      <h1 className="text-3xl font-bold mb-6">Your Workspaces</h1>

      <div className="space-y-4">
        {workspaces.map((workspace) => (
          <Link key={workspace.ID} href={`/tasks/${workspace.ID}`}>
            <div className="p-4 bg-gray-400 rounded shadow-md flex justify-between items-center hover:bg-gray-100 text-black">
              <span className="font-bold">{workspace.name}</span>
            </div>
          </Link>
        ))}
      </div>
    </div>
  )
}
