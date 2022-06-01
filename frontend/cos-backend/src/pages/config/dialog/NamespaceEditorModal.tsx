import { Button, Form, Input, Modal, message } from 'antd';
import useMergedState from 'rc-util/es/hooks/useMergedState';
import React, { useRef, useEffect, useState } from 'react';
import { useIntl } from '@@/plugin-locale/localeExports';
import { ProForm } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';
import { Namespace } from '@/services/cos/namespace';

export type NamespaceEditorProp = {
  visible?: boolean;
  onChangeVisible: (v: boolean) => void;
  oldData?: API.Namespace.SaveReq;
  trigger?: JSX.Element | undefined;
  onSuccess?: () => void;
};

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

const NamespaceEditorModal: React.FC<NamespaceEditorProp> = (props) => {
  const intl = useIntl();
  const { visible, onChangeVisible, onSuccess } = props;
  const formRef = useRef<ProFormInstance<API.Namespace.SaveReq>|undefined>();

  const [namespaceData, setNamespaceData] = useMergedState<API.Namespace.SaveReq>(
    { } as API.Namespace.SaveReq,
    {
      value: props.oldData,
    },
  );

  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    formRef.current?.setFieldsValue(namespaceData);
  });

  const onSubmit = async () => {
    try {
      setLoading(true);
      const params = formRef.current?.getFieldsValue() || ({} as API.Namespace.SaveReq);
      const resp = await Namespace.save({
        ...params,
        id: props.oldData?.id || undefined
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
      <ProForm<API.Namespace.SaveReq>
        {...layout}
        requiredMark={true}
        formRef={formRef}
        submitter={false}
      >
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.config.namespace.column.namespace' })}
          name="namespace"
        >
          <Input
            value={namespaceData.namespace}
            onChange={({ target: { value } }) => {
              setNamespaceData(Object.assign(namespaceData, { namespace: value }));
            }}
          />
        </Form.Item>
      </ProForm>
    </Modal>
  );
};

export default NamespaceEditorModal;
