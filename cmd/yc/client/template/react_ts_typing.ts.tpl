import type * as dt from '@juxuny/yc-ts-data-type/typing';

{{range $enum := .Enums}}export type {{$enum.EnumName}} = {{$enum.ValueSet}};

{{if ne $enum.Desc "" }}// {{$enum.Desc}}
{{end}}export class {{$enum.EnumName}}Enum {
{{range $field := $enum.Fields}}{{if ne $field.Desc ""}}  // {{$field.Desc}}{{end}}
  static {{$field.FieldName|upperFirst}}: {{$enum.EnumName}} = {{$field.Value}};
{{end}}
}

{{end}}{{range $msg := .Messages}}{{if ne $msg.Desc ""}}
// {{$msg.Desc}}{{end}}
export type {{$msg.Name}} = {
{{range $field := $msg.Fields}}{{if ne $field.Desc ""}}  // {{$field.Desc}}{{end}}
  {{$field.Name|lowerFirst}}{{if not $field.Required}}?{{end}}: {{$field.Type}};
{{end}}
}
{{end}}
