package util

import (
	"design-api/model"
	"encoding/json"
	"reflect"
	"sort"
)

/**
struct转map
*/
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

//申明一个排序类型
type sortByBelong []models.Category

func (s sortByBelong) Len() int {
	return len(s)
}

func (s sortByBelong) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByBelong) Less(i, j int) bool {
	return s[i].Belong < s[j].Belong
}

/**
切片或数组按照指定字段分组排序，必须实现sort的三个方法，这其实是一个冒泡排序
*/
func SplitSlice(list interface{}) interface{} {

	switch list.(type) {
	case []models.Category:
		v := list.([]models.Category)
		sort.Sort(sortByBelong(v))

		//returnData := make([][]models.Category, 0)
		returnData := make(map[string][]models.Category) //生成一个带key的map
		i := 0
		var j int
		for {
			if i >= len(v) {
				break
			}

			for j = i + 1; j < len(v) && v[i].Belong == v[j].Belong; j++ {
			}

			//returnData = append(returnData, v[i:j])
			returnData[v[i].Belong] = v[i:j]

			i = j
		}

		return returnData
	default:
		return nil
	}
}

/**
json转map
 */
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		return nil, err
	}

	return mapResult, nil
}
