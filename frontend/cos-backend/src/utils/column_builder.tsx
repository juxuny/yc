import {Tag} from "antd";
import { useIntl } from 'umi';
import React from 'react';
import type { ProColumns } from '@ant-design/pro-table';
import {FormattedMessage} from "@@/plugin-locale/localeExports";

export class ColumnBuilder<Type extends { isDisabled?: boolean }> {
  id = (): ProColumns<Type> => {
    const intl = useIntl();
    return {
      title: intl.formatMessage({ id: 'pages.column.id' }),
      dataIndex: 'id',
      hideInSearch: true,
    };
  }

  searchKey = (): ProColumns<Type> => {
    const intl = useIntl();
    return {
      title: intl.formatMessage({ id: 'pages.action.search' }),
      dataIndex: 'searchKey',
      hideInTable: true,
    };
  }

  isDisabled = (): ProColumns<Type> => {
    const intl = useIntl();
    return {
      title: intl.formatMessage({ id: 'pages.column.isDisabled' }),
        dataIndex: 'isDisabled',
      valueType: 'select',
      valueEnum: {
      all: { text: intl.formatMessage({ id: 'pages.status.all' }), status: 'All' },
      enabled: { text: intl.formatMessage({ id: 'pages.status.enable' }), status: 'Enabled' },
      disabled: { text: intl.formatMessage({ id: 'pages.status.disable' }), status: 'Disabled' },
    },
      search: {
        transform: (value: string | undefined) => {
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
      render: (node: React.ReactNode, record: Type) => {
          return (<Tag color={record.isDisabled ? 'error' : 'success'}><FormattedMessage id={record.isDisabled ? 'pages.status.disable' : 'pages.status.enable'}/></Tag>);
      },
    };
  }
}
