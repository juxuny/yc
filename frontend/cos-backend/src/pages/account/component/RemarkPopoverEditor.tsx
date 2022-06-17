import React, {useEffect, useState} from "react";
import {Button, Form, Input, Popover} from "antd";
import { useIntl } from 'umi';
import useMergedState from 'rc-util/es/hooks/useMergedState';
import {EditOutlined} from "@ant-design/icons";
import {FormattedMessage} from "@@/plugin-locale/localeExports";
import {useForm} from "antd/es/form/Form";
import type {SetAccessKeyRemarkRequest} from "@/services/api/typing";
import {cos} from "@/services/api";

export type RemarkPopoverEditorProp = {
  showPopup: boolean;
  data: SetAccessKeyRemarkRequest
  onSuccess: (remark: string) => void;
  onChangeVisible: (accessKeyId: number | string | undefined) => void;
}

export const RemarkPopoverEditor: React.FC<RemarkPopoverEditorProp> = (props) => {
  const {data, showPopup, onChangeVisible, onSuccess} = props;
  const intl = useIntl();
  const [formRef] = useForm<API.User.SetRemarkAccessKeyReq>();
  const [loading, setLoading] = useState<number>(0);
  const [formValue, setFormValue] = useMergedState<SetAccessKeyRemarkRequest>({} as SetAccessKeyRemarkRequest, {
    value: data
  });

  useEffect(() => {
    formRef.setFieldsValue(formValue);
  });

  const onSubmit = async () => {
    try {
      setLoading(v => v + 1);
      const resp = await cos.setRemarkAccessKey({
        id: data.id,
        remark: formValue.remark
      });
      if (resp && resp.code === 0) {
        onChangeVisible(undefined);
        onSuccess(formValue.remark || '');
      }
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(v => v - 1);
    }
  }

  return (<>
    {data.remark} <Popover visible={showPopup} trigger="click" title={intl.formatMessage({ id: 'pages.account.access-key.title.setRemark' })} content={() => {
    return (<Form<API.User.SetRemarkAccessKeyReq> layout={'inline'} form={formRef}>
      <Form.Item name='remark'>
        <Input key={'remark'} value={formValue.remark} onChange={({target: {value}}) => {
          setFormValue(Object.assign(formValue, {remark: value}));
        }}/>
      </Form.Item>
      <Form.Item>
        <Button type={'default'} onClick={() => onChangeVisible(undefined)}>
          <FormattedMessage id={'pages.action.cancel'}/>
        </Button>
      </Form.Item>
      <Form.Item>
        <Button type={'primary'} onClick={onSubmit} loading={loading > 0}>
          <FormattedMessage id={'pages.action.ok'}/>
        </Button>
      </Form.Item>
    </Form>);
  }}>
    <a onClick={() => onChangeVisible(data.id)}><EditOutlined/></a>
  </Popover>
  </>);
}
export default RemarkPopoverEditor;
