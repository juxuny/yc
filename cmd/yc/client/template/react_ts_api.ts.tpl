import type * as typing from './typing';
import {doRequest} from '@/services/dt';
import { prefix } from './index'

export class cos {
{{range $method := .Methods}}{{if ne $method.Desc ""}}// {{$method.Desc}}{{end}}
  static {{$method.MethodName|lowerFirst}} (body: typing.{{$method.Request}}, options?: { [key: string]: any }) {
    return doRequest{{.Lt}}typing.{{$method.Response}}{{.Gt}}(prefix + '/{{$method.Api}}', {
      method: 'POST',
      data: body,
      ...(options || {}),
    });
  }
{{end}}
}
