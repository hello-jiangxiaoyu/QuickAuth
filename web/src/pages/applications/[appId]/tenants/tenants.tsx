import React from 'react';
import {Button, Card, Table, TableColumnProps} from "@arco-design/web-react";
import Link from "next/link";

function Tenants(props:{appId:string}) {
  const columns: TableColumnProps[] = [
    {title: '租户名', dataIndex: 'name', align:'center', render:(_, record)=>(
        <Link href={`/applications/${props.appId}/tenants/${record.name}`}><a>{record.name}</a></Link>
      )},
    {title: '租户ID', dataIndex: 'tenantID', align:'center'},
    {title: '来源', dataIndex: 'from', align:'center'},
    {title: '用户池', dataIndex: 'userPoolID', align:'center'},
    {title: '操作', dataIndex: 'op', align:'center', render: () => (
        <>
          <Button type='primary' size='small' status='danger'>删除</Button>
          <Button type='text' size='small' status='warning'>禁用</Button>
          <Button type='text' size='small'>登录租户</Button>
        </>
      )},
  ];
  const data = [
    {key: '1', name: 'jiang', tenantID: 'e22d60b37fce2de059208602e4d8237f', from: 'default', userPoolID:1},
    {key: '2', name: 'zhao', tenantID: 'ri08JMbndpI05SDOYqFl_9VKuyUvcvTahR', from: 'm2m', userPoolID:1},
    {key: '3', name: 'feng', tenantID: '91b6dc86ee44d2f689e8beb4a447dfbd', from: 'm2m', userPoolID:1},
    {key: '4', name: 'gateway', tenantID: '67a339640d86c5ddcc74f30e93ce3326', from: 'm2m', userPoolID:1},
  ];

  return (
    <Card style={{minHeight:'80vh', width:'100%'}}>
      <Table columns={columns} data={data} />
    </Card>
  );
}

export default Tenants;
