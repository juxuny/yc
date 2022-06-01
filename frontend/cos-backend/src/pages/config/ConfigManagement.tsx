import React, { useRef, useState } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Config } from '@/services/cos/config';
import { useIntl } from 'umi';
import { Button, Popconfirm, Space, Tag } from 'antd';
import { FormattedMessage } from '@@/plugin-locale/localeExports';
import { PlusOutlined } from '@ant-design/icons';
import ProTable from '@ant-design/pro-table';
import type { ActionType, ProColumns } from '@ant-design/pro-table';
import ConfigEditorModal from '@/pages/config/component/ConfigEditorModal';
import { Formatter } from '@/utils/formatter';
import AsyncSelector from '@/components/AsyncSelector';
import { Namespace } from '@/services/cos/namespace';

export default (): React.ReactNode => {
  const intl = useIntl();
  const actionRef = useRef<ActionType>();
  const [visible, setVisible] = useState<boolean>(false);
  const [selectedData, setSelectedData] = useState<API.Config.SaveReq>();
  const loadData = async (
    params: API.QueryParams<API.Config.ListReq>,
  ): Promise<{ data: API.Config.ListItem[]; success: boolean; total: number }> => {
    const { current, pageSize, ...args } = params;
    try {
      const resp = await Config.list({
        ...args,
        pagination: { pageNum: current, pageSize: pageSize },
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

  const showEditor = (record: API.Config.ListItem) => {
    setSelectedData({
      id: record.id,
      configId: record.configId,
      namespaceId: '',
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

  const deleteNamespace = async (record: API.Config.ListItem) => {
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
      title: intl.formatMessage({ id: 'pages.config.namespace.column.namespace' }),
      dataIndex: 'namespace',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.action.search' }),
      dataIndex: 'searchKey',
      hideInTable: true,
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
          {record.isDisabled ? (
            <Popconfirm
              key={'enable'}
              title={intl.formatMessage({ id: 'pages.config.namespace.confirm.enable' })}
              cancelText={intl.formatMessage({ id: 'pages.confirm.cancel' })}
              okText={intl.formatMessage({ id: 'pages.confirm.ok' })}
              onConfirm={async () => {
                await updateStatus(record, false);
              }}
            >
              <a>
                <FormattedMessage id={'pages.action.enable'} />
              </a>
            </Popconfirm>
          ) : (
            <Popconfirm
              key={'disable'}
              title={intl.formatMessage({ id: 'pages.config.namespace.confirm.disable' })}
              cancelText={intl.formatMessage({ id: 'pages.confirm.cancel' })}
              okButtonProps={{ type: 'primary' }}
              okType={'danger'}
              okText={intl.formatMessage({ id: 'pages.confirm.ok' })}
              onConfirm={async () => {
                await updateStatus(record, true);
              }}
            >
              <a style={{ color: 'red' }}>
                <FormattedMessage id={'pages.action.disable'} />
              </a>
            </Popconfirm>
          )}
          <Popconfirm
            key={'delete'}
            title={intl.formatMessage({ id: 'pages.config.namespace.confirm.delete' })}
            cancelText={intl.formatMessage({ id: 'pages.confirm.cancel' })}
            okButtonProps={{ type: 'primary' }}
            okType={'danger'}
            okText={intl.formatMessage({ id: 'pages.confirm.ok' })}
            onConfirm={async () => await deleteNamespace(record)}
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
            key="button"
            icon={<PlusOutlined />}
            type="primary"
            onClick={() => {
              showEditor({} as API.Config.ListItem);
            }}
          >
            <FormattedMessage id="pages.action.create" />
          </Button>,
          <AsyncSelector
            all={true}
            request={async () => {
              try {
                const resp = await Namespace.selector();
                if (resp && resp.code === 0) {
                  return Promise.resolve(resp.data || ({} as API.SelectorResp));
                }
              } catch (err) {
                console.error(err);
              }
              return Promise.resolve({} as API.SelectorResp);
            }}
          />,
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
