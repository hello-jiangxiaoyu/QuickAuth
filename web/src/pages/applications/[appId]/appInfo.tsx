import {Button, Card, Descriptions, Form, Input, Link, Message, Space} from "@arco-design/web-react";
import React, {useEffect, useState} from "react";
import {IconDelete} from "@arco-design/web-react/icon";
import {apps} from "@/store/mobx";
import {fetchTenant, Tenant} from "@/http/tenant";
import {isIPAddress} from "@/utils/is";
import {useRouter} from "next/router";
import {getRouterPara} from "@/utils/stringTools";
import {observer} from "mobx-react";
import useStorage from "@/utils/useStorage";

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


function getHostWithScheme(host:string):string {
  if (typeof host !== 'string') {
    return '';
  }
  if (host.startsWith("localhost")) {
    return 'http://' + host;
  }
  if (isIPAddress(host)) {
    return 'http://' + host;
  }
  return 'https://' + host;
}

function AuthInformation() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const [domain, setDomain] = useState('')

  useEffect(() => {
    fetchTenant(appId, 1).then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        setDomain(getHostWithScheme(r.data.host))
      }
    });
  }, [appId])

  const data = [
    {
      label: 'App ID',
      value: appId,
    }, {
      label: 'JWKS 公钥',
      value: <Link style={{paddingLeft:0}} href={`${domain}/api/quick/.well-known/jwks.json`}>{domain}/api/quick/.well-known/jwks.json</Link>,
    }, {
      label: 'Token 接口',
      value: `${domain}/api/quick/token`,
    }, {
      label: 'Issuer',
      value: `${domain}`,
    }, {
      label: '用户信息接口',
      value: `${domain}/api/quick/me`,
    }, {
      label: '服务发现接口',
      value: <Link style={{paddingLeft:0}} href={`${domain}/api/quick/.well-known/openid-configuration`}>{domain}/api/quick/.well-known/openid-configuration</Link>,
    }, {
      label: '登录页',
      value: <Link style={{paddingLeft:0}} href={`${domain}/api/quick/login`}>{domain}/api/quick/login</Link>,
    }, {
      label: '登出接口',
      value: `${domain}/api/quick/oidc/session/end`,
    }
  ];
  return (
    <div style={{marginLeft:20}}>
      <h3>认证信息</h3>
      <Descriptions column={2} colon=':' data={data}
        style={{width:'100%', marginLeft:20}} labelStyle={{ paddingRight:25 }}
      />
    </div>
  );
}

function AppInfo() {
  return (
    <>
      <BasicInfo></BasicInfo>
      <AuthInformation></AuthInformation>
      <Card style={{ width:500, height:80, marginTop:50, backgroundColor:'var(--color-fill-2)'}}>
        <Space size={80}>
          <div>此操作不可逆，请谨慎操作</div>
          <Button type='primary' status='danger' icon={<IconDelete />}>删除应用</Button>
        </Space>
      </Card>
    </>
  );
}

export default observer(AppInfo)