func (t *wrapper) {{.MethodName}}(ctx context.Context, req *{{.PackageAlias}}.{{.Request}}) (resp *{{.PackageAlias}}.{{.Response}}, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
	defer func() {
		if recoverError := recover(); recoverError != nil {
			err = errors.SystemError.InternalError
				debug.PrintStack()
				handleRecover(ctx, recoverError)
			return
		}
	}(){{if .UseAuth}}
	isEnd, err = t.authHandler.Run(ctx)
	if err != nil {
		return nil, err
	}
	if isEnd {
		return nil, nil
	}{{end}}
	isEnd, err = t.beforeHandler.Run(ctx)
	if err != nil {
		return nil, err
	}
	if isEnd {
		return nil, nil
	}
	defer func () {
		_, err := t.afterHandler.Run(ctx)
		if err != nil {
			log.Error(err)
		}
	} ()
	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, err
	}
	defer func () {
		if err != nil {
			log.Error(err)
		}
	} ()
	return t.handler.{{.MethodName}}(ctx, req)
}

