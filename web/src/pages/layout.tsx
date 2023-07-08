import React, { ReactNode } from 'react';
import { Layout, Spin } from '@arco-design/web-react';
import cs from 'classnames';
import { useSelector } from 'react-redux';
import { useRouter } from 'next/router';
import Navbar from '@/components/NavBar';
import Footer from '@/components/Footer';
import { GlobalState } from '@/store';
import styles from '@/style/layout.module.less';
import NoAccess from '@/pages/exception/403';
import ApplicationSiderWithRouter from "@/router/sider";

function PageLayout({ children }: { children: ReactNode }) {
  const router = useRouter();
  const pathname = router.pathname;
  const { settings, userLoading } = useSelector((state: GlobalState) => state);
  const showNavbar = settings?.navbar;
  const showMenu = settings?.menu;
  const showFooter = settings?.footer;
  const paddingLeft = showMenu ? { paddingLeft: 48 } : {};
  const paddingTop = showNavbar ? { paddingTop: 60 } : {};
  const paddingStyle = { ...paddingLeft, ...paddingTop };

  return (
    <Layout className={styles.layout}>
      <div className={cs(styles['layout-navbar'], {[styles['layout-navbar-hidden']]: !showNavbar})}>
        <Navbar show={showNavbar} />
      </div>
      {userLoading ? (<Spin className={styles['spin']} />) : (
        <Layout>
          {showMenu && (<ApplicationSiderWithRouter/>)}
          <Layout className={styles['layout-content']} style={paddingStyle}>
            <div className={styles['layout-content-wrapper']}>
              <Layout.Content>
                {pathname !== '/_error' ? children : <NoAccess />/*routeMap.current.has(pathname) ? children : <NoAccess />*/}
              </Layout.Content>
            </div>
            {showFooter && <Footer />}
          </Layout>
        </Layout>
      )}
    </Layout>
  );
}

export default PageLayout;
