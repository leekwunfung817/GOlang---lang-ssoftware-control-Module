package main

import "strings"

func version_comparer_string_preprocess(versionString string) string {
	var comparer = versionString + ""
	comparer = strings.Replace(comparer, "-rc.", "", -1)
	comparer = strings.Replace(comparer, "rc.", "", -1)
	comparer = strings.Replace(comparer, "-rc", ".", -1)
	comparer = strings.Replace(comparer, "rc", ".", -1)

	comparer = strings.Replace(comparer, "-beta.", "", -1)
	comparer = strings.Replace(comparer, "beta.", "", -1)
	comparer = strings.Replace(comparer, "-beta", ".", -1)
	comparer = strings.Replace(comparer, "beta", ".", -1)

	comparer = strings.Replace(comparer, "..", ".", -1)
	comparer = strings.Replace(comparer, "..", ".", -1)
	comparer = strings.Replace(comparer, ".0.", ".", -1)
	return comparer
}

func high_level_version_preprocess(version string) string {
	version = high_level_version_string_preprocess(version)
	version = strings.Split(version, "-")[0]
	// version = strings.Split(version, ".")
	var index = 0
	for true {
		if !(index < len(version)) {
			break
		}
		var subresult = version[len(version)-index:]
		index++
		if strings.Contains(subresult, ".0") {
			version = version[:len(version)-index]
			index = 0
		}
	}
	return version
}

func high_level_version_string_preprocess(versionString string) string {
	var comparer = versionString + ""
	comparer = strings.Replace(comparer, "-rc.", "rc.", -1)
	comparer = strings.Replace(comparer, "rc.", "-rc.", -1)

	comparer = strings.Replace(comparer, "-beta.", "beta.", -1)
	comparer = strings.Replace(comparer, "beta", "-beta", -1)
	return comparer
}
