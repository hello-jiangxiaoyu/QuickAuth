import React, { ReactNode } from 'react';
import { Layout, Spin } from '@arco-design/web-react';
import Navbar from '@/components/NavBar';
import Footer from '@/components/Footer';
import NoAccess from '@/pages/exception/403';

import { useRouter } from 'next/router';
import styles from '@/style/layout.module.less';
import ApplicationSiderWithRouter from "@/router/sider";
import store from "@/store/mobx"
import {observer} from "mobx-react";

function PageLayout({ children }: { children: ReactNode }) {
  const router = useRouter();
  const pathname = router.pathname;

  return (
    <Layout className={styles.layout}>
      <div className={styles['layout-navbar']}>
        <Navbar/>
      </div>
      {store.userLoading ? (<Spin className={styles['spin']} />) : (
        <Layout>
          <ApplicationSiderWithRouter/>
          <Layout className={styles['layout-content']} style={{ paddingLeft:store.settings.menuWidth, paddingTop:60 }}>
            <div className={styles['layout-content-wrapper']}>
              <Layout.Content>
                {pathname !== '/_error' ? children : <NoAccess />/*routeMap.current.has(pathname) ? children : <NoAccess />*/}
              </Layout.Content>
            </div>
            <Footer />
          </Layout>
        </Layout>
      )}
    </Layout>
  );
}

export default observer(PageLayout);
