import React, {useEffect, useState} from 'react';
import {Button, Card, Message, Select, Table, TableColumnProps} from "@arco-design/web-react";
import {Pool, User} from "@/http/users";
import api from "@/http/api";

function Page() {
  const [pools, setPools] = useState([] as Array<Pool>);
  const [users, setUsers] = useState([] as Array<User>);
  const [selectId, setSelectId] = useState(1);
  useEffect(() => {
    updateUserPool();
  }, []);
  useEffect(() => {
    updateUser();
  }, [selectId])

  function updateUserPool() {
    api.fetchUserPoolList().then(r => {
      setPools(r.data);
    }).catch();
  }
  function updateUser() {
    api.fetchUserList(selectId).then(r => {
      setUsers(r.data);
    }).catch();
  }

  function UserPoolSelect() {
    return (
      <Select
        dropdownMenuStyle={{ maxHeight: 400 }} value={selectId} onChange={setSelectId} bordered={false}
        triggerProps={{autoAlignPopupWidth: false, autoAlignPopupMinWidth: true, position: 'bl'}}
        style={{width:'fit-content', minWidth:120, maxWidth:250, backgroundColor:'var(--color-fill-2)'}}
      >
        {pools && pools.map((option) => (
          <Select.Option key={option.id} value={option.id} style={{height:50, textAlign:'left', display:'block'}}>
            {option.name}
          </Select.Option>
        ))}
      </Select>
    )
  }

  function onDeleteUser(record:User) {
    api.deleteUser(selectId, record.id).then(() => {
      Message.success("Success !")
      updateUserPool();
    }).catch(e => Message.error(e.toString()));
  }
  const columns: TableColumnProps[] = [
    {title: 'id', dataIndex: 'id', align:'center'},
    {title: 'username', dataIndex: 'username', align:'center'},
    {title: 'displayName', dataIndex: 'displayName', align:'center'},
    {title: 'isDisabled', dataIndex: 'isDisabled', align:'center'},
    {title: '操作', dataIndex: 'op', align:'center', render: (_, record) => (
        <>
          <Button type='text' size='small' status='warning'>禁用</Button>
          <Button type='primary' size='small' status='danger' onClick={()=>onDeleteUser(record)}>删除</Button>
        </>
      )},
  ];

  return (
    <Card style={{minHeight:'80vh', width:'100%'}}>
      用户池: <UserPoolSelect></UserPoolSelect>
      <Table columns={columns} data={users} style={{marginTop:30}}/>
    </Card>
  );
}

export default Page;

