// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package jfrogproject

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/skycaptain/pulumi-jfrog-project/sdk/go/jfrog-project/internal"
)

var _ = internal.GetEnvOrDefault

type ProjectAdminPrivilege struct {
	// Enables a project admin to define the resources to be indexed by Xray
	IndexResources bool `pulumi:"indexResources"`
	// Allows the Project Admin to manage Platform users/groups as project members with different roles.
	ManageMembers bool `pulumi:"manageMembers"`
	// Allows the Project Admin to manage resources - repositories, builds and Pipelines resources on the project level.
	ManageResources bool `pulumi:"manageResources"`
}

// ProjectAdminPrivilegeInput is an input type that accepts ProjectAdminPrivilegeArgs and ProjectAdminPrivilegeOutput values.
// You can construct a concrete instance of `ProjectAdminPrivilegeInput` via:
//
//	ProjectAdminPrivilegeArgs{...}
type ProjectAdminPrivilegeInput interface {
	pulumi.Input

	ToProjectAdminPrivilegeOutput() ProjectAdminPrivilegeOutput
	ToProjectAdminPrivilegeOutputWithContext(context.Context) ProjectAdminPrivilegeOutput
}

type ProjectAdminPrivilegeArgs struct {
	// Enables a project admin to define the resources to be indexed by Xray
	IndexResources pulumi.BoolInput `pulumi:"indexResources"`
	// Allows the Project Admin to manage Platform users/groups as project members with different roles.
	ManageMembers pulumi.BoolInput `pulumi:"manageMembers"`
	// Allows the Project Admin to manage resources - repositories, builds and Pipelines resources on the project level.
	ManageResources pulumi.BoolInput `pulumi:"manageResources"`
}

func (ProjectAdminPrivilegeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectAdminPrivilege)(nil)).Elem()
}

func (i ProjectAdminPrivilegeArgs) ToProjectAdminPrivilegeOutput() ProjectAdminPrivilegeOutput {
	return i.ToProjectAdminPrivilegeOutputWithContext(context.Background())
}

func (i ProjectAdminPrivilegeArgs) ToProjectAdminPrivilegeOutputWithContext(ctx context.Context) ProjectAdminPrivilegeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectAdminPrivilegeOutput)
}

// ProjectAdminPrivilegeArrayInput is an input type that accepts ProjectAdminPrivilegeArray and ProjectAdminPrivilegeArrayOutput values.
// You can construct a concrete instance of `ProjectAdminPrivilegeArrayInput` via:
//
//	ProjectAdminPrivilegeArray{ ProjectAdminPrivilegeArgs{...} }
type ProjectAdminPrivilegeArrayInput interface {
	pulumi.Input

	ToProjectAdminPrivilegeArrayOutput() ProjectAdminPrivilegeArrayOutput
	ToProjectAdminPrivilegeArrayOutputWithContext(context.Context) ProjectAdminPrivilegeArrayOutput
}

type ProjectAdminPrivilegeArray []ProjectAdminPrivilegeInput

func (ProjectAdminPrivilegeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectAdminPrivilege)(nil)).Elem()
}

func (i ProjectAdminPrivilegeArray) ToProjectAdminPrivilegeArrayOutput() ProjectAdminPrivilegeArrayOutput {
	return i.ToProjectAdminPrivilegeArrayOutputWithContext(context.Background())
}

func (i ProjectAdminPrivilegeArray) ToProjectAdminPrivilegeArrayOutputWithContext(ctx context.Context) ProjectAdminPrivilegeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectAdminPrivilegeArrayOutput)
}

type ProjectAdminPrivilegeOutput struct{ *pulumi.OutputState }

func (ProjectAdminPrivilegeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectAdminPrivilege)(nil)).Elem()
}

func (o ProjectAdminPrivilegeOutput) ToProjectAdminPrivilegeOutput() ProjectAdminPrivilegeOutput {
	return o
}

func (o ProjectAdminPrivilegeOutput) ToProjectAdminPrivilegeOutputWithContext(ctx context.Context) ProjectAdminPrivilegeOutput {
	return o
}

// Enables a project admin to define the resources to be indexed by Xray
func (o ProjectAdminPrivilegeOutput) IndexResources() pulumi.BoolOutput {
	return o.ApplyT(func(v ProjectAdminPrivilege) bool { return v.IndexResources }).(pulumi.BoolOutput)
}

