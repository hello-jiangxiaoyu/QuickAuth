import {makeAutoObservable} from 'mobx'

export interface GlobalState {
  userInfo?: {
    name?: string;
    avatar?: string;
    job?: string;
    organization?: string;
    location?: string;
    email?: string;
    permissions: Record<string, string[]>;
  };
  userLoading?: boolean;
}

class GlobalStatus {
  constructor() {makeAutoObservable(this)}
  demo: GlobalState = {};
  userInfo: typeof this.demo.userInfo =  {permissions: {}};
  setUserInfo = (userInfo: GlobalState) => {this.userInfo = userInfo.userInfo};

  userLoading = true;
  setUserLoading = (userLoading: boolean) => {this.userLoading = userLoading};
}

const mobxStore = new GlobalStatus();
export default mobxStore;
