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

__all__ = ['IpaDomainArgs', 'IpaDomain']

@pulumi.input_type
class IpaDomainArgs:
    def __init__(__self__, *,
                 domain_name: pulumi.Input[str],
                 sources: pulumi.Input[Sequence[pulumi.Input['IpaDomainSourceArgs']]],
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IpaDomain resource.
        :param pulumi.Input[str] domain_name: The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        :param pulumi.Input[Sequence[pulumi.Input['IpaDomainSourceArgs']]] sources: Sources. See `sources` below.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        :param pulumi.Input[str] scope: The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        :param pulumi.Input[str] status: The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        """
        pulumi.set(__self__, "domain_name", domain_name)
        pulumi.set(__self__, "sources", sources)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Input[Sequence[pulumi.Input['IpaDomainSourceArgs']]]:
        """
        Sources. See `sources` below.
        """
        return pulumi.get(self, "sources")

    @sources.setter
    def sources(self, value: pulumi.Input[Sequence[pulumi.Input['IpaDomainSourceArgs']]]):
        pulumi.set(self, "sources", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _IpaDomainState:
    def __init__(__self__, *,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input['IpaDomainSourceArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpaDomain resources.
        :param pulumi.Input[str] domain_name: The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        :param pulumi.Input[str] resource_group_id: The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        :param pulumi.Input[str] scope: The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        :param pulumi.Input[Sequence[pulumi.Input['IpaDomainSourceArgs']]] sources: Sources. See `sources` below.
        :param pulumi.Input[str] status: The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        """
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if sources is not None:
            pulumi.set(__self__, "sources", sources)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpaDomainSourceArgs']]]]:
        """
        Sources. See `sources` below.
        """
        return pulumi.get(self, "sources")

    @sources.setter
    def sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpaDomainSourceArgs']]]]):
        pulumi.set(self, "sources", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class IpaDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpaDomainSourceArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a DCDN Ipa Domain resource.

        For information about DCDN Ipa Domain and how to use it, see [What is Ipa Domain](https://www.alibabacloud.com/help/en/doc-detail/130634.html).

        > **NOTE:** Available since v1.158.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        domain_name = config.get("domainName")
        if domain_name is None:
            domain_name = "example.com"
        default = alicloud.resourcemanager.get_resource_groups()
        example = alicloud.dcdn.IpaDomain("example",
            domain_name=domain_name,
            resource_group_id=default.groups[0].id,
            scope="global",
            status="online",
            sources=[alicloud.dcdn.IpaDomainSourceArgs(
                content="www.alicloud-provider.cn",
                port=80,
                priority="20",
                type="domain",
                weight=10,
            )])
        ```

        ## Import

        DCDN Ipa Domain can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dcdn/ipaDomain:IpaDomain example <domain_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        :param pulumi.Input[str] resource_group_id: The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        :param pulumi.Input[str] scope: The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpaDomainSourceArgs']]]] sources: Sources. See `sources` below.
        :param pulumi.Input[str] status: The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpaDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DCDN Ipa Domain resource.

        For information about DCDN Ipa Domain and how to use it, see [What is Ipa Domain](https://www.alibabacloud.com/help/en/doc-detail/130634.html).

        > **NOTE:** Available since v1.158.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        domain_name = config.get("domainName")
        if domain_name is None:
            domain_name = "example.com"
        default = alicloud.resourcemanager.get_resource_groups()
        example = alicloud.dcdn.IpaDomain("example",
            domain_name=domain_name,
            resource_group_id=default.groups[0].id,
            scope="global",
            status="online",
            sources=[alicloud.dcdn.IpaDomainSourceArgs(
                content="www.alicloud-provider.cn",
                port=80,
                priority="20",
                type="domain",
                weight=10,
            )])
        ```

        ## Import

        DCDN Ipa Domain can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dcdn/ipaDomain:IpaDomain example <domain_name>
        ```

        :param str resource_name: The name of the resource.
        :param IpaDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpaDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpaDomainSourceArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpaDomainArgs.__new__(IpaDomainArgs)

            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["scope"] = scope
            if sources is None and not opts.urn:
                raise TypeError("Missing required property 'sources'")
            __props__.__dict__["sources"] = sources
            __props__.__dict__["status"] = status
        super(IpaDomain, __self__).__init__(
            'alicloud:dcdn/ipaDomain:IpaDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpaDomainSourceArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'IpaDomain':
        """
        Get an existing IpaDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        :param pulumi.Input[str] resource_group_id: The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        :param pulumi.Input[str] scope: The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpaDomainSourceArgs']]]] sources: Sources. See `sources` below.
        :param pulumi.Input[str] status: The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpaDomainState.__new__(_IpaDomainState)

        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["scope"] = scope
        __props__.__dict__["sources"] = sources
        __props__.__dict__["status"] = status
        return IpaDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        The accelerated region. Valid values: `domestic`, `global`, `overseas`.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Output[Sequence['outputs.IpaDomainSource']]:
        """
        Sources. See `sources` below.
        """
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
        """
        return pulumi.get(self, "status")

