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
