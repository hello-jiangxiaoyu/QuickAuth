import React from 'react';
import {Card, Tabs, Typography} from "@arco-design/web-react";
import Router, {useRouter} from "next/router";
import {getRouterPara} from "@/utils/stringTools";
import TenantInfo from "@/pages/applications/[appId]/tenants/[tenantId]/info";


function Page() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  let tableKey = getRouterPara(router.query.tab);
  const tables = [
    {key: 'tenants', title: '租户信息', content: <TenantInfo appId={appId}></TenantInfo>},
  ];
  if (!tables.some(ele => ele.key === tableKey)) {
    tableKey = 'tenants'
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

export default Page;
