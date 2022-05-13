package main

import (
	"strings"
	"unicode"
)

type Package struct {
}

func (pack Package) apply_name_translations(upstream_name string) string {
	/**
	Applies some name hacks and corrections for certain components.

	:param upstream_name: the name as coming from the package data
	:return: the name after corrections applied to, if any
	*/
	var result_name = upstream_name
	var splitted_name = strings.Split(result_name, "_")
	if len(splitted_name) > 1 && splitted_name[len(splitted_name)-1] != "" && unicode.IsDigit((rune(splitted_name[len(splitted_name)-1][0]))) {
		result_name = strings.Join(splitted_name[:], "_")
	}

	for _, namehack := range [...]string{"openjdk-", "jdk-", "java-", "zulu-"} {
		if strings.HasPrefix(result_name, namehack) {
			part := strings.Split(result_name, namehack)
			if len(part) > 1 && unicode.IsDigit(rune(part[1][0])) {
				result_name = namehack[0:]
				break
			}
		}
	}
	return result_name
}

func (pack Package) apply_version_translations(libname, version string) string {

	if libname == "sapjvm" {
		version = ""
	}
	return version
}

func (pack Package) get_package_representative_file(package_contents, package_contents_fileobjs) {

}
