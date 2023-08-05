import App from "@/http/app";
import {Tenant, TenantDetail} from "@/http/tenant";
import {configureStore} from "@reduxjs/toolkit";
import env from "@/store/env.json";

export interface GlobalState{
  appList:  Array<App>;
  currentApp:  App;
  tenantList:  Array<Tenant>;
  currentTenant:  TenantDetail;

  collapsed:boolean;
  menuWidth:number;
}

// 全局变量初始值
const initialState:GlobalState = {
  appList:  [] as Array<App>,        // 全局app list
  currentApp: {} as App,             // 当前所在app
  tenantList: [] as Array<Tenant>,   // 当前app下的tenant list
  currentTenant: {} as TenantDetail, // 当前所在tenant

  collapsed: false,         // 用户池侧边单行菜单是否收起
  menuWidth: env.menuWidth, // 用户池侧边栏菜单宽度
};

// redux动作
export const DispatchAppList    = 'update-app-list';    // 更新app列表
export const DispatchApp        = 'update-app';         // 更新当前app
export const DispatchTenantList = 'update-tenant-list'; // 更新租户列表
export const DispatchTenant     = 'update-tenant';      // 更新当前租户信息
export const DispatchCollapsed  = 'set-menu-collapsed'; // 设置侧边栏收起或展开

export default function reducer(state = initialState, action) {
  console.log("--dispatch: ", action.type, action.payload);
  switch (action.type) {
    case DispatchAppList: {
      const { appList } = action.payload;
      return {...state, appList};
    }
    case DispatchApp: {
      const { currentApp } = action.payload;
      return {...state, currentApp};
    }
    case DispatchTenantList: {
      const { tenantList } = action.payload;
      return {...state, tenantList};
    }
    case DispatchTenant: {
      const { currentTenant } = action.payload;
      return {...state, currentTenant};
    }
    case DispatchCollapsed: {
      const { collapsed, menuWidth = env.menuWidth } = action.payload;
      return {...state, collapsed, menuWidth};
    }
    default:
      return state;
  }
}

// 设置app列表
export function dispatchAppList(data:Array<App>) {
  store.dispatch({
    type: DispatchAppList,
    payload: {appList: data},
  })
}

// 设置当前app
export function dispatchApp(data: App) {
  store.dispatch({
    type: DispatchApp,
    payload: {currentApp: data},
  })
}

// 设置当前app下的租户列表
export function dispatchTenantList(data:Array<Tenant>) {
  store.dispatch({
    type: DispatchTenantList,
    payload: {tenantList: data},
  })
}

// 设置当前租户
export function dispatchTenant(data:TenantDetail) {
  store.dispatch({
    type: DispatchTenant,
    payload: {currentTenant: data},
  })
}

// 设置侧边栏收起或展开
export function dispatchMenuCollapse(collapsed:boolean) {
  const menuWidth = collapsed ? env.menuCollapseWith : env.menuWidth;
  store.dispatch({
    type: DispatchCollapsed,
    payload: {collapsed: collapsed, menuWidth: menuWidth},
  })
}

export const store = configureStore({
  reducer,
  middleware: (getDefaultMiddleware) => {
    return getDefaultMiddleware()
  },
})
