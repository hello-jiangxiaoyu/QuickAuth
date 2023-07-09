import React, {useState} from 'react';
import {Button, Divider, Select} from "@arco-design/web-react";
import {IconPlus} from "@arco-design/web-react/icon";
import Router, {useRouter} from "next/router";
import {getRouterPara, replaceUriAppId} from "@/utils/stringTools";
import {observer} from "mobx-react";

function ApplicationSelector() {
  const [options, setOptions] = useState(['1', '2', '3', '4', '5', '6']);
  const [inputValue, setInputValue] = useState('');
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const visibility = appId === '' ? 'hidden' : 'visible';
  const addItem = () => {
    if (inputValue && options.indexOf(inputValue) === -1) {
      setOptions(options.concat([inputValue]));
      setInputValue('');
    }
  };

  function onChange(value: string) {
    const newUri = replaceUriAppId(value, router.asPath);
    if (newUri === router.asPath) {
      return
    }
    Router.push(newUri).then();
  }

  function CreateApplication() {
    return (
      <div style={{display: 'flex', alignItems: 'center', padding: '10px 12px'}}>
        <Button style={{ fontSize: 14, padding: '0 6px', marginRight:30, marginLeft:30, alignSelf:'center' }} type='text' size='mini' onClick={addItem}>
          <IconPlus />创建新应用{appId}
        </Button>
      </div>
    );
  }

  return (
    <Select dropdownMenuStyle={{ maxHeight: 400 }} value={appId} onChange={onChange} bordered={false}
      triggerProps={{autoAlignPopupWidth: false, autoAlignPopupMinWidth: true, position: 'bl'}}
      style={{width:'fit-content', minWidth:120, maxWidth:250, backgroundColor:'var(--color-fill-1)', visibility: visibility}}
      dropdownRender={(menu) => (<div>{menu}<Divider style={{ margin: 0 }} /><CreateApplication/></div>)}
    >
      {options.map((option) => (
        <Select.Option key={option} value={option} style={{height:50, textAlign:'left', display:'block'}}>
          {option}
        </Select.Option>
      ))}
    </Select>
  );
}

export default observer(ApplicationSelector);
