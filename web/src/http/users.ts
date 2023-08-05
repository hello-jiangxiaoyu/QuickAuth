import mobxStore from "@/store/mobx";
import {Root, GET, DELETE, PUT, POST} from "@/http/common";

export function fetchUserInfo() {
  mobxStore.setUserLoading(false)
  mobxStore.setUserInfo({userInfo: {
      name: '王立群',
      avatar: 'https://lf1-xgcdn-tos.pstatp.com/obj/vcloud/vadmin/start.8e0e4855ee346a46ccff8ff3e24db27b.png',
      job: 'frontend',
      organization: 'Frontend',
      location: 'beijing',
      email: 'wangliqun@email.com',
      permissions: {},
    }, userLoading: false });
}

export interface Pool {
  id: number;
  name: string;
  describe: string;
  isDisabled: boolean;
}

export async function fetchUserPoolList():Promise<Root<Pool[]>> {
  return await GET<Pool[]>('/api/quick/user-pools')
}

export async function fetchUserPool(poolId:number):Promise<Root<Pool[]>> {
  return await GET<Pool[]>(`/api/quick/user-pools/${poolId}`)
}

export async function createUserPool(data:Pool):Promise<Root<Pool>> {
  return await POST<Pool>(`/api/quick/user-pools`, data)
}

export async function modifyUserPool(poolId:number, data:Pool):Promise<Root<object>> {
  return await PUT(`/api/quick/user-pools/${poolId}`, data)
}

export async function deleteUserPool(poolId:number):Promise<Root<object>> {
  return await DELETE(`/api/quick/user-pools/${poolId}`)
}
