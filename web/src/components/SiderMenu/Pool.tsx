import React, {useState} from 'react';
import {
  IconApps,
  IconHistory,
  IconHome,
  IconLock, IconMenuFold,
  IconMenuUnfold,
  IconMessage,
  IconSafe
} from "@arco-design/web-react/icon";
import {IRoute} from "@/components/SiderMenu/index";
import {Layout, Menu} from "@arco-design/web-react";
import mobxStore from "@/store/mobx";
import styles from "@/style/layout.module.less";
import env from "@/store/env.json";
import Link from "next/link";
import useLocale from "@/utils/useLocale";
import {useRouter} from "next/router";

const iconStyle = {fontSize:'18px', verticalAlign:'text-bottom'};

export default function PoolSiderWithRouter() {
  const locale = useLocale();
  const router = useRouter();
  const [selectedKeys, setSelectedKeys] = useState<string[]>([router.asPath]);

  const poolsRoutes: IRoute[] = [
    {name: 'menu.dashboard', key: `/pools/dashboard/`, icon: <IconHome style={iconStyle}/>},
    {name: 'menu.applications', key: `/pools/`, icon: <IconApps style={iconStyle}/>},
    {name: 'menu.authentication', key: `/pools/authentication/`, icon: <IconSafe style={iconStyle}/>},
    {name: 'menu.messages', key: `/pools/messages/`, icon: <IconMessage style={iconStyle}/>},
    {name: 'menu.authorization', key: `/pools/authorization/`, icon: <IconLock style={iconStyle}/>},
    {name: 'menu.audit', key: `/pools/audit/`, icon: <IconHistory style={iconStyle}/>},
  ];
  return (
    <Layout.Sider
      collapsed={mobxStore.menuCollapsed} onCollapse={mobxStore.setCollapsed} collapsible
      className={styles['layout-sider']} style={{ paddingTop: 60 }} width={env.menuWidth}
      collapsedWidth={env.menuCollapseWith} breakpoint="xl" trigger={null}
    >
      <div className={styles['menu-wrapper']}>
        <Menu collapse={mobxStore.menuCollapsed} selectedKeys={selectedKeys}
              onClickMenuItem={(key)=>setSelectedKeys([key])}
        >
          {poolsRoutes.map(route => (
            <Menu.Item key={route.key}>
              <Link href={`${route.key}`}>
                <a>{route.icon} {locale[route.name] || route.name}</a>
              </Link>
            </Menu.Item>
          ))}
        </Menu>
      </div>
      <div className={styles['collapse-btn']} onClick={mobxStore.switchCollapsed}>
        {mobxStore.menuCollapsed ? <IconMenuUnfold /> : <IconMenuFold />}
      </div>
    </Layout.Sider>
  );
}
