package orm

import (
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/utils"
	"reflect"
)

type Column struct {
	reflect.Value
	Name string
}

type Row []Column

func (t Row) Reform(out reflect.Value) error {
	if out.Kind() != reflect.Ptr {
		return errors.SystemError.NotPointer
	}
	if len(t) == 0 {
		return errors.SystemError.DatabaseNoData
	}
	outElem := out.Elem()
	if outElem.Kind() == reflect.Struct {
		nm := getOrmColumnNameMap(outElem.Type())
		for _, col := range t {
			fieldName, b := nm[col.Name]
			if !b {
				fieldName = utils.ToHump(col.Name)
			}
			ft := outElem.FieldByName(fieldName)
			if !ft.IsValid() {
				continue
			}
			ft.Set(defaultConverter.Convert(col.Value, ft.Type()))
		}
	} else {
		outElem.Set(defaultConverter.Convert(t[0].Value, outElem.Type()))
	}
	return nil
}

type DataSet []Row

func NewDataSet() DataSet {
	return make(DataSet, 0)
}

func (t *DataSet) Reform(out interface{}) error {
	outType := reflect.TypeOf(out)
	if outType.Kind() != reflect.Ptr {
		return errors.SystemError.NotPointer
	}
	outValue := reflect.ValueOf(out)
	list := outValue.Elem()
	//fmt.Println("type of list: ", list.Type().String())
	if list.Kind() == reflect.Slice {
		elemType := list.Type().Elem()
		//fmt.Println("type of elem: ", elemType.String())
		for _, row := range *t {
			var item reflect.Value
			if elemType.Kind() == reflect.Ptr {
				item = reflect.Zero(elemType)
			} else {
				item = reflect.New(elemType)
			}
			if err := row.Reform(item); err != nil {
				return err
			}
			if elemType.Kind() == reflect.Ptr {
				list.Set(reflect.Append(list, item))
			} else {
				list.Set(reflect.Append(list, item.Elem()))
			}
		}
	} else if list.Kind() == reflect.Struct {
		if len(*t) == 0 {
			return errors.SystemError.DatabaseNoData
		}
		if err := (*t)[0].Reform(outValue); err != nil {
			return err
		}
	} else {
		if len(*t) == 0 {
			return errors.SystemError.DatabaseNoData
		}
		if err := (*t)[0].Reform(outValue); err != nil {
			return err
		}
	}
	return nil
}
