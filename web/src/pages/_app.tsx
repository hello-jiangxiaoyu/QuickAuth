import React, { useEffect, useMemo } from 'react';
import zhCN from '@arco-design/web-react/es/locale/zh-CN';
import enUS from '@arco-design/web-react/es/locale/en-US';
import '@/mock';
import '@/style/global.less';
import {ConfigProvider, Message} from '@arco-design/web-react';
import type { AppProps } from 'next/app';
import Head from 'next/head';
import cookies from 'next-cookies';
import Router, { useRouter } from 'next/router';
import useStorage from '@/utils/useStorage';
import { GlobalContext } from '@/context';
import {Provider} from 'react-redux';
import Layout from './layout';
import NProgress from 'nprogress';
import changeTheme from '@/utils/changeTheme';
import {dispatchApp, dispatchAppList, dispatchTenant, dispatchTenantList, store} from '@/store/redux';
import {getRouterPara} from "@/utils/stringTools";
import {Tenant, TenantDetail} from "@/http/tenant";
import api from "@/http/api";

async function updateAppAndTenant(appId:string):Promise<string> {
  const respApp = await api.fetchApp(appId);
  dispatchApp(respApp.data);

  if (respApp.data.tenant === null || respApp.data.tenant.length === 0) {
    dispatchTenantList([] as Array<Tenant>);
    dispatchTenant({} as TenantDetail);
    return
  }

  const respTenant = await api.fetchTenant(appId, respApp.data.tenant[0].id);
  dispatchTenantList(respApp.data.tenant);
  dispatchTenant(respTenant.data);
}

function MyApp({pageProps, Component, renderConfig}: AppProps & { renderConfig: {arcoLang?: string; arcoTheme?: string} }) {
  const { arcoLang, arcoTheme } = renderConfig;
  const [lang, setLang] = useStorage('arco-lang', arcoLang || 'en-US');
  const [theme, setTheme] = useStorage('arco-theme', arcoTheme || 'light');
  const contextValue = {lang, setLang, theme, setTheme};
  const locale = useMemo(() => {
    if (lang === 'en-US') {return enUS}
    return zhCN;
  }, [lang]);

  const router = useRouter();
  const appId = getRouterPara(router.query.appId);
  useEffect(() => { // 刷新页面时，检查登录状态
    api.fetchUserInfo();
    api.fetchMe().then().catch(() => {
      setTimeout(()=>{Router.push('/login').then()}, 2000); // 跳转到登录页
    })
    api.fetchAppList().then(r => dispatchAppList(r.data)).catch();
  }, []);

  useEffect(() => {changeTheme(theme)}, [lang, theme]);
  useEffect(() => { // 首次加载，以及appId发生变化
    if (typeof appId === 'string' && appId !== '') {
      updateAppAndTenant(appId).then().catch(() => {
        setTimeout(()=>{Router.push('/applications/').then()}, 2000);
      });
    }
  }, [appId]);

  useEffect(() => { // 页面加载进度条
    const handleStart = () => {
      NProgress.set(0.4);
      NProgress.start();
    };
    const handleStop = () => {NProgress.done()};
    router.events.on('routeChangeStart', handleStart);
    router.events.on('routeChangeComplete', handleStop);
    router.events.on('routeChangeError', handleStop);
    return () => {
      router.events.off('routeChangeStart', handleStart);
      router.events.off('routeChangeComplete', handleStop);
      router.events.off('routeChangeError', handleStop);
    };
  }, [router]);

  return (
    <>
      <Head>
        <link rel="shortcut icon" type="image/x-icon" href="https://unpkg.byted-static.com/latest/byted/arco-config/assets/favicon.ico"/>
      </Head>
      <ConfigProvider locale={locale} componentConfig={{Card: {bordered: false}, List: {bordered: false}, Table: {border: false}}}>
        <Provider store={store}>
          <GlobalContext.Provider value={contextValue}>
            {Component.displayName === 'LoginPage' ? (
              <Component {...pageProps} suppressHydrationWarning />
            ) : (
              <Layout><Component {...pageProps} suppressHydrationWarning /></Layout>
            )}
          </GlobalContext.Provider>
        </Provider>
      </ConfigProvider>
    </>
  );
}

// fix: next build ssr can't attach the localstorage
MyApp.getInitialProps = async (appContext) => {
  const { ctx } = appContext;
  const serverCookies = cookies(ctx);
  return {
    renderConfig: {
      arcoLang: serverCookies['arco-lang'],
      arcoTheme: serverCookies['arco-theme'],
    },
  };
};

export default MyApp
