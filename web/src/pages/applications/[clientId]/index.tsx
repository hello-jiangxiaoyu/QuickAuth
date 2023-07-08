import React  from "react";
import Router, { useRouter } from 'next/router';
import {Card, Tabs, Typography} from '@arco-design/web-react';
import TenantInfo from './tenant'
import LoginAuth from './login'
import AppInfo from './app'
import {getRouterPara} from "@/utils/getUrlParams";

export default function Page() {
  const router = useRouter();
  const clientId = getRouterPara(router.query.clientId);
  let tableKey = getRouterPara(router.query.tab);
  const tables = [
    {key: 'app', title: '应用信息', content: <AppInfo clientId={clientId}></AppInfo>},
    {key: 'login', title: '登录配置', content: <LoginAuth clientId={clientId}></LoginAuth>},
    {key: 'tenant', title: '租户信息', content: <TenantInfo clientId={clientId}></TenantInfo>}
  ];
  if (!tables.some(ele => ele.key === tableKey)) {
    tableKey = 'app'
  }

  return (
    <Card style={{minHeight:'80vh'}}>
      <h2>应用管理({clientId})</h2>
      <div style={{marginLeft:20}}>
        <Tabs defaultActiveTab={tableKey} onClickTab={(key: string) => Router.push(`${clientId}/?tab=${key}`).then()}>
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
