# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'ProjectAdminPrivilege',
    'ProjectGroup',
    'ProjectMember',
    'ProjectRole',
]

@pulumi.output_type
class ProjectAdminPrivilege(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "indexResources":
            suggest = "index_resources"
        elif key == "manageMembers":
            suggest = "manage_members"
        elif key == "manageResources":
            suggest = "manage_resources"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectAdminPrivilege. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectAdminPrivilege.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectAdminPrivilege.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 index_resources: bool,
                 manage_members: bool,
                 manage_resources: bool):
        """
        :param bool index_resources: Enables a project admin to define the resources to be indexed by Xray
        :param bool manage_members: Allows the Project Admin to manage Platform users/groups as project members with different roles.
        :param bool manage_resources: Allows the Project Admin to manage resources - repositories, builds and Pipelines resources on the project level.
        """
        pulumi.set(__self__, "index_resources", index_resources)
        pulumi.set(__self__, "manage_members", manage_members)
        pulumi.set(__self__, "manage_resources", manage_resources)

    @property
    @pulumi.getter(name="indexResources")
    def index_resources(self) -> bool:
        """
        Enables a project admin to define the resources to be indexed by Xray
        """
        return pulumi.get(self, "index_resources")

    @property
    @pulumi.getter(name="manageMembers")
    def manage_members(self) -> bool:
        """
        Allows the Project Admin to manage Platform users/groups as project members with different roles.
        """
        return pulumi.get(self, "manage_members")

    @property
    @pulumi.getter(name="manageResources")
    def manage_resources(self) -> bool:
        """
        Allows the Project Admin to manage resources - repositories, builds and Pipelines resources on the project level.
        """
        return pulumi.get(self, "manage_resources")


@pulumi.output_type
class ProjectGroup(dict):
    def __init__(__self__, *,
                 name: str,
                 roles: Sequence[str]):
        """
        :param str name: Must be existing Artifactory group
        :param Sequence[str] roles: List of pre-defined Project or custom roles
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Must be existing Artifactory group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def roles(self) -> Sequence[str]:
        """
        List of pre-defined Project or custom roles
        """
        return pulumi.get(self, "roles")


@pulumi.output_type
class ProjectMember(dict):
    def __init__(__self__, *,
                 name: str,
                 roles: Sequence[str]):
        """
        :param str name: Must be existing Artifactory user
        :param Sequence[str] roles: List of pre-defined Project or custom roles
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Must be existing Artifactory user
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def roles(self) -> Sequence[str]:
        """
        List of pre-defined Project or custom roles
        """
        return pulumi.get(self, "roles")


@pulumi.output_type
class ProjectRole(dict):
    def __init__(__self__, *,
                 actions: Sequence[str],
                 environments: Sequence[str],
                 name: str,
                 description: Optional[str] = None,
                 type: Optional[str] = None):
        """
        :param Sequence[str] actions: List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
        :param Sequence[str] environments: A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
        :param str type: Type of role. Only "CUSTOM" is supported
        """
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "environments", environments)
        pulumi.set(__self__, "name", name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def actions(self) -> Sequence[str]:
        """
        List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def environments(self) -> Sequence[str]:
        """
        A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
        """
        return pulumi.get(self, "environments")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Type of role. Only "CUSTOM" is supported
        """
        return pulumi.get(self, "type")


