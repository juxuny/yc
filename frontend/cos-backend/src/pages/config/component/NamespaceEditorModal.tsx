import { Button, Form, Input, Modal, message } from 'antd';
import React, { useRef, useEffect, useState } from 'react';
import { useIntl } from '@@/plugin-locale/localeExports';
import { ProForm } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';
import type {SaveNamespaceRequest} from "@/services/api/typing";
import {cos} from "@/services/api";

export type NamespaceEditorProp = {
  visible?: boolean;
  onChangeVisible: (v: boolean) => void;
  oldData?: SaveNamespaceRequest;
  trigger?: JSX.Element | undefined;
  onSuccess?: () => void;
};

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

const NamespaceEditorModal: React.FC<NamespaceEditorProp> = (props) => {
  const intl = useIntl();
  const { visible, onChangeVisible, onSuccess, oldData } = props;
  const formRef = useRef<ProFormInstance<SaveNamespaceRequest> | undefined>();

  const [loading, setLoading] = useState<boolean>(false);
  useEffect(() => {
    formRef.current?.setFieldsValue({...oldData});
  }, [oldData]);

  const onSubmit = async () => {
    try {
      setLoading(true);
      const params = formRef.current?.getFieldsValue() || ({} as SaveNamespaceRequest);
      const resp = await cos.saveNamespace({
        ...params,
        id: oldData?.id || undefined,
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
        <Button key="ok" type="primary" onClick={onSubmit} loading={loading}>
          {intl.formatMessage({ id: 'pages.confirm.ok' })}
        </Button>,
      ]}
    >
      <ProForm<SaveNamespaceRequest>
        {...layout}
        requiredMark={true}
        formRef={formRef}
        submitter={false}
        initialValues={oldData}
      >
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.config.namespace.column.namespace' })}
          name="namespace"
        >
          <Input />
        </Form.Item>
      </ProForm>
    </Modal>
  );
};

export default NamespaceEditorModal;
