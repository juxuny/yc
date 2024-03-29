import React, { useRef, useState } from 'react';
import { useIntl, history } from 'umi';
import { Button, Popconfirm, Space, Tag } from 'antd';
import { FormattedMessage } from '@@/plugin-locale/localeExports';
import { PlusOutlined } from '@ant-design/icons';
import ProTable from '@ant-design/pro-table';
import type { ActionType, ProColumns } from '@ant-design/pro-table';
import NamespaceEditorModal from '@/pages/config/component/NamespaceEditorModal';
import { Formatter } from '@/utils/formatter';
import type {ListNamespaceItem, ListNamespaceRequest, SaveNamespaceRequest} from "@/services/api/typing";
import type {QueryParams} from "@juxuny/yc-ts-data-type/typing";
import {cos} from "@/services/api";
import {PageContainer} from "@ant-design/pro-layout";

export default (): React.ReactNode => {
  const intl = useIntl();
  const actionRef = useRef<ActionType>();
  const [visible, setVisible] = useState<boolean>(false);
  const [selectedData, setSelectedData] = useState<SaveNamespaceRequest>();
  const loadData = async (
    params: QueryParams<ListNamespaceRequest>,
  ): Promise<{ data: ListNamespaceItem[]; success: boolean; total: number }> => {
    const { current, pageSize, ...args } = params;
    try {
      const resp = await cos.listNamespace({
        ...args,
        pagination: { pageNum: current || 1, pageSize: pageSize || 10 },
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

  const showEditor = (record: ListNamespaceItem) => {
    setSelectedData(record);
    setVisible(true);
  };

  const updateStatus = async (record: ListNamespaceItem, isDisabled: boolean) => {
    try {
      const resp = await cos.updateStatusNamespace({
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

  const deleteNamespace = async (record: ListNamespaceItem) => {
    try {
      const resp = await cos.deleteNamespace({
        id: record.id,
      });
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.error(err);
    }
  };

  const columns: ProColumns<ListNamespaceItem>[] = [
    {
      title: intl.formatMessage({ id: 'pages.config.namespace.column.id' }),
      dataIndex: 'id',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.config.namespace.column.namespace' }),
      dataIndex: 'namespace',
      hideInSearch: true,
      render: (node, record) => {
        return <a onClick={() => {
          history.push('/config/config-management?namespaceId=' + record.id);
        }}>
          { record.namespace }
        </a>
      }
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
              });
            }}
          >
            <FormattedMessage id={'pages.action.edit'} />
          </a>
          <Popconfirm
            key={'enable'}
            title={intl.formatMessage({ id: record.isDisabled ? 'pages.config.namespace.confirm.enable' : 'pages.config.namespace.confirm.disable' })}
            cancelText={intl.formatMessage({ id: 'pages.confirm.cancel' })}
            okButtonProps={{ type: 'primary' }}
            okText={intl.formatMessage({ id: 'pages.confirm.ok' })}
            okType={record.isDisabled ? 'primary' : 'danger'}
            onConfirm={async () => {
              await updateStatus(record, !record.isDisabled);
            }}
          >
            <a style={{color: record.isDisabled ? '': 'red'}}>
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
      <ProTable<ListNamespaceItem, ListNamespaceRequest>
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
              showEditor({});
            }}
          >
            <FormattedMessage id="pages.action.create" />
          </Button>,
        ]}
      />
      <NamespaceEditorModal
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
