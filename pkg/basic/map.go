package basic

import "fmt"

func Test_map() {
	map1 := make(map[string]string)

	/* map的插入 */
	map1["key1"] = "value1"
	map1["key2"] = "value2"
	map1["key3"] = "value3"
	map1["key4 "] = "value4"

	/* map的遍历 */
	for key := range map1 {
		fmt.Println(key, "value: ", map1[key])
	}

	/* 查找某个数据 */
	value, ok := map1["key1"]
	if ok {
		fmt.Println("key1的值是：", value)
	} else {
		fmt.Println("key1的值不存在")
	}

	/* 删除某个key-value */
	delete(map1, "key1")

	value, ok = map1["key1"]
	if ok {
		fmt.Println("key1的值是：", value)
	} else {
		fmt.Println("key1的值不存在")
	}
}
