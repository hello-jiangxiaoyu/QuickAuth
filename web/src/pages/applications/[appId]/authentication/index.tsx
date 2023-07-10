import React from 'react';
import {Card, Select, Space, Tabs, Typography} from "@arco-design/web-react";

function Page() {
  const tables = [
    {key: 'database', title: '数据库', content: <div></div>},
    {key: 'social', title: '社会身份', content: <div></div>},
    {key: 'enterprise', title: '企业身份', content: <div></div>},
    {key: 'sms', title: '短信认证', content: <div></div>},
    {key: 'email', title: '邮件认证', content: <div></div>}
  ];

  return (
    <Card style={{minHeight:'80vh'}}>
      <Space style={{marginBottom:15}}>
        <h4>租户:</h4>
        <Select style={{width:200}}></Select>
      </Space>
      <Tabs>
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
