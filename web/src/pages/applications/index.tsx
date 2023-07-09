import React from 'react';
import MyCard from './card'
import {Grid} from "@arco-design/web-react";

function Page() {
  const list = [
    {name: 'default', type: 'auth', icon: 'IconSafe', appId: '1'},
    {name: 'jiang', type: 'backend', icon: 'IconCloudDownload', appId: '2'},
    {name: 'zhao', type: 'frontend', icon: 'IconMessage', appId: '3'},
    {name: 'feng', type: 'insert', icon: 'IconBook', appId: '4'},
    {name: 'qiang', type: 'python', icon: 'IconRobot', appId: '5'},
  ];

  return (
    <div style={{ minHeight:'80vh', marginLeft:10 }}>
      <Grid.Row gutter={24} style={{minHeight:'200', width:'100%'}}>
        {list.map((item, index) => (
          <Grid.Col xs={24} sm={12} md={8} lg={6} xl={6} xxl={6} key={index} style={{ marginBottom: 16 }} flex='100px'>
            <MyCard appId={item.appId} name={item.name} type={item.type} icon={item.icon}></MyCard>
          </Grid.Col>
        ))}
      </Grid.Row>
    </div>
  );
}

Page.displayName = 'Application'
export default Page;
