import React, {useRef} from 'react';
import {Card, Input} from "@arco-design/web-react";

function Page() {
  const inputRef = useRef(null);
  return (
    <Card style={{minHeight:'80vh', width:'100%'}}>
      <Input ref={inputRef} style={{ width: 350 }} allowClear placeholder='Please Enter something' />
    </Card>
  );
}

export default Page;

