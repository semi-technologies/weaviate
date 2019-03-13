package getmeta

import (
	"testing"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	cf "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/common_filters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParamsHashing(t *testing.T) {

	params := func() Params {
		return Params{
			Analytics:        cf.AnalyticsProps{UseAnaltyicsEngine: true},
			IncludeMetaCount: true,
			ClassName:        schema.ClassName("MyBestClass"),
			Filters:          nil,
			Kind:             kind.THING_KIND,
			Properties: []MetaProperty{
				MetaProperty{
					Name:                schema.PropertyName("bestprop"),
					StatisticalAnalyses: []StatisticalAnalysis{Count},
				},
			},
		}
	}
	hash := func() string { return "18c6e361737a82cb99c5223fb2a61ba7" }

	t.Run("it generates a hash", func(t *testing.T) {
		p := params()
		h, err := p.AnalyticsHash()
		require.Nil(t, err)
		assert.Equal(t, h, hash())
	})

	t.Run("it generates the same hash if analytical meta props are changed", func(t *testing.T) {
		p := params()
		p.Analytics.ForceRecalculate = true
		h, err := p.AnalyticsHash()
		require.Nil(t, err)
		assert.Equal(t, hash(), h)
	})

	t.Run("it generates a different hash if a meta prop is changed", func(t *testing.T) {
		p := params()
		p.Properties[0].StatisticalAnalyses[0] = Maximum
		h, err := p.AnalyticsHash()
		require.Nil(t, err)
		assert.NotEqual(t, hash(), h)
	})

	t.Run("it generates a different hash if where filter is added", func(t *testing.T) {
		p := params()
		p.Filters = &cf.LocalFilter{Root: &cf.Clause{Value: &cf.Value{Value: "foo"}}}
		h, err := p.AnalyticsHash()
		require.Nil(t, err)
		assert.NotEqual(t, hash(), h)
	})
}