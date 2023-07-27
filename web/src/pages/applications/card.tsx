import React, {useState } from 'react';
import Router from "next/router"
import {Button, Popover, Card, Space, Modal, Form, Input, Select, Message} from '@arco-design/web-react';
import MyIcon from "@/components/Widget/StringIcon";
import {IconPlusCircle} from "@arco-design/web-react/icon";
import {App, createApp, deleteApp, fetchAppList} from "@/http/app";
import {apps} from "@/store/mobx";

// application card with dynamic icon
export default function MyCard(props: { appId: string, name: string, type: string, icon?: string}) {
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
        deleteApp(props.appId).then( r => {
          if (r.code !== 200) {Message.error(r.msg)} else {
            Message.success('Delete success !');
            fetchAppList().then(r => {
              if (r.code !== 200) {Message.error("Get app list err: " + r.msg)} else {
                apps.setAppList(r.data)
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


export function AddApp() {
  const [visible, setVisible] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [form] = Form.useForm();


  function onOk() {
    form.validate().then((app:App) => {
      setConfirmLoading(true);
      createApp(app).then(r => {
        if (r.code !== 200) {Message.error(r.msg)} else {
          Message.success('Success !');
          fetchAppList().then(r => {
            if (r.code !== 200) {Message.error("Get app list err: " + r.msg)} else {
              apps.setAppList(r.data)
            }
          })
          setVisible(false);
        }
        setConfirmLoading(false);
      }).catch()

    }).catch((err) => {
      Message.error(err.toString());
    });
  }

  return (
    <Popover content={"Add App"} position='bottom'>
      <Card hoverable style={{ width:330, height: 180, marginLeft:12, cursor:'pointer',
        boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 12px 0 rgba(0, 0, 0, 0.19)' }} onClick={() => setVisible(true)}
      >
        <IconPlusCircle style={{color:"#2f6af1", width:60, height: 60, position:'absolute', top:'30%', left:'38%'}}></IconPlusCircle>
      </Card>
      <Modal title='Create app' visible={visible} onOk={onOk} style={{width: 600}}
        confirmLoading={confirmLoading} onCancel={() => setVisible(false)}
      >
        <Form form={form} labelCol={{style: { flexBasis: 100 }}}
          wrapperCol={{style: { flexBasis: 'calc(100% - 100px)' }}} initialValues={{tag:'Single Tenant'}}
        >
          <Form.Item label='Name' field='name' rules={[{ required: true }]}>
            <Input placeholder='app name' />
          </Form.Item>
          <Form.Item label='Type' required field='tag' rules={[{ required: false }]}>
            <Select options={["single_tenant", "multi_tenant"]}/>
          </Form.Item>
          <Form.Item label='Icon' required field='icon' rules={[{ required: true }]}>
            <Input placeholder='icon url' />
          </Form.Item>
          <Form.Item label='Describe' required field='describe' rules={[{ required: false }]}>
            <Input.TextArea placeholder='' />
          </Form.Item>
        </Form>
      </Modal>
    </Popover>
  );
}
