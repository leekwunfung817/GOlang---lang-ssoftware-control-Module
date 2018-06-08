package main

import (
	"sort"

	"github.com/Masterminds/semver"
)

func VersionList_HighestPatchVersion_PostProcess(list []string) []string {
	var high_level_version = filter_out_high_level_version(list)
	// fmt.Println(high_level_version)
	return compare_HighestPatchVersion(high_level_version, list)
}

func compare_HighestPatchVersion(high_level_version_list []string, list []string) []string {
	var HighestPatchVersion []string

	for index := 0; index < len(high_level_version_list); index++ {
		c, err := semver.NewConstraint("~" + high_level_version_list[index])
		if err != nil {
			// Handle constraint not being parseable.
		} else {
			var sub_version []string
			for i := 0; i < len(list); i++ {
				v, _ := semver.NewVersion(list[i])
				if err != nil {
					// Handle version not being parseable.
				}
				// Validate a version against a constraint.
				a, msgs := c.Validate(v)
				if msgs == nil {
				}
				if a {
					sub_version = append(sub_version, list[i])
				}
			}
			sort.StringsAreSorted(sub_version)
			if len(sub_version) > 0 {
				if !list_contains(HighestPatchVersion, sub_version[0]) {
					HighestPatchVersion = append(HighestPatchVersion, sub_version[0])
				}

				// fmt.Println(sub_version[0])
			}

		}
	}
	sort.StringsAreSorted(HighestPatchVersion)
	return HighestPatchVersion
}

func filter_out_high_level_version(list []string) []string {
	var high_level_version []string
	// find all high_level_version
	for version := range list {
		var v = high_level_version_preprocess(list[version])
		if !list_contains(high_level_version, v) {
			high_level_version = append(high_level_version, v)
		}
	}
	return high_level_version
}
