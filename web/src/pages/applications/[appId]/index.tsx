import React, {useEffect} from "react";
import Router, { useRouter } from 'next/router';
import {Card, Message, Tabs, Typography} from '@arco-design/web-react';
import LoginAuth from '@/pages/applications/[appId]/login'
import AppInfo from '@/pages/applications/[appId]/appInfo'
import ClientCredential from "@/pages/applications/[appId]/credentials";
import Tenants from "@/pages/applications/[appId]/tenants"
import {getRouterPara} from "@/utils/stringTools";
import {fetchApp} from "@/http/app";
import {apps} from "@/store/mobx";
import {fetchTenantList} from "@/http/tenant";

function Page() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  let tableKey = getRouterPara(router.query.tab);
  const tables = [
    {key: 'app', title: '应用信息', content: <AppInfo appId={appId}></AppInfo>},
    {key: 'login', title: '登录控制', content: <LoginAuth appId={appId}></LoginAuth>},
    {key: 'credential', title: '访问凭证', content: <ClientCredential appId={appId}></ClientCredential>},
    {key: 'tenants', title: '租户管理', content: <Tenants appId={appId}></Tenants>},
  ];
  if (!tables.some(ele => ele.key === tableKey)) {
    tableKey = 'app';
  }

  useEffect(() => {
    fetchApp(appId).then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        apps.setCurrentApp(r.data);
      }
    })
    fetchTenantList(appId).then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        apps.setTenantList(r.data);
      }
    })
  }, [appId])

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
