import React from 'react';
import {Button, Table, TableColumnProps} from "@arco-design/web-react";

function ClientCredential(props:{appId:string}) {
  const columns: TableColumnProps[] = [
    {title: '名称', dataIndex: 'name', align:'center'},
    {title: '密钥', dataIndex: 'clientSecret', align:'center'},
    {title: 'Address', dataIndex: 'address', align:'center'},
    {title: '操作', dataIndex: 'op', align:'center', render: () => (
      <>
        <Button type='text' size='small' status='danger'>删除</Button>
        <Button type='text' size='small'>协议</Button>
        <Button type='text' size='small'>权限</Button>
      </>
      )},
  ];
  const data = [
    {key: '1', name: 'Gateway', clientSecret: 'e22d60b37fce2de059208602e4d8237f', address: '32 Park Road, London'},
    {key: '2', name: 'MES', clientSecret: 'ri08JMbndpI05SDOYqFl__9VKuyUvcvTahR', address: '35 Park Road, London'},
    {key: '3', name: 'OA', clientSecret: '91b6dc86ee44d2f689e8beb4a447dfbd', address: '31 Park Road, London'},
    {key: '4', name: 'Project', clientSecret: '67a339640d86c5ddcc74f30e93ce3326', address: '42 Park Road, London'},
  ];
  return (
    <>
      <h3>客户端凭证</h3>
      <Table columns={columns} data={data} />
      <h3>设备流</h3>
      <Table columns={columns} data={data} />
    </>
  );
}

export default ClientCredential;
