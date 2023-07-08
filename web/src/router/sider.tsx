import React, {ReactNode, useState} from 'react';
import {Layout, Menu} from '@arco-design/web-react';
import Link from "next/link";
import {IconApps, IconHistory, IconHome, IconIdcard, IconLock, IconMenuFold, IconMenuUnfold, IconMessage, IconSafe, IconUserGroup} from "@arco-design/web-react/icon";
import useLocale from "@/utils/useLocale";
import styles from "@/style/layout.module.less";
import {useSelector} from "react-redux";
import {GlobalState} from "@/store";
import {useRouter} from "next/router";
import qs from "query-string";
import useRoute from "@/routes";

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


export default function ApplicationSiderWithRouter() {
  const locale = useLocale();
  const { userInfo, settings } = useSelector((state: GlobalState) => state);
  const [routes, defaultRoute] = useRoute(userInfo?.permissions);
  const paddingTop = settings?.navbar !== false ? { paddingTop: 60 } : {};

  const currentComponent = qs.parseUrl(useRouter().pathname).url.slice(1);
  const paths = (currentComponent || defaultRoute).split('/');

  const defaultOpenKeys = paths.slice(0, paths.length - 1);
  const [openKeys, setOpenKeys] = useState<string[]>(defaultOpenKeys);

  const [collapsed, setCollapsed] = useState<boolean>(false);
  const menuWidth = collapsed ? 48 : settings?.menuWidth;

  const defaultSelectedKeys = [currentComponent || defaultRoute];
  const [selectedKeys, setSelectedKeys] = useState<string[]>(defaultSelectedKeys);


  return (
    <Layout.Sider
      className={styles['layout-sider']} style={paddingTop} width={menuWidth} breakpoint="xl" trigger={null}
      collapsed={collapsed} onCollapse={setCollapsed} collapsible
    >
      <div className={styles['menu-wrapper']}>
        <Menu
          collapse={collapsed} selectedKeys={selectedKeys} openKeys={openKeys}
          onClickMenuItem={(key)=>setSelectedKeys([key])} onClickSubMenu={(_, openKeys) => {setOpenKeys(openKeys)}}
        >
          {siderRoutes.map(route => (
            <Menu.Item key={route.key}>
              <Link href={`/${route.key}`}>
                <div>{route.icon} {locale[route.name] || route.name}</div>
              </Link>
            </Menu.Item>
          ))}
        </Menu>
      </div>
      <div className={styles['collapse-btn']} onClick={()=>setCollapsed((collapsed) => !collapsed)}>
        {collapsed ? <IconMenuUnfold /> : <IconMenuFold />}
      </div>
    </Layout.Sider>
  );
}
