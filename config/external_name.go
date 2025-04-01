/*
Copyright 2022 Upbound Inc.

Modifications Copyright 2025 BrainGu, LLC
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Note that commented out configurations exist in terraform as resources
	// to import, but haven't (yet) been implemented in this provider.
	// "sonarqube_gitlab_binding": config.IdentifierFromProvider,
	"sonarqube_group":               config.IdentifierFromProvider,
	"sonarqube_group_member":        config.IdentifierFromProvider,
	"sonarqube_permission_template": config.IdentifierFromProvider,
	"sonarqube_permissions":         config.IdentifierFromProvider,
	// "sonarqube_plugin":                        config.IdentifierFromProvider,
	// "sonarqube_portfolio":                     config.IdentifierFromProvider,
	"sonarqube_project":             config.IdentifierFromProvider,
	"sonarqube_project_main_branch": config.IdentifierFromProvider,
	// "sonarqube_qualitygate":                     config.IdentifierFromProvider,
	// "sonarqube_qualitygate_project_association": config.IdentifierFromProvider,
	// "sonarqube_qualityprofile":                     config.IdentifierFromProvider,
	// "sonarqube_qualityprofile_project_association": config.IdentifierFromProvider,
	"sonarqube_user":    config.IdentifierFromProvider,
	"sonarqube_webhook": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
