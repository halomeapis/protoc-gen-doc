package extensions

import (

	"github.com/pseudomuto/protoc-gen-doc/extensions"
	"google.golang.org/genproto/googleapis/api/visibility"
)
// HTTPRule represents a single HTTP rule from the (google.api.http) method option extension.
type VisibilityRule struct {
	Selector  string `json:"selector"`
	Restriction string `json:"restriction"`
}

// HTTPExtension contains the rules set by the (google.api.http) method option extension.
type GRPCExtension struct {
	Rules []VisibilityRule `json:"rules"`
}

func getRule(r *visibility.VisibilityRule) (rule VisibilityRule) {
	rule.Restriction = "INTERNAL"
	rule.Restriction = ""
	return
}

func init() {
	extensions.SetTransformer("google.api.method_visibility", func(payload interface{}) interface{} {
		_ , ok := payload.(*visibility.VisibilityRule)
		if !ok {
			return nil
		}

		return VisibilityRule{Restriction: "INTERNAL",Selector:""}
	})
}
