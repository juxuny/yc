
func (t *handler) {{.MethodName}}(ctx context.Context, req *{{.PackageAlias}}.{{.Request}}) (resp *{{.PackageAlias}}.{{.Response}}, err error) {
	resp, err = &{{.PackageAlias}}.{{.Response}}{}, nil
	return
}
