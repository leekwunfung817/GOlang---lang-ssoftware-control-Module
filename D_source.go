package main

import (
	"os"
	"strings"
	// Data Library
	"context"
	"fmt"
	"sort"

	// SemVer Library
	Masterminds_semver "github.com/Masterminds/semver"
	"github.com/coreos/go-semver/semver"
	// Github Library
	"github.com/google/go-github/github"
)

// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	var versionSlice []*semver.Version
	// This is just an example structure of the code, if you implement this interface, the test cases in main_test.go are very easy to run
	return versionSlice
}

func getLibraryAboveVersions(library_name string, min_version string) []string {
	// SemVer Condition
	c, err := Masterminds_semver.NewConstraint(">= " + min_version)
	if err != nil {
		// Handle constraint not being parseable.
	}
	client := github.NewClient(nil)
	ctx := context.Background()
	// https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=10
	// https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=1000000:
	opt := &github.ListOptions{PerPage: 1000000}
	var vd = strings.Split(library_name, "/")
	releases, _, err := client.Repositories.ListReleases(ctx, vd[0], vd[1], opt)
	if err != nil {
		// panic(err)
		// is this really a good way?
		// Answer: No, User must know the Internet is desconnected!
		// We should hide the URL to user
		fmt.Println("Please Check your Internet is Connected!")
		fmt.Println("If your Internet is Connected, Please Check GitHub server is available!")
		os.Exit(0)
	}
	var vs []string
	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}
		var comparer = version_comparer_string_preprocess(versionString)
		// SemVer Add New Version
		v, err := Masterminds_semver.NewVersion(comparer)
		if err != nil {
			// Handle version not being parseable.
			fmt.Errorf("Error parsing version: %s", err, i)
			fmt.Println("ERR: ", comparer, versionString, err)
			continue
		} else {
			// Check if the version meets the constraints. The a variable will be true.
			a := c.Check(v)
			if a {
				vs = append(vs, versionString)
			}
		}
	}

	// SemVer Array Sorting
	sort.StringsAreSorted(vs)
	fmt.Println(vs)
	var HighestPatchVersions = VersionList_HighestPatchVersion_PostProcess(vs)
	fmt.Println(HighestPatchVersions)
	fmt.Println("latest versions of "+library_name+":", HighestPatchVersions)
	return vs
}

func getVersions(version string, lenght int) []*Masterminds_semver.Version {
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
	// minVersion := semver.New(version)
	allReleases := make([]*semver.Version, len(releases))
	// fmt.Println(len(releases))

	var count = 0
	// var m map[string]int
	// var m = make(map[string]int)
	var version_array = make([]string, len(releases))

	// SemVer Object (Version) Array Initialize
	vs := make([]*Masterminds_semver.Version, len(releases))

	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}

		// SemVer Add New Version
		v, err := Masterminds_semver.NewVersion(versionString)
		if err != nil {
			fmt.Errorf("Error parsing version: %s", err)
		}
		vs[i] = v

		// strings.Split(versionString, "-")[0]
		version_array[i] = versionString
		allReleases[i] = semver.New(versionString)
		// m[versionString] = i
		// fmt.Println(count, versionString, len(versionString))
		// fmt.Println(versionString)
		count++
	}

	// SemVer Array Sorting
	sort.Sort(Masterminds_semver.Collection(vs))
	// semver

	sort.Strings(version_array)
	// fmt.Println(version_array)
	// versionSlice := LatestVersions(allReleases, minVersion)
	// fmt.Println("latest versions of kubernetes/kubernetes:", versionSlice, "Lenght: ", count)
	fmt.Println("latest versions of kubernetes/kubernetes:", vs)
	// fmt.Println(reflect.TypeOf(version_array))

	return vs
	// return version_array
	// return m
}

func getVersions_() []*Masterminds_semver.Version {
	return getVersions("1.8.0", 1000000)
}

// []*Masterminds_semver.Version