// Allows the Project Admin to manage Platform users/groups as project members with different roles.
func (o ProjectAdminPrivilegeOutput) ManageMembers() pulumi.BoolOutput {
	return o.ApplyT(func(v ProjectAdminPrivilege) bool { return v.ManageMembers }).(pulumi.BoolOutput)
}

// Allows the Project Admin to manage resources - repositories, builds and Pipelines resources on the project level.
func (o ProjectAdminPrivilegeOutput) ManageResources() pulumi.BoolOutput {
	return o.ApplyT(func(v ProjectAdminPrivilege) bool { return v.ManageResources }).(pulumi.BoolOutput)
}

type ProjectAdminPrivilegeArrayOutput struct{ *pulumi.OutputState }

func (ProjectAdminPrivilegeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectAdminPrivilege)(nil)).Elem()
}

func (o ProjectAdminPrivilegeArrayOutput) ToProjectAdminPrivilegeArrayOutput() ProjectAdminPrivilegeArrayOutput {
	return o
}

func (o ProjectAdminPrivilegeArrayOutput) ToProjectAdminPrivilegeArrayOutputWithContext(ctx context.Context) ProjectAdminPrivilegeArrayOutput {
	return o
}

func (o ProjectAdminPrivilegeArrayOutput) Index(i pulumi.IntInput) ProjectAdminPrivilegeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProjectAdminPrivilege {
		return vs[0].([]ProjectAdminPrivilege)[vs[1].(int)]
	}).(ProjectAdminPrivilegeOutput)
}

type ProjectGroup struct {
	// Must be existing Artifactory group
	Name string `pulumi:"name"`
	// List of pre-defined Project or custom roles
	Roles []string `pulumi:"roles"`
}

// ProjectGroupInput is an input type that accepts ProjectGroupArgs and ProjectGroupOutput values.
// You can construct a concrete instance of `ProjectGroupInput` via:
//
//	ProjectGroupArgs{...}
type ProjectGroupInput interface {
	pulumi.Input

	ToProjectGroupOutput() ProjectGroupOutput
	ToProjectGroupOutputWithContext(context.Context) ProjectGroupOutput
}

type ProjectGroupArgs struct {
	// Must be existing Artifactory group
	Name pulumi.StringInput `pulumi:"name"`
	// List of pre-defined Project or custom roles
	Roles pulumi.StringArrayInput `pulumi:"roles"`
}

func (ProjectGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectGroup)(nil)).Elem()
}

func (i ProjectGroupArgs) ToProjectGroupOutput() ProjectGroupOutput {
	return i.ToProjectGroupOutputWithContext(context.Background())
}

func (i ProjectGroupArgs) ToProjectGroupOutputWithContext(ctx context.Context) ProjectGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectGroupOutput)
}

// ProjectGroupArrayInput is an input type that accepts ProjectGroupArray and ProjectGroupArrayOutput values.
// You can construct a concrete instance of `ProjectGroupArrayInput` via:
//
//	ProjectGroupArray{ ProjectGroupArgs{...} }
type ProjectGroupArrayInput interface {
	pulumi.Input

	ToProjectGroupArrayOutput() ProjectGroupArrayOutput
	ToProjectGroupArrayOutputWithContext(context.Context) ProjectGroupArrayOutput
}

type ProjectGroupArray []ProjectGroupInput

func (ProjectGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectGroup)(nil)).Elem()
}

func (i ProjectGroupArray) ToProjectGroupArrayOutput() ProjectGroupArrayOutput {
	return i.ToProjectGroupArrayOutputWithContext(context.Background())
}

func (i ProjectGroupArray) ToProjectGroupArrayOutputWithContext(ctx context.Context) ProjectGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectGroupArrayOutput)
}

type ProjectGroupOutput struct{ *pulumi.OutputState }

func (ProjectGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectGroup)(nil)).Elem()
}

func (o ProjectGroupOutput) ToProjectGroupOutput() ProjectGroupOutput {
	return o
}

func (o ProjectGroupOutput) ToProjectGroupOutputWithContext(ctx context.Context) ProjectGroupOutput {
	return o
}

// Must be existing Artifactory group
func (o ProjectGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectGroup) string { return v.Name }).(pulumi.StringOutput)
}

// List of pre-defined Project or custom roles
func (o ProjectGroupOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProjectGroup) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

type ProjectGroupArrayOutput struct{ *pulumi.OutputState }

func (ProjectGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectGroup)(nil)).Elem()
}

func (o ProjectGroupArrayOutput) ToProjectGroupArrayOutput() ProjectGroupArrayOutput {
	return o
}

