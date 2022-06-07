import { Button, Form, Input, Modal, message, Select } from 'antd';
const { Option } = Select;
import useMergedState from 'rc-util/es/hooks/useMergedState';
import React, { useRef, useEffect, useState } from 'react';
import { useIntl } from '@@/plugin-locale/localeExports';
import { ProForm } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';
import { KeyValue } from '@/services/cos/key_value';

export type KeyValueEditorProp = {
  visible: boolean;
  onChangeVisible: (v: boolean) => void;
  oldData?: API.KeyValue.SaveReq;
  onSuccess?: () => void;
};

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

const KeyValueEditorModal: React.FC<KeyValueEditorProp> = (props) => {
  const intl = useIntl();
  const { visible, onChangeVisible, onSuccess, oldData } = props;
  const formRef = useRef<ProFormInstance<API.KeyValue.SaveReq> | undefined>();

  const [editingData, setEditingData] = useMergedState<API.KeyValue.SaveReq>(
    {} as API.KeyValue.SaveReq,
    {
      value: oldData,
    },
  );

  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    formRef.current?.setFieldsValue(editingData);
  });

  const onSubmit = async () => {
    try {
      setLoading(true);
      const params = formRef.current?.getFieldsValue() || ({} as API.KeyValue.SaveReq);
      const resp = await KeyValue.save({
        ...params,
        id: editingData.id,
        configId: editingData.configId
      });
      if (resp.code !== 0) {
        message.error(resp.msg);
      } else {
        if (onChangeVisible) onChangeVisible(false);
        if (onSuccess) onSuccess();
        formRef.current?.resetFields();
      }
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal
      title={intl.formatMessage({ id: 'pages.action.create' })}
      visible={visible}
      onCancel={() => onChangeVisible(false)}
      footer={[
        <Button key="cancel" onClick={() => onChangeVisible(false)}>
          {intl.formatMessage({ id: 'pages.confirm.cancel' })}
        </Button>,
        <Button key="ok" type="primary" onClick={onSubmit} loading={loading}>
          {intl.formatMessage({ id: 'pages.confirm.ok' })}
        </Button>,
      ]}
    >
      <ProForm<API.KeyValue.SaveReq>
        {...layout}
        requiredMark={true}
        formRef={formRef}
        submitter={false}
      >
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.config.key-value.column.configKey' })}
          name="configKey"
        >
          <Input
            disabled={editingData.id !== undefined}
            value={editingData.configKey}
            onChange={({ target: { value }}) => {
              setEditingData(Object.assign(editingData, { configKey: value }));
            }}
          />
        </Form.Item>
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.config.key-value.column.configValue' })}
          name="configValue"
        >
          <Input.TextArea
            autoSize={{ minRows: 3, maxRows: 6 }}
            value={editingData.configValue}
            onChange={({ target: { value } }) => {
              setEditingData(Object.assign(editingData, { configValue: value }));
            }}
          />
        </Form.Item>
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.config.key-value.column.isHot' })}
          name="isHot"
        >
          <Select>
            <Option key={'hot'} value={true}>{intl.formatMessage({ id: 'pages.config.key-value.column.isHot' })}</Option>
            <Option key={'cold'} value={false}>{intl.formatMessage({ id: 'pages.config.key-value.column.cold' })}</Option>
          </Select>
        </Form.Item>
      </ProForm>
    </Modal>
  );
};

export default KeyValueEditorModal;
