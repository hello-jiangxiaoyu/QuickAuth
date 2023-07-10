import React  from "react";
import Router, { useRouter } from 'next/router';
import {Card, Tabs, Typography} from '@arco-design/web-react';
import LoginAuth from './app/login'
import AppInfo from './app/app'
import {getRouterPara} from "@/utils/stringTools";
import ClientCredential from "@/pages/applications/[appId]/app/credentials";

function Page() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  let tableKey = getRouterPara(router.query.tab);
  const tables = [
    {key: 'app', title: '应用信息', content: <AppInfo appId={appId}></AppInfo>},
    {key: 'login', title: '登录控制', content: <LoginAuth appId={appId}></LoginAuth>},
    {key: 'credential', title: '访问权限', content: <ClientCredential appId={appId}></ClientCredential>},
  ];
  if (!tables.some(ele => ele.key === tableKey)) {
    tableKey = 'app'
  }

  return (
    <Card style={{minHeight:'80vh'}}>
      <Tabs defaultActiveTab={tableKey} onClickTab={(key: string) => Router.push(`${appId}/?tab=${key}`).then()}>
        {tables.map((item) => (
          <Tabs.TabPane key={item.key} title={item.title}>
            <Typography.Paragraph>{item.content}</Typography.Paragraph>
          </Tabs.TabPane>
        ))}
      </Tabs>
    </Card>
  );
}

Page.displayName = 'Application'
export default Page;
