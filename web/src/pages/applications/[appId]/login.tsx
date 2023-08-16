import React, {useState} from "react";
import {Button, Checkbox, Form, Grid, Input, InputNumber, Message, Space} from '@arco-design/web-react';
import {IconDelete, IconEdit} from '@arco-design/web-react/icon';
import {useSelector} from "react-redux";
import {dispatchTenant, GlobalState} from "@/store/redux";
import {TenantDetail} from "@/http/tenant";
import api from "@/http/api";

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
    const [form] = Form.useForm();
    const {currentApp, currentTenant} = useSelector((state: GlobalState) => state);
    function onSave() {
      form.validate().then((value:TenantDetail) => {
        const res = {} as TenantDetail;
        Object.assign(res, currentTenant);
        res.isCode = value.isCode;
        res.isRefresh = value.isRefresh;
        res.isPassword = value.isPassword;
        res.isCredential = value.isCredential;
        res.isDeviceFlow = value.isDeviceFlow;
        res.codeExpire = value.codeExpire;
        res.idExpire = value.idExpire;
        res.accessExpire = value.accessExpire;
        res.refreshExpire = value.refreshExpire;
        if (res === currentTenant) {
          Message.success('Success !');
          return
        }
        api.modifyTenant(currentTenant.appId, currentTenant.id, res).then(() => {
          Message.success('Success !');
          api.fetchTenant(currentApp.id, currentTenant.id).then(r => {dispatchTenant(r.data)})
        }).catch();
      }).catch((e) => Message.error("validator " + e.toString()));
    }

    return (
      <Form form={form} initialValues={currentTenant} style={{marginTop:50}} layout='inline'>
        <h3>协议配置</h3>
        <Grid.Col style={{ marginBottom:20, marginTop:10}}>
          <Space size={40}>
            <Form.Item label='' field='isCode' layout='inline' >
              <Checkbox defaultChecked={currentTenant.isCode}>authorization_code</Checkbox>
            </Form.Item>
            <Form.Item label='' field='isRefresh' layout='inline'>
              <Checkbox defaultChecked={currentTenant.isRefresh}>refresh_token</Checkbox>
            </Form.Item>
            <Form.Item label='' field='isPassword' layout='inline'>
              <Checkbox defaultChecked={currentTenant.isPassword}>password</Checkbox>
            </Form.Item>
            <Form.Item label='' field='isCredential' layout='inline'>
             <Checkbox defaultChecked={currentTenant.isCredential}>client_credentials</Checkbox>
            </Form.Item>
            <Form.Item label='' field='isDeviceFlow' layout='inline'>
              <Checkbox defaultChecked={currentTenant.isDeviceFlow}>device_flow</Checkbox>
            </Form.Item>
          </Space>
        </Grid.Col>

        <Grid.Row gutter={[24, 12]} style={{width:800}}>
          <Grid.Col span={12}>
            <Form.Item label='code过期时间:' field='codeExpire' layout='inline' style={{marginLeft:52}}>
              <InputNumber style={{ width: 100}} min={30} suffix='s'/>
            </Form.Item>
          </Grid.Col>
          <Grid.Col span={12}>
            <Form.Item label='id_token过期时间:' field='idExpire' layout='inline' style={{marginLeft:32}}>
              <InputNumber style={{ width: 100}} min={30} suffix='s'/>
            </Form.Item>
          </Grid.Col>
          <Grid.Col span={12}>
            <Form.Item label='access_token过期时间:' field='accessExpire' layout='inline'>
              <InputNumber style={{ width: 100}} min={30} suffix='s'/>
            </Form.Item>
          </Grid.Col>
          <Grid.Col span={12}>
            <Form.Item label='refresh_token过期时间:' field='refreshExpire' layout='inline'>
              <InputNumber style={{ width: 100}} min={30} suffix='s'/>
            </Form.Item>
          </Grid.Col>
        </Grid.Row>

        <Grid.Col style={{marginTop:30}}>
          <Form.Item>
            <Space size='medium'>
              <Button type='primary' onClick={onSave}>保存</Button>
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
