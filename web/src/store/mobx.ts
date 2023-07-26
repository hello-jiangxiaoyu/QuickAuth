import {makeAutoObservable} from 'mobx'
import env from './env.json'
import {App} from '@/http/app'
import {Tenant, TenantDetail} from "@/http/tenant";
import {isIPAddress} from "@/utils/is";

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

  setCurrentApp = (app: App) => {
    this.currentApp = app;
    this.multiTenant = (app.tag === 'multi_tenant');
  }

  tenantList: Array<Tenant> = [];
  currentTenant: Tenant;
  multiTenant = false;
  setTenantList = (tenants: Array<Tenant>) => {this.tenantList = tenants}
  setCurrentTenant = (tenant: Tenant) => {this.currentTenant = tenant};
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

