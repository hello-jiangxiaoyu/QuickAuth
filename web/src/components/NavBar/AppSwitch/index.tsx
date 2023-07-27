import React, {useEffect, useState} from 'react';
import {Button, Divider, Message, Select, Space, Typography} from "@arco-design/web-react";
import {IconPlus} from "@arco-design/web-react/icon";
import Router, {useRouter} from "next/router";
import {getRouterPara, replaceUriAppId} from "@/utils/stringTools";
import {observer} from "mobx-react";
import {fetchTenantList, Tenant} from "@/http/tenant";
import {App, fetchAppList} from "@/http/app";
import {apps} from "@/store/mobx";
import useStorage from "@/utils/useStorage";

function getAppById(id: string, appList:Array<App>): App {
  return appList.find((app: App) => app.id === id);
}

function ApplicationSelector() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const visibility = appId === '' ? 'hidden' : 'visible';
  const [appName, setAppName] = useState('');
  const [multiTenant, setMultiTenant] = useState(false);
  const [tenantName, setTenantName] = useState('');
  const [appList, setAppList] = useState([] as Array<App>);
  const [tenantList, setTenantList] = useState([] as Array<Tenant>);

  useEffect(() => { // 首次加载
    fetchAppList().then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        apps.setAppList(r.data)
        setAppList(r.data)
        setAppName(appId)
        setMultiTenant(getAppById(appId, r.data)?.tag === 'multi_tenant')
      }
    })
    updateTenantList(appId)
  }, [appId])

  function updateTenantList(appId:string) {
    if (typeof appId === 'string' && appId !== '') {
      fetchTenantList(appId).then(r => {
        if (r.code !== 200) {Message.error(r.msg)} else {
          setTenantList(r.data);
          if (r.data.length !== 0) {
            setTenantName(r.data[0].name);
          } else {
            setTenantName('');
          }
        }
      })
    }
  }

  function onAppChange(value: string) {
    setAppName(value)
    const newUri = replaceUriAppId(value, router.asPath);
    if (newUri === router.asPath) {
      return
    }
    Router.push(newUri).then();
  }

  const addItem = () => {console.log('add item')};

  function CreateApplication(props:{text:string}) {
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
                  dropdownRender={(menu) => (<div>{menu}<Divider style={{ margin: 0 }} /><CreateApplication text='创建应用'/></div>)}
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
                  dropdownRender={(menu) => (<div>{menu}<Divider style={{ margin: 0 }} /><CreateApplication text='创建租户'/></div>)}
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

export default observer(ApplicationSelector);
