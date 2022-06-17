package dt

import (
	"fmt"
	"github.com/juxuny/yc/errors"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
)

func (m *ID) Clone() *ID {
	return proto.Clone(m).(*ID)
}

func (m *NullInt64) Clone() *NullInt64 {
	return proto.Clone(m).(*NullInt64)
}

func (m *NullInt32) Clone() *NullInt32 {
	return proto.Clone(m).(*NullInt32)
}

func (m *NullBool) Clone() *NullBool {
	return proto.Clone(m).(*NullBool)
}

func (m *NullFloat32) Clone() *NullFloat32 {
	return proto.Clone(m).(*NullFloat32)
}

func (m *NullFloat64) Clone() *NullFloat64 {
	return proto.Clone(m).(*NullFloat64)
}

func (m *NullString) Clone() *NullString {
	return proto.Clone(m).(*NullString)
}

func (m *Pagination) Clone() *Pagination {
	return proto.Clone(m).(*Pagination)
}

func (m *PaginationResp) Clone() *PaginationResp {
	return proto.Clone(m).(*PaginationResp)
}

func (m *Error) Clone() *Error {
	return proto.Clone(m).(*Error)
}

func (m *ID) NewPointer() *ID {
	return NewIDPointer(m.Uint64)
}

func (m *ID) NumberAsString() string {
	return fmt.Sprintf("%d", m.Uint64)
}

func (m *ID) UnmarshalJSON(data []byte) error {
	input := strings.Trim(string(data), "\" ")
	if input == "" || strings.Contains(input, "null") {
		m.Valid = false
		return nil
	}
	if strings.ContainsAny(input, "-. \r\n") {
		return fmt.Errorf("invalid ID: %v", input)
	}
	id, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return err
	}
	m.Valid = true
	m.Uint64 = id
	return err
}

func (m ID) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", m.Uint64)), nil
}

func (m *ID) Equal(id *ID) bool {
	return m.Valid == id.Valid && m.Uint64 == id.Uint64
}

func (m *NullInt64) UnmarshalJSON(data []byte) error {
	input := strings.Trim(string(data), "\" ")
	if input == "" || strings.Contains(input, "null") {
		m.Valid = false
		return nil
	}
	id, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return err
	}
	m.Valid = true
	m.Int64 = id
	return err
}

func (m *NullInt64) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", m.Int64)), nil
}

func (m *NullInt64) Equal(id *NullInt64) bool {
	return m.Valid == id.Valid && m.Int64 == id.Int64
}

func (m *NullInt32) UnmarshalJSON(data []byte) error {
	input := strings.Trim(string(data), "\" ")
	if input == "" || strings.Contains(input, "null") {
		m.Valid = false
		return nil
	}
	id, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return err
	}
	m.Valid = true
	m.Int32 = int32(id)
	return err
}

func (m *NullInt32) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", m.Int32)), nil
}

func (m *NullInt32) Equal(id *NullInt32) bool {
	return m.Valid == id.Valid && m.Int32 == id.Int32
}

func (m *NullFloat32) UnmarshalJSON(data []byte) error {
	input := strings.Trim(string(data), "\" ")
	if input == "" || strings.Contains(input, "null") {
		m.Valid = false
		return nil
	}
	value, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return err
	}
	m.Valid = true
	m.Float32 = float32(value)
	return err
}

func (m *NullFloat32) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", m.Float32)), nil
}

func (m *NullFloat32) Equal(input *NullFloat32) bool {
	return m.Valid == input.Valid && m.Float32 == input.Float32
}

func (m *NullFloat64) UnmarshalJSON(data []byte) error {
	input := strings.Trim(string(data), "\" ")
	if input == "" || strings.Contains(input, "null") {
		m.Valid = false
		return nil
	}
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return err
	}
	m.Valid = true
	m.Float64 = value
	return err
}

func (m *NullFloat64) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", m.Float64)), nil
}

func (m *NullFloat64) Equal(input *NullFloat64) bool {
	return m.Valid == input.Valid && m.Float64 == input.Float64
}

func (m *NullBool) UnmarshalJSON(data []byte) error {
	input := strings.Trim(string(data), "\" ")
	if input == "" || strings.Contains(input, "null") {
		m.Valid = false
		return nil
	}
	value, err := strconv.ParseBool(input)
	if err != nil {
		return err
	}
	m.Valid = true
	m.Bool = value
	return err
}

func (m *NullBool) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%v", m.Bool)), nil
}

func (m *NullBool) Equal(input *NullBool) bool {
	return m.Valid == input.Valid && m.Bool == input.Bool
}

func (m *Pagination) ToResp(total int64) PaginationResp {
	return PaginationResp{
		PageNum:  m.PageNum,
		PageSize: m.PageSize,
		Total:    total,
	}
}

func (m *Pagination) ToRespPointer(total int64) *PaginationResp {
	return &PaginationResp{
		PageNum:  m.PageNum,
		PageSize: m.PageSize,
		Total:    total,
	}
}

func (m *Error) Error() errors.Error {
	data := make(map[string]interface{})
	for k, v := range m.Data {
		data[k] = v
	}
	return errors.Error{
		Code: m.Code,
		Msg:  m.Msg,
		Data: data,
	}
}

func FromError(err errors.Error) *Error {
	data := make(map[string]string)
	for k, v := range err.Data {
		data[k] = fmt.Sprintf("%v", v)
	}
	return &Error{
		Code: err.Code,
		Msg:  err.Msg,
		Data: data,
	}
}
