import App from "@/http/app";
import {Tenant, TenantDetail} from "@/http/tenant";
import {configureStore} from "@reduxjs/toolkit";

export interface GlobalState{
  appList:  Array<App>;
  currentApp:  App;
  tenantList:  Array<Tenant>;
  currentTenant:  TenantDetail;
}
const initialState:GlobalState = {
  appList:  [] as Array<App>,
  currentApp: {} as App,

  tenantList: [] as Array<Tenant>,
  currentTenant: {} as TenantDetail,
};

export const DispatchAppList = 'update-app-list';
export const DispatchApp = 'update-app';
export const DispatchTenantList = 'update-tenant-list';
export const DispatchTenant = 'update-tenant';

export default function reducer(state = initialState, action) {
  switch (action.type) {
    case DispatchAppList: {
      const { appList } = action.payload;
      return {
        ...state,
        appList,
      };
    }
    case DispatchApp: {
      const { currentApp } = action.payload;
      return {
        ...state,
        currentApp,
      };
    }
    case DispatchTenantList: {
      const { tenantList } = action.payload;
      return {
        ...state,
        tenantList,
      };
    }
    case DispatchTenant: {
      const { currentTenant } = action.payload;
      return {
        ...state,
        currentTenant,
      };
    }
    default:
      return state;
  }
}

export const store = configureStore({
  reducer,
  middleware: (getDefaultMiddleware) => {
    return getDefaultMiddleware()
  },
})

export function dispatchAppList(data: Array<App>) {
  store.dispatch({
    type: DispatchAppList,
    payload: {appList: data},
  })
}

export function dispatchApp(data: App) {
  store.dispatch({
    type: DispatchApp,
    payload: {currentApp: data},
  })
}

export function dispatchTenantList(data: Array<Tenant>) {
  console.log(data)
  store.dispatch({
    type: DispatchTenantList,
    payload: {tenantList: data},
  })
}

export function dispatchTenant(data: TenantDetail) {
  store.dispatch({
    type: DispatchTenant,
    payload: {currentTenant: data},
  })
}