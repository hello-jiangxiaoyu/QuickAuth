import React from 'react';
import {Card, Select, Space} from "@arco-design/web-react";
import store from "@/store/mobx";

function Page() {
  return (
    <Card style={{minHeight:'80vh'}}>
      <Space style={{marginBottom:15, display: store.multiTenant ? 'flex':'none'}} size='medium'>
        <h4>租户:</h4>
        <Select></Select>
      </Space>
    </Card>
  );
}

export default Page;
