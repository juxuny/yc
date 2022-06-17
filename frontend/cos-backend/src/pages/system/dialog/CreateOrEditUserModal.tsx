import { Button, Form, Input, Modal, message } from 'antd';
import React, { useRef, useEffect, useState } from 'react';
import { useIntl } from '@@/plugin-locale/localeExports';
import { ProForm } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';
import type {SaveOrCreateUserRequest} from "@/services/api/typing";
import {cos} from "@/services/api";

export type CreateOrEditUserProps = {
  visible?: boolean;
  onChangeVisible: (v: boolean) => void;
  oldData?: SaveOrCreateUserRequest,
  onSuccess?: () => void;
};

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

const CreateOrEditUserModal: React.FC<CreateOrEditUserProps> = (props) => {
  const intl = useIntl();
  const { visible, onChangeVisible, onSuccess, oldData } = props;
  const formRef = useRef<ProFormInstance<SaveOrCreateUserRequest> | null>();
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    formRef.current?.setFieldsValue({...oldData});
  }, [oldData]);

  const onSubmit = async () => {
    try {
      setLoading(true);
      const params = formRef.current?.getFieldsValue() || ({} as SaveOrCreateUserRequest);
      params.accountType = props.oldData?.accountType || 1;
      params.userId = props.oldData?.userId;
      const resp = await cos.saveOrCreateUser(params);
      if (resp.code !== 0) {
        message.error(resp.msg);
      } else {
        if (onChangeVisible) onChangeVisible(false);
        if (onSuccess) onSuccess();
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
      <ProForm<SaveOrCreateUserRequest>
        {...layout}
        requiredMark={true}
        formRef={formRef}
        submitter={false}
        initialValues={props.oldData}
      >
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.system.user-management.column.identifier' })}
          name="identifier"
        >
          <Input
            disabled={!!props.oldData?.userId}
          />
        </Form.Item>

        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.system.user-management.column.nick' })}
          name="nick"
        >
          <Input />
        </Form.Item>

        {!props.oldData?.userId && (
          <Form.Item
            required
            label={intl.formatMessage({ id: 'pages.field.password' })}
            name="credential"
          >
            <Input type="password" />
          </Form.Item>
        )}
      </ProForm>
    </Modal>
  );
};

export default CreateOrEditUserModal;
