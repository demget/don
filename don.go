package don

import (
	"time"

	"github.com/spf13/viper"
)

// Don represents a donation entry with particular scopes.
type Don struct {
	meta *viper.Viper
	scop map[string]bool

	// Name is don's name set from key.
	Name string
	// Level displays the don's weight. The higher the level, the better the don is.
	Level int
	// Inherit is used to specify the don's parent, works recursively.
	Inherit string
	// Scopes stores the all available scopes, including ones inherited from parent.
	Scopes []string
	// Meta contains additional don's payload. Also see: String, Int, Float, Duration.
	Meta map[string]interface{}
}

// Dons is a root.
type Dons struct {
	dons map[string]Don
}

// Parse parses the dons from pretty much any kind of file.
func Parse(path string) (*Dons, error) {
	vipr := viper.New()
	vipr.SetConfigFile(path)

	if err := vipr.ReadInConfig(); err != nil {
		return nil, err
	}

	dons := make(map[string]Don)
	if err := vipr.Unmarshal(&dons); err != nil {
		return nil, err
	}

	for key, don := range dons {
		vipr := vipr.Sub(key)
		don.meta = vipr.Sub("meta")
		don.Level = vipr.GetInt("level")
		don.Inherit = vipr.GetString("inherit")
		don.Scopes = vipr.GetStringSlice("scopes")
		don.Meta = vipr.GetStringMap("meta")

		don.Name = key
		dons[key] = don
	}

	for i, don := range dons {
		don.scop = make(map[string]bool)
		makeScopes(&don, dons)
		dons[i] = don
	}

	return &Dons{dons: dons}, nil
}

func makeScopes(don *Don, dons map[string]Don) {
	inherit := don.Inherit
	for inherit != "" && inherit != don.Name {
		par, ok := dons[inherit]
		if !ok {
			continue
		}

		inherit = par.Inherit
		for _, sc := range par.Scopes {
			don.scop[sc] = true
		}
	}

	for _, sc := range don.Scopes {
		don.scop[sc] = true
	}

	don.Scopes = make([]string, 0, len(don.scop))
	for sc := range don.scop {
		don.Scopes = append(don.Scopes, sc)
	}
}

// Get returns a don by the key.
func (ds *Dons) Get(key string) Don {
	return ds.dons[key]
}

// Scope returns true if the don has access to the specified scope.
func (d Don) Scope(key string) bool {
	return d.scop[key]
}

// String returns a string field from the meta.
func (d Don) String(key string) string {
	return d.meta.GetString(key)
}

// Int returns an integer field from the meta.
func (d Don) Int(key string) int {
	return d.meta.GetInt(key)
}

// Float returns a float field from the meta.
func (d Don) Float(key string) float64 {
	return d.meta.GetFloat64(key)
}

// Duration returns a duration field from the meta.
func (d Don) Duration(key string) time.Duration {
	return d.meta.GetDuration(key)
}
