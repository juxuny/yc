func (t *wrapper) {{.MethodName}}(ctx context.Context, req *{{.PackageAlias}}.{{.Request}}) (resp *{{.PackageAlias}}.{{.Response}}, err error) {
	trace.WithContext(ctx)
	defer trace.Clean()
	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, err
	}
	return t.handler.{{.MethodName}}(ctx, req)
}

