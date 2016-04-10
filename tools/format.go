package tools

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ToMap(o interface{}) map[string]interface{} {
	switch o.(type) {
	case map[string]interface{}:
		return o.(map[string]interface{})
	default:
		return nil
	}
}

func ToArray(o interface{}) []interface{} {
	switch o.(type) {
	case []interface{}:
		return o.([]interface{})
	default:
		return nil
	}
}

func ToBool(o interface{}) (bool, error) {
	switch b := o.(type) {
	case string:
		if b == "true" || b == "1" {
			return true, nil
		} else if b == "false" || b == "0" {
			return false, nil
		}
	case int, int32, int64:
		if b == 0 {
			return false, nil
		} else if b == 1 {
			return true, nil
		}
	case interface{}:
		return o.(bool), nil
	default:

	}

	return false, errors.New("Not a bool type")
}

func ToString(o interface{}) string {

	switch o.(type) {
	case string:
		return o.(string)
	default:
		v := fmt.Sprint(o)
		if v == "<nil>" {
			return ""
		} else if v == "nil" {
			return ""
		} else {
			return v
		}
	}
}

func ToInt(o interface{}) int {
	switch o.(type) {
	case string:
		v, err := strconv.Atoi(o.(string))
		if err != nil {
			return 0
		} else {
			return v
		}
	case int, int32:
		return o.(int)
	case int64:
		return int(o.(int64))
	case float64:
		return int(o.(float64))
	default:
		return 0
	}
}

func ToInt64(o interface{}) int64 {
	switch o.(type) {
	case string:
		v, err := strconv.ParseInt(o.(string), 10, 64)
		if err != nil {
			return 0
		} else {
			return v
		}
	case int, int32:
		return int64(o.(int))
	case int64:
		return o.(int64)
	case float64:
		return int64(o.(float64))
	default:
		return 0
	}
}

func ToFloat64(o interface{}) float64 {
	switch o.(type) {
	case string:
		v, err := strconv.ParseFloat(o.(string), 64)
		if err != nil {
			return 0
		} else {
			return v
		}
	case int, int32:
		return float64(o.(int))
	case int64:
		return float64(o.(int64))
	case float64:
		return o.(float64)
	default:
		return 0
	}
}

func Format_size(size int64) string {
	//size = 1073741824
	if size >= 1099511627776 {
		if size%1099511627776 == 0 {
			return strconv.FormatInt(size/1099511627776, 10) + " T"
		} else {
			return strconv.FormatFloat(float64(size)/float64(1099511627776), 'f', 2, 64) + " T"
		}

	} else if size >= 1073741824 {
		if size%1073741824 == 0 {
			return strconv.FormatInt(size/1073741824, 10) + " G"
		} else {
			return strconv.FormatFloat(float64(size)/float64(1073741824), 'f', 2, 64) + " G"
		}

	} else if size >= 1048576 {
		if size%1048576 == 0 {
			return strconv.FormatInt(size/1048576, 10) + " M"
		} else {
			return strconv.FormatFloat(float64(size)/float64(1048576), 'f', 2, 64) + " M"
		}

	} else if size >= 1024 {
		if size%1024 == 0 {
			return strconv.FormatInt(size/1024, 10) + " K"
		} else {
			return strconv.FormatFloat(float64(size)/float64(1024), 'f', 2, 64) + " K"
		}

	} else {
		return strconv.FormatInt(size, 10) + " B"
	}

}

func Int64ToDateString(inttime interface{}, format string) string {
	toint := ToInt64(inttime)
	t := time.Unix(toint, 0)
	//"2006-01-02 15:04"
	if format == "" {
		format = "2006-01-02 15:04"
	}
	return t.Format(format)
}

//简单的解释INI文件的方法
func ParseIniFile(filename string) (map[string]interface{}, error) {
	//实例化这个map
	tmpMap := make(map[string]interface{})
	//是否在段中
	section := ""
	//打开这个ini文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//读取文件到buffer里边
	buf := bufio.NewReader(file)
	for {
		//按照换行读取每一行
		l, err := buf.ReadString('\n')
		//判断退出循环
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(2)
			}

		}
		line := strings.TrimSpace(l)
		if len(line) == 0 {
			continue
		}

		if strings.Index(line, "#") == 0 {
			continue
		}

		if strings.Contains(line, "=") {
			slice := strings.Split(line, "=")
			tmpMap[section+strings.TrimSpace(slice[0])] = strings.TrimSpace(slice[1])
		} else if strings.Contains(line, "[") && strings.Contains(line, "]") {
			posl := strings.Index(line, "[")
			posr := strings.Index(line, "]")
			section_name := strings.TrimSpace(line[posl+1 : posr])
			if strings.ToLower(section_name) == "end" {
				section = ""
			} else {
				section = section_name + ":"
			}
		}

	}

	return tmpMap, nil
}

func ParseJsonFile(filename string) (map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	c := make(map[string]interface{})
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func JoinInt64Slice(str string, s []int64) string {

	sl := ""
	length := len(s)
	for i := 0; i < length; i++ {
		sl += ToString(s[i])
		if i < length-1 {
			sl += str
		}
	}

	return sl
}

func JoinIntSlice(str string, s []int) string {

	sl := ""
	length := len(s)
	for i := 0; i < length; i++ {
		sl += ToString(s[i])
		if i < length-1 {
			sl += str
		}
	}

	return sl
}

//Map 转 json
func MapToJson(data map[string]interface{}) (str string) {
	b, err := json.Marshal(data)
	if err == nil {
		str = string(b)
		return
	}

	return ""
}

//字符串数组转成int64数组
func Str2Int64Slice(s []string) []int64 {
	i := make([]int64, len(s))
	for _, v := range s {
		i = append(i, ToInt64(strings.TrimSpace(v)))
	}
	return i
}

//struct 转json
func Struct2Json(v interface{}) (s string, err error) {
	body, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	s = string(body)
	return s, nil
}

//Json字符串转成Map
func Json2Map(str string) (map[string]interface{}, error) {
	r := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &r)
	if err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

//Json字符串转成Struct
func Json2Struct(str string, v interface{}) error {
	err := json.Unmarshal([]byte(str), &v)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func Struct2Map(obj interface{}) map[string]interface{} {

	if obj == nil {
		return nil
	}
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return nil
	}
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}
