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

// Share a local or remote repository with all projects. Project Members of the target project are granted actions to the shared repository according to their Roles and Role actions assigned in the target Project. Requires a user assigned with the 'Administer the Platform' role.
//
// ->Only available for Artifactory 7.90.1 or later.
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
//			_, err := jfrogproject.NewShareRepositoryWithAll(ctx, "myprojectsharerepo", &jfrogproject.ShareRepositoryWithAllArgs{
//				RepoKey: pulumi.String("myrepo-generic-local"),
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
// $ pulumi import jfrog-project:index/shareRepositoryWithAll:ShareRepositoryWithAll myprojectsharerepo repo_key
// ```
type ShareRepositoryWithAll struct {
	pulumi.CustomResourceState

	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly pulumi.BoolOutput `pulumi:"readOnly"`
	// The key of the repository.
	RepoKey pulumi.StringOutput `pulumi:"repoKey"`
}

// NewShareRepositoryWithAll registers a new resource with the given unique name, arguments, and options.
func NewShareRepositoryWithAll(ctx *pulumi.Context,
	name string, args *ShareRepositoryWithAllArgs, opts ...pulumi.ResourceOption) (*ShareRepositoryWithAll, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepoKey == nil {
		return nil, errors.New("invalid value for required argument 'RepoKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ShareRepositoryWithAll
	err := ctx.RegisterResource("jfrog-project:index/shareRepositoryWithAll:ShareRepositoryWithAll", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShareRepositoryWithAll gets an existing ShareRepositoryWithAll resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShareRepositoryWithAll(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareRepositoryWithAllState, opts ...pulumi.ResourceOption) (*ShareRepositoryWithAll, error) {
	var resource ShareRepositoryWithAll
	err := ctx.ReadResource("jfrog-project:index/shareRepositoryWithAll:ShareRepositoryWithAll", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShareRepositoryWithAll resources.
type shareRepositoryWithAllState struct {
	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly *bool `pulumi:"readOnly"`
	// The key of the repository.
	RepoKey *string `pulumi:"repoKey"`
}

type ShareRepositoryWithAllState struct {
	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly pulumi.BoolPtrInput
	// The key of the repository.
	RepoKey pulumi.StringPtrInput
}

func (ShareRepositoryWithAllState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareRepositoryWithAllState)(nil)).Elem()
}

type shareRepositoryWithAllArgs struct {
	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly *bool `pulumi:"readOnly"`
	// The key of the repository.
	RepoKey string `pulumi:"repoKey"`
}

// The set of arguments for constructing a ShareRepositoryWithAll resource.
type ShareRepositoryWithAllArgs struct {
	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly pulumi.BoolPtrInput
	// The key of the repository.
	RepoKey pulumi.StringInput
}

func (ShareRepositoryWithAllArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareRepositoryWithAllArgs)(nil)).Elem()
}

type ShareRepositoryWithAllInput interface {
	pulumi.Input

	ToShareRepositoryWithAllOutput() ShareRepositoryWithAllOutput
	ToShareRepositoryWithAllOutputWithContext(ctx context.Context) ShareRepositoryWithAllOutput
}

func (*ShareRepositoryWithAll) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareRepositoryWithAll)(nil)).Elem()
}

func (i *ShareRepositoryWithAll) ToShareRepositoryWithAllOutput() ShareRepositoryWithAllOutput {
	return i.ToShareRepositoryWithAllOutputWithContext(context.Background())
}

func (i *ShareRepositoryWithAll) ToShareRepositoryWithAllOutputWithContext(ctx context.Context) ShareRepositoryWithAllOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareRepositoryWithAllOutput)
}

// ShareRepositoryWithAllArrayInput is an input type that accepts ShareRepositoryWithAllArray and ShareRepositoryWithAllArrayOutput values.
// You can construct a concrete instance of `ShareRepositoryWithAllArrayInput` via:
//
//	ShareRepositoryWithAllArray{ ShareRepositoryWithAllArgs{...} }
type ShareRepositoryWithAllArrayInput interface {
	pulumi.Input

	ToShareRepositoryWithAllArrayOutput() ShareRepositoryWithAllArrayOutput
	ToShareRepositoryWithAllArrayOutputWithContext(context.Context) ShareRepositoryWithAllArrayOutput
}

type ShareRepositoryWithAllArray []ShareRepositoryWithAllInput

func (ShareRepositoryWithAllArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShareRepositoryWithAll)(nil)).Elem()
}

func (i ShareRepositoryWithAllArray) ToShareRepositoryWithAllArrayOutput() ShareRepositoryWithAllArrayOutput {
	return i.ToShareRepositoryWithAllArrayOutputWithContext(context.Background())
}

