import {makeAutoObservable} from 'mobx'
import defaultSettings from "@/settings.json";

export interface GlobalState {
  settings?: typeof defaultSettings;
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
  settings = defaultSettings;

  setUserInfo = (userInfo: GlobalState) => {this.userInfo = userInfo.userInfo}
  setUserLoading = (userLoading: boolean) => {this.userLoading = userLoading}
  setCollapsed = (collapsed: boolean) => {this.settings.menuWidth = collapsed ? 48 : 220; this.settings.siderCollapsed = collapsed;}
  switchCollapsed = () => {
    this.settings.siderCollapsed = !this.settings.siderCollapsed;
    this.settings.menuWidth = this.settings.siderCollapsed ? 48 : 220;
  }
}

const store = new GlobalStatus()
export default store

