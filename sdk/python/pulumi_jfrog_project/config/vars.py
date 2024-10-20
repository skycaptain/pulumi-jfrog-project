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
from .. import _utilities

import types

__config__ = pulumi.Config('jfrog-project')


class _ExportableConfig(types.ModuleType):
    @property
    def access_token(self) -> Optional[str]:
        """
        This is a Bearer token that can be given to you by your admin under `Identity and Access`. This can also be sourced from
        the `PROJECT_ACCESS_TOKEN` or `JFROG_ACCESS_TOKEN` environment variable. Defauult to empty string if not set.
        """
        return __config__.get('accessToken') or _utilities.get_env('PROJECT_ACCESS_TOKEN', 'JFROG_ACCESS_TOKEN')

    @property
    def check_license(self) -> bool:
        """
        Toggle for pre-flight checking of Artifactory Enterprise license. Default to `true`.
        """
        return __config__.get_bool('checkLicense') or False

    @property
    def oidc_provider_name(self) -> Optional[str]:
        """
        OIDC provider name. See [Configure an OIDC
        Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
        more details.
        """
        return __config__.get('oidcProviderName')

    @property
    def tfc_credential_tag_name(self) -> Optional[str]:
        return __config__.get('tfcCredentialTagName')

    @property
    def url(self) -> str:
        """
        URL of Artifactory. This can also be sourced from the `PROJECT_URL` or `JFROG_URL` environment variable. Default to
        'http://localhost:8081' if not set.
        """
        return __config__.get('url') or (_utilities.get_env('PROJECT_URL', 'JFROG_URL', 'JFROG_PLATFORM_URL') or 'http://localhost:8081')

