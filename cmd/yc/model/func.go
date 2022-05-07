package model

import (
	"github.com/juxuny/yc/utils"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"strings"
)

func getIndexFromMessageCommentsOfProto(comments []*parser.Comment) (hasIndex bool) {
	for _, c := range comments {
		if c == nil {
			continue
		}
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
		if c == nil {
			continue
		}
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
		if item == nil {
			continue
		}
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

func getRefFromMessageOfProto(msg *parser.Message) (ref string, found bool) {
	for _, c := range msg.Comments {
		if c == nil {
			continue
		}
		for _, line := range c.Lines() {
			if strings.Contains(line, "@ref:") {
				ref = strings.Replace(line, "@ref:", "", 1)
				ref = strings.TrimSpace(ref)
				return ref, true
			}
		}
	}
	return "", false
}

func getOrmAliasFromMessageOfProto(field *parser.Field) (columnAlias string, found bool) {
	for _, c := range field.Comments {
		if c == nil {
			continue
		}
		for _, line := range c.Lines() {
			if strings.Contains(line, "@orm:") {
				columnAlias = strings.TrimSpace(strings.Replace(line, "@orm:", "", 1))
				return columnAlias, true
			}
		}
	}
	return "", false
}
