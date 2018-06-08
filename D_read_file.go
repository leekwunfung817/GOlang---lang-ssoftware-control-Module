package main

func ReadSourceFile() [][]string {

	var set = ReadSetting("F_Setting.txt")
	// for k, v := range set {
	// 	fmt.Println(k, v)
	// }
	var source = ReadCSV(set["source_file"])
	return source

	// ReadCSV(set["result_file"])

}

func getRepositoryAndMin_version() map[string]string {
	var source = ReadSourceFile()
	var result = make(map[string]string, len(source)-1)
	// var condition_expression []string
	for index := 0; index < len(source); index++ {
		if index == 0 {
			// condition_expression = source[index]
			for cell := 0; cell < len(source[index]); cell++ {
				// fmt.Println("Condition expression: ", source[index][cell])
			}
		} else {
			result[source[index][0]] = source[index][1]
			for cell := 0; cell < len(source[index]); cell++ {
				// fmt.Println("Data: ", source[index][cell], "(", condition_expression[cell], ")")
			}
		}
	}
	return result
}
