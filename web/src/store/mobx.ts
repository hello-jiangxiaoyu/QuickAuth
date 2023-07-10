import {makeAutoObservable} from 'mobx'
import env from './env.json'

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
  userLoading = true;
  menuWidth = env.menuWidth;
  menuCollapsed = false;

  setUserInfo = (userInfo: GlobalState) => {this.userInfo = userInfo.userInfo}
  setUserLoading = (userLoading: boolean) => {this.userLoading = userLoading}
  setCollapsed = (collapsed: boolean) => {this.menuWidth = collapsed ? env.menuCollapseWith : env.menuWidth; this.menuCollapsed = collapsed;}
  switchCollapsed = () => {
    this.setCollapsed(!this.menuCollapsed);
  }
}

const store = new GlobalStatus()
export default store

