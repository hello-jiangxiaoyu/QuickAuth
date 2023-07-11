import React, {useState} from "react";
import {Input, Button, Space } from '@arco-design/web-react';
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
    <div style={{marginTop:40}}>
      <h3>协议配置</h3>
    </div>
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
    <div style={{marginLeft:20}}>
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
