package extensions

import (

	"github.com/pseudomuto/protoc-gen-doc/extensions"
	"google.golang.org/genproto/googleapis/api/visibility"
	// "google.golang.org/genproto/googleapis/api/annotations"
)
// HTTPRule represents a single HTTP rule from the (google.api.http) method option extension.
type VisibilityRule struct {
	Selector  string `json:"selector"`
	Restriction string `json:"restriction"`
}

// HTTPExtension contains the rules set by the (google.api.http) method option extension.
type GRPCExtension struct {
	Rules VisibilityRule `json:"rules"`
}

func getRule(r *visibility.VisibilityRule) (rule VisibilityRule) {
	// switch r.GetRestriction().(type) {
	// case *annotations.HttpRule_Get:
	// 	rule.Selector = "selectr"
	// 	rule.Restriction = r.GetGet()
	// case *annotations.HttpRule_Put:
	// 	rule.Selector = "selectr"
	// 	rule.Restriction = r.GetPut()
	// case *annotations.HttpRule_Post:
	// 	rule.Selector = "selectr"
	// 	rule.Restriction = r.GetPost()
	// case *annotations.HttpRule_Custom:
	// 	custom := r.GetCustom()
	// 	rule.Selector = custom.GetKind()
	// 	rule.Restriction = custom.GetPath()
	// }
	rule.Restriction = r.GetRestriction()
	rule.Selector = ""
	return
}

func init() {
	extensions.SetTransformer("google.api.method_visibility", func(payload interface{}) interface{} {
		var result VisibilityRule
		rule , ok := payload.(*visibility.VisibilityRule)
		if !ok {
			return nil
		}

		result = getRule(rule)


		// rules = append(rules, getRule(rule))
		// result = getRule(rule)
		//NOTE: The option can only have one level of nested AdditionalBindings.
		// for _, rule := range rule.AdditionalBindings {
		// 	rules = append(rules, getRule(rule))
		// }

		return GRPCExtension{Rules:result}
	})
}
