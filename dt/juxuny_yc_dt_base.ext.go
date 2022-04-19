package dt

import (
	"fmt"
	"strconv"
	"strings"
)

func (m *ID) UnmarshalJSON(data []byte) error {
	input := strings.Trim(string(data), "\" ")
	if input == "" || strings.Contains(input, "null") {
		m.Valid = false
		return nil
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

func (m ID) Equal(id ID) bool {
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

func (m NullInt64) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", m.Int64)), nil
}

func (m NullInt64) Equal(id NullInt64) bool {
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

func (m NullInt32) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", m.Int32)), nil
}

func (m NullInt32) Equal(id NullInt32) bool {
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

func (m NullFloat32) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", m.Float32)), nil
}

func (m NullFloat32) Equal(input NullFloat32) bool {
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

func (m NullFloat64) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", m.Float64)), nil
}

func (m NullFloat64) Equal(input NullFloat64) bool {
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

func (m NullBool) MarshalJSON() (data []byte, err error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%v", m.Bool)), nil
}

func (m NullBool) Equal(input NullBool) bool {
	return m.Valid == input.Valid && m.Bool == input.Bool
}
