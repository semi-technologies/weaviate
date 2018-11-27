package local

import (
	"fmt"
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	local_get "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get"
	local_get_meta "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get_meta"
	"github.com/graphql-go/graphql"
)

// Build the local queries from the database schema.
func Build(dbSchema *schema.Schema) (*graphql.Field, error) {
	getField, err := local_get.Build(dbSchema)
	if err != nil {
		return nil, err
	}

	getMetaField, err := local_get_meta.Build(dbSchema)
	if err != nil {
		return nil, err
	}

	localObject := graphql.NewObject(graphql.ObjectConfig{
		Name: "WeaviateLocalObj",
		Fields: graphql.Fields{
			"Get":     getField,
			"GetMeta": getMetaField,
		},
		Description: "Type of query on the local Weaviate",
	})

	localField := graphql.Field{
		Type:        localObject,
		Description: "Query a local Weaviate instance",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			fmt.Printf("- localGetAndMetaObjectResolver (pass on source; the resolver)\n")
			// This step does nothing; all ways allow the resolver to continue
			return p.Source, nil
		},
	}

	return &localField, nil
}
