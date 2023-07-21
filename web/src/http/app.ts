import {GET, Root} from "@/http/common";

export interface App {
  id: string;
  name: string;
  describe: string;
  createTime: string;
  updateTime: string;
}

export async function fetchAppList():Promise<Root<App[]>> {
  return await GET<App[]>('/api/quick/apps')
}