
func (t *handler) {{.MethodName}}(ctx context.Context, req *{{.PackageAlias}}.{{.Request}}) (resp *{{.PackageAlias}}.{{.Response}}, err error) {
	return &{{.PackageAlias}}.{{.Response}}{}, nil
}
