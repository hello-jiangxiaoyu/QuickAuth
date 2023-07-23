import React, {useEffect, useState} from 'react';
import MyCard, {AddApp} from './card'
import {Grid, Message} from "@arco-design/web-react";
import {fetchAppList, App} from "@/http/app";

function Page() {
  //   {name: 'default', tag: 'auth', icon: 'IconSafe', appId: '1'},
  //   {name: 'jiang', tag: 'backend', icon: 'IconCloudDownload', appId: '2'},
  //   {name: 'zhao', tag: 'frontend', icon: 'IconMessage', appId: '3'},
  //   {name: 'feng', tag: 'insert', icon: 'IconBook', appId: '4'},
  //   {name: 'qiang', tag: 'python', icon: 'IconRobot', appId: '5'},

  const [apps, setApps] = useState<Array<App>>([])
  useEffect(() => {
    fetchAppList().then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        console.log(r.data)
        setApps(r.data);
      }
    })
  }, []);

  return (
    <div style={{ minHeight:'80vh' }}>
      <Grid.Row gutter={24} style={{minHeight:'200', width:'100%'}}>
        {apps.map((item, index) => (
          <Grid.Col xs={24} sm={12} md={8} lg={6} xl={6} xxl={6} key={index} style={{ marginBottom: 16 }} flex='100px'>
            <MyCard appId={item.id} name={item.name} type={item.describe} icon={item.icon}></MyCard>
          </Grid.Col>
        ))}
        <AddApp></AddApp>
      </Grid.Row>
    </div>
  );
}

export default Page;
