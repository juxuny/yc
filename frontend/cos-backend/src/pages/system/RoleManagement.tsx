
import type {ReactNode} from 'react';
import React, { useRef, useEffect, useState} from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import type {ActionType, ProColumns} from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import { useIntl } from 'umi';
import {System} from "@/services/shop/system";
import moment from "moment";
import { Button, Popover, Select, Tag} from "antd";
import {PlusOutlined} from "@ant-design/icons";
import {StatusEnumMap} from "@/services/shop/enum";
import CreateOrEditRoleModal from "@/pages/system/dialog/CreateOrEditRoleModal";

const { Option } = Select;

type PopoverStateMap = Record<string, boolean>;


export default (): React.ReactNode => {
  const intl = useIntl();
  const [pagination, setPagination] = useState<API.PaginationState>({page: 1, pageSize: 10});
  const [isLoading, setLoading] = useState<boolean>(false);
  const [popoverStateMap, setPopoverStateMap] = useState<PopoverStateMap>({});
  const tableRef = useRef<ActionType>();
  const [roleTypeValueEnumMap, setRoleTypeValueEnumMap] = useState<API.ValueEnum|undefined>(undefined);
  const [roleTypeValueEnumList, setRoleTypeValueEnumList] = useState<API.SelectorItem[]>([]);


  useEffect(() => {
    if (roleTypeValueEnumMap) return;
    System.getRoleTypeSelector().then(data => {
      if (data.code === 200 && data.result) {
        const valueEnum: API.ValueEnum = {};
        valueEnum[0] = { text: intl.formatMessage({id: 'pages.status.all'}) };
        for (let i = 0; data.result && data.result.list && i < data.result.list.length; i +=1) {
          const item = data.result.list[i];
          valueEnum[item.value !== undefined ? item.value : 'undefined'] = {text: item.label}
        }
        setRoleTypeValueEnumMap(valueEnum);
        if (data.result && data.result.list) setRoleTypeValueEnumList(data.result.list);
      }
    });
  })

  const [editorVisible, setEditorVisible] = useState<boolean>(false);
  const [editingData, setEditingData] = useState<API.SystemGetRoleItem>({});

  const onClickUpdateStatus = (e: React.MouseEvent<HTMLElement>, id?: number, status?: API.EnableStatus) => {
    const req: API.SystemUpdateRoleStatusReq = {};
    req.id = id;
    req.status = status === 1 ? 0: 1;
    setLoading(true);
    System.updateRoleStatus(req).then(() => {
      setPopoverStateMap({});
      tableRef.current?.reload();
    }).finally(() => {
      setLoading(false);
    });
  }

  const onHidePopover = () => {
    setPopoverStateMap({});
  }

  const getPopoverState = (id?: number) => {
    if (id !== undefined) {
      return popoverStateMap[id] === true;
    }
    return false;
  }

  const onClickEdit = (oldData: API.SystemGetRoleItem) => {
    setEditingData(oldData);
    setEditorVisible(true);
  }

  const genColumns = (): ProColumns<API.SystemGetRoleItem>[] => {
    return [
      {
        title: intl.formatMessage({id: 'pages.system.roleManagement.column.id', defaultMessage: 'ID'}),
        width: 80,
        dataIndex: 'id',
        search: false,
        render: (_) => <a>{_}</a>,
      },
      {
        title: intl.formatMessage({id: 'pages.system.roleManagement.column.type', defaultMessage: 'Type'}),
        dataIndex: 'type',
        width: 160,
        align: 'left',
        // valueType: 'select',
        initialValue: '0',
        valueEnum: roleTypeValueEnumMap,
        renderFormItem: (item, props) => {
          const {defaultRender, ...rest} = props;
          return <Select value={item.type} {...rest}>
            {
              roleTypeValueEnumList.filter(optionItem => {
                return optionItem && optionItem.value
              }).map((optionItem) => {
                return <Option key={optionItem.value} value={optionItem.value ? optionItem.value : ''}>{optionItem.label}</Option>
              })
            }
          </Select>
        }
      },
      {
        title: intl.formatMessage({id: 'pages.system.roleManagement.column.name', defaultMessage: 'Role Name'}),
        dataIndex: 'name',
        width: 180,
        align: 'left',
        search: false,
      },
      {
        title: intl.formatMessage({id: 'pages.system.roleManagement.column.status', defaultMessage: 'Status'}),
        width: 120,
        dataIndex: 'status',
        initialValue: "all",
        valueEnum: {
          all: { text: intl.formatMessage({id: 'pages.status.all'}), },
          enable: { text: intl.formatMessage({id: 'pages.status.enable'}) },
          disable: { text: intl.formatMessage({id: 'pages.status.disable'}) },
        },
        render: (_, item) => {
          return <Tag color={item.status === 1 ? 'green' : 'red'} key={item.id}>
            {item.status === 1 ? intl.formatMessage({id: 'pages.status.enable'}) : intl.formatMessage({id: 'pages.status.disable'})}
          </Tag>
        }
      },
      {
        title: intl.formatMessage({id: 'pages.system.roleManagement.column.createTime', defaultMessage: 'Create Time'}),
        width: 180,
        dataIndex: 'createTime',
        search: false,
        render: (_: ReactNode, item: API.SystemGetRoleItem) => <span>{moment(item.createTime).format("YYYY-DD-MM hh:mm:ss")}</span>
      },
      {
        title: intl.formatMessage({id: 'pages.system.roleManagement.column.updateTime', defaultMessage: 'Update Time'}),
        width: 180,
        dataIndex: 'updateTime',
        search: false,
        render: (_: ReactNode, item: API.SystemGetRoleItem) => <span>{moment(item.updateTime).format("YYYY-DD-MM hh:mm:ss")}</span>
      },
      {
        title: intl.formatMessage({id: 'pages.system.roleManagement.column.remark', defaultMessage: 'Remark'}),
        dataIndex: 'remark',
        search: false,
        ellipsis: true,
        copyable: false,
      },
      {
        title: intl.formatMessage({id: 'pages.system.roleManagement.column.action', defaultMessage: 'Action'}),
        width: 180,
        key: 'option',
        valueType: 'option',
        render: (_: ReactNode, item: API.SystemGetRoleItem) => [
          <Popover
            key={'status'}
            visible={getPopoverState(item.id)}
            content={[
              <Button key="cancel" type="default" size={'small'} onClick={onHidePopover}>{intl.formatMessage({id: 'pages.confirm.cancel'})}</Button>,
              <a key="blank">{ " " }</a>,
              <Button key='ok' type="primary" size={'small'} onClick={e => {onClickUpdateStatus(e, item.id, item.status)}}>{intl.formatMessage({id: 'pages.confirm.ok'})}</Button>
            ]}
            title={intl.formatMessage({id: item.status === 1 ? 'pages.status.disable.confirm' : 'pages.status.enable.confirm'})}
            trigger="click"
          >
            <a onClick={() => {
                 const m: PopoverStateMap = {};
                 if (item.id !== undefined) m[item.id] = true;
                 setPopoverStateMap(m);
               }}
               style={{color: item.status === 1 ? "red" : ""}}
            >
              {intl.formatMessage({
                id: item.status === 1 ? 'pages.status.disable' : 'pages.status.enable',
                defaultMessage: 'Enable'
              })}
            </a>
          </Popover>,
          <a key={'edit'} onClick={() => {onClickEdit(item)}}>{intl.formatMessage({id: 'pages.action.edit'})}</a>
        ],
      },
    ];
  }

  const onClickCreate = () => {
    setEditorVisible(true);
    setEditingData({});
  }

  return (
    <PageContainer>
      <ProTable<API.SystemGetRoleItem>
        columns={genColumns()}
        request={async (params) => {
          setLoading(true);
          const data = await System.getRoles({
            status: StatusEnumMap[params.status],
            type: parseInt(params.type, 10) === 0 ? -1 : parseInt(params.type, 10),
          }, {page: params.current, pageSize: params.pageSize});
          setLoading(false);
          const res = data?.result?.list;
          setPagination({page: params.current, pageSize: params.pageSize, total: data.result?.total});
          return Promise.resolve({data: res});
        }}
        rowKey="id"
        pagination={{
          showQuickJumper: true,
          pageSize: pagination.pageSize,
          total: pagination.total,
        }}
        actionRef={tableRef}
        search={{}}
        loading={isLoading}
        dateFormatter="string"
        toolBarRender={() => {
          return [
            <Button key="button" data-id={1} icon={<PlusOutlined />} type="primary" onClick={onClickCreate}>
              {intl.formatMessage({id: 'pages.action.create'})}
            </Button>
          ]
        }}
      />
      <CreateOrEditRoleModal visible={editorVisible} onChangeVisible={setEditorVisible} oldData={editingData} onSuccess={() => {
        setEditorVisible(false);
        tableRef.current?.reload();
      }}/>
    </PageContainer>
  );
};
