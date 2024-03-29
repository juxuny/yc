import React, { useRef, useState } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { useIntl } from 'umi';
import { Button, message } from 'antd';
import { FormattedMessage } from '@@/plugin-locale/localeExports';
import { ProForm, ProCard, ProFormText } from '@ant-design/pro-components';
import type { ProFormInstance } from '@ant-design/pro-components';
import type { ModifyPasswordRequest } from '@/services/api/typing';
import {cos} from "@/services/api";

export default (): React.ReactNode => {
  const intl = useIntl();
  const [loading, setLoading] = useState<number>(0);
  const formRef = useRef<ProFormInstance<ModifyPasswordRequest> | undefined>(undefined);

  const savePassword = async (update: ModifyPasswordRequest) => {
    try {
      setLoading((v) => v + 1);
      await formRef.current?.validateFields();
      const resp = await cos.modifyPassword(update);
      if (resp && resp.code === 0) {
        message.success(intl.formatMessage({ id: 'pages.result.modifySuccess' }));
        formRef.current?.resetFields();
      }
    } catch (err) {
      console.error(err);
    } finally {
      setLoading((v) => v - 1);
    }
  };

  return (
    <PageContainer>
      <ProCard>
        <ProForm<ModifyPasswordRequest>
          formRef={formRef}
          submitter={{
            render: () => {
              return [
                <Button
                  key={'ok'}
                  type={'primary'}
                  loading={loading > 0}
                  onClick={async () =>
                    await savePassword(
                      formRef.current?.getFieldsValue() || ({} as ModifyPasswordRequest),
                    )
                  }
                >
                  <FormattedMessage id={'pages.action.ok'} />
                </Button>,
                <Button key={'reset'} onClick={() => formRef.current?.resetFields()}>
                  <FormattedMessage id={'pages.action.reset'} />
                </Button>,
              ];
            },
          }}
        >
          <ProForm.Group>
            <ProFormText.Password
              rules={[
                { required: true },
                {
                  pattern: /^([a-zA-Z0-9]{6,22})$/,
                  message: intl.formatMessage({ id: 'pages.account.modify-password.tips' }),
                },
              ]}
              name="oldPassword"
              label={intl.formatMessage({ id: 'pages.account.modify-password.label.old' })}
              tooltip={intl.formatMessage({ id: 'pages.account.modify-password.tips' })}
              placeholder={intl.formatMessage({ id: 'pages.account.modify-password.placeholder' })}
            />
          </ProForm.Group>
          <ProForm.Group>
            <ProFormText.Password
              rules={[
                { required: true },
                {
                  pattern: /^([a-zA-Z0-9]{6,22})$/,
                  message: intl.formatMessage({ id: 'pages.account.modify-password.tips' }),
                },
              ]}
              name="newPassword"
              label={intl.formatMessage({ id: 'pages.account.modify-password.label.new' })}
              tooltip={intl.formatMessage({ id: 'pages.account.modify-password.tips' })}
              placeholder={intl.formatMessage({ id: 'pages.account.modify-password.placeholder' })}
            />
          </ProForm.Group>
        </ProForm>
      </ProCard>
    </PageContainer>
  );
};
