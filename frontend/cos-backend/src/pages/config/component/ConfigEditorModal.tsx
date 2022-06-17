import { Button, Form, Input, Modal, message } from 'antd';
import React, { useRef, useEffect, useState } from 'react';
import { useIntl } from '@@/plugin-locale/localeExports';
import { ProForm } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';
import type {SaveConfigRequest} from "@/services/api/typing";
import {cos} from "@/services/api";

export type ConfigEditorProp = {
  visible?: boolean;
  onChangeVisible: (v: boolean) => void;
  oldData?: SaveConfigRequest,
  onSuccess?: () => void;
  isClone: boolean;
};

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

const ConfigEditorModal: React.FC<ConfigEditorProp> = (props) => {
  const intl = useIntl();
  const { visible, onChangeVisible, onSuccess, isClone, oldData } = props;
  const formRef = useRef<ProFormInstance<SaveConfigRequest> | undefined>();

  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    formRef.current?.setFieldsValue({...oldData});
  }, [oldData]);

  const onSubmitSave = async () => {
    try {
      setLoading(true);
      const params = formRef.current?.getFieldsValue() || {};
      const resp = await cos.saveConfig({
        ...params,
        baseId: props.oldData?.baseId,
        namespaceId: props.oldData?.namespaceId,
        id: props.oldData?.id || undefined,
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

  const onSubmitClone = async () => {
    try {
      setLoading(true);
      const params = formRef.current?.getFieldsValue() || {};
      const resp = await cos.cloneConfig({
        id: oldData?.id || 0,
        newConfigId: params.configId
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
      onCancel={() => {
        onChangeVisible(false);
      }}
      footer={[
        <Button key="cancel" onClick={() => onChangeVisible(false)}>
          {intl.formatMessage({ id: 'pages.confirm.cancel' })}
        </Button>,
        <Button key="ok" type="primary" onClick={isClone ? onSubmitClone : onSubmitSave} loading={loading}>
          {intl.formatMessage({ id: 'pages.confirm.ok' })}
        </Button>,
      ]}
    >
      <ProForm<SaveConfigRequest>
        {...layout}
        requiredMark={true}
        formRef={formRef}
        submitter={false}
        initialValues={props.oldData}
      >
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.config.config-management.column.configId' })}
          name="configId"
        >
          <Input />
        </Form.Item>
      </ProForm>
    </Modal>
  );
};

export default ConfigEditorModal;
