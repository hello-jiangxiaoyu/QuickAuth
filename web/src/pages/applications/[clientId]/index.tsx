import React  from "react";
import { useRouter } from 'next/router';
import {Card, Tabs, Typography} from '@arco-design/web-react';
import TenantInfo from './tenant'
import LoginAuth from './login'
import AppInfo from './app'

export default function Page() {
  const router = useRouter();
  let clientId = '';
  if (typeof router.query.clientId === 'string') {
    clientId = router.query.clientId;
  } else {
    return (<>duplicate client ID</>);
  }

  return (
    <>
      <Card style={{minHeight: '80vh'}}>
        <h2>应用管理({clientId})</h2>
        <div style={{marginLeft:20}}>
          <Tabs defaultActiveTab='1'>
            <Tabs.TabPane key='1' title='应用信息'>
              <Typography.Paragraph><AppInfo clientId={clientId}></AppInfo></Typography.Paragraph>
            </Tabs.TabPane>
            <Tabs.TabPane key='2' title='登录配置'>
              <Typography.Paragraph><LoginAuth clientId={clientId}></LoginAuth></Typography.Paragraph>
            </Tabs.TabPane>
            <Tabs.TabPane key='3' title={'租户信息'}>
              <Typography.Paragraph><TenantInfo clientId={clientId}></TenantInfo></Typography.Paragraph>
            </Tabs.TabPane>
          </Tabs>
        </div>
      </Card>
    </>
  )
}
