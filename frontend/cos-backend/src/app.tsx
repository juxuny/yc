import type { Settings as LayoutSettings } from '@ant-design/pro-layout';
import { PageLoading } from '@ant-design/pro-layout';
import type { RunTimeLayoutConfig } from 'umi';
import { history, Link } from 'umi';
import RightContent from '@/components/RightContent';
import Footer from '@/components/Footer';
import { BookOutlined, LinkOutlined } from '@ant-design/icons';
import { User } from '@/services/cos/user';
import { timestampInMinute } from '@/utils/func';

const isDev = process.env.NODE_ENV === 'development';
const loginPath = '/user/login';

/** 获取用户信息比较慢的时候会展示一个 loading */
export const initialStateConfig = {
  loading: <PageLoading />,
};

/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * */
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  currentUser?: API.User.InfoResp;
  fetchUserInfo?: () => Promise<API.User.InfoResp | undefined>;
}> {
  const fetchUserInfo = async () => {
    try {
      const userInfoResp = await User.userInfo();
      return userInfoResp.data;
    } catch (error) {
      history.push(loginPath);
    }
    return undefined;
  };
  // 如果是登录页面，不执行
  if (history.location.pathname !== loginPath) {
    const currentUser = await fetchUserInfo();
    return {
      fetchUserInfo,
      currentUser,
      settings: {},
    };
  }
  return {
    fetchUserInfo,
    settings: {},
  };
}

// ProLayout 支持的api https://procomponents.ant.design/components/layout
export const layout: RunTimeLayoutConfig = ({ initialState }) => {
  return {
    rightContentRender: () => <RightContent />,
    disableContentMargin: false,
    waterMarkProps: {
      content: initialState?.currentUser?.nick,
    },
    footerRender: () => <Footer />,
    onPageChange: () => {
      const { location } = history;

      if (location.pathname.indexOf('_t_=') === -1) {
        if (location.pathname.indexOf('?') === -1) {
          history.push(location.pathname + '?_t_=' + timestampInMinute());
        } else {
          history.push(location.pathname + '&_t_=' + timestampInMinute());
        }
      }
      // 如果没有登录，重定向到 login
      if (!initialState?.currentUser && location.pathname.indexOf(loginPath) === -1) {
        history.push(loginPath + '?_t_=' + timestampInMinute());
      }
    },
    links: isDev
      ? [
          <Link to="/umi/plugin/openapi" target="_blank" key={'api'}>
            <LinkOutlined />
            <span>OpenAPI 文档</span>
          </Link>,
          <Link to="/~docs" key={'component'}>
            <BookOutlined />
            <span>业务组件文档</span>
          </Link>,
        ]
      : [],
    menuHeaderRender: undefined,
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    ...initialState?.settings,
  };
};
