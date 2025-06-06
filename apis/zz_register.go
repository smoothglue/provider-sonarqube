// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/smoothglue/provider-sonarqube/apis/group/v1alpha1"
	v1alpha1groupmember "github.com/smoothglue/provider-sonarqube/apis/groupmember/v1alpha1"
	v1alpha1permissions "github.com/smoothglue/provider-sonarqube/apis/permissions/v1alpha1"
	v1alpha1permissiontemplate "github.com/smoothglue/provider-sonarqube/apis/permissiontemplate/v1alpha1"
	v1alpha1project "github.com/smoothglue/provider-sonarqube/apis/project/v1alpha1"
	v1alpha1projectmainbranch "github.com/smoothglue/provider-sonarqube/apis/projectmainbranch/v1alpha1"
	v1alpha1user "github.com/smoothglue/provider-sonarqube/apis/user/v1alpha1"
	v1alpha1apis "github.com/smoothglue/provider-sonarqube/apis/v1alpha1"
	v1beta1 "github.com/smoothglue/provider-sonarqube/apis/v1beta1"
	v1alpha1webhook "github.com/smoothglue/provider-sonarqube/apis/webhook/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1groupmember.SchemeBuilder.AddToScheme,
		v1alpha1permissions.SchemeBuilder.AddToScheme,
		v1alpha1permissiontemplate.SchemeBuilder.AddToScheme,
		v1alpha1project.SchemeBuilder.AddToScheme,
		v1alpha1projectmainbranch.SchemeBuilder.AddToScheme,
		v1alpha1user.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1webhook.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
