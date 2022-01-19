// Package data 数据层
package data

//GetData 返回数据层 TODO
func GetData() (interface{}, error) {
	return []interface{}{Empty{}}, nil
}

type Empty struct{}
