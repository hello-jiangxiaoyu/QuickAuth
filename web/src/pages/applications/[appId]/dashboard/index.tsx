import React, {useRef} from 'react';
import {Button, Card, Grid, Input} from "@arco-design/web-react";

function Page() {
  const inputRef = useRef(null);
  return (
    <>
      <div style={{ minHeight:'80vh', marginLeft:10 }}>
        <Grid.Row gutter={24} style={{minHeight:'200', width:'100%'}}>
          <Card style={{minHeight:'80vh', width:'100%'}}>
            <Input ref={inputRef} style={{ width: 350 }} allowClear placeholder='Please Enter something' />
            <Button>Clear</Button>
          </Card>
        </Grid.Row>
      </div>
    </>
  );
}

Page.displayName = 'Application'
export default Page;

