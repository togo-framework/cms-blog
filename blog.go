// Package blog is the Blog content pack for togo: BlogPost / Category / Tag with reading-time, views, featured and slug detail pages.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package blog

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("blog", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
