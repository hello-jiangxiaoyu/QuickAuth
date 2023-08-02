import React, {useState} from 'react'
import {Form, Input, Message, Modal, Select} from "@arco-design/web-react";
import api from "@/http/api";
import Secret from "@/http/secret";
import {useRouter} from "next/router";
import {getRouterPara} from "@/utils/stringTools";

export default function CreateSecretDialog(props:{visible: boolean, setVisible: React.Dispatch<React.SetStateAction<boolean>>}) {
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [form] = Form.useForm();
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);

  function onOk() {
    form.validate().then((secret:Secret) => {
      setConfirmLoading(true);
      api.createSecret(appId, secret).then(r => {
        if (r.code !== 200) {Message.error(r.msg)} else {
          Message.success('Success !');
        }
        setConfirmLoading(false);
        props.setVisible(false);
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
        <Form.Item label='Scope' required field='scope' rules={[{ required: false }]}>
          <Select mode='multiple' options={["read_user", "modify_user", "delete_user"]}/>
        </Form.Item>
        <Form.Item label='Describe' required field='describe' rules={[{ required: true }]}>
          <Input.TextArea placeholder='' />
        </Form.Item>
      </Form>
    </Modal>
  )
}
