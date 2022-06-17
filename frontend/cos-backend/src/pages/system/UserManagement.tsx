import React, { useRef, useState } from 'react';
import ProTable from '@ant-design/pro-table';
import { Button, Tag, Space, Popconfirm } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import CreateOrEditUserModal from '@/pages/system/dialog/CreateOrEditUserModal';
import { useIntl } from 'umi';
import { FormattedMessage } from '@@/plugin-locale/localeExports';
import {cos} from "@/services/api";
import type {
  SaveOrCreateUserRequest,
  UserDeleteRequest,
  UserListItem,
  UserListRequest,
  UserUpdateStatusRequest
} from "@/services/api/typing";
import type {QueryParams} from "@juxuny/yc-ts-data-type/typing";
import {PageContainer} from "@ant-design/pro-layout";

export default (): React.ReactNode => {
  const intl = useIntl();
  const actionRef = useRef<ActionType>();
  const loadData = async (
    params: QueryParams<UserListRequest>,
  ): Promise<{ data: UserListItem[]; success: boolean; total: number }> => {
    const { current, pageSize, ...args } = params;
    const resp = await cos.userList({
      ...args,
      pagination: { pageNum: current || 1, pageSize: pageSize || 10},
    });
    return {
      data: resp.data?.list || [],
      success: true,
      total: resp.data?.pagination?.total || 0,
    };
  };

  const [editVisible, setEditVisible] = useState(false);
  const [selectedUserData, setSelectedUserData] = useState<SaveOrCreateUserRequest | undefined>(undefined);

  const showEditor = (userData: SaveOrCreateUserRequest, visible: boolean) => {
    setEditVisible(visible);
    console.log(userData);
    setSelectedUserData(userData);
  };

  const updateStatus = async (userData: UserListItem, isDisabled: boolean) => {
    const req: UserUpdateStatusRequest = {
      userId: userData.id || '',
      isDisabled: isDisabled || false,
    };
    try {
      const resp = await cos.userUpdateStatus(req);
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.log(err);
    }
  };

  const userDelete = async (userData: UserListItem) => {
    const req: UserDeleteRequest = {
      userId: userData.id || '',
    };
    try {
      const resp = await cos.userDelete(req);
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.error(err);
    }
  };

  const columns: ProColumns<UserListItem>[] = [
    {
      title: intl.formatMessage({ id: 'pages.system.user-management.column.id' }),
      dataIndex: 'id',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.user-management.column.identifier' }),
      dataIndex: 'identifier',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.user-management.column.nick' }),
      dataIndex: 'nick',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.user-management.query.search' }),
      dataIndex: 'searchKey',
      hideInTable: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.user-management.column.accountType' }),
      dataIndex: 'accountType',
      hideInTable: true,
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.user-management.column.isDisabled' }),
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
      title: intl.formatMessage({ id: 'pages.action' }),
      key: 'action',
      hideInSearch: true,
      render: (node, record) => (
        <Space>
          <a
            key={'edit'}
            onClick={() => {
              showEditor(
                {
                  ...record,
                  userId: record.id,
                },
                true,
              );
            }}
          >
            <FormattedMessage id={'pages.action.edit'} />
          </a>
          {record.isDisabled ? (
            <Popconfirm
              key={'enable'}
              title={intl.formatMessage({ id: 'pages.system.user-management.confirm.enable' })}
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
              title={intl.formatMessage({ id: 'pages.system.user-management.confirm.disable' })}
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
            title={intl.formatMessage({ id: 'pages.system.user-management.confirm.delete' })}
            cancelText={intl.formatMessage({ id: 'pages.confirm.cancel' })}
            okButtonProps={{ type: 'primary' }}
            okType={'danger'}
            okText={intl.formatMessage({ id: 'pages.confirm.ok' })}
            onConfirm={async () => await userDelete(record)}
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
      <ProTable<UserListItem, UserListRequest>
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
            type="primary"
            onClick={() => {
              showEditor(
                {
                  userId: undefined,
                  nick: '',
                  identifier: '',
                },
                true,
              );
            }}
          >
            <FormattedMessage id="pages.action.create" />
          </Button>,
        ]}
      />
      <CreateOrEditUserModal
        visible={editVisible}
        oldData={selectedUserData}
        onChangeVisible={setEditVisible}
        onSuccess={() => {
          actionRef.current?.reload();
        }}
      />
    </PageContainer>
  );
};
