import React from 'react';
import {Card, Select, Space} from "@arco-design/web-react";

export default function Page() {


  return (
    <Card style={{minHeight:'80vh'}}>
      <Space style={{marginBottom:15}} size='medium'>
        <h4>租户:</h4>
        <Select style={{width:200}}></Select>
      </Space>
      <Space style={{marginBottom:15, marginLeft:20}} size='medium'>
        <h4>消息源:</h4>
        <Select style={{width:200}}></Select>
      </Space>
    </Card>
  );
}
