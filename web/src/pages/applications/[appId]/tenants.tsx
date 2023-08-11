import React, {useState} from 'react';
import {Button, Grid, Message, Table, TableColumnProps} from "@arco-design/web-react";
import {useSelector} from "react-redux";
import {dispatchTenantList, GlobalState} from "@/store/redux";
import {Tenant} from "@/http/tenant";
import api from "@/http/api";
import {useRouter} from "next/router";
import {getRouterPara} from "@/utils/stringTools";
import CreateTenantDialog from "@/components/Dialog/Tenant";

function Tenants() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const [visible, setVisible] = useState(false);
  const {tenantList, currentApp} = useSelector((state: GlobalState) => state);
  const columns: TableColumnProps[] = [
    {title: '序号', dataIndex: 'key', align:'center'},
    {title: '租户ID', dataIndex: 'id', align:'center'},
    {title: '租户名', dataIndex: 'name', align:'center'},
    {title: '类型ID', dataIndex: 'type', align:'center'},
    {title: '用户池ID', dataIndex: 'userPoolId', align:'center'},
    {title: 'Host', dataIndex: 'host', align:'center'},
    {title: '公司名', dataIndex: 'company', align:'center'},
    {title: '操作', dataIndex: 'op', align:'center', render: (_, record) => (
      <>
        <Button type='text' size='small'>登录租户</Button>
        <Button type='text' size='small' status='warning'>禁用</Button>
        {currentApp.tag === 'Multi Tenant' &&
          <Button type='primary' size='small' status='danger' onClick={() => onDeleteTenant(record)}>删除</Button>
        }
      </>
    )},
  ];

  function updateTenantList() {
    api.fetchTenantList(appId).then(r => {
      r.data.forEach((obj, index) => {
        obj.key = index + 1;
      });
      dispatchTenantList(r.data);
    }).catch();
  }
  function onDeleteTenant(record:Tenant) {
    api.deleteTenant(appId, record.id).then(() => {
      Message.success("Success !");
      updateTenantList();
    }).catch()
  }

  return (
    <div style={{width:'97%', marginLeft:20}}>
      <Grid.Row className='grid-demo' style={{ marginBottom: 16 }}>
        <Grid.Col span={12}>
          <h3>客户端凭证</h3>
        </Grid.Col>
        {currentApp.tag === 'Multi Tenant' && <Grid.Col span={12}>
          <Button style={{float:'right'}} type="primary" onClick={()=>setVisible(true)}>Add</Button>
        </Grid.Col>}
      </Grid.Row>
      <Table columns={columns} data={tenantList} />
      <CreateTenantDialog visible={visible} setVisible={setVisible}/>
    </div>
  );
}

export default Tenants;
