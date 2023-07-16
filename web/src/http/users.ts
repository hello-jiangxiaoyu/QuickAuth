import store from "@/store/mobx";
import axios from "axios";
import {Root, fetchData} from "@/http/common";

export function fetchUserInfo() {
  store.setUserLoading(true)
  axios.get('/api/user/userInfo').then((res) => {
    store.setUserLoading(false)
    store.setUserInfo({userInfo: res.data, userLoading: false })
  });
}



export interface Keys {
  use: string;
  kty: string;
  kid: string;
  alg: string;
  n: string;
  e: string;
}

export interface Data {
  keys: Keys[];
}

export async function fetchOIDC():Promise<Root<Keys[]>> {
  return await fetchData<Keys[]>('http://localhost/api/quick/.well-known/jwks.json')
}


