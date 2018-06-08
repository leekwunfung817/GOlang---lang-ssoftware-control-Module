package main

/*
 編寫一個簡單的應用程序
 為我們提供最低版本和最高發布版本之間每個版本的最高補丁版本

 用Go編寫
 讀取Github發行版列表，
 使用SemVer進行比較Github發行版列表，
 執行時將文件路徑作為其第一個參數。它讀取這個文件，


written in Go,
reads the Github Releases list,
uses SemVer for comparison
takes a path to a file as its first argument when executed.

It reads this file, which is in the format of:
*/

// Here we implement the basics of communicating with github through the library as well as printing the version
// You will need to implement LatestVersions function as well as make this application support the file format outlined in the README
// Please use the format defined by the fmt.Printf line at the bottom, as we will define a passing coding challenge as one that outputs
// the correct information, including this line
func main() {
	// Github
	// fmt.Println("Github")

	// getVersions_()

	// ReadSourceFile()
	var m = getRepositoryAndMin_version()
	for k, v := range m {
		// fmt.Println(k, v)
		getLibraryAboveVersions(k, v)
	}
}
