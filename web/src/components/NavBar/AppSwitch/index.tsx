import React, {useState} from 'react';
import {Button, Divider, Select} from "@arco-design/web-react";
import {IconPlus} from "@arco-design/web-react/icon";
import {useRouter} from "next/router";
import {getRouterPara} from "@/utils/getUrlParams";

export default function ApplicationSelector() {
  const [options, setOptions] = useState(['Beijing', 'Shanghai', 'Guangzhou', 'Shenzhen', '1', '2', '3', '4', '5']);
  const [inputValue, setInputValue] = useState('');
  const router = useRouter();
  const clientId = getRouterPara(router.query.clientId);

  const addItem = () => {
    if (inputValue && options.indexOf(inputValue) === -1) {
      setOptions(options.concat([inputValue]));
      setInputValue('');
    }
  };

  return (
    <Select style={{width:'fit-content', minWidth:120, maxWidth:250, backgroundColor:'var(--color-fill-1)'}} bordered={false} dropdownRender={
      (menu) => (
        <div>
          {menu}<Divider style={{ margin: 0 }} />
          <div style={{display: 'flex', alignItems: 'center', padding: '10px 12px'}}>
            <Button style={{ fontSize: 14, padding: '0 6px', marginRight:30, marginLeft:30, alignSelf:'center' }} type='text' size='mini' onClick={addItem}>
              <IconPlus />创建新应用{clientId}
            </Button>
          </div>
        </div>
      )} dropdownMenuStyle={{ maxHeight: 400 }} triggerProps={{
        autoAlignPopupWidth: false,
        autoAlignPopupMinWidth: true,
        position: 'bl',
      }} defaultValue={options[0]}
    >
      {options.map((option) => (
        <Select.Option key={option} value={option} style={{height:50, textAlign:'left', display:'block'}}>
          {option}
        </Select.Option>
      ))}
    </Select>
  )
}

