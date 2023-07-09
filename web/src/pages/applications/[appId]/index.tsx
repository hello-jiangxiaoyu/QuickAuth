import React  from "react";
import Router, { useRouter } from 'next/router';
import {Card, Tabs, Typography} from '@arco-design/web-react';
import TenantInfo from './tenant'
import LoginAuth from './login'
import AppInfo from './app'
import {getRouterPara} from "@/utils/stringTools";

function Page() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  let tableKey = getRouterPara(router.query.tab);
  const tables = [
    {key: 'app', title: '应用信息', content: <AppInfo appId={appId}></AppInfo>},
    {key: 'login', title: '登录配置', content: <LoginAuth appId={appId}></LoginAuth>},
    {key: 'tenant', title: '租户信息', content: <TenantInfo appId={appId}></TenantInfo>}
  ];
  if (!tables.some(ele => ele.key === tableKey)) {
    tableKey = 'app'
  }

  return (
    <Card style={{minHeight:'80vh'}}>
      <h2>应用管理({appId})</h2>
      <div style={{marginLeft:20}}>
        <Tabs defaultActiveTab={tableKey} onClickTab={(key: string) => Router.push(`${appId}/?tab=${key}`).then()}>
          {tables.map((item) => (
            <Tabs.TabPane key={item.key} title={item.title}>
              <Typography.Paragraph>{item.content}</Typography.Paragraph>
            </Tabs.TabPane>
          ))}
        </Tabs>
      </div>
    </Card>
  );
}

Page.displayName = 'Application'
export default Page;
