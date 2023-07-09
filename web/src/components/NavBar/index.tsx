import React, { useContext } from 'react';
import {Tooltip, Avatar, Select, Dropdown, Menu, Divider, Message} from '@arco-design/web-react';
import {IconLanguage, IconNotification, IconSunFill, IconMoonFill, IconPoweroff} from '@arco-design/web-react/icon';
import ApplicationSelector from "@/components/NavBar/AppSwitch";
import MessageBox from '@/components/MessageBox';
import IconButton from './IconButton';
import Logo from '@/assets/logo.svg';

import styles from './style/index.module.less';
import { GlobalContext } from '@/context';
import useLocale from '@/utils/useLocale';
import defaultLocale from '@/locale';
import useStorage from '@/utils/useStorage';
import store from "@/store/mobx";
import {observer} from "mobx-react";

function Navbar() {
  const t = useLocale();
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [_, setUserStatus] = useStorage('userStatus');
  const { setLang, lang, theme, setTheme } = useContext(GlobalContext);
  function onMenuItemClick(key) {
    if (key === 'logout') {
      setUserStatus('logout');
      window.location.href = '/login';
    }
  }

  const dropList = ( // 头像下拉框
    <Menu onClickMenuItem={onMenuItemClick}>
      <Divider style={{ margin: '4px 0' }} />
      <Menu.Item key="logout">
        <IconPoweroff className={styles['dropdown-icon']} />
        {t['navbar.logout']}
      </Menu.Item>
    </Menu>
  );

  return (
    <div className={styles.navbar}>
      <div style={{display:'flex',alignItems:'center', width:'fit-content'}}>
        <div style={{display:'flex', alignItems:'center', width:'200px', paddingLeft:'20px', boxSizing:'border-box'}}>
          <Logo />
          <div className={styles['logo-name']}>Quick Auth</div>
          <div style={{height:25, width:15, borderRight:2, borderRightStyle:'solid', borderRightColor:'#BBBBBB'}}/>
        </div>
        <ApplicationSelector/>
      </div>

      <ul className={styles.right}>
        <li>
          <Select
            value={lang} trigger="hover" triggerElement={<IconButton icon={<IconLanguage />} />}
            triggerProps={{autoAlignPopupWidth: false, autoAlignPopupMinWidth: true, position: 'br'}}
            options={[{ label: '中文', value: 'zh-CN' }, { label: 'English', value: 'en-US' }]}
            onChange={(value) => {
              setLang(value);
              const nextLang = defaultLocale[value];
              Message.info(`${nextLang['message.lang.tips']}${value}`);
            }}
          />
        </li>
        <li>
          <Tooltip content={theme === 'light' ? t['settings.navbar.theme.toDark'] : t['settings.navbar.theme.toLight']}>
            <IconButton icon={theme !== 'dark' ? <IconMoonFill /> : <IconSunFill />}
              onClick={() => setTheme(theme === 'light' ? 'dark' : 'light')}
            />
          </Tooltip>
        </li>
        <li>
          <MessageBox>
            <IconButton icon={<IconNotification />} />
          </MessageBox>
        </li>
        {store.userInfo && (
          <li>
            <Dropdown droplist={dropList} position="br">
              <Avatar size={32} style={{ cursor: 'pointer' }}>
                <img alt="avatar" src={store.userInfo.avatar} />
              </Avatar>
            </Dropdown>
          </li>
        )}
      </ul>
    </div>
  );
}

export default observer(Navbar);
