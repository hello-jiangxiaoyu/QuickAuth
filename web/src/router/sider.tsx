import React, {ReactNode, useState} from 'react';
import {Layout, Menu} from '@arco-design/web-react';
import Link from "next/link";
import {IconApps, IconHistory, IconHome, IconIdcard, IconLock, IconMenuFold, IconMenuUnfold, IconMessage, IconSafe, IconUserGroup} from "@arco-design/web-react/icon";
import useLocale from "@/utils/useLocale";
import styles from "@/style/layout.module.less";
import store from "@/store/mobx";
import {observer} from "mobx-react";

export type IRoute = {
   name: string;
  key: string;
  breadcrumb?: boolean;
  icon: ReactNode;
};

const iconStyle = {fontSize:'18px', verticalAlign:'text-bottom'}
const siderRoutes: IRoute[] = [
  {name: 'menu.dashboard', key: 'dashboard', icon: <IconHome style={iconStyle}/>},
  {name: 'menu.applications', key: 'applications', icon: <IconApps style={iconStyle}/>},
  {name: 'menu.tenants', key: 'tenants', icon: <IconIdcard style={iconStyle}/>},
  {name: 'menu.authentication', key: 'authentication', icon: <IconSafe style={iconStyle}/>},
  {name: 'menu.authorization', key: 'authorization', icon: <IconLock style={iconStyle}/>},
  {name: 'menu.messages', key: 'messages', icon: <IconMessage style={iconStyle}/>},
  {name: 'menu.pools', key: 'pools', icon: <IconUserGroup style={iconStyle}/>},
  {name: 'menu.audit', key: 'audit', icon: <IconHistory style={iconStyle}/>},
];

function ApplicationSiderWithRouter() {
  const locale = useLocale();
  const paddingTop = { paddingTop: 60 };
  const [selectedKeys, setSelectedKeys] = useState<string[]>(["dashboard"]);

  return (
    <Layout.Sider collapsed={store.settings.siderCollapsed} onCollapse={store.setCollapsed} collapsible
      className={styles['layout-sider']} style={paddingTop} width={store.settings.menuWidth} breakpoint="xl" trigger={null}
    >
      <div className={styles['menu-wrapper']}>
        <Menu collapse={store.settings.siderCollapsed} selectedKeys={selectedKeys}
          onClickMenuItem={(key)=>setSelectedKeys([key])}
        >
          {siderRoutes.map(route => (
            <Menu.Item key={route.key}>
              <Link href={`/${route.key}`}>
                <a>{route.icon} {locale[route.name] || route.name}</a>
              </Link>
            </Menu.Item>
          ))}
        </Menu>
      </div>
      <div className={styles['collapse-btn']} onClick={store.switchCollapsed}>
        {store.settings.siderCollapsed ? <IconMenuUnfold /> : <IconMenuFold />}
      </div>
    </Layout.Sider>
  );
}

export default observer(ApplicationSiderWithRouter)
