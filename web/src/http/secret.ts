import {DELETE, GET, POST, PUT, Root} from "@/http/common";

export default interface Secret {
  id: number;
  appId: string;
  secret: string;
  scope: string;
  accessExpire: number;
  refreshExpire: number;
  describe: string;
  createdAt: string;
  updatedAt: string;
}

export async function fetchSecretList(appId:string):Promise<Root<Secret[]>> {
  return await GET<Secret[]>(`/api/quick/apps/${appId}/secrets`);
}

export async function fetchSecret(appId:string, secretId:number):Promise<Root<Secret>> {
  return await GET<Secret>(`/api/quick/apps/${appId}/secrets/${secretId}`);
}

export async function createSecret(appId:string, data:Secret):Promise<Root<Secret>> {
  return await POST<Secret>(`/api/quick/apps/${appId}/secrets`, data);
}

export async function deleteSecret(appId:string, secretId:string):Promise<Root<object>> {
  return await DELETE(`/api/quick/apps/${appId}/secrets/${secretId}`);
}
