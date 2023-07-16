import * as util from "util";

export interface Root<T> {
  code: number;
  msg: string;
  data: T;
}

export async function fetchData<T>(url: string):Promise<Root<T>> {
  const result  = await fetch(url).then((resp) => resp.json())

  if (result.code !== 200) {
    result.msg = util.format("%s (%d)", result?.msg, result?.code)
  }

  return result
}
