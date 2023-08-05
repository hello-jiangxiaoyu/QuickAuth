import {DELETE, GET, POST, PUT, Root} from "@/http/common";

export interface TenantDetail {
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
  isCode: boolean;
  isRefresh: boolean;
  isPassword: boolean;
  isCredential: boolean;
  isDeviceFlow: boolean;
  describe: string;
  isDisabled: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface Tenant {
  key: number;
  id: number;
  appId: string;
  userPoolId: number;
  type: number;
  name: string;
  host: string;
  company: string;
}

export async function fetchTenantList(appId:string):Promise<Root<Tenant[]>> {
  return await GET<Tenant[]>(`/api/quick/apps/${appId}/tenants`, 'Get tenant list');
}

export async function fetchTenant(appId:string, tenantId:number):Promise<Root<TenantDetail>> {
  return await GET<TenantDetail>(`/api/quick/apps/${appId}/tenants/${tenantId}`, 'Get tenant');
}

export async function createTenant(appId:string, data:Tenant):Promise<Root<Tenant>> {
  return await POST<Tenant>(`/api/quick/apps/${appId}/tenants`, data, 'Create tenant');
}

export async function modifyTenant(appId:string, tenantId:number, data:TenantDetail):Promise<Root<object>> {
  return await PUT(`/api/quick/apps/${appId}/tenants/${tenantId}`, data, 'Modify tenant');
}

export async function deleteTenant(appId:string, tenantId:number):Promise<Root<object>> {
  return await DELETE(`/api/quick/apps/${appId}/tenants/${tenantId}`, 'Delete tenant');
}
