export function isArray(val): boolean {
  return Object.prototype.toString.call(val) === '[object Array]';
}
export function isObject(val): boolean {
  return Object.prototype.toString.call(val) === '[object Object]';
}
export function isString(val): boolean {
  return Object.prototype.toString.call(val) === '[object String]';
}

export const isSSR = (function () {
  try {
    return !(typeof window !== 'undefined' && document !== undefined);
  } catch (e) {
    return true;
  }
})();

export function isValidUri(uri: string, path: string): boolean {
  const regexPath = path.replace(/{.*?}/g, '[\\w\\-_]*$');
  const regex = new RegExp(`^${regexPath}$`);
  return regex.test(uri);
}


export function isIPAddress(str: string): boolean {
  const ipAddressRegex = /^((\d{1,3}\.){3}\d{1,3})$/;
  return ipAddressRegex.test(str);
}