func (o ProjectGroupArrayOutput) ToProjectGroupArrayOutputWithContext(ctx context.Context) ProjectGroupArrayOutput {
	return o
}

func (o ProjectGroupArrayOutput) Index(i pulumi.IntInput) ProjectGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProjectGroup {
		return vs[0].([]ProjectGroup)[vs[1].(int)]
	}).(ProjectGroupOutput)
}

type ProjectMember struct {
	// Must be existing Artifactory user
	Name string `pulumi:"name"`
	// List of pre-defined Project or custom roles
	Roles []string `pulumi:"roles"`
}

// ProjectMemberInput is an input type that accepts ProjectMemberArgs and ProjectMemberOutput values.
// You can construct a concrete instance of `ProjectMemberInput` via:
//
//	ProjectMemberArgs{...}
type ProjectMemberInput interface {
	pulumi.Input

	ToProjectMemberOutput() ProjectMemberOutput
	ToProjectMemberOutputWithContext(context.Context) ProjectMemberOutput
}

type ProjectMemberArgs struct {
	// Must be existing Artifactory user
	Name pulumi.StringInput `pulumi:"name"`
	// List of pre-defined Project or custom roles
	Roles pulumi.StringArrayInput `pulumi:"roles"`
}

func (ProjectMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectMember)(nil)).Elem()
}

func (i ProjectMemberArgs) ToProjectMemberOutput() ProjectMemberOutput {
	return i.ToProjectMemberOutputWithContext(context.Background())
}

func (i ProjectMemberArgs) ToProjectMemberOutputWithContext(ctx context.Context) ProjectMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberOutput)
}

// ProjectMemberArrayInput is an input type that accepts ProjectMemberArray and ProjectMemberArrayOutput values.
// You can construct a concrete instance of `ProjectMemberArrayInput` via:
//
//	ProjectMemberArray{ ProjectMemberArgs{...} }
type ProjectMemberArrayInput interface {
	pulumi.Input

	ToProjectMemberArrayOutput() ProjectMemberArrayOutput
	ToProjectMemberArrayOutputWithContext(context.Context) ProjectMemberArrayOutput
}

type ProjectMemberArray []ProjectMemberInput

func (ProjectMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectMember)(nil)).Elem()
}

func (i ProjectMemberArray) ToProjectMemberArrayOutput() ProjectMemberArrayOutput {
	return i.ToProjectMemberArrayOutputWithContext(context.Background())
}

func (i ProjectMemberArray) ToProjectMemberArrayOutputWithContext(ctx context.Context) ProjectMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberArrayOutput)
}

type ProjectMemberOutput struct{ *pulumi.OutputState }

func (ProjectMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectMember)(nil)).Elem()
}

func (o ProjectMemberOutput) ToProjectMemberOutput() ProjectMemberOutput {
	return o
}

func (o ProjectMemberOutput) ToProjectMemberOutputWithContext(ctx context.Context) ProjectMemberOutput {
	return o
}

// Must be existing Artifactory user
func (o ProjectMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectMember) string { return v.Name }).(pulumi.StringOutput)
}

// List of pre-defined Project or custom roles
func (o ProjectMemberOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProjectMember) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

type ProjectMemberArrayOutput struct{ *pulumi.OutputState }

func (ProjectMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectMember)(nil)).Elem()
}

func (o ProjectMemberArrayOutput) ToProjectMemberArrayOutput() ProjectMemberArrayOutput {
	return o
}

func (o ProjectMemberArrayOutput) ToProjectMemberArrayOutputWithContext(ctx context.Context) ProjectMemberArrayOutput {
	return o
}

func (o ProjectMemberArrayOutput) Index(i pulumi.IntInput) ProjectMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProjectMember {
		return vs[0].([]ProjectMember)[vs[1].(int)]
	}).(ProjectMemberOutput)
}

type ProjectRole struct {
	// List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
	Actions     []string `pulumi:"actions"`
	Description *string  `pulumi:"description"`
	// A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
	Environments []string `pulumi:"environments"`
	Name         string   `pulumi:"name"`
	// Type of role. Only "CUSTOM" is supported
	Type *string `pulumi:"type"`
}

// ProjectRoleInput is an input type that accepts ProjectRoleArgs and ProjectRoleOutput values.
// You can construct a concrete instance of `ProjectRoleInput` via:
//
//	ProjectRoleArgs{...}
type ProjectRoleInput interface {
	pulumi.Input

	ToProjectRoleOutput() ProjectRoleOutput
	ToProjectRoleOutputWithContext(context.Context) ProjectRoleOutput
}

