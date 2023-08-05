import React, {useState} from 'react';
import {IconUserGroup, IconIdcard, IconMenuFold, IconMenuUnfold} from "@arco-design/web-react/icon";
import {IRoute} from "@/components/SiderMenu/index";
import {Layout, Menu} from "@arco-design/web-react";
import styles from "@/style/layout.module.less";
import env from "@/store/env.json";
import Link from "next/link";
import useLocale from "@/utils/useLocale";
import {useRouter} from "next/router";
import {useSelector} from "react-redux";
import {dispatchMenuCollapse, GlobalState} from "@/store/redux";

const iconStyle = {fontSize:'18px', verticalAlign:'text-bottom'};

export default function PoolSiderWithRouter() {
  const locale = useLocale();
  const router = useRouter();
  const [selectedKeys, setSelectedKeys] = useState<string[]>([router.asPath]);
  const {collapsed} = useSelector((state: GlobalState) => state);

  if (!router.pathname.startsWith('/pools')) {
    return <></>;
  }

  const poolsRoutes: IRoute[] = [
    {name: 'menu.pool', key: `/pools/`, icon: <IconUserGroup style={iconStyle}/>},
    {name: 'menu.user', key: `/pools/users/`, icon: <IconIdcard style={iconStyle}/>},
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
          {poolsRoutes.map(route => (
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
