import React from 'react';
import MyCard from './card'
import {Grid} from "@arco-design/web-react";

function Example() {
  const list = [
    {name: 'default', type: 'auth', icon: 'IconSafe', clientId: '1'},
    {name: 'jiang', type: 'backend', icon: 'IconCloudDownload', clientId: '2'},
    {name: 'zhao', type: 'frontend', icon: 'IconMessage', clientId: '3'},
    {name: 'feng', type: 'insert', icon: 'IconBook', clientId: '4'},
    {name: 'qiang', type: 'python', icon: 'IconRobot', clientId: '5'},
  ];

  return (
    <div style={{ minHeight:'80vh', marginLeft:10 }}>
      <Grid.Row gutter={24} style={{minHeight:'200', width:'100%'}}>
        {list.map((item, index) => (
          <Grid.Col xs={24} sm={12} md={8} lg={6} xl={6} xxl={6} key={index} style={{ marginBottom: 16 }} flex='100px'>
            <MyCard clientId={item.clientId} name={item.name} type={item.type} icon={item.icon}></MyCard>
          </Grid.Col>
        ))}
      </Grid.Row>
    </div>
  );
}

export default Example;
