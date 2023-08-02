import React, {useEffect, useState} from 'react';
import {Button, Message, Table, TableColumnProps, Grid} from "@arco-design/web-react";
import Secret, {deleteSecret} from "@/http/secret";
import {useRouter} from "next/router";
import {getRouterPara} from "@/utils/stringTools";
import api from "@/http/api";
import CreateSecretDialog from "@/components/Dialog/secret";

function ClientCredential() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const [secrets, setSecrets] = useState([] as Array<Secret>);

  function onDeleteSecret(record:Secret) {
    deleteSecret(appId, record.id).then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        Message.success("Success !")
        api.fetchSecretList(appId).then(r => {
          if (r.code !== 200) {Message.error(r.msg)} else {
            setSecrets(r.data);
          }
        })
      }
    })
  }

  const columns: TableColumnProps[] = [
    {title: 'App Id', dataIndex: 'appId', align:'center'},
    {title: '密钥', dataIndex: 'secret', align:'center'},
    {title: '创建时间', dataIndex: 'createdAt', align:'center'},
    {title: '操作', dataIndex: 'op', align:'center', render: (_, record) => (
      <>
        <Button type='text' size='small' status='danger' onClick={()=>onDeleteSecret(record)}>删除</Button>
        <Button type='text' size='small'>权限</Button>
      </>
      )},
  ];

  useEffect(() => {
    api.fetchSecretList(appId).then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        setSecrets(r.data);
      }
    })
  }, [appId]);

  const [visible, setVisible] = useState(false);
  return (
    <>
      <Grid.Row className='grid-demo' style={{ marginBottom: 16 }}>
        <Grid.Col span={12}>
          <h3>客户端凭证</h3>
        </Grid.Col>
        <Grid.Col span={12}>
          <Button style={{float:'right'}} type="primary" onClick={()=>setVisible(true)}>Add</Button>
        </Grid.Col>
      </Grid.Row>
      <Table columns={columns} data={secrets} />
      <CreateSecretDialog visible={visible} setVisible={setVisible}></CreateSecretDialog>
    </>
  );
}

export default ClientCredential;
