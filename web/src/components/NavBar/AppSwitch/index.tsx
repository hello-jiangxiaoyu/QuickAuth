import React, {useState} from 'react';
import {Button, Divider, Message, Select, Space, Typography} from "@arco-design/web-react";
import {IconPlus} from "@arco-design/web-react/icon";
import Router, {useRouter} from "next/router";
import {getRouterPara, replaceUriAppId} from "@/utils/stringTools";
import {useSelector} from "react-redux";
import {dispatchTenant, GlobalState} from "@/store/redux";
import CreateAppDialog from "@/components/Dialog/App";
import CreateTenantDialog from "@/components/Dialog/Tenant";
import api from "@/http/api";

export default function ApplicationSelector() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const {appList, currentApp, tenantList, currentTenant} = useSelector((state: GlobalState) => state);

  function onAppChange(value: string) {
    const newUri = replaceUriAppId(value, router.asPath);
    if (newUri === router.asPath) {
      return
    }
    Router.push(newUri).then();
  }

  function onTenantChange(value: number) {
    api.fetchTenant(appId, value).then(r => {
      dispatchTenant(r.data);
    }).catch(e => Message.error(e.toString()));
  }

  function CreateItem(props:{text:string, isApp:boolean}) {
    const [visible, setVisible] = useState(false);
    return (
      <div style={{display: 'flex', alignItems: 'center', padding: '10px 12px'}}>
        <Button style={{ fontSize: 14, padding: '0 6px', marginRight:30, marginLeft:30, alignSelf:'center' }} type='text' size='mini' onClick={()=>setVisible(true)}>
          <IconPlus />{props.text}
        </Button>
        {props.isApp ?
          <CreateAppDialog visible={visible} setVisible={setVisible}/> :
          <CreateTenantDialog visible={visible} setVisible={setVisible}/>
        }
      </div>
    );
  }

  return (
    <div style={{marginLeft:10}}>
      <Space>
        <Typography.Text>应用:</Typography.Text>
        <Select
          dropdownMenuStyle={{ maxHeight: 400 }} value={currentApp?.id} onChange={onAppChange} bordered={false}
          triggerProps={{autoAlignPopupWidth: false, autoAlignPopupMinWidth: true, position: 'bl'}}
          style={{width:'fit-content', minWidth:120, maxWidth:250, backgroundColor:'var(--color-fill-2)'}}
          dropdownRender={(menu) => (<div>{menu}<Divider style={{ margin: 0 }} /><CreateItem text='创建应用' isApp={true}/></div>)}
        >
          {appList && appList.map((option) => (
            <Select.Option key={option.id} value={option.id} style={{height:50, textAlign:'left', display:'block'}}>
              {option.name}
            </Select.Option>
          ))}
        </Select>
      </Space>

      {currentApp?.tag==='Multi Tenant' && (
        <Space style={{marginLeft:20}}>
          <Typography.Text>租户:</Typography.Text>
          <Select
            dropdownMenuStyle={{ maxHeight: 400 }} value={currentTenant?.id} onChange={onTenantChange} bordered={false}
            triggerProps={{autoAlignPopupWidth: false, autoAlignPopupMinWidth: true, position: 'bl'}}
            style={{width:'fit-content', minWidth:120, maxWidth:250, backgroundColor:'var(--color-fill-2)'}}
            dropdownRender={(menu) => (<div>{menu}<Divider style={{ margin: 0 }} /><CreateItem text='创建租户' isApp={false}/></div>)}
          >
            {tenantList.map((option) => (
              <Select.Option key={option.id} value={option.id} style={{height:50, textAlign:'left', display:'block'}}>
                {option.name}
              </Select.Option>
            ))}
          </Select>
        </Space>
      )}
    </div>
  );
}
