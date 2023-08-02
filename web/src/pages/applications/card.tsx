import React, {useState } from 'react';
import Router from "next/router"
import {Button, Popover, Card, Space, Modal, Message} from '@arco-design/web-react';
import MyIcon from "@/components/Widget/StringIcon";
import {IconPlusCircle} from "@arco-design/web-react/icon";
import {dispatchAppList} from "@/store/redux";
import api from "@/http/api";
import CreateAppDialog from "@/components/Dialog/app";

// 应用展示选项卡
export default function ApplicationCard(props: { appId: string, name: string, type: string, icon?: string}) {
  let icon = 'IconCodeSandbox';
  if (typeof props.icon == 'string' && props.icon.startsWith('Icon')) {
    icon = props.icon;
  }

  function onClickCard() {
    Router.push(`applications/${props.appId}`).then();
  }

  function confirm() {
    Modal.confirm({title: 'Confirm deletion',
      content: 'Are you sure you want to delete the app.',
      okButtonProps: {status: 'danger'},
      onOk: () => {
        api.deleteApp(props.appId).then( r => { // 删除app
          if (r.code !== 200) {Message.error(r.msg)} else {
            Message.success('Delete success !');
            api.fetchAppList().then(r => {  // 刷新页面app列表
              if (r.code !== 200) {Message.error(r.msg)} else {
                dispatchAppList(r.data)
              }
            })
          }
        })
      },
    });
  }

  function MoreButton(props: { appId: string}) {
    return (
      <Popover position='top'
        content={
          <Button type='text' status='danger' style={{height:25}} key={props.appId} onClick={confirm}>
            删除应用
          </Button>
        }
      >
        <Button type={'text'} style={{height:25}} key={props.appId}>
          <MyIcon name={'IconMore'}/>
        </Button>
      </Popover>
    );
  }

  return (
    <>
      <Card hoverable actions={[<MoreButton key={1} appId={props.appId}></MoreButton>]}
        style={{ width:330, height: 180, cursor:'pointer', boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 12px 0 rgba(0, 0, 0, 0.19)' }}
      >
        <div onClick={onClickCard}>
          <Space >
            <MyIcon name={icon} color={'#2f6af1'} size={50}/>
            <div style={{margin:20}}>
              <h2>{props.name}</h2>
              <div>{props.type}</div>
            </div>
          </Space>
        </div>
      </Card>
    </>
  );
}

// 创建一个应用
export function AddApplication() {
  const [visible, setVisible] = useState(false);
  return (
    <Popover content={"Add App"} position='bottom'>
      <Card hoverable style={{ width:330, height: 180, marginLeft:12, cursor:'pointer',
        boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 12px 0 rgba(0, 0, 0, 0.19)' }} onClick={() => setVisible(true)}
      >
        <IconPlusCircle style={{color:"#2f6af1", width:60, height: 60, position:'absolute', top:'30%', left:'38%'}}></IconPlusCircle>
      </Card>
      <CreateAppDialog visible={visible} setVisible={setVisible}></CreateAppDialog>
    </Popover>
  );
}
