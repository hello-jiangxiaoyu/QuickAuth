import {DELETE, GET, POST, PUT, Root} from "@/http/common";

export default interface App {id?:string, name:string, host?:string, tag:string, icon:string, describe:string}

export async function fetchAppList():Promise<Root<App[]>> {
  return await GET<App[]>('/api/quick/apps', 'Get application list');
}

export async function fetchApp(appId:string):Promise<Root<App>> {
  return await GET<App>(`/api/quick/apps/${appId}`, 'Get application details');
}

export async function createApp(data:App):Promise<Root<App>> {
  return await POST<App>('/api/quick/apps', data, 'Create application');
}

export async function modifyApp(appId:string, data:App):Promise<Root<object>> {
  return await PUT(`/api/quick/apps/${appId}`, data, 'Modify application');
}

export async function deleteApp(appId:string):Promise<Root<object>> {
  return await DELETE(`/api/quick/apps/${appId}`, 'Delete application');
}


