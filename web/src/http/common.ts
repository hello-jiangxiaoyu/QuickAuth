import {format} from "util";
import env from "../store/env.json"
import {Message} from "@arco-design/web-react";

export interface Root<T> {
  code: number;
  msg: string;
  data?: T;
}

async function SendHttpRequest<T>(
  method: 'GET' | 'POST' | 'PUT' | 'DELETE',
  uri:string,
  data?:string|FormData
):Promise<Root<T>> {
  const url = env.devHost + uri;
  let err = ""
  const response:Root<T> = await fetch(url, {method, body: data}).then((resp) => resp.json()).catch(e => {err = e.toString();});
  if (err !== "") {
    Message.error(err);
    return Promise.reject("fetch error");
  }

  if (typeof response !== 'object') {
    Message.error("Invalid server response type");
    return Promise.reject("Invalid server response type");
  }

  if (response?.code > 308) {
    response.msg = format("%s (%d)", response?.msg, response?.code);
    Message.error(response.msg);
    return Promise.reject(response.msg);
  }

  return response;
}

export async function GET<T>(uri: string):Promise<Root<T>> {
  return await SendHttpRequest<T>('GET', uri);
}

export async function POST<T>(uri: string, data:object):Promise<Root<T>> {
  return await SendHttpRequest<T>('POST', uri, JSON.stringify(data));
}


export async function PUT(uri: string, data:object):Promise<Root<object>> {
  return await SendHttpRequest<object>('PUT', uri, JSON.stringify(data));
}

export async function DELETE(uri: string):Promise<Root<object>> {
  return await SendHttpRequest<object>('DELETE', uri);
}

export async function POSTForm<T>(uri: string, data:object):Promise<Root<T>> {
  const formData = new FormData(this);
  for (const key in data) {
    if (data.hasOwnProperty(key)) {
      formData.append(key, data[key]);
    }
  }
  return await SendHttpRequest<T>('POST', uri, formData);
}
