import React, {ReactNode, useEffect, useState} from 'react';
import {Layout, Menu} from '@arco-design/web-react';
import Link from "next/link";
import {IconApps, IconHistory, IconHome, IconLock, IconMenuFold, IconMenuUnfold, IconMessage, IconSafe} from "@arco-design/web-react/icon";
import useLocale from "@/utils/useLocale";
import styles from "@/style/layout.module.less";
import env from "@/store/env.json";
import {getRouterPara, removeQueryParams} from "@/utils/stringTools";
import {useRouter} from "next/router";
import {useSelector} from "react-redux";
import {dispatchMenuCollapse, GlobalState} from "@/store/redux";

export type IRoute = {
  name: string;
  key: string;
  breadcrumb?: boolean;
  icon: ReactNode;
};

const iconStyle = {fontSize:'18px', verticalAlign:'text-bottom'}

export default function ApplicationSiderWithRouter() {
  const locale = useLocale();
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const [selectedKeys, setSelectedKeys] = useState<string[]>([removeQueryParams(router.asPath)]);
  useEffect(() => {
    setSelectedKeys([removeQueryParams(router.asPath)]);
  }, [router.asPath])

  const {collapsed} = useSelector((state: GlobalState) => state);
  if (appId === "") {
    return <></>;
  }

  const sideRoutes: IRoute[] = [
    {name: 'menu.dashboard', key: `/applications/${appId}/dashboard/`, icon: <IconHome style={iconStyle}/>},
    {name: 'menu.applications', key: `/applications/${appId}/`, icon: <IconApps style={iconStyle}/>},
    {name: 'menu.authentication', key: `/applications/${appId}/authentication/`, icon: <IconSafe style={iconStyle}/>},
    {name: 'menu.messages', key: `/applications/${appId}/messages/`, icon: <IconMessage style={iconStyle}/>},
    {name: 'menu.authorization', key: `/applications/${appId}/authorization/`, icon: <IconLock style={iconStyle}/>},
    {name: 'menu.audit', key: `/applications/${appId}/audit/`, icon: <IconHistory style={iconStyle}/>},
  ];

  return (
    <Layout.Sider
      collapsed={collapsed} onCollapse={dispatchMenuCollapse} collapsible
      className={styles['layout-sider']} style={{ paddingTop: 60 }} width={env.menuWidth}
      collapsedWidth={env.menuCollapseWith} breakpoint="xl" trigger={null}
    >
      <div className={styles['menu-wrapper']}>
        <Menu collapse={collapsed} selectedKeys={selectedKeys}
              onClickMenuItem={(key)=>setSelectedKeys([key])}
        >
          {sideRoutes.map(route => (
            <Menu.Item key={route.key}>
              <Link href={`${route.key}`}>
                <a>{route.icon} {locale[route.name] || route.name}</a>
              </Link>
            </Menu.Item>
          ))}
        </Menu>
      </div>
      <div className={styles['collapse-btn']} onClick={()=>dispatchMenuCollapse(!collapsed)}>
        {collapsed ? <IconMenuUnfold /> : <IconMenuFold />}
      </div>
    </Layout.Sider>
  );
}
