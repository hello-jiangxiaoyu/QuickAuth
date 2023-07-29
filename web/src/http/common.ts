import {format} from "util";
import env from "../store/env.json"

export interface Root<T> {
  code: number;
  msg: string;
  data: T;
}

async function SendHttpRequest<T>(method:string, uri:string, wrapMsg?:string, data?:object):Promise<Root<T>> {
  let errorReason = '';
  const url = env.devHost + uri;
  let response:Root<T>;
  if (method === 'GET' || method === 'DELETE') {
    response = await fetch(url, {method:method}).then((resp) => resp.json()).catch((reason) => {
      errorReason = reason;
    });
  } else {
    response = await fetch(url, {method:method, body:JSON.stringify(data)}).then((resp) => resp.json()).catch((reason) => {
      errorReason = reason;
    });
  }

  if (typeof response !== 'object') {
    return { code:500, msg:"Server error: " + errorReason, data: null};
  }

  if (response?.code !== 200) {
    response.msg = format("%s (%d)", response?.msg, response?.code);
    if (typeof wrapMsg === 'string' && wrapMsg !== '') {
      response.msg = wrapMsg + ' error: ' + response.msg
    }
  }

  return response;
}

export async function GET<T>(uri: string, wrapMsg?:string):Promise<Root<T>> {
  return await SendHttpRequest<T>('GET', uri, wrapMsg);
}

export async function POST<T>(uri: string, data:object, wrapMsg?:string):Promise<Root<T>> {
  return await SendHttpRequest<T>('POST', uri, wrapMsg, data);
}

export async function PUT(uri: string, data:object, wrapMsg?:string):Promise<Root<object>> {
  return await SendHttpRequest<object>('PUT', uri, wrapMsg, data);
}

export async function DELETE(uri: string, wrapMsg?:string):Promise<Root<object>> {
  return await SendHttpRequest<object>('DELETE', uri, wrapMsg);
}
