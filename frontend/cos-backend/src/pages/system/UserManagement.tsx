import React, {useRef, useState} from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import { User } from '@/services/cos/user';
import CreateOrEditUserModal from '@/pages/system/dialog/CreateOrEditUserModal';
import { useIntl } from 'umi';

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

  const [ editVisible, setEditVisible ] = useState(false);

  const columns: ProColumns<API.UserListItem>[] = [
    {
      title: intl.formatMessage({id: 'pages.system.userManagement.column.id'}),
      dataIndex: 'id',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({id: 'pages.system.userManagement.column.identifier'}),
      dataIndex: 'identifier',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({id: 'pages.system.userManagement.column.nick'}),
      dataIndex: 'nick',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({id: 'pages.system.userManagement.query.search'}),
      dataIndex: 'searchKey',
      hideInTable: true,
    }
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
      />
      <CreateOrEditUserModal visible={editVisible} onChangeVisible={setEditVisible} />
    </PageContainer>
  );
};
