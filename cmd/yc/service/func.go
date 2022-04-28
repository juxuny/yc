package service

import (
	"github.com/juxuny/yc/utils"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"strings"
)

func getGroupNameFromRpcCommentsOfProto(comments []*parser.Comment) (groupName string, ok bool) {
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
