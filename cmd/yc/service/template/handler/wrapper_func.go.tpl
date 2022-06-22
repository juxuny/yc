func (t *wrapper) {{.MethodName}}(ctx context.Context, req *{{.PackageAlias}}.{{.Request}}) (resp *{{.PackageAlias}}.{{.Response}}, err error) {
	if err := t.runMiddle(ctx, {{if .UseAuth}}true{{else}}false{{end}}, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.{{.MethodName}}(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

