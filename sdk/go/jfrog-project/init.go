// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package jfrogproject

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/skycaptain/pulumi-jfrog-project/sdk/go/jfrog-project/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "jfrog-project:index/environment:Environment":
		r = &Environment{}
	case "jfrog-project:index/group:Group":
		r = &Group{}
	case "jfrog-project:index/project:Project":
		r = &Project{}
	case "jfrog-project:index/repository:Repository":
		r = &Repository{}
	case "jfrog-project:index/role:Role":
		r = &Role{}
	case "jfrog-project:index/user:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:jfrog-project" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"jfrog-project",
		"index/environment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"jfrog-project",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"jfrog-project",
		"index/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"jfrog-project",
		"index/repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"jfrog-project",
		"index/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"jfrog-project",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"jfrog-project",
		&pkg{version},
	)
}