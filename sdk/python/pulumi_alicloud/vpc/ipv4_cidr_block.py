# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['Ipv4CidrBlockArgs', 'Ipv4CidrBlock']

@pulumi.input_type
class Ipv4CidrBlockArgs:
    def __init__(__self__, *,
                 secondary_cidr_block: pulumi.Input[str],
                 vpc_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a Ipv4CidrBlock resource.
        :param pulumi.Input[str] secondary_cidr_block: The IPv4 CIDR block. Take note of the following requirements:
               * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
               * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
               * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
               * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        pulumi.set(__self__, "secondary_cidr_block", secondary_cidr_block)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="secondaryCidrBlock")
    def secondary_cidr_block(self) -> pulumi.Input[str]:
        """
        The IPv4 CIDR block. Take note of the following requirements:
        * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
        * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
        * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
        * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        """
        return pulumi.get(self, "secondary_cidr_block")

    @secondary_cidr_block.setter
    def secondary_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "secondary_cidr_block", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _Ipv4CidrBlockState:
    def __init__(__self__, *,
                 secondary_cidr_block: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ipv4CidrBlock resources.
        :param pulumi.Input[str] secondary_cidr_block: The IPv4 CIDR block. Take note of the following requirements:
               * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
               * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
               * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
               * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        if secondary_cidr_block is not None:
            pulumi.set(__self__, "secondary_cidr_block", secondary_cidr_block)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="secondaryCidrBlock")
    def secondary_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 CIDR block. Take note of the following requirements:
        * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
        * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
        * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
        * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        """
        return pulumi.get(self, "secondary_cidr_block")

    @secondary_cidr_block.setter
    def secondary_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_cidr_block", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class Ipv4CidrBlock(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 secondary_cidr_block: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VPC Ipv4 Cidr Block resource.

        For information about VPC Ipv4 Cidr Block and how to use it, see [What is Ipv4 Cidr Block](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/associatevpccidrblock).

        > **NOTE:** Available in v1.185.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.vpc.Network("default",
            cidr_block="192.168.0.0/24",
            vpc_name=var["name"])
        example = alicloud.vpc.Ipv4CidrBlock("example",
            vpc_id=default.id,
            secondary_cidr_block="192.163.0.0/16")
        ```

        ## Import

        VPC Ipv4 Cidr Block can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock example <vpc_id>:<secondary_cidr_block>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] secondary_cidr_block: The IPv4 CIDR block. Take note of the following requirements:
               * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
               * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
               * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
               * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ipv4CidrBlockArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Ipv4 Cidr Block resource.

        For information about VPC Ipv4 Cidr Block and how to use it, see [What is Ipv4 Cidr Block](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/associatevpccidrblock).

        > **NOTE:** Available in v1.185.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.vpc.Network("default",
            cidr_block="192.168.0.0/24",
            vpc_name=var["name"])
        example = alicloud.vpc.Ipv4CidrBlock("example",
            vpc_id=default.id,
            secondary_cidr_block="192.163.0.0/16")
        ```

        ## Import

        VPC Ipv4 Cidr Block can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock example <vpc_id>:<secondary_cidr_block>
        ```

        :param str resource_name: The name of the resource.
        :param Ipv4CidrBlockArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ipv4CidrBlockArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 secondary_cidr_block: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Ipv4CidrBlockArgs.__new__(Ipv4CidrBlockArgs)

            if secondary_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'secondary_cidr_block'")
            __props__.__dict__["secondary_cidr_block"] = secondary_cidr_block
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(Ipv4CidrBlock, __self__).__init__(
            'alicloud:vpc/ipv4CidrBlock:Ipv4CidrBlock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            secondary_cidr_block: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'Ipv4CidrBlock':
        """
        Get an existing Ipv4CidrBlock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] secondary_cidr_block: The IPv4 CIDR block. Take note of the following requirements:
               * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
               * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
               * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
               * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ipv4CidrBlockState.__new__(_Ipv4CidrBlockState)

        __props__.__dict__["secondary_cidr_block"] = secondary_cidr_block
        __props__.__dict__["vpc_id"] = vpc_id
        return Ipv4CidrBlock(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="secondaryCidrBlock")
    def secondary_cidr_block(self) -> pulumi.Output[str]:
        """
        The IPv4 CIDR block. Take note of the following requirements:
        * You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
        * You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
        * The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
        * The secondary CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block.
        """
        return pulumi.get(self, "secondary_cidr_block")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

