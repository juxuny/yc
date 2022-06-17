import { Button, Form, Input, Modal, message, Select } from 'antd';
const { Option } = Select;
import React, { useRef, useEffect, useState } from 'react';
import { useIntl } from '@@/plugin-locale/localeExports';
import { ProForm } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';
import type {SaveValueRequest} from "@/services/api/typing";
import {cos} from "@/services/api";

export type KeyValueEditorProp = {
  visible: boolean;
  onChangeVisible: (v: boolean) => void;
  oldData?: SaveValueRequest;
  onSuccess?: () => void;
};

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

const KeyValueEditorModal: React.FC<KeyValueEditorProp> = (props) => {
  const intl = useIntl();
  const { visible, onChangeVisible, onSuccess, oldData } = props;
  const formRef = useRef<ProFormInstance<SaveValueRequest> | undefined>();

  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    formRef.current?.setFieldsValue({...oldData});
  }, [oldData]);

  const onSubmit = async () => {
    try {
      setLoading(true);
      const params = formRef.current?.getFieldsValue() || ({} as SaveValueRequest);
      const resp = await cos.saveValue({
        ...params,
        configId: oldData?.configId
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
      <ProForm<SaveValueRequest>
        {...layout}
        requiredMark={true}
        formRef={formRef}
        submitter={false}
        initialValues={oldData}
      >
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.config.key-value.column.configKey' })}
          name="configKey"
        >
          <Input disabled={oldData?.configKey !== undefined && oldData?.configKey !== ''} />
        </Form.Item>
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.config.key-value.column.configValue' })}
          name="configValue"
        >
          <Input.TextArea
            autoSize={{ minRows: 3, maxRows: 6 }}
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
