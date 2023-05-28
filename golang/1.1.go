package main

//模块一作业1.1
//给定一个字符串数组
//[“I”,“am”,“stupid”,“and”,“weak”]
//用 for 循环遍历该数组并修改为
//[“I”,“am”,“smart”,“and”,“strong”]
//并打印修改后的数组
func main() {
	var arr = [5]string{"I", "am", "stupid", "and", "weak"}

	// define a map
	convert := map[string]string{
		"stupid": "smart",
		"weak":   "strong",
	}

	for i := 0; i < len(arr); i++ {
		if val, ok := convert[arr[i]]; ok {
			arr[i] = val
		}

		println(arr[i])
	}
}
