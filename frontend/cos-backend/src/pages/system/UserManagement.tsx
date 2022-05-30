import React, { useRef, useState } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { PlusOutlined } from '@ant-design/icons';
import { Button, Tag, Space, Popconfirm } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import { User } from '@/services/cos/user';
import CreateOrEditUserModal from '@/pages/system/dialog/CreateOrEditUserModal';
import { useIntl } from 'umi';
import { FormattedMessage } from '@@/plugin-locale/localeExports';

export default (): React.ReactNode => {
  const intl = useIntl();
  const actionRef = useRef<ActionType>();
  const loadData = async (
    params: API.QueryParams<API.UserListReq>,
  ): Promise<{ data: API.UserListItem[]; success: boolean; total: number }> => {
    const { current, pageSize, ...args } = params;
    const resp = await User.userList({
      ...args,
      pagination: { pageNum: current, pageSize: pageSize },
    });
    return {
      data: resp.data?.list || [],
      success: true,
      total: resp.data?.pagination.total || 0,
    };
  };

  const [editVisible, setEditVisible] = useState(false);
  const [selectedUserData, setSelectedUserData] = useState<API.SaveUserInfoReq | undefined>(
    undefined,
  );

  const showEditor = (userData: API.SaveUserInfoReq, visible: boolean) => {
    setEditVisible(visible);
    console.log(userData);
    setSelectedUserData(userData);
  };

  const updateStatus = async (userData: API.UserListItem, isDisabled: boolean) => {
    const req = {
      userId: userData.id || '',
      isDisabled: isDisabled || false,
    }
    try {
      const resp = await User.updateStatus(req);
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.log(err);
    }
  };

  const columns: ProColumns<API.UserListItem>[] = [
    {
      title: intl.formatMessage({ id: 'pages.system.userManagement.column.id' }),
      dataIndex: 'id',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.userManagement.column.identifier' }),
      dataIndex: 'identifier',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.userManagement.column.nick' }),
      dataIndex: 'nick',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.userManagement.query.search' }),
      dataIndex: 'searchKey',
      hideInTable: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.userManagement.column.accountType' }),
      dataIndex: 'accountType',
      hideInTable: true,
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.system.userManagement.column.isDisabled' }),
      dataIndex: 'isDisabled',
      valueType: 'select',
      valueEnum: {
        all: { text: intl.formatMessage({ id: 'pages.status.all' }), status: 'All'},
        enabled: { text: intl.formatMessage({ id: 'pages.status.enable' }), status: 'Enabled'},
        disabled: { text: intl.formatMessage({ id: 'pages.status.disable' }), status: 'Disabled'},
      },
      search: {
        transform: (value) => {
          if (value === "all") {
            return { isDisabled: undefined }
          } else if (value === "enabled") {
            return { isDisabled: 0 }
          } else if (value == "disabled") {
            return { isDisabled: 1 }
          } else {
            return {}
          }
        }
      },
      hideInSearch: false,
      render: (node, record) => {
        return <Tag color={record.isDisabled ? 'error' : 'success'}>{record.isDisabled ? '禁用' : '启用'}</Tag>
      },
    },
    {
      title: intl.formatMessage({ id: 'pages.action' }),
      key: 'action',
      hideInSearch: true,
      render: (node, record) => <Space>
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
        {
          record.isDisabled ?
            <Popconfirm key={'enable'}
                        title={intl.formatMessage({id: 'pages.system.userManagement.confirm.enable'})}
                        cancelText={intl.formatMessage({id: 'pages.confirm.cancel'})}
                        okText={intl.formatMessage({id: 'pages.confirm.ok'})}
                        onConfirm={async () => {
                          await updateStatus(record, false);
            }}>
              <a><FormattedMessage id={'pages.action.enable'}/></a>
            </Popconfirm>
             :
            <Popconfirm key={'disable'}
                        title={intl.formatMessage({id: 'pages.system.userManagement.confirm.disable'})}
                        cancelText={intl.formatMessage({id: 'pages.confirm.cancel'})}
                        okButtonProps={{type: 'primary'}}
                        okType={'danger'}
                        okText={intl.formatMessage({id: 'pages.confirm.ok'})} onConfirm={async () => {
                          await updateStatus(record, true);
            }}>
              <a style={{color: 'red'}}><FormattedMessage id={'pages.action.disable'}/></a>
            </Popconfirm>
        }
      </Space>,
    },
  ];

  return (
    <PageContainer>
      <ProTable<API.UserListItem, API.UserListReq>
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
