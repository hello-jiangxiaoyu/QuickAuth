import {makeAutoObservable} from 'mobx'
import env from './env.json'
import {App, fetchAppList} from '@/http/app'
import {Message} from "@arco-design/web-react";

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

  menuCollapsed = false;
  menuWidth = env.menuWidth;
  switchCollapsed = () => {this.setCollapsed(!this.menuCollapsed)}
  setCollapsed = (collapsed: boolean) => {
    this.menuWidth = collapsed ? env.menuCollapseWith : env.menuWidth; this.menuCollapsed = collapsed;
  }
}

class GlobalApplications {
  constructor() {makeAutoObservable(this)}

  appList: Array<App> = [];
  currentApp:App;
  updateApps = (apps: Array<App>) => {
    this.appList = apps;
  }

  setCurrentApp = (appId: string) => {
    for (const i of this.appList) {
      if (i.id === appId) {
        this.currentApp = i;
        return true
      }
    }
    return false;
  }

  tenantList = '';
  currentTenant = '';
  multiTenant = false;
  setMultiTenant = (multiTenant: boolean) => {this.multiTenant = multiTenant};

  setCurrentTenant = (tenant: string) => {this.currentTenant = tenant};
}

class GlobalUser {
  constructor() {makeAutoObservable(this)}

  apps: Array<App> = [];
  updateApps = (apps: Array<App>) => {
    this.apps = apps;
  }
}

export const apps = new GlobalApplications();
export const user = new GlobalUser();

const store = new GlobalStatus();
export default store;

