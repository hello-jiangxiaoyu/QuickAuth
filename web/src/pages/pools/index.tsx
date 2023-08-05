import React, {useEffect, useState} from 'react';
import {Button, Card, Message, Table, TableColumnProps} from "@arco-design/web-react";
import api from "@/http/api";
import {Pool} from "@/http/users";

function Page() {
  const [pools, setPools] = useState([] as Array<Pool>);
  useEffect(() => {
    updateUserPool();
  }, []);

  function updateUserPool() {
    api.fetchUserPoolList().then(r => {
      setPools(r.data);
    }).catch(e => Message.error(e.toString()));
  }
  function onDeletePool(record:Pool) {
    api.deleteUserPool(record.id).then(() => {
      Message.success("Success !")
      updateUserPool();
    }).catch(e => Message.error(e.toString()));
  }

  const columns: TableColumnProps[] = [
    {title: 'id', dataIndex: 'id', align:'center'},
    {title: 'name', dataIndex: 'name', align:'center'},
    {title: 'describe', dataIndex: 'describe', align:'center'},
    {title: '操作', dataIndex: 'op', align:'center', render: (_, record) => (
        <>
          <Button type='text' size='small'>关联租户</Button>
          <Button type='primary' size='small' status='danger' onClick={()=>onDeletePool(record)}>删除</Button>
        </>
      )},
  ];

  return (
    <Card style={{minHeight:'80vh', width:'100%'}}>
      <h3>用户池列表</h3>
      <Table columns={columns} data={pools} />
    </Card>
  );
}

export default Page;

