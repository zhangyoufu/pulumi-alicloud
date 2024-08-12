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

__all__ = ['VvpInstanceArgs', 'VvpInstance']

@pulumi.input_type
class VvpInstanceArgs:
    def __init__(__self__, *,
                 payment_type: pulumi.Input[str],
                 storage: pulumi.Input['VvpInstanceStorageArgs'],
                 vpc_id: pulumi.Input[str],
                 vswitch_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 vvp_instance_name: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 duration: Optional[pulumi.Input[int]] = None,
                 pricing_cycle: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 resource_spec: Optional[pulumi.Input['VvpInstanceResourceSpecArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a VvpInstance resource.
        :param pulumi.Input[str] payment_type: The payment type of the resource.
        :param pulumi.Input['VvpInstanceStorageArgs'] storage: Store information. See `storage` below.
        :param pulumi.Input[str] vpc_id: The VPC ID of the user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: Virtual Switch ID.
        :param pulumi.Input[str] vvp_instance_name: The name of the workspace.
        :param pulumi.Input[str] zone_id: The zone ID of the resource.
        :param pulumi.Input[int] duration: The number of subscription periods. If the payment type is PRE, this parameter is required.
        :param pulumi.Input[str] pricing_cycle: The subscription period. If the payment type is PRE, this parameter is required.
        :param pulumi.Input[str] resource_group_id: The resource group to which the newly purchased instance belongs.
        :param pulumi.Input['VvpInstanceResourceSpecArgs'] resource_spec: Resource specifications. See `resource_spec` below.
        :param pulumi.Input[Mapping[str, Any]] tags: The tags of the resource.
        """
        pulumi.set(__self__, "payment_type", payment_type)
        pulumi.set(__self__, "storage", storage)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitch_ids", vswitch_ids)
        pulumi.set(__self__, "vvp_instance_name", vvp_instance_name)
        pulumi.set(__self__, "zone_id", zone_id)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if pricing_cycle is not None:
            pulumi.set(__self__, "pricing_cycle", pricing_cycle)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if resource_spec is not None:
            pulumi.set(__self__, "resource_spec", resource_spec)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> pulumi.Input[str]:
        """
        The payment type of the resource.
        """
        return pulumi.get(self, "payment_type")

    @payment_type.setter
    def payment_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "payment_type", value)

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Input['VvpInstanceStorageArgs']:
        """
        Store information. See `storage` below.
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: pulumi.Input['VvpInstanceStorageArgs']):
        pulumi.set(self, "storage", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The VPC ID of the user.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Virtual Switch ID.
        """
        return pulumi.get(self, "vswitch_ids")

    @vswitch_ids.setter
    def vswitch_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "vswitch_ids", value)

    @property
    @pulumi.getter(name="vvpInstanceName")
    def vvp_instance_name(self) -> pulumi.Input[str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "vvp_instance_name")

    @vvp_instance_name.setter
    def vvp_instance_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vvp_instance_name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The zone ID of the resource.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        The number of subscription periods. If the payment type is PRE, this parameter is required.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="pricingCycle")
    def pricing_cycle(self) -> Optional[pulumi.Input[str]]:
        """
        The subscription period. If the payment type is PRE, this parameter is required.
        """
        return pulumi.get(self, "pricing_cycle")

    @pricing_cycle.setter
    def pricing_cycle(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pricing_cycle", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource group to which the newly purchased instance belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="resourceSpec")
    def resource_spec(self) -> Optional[pulumi.Input['VvpInstanceResourceSpecArgs']]:
        """
        Resource specifications. See `resource_spec` below.
        """
        return pulumi.get(self, "resource_spec")

    @resource_spec.setter
    def resource_spec(self, value: Optional[pulumi.Input['VvpInstanceResourceSpecArgs']]):
        pulumi.set(self, "resource_spec", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VvpInstanceState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 pricing_cycle: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 resource_spec: Optional[pulumi.Input['VvpInstanceResourceSpecArgs']] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 storage: Optional[pulumi.Input['VvpInstanceStorageArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vvp_instance_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VvpInstance resources.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[int] duration: The number of subscription periods. If the payment type is PRE, this parameter is required.
        :param pulumi.Input[str] payment_type: The payment type of the resource.
        :param pulumi.Input[str] pricing_cycle: The subscription period. If the payment type is PRE, this parameter is required.
        :param pulumi.Input[str] resource_group_id: The resource group to which the newly purchased instance belongs.
        :param pulumi.Input['VvpInstanceResourceSpecArgs'] resource_spec: Resource specifications. See `resource_spec` below.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input['VvpInstanceStorageArgs'] storage: Store information. See `storage` below.
        :param pulumi.Input[Mapping[str, Any]] tags: The tags of the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID of the user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: Virtual Switch ID.
        :param pulumi.Input[str] vvp_instance_name: The name of the workspace.
        :param pulumi.Input[str] zone_id: The zone ID of the resource.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if payment_type is not None:
            pulumi.set(__self__, "payment_type", payment_type)
        if pricing_cycle is not None:
            pulumi.set(__self__, "pricing_cycle", pricing_cycle)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if resource_spec is not None:
            pulumi.set(__self__, "resource_spec", resource_spec)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if storage is not None:
            pulumi.set(__self__, "storage", storage)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_ids is not None:
            pulumi.set(__self__, "vswitch_ids", vswitch_ids)
        if vvp_instance_name is not None:
            pulumi.set(__self__, "vvp_instance_name", vvp_instance_name)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        The number of subscription periods. If the payment type is PRE, this parameter is required.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> Optional[pulumi.Input[str]]:
        """
        The payment type of the resource.
        """
        return pulumi.get(self, "payment_type")

    @payment_type.setter
    def payment_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_type", value)

    @property
    @pulumi.getter(name="pricingCycle")
    def pricing_cycle(self) -> Optional[pulumi.Input[str]]:
        """
        The subscription period. If the payment type is PRE, this parameter is required.
        """
        return pulumi.get(self, "pricing_cycle")

    @pricing_cycle.setter
    def pricing_cycle(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pricing_cycle", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource group to which the newly purchased instance belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="resourceSpec")
    def resource_spec(self) -> Optional[pulumi.Input['VvpInstanceResourceSpecArgs']]:
        """
        Resource specifications. See `resource_spec` below.
        """
        return pulumi.get(self, "resource_spec")

    @resource_spec.setter
    def resource_spec(self, value: Optional[pulumi.Input['VvpInstanceResourceSpecArgs']]):
        pulumi.set(self, "resource_spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def storage(self) -> Optional[pulumi.Input['VvpInstanceStorageArgs']]:
        """
        Store information. See `storage` below.
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: Optional[pulumi.Input['VvpInstanceStorageArgs']]):
        pulumi.set(self, "storage", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC ID of the user.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Virtual Switch ID.
        """
        return pulumi.get(self, "vswitch_ids")

    @vswitch_ids.setter
    def vswitch_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vswitch_ids", value)

    @property
    @pulumi.getter(name="vvpInstanceName")
    def vvp_instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "vvp_instance_name")

    @vvp_instance_name.setter
    def vvp_instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vvp_instance_name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The zone ID of the resource.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class VvpInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 pricing_cycle: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 resource_spec: Optional[pulumi.Input[Union['VvpInstanceResourceSpecArgs', 'VvpInstanceResourceSpecArgsDict']]] = None,
                 storage: Optional[pulumi.Input[Union['VvpInstanceStorageArgs', 'VvpInstanceStorageArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vvp_instance_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Realtime Compute Vvp Instance can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:realtimecompute/vvpInstance:VvpInstance example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] duration: The number of subscription periods. If the payment type is PRE, this parameter is required.
        :param pulumi.Input[str] payment_type: The payment type of the resource.
        :param pulumi.Input[str] pricing_cycle: The subscription period. If the payment type is PRE, this parameter is required.
        :param pulumi.Input[str] resource_group_id: The resource group to which the newly purchased instance belongs.
        :param pulumi.Input[Union['VvpInstanceResourceSpecArgs', 'VvpInstanceResourceSpecArgsDict']] resource_spec: Resource specifications. See `resource_spec` below.
        :param pulumi.Input[Union['VvpInstanceStorageArgs', 'VvpInstanceStorageArgsDict']] storage: Store information. See `storage` below.
        :param pulumi.Input[Mapping[str, Any]] tags: The tags of the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID of the user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: Virtual Switch ID.
        :param pulumi.Input[str] vvp_instance_name: The name of the workspace.
        :param pulumi.Input[str] zone_id: The zone ID of the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VvpInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Realtime Compute Vvp Instance can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:realtimecompute/vvpInstance:VvpInstance example <id>
        ```

        :param str resource_name: The name of the resource.
        :param VvpInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VvpInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 pricing_cycle: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 resource_spec: Optional[pulumi.Input[Union['VvpInstanceResourceSpecArgs', 'VvpInstanceResourceSpecArgsDict']]] = None,
                 storage: Optional[pulumi.Input[Union['VvpInstanceStorageArgs', 'VvpInstanceStorageArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vvp_instance_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VvpInstanceArgs.__new__(VvpInstanceArgs)

            __props__.__dict__["duration"] = duration
            if payment_type is None and not opts.urn:
                raise TypeError("Missing required property 'payment_type'")
            __props__.__dict__["payment_type"] = payment_type
            __props__.__dict__["pricing_cycle"] = pricing_cycle
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["resource_spec"] = resource_spec
            if storage is None and not opts.urn:
                raise TypeError("Missing required property 'storage'")
            __props__.__dict__["storage"] = storage
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if vswitch_ids is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_ids'")
            __props__.__dict__["vswitch_ids"] = vswitch_ids
            if vvp_instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'vvp_instance_name'")
            __props__.__dict__["vvp_instance_name"] = vvp_instance_name
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["status"] = None
        super(VvpInstance, __self__).__init__(
            'alicloud:realtimecompute/vvpInstance:VvpInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            duration: Optional[pulumi.Input[int]] = None,
            payment_type: Optional[pulumi.Input[str]] = None,
            pricing_cycle: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            resource_spec: Optional[pulumi.Input[Union['VvpInstanceResourceSpecArgs', 'VvpInstanceResourceSpecArgsDict']]] = None,
            status: Optional[pulumi.Input[str]] = None,
            storage: Optional[pulumi.Input[Union['VvpInstanceStorageArgs', 'VvpInstanceStorageArgsDict']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vswitch_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            vvp_instance_name: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'VvpInstance':
        """
        Get an existing VvpInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[int] duration: The number of subscription periods. If the payment type is PRE, this parameter is required.
        :param pulumi.Input[str] payment_type: The payment type of the resource.
        :param pulumi.Input[str] pricing_cycle: The subscription period. If the payment type is PRE, this parameter is required.
        :param pulumi.Input[str] resource_group_id: The resource group to which the newly purchased instance belongs.
        :param pulumi.Input[Union['VvpInstanceResourceSpecArgs', 'VvpInstanceResourceSpecArgsDict']] resource_spec: Resource specifications. See `resource_spec` below.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[Union['VvpInstanceStorageArgs', 'VvpInstanceStorageArgsDict']] storage: Store information. See `storage` below.
        :param pulumi.Input[Mapping[str, Any]] tags: The tags of the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID of the user.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: Virtual Switch ID.
        :param pulumi.Input[str] vvp_instance_name: The name of the workspace.
        :param pulumi.Input[str] zone_id: The zone ID of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VvpInstanceState.__new__(_VvpInstanceState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["duration"] = duration
        __props__.__dict__["payment_type"] = payment_type
        __props__.__dict__["pricing_cycle"] = pricing_cycle
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["resource_spec"] = resource_spec
        __props__.__dict__["status"] = status
        __props__.__dict__["storage"] = storage
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vswitch_ids"] = vswitch_ids
        __props__.__dict__["vvp_instance_name"] = vvp_instance_name
        __props__.__dict__["zone_id"] = zone_id
        return VvpInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[Optional[int]]:
        """
        The number of subscription periods. If the payment type is PRE, this parameter is required.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> pulumi.Output[str]:
        """
        The payment type of the resource.
        """
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter(name="pricingCycle")
    def pricing_cycle(self) -> pulumi.Output[Optional[str]]:
        """
        The subscription period. If the payment type is PRE, this parameter is required.
        """
        return pulumi.get(self, "pricing_cycle")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The resource group to which the newly purchased instance belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="resourceSpec")
    def resource_spec(self) -> pulumi.Output['outputs.VvpInstanceResourceSpec']:
        """
        Resource specifications. See `resource_spec` below.
        """
        return pulumi.get(self, "resource_spec")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Output['outputs.VvpInstanceStorage']:
        """
        Store information. See `storage` below.
        """
        return pulumi.get(self, "storage")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The VPC ID of the user.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Virtual Switch ID.
        """
        return pulumi.get(self, "vswitch_ids")

    @property
    @pulumi.getter(name="vvpInstanceName")
    def vvp_instance_name(self) -> pulumi.Output[str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "vvp_instance_name")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The zone ID of the resource.
        """
        return pulumi.get(self, "zone_id")

