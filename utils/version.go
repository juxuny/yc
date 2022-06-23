package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Stage string

const (
	StageAlpha = Stage("alpha")
	StageBeta  = Stage("beta")
	StageGamma = Stage("gamma")
	StageRc    = Stage("rc")
	StageFinal = Stage("final")
)

type Version struct {
	LastError error
	Valid     bool
	Data      []int
	Stage     Stage
}

func ParseVersion(version string) Version {
	ret := Version{
		Data:  make([]int, 3),
		Stage: StageAlpha,
	}
	version = strings.ToLower(version)
	if strings.Contains(version, string(StageAlpha)) {
		ret.Stage = StageAlpha
	} else if strings.Contains(version, string(StageBeta)) {
		ret.Stage = StageBeta
	} else if strings.Contains(version, string(StageGamma)) {
		ret.Stage = StageGamma
	} else if strings.Contains(version, string(StageRc)) {
		ret.Stage = StageRc
	} else if strings.Contains(version, string(StageFinal)) {
		ret.Stage = StageFinal
	} else {
		ret.Valid = false
		return ret
	}
	version = StringHelper.DropLetters(version)
	version = strings.Trim(version, "- ")
	l := strings.Split(version, ".")
	for _, item := range l {
		v, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			ret.LastError = err
			return ret
		}
		ret.Data = append(ret.Data, int(v))
	}
	for len(ret.Data) < 3 {
		ret.Data = append(ret.Data, 0)
	}
	ret.Valid = true
	return ret
}

func (t Version) String() string {
	data := make([]string, 0)
	for _, item := range t.Data {
		data = append(data, fmt.Sprintf("%d", item))
	}
	vs := strings.Join(data, ".")
	ret := "v" + vs
	if t.Stage != "" {
		ret += "-" + string(t.Stage)
	}
	return ret
}

func (t Version) IsEqual(v Version) bool {
	if !t.Valid || !v.Valid {
		return false
	}
	if t.Stage != v.Stage {
		return false
	}

	minLength := len(t.Data)
	if len(v.Data) < minLength {
		minLength = len(v.Data)
	}
	for i := 0; i < minLength; i++ {
		if t.Data[i] != v.Data[i] {
			return false
		}
	}
	if len(t.Data) > minLength {
		for i := minLength; i < len(t.Data); i++ {
			if t.Data[i] != 0 {
				return false
			}
		}
	}
	if len(v.Data) > minLength {
		for i := minLength; i < len(v.Data); i++ {
			if v.Data[i] != 0 {
				return false
			}
		}
	}
	return true
}

func (t Version) Less(v Version) bool {
	if !t.Valid || !v.Valid {
		return false
	}
	if t.Stage != v.Stage {
		return false
	}
	minLength := len(t.Data)
	if len(v.Data) < minLength {
		minLength = len(v.Data)
	}
	for i := 0; i < minLength; i++ {
		if t.Data[i] < v.Data[i] {
			return true
		}
	}
	if len(v.Data) > minLength {
		for i := minLength; i < len(t.Data); i++ {
			if v.Data[i] > 0 {
				return true
			}
		}
	}
	return false
}

func (t Version) Greater(v Version) bool {
	if !t.Valid || !v.Valid {
		return false
	}
	if t.Stage != v.Stage {
		return false
	}
	minLength := len(t.Data)
	if len(v.Data) < minLength {
		minLength = len(v.Data)
	}
	for i := 0; i < minLength; i++ {
		if t.Data[i] > v.Data[i] {
			return true
		}
	}
	if len(t.Data) > minLength {
		for i := minLength; i < len(t.Data); i++ {
			if t.Data[i] > 0 {
				return true
			}
		}
	}
	return false
}

func (t Version) GreaterOrEqual(v Version) bool {
	return t.Greater(v) || t.IsEqual(v)
}

func (t Version) LessOrEqual(v Version) bool {
	return t.Less(v) || t.IsEqual(v)
}

var VersionHelper = versionHelper{}

type versionHelper struct{}

func (versionHelper) IsEqual(v1, v2 string) bool {
	return ParseVersion(v1).IsEqual(ParseVersion(v2))
}

func (versionHelper) Greater(v1, v2 string) bool {
	return ParseVersion(v1).Greater(ParseVersion(v2))
}

func (versionHelper) Less(v1, v2 string) bool {
	return ParseVersion(v1).Less(ParseVersion(v2))
}

func (versionHelper) GreaterOrEqual(v1, v2 string) bool {
	return ParseVersion(v1).GreaterOrEqual(ParseVersion(v2))
}

func (versionHelper) LessOrEqual(v1, v2 string) bool {
	return ParseVersion(v1).LessOrEqual(ParseVersion(v2))
}
