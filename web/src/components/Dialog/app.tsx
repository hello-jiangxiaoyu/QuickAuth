import React, {useState} from 'react';
import {Form, Input, Message, Modal, Select} from "@arco-design/web-react";
import App from "@/http/app";
import api from "@/http/api";
import {dispatchAppList} from "@/store/redux";

export default function CreateAppDialog(props:{visible: boolean, setVisible: React.Dispatch<React.SetStateAction<boolean>>}) {
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [form] = Form.useForm();

  function onOk() {
    form.validate().then((app:App) => {
      setConfirmLoading(true);
      api.createApp(app).then(r => {
        if (r.code !== 200) {Message.error(r.msg)} else {
          Message.success('Success !');
          api.fetchAppList().then(r => {
            if (r.code !== 200) {Message.error(r.msg)} else {
              dispatchAppList(r.data)
            }
          })
          props.setVisible(false);
        }
        setConfirmLoading(false);
      }).catch();
    }).catch((err) => {
      Message.error(err.toString());
    });
  }

  return (
    <Modal title='Create app' visible={props.visible} onOk={onOk} style={{width: 600}}
           confirmLoading={confirmLoading} onCancel={() => props.setVisible(false)}
    >
      <Form form={form} labelCol={{style: { flexBasis: 100 }}}
            wrapperCol={{style: { flexBasis: 'calc(100% - 100px)' }}} initialValues={{tag:'Single Tenant'}}
      >
        <Form.Item label='Name' field='name' rules={[{ required: true }]}>
          <Input placeholder='app name' />
        </Form.Item>
        <Form.Item label='Type' required field='tag' rules={[{ required: false }]}>
          <Select options={["Single Tenant", "Multi Tenant"]}/>
        </Form.Item>
        <Form.Item label='Host' required field='host' rules={[{ required: false }]}>
          <Input placeholder='do not start with http:// or https://' />
        </Form.Item>
        <Form.Item label='Icon' required field='icon' rules={[{ required: true }]}>
          <Input placeholder='icon url' />
        </Form.Item>
        <Form.Item label='Describe' required field='describe' rules={[{ required: true }]}>
          <Input.TextArea placeholder='' />
        </Form.Item>
      </Form>
    </Modal>
  )
}
