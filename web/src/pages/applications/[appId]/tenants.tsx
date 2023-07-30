import React from 'react';
import {Button, Card, Table, TableColumnProps} from "@arco-design/web-react";
import {useSelector} from "react-redux";
import {GlobalState} from "@/store/redux";

function Tenants() {
  const tenantList = useSelector((state: GlobalState) => state.tenantList);
  const columns: TableColumnProps[] = [
    {title: '租户ID', dataIndex: 'id', align:'center'},
    {title: '租户名', dataIndex: 'name', align:'center'},
    {title: '类型ID', dataIndex: 'type', align:'center'},
    {title: '用户池ID', dataIndex: 'userPoolId', align:'center'},
    {title: 'Host', dataIndex: 'host', align:'center'},
    {title: '公司名', dataIndex: 'company', align:'center'},
    {title: '操作', dataIndex: 'op', align:'center', render: () => (
        <>
          <Button type='primary' size='small' status='danger'>删除</Button>
          <Button type='text' size='small' status='warning'>禁用</Button>
          <Button type='text' size='small'>登录租户</Button>
        </>
      )},
  ];

  return (
    <Card style={{minHeight:'80vh', width:'100%'}}>
      <Table columns={columns} data={tenantList} />
    </Card>
  );
}

export default Tenants;
