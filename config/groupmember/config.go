package groupmember

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_group_member", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "sonarqube"
		r.ShortGroup = "groupmember"
		// This resource needs the group in which the member belongs
		// as an input. And by defining it as a reference to Group
		// object, we can build cross resource referencing.
		r.References["group"] = config.Reference{
			TerraformName: "sonarqube_group",
		}
		// This resource needs the user who is the group member
		// as an input. And by defining it as a reference to User
		// object, we can build cross resource referencing.
		r.References["user"] = config.Reference{
			TerraformName: "sonarqube_user",
		}
	})
}
