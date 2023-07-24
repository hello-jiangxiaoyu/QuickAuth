import React, {useRef} from 'react';
import {Button, Card, Input} from "@arco-design/web-react";
import {apps} from "@/store/mobx";
import {observer} from "mobx-react";

function Page() {
  const inputRef = useRef(null);
  return (
    <Card style={{minHeight:'80vh', width:'100%'}}>
      <Input ref={inputRef} style={{ width: 350 }} allowClear placeholder='Please Enter something' />
      <Button onClick={()=>{apps.setMultiTenant(!apps.multiTenant)}}>Clear</Button>
    </Card>
  );
}

export default observer(Page);

