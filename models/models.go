package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Conf struct {
	Id    int
	Lowlimit int
	Uplimit int
}

type Data struct {
	Id    int
	Date  time.Time
	Value int
}

func init() {
	orm.RegisterModel(new(Conf), new(Data))
}
