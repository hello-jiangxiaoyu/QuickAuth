import React, {useState} from "react";
import {Input, Button, Space, List, Grid } from '@arco-design/web-react';
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
    <Grid.Row>
      {edit ? <> {/*编辑状态*/}
        <Input style={{ width:650, height:30 }} allowClear  placeholder='请填写 HTTP/HTTPS 开头的 URL' defaultValue={props.uri} />
        <Button type='primary' style={{marginLeft:10}} onClick={onClickSave} >保存</Button>
        <Button type='secondary' style={{marginLeft:10}} onClick={onClickCansel} >取消</Button>
      </> : <>
        <div style={{width:600, height:30}}>{props.uri}</div>
        <Button type='text' style={{marginLeft:10}} icon={<IconEdit />} onClick={onClickEdit} />
        <Button type='text' icon={<IconDelete />} onClick={onClickDelete} />
      </>}
    </Grid.Row>
  );
}

export default function LoginAuth(props: {appId: string}) {
  const dataSource = new Array(4).fill({
    title: 'Beijing Bytedance Technology Co., Ltd.',
    description: 'Beijing ByteDance Technology Co., Ltd. is an enterprise located in China.',
  });

  return (
    <>
      <h3>回调域名</h3>
      <Space size='small'>
        <Input style={{ width:800 }} allowClear  placeholder='请填写 HTTP/HTTPS 开头的 URL' />
        <Button type='primary'>添加</Button>
      </Space>

      <List style={{ width:900, marginTop:20 }} size='small' dataSource={dataSource} split={false} render={(item, index) => (
          <List.Item key={index}>
            <RedirectUriComponent uri='https://jiangzhaofeng.online/login'></RedirectUriComponent>
          </List.Item>
        )}
      />
    </>
  );
}
