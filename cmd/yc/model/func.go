package model

import (
	"github.com/juxuny/yc/utils"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"strings"
)

func getIndexFromMessageCommentsOfProto(comments []*parser.Comment) (hasIndex bool) {
	for _, c := range comments {
		lines := c.Lines()
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.Index(line, "@index") == 0 {
				return true
			}
		}
	}
	return false
}

func getIgnoreFromMessageCommentsOfProto(comments []*parser.Comment) (hasIgnore bool) {
	for _, c := range comments {
		lines := c.Lines()
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.Index(line, "@ignore-proto") == 0 {
				return true
			}
		}
	}
	return false
}

func findDeletedAtFromMessageOfProto(msg *parser.Message) (found bool) {
	for _, item := range msg.MessageBody {
		v, ok := item.(*parser.Field)
		if !ok {
			continue
		}
		fieldName := utils.ToUnderLine(v.FieldName)
		if fieldName == "deleted_at" {
			return true
		}
	}
	return false
}
