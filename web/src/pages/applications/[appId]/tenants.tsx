import React, {useState} from 'react';
import {Button, Card, Grid, Table, TableColumnProps} from "@arco-design/web-react";
import {useSelector} from "react-redux";
import {GlobalState} from "@/store/redux";

function Tenants() {
  const [visible, setVisible] = useState(false);
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
        <Button type='text' size='small'>登录租户</Button>
        <Button type='text' size='small' status='warning'>禁用</Button>
        <Button type='primary' size='small' status='danger'>删除</Button>
      </>
    )},
  ];

  return (
    <div style={{width:'97%', marginLeft:20}}>
      <Grid.Row className='grid-demo' style={{ marginBottom: 16 }}>
        <Grid.Col span={12}>
          <h3>客户端凭证</h3>
        </Grid.Col>
        <Grid.Col span={12}>
          <Button style={{float:'right'}} type="primary" onClick={()=>setVisible(true)}>Add</Button>
        </Grid.Col>
      </Grid.Row>
      <Table columns={columns} data={tenantList} />
    </div>
  );
}

export default Tenants;
