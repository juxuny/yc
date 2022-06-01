import React, { useEffect, useState } from 'react';
import { useIntl } from 'umi';

import { Select } from 'antd';
const { Option } = Select;

export type AsyncSelectorProp = {
  all?: boolean;
  defaultFirst?: boolean;
  defaultValue?: any;
  request?: () => Promise<API.SelectorResp>;
};

export const AsyncSelector: React.FC<AsyncSelectorProp> = (props) => {
  const intl = useIntl();
  const { all, defaultFirst, defaultValue, request } = props;
  const [items, setItems] = useState<API.SelectorItem[]>([]);
  const [currentValue, setCurrentValue] = useState<any>();

  useEffect(() => {
    const initValue = (list: API.SelectorItem[]) => {
      let finalList = list;
      if (all)
        finalList = [
          { label: intl.formatMessage({ id: 'pages.label.all' }), value: 'undefined' },
          ...list,
        ];
      if (defaultFirst || defaultValue) {
        setCurrentValue(defaultFirst || defaultValue);
      }
      if (all && !defaultFirst && !defaultValue && list.length > 0) {
        setCurrentValue(finalList[0].value);
      }
      if (items.length === 0) setItems(finalList);
    };
    if (!request) return;
    request().then((res) => {
      initValue(res.list);
    });
  });
  if (items.length === 0) {
    return <div />;
  } else {
    return (
      <Select defaultValue={currentValue}>
        {items.map((item) => (
          <Option key={item.value} value={item.value}>
            {item.label}
          </Option>
        ))}
      </Select>
    );
  }
};

export default AsyncSelector;
