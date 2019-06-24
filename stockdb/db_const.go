package dbapi


import "fmt"

const DB_NAME ="stockauto"
const COLLECT_NAME_TEST ="xuangu"

type ErrorDB struct{
	When string
	What string
}

func (err ErrorDB) Error()string{
	ret := fmt.Sprintf("%s %s", err.When, err.What )
	return ret
}

 