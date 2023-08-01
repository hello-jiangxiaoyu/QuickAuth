import React, {useState} from 'react';
import {Form, Input, Message, Modal} from "@arco-design/web-react";
import App from "@/http/app";
import api from "@/http/api";
import {dispatchTenantList, GlobalState} from "@/store/redux";
import {useSelector} from "react-redux";

export default function CreateTenantDialog(props:{visible: boolean, setVisible: React.Dispatch<React.SetStateAction<boolean>>}) {
  const [confirmLoading, setConfirmLoading] = useState(false);

  const currentApp = useSelector((state: GlobalState) => state.currentApp);
  function onOk() {
    form.validate().then((app:App) => {
      setConfirmLoading(true);
      api.createApp(app).then(r => {
        if (r.code !== 200) {Message.error(r.msg)} else {
          Message.success('Success !');
          api.fetchTenantList(currentApp.id).then(r => {
            if (r.code !== 200) {Message.error(r.msg)} else {
              dispatchTenantList(r.data);
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

  const [form] = Form.useForm();
  return (
    <Modal title='Create tenant' visible={props.visible} onOk={onOk} style={{width: 600}}
           confirmLoading={confirmLoading} onCancel={() => props.setVisible(false)}
    >
      <Form form={form} labelCol={{style: { flexBasis: 100 }}}
            wrapperCol={{style: { flexBasis: 'calc(100% - 100px)' }}} initialValues={{tag:'Single Tenant'}}
      >
        <Form.Item label='Name' field='name' rules={[{ required: true }]}>
          <Input placeholder='app name' />
        </Form.Item>
      </Form>
    </Modal>
  )
}
