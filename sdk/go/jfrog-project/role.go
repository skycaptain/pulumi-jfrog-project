// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package jfrogproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/skycaptain/pulumi-jfrog-project/sdk/go/jfrog-project/internal"
)

// Create a project role. Element has one to one mapping with the [JFrog Project Roles API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-AddaNewRole). Requires a user assigned with the 'Administer the Platform' role or Project Admin permissions if `admin_privileges.manage_resoures` is enabled.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	jfrogproject "github.com/skycaptain/pulumi-jfrog-project/sdk/go/jfrog-project"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := jfrogproject.NewRole(ctx, "myrole", &jfrogproject.RoleArgs{
//				Type:       pulumi.String("CUSTOM"),
//				ProjectKey: pulumi.Any(project.Myproject.Key),
//				Environments: pulumi.StringArray{
//					pulumi.String("DEV"),
//				},
//				Actions: pulumi.StringArray{
//					pulumi.String("READ_REPOSITORY"),
//					pulumi.String("ANNOTATE_REPOSITORY"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import jfrog-project:index/role:Role myrole project_key:role_name
// ```
type Role struct {
	pulumi.CustomResourceState

	// List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
	Actions pulumi.StringArrayOutput `pulumi:"actions"`
	// A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
	Environments pulumi.StringArrayOutput `pulumi:"environments"`
	Name         pulumi.StringOutput      `pulumi:"name"`
	// Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
	ProjectKey pulumi.StringOutput `pulumi:"projectKey"`
	// Type of role. Only "CUSTOM" is supported
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Environments == nil {
		return nil, errors.New("invalid value for required argument 'Environments'")
	}
	if args.ProjectKey == nil {
		return nil, errors.New("invalid value for required argument 'ProjectKey'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Role
	err := ctx.RegisterResource("jfrog-project:index/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("jfrog-project:index/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
	Actions []string `pulumi:"actions"`
	// A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
	Environments []string `pulumi:"environments"`
	Name         *string  `pulumi:"name"`
	// Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
	ProjectKey *string `pulumi:"projectKey"`
	// Type of role. Only "CUSTOM" is supported
	Type *string `pulumi:"type"`
}

type RoleState struct {
	// List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
	Actions pulumi.StringArrayInput
	// A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
	Environments pulumi.StringArrayInput
	Name         pulumi.StringPtrInput
	// Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
	ProjectKey pulumi.StringPtrInput
	// Type of role. Only "CUSTOM" is supported
	Type pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
	Actions []string `pulumi:"actions"`
	// A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
	Environments []string `pulumi:"environments"`
	Name         *string  `pulumi:"name"`
	// Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
	ProjectKey string `pulumi:"projectKey"`
	// Type of role. Only "CUSTOM" is supported
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
	Actions pulumi.StringArrayInput
	// A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
	Environments pulumi.StringArrayInput
	Name         pulumi.StringPtrInput
	// Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
	ProjectKey pulumi.StringInput
	// Type of role. Only "CUSTOM" is supported
	Type pulumi.StringInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//	RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//	RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

// List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
func (o RoleOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.Actions }).(pulumi.StringArrayOutput)
}

// A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
func (o RoleOutput) Environments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.Environments }).(pulumi.StringArrayOutput)
}

func (o RoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
func (o RoleOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// Type of role. Only "CUSTOM" is supported
func (o RoleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}
