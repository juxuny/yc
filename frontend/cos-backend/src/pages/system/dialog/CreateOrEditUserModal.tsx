import { Button, Form, Input, Modal, message } from 'antd';
import useMergedState from 'rc-util/es/hooks/useMergedState';
import React, { useRef, useEffect, useState } from 'react';
import { useIntl } from '@@/plugin-locale/localeExports';
import { ProForm } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';
import { User } from '@/services/cos/user';

export type CreateOrEditUserProps = {
  visible?: boolean;
  onChangeVisible: (v: boolean) => void;
  oldData?: API.SaveUserInfoReq;
  trigger?: JSX.Element | undefined;
  onSuccess?: () => void;
};

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

const CreateOrEditUserModal: React.FC<CreateOrEditUserProps> = (props) => {
  const intl = useIntl();
  const { visible, onChangeVisible, onSuccess } = props;
  const formRef = useRef<ProFormInstance<API.SaveUserInfoReq> | null>();

  const [userInfo, setUserInfo] = useMergedState<API.SaveUserInfoReq>(
    { userId: '', identifier: '', nick: '', credential: '' },
    {
      value: props.oldData,
    },
  );

  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    formRef.current?.setFieldsValue(userInfo);
  });

  const onSubmit = async () => {
    try {
      setLoading(true);
      const params = formRef.current?.getFieldsValue() || ({} as API.SaveUserInfoReq);
      params.accountType = props.oldData?.accountType || 1;
      params.userId = props.oldData?.userId;
      const resp = await User.saveOrCreateUser(params);
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
      <ProForm<API.SaveUserInfoReq>
        {...layout}
        requiredMark={true}
        formRef={formRef}
        submitter={false}
      >
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.system.user-management.column.identifier' })}
          name="identifier"
        >
          <Input
            disabled={!!userInfo.userId}
            value={userInfo.identifier}
            onChange={({ target: { value } }) => {
              setUserInfo(Object.assign(userInfo, { identifier: value }));
            }}
          />
        </Form.Item>

        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.system.user-management.column.nick' })}
          name="nick"
        >
          <Input
            value={userInfo.nick}
            onChange={({ target: { value } }) => {
              setUserInfo(Object.assign(userInfo, { nick: value }));
            }}
          />
        </Form.Item>

        {!userInfo.userId && (
          <Form.Item
            required
            label={intl.formatMessage({ id: 'pages.field.password' })}
            name="credential"
          >
            <Input
              type="password"
              value={userInfo.credential}
              onChange={({ target: { value } }) => {
                setUserInfo(Object.assign(userInfo, { credential: value }));
              }}
            />
          </Form.Item>
        )}
      </ProForm>
    </Modal>
  );
};

export default CreateOrEditUserModal;
