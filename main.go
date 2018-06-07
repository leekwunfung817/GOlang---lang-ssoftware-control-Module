package main

import (
	"context"
	"fmt"

	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
)

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

// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	var versionSlice []*semver.Version
	// This is just an example structure of the code, if you implement this interface, the test cases in main_test.go are very easy to run
	return versionSlice
}

// func ()  {

// }

func getVersions(version string, lenght int) []string {
	client := github.NewClient(nil)
	ctx := context.Background()
	// https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=10
	// https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=1000000:
	// opt := &github.ListOptions{PerPage: 1}
	opt := &github.ListOptions{PerPage: lenght}
	releases, _, err := client.Repositories.ListReleases(ctx, "kubernetes", "kubernetes", opt)
	if err != nil {
		panic(err) // is this really a good way?
	}
	// minVersion := semver.New("1.8.0")
	minVersion := semver.New(version)
	allReleases := make([]*semver.Version, len(releases))
	// fmt.Println(len(releases))
	var count = 0

	var version_array = make([]string, len(releases))
	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}
		// strings.Split(versionString, "-")[0]
		version_array[i] = versionString
		allReleases[i] = semver.New(versionString)
		// fmt.Println(count, versionString, len(versionString))
		fmt.Println(versionString)
		count++
	}
	// fmt.Println(version_array)
	versionSlice := LatestVersions(allReleases, minVersion)
	fmt.Println("latest versions of kubernetes/kubernetes:", versionSlice, "Lenght: ", count)
	// fmt.Println(reflect.TypeOf(version_array))

	return version_array
}
func getVersions_() []string {
	return getVersions("1.8.0", 1000000)
}

// Here we implement the basics of communicating with github through the library as well as printing the version
// You will need to implement LatestVersions function as well as make this application support the file format outlined in the README
// Please use the format defined by the fmt.Printf line at the bottom, as we will define a passing coding challenge as one that outputs
// the correct information, including this line
func main() {
	// Github
	fmt.Println("Github")
	arr := getVersions_()
	fmt.Println(len(arr))

	// arr_1 := getVersions("1.8.0", 1000000)
	// arr_2 := getVersions("1.10.4", 1000000)
	// arr_2 := getVersions("1.13.4", 1000000)

	// fmt.Println(reflect.DeepEqual(arr_1, arr_2))

}
