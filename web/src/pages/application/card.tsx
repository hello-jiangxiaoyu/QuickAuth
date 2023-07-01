import React from 'react';
import Router from "next/router"
import {Card, Space} from '@arco-design/web-react';
import * as Icon from '@arco-design/web-react/icon';

function MyIcon(props: {name: string, color?: string}) {
  const iconStyle = {height: 50, width: 50, color: '#2f6af1'}
  if (props.color != undefined && props.color != '') {
    iconStyle.color = props.color
  }
  const allIcon = (name: string) => React.createElement(Icon && (Icon as any)[name], {
      style: iconStyle,
    }
  )
  return <>{allIcon(props.name)}</>
}

// application card with dynamic icon
export default function MyCard(props: { clientId: string, name: string, type: string, icon?: string}) {
  let icon = 'IconCodeSandbox'
  if (props.icon != undefined && props.icon != '') {
    icon = props.icon
  }

  function onClickCard() {
    Router.push(`application/${props.clientId}`).then();
  }
  return (
    <>
      <Card hoverable onClick={onClickCard}
            style={{ width: 330, height: 180, boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 12px 0 rgba(0, 0, 0, 0.19)' }}
      >
        <Space>
          <MyIcon name={icon}></MyIcon>
          <div style={{margin:20}}>
            <h2>{props.name}</h2>
            <div>{props.type}</div>
          </div>
        </Space>
      </Card>
    </>
  )
}
