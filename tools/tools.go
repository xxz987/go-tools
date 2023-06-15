package gotools

import "encoding/json"

// MergeTwoJson 合并json
// eg: params s1:{"a":{"b":1,"d":4}} s2:{"a":{"b":2,"c":3},"b":{"c":"cc"}}
// return {"a":{"b":2,"c":3,"d":4},"b":{"c":"cc"}}
func MergeTwoJson(s1, s2 string) (res string, err error) {
	m1 := map[string]interface{}{}
	m2 := map[string]interface{}{}
	res = ""
	err = nil
	err = json.Unmarshal([]byte(s1), &m1)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(s2), &m2)
	if err != nil {
		return
	}
	for k, v := range m2 {
		// 检查是否存在指定的键
		if v_m1, ok := m1[k]; ok {
			vFloat, isFloat := v.(float64)
			if isFloat {
				m1[k] = vFloat
				continue
			}
			vStr, isStr := v.(string)
			if isStr {
				m1[k] = vStr
				continue
			}
			vFloat, isFloat = v_m1.(float64)
			if isFloat {
				m1[k] = vFloat
				continue
			}
			vStr, isStr = v_m1.(string)
			if isStr {
				m1[k] = vStr
				continue
			}
			var str1, str2 []byte
			str1, err = json.Marshal(m1[k])
			if err != nil {
				return
			}
			str2, err = json.Marshal(v)
			if err != nil {
				return
			}
			m1[k], err = MergeTwoJsonByTree(string(str1), string(str2))
			if err != nil {
				return
			}
		} else {
			m1[k] = v
		}
	}
	marshal, err := json.Marshal(m1)
	if err != nil {
		return
	}
	res = string(marshal)
	return
}
