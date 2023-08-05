import {configureStore} from "@reduxjs/toolkit";
import env from './env.json'

export interface GlobalState{
  menuWidth:number;
}
const initialState:GlobalState = {
  menuWidth: env.menuWidth,
};

const DispatchMenu = 'menu width';

export default function reducer(state = initialState, action) {
  switch (action.type) {
    case DispatchMenu: {
      const { appList } = action.payload;
      return {
        ...state,
        appList,
      };
    }
    default:
      return state;
  }
}

export const poolStore = configureStore({
  reducer,
  middleware: (getDefaultMiddleware) => {
    return getDefaultMiddleware()
  },
})

export function dispatchMenuCollapse(data) {
  poolStore.dispatch({
    type: DispatchMenu,
    payload: {appList: data},
  })
}

