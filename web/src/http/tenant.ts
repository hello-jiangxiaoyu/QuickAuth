import {DELETE, GET, POST, PUT, Root} from "@/http/common";

export interface Tenant {
  id: number;
  appId: string;
  userPoolId: number;
  type: number;
  name: string;
  host: string;
  company: string;
  grantTypes?: Array<string>;
  redirectUris?: Array<string>;
  codeExpire: number;
  idExpire: number;
  accessExpire: number;
  refreshExpire: number;
  isCode: number;
  isRefresh: number;
  isPassword: number;
  isCredential: number;
  isDeviceFlow: number;
  describe: string;
  isDisabled: number;
  createdAt: string;
  updatedAt: string;
}

export async function fetchTenantList(appId:string):Promise<Root<Tenant[]>> {
  return await GET<Tenant[]>(`/api/quick/apps/${appId}/tenants`)
}

export async function fetchTenant(appId:string, tenantId:string):Promise<Root<Tenant>> {
  return await GET<Tenant>(`/api/quick/apps/${appId}/tenants/${tenantId}`)
}

export async function createTenant(appId:string, data:Tenant):Promise<Root<Tenant>> {
  return await POST<Tenant>(`/api/quick/apps/${appId}/tenants`, data)
}

export async function modifyTenant(appId:string, tenantId:string, data:Tenant):Promise<Root<object>> {
  return await PUT(`/api/quick/apps/${appId}/tenants/${tenantId}`, data)
}

export async function deleteTenant(appId:string, tenantId:string):Promise<Root<object>> {
  return await DELETE(`/api/quick/apps/${appId}/tenants/${tenantId}`)
}