func (i ShareRepositoryWithAllArray) ToShareRepositoryWithAllArrayOutputWithContext(ctx context.Context) ShareRepositoryWithAllArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareRepositoryWithAllArrayOutput)
}

// ShareRepositoryWithAllMapInput is an input type that accepts ShareRepositoryWithAllMap and ShareRepositoryWithAllMapOutput values.
// You can construct a concrete instance of `ShareRepositoryWithAllMapInput` via:
//
//	ShareRepositoryWithAllMap{ "key": ShareRepositoryWithAllArgs{...} }
type ShareRepositoryWithAllMapInput interface {
	pulumi.Input

	ToShareRepositoryWithAllMapOutput() ShareRepositoryWithAllMapOutput
	ToShareRepositoryWithAllMapOutputWithContext(context.Context) ShareRepositoryWithAllMapOutput
}

type ShareRepositoryWithAllMap map[string]ShareRepositoryWithAllInput

func (ShareRepositoryWithAllMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShareRepositoryWithAll)(nil)).Elem()
}

func (i ShareRepositoryWithAllMap) ToShareRepositoryWithAllMapOutput() ShareRepositoryWithAllMapOutput {
	return i.ToShareRepositoryWithAllMapOutputWithContext(context.Background())
}

func (i ShareRepositoryWithAllMap) ToShareRepositoryWithAllMapOutputWithContext(ctx context.Context) ShareRepositoryWithAllMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareRepositoryWithAllMapOutput)
}

type ShareRepositoryWithAllOutput struct{ *pulumi.OutputState }

func (ShareRepositoryWithAllOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareRepositoryWithAll)(nil)).Elem()
}

func (o ShareRepositoryWithAllOutput) ToShareRepositoryWithAllOutput() ShareRepositoryWithAllOutput {
	return o
}

func (o ShareRepositoryWithAllOutput) ToShareRepositoryWithAllOutputWithContext(ctx context.Context) ShareRepositoryWithAllOutput {
	return o
}

// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
func (o ShareRepositoryWithAllOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *ShareRepositoryWithAll) pulumi.BoolOutput { return v.ReadOnly }).(pulumi.BoolOutput)
}

// The key of the repository.
func (o ShareRepositoryWithAllOutput) RepoKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareRepositoryWithAll) pulumi.StringOutput { return v.RepoKey }).(pulumi.StringOutput)
}

type ShareRepositoryWithAllArrayOutput struct{ *pulumi.OutputState }

func (ShareRepositoryWithAllArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShareRepositoryWithAll)(nil)).Elem()
}

func (o ShareRepositoryWithAllArrayOutput) ToShareRepositoryWithAllArrayOutput() ShareRepositoryWithAllArrayOutput {
	return o
}

func (o ShareRepositoryWithAllArrayOutput) ToShareRepositoryWithAllArrayOutputWithContext(ctx context.Context) ShareRepositoryWithAllArrayOutput {
	return o
}

func (o ShareRepositoryWithAllArrayOutput) Index(i pulumi.IntInput) ShareRepositoryWithAllOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ShareRepositoryWithAll {
		return vs[0].([]*ShareRepositoryWithAll)[vs[1].(int)]
	}).(ShareRepositoryWithAllOutput)
}

type ShareRepositoryWithAllMapOutput struct{ *pulumi.OutputState }

func (ShareRepositoryWithAllMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShareRepositoryWithAll)(nil)).Elem()
}

func (o ShareRepositoryWithAllMapOutput) ToShareRepositoryWithAllMapOutput() ShareRepositoryWithAllMapOutput {
	return o
}

func (o ShareRepositoryWithAllMapOutput) ToShareRepositoryWithAllMapOutputWithContext(ctx context.Context) ShareRepositoryWithAllMapOutput {
	return o
}

func (o ShareRepositoryWithAllMapOutput) MapIndex(k pulumi.StringInput) ShareRepositoryWithAllOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ShareRepositoryWithAll {
		return vs[0].(map[string]*ShareRepositoryWithAll)[vs[1].(string)]
	}).(ShareRepositoryWithAllOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShareRepositoryWithAllInput)(nil)).Elem(), &ShareRepositoryWithAll{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareRepositoryWithAllArrayInput)(nil)).Elem(), ShareRepositoryWithAllArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareRepositoryWithAllMapInput)(nil)).Elem(), ShareRepositoryWithAllMap{})
	pulumi.RegisterOutputType(ShareRepositoryWithAllOutput{})
	pulumi.RegisterOutputType(ShareRepositoryWithAllArrayOutput{})
	pulumi.RegisterOutputType(ShareRepositoryWithAllMapOutput{})
}
