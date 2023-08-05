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
export interface User {
  id: string;
  userPoolId: number;
  username: string;
  displayName: string;
  email: string;
  phone: string;
  type: number;
  isDisabled: boolean;
  createdAt: string;
  updatedAt: string;
}

// 用户池列表
export async function fetchUserPoolList():Promise<Root<Pool[]>> {
  return await GET<Pool[]>('/api/quick/user-pools', 'Get user pool list');
}
// 用户池详情
export async function fetchUserPool(poolId:number):Promise<Root<Pool[]>> {
  return await GET<Pool[]>(`/api/quick/user-pools/${poolId}`, 'Get user pool');
}
// 创建新的用户池
export async function createUserPool(data:Pool):Promise<Root<Pool>> {
  return await POST<Pool>(`/api/quick/user-pools`, data, 'Create user pool');
}
// 修改用户池信息
export async function modifyUserPool(poolId:number, data:Pool):Promise<Root<object>> {
  return await PUT(`/api/quick/user-pools/${poolId}`, data, 'Modify user pool');
}
// 删除用户池
export async function deleteUserPool(poolId:number):Promise<Root<object>> {
  return await DELETE(`/api/quick/user-pools/${poolId}`, 'Delete user pool');
}


// 用户列表
export async function fetchUserList(poolId:number):Promise<Root<User[]>> {
  return await GET<User[]>(`/api/quick/user-pools/${poolId}/users`, 'Get user list');
}
// 用户详情
export async function fetchUser(poolId:number, userId:string):Promise<Root<User>> {
  return await GET<User>(`/api/quick/user-pools/${poolId}/users/${userId}`, 'Get user');
}
// 创建用户
export async function createUser(poolId:number, data:User):Promise<Root<User>> {
  return await POST<User>(`/api/quick/user-pools/${poolId}/users`, data, 'Create user');
}
// 修改用户信息
export async function modifyUser(poolId:number, userId:string, data:User):Promise<Root<object>> {
  return await PUT(`/api/quick/user-pools/${poolId}/users/${userId}`, data, 'Modify user');
}
//删除用户
export async function deleteUser(poolId:number, userId:string):Promise<Root<object>> {
  return await DELETE(`/api/quick/user-pools/${poolId}/users/${userId}`, 'Delete user');
}