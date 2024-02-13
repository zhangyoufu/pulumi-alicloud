# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DirectoryArgs', 'Directory']

@pulumi.input_type
class DirectoryArgs:
    def __init__(__self__, *,
                 directory_name: Optional[pulumi.Input[str]] = None,
                 mfa_authentication_status: Optional[pulumi.Input[str]] = None,
                 saml_identity_provider_configuration: Optional[pulumi.Input['DirectorySamlIdentityProviderConfigurationArgs']] = None,
                 scim_synchronization_status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Directory resource.
        :param pulumi.Input[str] directory_name: The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        :param pulumi.Input[str] mfa_authentication_status: The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        :param pulumi.Input['DirectorySamlIdentityProviderConfigurationArgs'] saml_identity_provider_configuration: The saml identity provider configuration. See `saml_identity_provider_configuration` below.
               
               > **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        :param pulumi.Input[str] scim_synchronization_status: The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        """
        if directory_name is not None:
            pulumi.set(__self__, "directory_name", directory_name)
        if mfa_authentication_status is not None:
            pulumi.set(__self__, "mfa_authentication_status", mfa_authentication_status)
        if saml_identity_provider_configuration is not None:
            pulumi.set(__self__, "saml_identity_provider_configuration", saml_identity_provider_configuration)
        if scim_synchronization_status is not None:
            pulumi.set(__self__, "scim_synchronization_status", scim_synchronization_status)

    @property
    @pulumi.getter(name="directoryName")
    def directory_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        """
        return pulumi.get(self, "directory_name")

    @directory_name.setter
    def directory_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_name", value)

    @property
    @pulumi.getter(name="mfaAuthenticationStatus")
    def mfa_authentication_status(self) -> Optional[pulumi.Input[str]]:
        """
        The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        """
        return pulumi.get(self, "mfa_authentication_status")

    @mfa_authentication_status.setter
    def mfa_authentication_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_authentication_status", value)

    @property
    @pulumi.getter(name="samlIdentityProviderConfiguration")
    def saml_identity_provider_configuration(self) -> Optional[pulumi.Input['DirectorySamlIdentityProviderConfigurationArgs']]:
        """
        The saml identity provider configuration. See `saml_identity_provider_configuration` below.

        > **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        """
        return pulumi.get(self, "saml_identity_provider_configuration")

    @saml_identity_provider_configuration.setter
    def saml_identity_provider_configuration(self, value: Optional[pulumi.Input['DirectorySamlIdentityProviderConfigurationArgs']]):
        pulumi.set(self, "saml_identity_provider_configuration", value)

    @property
    @pulumi.getter(name="scimSynchronizationStatus")
    def scim_synchronization_status(self) -> Optional[pulumi.Input[str]]:
        """
        The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        """
        return pulumi.get(self, "scim_synchronization_status")

    @scim_synchronization_status.setter
    def scim_synchronization_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scim_synchronization_status", value)


@pulumi.input_type
class _DirectoryState:
    def __init__(__self__, *,
                 directory_name: Optional[pulumi.Input[str]] = None,
                 mfa_authentication_status: Optional[pulumi.Input[str]] = None,
                 saml_identity_provider_configuration: Optional[pulumi.Input['DirectorySamlIdentityProviderConfigurationArgs']] = None,
                 scim_synchronization_status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Directory resources.
        :param pulumi.Input[str] directory_name: The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        :param pulumi.Input[str] mfa_authentication_status: The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        :param pulumi.Input['DirectorySamlIdentityProviderConfigurationArgs'] saml_identity_provider_configuration: The saml identity provider configuration. See `saml_identity_provider_configuration` below.
               
               > **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        :param pulumi.Input[str] scim_synchronization_status: The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        """
        if directory_name is not None:
            pulumi.set(__self__, "directory_name", directory_name)
        if mfa_authentication_status is not None:
            pulumi.set(__self__, "mfa_authentication_status", mfa_authentication_status)
        if saml_identity_provider_configuration is not None:
            pulumi.set(__self__, "saml_identity_provider_configuration", saml_identity_provider_configuration)
        if scim_synchronization_status is not None:
            pulumi.set(__self__, "scim_synchronization_status", scim_synchronization_status)

    @property
    @pulumi.getter(name="directoryName")
    def directory_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        """
        return pulumi.get(self, "directory_name")

    @directory_name.setter
    def directory_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_name", value)

    @property
    @pulumi.getter(name="mfaAuthenticationStatus")
    def mfa_authentication_status(self) -> Optional[pulumi.Input[str]]:
        """
        The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        """
        return pulumi.get(self, "mfa_authentication_status")

    @mfa_authentication_status.setter
    def mfa_authentication_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_authentication_status", value)

    @property
    @pulumi.getter(name="samlIdentityProviderConfiguration")
    def saml_identity_provider_configuration(self) -> Optional[pulumi.Input['DirectorySamlIdentityProviderConfigurationArgs']]:
        """
        The saml identity provider configuration. See `saml_identity_provider_configuration` below.

        > **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        """
        return pulumi.get(self, "saml_identity_provider_configuration")

    @saml_identity_provider_configuration.setter
    def saml_identity_provider_configuration(self, value: Optional[pulumi.Input['DirectorySamlIdentityProviderConfigurationArgs']]):
        pulumi.set(self, "saml_identity_provider_configuration", value)

    @property
    @pulumi.getter(name="scimSynchronizationStatus")
    def scim_synchronization_status(self) -> Optional[pulumi.Input[str]]:
        """
        The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        """
        return pulumi.get(self, "scim_synchronization_status")

    @scim_synchronization_status.setter
    def scim_synchronization_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scim_synchronization_status", value)


class Directory(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 directory_name: Optional[pulumi.Input[str]] = None,
                 mfa_authentication_status: Optional[pulumi.Input[str]] = None,
                 saml_identity_provider_configuration: Optional[pulumi.Input[pulumi.InputType['DirectorySamlIdentityProviderConfigurationArgs']]] = None,
                 scim_synchronization_status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud SSO Directory resource.

        For information about Cloud SSO Directory and how to use it, see [What is Directory](https://www.alibabacloud.com/help/en/cloudsso/latest/api-cloudsso-2021-05-15-createdirectory).

        > **NOTE:** Available since v1.135.0.

        > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region

        ## Import

        Cloud SSO Directory can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cloudsso/directory:Directory example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] directory_name: The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        :param pulumi.Input[str] mfa_authentication_status: The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        :param pulumi.Input[pulumi.InputType['DirectorySamlIdentityProviderConfigurationArgs']] saml_identity_provider_configuration: The saml identity provider configuration. See `saml_identity_provider_configuration` below.
               
               > **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        :param pulumi.Input[str] scim_synchronization_status: The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DirectoryArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud SSO Directory resource.

        For information about Cloud SSO Directory and how to use it, see [What is Directory](https://www.alibabacloud.com/help/en/cloudsso/latest/api-cloudsso-2021-05-15-createdirectory).

        > **NOTE:** Available since v1.135.0.

        > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region

        ## Import

        Cloud SSO Directory can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cloudsso/directory:Directory example <id>
        ```

        :param str resource_name: The name of the resource.
        :param DirectoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DirectoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 directory_name: Optional[pulumi.Input[str]] = None,
                 mfa_authentication_status: Optional[pulumi.Input[str]] = None,
                 saml_identity_provider_configuration: Optional[pulumi.Input[pulumi.InputType['DirectorySamlIdentityProviderConfigurationArgs']]] = None,
                 scim_synchronization_status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DirectoryArgs.__new__(DirectoryArgs)

            __props__.__dict__["directory_name"] = directory_name
            __props__.__dict__["mfa_authentication_status"] = mfa_authentication_status
            __props__.__dict__["saml_identity_provider_configuration"] = saml_identity_provider_configuration
            __props__.__dict__["scim_synchronization_status"] = scim_synchronization_status
        super(Directory, __self__).__init__(
            'alicloud:cloudsso/directory:Directory',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            directory_name: Optional[pulumi.Input[str]] = None,
            mfa_authentication_status: Optional[pulumi.Input[str]] = None,
            saml_identity_provider_configuration: Optional[pulumi.Input[pulumi.InputType['DirectorySamlIdentityProviderConfigurationArgs']]] = None,
            scim_synchronization_status: Optional[pulumi.Input[str]] = None) -> 'Directory':
        """
        Get an existing Directory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] directory_name: The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        :param pulumi.Input[str] mfa_authentication_status: The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        :param pulumi.Input[pulumi.InputType['DirectorySamlIdentityProviderConfigurationArgs']] saml_identity_provider_configuration: The saml identity provider configuration. See `saml_identity_provider_configuration` below.
               
               > **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        :param pulumi.Input[str] scim_synchronization_status: The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DirectoryState.__new__(_DirectoryState)

        __props__.__dict__["directory_name"] = directory_name
        __props__.__dict__["mfa_authentication_status"] = mfa_authentication_status
        __props__.__dict__["saml_identity_provider_configuration"] = saml_identity_provider_configuration
        __props__.__dict__["scim_synchronization_status"] = scim_synchronization_status
        return Directory(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="directoryName")
    def directory_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        """
        return pulumi.get(self, "directory_name")

    @property
    @pulumi.getter(name="mfaAuthenticationStatus")
    def mfa_authentication_status(self) -> pulumi.Output[str]:
        """
        The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        """
        return pulumi.get(self, "mfa_authentication_status")

    @property
    @pulumi.getter(name="samlIdentityProviderConfiguration")
    def saml_identity_provider_configuration(self) -> pulumi.Output['outputs.DirectorySamlIdentityProviderConfiguration']:
        """
        The saml identity provider configuration. See `saml_identity_provider_configuration` below.

        > **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        """
        return pulumi.get(self, "saml_identity_provider_configuration")

    @property
    @pulumi.getter(name="scimSynchronizationStatus")
    def scim_synchronization_status(self) -> pulumi.Output[str]:
        """
        The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        """
        return pulumi.get(self, "scim_synchronization_status")

