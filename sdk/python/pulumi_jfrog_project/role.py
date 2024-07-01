# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RoleArgs', 'Role']

@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 actions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 environments: pulumi.Input[Sequence[pulumi.Input[str]]],
                 project_key: pulumi.Input[str],
                 type: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Role resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] actions: List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
        :param pulumi.Input[str] project_key: Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
        :param pulumi.Input[str] type: Type of role. Only "CUSTOM" is supported
        """
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "environments", environments)
        pulumi.set(__self__, "project_key", project_key)
        pulumi.set(__self__, "type", type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def environments(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
        """
        return pulumi.get(self, "environments")

    @environments.setter
    def environments(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "environments", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> pulumi.Input[str]:
        """
        Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of role. Only "CUSTOM" is supported
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RoleState:
    def __init__(__self__, *,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Role resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] actions: List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
        :param pulumi.Input[str] project_key: Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
        :param pulumi.Input[str] type: Type of role. Only "CUSTOM" is supported
        """
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if environments is not None:
            pulumi.set(__self__, "environments", environments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def environments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
        """
        return pulumi.get(self, "environments")

    @environments.setter
    def environments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "environments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        """
        Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of role. Only "CUSTOM" is supported
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Role(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a project role. Element has one to one mapping with the [JFrog Project Roles API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-AddaNewRole). Requires a user assigned with the 'Administer the Platform' role or Project Admin permissions if `admin_privileges.manage_resoures` is enabled.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_jfrog_project as jfrog_project

        myrole = jfrog_project.Role("myrole",
            type="CUSTOM",
            project_key=project["myproject"]["key"],
            environments=["DEV"],
            actions=[
                "READ_REPOSITORY",
                "ANNOTATE_REPOSITORY",
            ])
        ```

        ## Import

        ```sh
        $ pulumi import jfrog-project:index/role:Role myrole project_key:role_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] actions: List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
        :param pulumi.Input[str] project_key: Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
        :param pulumi.Input[str] type: Type of role. Only "CUSTOM" is supported
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a project role. Element has one to one mapping with the [JFrog Project Roles API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-AddaNewRole). Requires a user assigned with the 'Administer the Platform' role or Project Admin permissions if `admin_privileges.manage_resoures` is enabled.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_jfrog_project as jfrog_project

        myrole = jfrog_project.Role("myrole",
            type="CUSTOM",
            project_key=project["myproject"]["key"],
            environments=["DEV"],
            actions=[
                "READ_REPOSITORY",
                "ANNOTATE_REPOSITORY",
            ])
        ```

        ## Import

        ```sh
        $ pulumi import jfrog-project:index/role:Role myrole project_key:role_name
        ```

        :param str resource_name: The name of the resource.
        :param RoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleArgs.__new__(RoleArgs)

            if actions is None and not opts.urn:
                raise TypeError("Missing required property 'actions'")
            __props__.__dict__["actions"] = actions
            if environments is None and not opts.urn:
                raise TypeError("Missing required property 'environments'")
            __props__.__dict__["environments"] = environments
            __props__.__dict__["name"] = name
            if project_key is None and not opts.urn:
                raise TypeError("Missing required property 'project_key'")
            __props__.__dict__["project_key"] = project_key
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(Role, __self__).__init__(
            'jfrog-project:index/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            environments: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_key: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] actions: List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] environments: A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
        :param pulumi.Input[str] project_key: Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
        :param pulumi.Input[str] type: Type of role. Only "CUSTOM" is supported
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleState.__new__(_RoleState)

        __props__.__dict__["actions"] = actions
        __props__.__dict__["environments"] = environments
        __props__.__dict__["name"] = name
        __props__.__dict__["project_key"] = project_key
        __props__.__dict__["type"] = type
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Sequence[str]]:
        """
        List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def environments(self) -> pulumi.Output[Sequence[str]]:
        """
        A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
        """
        return pulumi.get(self, "environments")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> pulumi.Output[str]:
        """
        Project key for this environment. This field supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter.
        """
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of role. Only "CUSTOM" is supported
        """
        return pulumi.get(self, "type")

