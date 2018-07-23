package common

import (
	"strings"
)

type Tag struct {
	Key   string `yaml:"Key"`
	Value string `yaml:"Value"`
}

func SplitStrArray(asset string) []string {
	if len(asset) > 0 {
		strArray := strings.Split(asset, ",")
		return strArray
	} else {
		return nil
	}
}

func GenMap(asset map[string]string) []map[string]string {
	arrayMap := make([]map[string]string, 0)
	arrayMap = append(arrayMap, asset)
	return arrayMap
}

func GenTags(Tags map[string]string) []Tag {
	arrayTags := make([]Tag, 0)
	for k, v := range Tags {
		arrayTags = append(arrayTags, Tag{Key: k, Value: v})
	}
	return arrayTags
}
