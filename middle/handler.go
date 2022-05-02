package middle

import "context"

type Handler interface {
	Run(context context.Context) (isEnd bool, err error)
}

type Group interface {
	Run(context context.Context) (isEnd bool, err error)
	Add(h ...Handler) Group
}

type group struct {
	handlers []Handler
}

func (t *group) Run(context context.Context) (isEnd bool, err error) {
	for _, h := range t.handlers {
		isEnd, err := h.Run(context)
		if err != nil {
			return true, err
		}
		if isEnd {
			return isEnd, nil
		}
	}
	return false, nil
}

func (t *group) Add(h ...Handler) Group {
	t.handlers = append(t.handlers, h...)
	return t
}

func NewGroup(h ...Handler) Group {
	ret := &group{
		handlers: h,
	}
	return ret
}
