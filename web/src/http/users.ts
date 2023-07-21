import store from "@/store/mobx";
import {Root, GET} from "@/http/common";

export function fetchUserInfo() {
  store.setUserLoading(false)
  store.setUserInfo({userInfo: {
      name: '王立群',
      avatar: 'https://lf1-xgcdn-tos.pstatp.com/obj/vcloud/vadmin/start.8e0e4855ee346a46ccff8ff3e24db27b.png',
      job: 'frontend',
      organization: 'Frontend',
      location: 'beijing',
      email: 'wangliqun@email.com',
      permissions: {},
    }, userLoading: false });
}

export interface Keys {
  use: string;
  kty: string;
  kid: string;
  alg: string;
  n: string;
  e: string;
}

export async function fetchOIDC():Promise<Root<Keys[]>> {
  return await GET<Keys[]>('/api/quick/.well-known/openid-configuration')
}


