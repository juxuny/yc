import { Button, Form, Input, Modal } from 'antd';
import useMergedState from 'rc-util/es/hooks/useMergedState';
import React, { useRef, useEffect, useState } from 'react';
import { useIntl } from '@@/plugin-locale/localeExports';
import { ProForm } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';
import { User } from '@/services/cos/user';

export type CreateAccessKeyProps = {
  visible?: boolean;
  onChangeVisible: (v: boolean) => void;
  data?: API.User.CreateAccessKeyReq;
  onSuccess?: () => void;
};

const layout = {
  labelCol: { span: 6 },
  wrapperCol: { span: 16 },
};

const CreateAccessKeyModal: React.FC<CreateAccessKeyProps> = (props) => {
  const intl = useIntl();
  const { visible, onChangeVisible, onSuccess, data} = props;
  const formRef = useRef<ProFormInstance<API.User.CreateAccessKeyReq> | null>();

  const [currentData, setCurrentData] = useMergedState<API.User.CreateAccessKeyReq>(
    {} as API.User.CreateAccessKeyReq,
    {
      value: data,
    },
  );

  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    formRef.current?.setFieldsValue(currentData);
  });

  const onSubmit = async () => {
    try {
      setLoading(true);
      const params = formRef.current?.getFieldsValue() || ({} as API.User.CreateAccessKeyReq);
      const resp = await User.createAccessKey(params);
      if (resp && resp.code === 0) {
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
      <ProForm<API.User.CreateAccessKeyReq>
        {...layout}
        requiredMark={true}
        formRef={formRef}
        submitter={false}
      >
        <Form.Item
          required
          label={intl.formatMessage({ id: 'pages.account.access-key.column.remark' })}
          name="remark"
        >
          <Input
            value={currentData.remark}
            onChange={({ target: { value } }) => {
              setCurrentData(Object.assign(currentData, { remark: value }));
            }}
          />
        </Form.Item>
      </ProForm>
    </Modal>
  );
};

export default CreateAccessKeyModal;
