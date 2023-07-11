import {Button, Card, Descriptions, Form, Input, Link, Space} from "@arco-design/web-react";
import React from "react";
import {IconDelete} from "@arco-design/web-react/icon";

function BasicInfo() {
  return (
    <Form style={{ width:600 }} autoComplete='off'>
      <h3 style={{marginLeft:20}}>基本信息</h3>
      <Form.Item label={'应用名称'} rules={[{ required: true }]}>
        <Input placeholder='please enter your app name...' />
      </Form.Item>
      <Form.Item label={'应用描述'}>
        <Input.TextArea placeholder='please enter app description...' />
      </Form.Item>
      <Form.Item wrapperCol={{ offset: 5 }}>
        <Space size='medium'>
          <Button type='primary'>保存</Button>
          <Button type='secondary'>重置</Button>
        </Space>
      </Form.Item>
    </Form>
  );
}

function AuthInformation(props:{appId:string}) {
  const data = [
    {
      label: 'App ID',
      value: '648ed50f20231ecfee93ad87',
    }, {
      label: 'JWKS 公钥',
      value: <Link style={{paddingLeft:0}} href='https://jiangzhaofeng.online/oidc/.well-known/jwks.json'>https://jiangzhaofeng.online/oidc/.well-known/jwks.json</Link>,
    }, {
      label: 'Token 接口',
      value: 'https://jiangzhaofeng.online/oidc/token',
    }, {
      label: 'Issuer',
      value: 'https://jiangzhaofeng.online/oidc',
    }, {
      label: '用户信息接口',
      value: 'https://jiangzhaofeng.online/oidc/me',
    }, {
      label: '服务发现接口',
      value: <Link style={{paddingLeft:0}} href='https://jiangzhaofeng.online/oidc/.well-known/openid-configuration'>https://jiangzhaofeng.online/oidc/.well-known/openid-configuration</Link>,
    }, {
      label: '登录页',
      value: <Link style={{paddingLeft:0}} href='https://jiangzhaofeng.online/login'>https://jiangzhaofeng.online/oidc/.well-known/openid-configuration</Link>,
    }, {
      label: '登出接口',
      value: 'https://jiangzhaofeng.online/oidc/session/end',
    }
  ];
  return (
    <Card style={{width:'100%'}}>
      <Descriptions
        column={2} colon=':' title={<h3>认证信息</h3>} data={data}
        style={{width:'100%'}} labelStyle={{ paddingRight:25 }}
      />
    </Card>
  );
}

export default function AppInfo(props: {appId: string}) {
  return (
    <>
      <BasicInfo></BasicInfo>
      <AuthInformation appId={props.appId}></AuthInformation>
      <Card style={{ width:500, height:80, marginTop:40, backgroundColor:'var(--color-fill-2)'}}>
        <Space size={80}>
          <div>此操作不可逆，请谨慎操作</div>
          <Button type='primary' status='danger' icon={<IconDelete />}>删除应用</Button>
        </Space>
      </Card>
    </>
  );
}
