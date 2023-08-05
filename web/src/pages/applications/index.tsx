import React from 'react';
import ApplicationCard, {AddApplication} from './card'
import {Grid} from "@arco-design/web-react";
import {useSelector} from "react-redux";
import {GlobalState} from "@/store/redux";

function Page() {
  const appList = useSelector((state: GlobalState) => state.appList);
  return (
    <div style={{ minHeight:'80vh', marginLeft:100 }}>
      <Grid.Row gutter={24} style={{minHeight:'200', width:'100%'}}>
        {appList && appList.map((item, index) => (
          <Grid.Col xs={24} sm={12} md={8} lg={6} xl={6} xxl={6} key={index} style={{ marginBottom: 16 }} flex='100px'>
            <ApplicationCard appId={item.id} name={item.name} type={item.tag} icon={item.icon}></ApplicationCard>
          </Grid.Col>
        ))}
        <AddApplication></AddApplication>
      </Grid.Row>
    </div>
  );
}

export default Page;
