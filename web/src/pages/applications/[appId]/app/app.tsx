import {Button, Card, Form, Input, Space} from "@arco-design/web-react";
import React from "react";
import {IconDelete} from "@arco-design/web-react/icon";

function BasicInfo() {
  return (
    <Form style={{ width:600 }} autoComplete='off'>
      <h3>基本信息</h3>
      <Form.Item label={'应用名称'} rules={[{ required: true }]}>
        <Input placeholder='please enter your app name...' />
      </Form.Item>
      <Form.Item label={'应用描述'}>
        <Input.TextArea placeholder='please enter app description...' />
      </Form.Item>
      <Form.Item wrapperCol={{ offset: 5 }}>
        <Space size='medium'>
          <Button type='primary'>保存</Button>
          <Button type='secondary'>重置</Button>
        </Space>
      </Form.Item>
    </Form>
  );
}

export default function AppInfo(props: {appId: string}) {
  return (
    <>
      <BasicInfo></BasicInfo>

      <Card style={{ width:500, height:80, marginTop:60, marginBottom:50, backgroundColor:'var(--color-fill-2)'}}>
        <Space size={80}>
          <div>此操作不可逆，请谨慎操作</div>
          <Button type='primary' status='danger' icon={<IconDelete />}>删除应用</Button>
        </Space>
      </Card>
    </>
  );
}
