import {POSTForm, Root} from "@/http/common";

export interface loginForm {
  username: string;
  password: string;
}

export async function login(data:loginForm):Promise<Root<object>> {
  return await POSTForm<object>('/api/quick/login', data);
}
