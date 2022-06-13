// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package flex

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-openapi/strfmt"
)

func NormalizeJSONString(jsonString interface{}) (string, error) {
	var j interface{}
	if jsonString == nil || jsonString.(string) == "" {
		return "", nil
	}
	s := jsonString.(string)
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return s, err
	}
	bytes, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(bytes[:]), nil
}

func SepIdParts(id string, separator string) ([]string, error) {
	if strings.Contains(id, separator) {
		parts := strings.Split(id, separator)
		return parts, nil
	}
	return []string{}, fmt.Errorf("The given id %s does not contain %s please check documentation on how to provider id during import command", id, separator)
}

func ptrToInt(i int) *int {
	return &i
}

func PtrToString(s string) *string {
	return &s
}

func IntValue(i64 *int64) (i int) {
	if i64 != nil {
		i = int(*i64)
	}
	return
}

func float64Value(f32 *float32) (f float64) {
	if f32 != nil {
		f = float64(*f32)
	}
	return
}

func dateToString(d *strfmt.Date) (s string) {
	if d != nil {
		s = d.String()
	}
	return
}

func DateTimeToString(dt *strfmt.DateTime) (s string) {
	if dt != nil {
		s = dt.String()
	}
	return
}
