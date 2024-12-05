package dhtml

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const (
	TAG_NAME_REGEXP       = `^[a-z][a-z0-9\-]{0,255}$`
	ATTRIBUTE_NAME_REGEXP = `^[a-z][a-z0-9\-\:]{0,255}$`
	HTML_ID_REGEXP        = `^[a-z][a-z0-9\_]{0,255}$`
	CLASS_NAME_REGEXP     = `^\-?[\_a-z][\_\-a-z0-9\-\:]{0,255}$`
)

func CheckTagName(name string) error {
	if !regexp.MustCompile(TAG_NAME_REGEXP).MatchString(name) {
		return fmt.Errorf("Wrong tag name: <%s>", name)
	}

	return nil
}

func SafeTagName(name string) string {
	name = strings.ToLower(name)

	if err := CheckTagName(name); err != nil {
		log.Fatalln(err)
	}

	return name
}

func CheckAttributeName(name string) error {
	if !regexp.MustCompile(ATTRIBUTE_NAME_REGEXP).MatchString(name) {
		return fmt.Errorf("Wrong attribute name '%s'", name)
	}

	return nil
}

func SafeAttributeName(name string) string {
	name = strings.ToLower(name)

	if err := CheckAttributeName(name); err != nil {
		log.Fatalln(err)
	}

	return name
}

func CheckClassName(name string) error {
	if !regexp.MustCompile(CLASS_NAME_REGEXP).MatchString(name) {
		return fmt.Errorf("Wrong class name '%s'", name)
	}

	return nil
}

func SafeClassName(name string) string {
	name = strings.ToLower(name)

	if err := CheckClassName(name); err != nil {
		log.Fatalln(err)
	}

	return name
}

func CheckId(s string) error {
	if !regexp.MustCompile(HTML_ID_REGEXP).MatchString(s) {
		return fmt.Errorf("Wrong id: '%s'", s)
	}

	return nil
}

func SafeId(s string) string {
	s = strings.ToLower(s)

	if err := CheckId(s); err != nil {
		log.Fatalln(err)
	}

	return s
}
