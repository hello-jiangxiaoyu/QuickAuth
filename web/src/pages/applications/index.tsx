import React from 'react';
import MyCard, {AddApp} from './card'
import {Grid} from "@arco-design/web-react";
import {observer} from "mobx-react";
import {useSelector} from "react-redux";
import {GlobalState} from "@/store/redux";

function Page() {
  //   {name: 'default', tag: 'auth', icon: 'IconSafe', appId: '1'},
  //   {name: 'jiang', tag: 'backend', icon: 'IconCloudDownload', appId: '2'},
  //   {name: 'zhao', tag: 'frontend', icon: 'IconMessage', appId: '3'},
  //   {name: 'feng', tag: 'insert', icon: 'IconBook', appId: '4'},
  //   {name: 'qiang', tag: 'python', icon: 'IconRobot', appId: '5'},

  const appList = useSelector((state: GlobalState) => state.appList);
  return (
    <div style={{ minHeight:'80vh' }}>
      <Grid.Row gutter={24} style={{minHeight:'200', width:'100%'}}>
        {appList && appList.map((item, index) => (
          <Grid.Col xs={24} sm={12} md={8} lg={6} xl={6} xxl={6} key={index} style={{ marginBottom: 16 }} flex='100px'>
            <MyCard appId={item.id} name={item.name} type={item.describe} icon={item.icon}></MyCard>
          </Grid.Col>
        ))}
        <AddApp></AddApp>
      </Grid.Row>
    </div>
  );
}

export default observer(Page);
