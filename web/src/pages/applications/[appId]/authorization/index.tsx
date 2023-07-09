import React from 'react';
import {Card, Grid} from "@arco-design/web-react";

function Page() {
  return (
    <>
      <div style={{ minHeight:'80vh', marginLeft:10 }}>
        <Grid.Row gutter={24} style={{minHeight:'200', width:'100%'}}>
          <Card style={{minHeight:'80vh', width:'100%'}}></Card>
        </Grid.Row>
      </div>
    </>
  );
}

Page.displayName = 'Application'
export default Page;
