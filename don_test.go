package don

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDons(t *testing.T) {
	dons, err := Parse("don.yml")
	if err != nil {
		t.Fatal(err)
	}

	premium := dons.Get("premium")
	assert.Equal(t, premium.Level, 1)
	assert.Equal(t, "4.99", premium.String("price"))
	assert.Equal(t, 1, premium.Int("accounts"))
	assert.Contains(t, premium.Scopes, "adsFree")
	assert.Contains(t, premium.Scopes, "offline")
	assert.True(t, premium.Scope("adsFree"))
	assert.True(t, premium.Scope("offline"))

	student := dons.Get("student")
	assert.Equal(t, 2, student.Level)
	assert.Equal(t, "premium", student.Inherit)
	assert.Equal(t, "2.49", student.String("price"))
	assert.Equal(t, 1, student.Int("accounts"))
	assert.True(t, student.Scope("adsFree"))
	assert.True(t, student.Scope("offline"))
	assert.True(t, student.Scope("discount"))

	family := dons.Get("family")
	assert.Equal(t, 3, family.Level)
	assert.Equal(t, "premium", family.Inherit)
	assert.Equal(t, "7.99", family.String("price"))
	assert.Equal(t, 6, family.Int("accounts"))

	assert.Equal(t, map[string]interface{}{
		"price":    "7.99",
		"accounts": 6,
	}, family.Meta)

	for _, scope := range premium.Scopes {
		assert.Contains(t, family.Scopes, scope)
	}
}
