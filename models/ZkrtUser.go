package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"github.com/astaxie/beego/orm"
)

type ZkrtUser struct {
	Id               int    `orm:"column(Zkrt_id);auto"`
	ZkrtUser         string `orm:"column(Zkrt_user);size(45)"`
	ZkrtPassword     string `orm:"column(Zkrt_password);size(45)"`
	ZkrtFastName     string `orm:"column(Zkrt_fastName);size(45)"`
	ZkrtLastName     string `orm:"column(Zkrt_lastName);size(45)"`
	ZkrtJurisdiction int    `orm:"column(Zkrt_Jurisdiction)"`
	ZkrtTime         string `orm:"column(Zkrt_time);size(45)"`
}

func (t *ZkrtUser) TableName() string {
	return "ZkrtUser"
}

func init() {
	orm.RegisterModel(new(ZkrtUser))
}

/*func FindUserByToken(token string) (bool, ZkrtUser) {
	o := orm.NewOrm()
	var user ZkrtUser
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}*/


// AddZkrtUser insert a new ZkrtUser into database and returns
// last inserted Id on success.
func AddZkrtUser(m *ZkrtUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetZkrtUserById retrieves ZkrtUser by Id. Returns error if
// Id doesn't exist
func GetZkrtUserById(id int) (v *ZkrtUser, err error) {
	o := orm.NewOrm()
	v = &ZkrtUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//登陆获取基本的信息
func GetWaiterForLogin(username string, password string) (bool, ZkrtUser) {
	o := orm.NewOrm()
	var user ZkrtUser
	err := o.QueryTable(user).Filter("Zkrt_user", username).Filter("Zkrt_password", password).One(&user)
	return err != orm.ErrNoRows, user
}

// GetAllZkrtUser retrieves all ZkrtUser matches certain condition. Returns empty list if
// no records exist
func GetAllZkrtUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ZkrtUser))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
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

	var l []ZkrtUser
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

// UpdateZkrtUser updates ZkrtUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateZkrtUserById(m *ZkrtUser) (err error) {
	o := orm.NewOrm()
	v := ZkrtUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteZkrtUser deletes ZkrtUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteZkrtUser(id int) (err error) {
	o := orm.NewOrm()
	v := ZkrtUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ZkrtUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func RegisterDataBase() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/zkrt_db?charset=utf8")
}
