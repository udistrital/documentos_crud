package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type FirmaElectronica struct {
	Id                 string `orm:"column(id);pk"`
	CodigoAutenticidad string `orm:"column(codigo_autenticidad)"`
	Firmantes          string `orm:"column(firmantes);type(json)"`
	FirmaEncriptada    string `orm:"column(firma_encriptada)"`
	Llaves             string `orm:"column(llaves);type(json)"`
	Activo             bool   `orm:"column(activo)"`
	FechaCreacion      string `orm:"column(fecha_creacion)"`
	FechaModificacion  string `orm:"column(fecha_modificacion)"`
}

func (t *FirmaElectronica) TableName() string {
	return "firma_electronica"
}

func init() {
	orm.RegisterModel(new(FirmaElectronica))
}

// AddFirmaElectronica insert a new FirmaElectronica into database and returns
// last inserted Id on success.
func AddFirmaElectronica(m *FirmaElectronica) (id string, err error) {
	o := orm.NewOrm()

	const script = `
	INSERT INTO documento.firma_electronica (
		codigo_autenticidad,
		llaves,
		firma_encriptada,
		firmantes,
		activo,
		fecha_creacion,
		fecha_modificacion)
	VALUES(?, ?, ?, ?, ?, now(), now())
	RETURNING id;`

	err = o.Raw(script, m.CodigoAutenticidad, m.Llaves, m.FirmaEncriptada, m.Firmantes, m.Activo).QueryRow(&m.Id)

	return
}

// GetFirmaElectronicaById retrieves FirmaElectronica by Id. Returns error if
// Id doesn't exist
func GetFirmaElectronicaById(id string) (v *FirmaElectronica, err error) {
	o := orm.NewOrm()
	v = &FirmaElectronica{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFirmaElectronica retrieves all FirmaElectronica matches certain condition. Returns empty list if
// no records exist
func GetAllFirmaElectronica(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FirmaElectronica)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.HasSuffix(k, "__in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []FirmaElectronica
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateFirmaElectronica updates FirmaElectronica by Id and returns error if
// the record to be updated doesn't exist
func UpdateFirmaElectronicaById(m *FirmaElectronica) (err error) {
	o := orm.NewOrm()
	v := FirmaElectronica{Id: m.Id}
	// ascertain id exists in the database
	err = o.Read(&v)
	if err != nil {
		return
	}

	const script = `
	UPDATE documento.firma_electronica
	SET codigo_autenticidad = ?,
		llaves = ?,
		firmantes = ?,
		firma_encriptada = ?,
		activo = ?,
		fecha_modificacion = now()
	WHERE id = ?;`

	_, err = o.Raw(script, m.CodigoAutenticidad, m.Llaves, m.Firmantes, m.FirmaEncriptada, m.Activo, m.Id).Exec()

	return
}

// DeleteFirmaElectronica deletes FirmaElectronica by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFirmaElectronica(id string) (err error) {
	o := orm.NewOrm()
	v := FirmaElectronica{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FirmaElectronica{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
