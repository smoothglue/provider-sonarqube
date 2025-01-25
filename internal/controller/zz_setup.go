// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	group "github.com/smoothglue/provider-sonarqube/internal/controller/group/group"
	member "github.com/smoothglue/provider-sonarqube/internal/controller/groupmember/member"
	permissions "github.com/smoothglue/provider-sonarqube/internal/controller/permissions/permissions"
	template "github.com/smoothglue/provider-sonarqube/internal/controller/permissiontemplate/template"
	project "github.com/smoothglue/provider-sonarqube/internal/controller/project/project"
	mainbranch "github.com/smoothglue/provider-sonarqube/internal/controller/projectmainbranch/mainbranch"
	providerconfig "github.com/smoothglue/provider-sonarqube/internal/controller/providerconfig"
	user "github.com/smoothglue/provider-sonarqube/internal/controller/user/user"
	webhook "github.com/smoothglue/provider-sonarqube/internal/controller/webhook/webhook"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		member.Setup,
		permissions.Setup,
		template.Setup,
		project.Setup,
		mainbranch.Setup,
		providerconfig.Setup,
		user.Setup,
		webhook.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
