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

// Share a local or remote repository with a list of projects. Project Members of the target project are granted actions to the shared repository according to their Roles and Role actions assigned in the target Project. Requires a user assigned with the 'Administer the Platform' role.
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
//			_, err := jfrogproject.NewShareRepository(ctx, "myprojectsharerepo", &jfrogproject.ShareRepositoryArgs{
//				RepoKey:          pulumi.String("myrepo-generic-local"),
//				TargetProjectKey: pulumi.String("myproj"),
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
// $ pulumi import jfrog-project:index/shareRepository:ShareRepository myprojectsharerepo repo_key:project_key
// ```
type ShareRepository struct {
	pulumi.CustomResourceState

	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly pulumi.BoolOutput `pulumi:"readOnly"`
	// The key of the repository.
	RepoKey pulumi.StringOutput `pulumi:"repoKey"`
	// The project key to which the repository should be shared with.
	TargetProjectKey pulumi.StringOutput `pulumi:"targetProjectKey"`
}

// NewShareRepository registers a new resource with the given unique name, arguments, and options.
func NewShareRepository(ctx *pulumi.Context,
	name string, args *ShareRepositoryArgs, opts ...pulumi.ResourceOption) (*ShareRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepoKey == nil {
		return nil, errors.New("invalid value for required argument 'RepoKey'")
	}
	if args.TargetProjectKey == nil {
		return nil, errors.New("invalid value for required argument 'TargetProjectKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ShareRepository
	err := ctx.RegisterResource("jfrog-project:index/shareRepository:ShareRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShareRepository gets an existing ShareRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShareRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareRepositoryState, opts ...pulumi.ResourceOption) (*ShareRepository, error) {
	var resource ShareRepository
	err := ctx.ReadResource("jfrog-project:index/shareRepository:ShareRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShareRepository resources.
type shareRepositoryState struct {
	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly *bool `pulumi:"readOnly"`
	// The key of the repository.
	RepoKey *string `pulumi:"repoKey"`
	// The project key to which the repository should be shared with.
	TargetProjectKey *string `pulumi:"targetProjectKey"`
}

type ShareRepositoryState struct {
	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly pulumi.BoolPtrInput
	// The key of the repository.
	RepoKey pulumi.StringPtrInput
	// The project key to which the repository should be shared with.
	TargetProjectKey pulumi.StringPtrInput
}

func (ShareRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareRepositoryState)(nil)).Elem()
}

type shareRepositoryArgs struct {
	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly *bool `pulumi:"readOnly"`
	// The key of the repository.
	RepoKey string `pulumi:"repoKey"`
	// The project key to which the repository should be shared with.
	TargetProjectKey string `pulumi:"targetProjectKey"`
}

// The set of arguments for constructing a ShareRepository resource.
type ShareRepositoryArgs struct {
	// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
	ReadOnly pulumi.BoolPtrInput
	// The key of the repository.
	RepoKey pulumi.StringInput
	// The project key to which the repository should be shared with.
	TargetProjectKey pulumi.StringInput
}

func (ShareRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareRepositoryArgs)(nil)).Elem()
}

type ShareRepositoryInput interface {
	pulumi.Input

	ToShareRepositoryOutput() ShareRepositoryOutput
	ToShareRepositoryOutputWithContext(ctx context.Context) ShareRepositoryOutput
}

func (*ShareRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareRepository)(nil)).Elem()
}

func (i *ShareRepository) ToShareRepositoryOutput() ShareRepositoryOutput {
	return i.ToShareRepositoryOutputWithContext(context.Background())
}

func (i *ShareRepository) ToShareRepositoryOutputWithContext(ctx context.Context) ShareRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareRepositoryOutput)
}

// ShareRepositoryArrayInput is an input type that accepts ShareRepositoryArray and ShareRepositoryArrayOutput values.
// You can construct a concrete instance of `ShareRepositoryArrayInput` via:
//
//	ShareRepositoryArray{ ShareRepositoryArgs{...} }
type ShareRepositoryArrayInput interface {
	pulumi.Input

	ToShareRepositoryArrayOutput() ShareRepositoryArrayOutput
	ToShareRepositoryArrayOutputWithContext(context.Context) ShareRepositoryArrayOutput
}

