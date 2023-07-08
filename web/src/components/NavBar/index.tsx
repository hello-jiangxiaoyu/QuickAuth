import React, { useContext, useEffect } from 'react';
import {Tooltip, Avatar, Select, Dropdown, Menu, Divider, Message, Button} from '@arco-design/web-react';
import {IconLanguage, IconNotification, IconSunFill, IconMoonFill, IconSettings, IconPoweroff} from '@arco-design/web-react/icon';
import { useSelector, useDispatch } from 'react-redux';
import { GlobalState } from '@/store';
import { GlobalContext } from '@/context';
import useLocale from '@/utils/useLocale';
import Logo from '@/assets/logo.svg';
import MessageBox from '@/components/MessageBox';
import IconButton from './IconButton';
import Settings from '../Settings';
import styles from './style/index.module.less';
import defaultLocale from '@/locale';
import useStorage from '@/utils/useStorage';
import { generatePermission } from '@/routes';
import ApplicationSelector from "@/components/NavBar/AppSwitch";

function Navbar({ show }: { show: boolean }) {
  const t = useLocale();
  const userInfo = useSelector((state: GlobalState) => state.userInfo);
  const dispatch = useDispatch();

  const [_, setUserStatus] = useStorage('userStatus');
  const [role, setRole] = useStorage('userRole', 'admin');

  const { setLang, lang, theme, setTheme } = useContext(GlobalContext);

  function logout() {
    setUserStatus('logout');
    window.location.href = '/login';
  }

  function onMenuItemClick(key) {
    if (key === 'logout') {
      logout();
    } else {
      Message.info(`You clicked ${key}`);
    }
  }

  useEffect(() => {
    dispatch({
      type: 'update-userInfo',
      payload: {
        userInfo: {
          ...userInfo,
          permissions: generatePermission(role),
        },
      },
    });
  }, [role]);

  if (!show) {
    return (
      <div className={styles['fixed-settings']}>
        <Settings trigger={<Button icon={<IconSettings />} type="primary" size="large" />}/>
      </div>
    );
  }

  const dropList = (
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
          <MessageBox>
            <IconButton icon={<IconNotification />} />
          </MessageBox>
        </li>
        <li>
          <Tooltip content={theme === 'light' ? t['settings.navbar.theme.toDark'] : t['settings.navbar.theme.toLight']}>
            <IconButton
              icon={theme !== 'dark' ? <IconMoonFill /> : <IconSunFill />}
              onClick={() => setTheme(theme === 'light' ? 'dark' : 'light')}
            />
          </Tooltip>
        </li>
        <Settings />
        {userInfo && (
          <li>
            <Dropdown droplist={dropList} position="br">
              <Avatar size={32} style={{ cursor: 'pointer' }}>
                <img alt="avatar" src={userInfo.avatar} />
              </Avatar>
            </Dropdown>
          </li>
        )}
      </ul>
    </div>
  );
}

export default Navbar;
