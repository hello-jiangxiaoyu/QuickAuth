import Mock from 'mockjs';
import setupMock from '@/utils/setupMock';

const haveReadIds = [];
const getMessageList = () => {
  return [
    {
      id: 1,
      type: 'message',
      title: '郑曦月',
      subTitle: '的私信',
      avatar:
        '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/8361eeb82904210b4f55fab888fe8416.png~tplv-uwbnlip3yd-webp.webp',
      content: '审批请求已发送，请查收',
      time: '今天 12:30:01',
    },
  ].map((item) => ({
    ...item,
    status: haveReadIds.indexOf(item.id) === -1 ? 0 : 1,
  }));
};

setupMock({
  setup: () => {
    Mock.mock(new RegExp('/api/message/list'), () => {
      return getMessageList();
    });

    Mock.mock(new RegExp('/api/message/read'), (params) => {
      const { ids } = JSON.parse(params.body);
      haveReadIds.push(...(ids || []));
      return true;
    });
  },
});
