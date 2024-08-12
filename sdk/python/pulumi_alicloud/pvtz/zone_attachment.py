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

__all__ = ['ZoneAttachmentArgs', 'ZoneAttachment']

@pulumi.input_type
class ZoneAttachmentArgs:
    def __init__(__self__, *,
                 zone_id: pulumi.Input[str],
                 lang: Optional[pulumi.Input[str]] = None,
                 user_client_ip: Optional[pulumi.Input[str]] = None,
                 vpc_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAttachmentVpcArgs']]]] = None):
        """
        The set of arguments for constructing a ZoneAttachment resource.
        :param pulumi.Input[str] zone_id: The name of the Private Zone Record.
        :param pulumi.Input[str] lang: The language of code.
        :param pulumi.Input[str] user_client_ip: The user custom IP address.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_ids: The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
        :param pulumi.Input[Sequence[pulumi.Input['ZoneAttachmentVpcArgs']]] vpcs: See `vpcs` below.Recommend to use `vpcs`.
        """
        pulumi.set(__self__, "zone_id", zone_id)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if user_client_ip is not None:
            pulumi.set(__self__, "user_client_ip", user_client_ip)
        if vpc_ids is not None:
            pulumi.set(__self__, "vpc_ids", vpc_ids)
        if vpcs is not None:
            pulumi.set(__self__, "vpcs", vpcs)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The name of the Private Zone Record.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The language of code.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter(name="userClientIp")
    def user_client_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The user custom IP address.
        """
        return pulumi.get(self, "user_client_ip")

    @user_client_ip.setter
    def user_client_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_client_ip", value)

    @property
    @pulumi.getter(name="vpcIds")
    def vpc_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
        """
        return pulumi.get(self, "vpc_ids")

    @vpc_ids.setter
    def vpc_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_ids", value)

    @property
    @pulumi.getter
    def vpcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAttachmentVpcArgs']]]]:
        """
        See `vpcs` below.Recommend to use `vpcs`.
        """
        return pulumi.get(self, "vpcs")

    @vpcs.setter
    def vpcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAttachmentVpcArgs']]]]):
        pulumi.set(self, "vpcs", value)


@pulumi.input_type
class _ZoneAttachmentState:
    def __init__(__self__, *,
                 lang: Optional[pulumi.Input[str]] = None,
                 user_client_ip: Optional[pulumi.Input[str]] = None,
                 vpc_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAttachmentVpcArgs']]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ZoneAttachment resources.
        :param pulumi.Input[str] lang: The language of code.
        :param pulumi.Input[str] user_client_ip: The user custom IP address.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_ids: The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
        :param pulumi.Input[Sequence[pulumi.Input['ZoneAttachmentVpcArgs']]] vpcs: See `vpcs` below.Recommend to use `vpcs`.
        :param pulumi.Input[str] zone_id: The name of the Private Zone Record.
        """
        if lang is not None:
            pulumi.set(__self__, "lang", lang)
        if user_client_ip is not None:
            pulumi.set(__self__, "user_client_ip", user_client_ip)
        if vpc_ids is not None:
            pulumi.set(__self__, "vpc_ids", vpc_ids)
        if vpcs is not None:
            pulumi.set(__self__, "vpcs", vpcs)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The language of code.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)

    @property
    @pulumi.getter(name="userClientIp")
    def user_client_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The user custom IP address.
        """
        return pulumi.get(self, "user_client_ip")

    @user_client_ip.setter
    def user_client_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_client_ip", value)

    @property
    @pulumi.getter(name="vpcIds")
    def vpc_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
        """
        return pulumi.get(self, "vpc_ids")

    @vpc_ids.setter
    def vpc_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vpc_ids", value)

    @property
    @pulumi.getter
    def vpcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAttachmentVpcArgs']]]]:
        """
        See `vpcs` below.Recommend to use `vpcs`.
        """
        return pulumi.get(self, "vpcs")

    @vpcs.setter
    def vpcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAttachmentVpcArgs']]]]):
        pulumi.set(self, "vpcs", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Private Zone Record.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class ZoneAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 user_client_ip: Optional[pulumi.Input[str]] = None,
                 vpc_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ZoneAttachmentVpcArgs', 'ZoneAttachmentVpcArgsDict']]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        Using `vpc_ids` to attach being in same region several vpc instances to a private zone

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        zone = alicloud.pvtz.Zone("zone", zone_name="foo.example.com")
        first = alicloud.vpc.Network("first",
            vpc_name="the-first-vpc",
            cidr_block="172.16.0.0/12")
        second = alicloud.vpc.Network("second",
            vpc_name="the-second-vpc",
            cidr_block="172.16.0.0/16")
        zone_attachment = alicloud.pvtz.ZoneAttachment("zone-attachment",
            zone_id=zone.id,
            vpc_ids=[
                first.id,
                second.id,
            ])
        ```

        Using `vpcs` to attach being in same region several vpc instances to a private zone

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        zone = alicloud.pvtz.Zone("zone", zone_name="foo.example.com")
        first = alicloud.vpc.Network("first",
            vpc_name="the-first-vpc",
            cidr_block="172.16.0.0/12")
        second = alicloud.vpc.Network("second",
            vpc_name="the-second-vpc",
            cidr_block="172.16.0.0/16")
        zone_attachment = alicloud.pvtz.ZoneAttachment("zone-attachment",
            zone_id=zone.id,
            vpcs=[
                {
                    "vpc_id": first.id,
                },
                {
                    "vpc_id": second.id,
                },
            ])
        ```

        Using `vpcs` to attach being in different regions several vpc instances to a private zone

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        zone = alicloud.pvtz.Zone("zone", zone_name="foo.example.com")
        first = alicloud.vpc.Network("first",
            vpc_name="the-first-vpc",
            cidr_block="172.16.0.0/12")
        second = alicloud.vpc.Network("second",
            vpc_name="the-second-vpc",
            cidr_block="172.16.0.0/16")
        third = alicloud.vpc.Network("third",
            vpc_name="the-third-vpc",
            cidr_block="172.16.0.0/16")
        zone_attachment = alicloud.pvtz.ZoneAttachment("zone-attachment",
            zone_id=zone.id,
            vpcs=[
                {
                    "vpc_id": first.id,
                },
                {
                    "vpc_id": second.id,
                },
                {
                    "region_id": "eu-central-1",
                    "vpc_id": third.id,
                },
            ])
        ```

        ## Import

        Private Zone attachment can be imported using the id(same with `zone_id`), e.g.

        ```sh
        $ pulumi import alicloud:pvtz/zoneAttachment:ZoneAttachment example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lang: The language of code.
        :param pulumi.Input[str] user_client_ip: The user custom IP address.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_ids: The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
        :param pulumi.Input[Sequence[pulumi.Input[Union['ZoneAttachmentVpcArgs', 'ZoneAttachmentVpcArgsDict']]]] vpcs: See `vpcs` below.Recommend to use `vpcs`.
        :param pulumi.Input[str] zone_id: The name of the Private Zone Record.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ZoneAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Using `vpc_ids` to attach being in same region several vpc instances to a private zone

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        zone = alicloud.pvtz.Zone("zone", zone_name="foo.example.com")
        first = alicloud.vpc.Network("first",
            vpc_name="the-first-vpc",
            cidr_block="172.16.0.0/12")
        second = alicloud.vpc.Network("second",
            vpc_name="the-second-vpc",
            cidr_block="172.16.0.0/16")
        zone_attachment = alicloud.pvtz.ZoneAttachment("zone-attachment",
            zone_id=zone.id,
            vpc_ids=[
                first.id,
                second.id,
            ])
        ```

        Using `vpcs` to attach being in same region several vpc instances to a private zone

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        zone = alicloud.pvtz.Zone("zone", zone_name="foo.example.com")
        first = alicloud.vpc.Network("first",
            vpc_name="the-first-vpc",
            cidr_block="172.16.0.0/12")
        second = alicloud.vpc.Network("second",
            vpc_name="the-second-vpc",
            cidr_block="172.16.0.0/16")
        zone_attachment = alicloud.pvtz.ZoneAttachment("zone-attachment",
            zone_id=zone.id,
            vpcs=[
                {
                    "vpc_id": first.id,
                },
                {
                    "vpc_id": second.id,
                },
            ])
        ```

        Using `vpcs` to attach being in different regions several vpc instances to a private zone

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        zone = alicloud.pvtz.Zone("zone", zone_name="foo.example.com")
        first = alicloud.vpc.Network("first",
            vpc_name="the-first-vpc",
            cidr_block="172.16.0.0/12")
        second = alicloud.vpc.Network("second",
            vpc_name="the-second-vpc",
            cidr_block="172.16.0.0/16")
        third = alicloud.vpc.Network("third",
            vpc_name="the-third-vpc",
            cidr_block="172.16.0.0/16")
        zone_attachment = alicloud.pvtz.ZoneAttachment("zone-attachment",
            zone_id=zone.id,
            vpcs=[
                {
                    "vpc_id": first.id,
                },
                {
                    "vpc_id": second.id,
                },
                {
                    "region_id": "eu-central-1",
                    "vpc_id": third.id,
                },
            ])
        ```

        ## Import

        Private Zone attachment can be imported using the id(same with `zone_id`), e.g.

        ```sh
        $ pulumi import alicloud:pvtz/zoneAttachment:ZoneAttachment example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param ZoneAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 user_client_ip: Optional[pulumi.Input[str]] = None,
                 vpc_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ZoneAttachmentVpcArgs', 'ZoneAttachmentVpcArgsDict']]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ZoneAttachmentArgs.__new__(ZoneAttachmentArgs)

            __props__.__dict__["lang"] = lang
            __props__.__dict__["user_client_ip"] = user_client_ip
            __props__.__dict__["vpc_ids"] = vpc_ids
            __props__.__dict__["vpcs"] = vpcs
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(ZoneAttachment, __self__).__init__(
            'alicloud:pvtz/zoneAttachment:ZoneAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            lang: Optional[pulumi.Input[str]] = None,
            user_client_ip: Optional[pulumi.Input[str]] = None,
            vpc_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ZoneAttachmentVpcArgs', 'ZoneAttachmentVpcArgsDict']]]]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'ZoneAttachment':
        """
        Get an existing ZoneAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lang: The language of code.
        :param pulumi.Input[str] user_client_ip: The user custom IP address.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_ids: The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
        :param pulumi.Input[Sequence[pulumi.Input[Union['ZoneAttachmentVpcArgs', 'ZoneAttachmentVpcArgsDict']]]] vpcs: See `vpcs` below.Recommend to use `vpcs`.
        :param pulumi.Input[str] zone_id: The name of the Private Zone Record.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneAttachmentState.__new__(_ZoneAttachmentState)

        __props__.__dict__["lang"] = lang
        __props__.__dict__["user_client_ip"] = user_client_ip
        __props__.__dict__["vpc_ids"] = vpc_ids
        __props__.__dict__["vpcs"] = vpcs
        __props__.__dict__["zone_id"] = zone_id
        return ZoneAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def lang(self) -> pulumi.Output[Optional[str]]:
        """
        The language of code.
        """
        return pulumi.get(self, "lang")

    @property
    @pulumi.getter(name="userClientIp")
    def user_client_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The user custom IP address.
        """
        return pulumi.get(self, "user_client_ip")

    @property
    @pulumi.getter(name="vpcIds")
    def vpc_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].
        """
        return pulumi.get(self, "vpc_ids")

    @property
    @pulumi.getter
    def vpcs(self) -> pulumi.Output[Sequence['outputs.ZoneAttachmentVpc']]:
        """
        See `vpcs` below.Recommend to use `vpcs`.
        """
        return pulumi.get(self, "vpcs")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The name of the Private Zone Record.
        """
        return pulumi.get(self, "zone_id")

