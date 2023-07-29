import React, {useState} from "react";
import {Input, Button, Space, Checkbox, InputNumber, Form, Grid} from '@arco-design/web-react';
import { IconDelete, IconEdit } from '@arco-design/web-react/icon';
import {useSelector} from "react-redux";
import {GlobalState} from "@/store/redux";

export default function LoginAuth() {
  const {currentTenant} = useSelector((state: GlobalState) => state);
  const dataSource = currentTenant.redirectUris;

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
            <Checkbox defaultChecked={currentTenant.isCode === 1}>authorization_code</Checkbox>
            <Checkbox defaultChecked={currentTenant.isRefresh === 1}>refresh_token</Checkbox>
            <Checkbox defaultChecked={currentTenant.isPassword === 1}>password</Checkbox>
            <Checkbox defaultChecked={currentTenant.isCredential === 1}>client_credentials</Checkbox>
            <Checkbox defaultChecked={currentTenant.isDeviceFlow === 1}>device_flow</Checkbox>
          </Space>
        </Grid.Col>

        <Grid.Row gutter={[24, 12]} style={{width:800}}>
          <Grid.Col span={12}>
            <Form.Item label='code过期时间:' layout='inline' style={{marginLeft:52}}>
              <InputNumber style={{ width: 100}} min={30} defaultValue={currentTenant.codeExpire} suffix='s'/>
            </Form.Item>
          </Grid.Col>
          <Grid.Col span={12}>
            <Form.Item label='id_token过期时间:' layout='inline' style={{marginLeft:32}}>
              <InputNumber style={{ width: 100}} min={30} defaultValue={currentTenant.idExpire} suffix='s'/>
            </Form.Item>
          </Grid.Col>
          <Grid.Col span={12}>
            <Form.Item label='access_token过期时间:' layout='inline'>
              <InputNumber style={{ width: 100}} min={30} defaultValue={currentTenant.accessExpire} suffix='s'/>
            </Form.Item>
          </Grid.Col>
          <Grid.Col span={12}>
            <Form.Item label='refresh_token过期时间:' layout='inline'>
              <InputNumber style={{ width: 100}} min={30} defaultValue={currentTenant.refreshExpire} suffix='s'/>
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

  return (
    <div style={{marginLeft:20, marginBottom:40}}>
      <h3>回调域名</h3>
      <Space size='small' style={{marginBottom:10}}>
        <Input style={{ width:800 }} allowClear  placeholder='请填写 HTTP/HTTPS 开头的 URL' />
        <Button type='primary'>添加</Button>
      </Space>

      {dataSource && dataSource.map((item)=>(
        <RedirectUriComponent key={item} uri={item}></RedirectUriComponent>
      ))}
      <ProtoConfig></ProtoConfig>
    </div>
  );
}
