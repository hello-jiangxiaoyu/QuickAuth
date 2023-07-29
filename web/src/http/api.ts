import {fetchAppList, fetchApp, createApp, modifyApp, deleteApp} from "@/http/app";
import {fetchTenantList, fetchTenant, createTenant, modifyTenant, deleteTenant} from "@/http/tenant";
import {fetchUserInfo} from "@/http/users";

// backend http api
class QuickAuthBackendApi {
  fetchUserInfo = fetchUserInfo;

  // applications
  fetchAppList = fetchAppList;
  fetchApp = fetchApp;
  createApp = createApp;
  modifyApp = modifyApp;
  deleteApp = deleteApp;

  // tenants
  fetchTenantList = fetchTenantList;
  fetchTenant = fetchTenant;
  createTenant = createTenant;
  modifyTenant = modifyTenant;
  deleteTenant = deleteTenant;
}

const api = new QuickAuthBackendApi

export default api;
