package params

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type validationRule string

const (
	ruleEmail validationRule = "email"
	ruleUUID  validationRule = "uuid"
)

var allRules = []validationRule{ruleEmail, ruleUUID}

var patterns = map[validationRule]*regexp.Regexp{
	ruleEmail: regexp.MustCompile(`^[^\s]+@[^\s]+$`),
	ruleUUID:  regexp.MustCompile(`^[0-9a-fA-F]{8}-?[0-9a-fA-F]{4}-?[0-9a-fA-F]{4}-?[0-9a-fA-F]{4}-?[0-9a-fA-F]{12}$`),
}

type field struct {
	v     reflect.Value
	rules []validationRule
}

func parseValidateField(field string) ([]validationRule, error) {
	rules := []validationRule{}
	for _, s := range strings.Split(field, ",") {
		matched := false
		for _, rule := range allRules {
			if string(rule) == s {
				rules = append(rules, rule)
				matched = true
				break
			}
		}
		if !matched {
			return nil, fmt.Errorf("unsupported validation rule: %s", s)
		}
	}
	return rules, nil
}

func validate(rules []validationRule, name, value string) error {
	for _, rule := range rules {
		if matched := patterns[rule].MatchString(value); !matched {
			return fmt.Errorf("query %q must be %s, got: %q", name, rule, value)
		}
	}
	return nil
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]*field)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		rules, err := parseValidateField(tag.Get("validate"))
		if err != nil {
			return err
		}
		fields[name] = &field{v.Field(i), rules}
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.v.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if err := validate(f.rules, name, value); err != nil {
				return err
			}
			if f.v.Kind() == reflect.Slice {
				elem := reflect.New(f.v.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.v.Set(reflect.Append(f.v, elem))
			} else {
				if err := populate(f.v, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
