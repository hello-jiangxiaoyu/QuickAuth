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
import {checkLogin} from '@/store/localStorage';
import changeTheme from '@/utils/changeTheme';

import {dispatchApp, dispatchAppList, dispatchTenant, dispatchTenantList, store} from '@/store/redux';
import {getRouterPara} from "@/utils/stringTools";
import {TenantDetail} from "@/http/tenant";
import api from "@/http/api";

async function updateAppAndTenant(appId:string):Promise<string> {
  const respApp = await api.fetchApp(appId);
  dispatchApp(respApp.data);
  if (respApp.code !== 200) {
    setTimeout(()=>{
      Router.push('/applications/').then();
    }, 2000);
    return respApp.msg;
  }

  const respTenantList = await api.fetchTenantList(appId);
  dispatchTenantList(respTenantList.data);
  if (respTenantList.code !== 200) {
    return respTenantList.msg;
  }

  if (respTenantList.data.length === 0) {
    dispatchTenant({} as TenantDetail);
    return respTenantList.msg;
  }

  const respTenant = await api.fetchTenant(appId, respTenantList.data[0].id);
  dispatchTenant(respTenant.data);
  if (respTenant.code !== 200) {
    return respTenant.msg;
  }

  return ''
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
  useEffect(() => {changeTheme(theme)}, [lang, theme]);
  useEffect(() => { // 首次加载，以及appId发生变化
    if (checkLogin()) {
      api.fetchUserInfo();
    } else if (window.location.pathname.replace(/\//g, '') !== 'login') {
      window.location.pathname = '/login';
    }
    api.fetchAppList().then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        dispatchAppList(r.data);
      }
    });
    if (typeof appId === 'string' && appId !== '') {
      updateAppAndTenant(appId).then(msg => {
        if (typeof msg === 'string' && msg !== '') {
          Message.error(msg);
        }
      })
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
