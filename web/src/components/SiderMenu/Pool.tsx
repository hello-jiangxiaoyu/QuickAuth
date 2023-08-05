import React, {useState} from 'react';
import {IconUserGroup, IconIdcard, IconMenuFold, IconMenuUnfold} from "@arco-design/web-react/icon";
import {IRoute} from "@/components/SiderMenu/index";
import {Layout, Menu} from "@arco-design/web-react";
import mobxStore from "@/store/mobx";
import styles from "@/style/layout.module.less";
import env from "@/store/env.json";
import Link from "next/link";
import useLocale from "@/utils/useLocale";
import {useRouter} from "next/router";
import {getRouterPara} from "@/utils/stringTools";
import {observer} from "mobx-react";

const iconStyle = {fontSize:'18px', verticalAlign:'text-bottom'};

function PoolSiderWithRouter() {
  const locale = useLocale();
  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  const [selectedKeys, setSelectedKeys] = useState<string[]>([router.asPath]);

  if (!router.pathname.startsWith('/pools')) {
    return <></>;
  }

  const poolsRoutes: IRoute[] = [
    {name: 'menu.user', key: `/pools/users/`, icon: <IconIdcard style={iconStyle}/>},
    {name: 'menu.pool', key: `/pools/`, icon: <IconUserGroup style={iconStyle}/>},
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

export default observer(PoolSiderWithRouter);
