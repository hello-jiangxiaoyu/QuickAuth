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

  list: Array<App> = [];
  updateApps = (apps: Array<App>) => {
    this.list = apps;
  }

  multiTenant = false;
  setMultiTenant = (multiTenant: boolean) => {this.multiTenant = multiTenant};
}

class GlobalUser {
  constructor() {makeAutoObservable(this)}

  apps: Array<App> = [];
  updateApps = (apps: Array<App>) => {
    this.apps = apps;
  }

  multiTenant = false;
  setMultiTenant = (multiTenant: boolean) => {this.multiTenant = multiTenant};
}

export const apps = new GlobalApplications();
export const user = new GlobalUser();

const store = new GlobalStatus();
export default store;

