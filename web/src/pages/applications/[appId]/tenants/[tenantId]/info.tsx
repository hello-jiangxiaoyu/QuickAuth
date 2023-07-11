import React from "react";
import {Card, Descriptions, Link} from "@arco-design/web-react";

function TenantInfo(props:{appId:string}) {
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
      label: '登出接口',
      value: 'https://jiangzhaofeng.online/oidc/session/end',
    }
  ];
  return (
    <Card style={{minHeight:'80vh', width:'100%'}}>
      <Descriptions
        column={2} colon=':' title={<h3>认证信息</h3>} data={data}
        style={{width:'90%'}} labelStyle={{ paddingRight:25 }}
      />
    </Card>
  );
}

export default TenantInfo;
