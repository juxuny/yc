import { Button, Modal } from 'antd';
import React  from 'react';
import {FormattedMessage, useIntl} from '@@/plugin-locale/localeExports';
import { Typography } from 'antd';
const { Text, Paragraph } = Typography;

import type {CreateResult} from "@/pages/account/component/CreateAccessKeyModal";


export type CreateResultProps = {
  visible?: boolean;
  onChangeVisible: (v: boolean) => void;
  data?: CreateResult;
};

const CreateResultModal: React.FC<CreateResultProps> = (props) => {
  const intl = useIntl();
  const { visible, onChangeVisible, data} = props;

  return (
    <Modal
      width={'680px'}
      title={intl.formatMessage({ id: 'pages.action.create' })}
      visible={visible}
      onCancel={() => {
        onChangeVisible(false);
      }}
      footer={[
        <Button key="cancel" onClick={() => onChangeVisible(false)}>
          {intl.formatMessage({ id: 'pages.confirm.cancel' })}
        </Button>
      ]}
    >
      <Text type='danger'>
        <FormattedMessage id={'pages.account.access-key.create-result.warning'}/>
      </Text>
      <br/>
      <Text strong>AccessKey: </Text>
      <br/>
      <Paragraph copyable>{data?.accessKey}</Paragraph>
      <br/>
      <Text strong>Secret: </Text>
      <br/>
      <Paragraph copyable>{data?.secret}</Paragraph>
    </Modal>
  );
};

export default CreateResultModal;
