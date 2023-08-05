// 仅用于线上预览，实际使用中可以将此逻辑删除

export function getRouterPara(temp: unknown): string {
  let res = '';
  if (typeof temp === 'string') {
    res = temp;
  } else {
    return "";
  }

  return res;
}

export function replaceUriAppId(appId: string, uri: string): string {
  if (!uri.startsWith('/applications/')) {
    return uri
  }
  for (let i = 14; i < uri.length; i++) {
    if (uri[i] === '/') {
      uri = uri.substring(0, 14) + appId + uri.substring(i);
      break;
    } else if (i == uri.length - 1) {
      uri = uri.substring(0, 14) + appId;
    }
  }

  return uri;
}

export function removeQueryParams(uri: string): string {
  const index = uri.indexOf('?');
  if (index !== -1) {
    return uri.substring(0, index);
  }
  return uri;
}
