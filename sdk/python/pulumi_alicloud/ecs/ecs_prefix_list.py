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

__all__ = ['EcsPrefixListArgs', 'EcsPrefixList']

@pulumi.input_type
class EcsPrefixListArgs:
    def __init__(__self__, *,
                 address_family: pulumi.Input[str],
                 entries: pulumi.Input[Sequence[pulumi.Input['EcsPrefixListEntryArgs']]],
                 max_entries: pulumi.Input[int],
                 prefix_list_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EcsPrefixList resource.
        :param pulumi.Input[str] address_family: The IP address family. Valid values: `IPv4`,`IPv6`.
        :param pulumi.Input[Sequence[pulumi.Input['EcsPrefixListEntryArgs']]] entries: The Entry. The details see Block `entry`.
        :param pulumi.Input[int] max_entries: The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
        :param pulumi.Input[str] prefix_list_name: The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        :param pulumi.Input[str] description: The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        """
        pulumi.set(__self__, "address_family", address_family)
        pulumi.set(__self__, "entries", entries)
        pulumi.set(__self__, "max_entries", max_entries)
        pulumi.set(__self__, "prefix_list_name", prefix_list_name)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Input[str]:
        """
        The IP address family. Valid values: `IPv4`,`IPv6`.
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: pulumi.Input[str]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter
    def entries(self) -> pulumi.Input[Sequence[pulumi.Input['EcsPrefixListEntryArgs']]]:
        """
        The Entry. The details see Block `entry`.
        """
        return pulumi.get(self, "entries")

    @entries.setter
    def entries(self, value: pulumi.Input[Sequence[pulumi.Input['EcsPrefixListEntryArgs']]]):
        pulumi.set(self, "entries", value)

    @property
    @pulumi.getter(name="maxEntries")
    def max_entries(self) -> pulumi.Input[int]:
        """
        The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
        """
        return pulumi.get(self, "max_entries")

    @max_entries.setter
    def max_entries(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_entries", value)

    @property
    @pulumi.getter(name="prefixListName")
    def prefix_list_name(self) -> pulumi.Input[str]:
        """
        The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        """
        return pulumi.get(self, "prefix_list_name")

    @prefix_list_name.setter
    def prefix_list_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "prefix_list_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _EcsPrefixListState:
    def __init__(__self__, *,
                 address_family: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input['EcsPrefixListEntryArgs']]]] = None,
                 max_entries: Optional[pulumi.Input[int]] = None,
                 prefix_list_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EcsPrefixList resources.
        :param pulumi.Input[str] address_family: The IP address family. Valid values: `IPv4`,`IPv6`.
        :param pulumi.Input[str] description: The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        :param pulumi.Input[Sequence[pulumi.Input['EcsPrefixListEntryArgs']]] entries: The Entry. The details see Block `entry`.
        :param pulumi.Input[int] max_entries: The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
        :param pulumi.Input[str] prefix_list_name: The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        """
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if entries is not None:
            pulumi.set(__self__, "entries", entries)
        if max_entries is not None:
            pulumi.set(__self__, "max_entries", max_entries)
        if prefix_list_name is not None:
            pulumi.set(__self__, "prefix_list_name", prefix_list_name)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address family. Valid values: `IPv4`,`IPv6`.
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EcsPrefixListEntryArgs']]]]:
        """
        The Entry. The details see Block `entry`.
        """
        return pulumi.get(self, "entries")

    @entries.setter
    def entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EcsPrefixListEntryArgs']]]]):
        pulumi.set(self, "entries", value)

    @property
    @pulumi.getter(name="maxEntries")
    def max_entries(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
        """
        return pulumi.get(self, "max_entries")

    @max_entries.setter
    def max_entries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_entries", value)

    @property
    @pulumi.getter(name="prefixListName")
    def prefix_list_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        """
        return pulumi.get(self, "prefix_list_name")

    @prefix_list_name.setter
    def prefix_list_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix_list_name", value)


class EcsPrefixList(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EcsPrefixListEntryArgs']]]]] = None,
                 max_entries: Optional[pulumi.Input[int]] = None,
                 prefix_list_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ECS Prefix List resource.

        For information about ECS Prefix List and how to use it, see [What is Prefix List.](https://www.alibabacloud.com/help/en/doc-detail/207969.html).

        > **NOTE:** Available in v1.152.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.ecs.EcsPrefixList("default",
            address_family="IPv4",
            max_entries=2,
            prefix_list_name="tftest",
            description="description",
            entries=[alicloud.ecs.EcsPrefixListEntryArgs(
                cidr="192.168.0.0/24",
                description="description",
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ECS Prefix List can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsPrefixList:EcsPrefixList example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: The IP address family. Valid values: `IPv4`,`IPv6`.
        :param pulumi.Input[str] description: The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EcsPrefixListEntryArgs']]]] entries: The Entry. The details see Block `entry`.
        :param pulumi.Input[int] max_entries: The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
        :param pulumi.Input[str] prefix_list_name: The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EcsPrefixListArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ECS Prefix List resource.

        For information about ECS Prefix List and how to use it, see [What is Prefix List.](https://www.alibabacloud.com/help/en/doc-detail/207969.html).

        > **NOTE:** Available in v1.152.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.ecs.EcsPrefixList("default",
            address_family="IPv4",
            max_entries=2,
            prefix_list_name="tftest",
            description="description",
            entries=[alicloud.ecs.EcsPrefixListEntryArgs(
                cidr="192.168.0.0/24",
                description="description",
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        ECS Prefix List can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsPrefixList:EcsPrefixList example <id>
        ```

        :param str resource_name: The name of the resource.
        :param EcsPrefixListArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EcsPrefixListArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EcsPrefixListEntryArgs']]]]] = None,
                 max_entries: Optional[pulumi.Input[int]] = None,
                 prefix_list_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EcsPrefixListArgs.__new__(EcsPrefixListArgs)

            if address_family is None and not opts.urn:
                raise TypeError("Missing required property 'address_family'")
            __props__.__dict__["address_family"] = address_family
            __props__.__dict__["description"] = description
            if entries is None and not opts.urn:
                raise TypeError("Missing required property 'entries'")
            __props__.__dict__["entries"] = entries
            if max_entries is None and not opts.urn:
                raise TypeError("Missing required property 'max_entries'")
            __props__.__dict__["max_entries"] = max_entries
            if prefix_list_name is None and not opts.urn:
                raise TypeError("Missing required property 'prefix_list_name'")
            __props__.__dict__["prefix_list_name"] = prefix_list_name
        super(EcsPrefixList, __self__).__init__(
            'alicloud:ecs/ecsPrefixList:EcsPrefixList',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_family: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EcsPrefixListEntryArgs']]]]] = None,
            max_entries: Optional[pulumi.Input[int]] = None,
            prefix_list_name: Optional[pulumi.Input[str]] = None) -> 'EcsPrefixList':
        """
        Get an existing EcsPrefixList resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: The IP address family. Valid values: `IPv4`,`IPv6`.
        :param pulumi.Input[str] description: The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EcsPrefixListEntryArgs']]]] entries: The Entry. The details see Block `entry`.
        :param pulumi.Input[int] max_entries: The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
        :param pulumi.Input[str] prefix_list_name: The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EcsPrefixListState.__new__(_EcsPrefixListState)

        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["description"] = description
        __props__.__dict__["entries"] = entries
        __props__.__dict__["max_entries"] = max_entries
        __props__.__dict__["prefix_list_name"] = prefix_list_name
        return EcsPrefixList(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Output[str]:
        """
        The IP address family. Valid values: `IPv4`,`IPv6`.
        """
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def entries(self) -> pulumi.Output[Sequence['outputs.EcsPrefixListEntry']]:
        """
        The Entry. The details see Block `entry`.
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter(name="maxEntries")
    def max_entries(self) -> pulumi.Output[int]:
        """
        The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
        """
        return pulumi.get(self, "max_entries")

    @property
    @pulumi.getter(name="prefixListName")
    def prefix_list_name(self) -> pulumi.Output[str]:
        """
        The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
        """
        return pulumi.get(self, "prefix_list_name")

