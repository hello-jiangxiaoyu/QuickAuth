import React  from "react";
import { useRouter } from 'next/router';
import {Card, Tabs, Typography, Form, Input, Button, Space} from '@arco-design/web-react';
import TenantInfo from './tenant'
import LoginAuth from './login'

function AppInfo() {
  return (
    <>
      <Form style={{ width: 600 }} autoComplete='off'>
        <h3>基本信息</h3>
        <Form.Item label={'app name'} rules={[{ required: true }]}>
          <Input placeholder='please enter your app name...' />
        </Form.Item>
        <Form.Item label={'应用描述'}>
          <Input.TextArea placeholder='please enter app description...' />
        </Form.Item>
        <Form.Item wrapperCol={{ offset: 5 }}>
          <Space>
            <Button type='primary'>保存</Button>
            <Button type='secondary'>重置</Button>
          </Space>
        </Form.Item>
      </Form>

      <Form style={{ width: 600 }} autoComplete='off'>
        <h3>端点信息</h3>
        <Form.Item label={'app name'} rules={[{ required: true }]}>
          <Input placeholder='please enter your app name...' />
        </Form.Item>
        <Form.Item label={'应用描述'}>
          <Input.TextArea placeholder='please enter app description...' />
        </Form.Item>
        <Form.Item wrapperCol={{ offset: 5 }}>
          <Space>
            <Button type='primary'>保存</Button>
            <Button type='secondary'>重置</Button>
          </Space>
        </Form.Item>
      </Form>
    </>
  )
}

export default function Page() {
  const router = useRouter();
  let clientId = '';
  if (typeof router.query.clientId === 'string') {
    clientId = router.query.clientId
  } else {
    return (<>duplicate client ID</>)
  }

  return (
    <>
      <Card style={{minHeight: '80vh'}}>
        <h2>应用管理</h2>
        <Tabs defaultActiveTab='1'>
          <Tabs.TabPane key='1' title='应用信息'>
            <Typography.Paragraph><AppInfo></AppInfo></Typography.Paragraph>
          </Tabs.TabPane>
          <Tabs.TabPane key='2' title='登录权限'>
            <Typography.Paragraph><LoginAuth clientId={clientId}></LoginAuth></Typography.Paragraph>
          </Tabs.TabPane>
          <Tabs.TabPane key='3' title='租户信息'>
            <Typography.Paragraph><TenantInfo clientId={clientId}></TenantInfo></Typography.Paragraph>
          </Tabs.TabPane>
        </Tabs>
      </Card>
    </>
  )
}
