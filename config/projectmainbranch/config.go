package projectmainbranch

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_project_main_branch", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "sonarqube"
		r.ShortGroup = "projectmainbranch"
		// This resource can be applied to a project that is specified
		// as an input. And by defining it as a reference to Project
		// object, we can build cross resource referencing.
		r.References["project"] = config.Reference{
			TerraformName: "sonarqube_project",
		}
	})
}
