# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 access_token: Optional[pulumi.Input[str]] = None,
                 check_license: Optional[pulumi.Input[bool]] = None,
                 oidc_provider_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] access_token: This is a Bearer token that can be given to you by your admin under `Identity and Access`. This can also be sourced from
               the `PROJECT_ACCESS_TOKEN` or `JFROG_ACCESS_TOKEN` environment variable. Defauult to empty string if not set.
        :param pulumi.Input[bool] check_license: Toggle for pre-flight checking of Artifactory Enterprise license. Default to `true`.
        :param pulumi.Input[str] oidc_provider_name: OIDC provider name. See [Configure an OIDC
               Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
               more details.
        :param pulumi.Input[str] url: URL of Artifactory. This can also be sourced from the `PROJECT_URL` or `JFROG_URL` environment variable. Default to
               'http://localhost:8081' if not set.
        """
        if access_token is None:
            access_token = _utilities.get_env('PROJECT_ACCESS_TOKEN', 'JFROG_ACCESS_TOKEN')
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if check_license is None:
            check_license = False
        if check_license is not None:
            pulumi.set(__self__, "check_license", check_license)
        if oidc_provider_name is not None:
            pulumi.set(__self__, "oidc_provider_name", oidc_provider_name)
        if url is None:
            url = (_utilities.get_env('PROJECT_URL', 'JFROG_URL', 'JFROG_PLATFORM_URL') or 'http://localhost:8081')
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[pulumi.Input[str]]:
        """
        This is a Bearer token that can be given to you by your admin under `Identity and Access`. This can also be sourced from
        the `PROJECT_ACCESS_TOKEN` or `JFROG_ACCESS_TOKEN` environment variable. Defauult to empty string if not set.
        """
        return pulumi.get(self, "access_token")

    @access_token.setter
    def access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token", value)

    @property
    @pulumi.getter(name="checkLicense")
    def check_license(self) -> Optional[pulumi.Input[bool]]:
        """
        Toggle for pre-flight checking of Artifactory Enterprise license. Default to `true`.
        """
        return pulumi.get(self, "check_license")

    @check_license.setter
    def check_license(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "check_license", value)

    @property
    @pulumi.getter(name="oidcProviderName")
    def oidc_provider_name(self) -> Optional[pulumi.Input[str]]:
        """
        OIDC provider name. See [Configure an OIDC
        Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
        more details.
        """
        return pulumi.get(self, "oidc_provider_name")

    @oidc_provider_name.setter
    def oidc_provider_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oidc_provider_name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL of Artifactory. This can also be sourced from the `PROJECT_URL` or `JFROG_URL` environment variable. Default to
        'http://localhost:8081' if not set.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token: Optional[pulumi.Input[str]] = None,
                 check_license: Optional[pulumi.Input[bool]] = None,
                 oidc_provider_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the project package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token: This is a Bearer token that can be given to you by your admin under `Identity and Access`. This can also be sourced from
               the `PROJECT_ACCESS_TOKEN` or `JFROG_ACCESS_TOKEN` environment variable. Defauult to empty string if not set.
        :param pulumi.Input[bool] check_license: Toggle for pre-flight checking of Artifactory Enterprise license. Default to `true`.
        :param pulumi.Input[str] oidc_provider_name: OIDC provider name. See [Configure an OIDC
               Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
               more details.
        :param pulumi.Input[str] url: URL of Artifactory. This can also be sourced from the `PROJECT_URL` or `JFROG_URL` environment variable. Default to
               'http://localhost:8081' if not set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the project package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token: Optional[pulumi.Input[str]] = None,
                 check_license: Optional[pulumi.Input[bool]] = None,
                 oidc_provider_name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if access_token is None:
                access_token = _utilities.get_env('PROJECT_ACCESS_TOKEN', 'JFROG_ACCESS_TOKEN')
            __props__.__dict__["access_token"] = None if access_token is None else pulumi.Output.secret(access_token)
            if check_license is None:
                check_license = False
            __props__.__dict__["check_license"] = pulumi.Output.from_input(check_license).apply(pulumi.runtime.to_json) if check_license is not None else None
            __props__.__dict__["oidc_provider_name"] = oidc_provider_name
            if url is None:
                url = (_utilities.get_env('PROJECT_URL', 'JFROG_URL', 'JFROG_PLATFORM_URL') or 'http://localhost:8081')
            __props__.__dict__["url"] = url
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accessToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'jfrog-project',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> pulumi.Output[Optional[str]]:
        """
        This is a Bearer token that can be given to you by your admin under `Identity and Access`. This can also be sourced from
        the `PROJECT_ACCESS_TOKEN` or `JFROG_ACCESS_TOKEN` environment variable. Defauult to empty string if not set.
        """
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter(name="oidcProviderName")
    def oidc_provider_name(self) -> pulumi.Output[Optional[str]]:
        """
        OIDC provider name. See [Configure an OIDC
        Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
        more details.
        """
        return pulumi.get(self, "oidc_provider_name")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        URL of Artifactory. This can also be sourced from the `PROJECT_URL` or `JFROG_URL` environment variable. Default to
        'http://localhost:8081' if not set.
        """
        return pulumi.get(self, "url")

