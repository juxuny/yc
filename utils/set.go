package utils

type StringSet interface {
	Add(in ...string) StringSet
	Sub(set StringSet) StringSet
	Delete(in ...string) StringSet
	Union(set StringSet) StringSet
	Data() []string
	Clone() StringSet
	Exists(k string) bool
}

type stringSet map[string]bool

func (t stringSet) Data() []string {
	ret := make([]string, 0)
	for k := range t {
		ret = append(ret, k)
	}
	return ret
}

func (t stringSet) Add(in ...string) StringSet {
	for _, item := range in {
		t[item] = true
	}
	return t
}

func (t stringSet) Delete(in ...string) StringSet {
	for _, s := range in {
		delete(t, s)
	}
	return t
}

func (t stringSet) Sub(set StringSet) StringSet {
	if set == nil {
		return t
	}
	return t.Delete(set.Data()...)
}

func (t stringSet) Union(set StringSet) StringSet {
	if set == nil {
		return t
	}
	return t.Add(set.Data()...)
}

func (t stringSet) Clone() StringSet {
	return NewStringSet(t.Data()...)
}

func (t stringSet) Exists(k string) bool {
	return t[k]
}

func NewStringSet(in ...string) StringSet {
	return make(stringSet).Add(in...)
}
