import store from "@/store/mobx";
import axios from "axios";

export function fetchUserInfo() {
  store.setUserLoading(true)
  axios.get('/api/user/userInfo').then((res) => {
    store.setUserLoading(false)
    store.setUserInfo({userInfo: res.data, userLoading: false })
  });
}

