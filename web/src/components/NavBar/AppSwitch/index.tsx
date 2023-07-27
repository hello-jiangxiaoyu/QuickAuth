import React, {useEffect, useState} from 'react';
import {Button, Divider, Message, Select, Space, Typography} from "@arco-design/web-react";
import {IconPlus} from "@arco-design/web-react/icon";
import Router, {useRouter} from "next/router";
import {getRouterPara, replaceUriAppId} from "@/utils/stringTools";
import {observer} from "mobx-react";
import {apps} from "@/store/mobx";
import {fetchTenant, fetchTenantList} from "@/http/tenant";
import {fetchApp} from "@/http/app";

function ApplicationSelector() {
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const visibility = appId === '' ? 'hidden' : 'visible';
  const [appName, setApp] = useState('')
  const [tenantName, setTenant] = useState('')

  function updateTenant() {
    fetchTenantList(appId).then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        apps.setTenantList(r.data);
        if (r.data.length > 0) {
          apps.setCurrentTenant(r.data[0]);
        } else {
          setTenant('');
        }
      }
    })
  }
  useEffect(() => {
    if (typeof appId === 'string' && appId !== '') {
      fetchApp(appId).then(r => {
        if (r.code !== 200) {Message.error(r.msg)} else {
          setApp(r.data.name);
          apps.setCurrentApp(r.data);
        }
      })
      updateTenant()
    }
  }, [appId])

  function onAppChange(value: string) {
    const newUri = replaceUriAppId(value, router.asPath);
    if (newUri === router.asPath) {
      return
    }
    Router.push(newUri).then(() => {
      updateTenant()
    });
  }

  function onTenantChange(value: number) {
    fetchTenant(appId, value).then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        apps.setCurrentTenant(r.data);
        setTenant(r.data[0]);
      }
    })
  }

  const addItem = () => {
    console.log('add item')
  };

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
            {apps.appList.map((option) => (
              <Select.Option key={option.id} value={option.id} style={{height:50, textAlign:'left', display:'block'}}>
                {option.name}
              </Select.Option>
            ))}
          </Select>
        </Space>
      )}

      {apps.multiTenant && (
        <Space style={{marginLeft:20}}>
          <Typography.Text>租户:</Typography.Text>
          <Select dropdownMenuStyle={{ maxHeight: 400 }} value={tenantName} onChange={onTenantChange} bordered={false}
                  triggerProps={{autoAlignPopupWidth: false, autoAlignPopupMinWidth: true, position: 'bl'}}
                  style={{width:'fit-content', minWidth:120, maxWidth:250, backgroundColor:'var(--color-fill-2)', visibility: visibility}}
                  dropdownRender={(menu) => (<div>{menu}<Divider style={{ margin: 0 }} /><CreateApplication text='创建租户'/></div>)}
          >
            {apps.tenantList.map((option) => (
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

export default observer(ApplicationSelector);
