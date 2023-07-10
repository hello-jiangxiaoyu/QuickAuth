import React from "react";
import * as Icon from "@arco-design/web-react/icon";

export default function MyIcon(props: {name: string, color?: string, size?: number}) {
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
  );
  return <>{allIcon(props.name)}</>
}
