package permissions

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sonarqube_permissions", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "sonarqube"
		r.ShortGroup = "permissions"
		// This resource can be applied to a group that is specified
		// as an input. And by defining it as a reference to Group
		// object, we can build cross resource referencing.
		r.References["group"] = config.Reference{
			TerraformName: "sonarqube_group",
		}
		// This resource can be applied to a user that is specified
		// as an input. And by defining it as a reference to User
		// object, we can build cross resource referencing.
		r.References["user"] = config.Reference{
			TerraformName: "sonarqube_user",
		}
		// This resource can be applied to a project that is specified
		// as an input. And by defining it as a reference to Project
		// object, we can build cross resource referencing.
		r.References["project"] = config.Reference{
			TerraformName: "sonarqube_project",
		}
		// This resource can be applied to a permission template that is specified
		// as an input. And by defining it as a reference to Template
		// object, we can build cross resource referencing.
		r.References["permissiontemplate"] = config.Reference{
			TerraformName: "sonarqube_permission_template",
		}
	})
}
