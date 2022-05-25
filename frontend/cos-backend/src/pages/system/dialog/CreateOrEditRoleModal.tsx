import {Button, Form, FormInstance, Input, message, Modal, Select} from 'antd';
import useMergedState from 'rc-util/es/hooks/useMergedState';
import React, {createRef, useEffect, useState} from "react";
import {useIntl} from "@@/plugin-locale/localeExports";
import {System} from "@/services/shop/system";

const { Option } = Select;
const {TextArea} = Input;

export type CreateOrEditRoleProps = {
  visible?: boolean;
  onChangeVisible: (v: boolean) => void;
  oldData?: API.SystemGetRoleItem;
  trigger?: JSX.Element|undefined;
  onSuccess?: () => void;
};

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

const CreateOrEditRoleModal: React.FC<CreateOrEditRoleProps> = (props) => {
  const intl = useIntl();
  const {
    visible,
    onChangeVisible,
    onSuccess,
  } = props;
  const formRef = createRef<FormInstance<API.SystemGetRoleItem>>();

  const [roleInfo, setRoleInfo] = useMergedState<API.SystemGetRoleItem>({}, {
    value: props.oldData,
  });

  const [roleTypeOptions, setRoleTypeOptions] = useState<API.SelectorItem[]>([]);

  useEffect(() => {
    if (roleTypeOptions && roleTypeOptions.length > 0) return;
    System.getRoleTypeSelector().then(data => {
      if (data.result?.list) {
        setRoleTypeOptions(data.result.list);
      }
    })
  });

  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    if (Object.keys(roleInfo).length === 0) formRef.current?.resetFields();
    else formRef.current?.setFieldsValue(roleInfo);
  }, [roleInfo]);

  return (
    <Modal
      title={intl.formatMessage({id: 'pages.action.create'})}
      visible={visible}
      onCancel={() => {onChangeVisible(false)}}
      footer={
        [
          <Button key='cancel' onClick={() => onChangeVisible(false)}>{intl.formatMessage({id: 'pages.confirm.cancel'})}</Button>,
          <Button key='ok' type='primary' onClick={() => {
            setLoading(true);
            System.saveRole({...roleInfo}).then(data => {
              if (data && data.code === 200) {
                if (onSuccess) onSuccess();
                message.success(intl.formatMessage({id: 'pages.result.saveSuccess'})).then(() => {
                  formRef.current?.resetFields();
                });
              }
            }).finally(() => {
              setLoading(false);
            })
          }} loading={loading}>{intl.formatMessage({id: 'pages.confirm.ok'})}</Button>
        ]
      }
    >
      <Form<API.SystemGetRoleItem> {...layout} requiredMark={true} ref={formRef}>
        {
          roleInfo.id &&
          <Form.Item
            label={intl.formatMessage({id: 'pages.system.roleManagement.column.id'})}
            name='id'
            initialValue={roleInfo.id}
          >
            <Input disabled value={roleInfo.id} />
          </Form.Item>
        }

        <Form.Item
          required
          label={intl.formatMessage({id: 'pages.system.roleManagement.column.type'})}
          name='type'
        >
          <Select value={roleInfo.type} onChange={v => {
            setRoleInfo(Object.assign(roleInfo, {type: v}));
          }}
          >
            {
              roleTypeOptions.map(item => <Option value={item.value} key={item.value}>{item.label}</Option>)
            }
          </Select>
        </Form.Item>
        <Form.Item
          required
          label={intl.formatMessage({id: 'pages.system.roleManagement.column.name'})}
          name='name'
        >
          <Input value={roleInfo.name} onChange={({target: {value}}) => {
            setRoleInfo(Object.assign(roleInfo, {name: value}));
          }}/>
        </Form.Item>

        <Form.Item
          label={intl.formatMessage({id: 'pages.system.roleManagement.column.remark'})}
          name='remark'
        >
          <TextArea value={roleInfo.remark} rows={5} onChange={({target: {value}}) => {
            setRoleInfo(Object.assign(roleInfo, {remark: value}));
          }}/>
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreateOrEditRoleModal;
