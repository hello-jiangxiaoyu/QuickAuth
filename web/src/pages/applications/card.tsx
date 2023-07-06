import React from 'react';
import Router from "next/router"
import {Button, Popover, Card, Space} from '@arco-design/web-react';
import * as Icon from '@arco-design/web-react/icon';

function MyIcon(props: {name: string, color?: string, size?: number}) {
  const iconStyle = {color: undefined, height: undefined, width: undefined}
  if (typeof props.color === 'string' && props.color !== '') {
    iconStyle.color = props.color
  }
  if (typeof props.size === 'number' && props.size !== 0) {
    iconStyle.height = props.size
    iconStyle.width = props.size
  }
  const allIcon = (name: string) => React.createElement(Icon && (Icon as any)[name], {
      style: iconStyle,
    }
  )
  return <>{allIcon(props.name)}</>
}

// application card with dynamic icon
export default function MyCard(props: { clientId: string, name: string, type: string, icon?: string}) {
  let icon = 'IconCodeSandbox';
  if (props.icon != undefined && props.icon != '') {
    icon = props.icon;
  }

  function onClickCard() {
    Router.push(`applications/${props.clientId}`).then();
  }

  function MoreButton(props: { clientId: string}) {
    return (
      <Popover
        content={
          <Button type='text' status='danger' style={{height:25}} key={props.clientId}>
            删除应用
          </Button>
        }
      >
        <Button type={'text'} style={{height:25}} key={props.clientId}>
          <MyIcon name={'IconMore'}/>
        </Button>
      </Popover>
    )
  }

  return (
    <>
      <Card
        hoverable style={{ width: 330, height: 180, boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 12px 0 rgba(0, 0, 0, 0.19)' }}
        actions={[<MoreButton key={1} clientId={props.clientId}></MoreButton>]}
      >
        <div onClick={onClickCard}>
          <Space >
            <MyIcon name={icon} color={'#2f6af1'} size={50}/>
            <div style={{margin:20}}>
              <h2>{props.name}</h2>
              <div>{props.type}</div>
            </div>
          </Space>
        </div>
      </Card>
    </>
  );
}
