import React, {useRef, useState, useEffect} from 'react';
import { useIntl } from 'umi';
import {Button, Popconfirm, Space, Tag, message, Drawer} from 'antd';
import type { FormInstance } from 'antd';
import { FormattedMessage } from '@@/plugin-locale/localeExports';
import ProTable from '@ant-design/pro-table';
import type { ActionType, ProColumns } from '@ant-design/pro-table';
import ConfigEditorModal from '@/pages/config/component/ConfigEditorModal';
import { Formatter } from '@/utils/formatter';
import { history } from 'umi';
import KeyValuePairs from "@/pages/config/component/KeyValuePairs";
import {cos} from "@/services/api";
import {
  ListConfigItem,
  ListConfigRequest,
  ListValueRequest,
  SaveConfigRequest,
  UpdateStatusConfigRequest
} from "@/services/api/typing";
import type {QueryParams} from "@juxuny/yc-ts-data-type/typing";
import {PageContainer} from "@ant-design/pro-layout";

const calculateDrawerWidth = () => {
  return window.innerWidth * 0.8;
}

export default (): React.ReactNode => {
  const intl = useIntl();
  const actionRef = useRef<ActionType>();
  const formRef = useRef<FormInstance<ListConfigRequest> | undefined>();
  const [editorVisible, setEditorVisible] = useState<boolean>(false);
  const [drawerVisible, setDrawerVisible] = useState<boolean>(false);
  const [selectedData, setSelectedData] = useState<SaveConfigRequest>();
  const [drawerWidth, setDrawerWidth] = useState<number>(calculateDrawerWidth);
  const [keyValueListReq, setKeyValueListReq] = useState<ListValueRequest | undefined>(undefined);
  const [isCloneEditor, setIsCloneEditor] = useState<boolean>(false);
  const loadData = async (
    params: QueryParams<ListConfigRequest>,
  ): Promise<{ data: ListConfigItem[]; success: boolean; total: number }> => {
    const { current, pageSize, ...args } = params;
    const selectedNamespaceId = formRef.current?.getFieldsValue().namespaceId;
    if (selectedNamespaceId === undefined || selectedNamespaceId === 'undefined') {
      message.error(intl.formatMessage({ id: 'pages.config.config-management.error.missingNamespaceId' }));
      return { data: [], success: false, total: 0 };
    }
    try {
      const resp = await cos.listConfig({
        ...args,
        pagination: { pageNum: current || 1, pageSize: pageSize || 10},
        namespaceId: selectedNamespaceId,
      });
      return {
        data: resp.data?.list || [],
        success: true,
        total: resp.data?.pagination?.total || 0,
      };
    } catch (err) {
      console.error(err);
    }
    return { data: [], success: false, total: 0 };
  };

  useEffect(() => {
    const { location } = history;
    const namespaceId = location.query?.namespaceId;
    if (namespaceId) {
      if (Array.isArray(namespaceId)) {
        formRef.current?.setFieldsValue({namespaceId: namespaceId[0]});
      } else {
        formRef.current?.setFieldsValue({namespaceId});
      }
    }
  });

  // auto resize drawer
  const updateDrawerWidth = () => {
    setDrawerWidth(calculateDrawerWidth());
  }
  useEffect(() => {
    window.addEventListener('resize', updateDrawerWidth);
    return () => {
      window.removeEventListener('resize', updateDrawerWidth);
    }
  });

  const showKeyValueDrawer = (configId: string) => {
    setDrawerVisible(true);
    setKeyValueListReq({
      configId: configId,
      searchKey: '',
      pagination: { pageNum: 1, pageSize: 10 }
    } as ListValueRequest);
  }

  const showEditor = (record: SaveConfigRequest, isClone: boolean) => {
    const selectedNamespaceId = formRef.current?.getFieldsValue().namespaceId;
    setIsCloneEditor(isClone);
    setSelectedData({
      id: record.id,
      configId: record.configId,
      namespaceId: selectedNamespaceId || 0,
      baseId: record.baseId,
    });
    setEditorVisible(true);
  };

  const updateStatus = async (record: UpdateStatusConfigRequest, isDisabled: boolean) => {
    try {
      const resp = await cos.updateStatusConfig({
        id: record.id ,
        isDisabled: isDisabled,
      });
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.error(err);
    }
  };

  const deleteConfig = async (record: ListConfigItem) => {
    try {
      const resp = await cos.deleteConfig({
        id: record.id,
      });
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.error(err);
    }
  };

  const columns: ProColumns<ListConfigItem>[] = [
    {
      title: intl.formatMessage({ id: 'pages.config.namespace.column.id' }),
      dataIndex: 'id',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.config.config-management.column.configId' }),
      dataIndex: 'configId',
      hideInSearch: true,
      render: (node, record) => {
        return <a onClick={() => showKeyValueDrawer(record.id ? record.id.toString() : '')}>
          { record.configId }
        </a>
      }
    },
    {
      title: intl.formatMessage({ id: 'pages.action.search' }),
      dataIndex: 'searchKey',
      hideInTable: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.config.namespace.column.namespace' }),
      dataIndex: 'namespaceId',
      hideInTable: true,
      valueType: 'select',
      request: async () => {
        try {
          const resp = await cos.selectorNamespace({});
          if (resp && resp.code === 0) {
            return resp.data?.list || [];
          }
          return [];
        } catch (e) {
          console.error(e);
        }
        return [];
      }
    },
    {
      title: intl.formatMessage({ id: 'pages.config.namespace.column.isDisabled' }),
      dataIndex: 'isDisabled',
      valueType: 'select',
      valueEnum: {
        all: { text: intl.formatMessage({ id: 'pages.status.all' }), status: 'All' },
        enabled: { text: intl.formatMessage({ id: 'pages.status.enable' }), status: 'Enabled' },
        disabled: { text: intl.formatMessage({ id: 'pages.status.disable' }), status: 'Disabled' },
      },
      search: {
        transform: (value) => {
          if (value === 'all') {
            return { isDisabled: undefined };
          } else if (value === 'enabled') {
            return { isDisabled: 0 };
          } else if (value == 'disabled') {
            return { isDisabled: 1 };
          } else {
            return {};
          }
        },
      },
      hideInSearch: false,
      render: (node, record) => {
        return (
          <Tag color={record.isDisabled ? 'error' : 'success'}>
            {record.isDisabled ? '禁用' : '启用'}
          </Tag>
        );
      },
    },
    {
      title: intl.formatMessage({ id: 'pages.column.createTime' }),
      dataIndex: 'createTime',
      hideInTable: false,
      hideInSearch: true,
      renderText: Formatter.convertTimestampFromMillionSeconds,
    },
    {
      title: intl.formatMessage({ id: 'pages.column.updateTime' }),
      dataIndex: 'updateTime',
      hideInTable: false,
      hideInSearch: true,
      renderText: Formatter.convertTimestampFromMillionSeconds,
    },
    {
      title: intl.formatMessage({ id: 'pages.action' }),
      key: 'action',
      hideInSearch: true,
      render: (node, record) => (
        <Space>
          <a
            key={'edit'}
            onClick={() => {
              showEditor({
                ...record,
                baseId: undefined
              }, false);
            }}
          >
            <FormattedMessage id={'pages.action.edit'} />
          </a>
          <a
            key={'derive'}
            onClick={() => {
              showEditor({
                ...record,
                configId: record.configId + '(copy)',
                baseId: record.id,
                id: undefined
              }, false);
            }}
          >
            <FormattedMessage id={'pages.action.derive'} />
          </a>
          <a
            key={'clone'}
            onClick={() => {
              showEditor({
                ...record,
                configId: record.configId + '(copy)',
                id: record.id
              }, true);
            }}
          >
            <FormattedMessage id={'pages.action.clone'} />
          </a>
          <Popconfirm
            key={'enable'}
            title={intl.formatMessage({ id: record.isDisabled ? 'pages.config.config-management.confirm.enable' : 'pages.config.config-management.confirm.disable' })}
            cancelText={intl.formatMessage({ id: 'pages.confirm.cancel' })}
            okText={intl.formatMessage({ id: 'pages.confirm.ok' })}
            onConfirm={async () => {
              await updateStatus(record, !record.isDisabled);
            }}
          >
            <a style={{color: record.isDisabled ? '' : 'red'}}>
              <FormattedMessage id={record.isDisabled ? 'pages.action.enable' : 'pages.action.disable'} />
            </a>
          </Popconfirm>
          <Popconfirm
            key={'delete'}
            title={intl.formatMessage({ id: 'pages.config.config-management.confirm.delete' })}
            cancelText={intl.formatMessage({ id: 'pages.confirm.cancel' })}
            okButtonProps={{ type: 'primary' }}
            okType={'danger'}
            okText={intl.formatMessage({ id: 'pages.confirm.ok' })}
            onConfirm={async () => await deleteConfig(record)}
          >
            <a style={{ color: 'red' }}>
              <FormattedMessage id={'pages.action.delete'} />
            </a>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <PageContainer>
      <ProTable<ListConfigItem, ListConfigRequest>
        formRef={formRef}
        request={loadData}
        actionRef={actionRef}
        columns={columns}
        rowKey="id"
        pagination={{
          showQuickJumper: true,
        }}
        dateFormatter="string"
        headerTitle={false}
        toolBarRender={() => [
          <Button
            key={'edit'}
            type="primary"
            onClick={() => {
              showEditor({} as SaveConfigRequest, false);
            }}
          >
            <FormattedMessage id="pages.action.create" />
          </Button>
        ]}
      />
      <ConfigEditorModal
        isClone={isCloneEditor}
        visible={editorVisible}
        onChangeVisible={setEditorVisible}
        oldData={selectedData}
        onSuccess={() => {
          actionRef.current?.reload();
        }}
      />
      <Drawer
        title={intl.formatMessage({ id: 'pages.config.config-management.drawer.title' })}
        placement="right"
        width={drawerWidth}
        onClose={() => setDrawerVisible(false)}
        visible={drawerVisible}
      >
        { drawerVisible && <KeyValuePairs reqData={keyValueListReq} /> }
      </Drawer>
    </PageContainer>
  );
};
