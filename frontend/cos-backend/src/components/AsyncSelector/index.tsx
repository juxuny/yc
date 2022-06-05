import React, { useEffect, useState } from 'react';
import { useIntl } from 'umi';

import { Select } from 'antd';
const { Option } = Select;

export type AsyncSelectorProp = {
  all?: boolean;
  defaultFirst?: boolean;
  defaultValue?: any;
  request?: () => Promise<API.SelectorResp>;
  onChangeValue?: (v: number | string | undefined) => void;
};

export const AsyncSelector: React.FC<AsyncSelectorProp> = (props) => {
  const intl = useIntl();
  const { all, defaultFirst, defaultValue, request, onChangeValue } = props;
  const [items, setItems] = useState<API.SelectorItem[]>([]);
  const [currentValue, setCurrentValue] = useState<any>();

  useEffect(() => {
    console.log('AsyncSelector useEffect');
    const initValue = (list: API.SelectorItem[]) => {
      let finalList = list;
      if (all) {
        finalList = [
          { label: intl.formatMessage({ id: 'pages.label.all' }), value: 'undefined' },
          ...list,
        ];
      }
      if (defaultFirst || defaultValue) {
        setCurrentValue(defaultFirst || defaultValue);
      } else {
        setCurrentValue(finalList[0].value);
        if (onChangeValue) onChangeValue(finalList[0].value);
      }
      if (items.length === 0) setItems(finalList);
    };
    if (!request) return;
    request().then((res) => {
      initValue(res.list);
    });
  });

  return (<Select value={currentValue} onChange={onChangeValue}>
    {items.map((item) => (
      <Option key={item.value} value={item.value}>
        {item.label}
      </Option>
    ))}
  </Select>);
};

export default AsyncSelector;
