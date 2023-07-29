import React, {useState} from 'react';
import {Button, Divider, Select, Space, Typography} from "@arco-design/web-react";
import {IconPlus} from "@arco-design/web-react/icon";
import Router, {useRouter} from "next/router";
import {getRouterPara, replaceUriAppId} from "@/utils/stringTools";
import {useSelector} from "react-redux";
import {GlobalState} from "@/store/redux";


function ApplicationSelector() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const visibility = appId === '' ? 'hidden' : 'visible';
  const [appName, setAppName] = useState('');
  const {appList, tenantList, currentApp} = useSelector((state: GlobalState) => state);
  const multiTenant = currentApp?.tag === 'multiTenant';
  const [tenantName, setTenantName] = useState('');

  function onAppChange(value: string) {
    setAppName(value);
    const newUri = replaceUriAppId(value, router.asPath);
    if (newUri === router.asPath) {
      return
    }
    Router.push(newUri).then();
  }

  const addItem = () => {console.log('add item')};
  function CreateItem(props:{text:string}) {
    return (
      <div style={{display: 'flex', alignItems: 'center', padding: '10px 12px'}}>
        <Button style={{ fontSize: 14, padding: '0 6px', marginRight:30, marginLeft:30, alignSelf:'center' }} type='text' size='mini' onClick={addItem}>
          <IconPlus />{props.text}
        </Button>
      </div>
    );
  }

  return (
    <div style={{marginLeft:10}}>
      {appId !== '' && (
        <Space>
          <Typography.Text>应用:</Typography.Text>
          <Select dropdownMenuStyle={{ maxHeight: 400 }} value={appName} onChange={onAppChange} bordered={false}
                  triggerProps={{autoAlignPopupWidth: false, autoAlignPopupMinWidth: true, position: 'bl'}}
                  style={{width:'fit-content', minWidth:120, maxWidth:250, backgroundColor:'var(--color-fill-2)'}}
                  dropdownRender={(menu) => (<div>{menu}<Divider style={{ margin: 0 }} /><CreateItem text='创建应用'/></div>)}
          >
            {appList.map((option) => (
              <Select.Option key={option.id} value={option.id} style={{height:50, textAlign:'left', display:'block'}}>
                {option.name}
              </Select.Option>
            ))}
          </Select>
        </Space>
      )}

      {multiTenant && (
        <Space style={{marginLeft:20}}>
          <Typography.Text>租户:</Typography.Text>
          <Select dropdownMenuStyle={{ maxHeight: 400 }} value={tenantName} onChange={setTenantName} bordered={false}
                  triggerProps={{autoAlignPopupWidth: false, autoAlignPopupMinWidth: true, position: 'bl'}}
                  style={{width:'fit-content', minWidth:120, maxWidth:250, backgroundColor:'var(--color-fill-2)', visibility: visibility}}
                  dropdownRender={(menu) => (<div>{menu}<Divider style={{ margin: 0 }} /><CreateItem text='创建租户'/></div>)}
          >
            {tenantList.map((option) => (
              <Select.Option key={option.id} value={option.name} style={{height:50, textAlign:'left', display:'block'}}>
                {option.name}
              </Select.Option>
            ))}
          </Select>
        </Space>
      )}
    </div>

  );
}

export default ApplicationSelector;
