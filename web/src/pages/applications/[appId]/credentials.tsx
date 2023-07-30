import React, {useEffect, useState} from 'react';
import {Button, Message, Table, TableColumnProps} from "@arco-design/web-react";
import Secret from "@/http/secret";
import {useRouter} from "next/router";
import {getRouterPara} from "@/utils/stringTools";
import api from "@/http/api";

function ClientCredential() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const [secrets, setSecrets] = useState([] as Array<Secret>);
  const columns: TableColumnProps[] = [
    {title: 'App Id', dataIndex: 'appId', align:'center'},
    {title: '密钥', dataIndex: 'secret', align:'center'},
    {title: '创建时间', dataIndex: 'createdAt', align:'center'},
    {title: '操作', dataIndex: 'op', align:'center', render: () => (
      <>
        <Button type='text' size='small' status='danger'>删除</Button>
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

  return (
    <>
      <h3>客户端凭证</h3>
      <Table columns={columns} data={secrets} />
    </>
  );
}

export default ClientCredential;
