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

import Layout from './layout';
import NProgress from 'nprogress';
import {checkLogin} from '@/store/localStorage';
import changeTheme from '@/utils/changeTheme';
import {fetchUserInfo} from "@/http/users";
import {fetchApp, fetchAppList} from "@/http/app";
import {apps} from "@/store/mobx";
import {getRouterPara} from "@/utils/stringTools";
import {fetchTenantList} from "@/http/tenant";

interface RenderConfig {
  arcoLang?: string;
  arcoTheme?: string;
}

function MyApp({pageProps, Component, renderConfig}: AppProps & { renderConfig: RenderConfig }) {
  const { arcoLang, arcoTheme } = renderConfig;
  const [lang, setLang] = useStorage('arco-lang', arcoLang || 'en-US');
  const [theme, setTheme] = useStorage('arco-theme', arcoTheme || 'light');
  const contextValue = {lang, setLang, theme, setTheme};
  const locale = useMemo(() => {
    if (lang === 'en-US') {return enUS}
    return zhCN;
  }, [lang]);

  const router = useRouter();
  useEffect(() => {changeTheme(theme)}, [lang, theme]);
  useEffect(() => {
    if (checkLogin()) {
      fetchUserInfo();
    } else if (window.location.pathname.replace(/\//g, '') !== 'login') {
      window.location.pathname = '/login';
    }

    fetchAppList().then(r => {
      if (r.code !== 200) {Message.error(r.msg)} else {
        apps.updateApps(r.data);
      }
    })
  }, []);

  useEffect(() => { // 页面渲染进度条
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
        <GlobalContext.Provider value={contextValue}>
          {Component.displayName === 'LoginPage' ? (
            <Component {...pageProps} suppressHydrationWarning />
          ) : (
            <Layout><Component {...pageProps} suppressHydrationWarning /></Layout>
          )}
        </GlobalContext.Provider>
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
