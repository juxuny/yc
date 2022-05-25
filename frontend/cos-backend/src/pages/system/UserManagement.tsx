import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { useIntl } from 'umi';

export default (): React.ReactNode => {
  const intl = useIntl();
  return (
    <PageContainer>
      {intl.formatMessage({id: 'pages.welcome.link', defaultMessage: 'welcome'})}
    </PageContainer>
  );
};