type ShareRepositoryArray []ShareRepositoryInput

func (ShareRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShareRepository)(nil)).Elem()
}

func (i ShareRepositoryArray) ToShareRepositoryArrayOutput() ShareRepositoryArrayOutput {
	return i.ToShareRepositoryArrayOutputWithContext(context.Background())
}

func (i ShareRepositoryArray) ToShareRepositoryArrayOutputWithContext(ctx context.Context) ShareRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareRepositoryArrayOutput)
}

// ShareRepositoryMapInput is an input type that accepts ShareRepositoryMap and ShareRepositoryMapOutput values.
// You can construct a concrete instance of `ShareRepositoryMapInput` via:
//
//	ShareRepositoryMap{ "key": ShareRepositoryArgs{...} }
type ShareRepositoryMapInput interface {
	pulumi.Input

	ToShareRepositoryMapOutput() ShareRepositoryMapOutput
	ToShareRepositoryMapOutputWithContext(context.Context) ShareRepositoryMapOutput
}

type ShareRepositoryMap map[string]ShareRepositoryInput

func (ShareRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShareRepository)(nil)).Elem()
}

func (i ShareRepositoryMap) ToShareRepositoryMapOutput() ShareRepositoryMapOutput {
	return i.ToShareRepositoryMapOutputWithContext(context.Background())
}

func (i ShareRepositoryMap) ToShareRepositoryMapOutputWithContext(ctx context.Context) ShareRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareRepositoryMapOutput)
}

type ShareRepositoryOutput struct{ *pulumi.OutputState }

func (ShareRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareRepository)(nil)).Elem()
}

func (o ShareRepositoryOutput) ToShareRepositoryOutput() ShareRepositoryOutput {
	return o
}

func (o ShareRepositoryOutput) ToShareRepositoryOutputWithContext(ctx context.Context) ShareRepositoryOutput {
	return o
}

// Share repository with a Project in Read-Only mode to avoid any changes or modifications of the shared content.
func (o ShareRepositoryOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *ShareRepository) pulumi.BoolOutput { return v.ReadOnly }).(pulumi.BoolOutput)
}

// The key of the repository.
func (o ShareRepositoryOutput) RepoKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareRepository) pulumi.StringOutput { return v.RepoKey }).(pulumi.StringOutput)
}

// The project key to which the repository should be shared with.
func (o ShareRepositoryOutput) TargetProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareRepository) pulumi.StringOutput { return v.TargetProjectKey }).(pulumi.StringOutput)
}

type ShareRepositoryArrayOutput struct{ *pulumi.OutputState }

func (ShareRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShareRepository)(nil)).Elem()
}

func (o ShareRepositoryArrayOutput) ToShareRepositoryArrayOutput() ShareRepositoryArrayOutput {
	return o
}

func (o ShareRepositoryArrayOutput) ToShareRepositoryArrayOutputWithContext(ctx context.Context) ShareRepositoryArrayOutput {
	return o
}

func (o ShareRepositoryArrayOutput) Index(i pulumi.IntInput) ShareRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ShareRepository {
		return vs[0].([]*ShareRepository)[vs[1].(int)]
	}).(ShareRepositoryOutput)
}

type ShareRepositoryMapOutput struct{ *pulumi.OutputState }

func (ShareRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShareRepository)(nil)).Elem()
}

func (o ShareRepositoryMapOutput) ToShareRepositoryMapOutput() ShareRepositoryMapOutput {
	return o
}

func (o ShareRepositoryMapOutput) ToShareRepositoryMapOutputWithContext(ctx context.Context) ShareRepositoryMapOutput {
	return o
}

func (o ShareRepositoryMapOutput) MapIndex(k pulumi.StringInput) ShareRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ShareRepository {
		return vs[0].(map[string]*ShareRepository)[vs[1].(string)]
	}).(ShareRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShareRepositoryInput)(nil)).Elem(), &ShareRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareRepositoryArrayInput)(nil)).Elem(), ShareRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareRepositoryMapInput)(nil)).Elem(), ShareRepositoryMap{})
	pulumi.RegisterOutputType(ShareRepositoryOutput{})
	pulumi.RegisterOutputType(ShareRepositoryArrayOutput{})
	pulumi.RegisterOutputType(ShareRepositoryMapOutput{})
}
