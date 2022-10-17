import React, {useRef, useState} from 'react';
import {PageContainer} from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import {PlusOutlined} from '@ant-design/icons';
import {Button, Space, Popconfirm} from 'antd';
import type {ProColumns, ActionType} from '@ant-design/pro-table';
import CreateAccessKeyModal from '@/pages/account/component/CreateAccessKeyModal';
import type {CreateResult} from '@/pages/account/component/CreateAccessKeyModal';
import {useIntl} from 'umi';
import CreateResultModal from '@/pages/account/component/CreateResultModal';
import {FormattedMessage} from '@@/plugin-locale/localeExports';
import {ColumnBuilder} from "@/utils/column_builder";
import {Formatter} from "@/utils/formatter";
import RemarkPopoverEditor from "@/pages/account/component/RemarkPopoverEditor";
import { cos } from '@/services/api';
import type { AccessKeyListRequest, AccessKeyItem, DeleteAccessKeyRequest, UpdateStatusAccessKeyRequest, CreateAccessKeyRequest} from '@/services/api/typing';
import type {QueryParams} from "@juxuny/yc-ts-data-type/typing";

export default (): React.ReactNode => {
  const intl = useIntl();
  const actionRef = useRef<ActionType>();
  const loadData = async (
    params: QueryParams<AccessKeyListRequest>,
  ): Promise<{ data: AccessKeyItem[]; success: boolean; total: number }> => {
    const {current, pageSize, ...args} = params;
    const resp = await cos.accessKeyList({
      ...args,
      pagination: {pageNum: current || 1, pageSize: pageSize || 10},
    } as AccessKeyListRequest);
    return {
      data: resp.data?.list || [],
      success: true,
      total: resp.data?.pagination?.total || 0,
    };
  };

  const [editVisible, setEditVisible] = useState(false);
  const [visibleAccessKeyId, setVisibleAccessKeyId] = useState<string | number | undefined>();
  const [selectedUserData, setSelectedUserData] = useState<CreateAccessKeyRequest | undefined>(undefined);
  const [createResult, setCreateResult] = useState<CreateResult | undefined>(undefined);

  const showEditor = (userData: CreateAccessKeyRequest) => {
    setEditVisible(true);
    setSelectedUserData(userData);
  };

  const updateStatus = async (userData: AccessKeyItem, isDisabled: boolean) => {
    const req: UpdateStatusAccessKeyRequest = {
      id: userData.id,
      isDisabled: isDisabled || false,
    };
    try {
      const resp = await cos.updateStatusAccessKey(req);
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.log(err);
    }
  };

  const deleteAccessKey = async (userData: AccessKeyItem) => {
    const req: DeleteAccessKeyRequest = {
      id: userData.id,
    };
    try {
      const resp = await cos.deleteAccessKey(req);
      if (resp && resp.code === 0) {
        actionRef.current?.reload();
      }
    } catch (err) {
      console.error(err);
    }
  };

  const columnBuilder = new ColumnBuilder<AccessKeyItem>();

  const columns: ProColumns<AccessKeyItem>[] = [
    columnBuilder.id(),
    {
      title: intl.formatMessage({id: 'pages.account.access-key.column.accessKey'}),
      dataIndex: 'accessKey',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({id: 'pages.account.access-key.column.remark'}),
      dataIndex: 'remark',
      hideInSearch: true,
      render: (node, record) => <RemarkPopoverEditor data={record}
                                                     showPopup={visibleAccessKeyId === record.id}
                                                     onChangeVisible={setVisibleAccessKeyId}
                                                     onSuccess={() => actionRef.current?.reload()}/>,
    },
    {
      title: intl.formatMessage({id: 'pages.account.access-key.column.validStartTime'}),
      dataIndex: 'validStartTime',
      hideInSearch: true,
      renderText: Formatter.convertTimestampFromMillionSeconds,
    },
    {
      title: intl.formatMessage({id: 'pages.account.access-key.column.validEndTime'}),
      dataIndex: 'validEndTime',
      hideInSearch: true,
      renderText: Formatter.convertTimestampFromMillionSeconds,
    },
    columnBuilder.searchKey(),
    columnBuilder.isDisabled(),
    {
      title: intl.formatMessage({id: 'pages.column.createTime'}),
      dataIndex: 'createTime',
      hideInTable: false,
      hideInSearch: true,
      renderText: Formatter.convertTimestampFromMillionSeconds,
    },
    {
      title: intl.formatMessage({id: 'pages.column.updateTime'}),
      dataIndex: 'updateTime',
      hideInTable: false,
      hideInSearch: true,
      renderText: Formatter.convertTimestampFromMillionSeconds,
    },
    {
      title: intl.formatMessage({id: 'pages.action'}),
      key: 'action',
      hideInSearch: true,
      render: (node, record) => (
        <Space>
          <Popconfirm
            key={'enable'}
            title={intl.formatMessage({id: record.isDisabled ? 'pages.status.enable' : 'pages.status.disable'})}
            cancelText={intl.formatMessage({id: 'pages.confirm.cancel'})}
            okText={intl.formatMessage({id: 'pages.confirm.ok'})}
            onConfirm={async () => {
              await updateStatus(record, !record.isDisabled);
            }}
          >
            <a style={{color: record.isDisabled ? '' : 'red'}}>
              <FormattedMessage id={record.isDisabled ? 'pages.action.enable' : 'pages.action.disable'}/>
            </a>
          </Popconfirm>
          <Popconfirm
            key={'delete'}
            title={intl.formatMessage({id: 'pages.system.user-management.confirm.delete'})}
            cancelText={intl.formatMessage({id: 'pages.confirm.cancel'})}
            okButtonProps={{type: 'primary'}}
            okType={'danger'}
            okText={intl.formatMessage({id: 'pages.confirm.ok'})}
            onConfirm={async () => await deleteAccessKey(record)}
          >
            <a style={{color: 'red'}}>
              <FormattedMessage id={'pages.action.delete'}/>
            </a>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <PageContainer>
      <ProTable<AccessKeyItem, AccessKeyListRequest>
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
            icon={<PlusOutlined/>}
            type="primary"
            onClick={() => {
              showEditor(
                {} as AccessKeyItem
              );
            }}
          >
            <FormattedMessage id="pages.action.create"/>
          </Button>,
        ]}
      />
      <CreateAccessKeyModal
        visible={editVisible}
        data={selectedUserData}
        onChangeVisible={setEditVisible}
        onSuccess={(result: CreateResult) => {
          actionRef.current?.reload();
          setCreateResult(result);
        }}
      />
      <CreateResultModal
        visible={createResult !== undefined}
        onChangeVisible={v => {
          setCreateResult(v ? createResult : undefined);
        }}
        data={createResult}
      />
    </PageContainer>
  );
};
