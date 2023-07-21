import {format} from "util";
import env from "../store/env.json"

export interface Root<T> {
  code: number;
  msg: string;
  data: T;
}

export async function GET<T>(uri: string):Promise<Root<T>> {
  let errorReason = ''
  const url = env.devHost + uri;
  const response:Root<T> = await fetch(url).then((resp) => resp.json()).catch((reason) => {
    errorReason = reason;
  });

  if (typeof response !== 'object') {
    return { code:500, msg:"Server error: " + errorReason, data: null};
  }

  if (response?.code !== 200) {
    response.msg = format("%s (%d)", response?.msg, response?.code);
  }

  return response;
}

export async function POST<T>(uri: string, data:object = {}):Promise<Root<T>> {
  let errorReason = ''
  const url = env.devHost + uri;
  const response  = await fetch(url, {method:'POST', body:JSON.stringify(data)}).then((resp) => resp.json()).catch((reason) => {
    errorReason = reason;
  });

  if (typeof response !== 'object') {
    return { code:500, msg:"Server error: " + errorReason, data: null};
  }


  if (response.code !== 200) {
    response.msg = format("%s (%d)", response?.msg, response?.code);
  }

  return response;
}

export async function PUT(uri: string, data:object = {}):Promise<Root<object>> {
  let errorReason = ''
  const url = env.devHost + uri;
  const response  = await fetch(url, {method:'PUT', body:JSON.stringify(data)}).then((resp) => resp.json()).catch((reason) => {
    errorReason = reason;
  });

  if (typeof response !== 'object') {
    return { code:500, msg:"Server error: " + errorReason, data: null};
  }


  if (response.code !== 200) {
    response.msg = format("%s (%d)", response?.msg, response?.code);
  }

  return response;
}

export async function DELETE(uri: string):Promise<Root<object>> {
  let errorReason = ''
  const url = env.devHost + uri;
  const response  = await fetch(url, {method:'DELETE'}).then((resp) => resp.json()).catch((reason) => {
    errorReason = reason;
  });

  if (typeof response !== 'object') {
    return { code:500, msg:"Server error: " + errorReason, data: null};
  }


  if (response.code !== 200) {
    response.msg = format("%s (%d)", response?.msg, response?.code);
  }

  return response;
}
