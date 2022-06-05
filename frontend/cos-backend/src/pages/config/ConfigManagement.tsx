import React, {useRef, useState, useEffect} from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Config } from '@/services/cos/config';
import { useIntl } from 'umi';
import {Button, Popconfirm, Space, Tag, message} from 'antd';
import type { FormInstance } from 'antd';
import { FormattedMessage } from '@@/plugin-locale/localeExports';
import { PlusOutlined } from '@ant-design/icons';
import ProTable from '@ant-design/pro-table';
import type { ActionType, ProColumns } from '@ant-design/pro-table';
import ConfigEditorModal from '@/pages/config/component/ConfigEditorModal';
import { Formatter } from '@/utils/formatter';
import { Namespace } from '@/services/cos/namespace';
import { history } from 'umi';

export default (): React.ReactNode => {
  const intl = useIntl();
  const actionRef = useRef<ActionType>();
  const formRef = useRef<FormInstance<API.Config.ListReq> | undefined>();
  const [visible, setVisible] = useState<boolean>(false);
  const [selectedData, setSelectedData] = useState<API.Config.SaveReq>();
  const loadData = async (
    params: API.QueryParams<API.Config.ListReq>,
  ): Promise<{ data: API.Config.ListItem[]; success: boolean; total: number }> => {
    const { current, pageSize, ...args } = params;
    const selectedNamespaceId = formRef.current?.getFieldsValue().namespaceId;
    if (selectedNamespaceId === undefined || selectedNamespaceId === 'undefined') {
      message.error(intl.formatMessage({ id: 'pages.config.config-management.error.missingNamespaceId' }));
      return { data: [], success: false, total: 0 };
    }
    try {
      const resp = await Config.list({
        ...args,
        pagination: { pageNum: current, pageSize: pageSize },
        namespaceId: selectedNamespaceId,
      });
      return {
        data: resp.data?.list || [],
        success: true,
        total: resp.data?.pagination.total || 0,
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

  const showEditor = (record: API.Config.ListItem) => {
    const selectedNamespaceId = formRef.current?.getFieldsValue().namespaceId;
    setSelectedData({
      id: record.id,
      configId: record.configId,
      namespaceId: selectedNamespaceId || 0,
      baseId: undefined,
    });
    setVisible(true);
  };

  const updateStatus = async (record: API.Config.ListItem, isDisabled: boolean) => {
    try {
      const resp = await Config.updateStatus({
        id: record.id,
        isDisabled: isDisabled,
      });
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.error(err);
    }
  };

  const deleteConfig = async (record: API.Config.ListItem) => {
    try {
      const resp = await Config.deleteOne({
        id: record.id,
      });
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.error(err);
    }
  };

  const columns: ProColumns<API.Config.ListItem>[] = [
    {
      title: intl.formatMessage({ id: 'pages.config.namespace.column.id' }),
      dataIndex: 'id',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.config.config-management.column.configId' }),
      dataIndex: 'configId',
      hideInSearch: true,
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
          const resp = await Namespace.selector();
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
      title: intl.formatMessage({ id: 'pages.config.namespace.column.createTime' }),
      dataIndex: 'createTime',
      hideInTable: false,
      hideInSearch: true,
      renderText: Formatter.convertTimestampFromMillionSeconds,
    },
    {
      title: intl.formatMessage({ id: 'pages.config.namespace.column.updateTime' }),
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
              });
            }}
          >
            <FormattedMessage id={'pages.action.edit'} />
          </a>
          <Popconfirm
            key={'enable'}
            title={intl.formatMessage({ id: record.isDisabled ? 'pages.config.namespace.confirm.enable' : 'pages.config.namespace.confirm.disable' })}
            cancelText={intl.formatMessage({ id: 'pages.confirm.cancel' })}
            okText={intl.formatMessage({ id: 'pages.confirm.ok' })}
            onConfirm={async () => {
              await updateStatus(record, !record.isDisabled);
            }}
          >
            <a>
              <FormattedMessage id={record.isDisabled ? 'pages.action.enable' : 'pages.action.disable'} />
            </a>
          </Popconfirm>
          <Popconfirm
            key={'delete'}
            title={intl.formatMessage({ id: 'pages.config.namespace.confirm.delete' })}
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
      <ProTable<API.Config.ListItem, API.Config.ListReq>
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
            icon={<PlusOutlined />}
            type="primary"
            onClick={() => {
              showEditor({} as API.Config.ListItem);
            }}
          >
            <FormattedMessage id="pages.action.create" />
          </Button>
        ]}
      />
      <ConfigEditorModal
        visible={visible}
        onChangeVisible={setVisible}
        oldData={selectedData}
        onSuccess={() => {
          actionRef.current?.reload();
        }}
      />
    </PageContainer>
  );
};
