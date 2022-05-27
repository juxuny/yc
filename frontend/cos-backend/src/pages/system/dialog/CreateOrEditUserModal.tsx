import { Button, Form, Input, Modal } from 'antd';
import useMergedState from 'rc-util/es/hooks/useMergedState';
import React, { useRef, useEffect, useState } from 'react';
import { useIntl } from '@@/plugin-locale/localeExports';
import { ProForm } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';

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
    { userId: '', identifier: '', nick: '' },
    {
      value: props.oldData,
    },
  );

  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    if (Object.keys(userInfo).length === 0) formRef.current?.resetFields();
    else formRef.current?.setFieldsValue(userInfo);
  }, []);

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
        <Button
          key="ok"
          type="primary"
          onClick={() => {
            setLoading(true);
            console.log(formRef.current?.getFieldsValue());
            setLoading(false);
            if (onSuccess) onSuccess();
          }}
          loading={loading}
        >
          {intl.formatMessage({ id: 'pages.confirm.ok' })}
        </Button>,
      ]}
    >
      <ProForm<API.SaveUserInfoReq> {...layout} requiredMark={true} formRef={formRef} submitter={false}>
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.system.userManagement.column.nick' })}
          name="nick"
        >
          <Input
            value={userInfo.nick}
            onChange={({ target: { value } }) => {
              setUserInfo(Object.assign({ nick: value }, userInfo));
            }}
          />
        </Form.Item>

        <Form.Item
          label={intl.formatMessage({ id: 'pages.system.userManagement.column.nick' })}
          name="nick"
        >
          <Input
            value={userInfo.nick}
            onChange={({ target: { value } }) => {
              setUserInfo(Object.assign({ nick: value }, userInfo));
            }}
          />
        </Form.Item>
      </ProForm>
    </Modal>
  );
};

export default CreateOrEditUserModal;
