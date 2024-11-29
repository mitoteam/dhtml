package dhtml

import (
	"errors"
	"log"
	"regexp"
)

const (
	TAG_NAME_REGEXP       = `^(?i)[a-z][a-z0-9\-]{0,255}$`
	ATTRIBUTE_NAME_REGEXP = `^(?i)[a-z][a-z0-9\-\:]{0,255}$`
	CLASS_NAME_REGEXP     = `^(?i)\-?[\_a-z][\_\-a-z0-9\-\:]{0,255}$`
)

func CheckTagName(name string) error {
	if !regexp.MustCompile(TAG_NAME_REGEXP).MatchString(name) {
		return errors.New("Wrong tag name: " + name)
	}

	return nil
}

func SafeTagName(name string) string {
	if err := CheckTagName(name); err != nil {
		log.Fatalln(err)
	}

	return name
}

func CheckAttributeName(name string) error {
	if !regexp.MustCompile(ATTRIBUTE_NAME_REGEXP).MatchString(name) {
		return errors.New("Wrong attribute name: " + name)
	}

	return nil
}

func SafeAttributeName(name string) string {
	if err := CheckAttributeName(name); err != nil {
		log.Fatalln(err)
	}

	return name
}

func CheckClassName(name string) error {
	if !regexp.MustCompile(CLASS_NAME_REGEXP).MatchString(name) {
		return errors.New("Wrong class name: " + name)
	}

	return nil
}

func SafeClassName(name string) string {
	if err := CheckClassName(name); err != nil {
		log.Fatalln(err)
	}

	return name
}
