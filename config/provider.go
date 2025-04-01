/*
Copyright 2021 Upbound Inc.

Modifications Copyright 2025 BrainGu, LLC
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/smoothglue/provider-sonarqube/config/group"
	"github.com/smoothglue/provider-sonarqube/config/groupmember"
	"github.com/smoothglue/provider-sonarqube/config/permissions"
	"github.com/smoothglue/provider-sonarqube/config/permissiontemplate"
	"github.com/smoothglue/provider-sonarqube/config/project"
	"github.com/smoothglue/provider-sonarqube/config/projectmainbranch"
	"github.com/smoothglue/provider-sonarqube/config/user"
	"github.com/smoothglue/provider-sonarqube/config/webhook"
)

const (
	resourcePrefix = "sonarqube"
	modulePath     = "github.com/smoothglue/provider-sonarqube"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("sonarqube.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		group.Configure,
		groupmember.Configure,
		permissions.Configure,
		permissiontemplate.Configure,
		project.Configure,
		projectmainbranch.Configure,
		user.Configure,
		webhook.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
