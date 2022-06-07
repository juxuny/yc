import {Button, message, Popconfirm, Space, Tag} from 'antd';
import useMergedState from 'rc-util/es/hooks/useMergedState';
import React, { useRef, useState } from 'react';
import {FormattedMessage, useIntl} from '@@/plugin-locale/localeExports';
import type { ProFormInstance } from '@ant-design/pro-components';
import { KeyValue } from '@/services/cos/key_value';
import type { ActionType } from '@ant-design/pro-components';
import ProTable from "@ant-design/pro-table";
import type {ProColumns} from '@ant-design/pro-table';
import {PlusOutlined} from "@ant-design/icons";
import KeyValueEditorModel  from "@/pages/config/component/KeyValueEidtorModal";
import {Formatter} from "@/utils/formatter";

export type KeyValuePairsProps = {
  reqData?: API.KeyValue.ListReq;
};

const KeyValuePairs: React.FC<KeyValuePairsProps> = (props) => {
  const intl = useIntl();
  const formRef = useRef<ProFormInstance<API.Config.SaveReq> | undefined>();
  const actionRef = useRef<ActionType>();
  const [editorVisible, setEditorVisible] = useState<boolean>(false);
  const [selectedData, setSelectedData] = useState<API.KeyValue.SaveReq|undefined>();
  const [reqData]  = useMergedState<API.KeyValue.ListAllReq>(
    {} as API.KeyValue.ListAllReq,
    {
      value: props.reqData
    }
  );

  const loadData = async (
    params: API.QueryParams<API.KeyValue.ListReq>,
  ): Promise<{ data: API.KeyValue.ListItem[]; success: boolean; total: number }> => {
    const { current, pageSize, ...args } = params;
    try {
      if (!reqData || !reqData.configId) {
        message.error(intl.formatMessage({ id: 'pages.config.key-value.missing.configId' }));
        return { data: [], success: false, total: 0 };
      }
      const resp = await KeyValue.listAll({
        ...args,
        configId: reqData.configId
      });
      return {
        data: resp.data?.list || [],
        success: true,
        total: resp.data?.list?.length || 0,
      };
    } catch (err) {
      console.error(err);
    }
    return { data: [], success: false, total: 0 };
  };

  const updateStatus = async (record: API.KeyValue.ListItem, isDisabled: boolean) => {
    try {
      const resp = await KeyValue.updateStatus({
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

  const columns: ProColumns<API.KeyValue.ListItem>[] = [
    {
      title: intl.formatMessage({ id: 'pages.column.id' }),
      dataIndex: 'id',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.config.key-value.column.configKey' }),
      dataIndex: 'configKey',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.config.key-value.column.configValue' }),
      dataIndex: 'configValue',
      hideInSearch: true,
    },
    {
      title: intl.formatMessage({ id: 'pages.config.key-value.column.isHot' }),
      dataIndex: 'isHot',
      valueType: 'select',
      valueEnum: {
        all: { text: intl.formatMessage({ id: 'pages.status.all' }), status: 'All' },
        hot: { text: intl.formatMessage({ id: 'pages.config.key-value.column.isHot' }), status: 'hot' },
        cold: { text: intl.formatMessage({ id: 'pages.config.key-value.column.cold' }), status: 'cold' },
      },
      search: {
        transform: (value) => {
          if (value === 'all') {
            return { isHot: undefined };
          } else if (value === 'hot') {
            return { isHot: true };
          } else if (value == 'cold') {
            return { isHot: false };
          } else {
            return {};
          }
        },
      },
      hideInSearch: false,
      render: (node, record) => {
        return (
          <Tag color={record.isHot ? 'error' : 'success'}>
            <FormattedMessage id={record.isHot ? 'pages.config.key-value.column.isHot' : 'pages.config.key-value.column.cold'}/>
          </Tag>
        );
      },
    },
    {
      title: intl.formatMessage({ id: 'pages.column.isDisabled' }),
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
            <FormattedMessage id={record.isDisabled ? 'pages.status.disable': 'pages.status.enable'}/>
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
      title: intl.formatMessage({ id: 'pages.action.search' }),
      dataIndex: 'searchKey',
      hideInTable: true,
    },
    {
      title: intl.formatMessage({id: 'pages.action'}),
      key: 'action',
      hideInSearch: true,
      render: (node, record) => (
        <Space>
          <a
            key={'edit'}
            onClick={() => {
              setSelectedData({
                id: record.id,
                configId: record.configId,
                configKey: record.configKey,
                configValue: record.configValue,
                isHot: record.isHot,
              } as API.KeyValue.SaveReq);
              setEditorVisible(true);
            }}
          >
            <FormattedMessage id={'pages.action.edit'} />
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
        </Space>
      )
    }
  ];
  return (
    <>
      <ProTable
        request={loadData}
        rowKey='id'
        formRef={formRef}
        actionRef={actionRef}
        columns={columns}
        dateFormatter="string"
        headerTitle={false}
        pagination={false}
        toolBarRender={() => [
          <Button
            key="button"
            icon={<PlusOutlined />}
            type="primary"
            onClick={() => {
              setSelectedData({
                configId: props.reqData?.configId || '',
                id: undefined,
                configKey: '',
                configValue: ''
              } as API.KeyValue.SaveReq);
              setEditorVisible(true);
            }}
          >
            <FormattedMessage id="pages.action.create" />
          </Button>,
        ]}
      />
      <KeyValueEditorModel visible={editorVisible} onChangeVisible={setEditorVisible} oldData={selectedData} onSuccess={() => {
        actionRef.current?.reload();
      }}/>
    </>
  );
};

export default KeyValuePairs;
