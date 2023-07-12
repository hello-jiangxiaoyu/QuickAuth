import React, {useState} from "react";
import {Input, Button, Space, Checkbox, InputNumber, Form, Grid} from '@arco-design/web-react';
import { IconDelete, IconEdit } from '@arco-design/web-react/icon';

function RedirectUriComponent(props: {uri: string}) {
  const [edit, setEdit] = useState(false);
  function onClickEdit() {
    setEdit(true);
  }
  function onClickCansel() {
    setEdit(false);
  }
  function onClickDelete() {
    setEdit(false);
  }
  function onClickSave() {
    setEdit(false);
  }
  return (
    <div style={{marginTop:10}}>
      {edit ? <Space> {/*编辑状态*/}
        <Input style={{ width:650, height:30 }} allowClear  placeholder='请填写 HTTP/HTTPS 开头的 URL' defaultValue={props.uri} />
        <Button type='primary' style={{marginLeft:10}} onClick={onClickSave} >保存</Button>
        <Button type='secondary' style={{marginLeft:10}} onClick={onClickCansel} >取消</Button>
      </Space> : <Space>
        <div style={{width:600, height:30}}>{props.uri}</div>
        <Button type='text' style={{marginLeft:10}} icon={<IconEdit />} onClick={onClickEdit} />
        <Button type='text' icon={<IconDelete />} onClick={onClickDelete} />
      </Space>}
    </div>
  );
}

function ProtoConfig() {
  return (
    <Form style={{marginTop:50}} layout='inline'>
      <h3>协议配置</h3>
      <Grid.Col style={{ marginBottom:20, marginTop:10}}>
        <Space size={40}>
          <Checkbox>authorization_code</Checkbox>
          <Checkbox>refresh_token</Checkbox>
          <Checkbox>password</Checkbox>
          <Checkbox>client_credentials</Checkbox>
          <Checkbox>device_flow</Checkbox>
        </Space>
      </Grid.Col>

      <Grid.Row gutter={[24, 12]} style={{width:800}}>
        <Grid.Col span={12}>
          <Form.Item label='code过期时间:' layout='inline' style={{marginLeft:52}}>
            <InputNumber style={{ width: 100}} min={30} defaultValue={120} suffix='s'/>
          </Form.Item>
        </Grid.Col>
        <Grid.Col span={12}>
          <Form.Item label='id_token过期时间:' layout='inline' style={{marginLeft:32}}>
            <InputNumber style={{ width: 100}} min={30} defaultValue={1209600} suffix='s'/>
          </Form.Item>
        </Grid.Col>
        <Grid.Col span={12}>
          <Form.Item label='access_token过期时间:' layout='inline'>
            <InputNumber style={{ width: 100}} min={30} defaultValue={1209600} suffix='s'/>
          </Form.Item>
        </Grid.Col>
        <Grid.Col span={12}>
          <Form.Item label='refresh_token过期时间:' layout='inline'>
            <InputNumber style={{ width: 100}} min={30} defaultValue={2592000} suffix='s'/>
          </Form.Item>
        </Grid.Col>
      </Grid.Row>

      <Grid.Col style={{marginTop:30}}>
        <Form.Item>
          <Space size='medium'>
            <Button type='primary'>保存</Button>
            <Button type='secondary'>重置</Button>
          </Space>
        </Form.Item>
      </Grid.Col>
    </Form>
  )
}

export default function LoginAuth(props: {appId: string}) {
  const dataSource = [
    {key:1, name: 'jiang', uri: 'https://jiangzhaofeng.online/login'},
    {key:2, name: 'jiang', uri: 'https://jiangzhaofeng.online/login'},
    {key:3, name: 'jiang', uri: 'https://jiangzhaofeng.online/login'},
    {key:4, name: 'jiang', uri: 'https://jiangzhaofeng.online/login'},
  ];

  return (
    <div style={{marginLeft:20, marginBottom:40}}>
      <h3>回调域名</h3>
      <Space size='small' style={{marginBottom:10}}>
        <Input style={{ width:800 }} allowClear  placeholder='请填写 HTTP/HTTPS 开头的 URL' />
        <Button type='primary'>添加</Button>
      </Space>

      {dataSource.map((item)=>(
        <RedirectUriComponent key={item.key} uri={item.uri}></RedirectUriComponent>
      ))}
      <ProtoConfig></ProtoConfig>
    </div>
  );
}