type ProjectRoleArgs struct {
	// List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
	Actions     pulumi.StringArrayInput `pulumi:"actions"`
	Description pulumi.StringPtrInput   `pulumi:"description"`
	// A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
	Environments pulumi.StringArrayInput `pulumi:"environments"`
	Name         pulumi.StringInput      `pulumi:"name"`
	// Type of role. Only "CUSTOM" is supported
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ProjectRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectRole)(nil)).Elem()
}

func (i ProjectRoleArgs) ToProjectRoleOutput() ProjectRoleOutput {
	return i.ToProjectRoleOutputWithContext(context.Background())
}

func (i ProjectRoleArgs) ToProjectRoleOutputWithContext(ctx context.Context) ProjectRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectRoleOutput)
}

// ProjectRoleArrayInput is an input type that accepts ProjectRoleArray and ProjectRoleArrayOutput values.
// You can construct a concrete instance of `ProjectRoleArrayInput` via:
//
//	ProjectRoleArray{ ProjectRoleArgs{...} }
type ProjectRoleArrayInput interface {
	pulumi.Input

	ToProjectRoleArrayOutput() ProjectRoleArrayOutput
	ToProjectRoleArrayOutputWithContext(context.Context) ProjectRoleArrayOutput
}

type ProjectRoleArray []ProjectRoleInput

func (ProjectRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectRole)(nil)).Elem()
}

func (i ProjectRoleArray) ToProjectRoleArrayOutput() ProjectRoleArrayOutput {
	return i.ToProjectRoleArrayOutputWithContext(context.Background())
}

func (i ProjectRoleArray) ToProjectRoleArrayOutputWithContext(ctx context.Context) ProjectRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectRoleArrayOutput)
}

type ProjectRoleOutput struct{ *pulumi.OutputState }

func (ProjectRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectRole)(nil)).Elem()
}

func (o ProjectRoleOutput) ToProjectRoleOutput() ProjectRoleOutput {
	return o
}

func (o ProjectRoleOutput) ToProjectRoleOutputWithContext(ctx context.Context) ProjectRoleOutput {
	return o
}

// List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
func (o ProjectRoleOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProjectRole) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o ProjectRoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectRole) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
func (o ProjectRoleOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProjectRole) []string { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o ProjectRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectRole) string { return v.Name }).(pulumi.StringOutput)
}

// Type of role. Only "CUSTOM" is supported
func (o ProjectRoleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectRole) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ProjectRoleArrayOutput struct{ *pulumi.OutputState }

func (ProjectRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectRole)(nil)).Elem()
}

func (o ProjectRoleArrayOutput) ToProjectRoleArrayOutput() ProjectRoleArrayOutput {
	return o
}

func (o ProjectRoleArrayOutput) ToProjectRoleArrayOutputWithContext(ctx context.Context) ProjectRoleArrayOutput {
	return o
}

func (o ProjectRoleArrayOutput) Index(i pulumi.IntInput) ProjectRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProjectRole {
		return vs[0].([]ProjectRole)[vs[1].(int)]
	}).(ProjectRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectAdminPrivilegeInput)(nil)).Elem(), ProjectAdminPrivilegeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectAdminPrivilegeArrayInput)(nil)).Elem(), ProjectAdminPrivilegeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectGroupInput)(nil)).Elem(), ProjectGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectGroupArrayInput)(nil)).Elem(), ProjectGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberInput)(nil)).Elem(), ProjectMemberArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberArrayInput)(nil)).Elem(), ProjectMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectRoleInput)(nil)).Elem(), ProjectRoleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectRoleArrayInput)(nil)).Elem(), ProjectRoleArray{})
	pulumi.RegisterOutputType(ProjectAdminPrivilegeOutput{})
	pulumi.RegisterOutputType(ProjectAdminPrivilegeArrayOutput{})
	pulumi.RegisterOutputType(ProjectGroupOutput{})
	pulumi.RegisterOutputType(ProjectGroupArrayOutput{})
	pulumi.RegisterOutputType(ProjectMemberOutput{})
	pulumi.RegisterOutputType(ProjectMemberArrayOutput{})
	pulumi.RegisterOutputType(ProjectRoleOutput{})
	pulumi.RegisterOutputType(ProjectRoleArrayOutput{})
}
