package services

import (
	"github.com/juxuny/yc/utils"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"strings"
)

func getLinesContainTag(comments []*parser.Comment, tag string) (lines []string) {
	for _, c := range comments {
		lines := c.Lines()
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.Index(line, tag) == 0 {
				lines = append(lines, strings.Trim(strings.Replace(line, tag, "", 1), ": "))
			}
		}
	}
	return
}

func GetGroupNameFromRpcCommentsOfProto(comments []*parser.Comment) (groupName string, ok bool) {
	for _, c := range comments {
		lines := c.Lines()
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.Index(line, "@group") == 0 {
				groupName = strings.Replace(line, "@group:", "", 1)
				groupName = utils.ToUnderLine(strings.TrimSpace(groupName))
				return groupName, true
			}
		}
	}
	return
}

func GetAuthFromRpcCommentsOfProto(comments []*parser.Comment) (auth bool) {
	for _, c := range comments {
		lines := c.Lines()
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.Index(line, "@auth") == 0 {
				return true
			}

			if strings.Index(line, "@ignore-auth") == 0 {
				return false
			}
		}
	}
	return true
}

func GetDescFromFieldCommentsOfProto(comments []*parser.Comment) string {
	for _, c := range comments {
		lines := c.Lines()
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.Index(line, "@desc:") == 0 {
				return strings.TrimSpace(strings.Replace(line, "@desc:", "", 1))
			}
		}
	}
	return ""
}

func CheckIfRequired(comments []*parser.Comment) bool {
	lines := getLinesContainTag(comments, "@required")
	return len(lines) > 0
}

func GetContentByProtoTag(protoTag ProtoTag, comments []*parser.Comment) []string {
	ret := make([]string, 0)
	for _, c := range comments {
		lines := c.Lines()
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, protoTag.String()) {
				ret = append(ret, strings.TrimLeft(utils.StringHelper.TrimSubSequenceLeft(line, protoTag.String()), ": "))
			}
		}
	}
	return ret
}

func GetContentByProtoTagFistOne(protoTag ProtoTag, comments []*parser.Comment) string {
	for _, c := range comments {
		lines := c.Lines()
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, protoTag.String()) {
				return strings.Trim(utils.StringHelper.TrimSubSequenceLeft(line, protoTag.String()), ": ")
			}
		}
	}
	return ""
}

func CheckIfContainProtoTag(protoTag ProtoTag, comments []*parser.Comment) bool {
	ret := GetContentByProtoTag(protoTag, comments)
	return len(ret) > 0
}
