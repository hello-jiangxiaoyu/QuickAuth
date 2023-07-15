import React from 'react';
import {Button, Table, TableColumnProps} from "@arco-design/web-react";

function ClientCredential(props:{appId:string}) {
  const columns: TableColumnProps[] = [
    {title: '名称', dataIndex: 'name', align:'center'},
    {title: '密钥', dataIndex: 'clientSecret', align:'center'},
    {title: '创建时间', dataIndex: 'create_time', align:'center'},
    {title: '操作', dataIndex: 'op', align:'center', render: () => (
      <>
        <Button type='text' size='small' status='danger'>删除</Button>
        <Button type='text' size='small'>权限</Button>
      </>
      )},
  ];
  const data = [
    {key: '1', name: 'Gateway', clientSecret: 'e22d60b37fce2de059208602e4d8237f', create_time: '2023-07-10 10:30:20'},
    {key: '2', name: 'MES', clientSecret: 'ri08JMndpI05DOYqFl__9VKuyUvcvTahR', create_time: '2023-07-11 10:01:34'},
    {key: '3', name: 'OA', clientSecret: '91b6dc86ee44d2f689e8beb4a447dfbd', create_time: '2023-07-15 10:50:04'},
    {key: '4', name: 'Project', clientSecret: '67a339640d86c5ddcc74f30e93ce3326', create_time: '2023-07-15 22:11:38'},
  ];
  return (
    <>
      <h3>客户端凭证</h3>
      <Table columns={columns} data={data} />
    </>
  );
}

export default ClientCredential;
