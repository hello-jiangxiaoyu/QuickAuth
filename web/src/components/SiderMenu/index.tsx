import React, {ReactNode, useState} from 'react';
import {Layout, Menu} from '@arco-design/web-react';
import Link from "next/link";
import {IconApps, IconHistory, IconHome, IconIdcard, IconLock, IconMenuFold, IconMenuUnfold, IconMessage, IconSafe, IconUserGroup} from "@arco-design/web-react/icon";
import useLocale from "@/utils/useLocale";
import styles from "@/style/layout.module.less";
import store from "@/store/mobx";
import {observer} from "mobx-react";
import {getRouterPara} from "@/utils/stringTools";
import {useRouter} from "next/router";

export type IRoute = {
  name: string;
  key: string;
  breadcrumb?: boolean;
  icon: ReactNode;
};

const iconStyle = {fontSize:'18px', verticalAlign:'text-bottom'}

function ApplicationSiderWithRouter() {
  const locale = useLocale();
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const [selectedKeys, setSelectedKeys] = useState<string[]>([router.asPath]); // todo: default key start with router path
  if (appId === "") {
    return <></>;
  }

  const siderRoutes: IRoute[] = [
    {name: 'menu.dashboard', key: `/applications/${appId}/dashboard/`, icon: <IconHome style={iconStyle}/>},
    {name: 'menu.applications', key: `/applications/${appId}/`, icon: <IconApps style={iconStyle}/>},
    {name: 'menu.tenants', key: `/applications/${appId}/tenants/`, icon: <IconIdcard style={iconStyle}/>},
    {name: 'menu.authentication', key: `/applications/${appId}/authentication/`, icon: <IconSafe style={iconStyle}/>},
    {name: 'menu.authorization', key: `/applications/${appId}/authorization/`, icon: <IconLock style={iconStyle}/>},
    {name: 'menu.messages', key: `/applications/${appId}/messages/`, icon: <IconMessage style={iconStyle}/>},
    {name: 'menu.pools', key: `/applications/${appId}/pools/`, icon: <IconUserGroup style={iconStyle}/>},
    {name: 'menu.audit', key: `/applications/${appId}/audit/`, icon: <IconHistory style={iconStyle}/>},
  ];

  return (
    <Layout.Sider collapsed={store.settings.siderCollapsed} onCollapse={store.setCollapsed} collapsible
      className={styles['layout-sider']} style={{ paddingTop: 60 }} width={store.settings.menuWidth} breakpoint="xl" trigger={null}
    >
      <div className={styles['menu-wrapper']}>
        <Menu collapse={store.settings.siderCollapsed} selectedKeys={selectedKeys}
          onClickMenuItem={(key)=>setSelectedKeys([key])}
        >
          <Menu.Item key="1">
          </Menu.Item>
          {siderRoutes.map(route => (
            <Menu.Item key={route.key}>
              <Link href={`${route.key}`}>
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

export default observer(ApplicationSiderWithRouter);
