package utils

import "github.com/graphql-go/graphql"

func MergeGqlFieldMaps(first, second graphql.Fields) graphql.Fields {
	for k, v := range first {
		second[k] = v
	}
	return second
}
