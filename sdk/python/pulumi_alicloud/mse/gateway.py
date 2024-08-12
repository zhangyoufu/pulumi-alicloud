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

__all__ = ['GatewayArgs', 'Gateway']

@pulumi.input_type
class GatewayArgs:
    def __init__(__self__, *,
                 replica: pulumi.Input[int],
                 spec: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str],
                 backup_vswitch_id: Optional[pulumi.Input[str]] = None,
                 delete_slb: Optional[pulumi.Input[bool]] = None,
                 enterprise_security_group: Optional[pulumi.Input[bool]] = None,
                 gateway_name: Optional[pulumi.Input[str]] = None,
                 internet_slb_spec: Optional[pulumi.Input[str]] = None,
                 slb_spec: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Gateway resource.
        :param pulumi.Input[int] replica: Number of Gateway Nodes.
        :param pulumi.Input[str] spec: Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        :param pulumi.Input[str] vpc_id: The ID of the vpc.
        :param pulumi.Input[str] vswitch_id: The ID of the vswitch.
        :param pulumi.Input[str] backup_vswitch_id: The backup vswitch id.
        :param pulumi.Input[bool] delete_slb: Whether to delete the SLB purchased on behalf of the gateway at the same time.
        :param pulumi.Input[bool] enterprise_security_group: Whether the enterprise security group type.
        :param pulumi.Input[str] gateway_name: The name of the Gateway .
        :param pulumi.Input[str] internet_slb_spec: Public network SLB specifications.
        :param pulumi.Input[str] slb_spec: Private network SLB specifications.
        """
        pulumi.set(__self__, "replica", replica)
        pulumi.set(__self__, "spec", spec)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if backup_vswitch_id is not None:
            pulumi.set(__self__, "backup_vswitch_id", backup_vswitch_id)
        if delete_slb is not None:
            pulumi.set(__self__, "delete_slb", delete_slb)
        if enterprise_security_group is not None:
            pulumi.set(__self__, "enterprise_security_group", enterprise_security_group)
        if gateway_name is not None:
            pulumi.set(__self__, "gateway_name", gateway_name)
        if internet_slb_spec is not None:
            pulumi.set(__self__, "internet_slb_spec", internet_slb_spec)
        if slb_spec is not None:
            pulumi.set(__self__, "slb_spec", slb_spec)

    @property
    @pulumi.getter
    def replica(self) -> pulumi.Input[int]:
        """
        Number of Gateway Nodes.
        """
        return pulumi.get(self, "replica")

    @replica.setter
    def replica(self, value: pulumi.Input[int]):
        pulumi.set(self, "replica", value)

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Input[str]:
        """
        Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: pulumi.Input[str]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the vpc.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        The ID of the vswitch.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="backupVswitchId")
    def backup_vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The backup vswitch id.
        """
        return pulumi.get(self, "backup_vswitch_id")

    @backup_vswitch_id.setter
    def backup_vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_vswitch_id", value)

    @property
    @pulumi.getter(name="deleteSlb")
    def delete_slb(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to delete the SLB purchased on behalf of the gateway at the same time.
        """
        return pulumi.get(self, "delete_slb")

    @delete_slb.setter
    def delete_slb(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_slb", value)

    @property
    @pulumi.getter(name="enterpriseSecurityGroup")
    def enterprise_security_group(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the enterprise security group type.
        """
        return pulumi.get(self, "enterprise_security_group")

    @enterprise_security_group.setter
    def enterprise_security_group(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enterprise_security_group", value)

    @property
    @pulumi.getter(name="gatewayName")
    def gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Gateway .
        """
        return pulumi.get(self, "gateway_name")

    @gateway_name.setter
    def gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_name", value)

    @property
    @pulumi.getter(name="internetSlbSpec")
    def internet_slb_spec(self) -> Optional[pulumi.Input[str]]:
        """
        Public network SLB specifications.
        """
        return pulumi.get(self, "internet_slb_spec")

    @internet_slb_spec.setter
    def internet_slb_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_slb_spec", value)

    @property
    @pulumi.getter(name="slbSpec")
    def slb_spec(self) -> Optional[pulumi.Input[str]]:
        """
        Private network SLB specifications.
        """
        return pulumi.get(self, "slb_spec")

    @slb_spec.setter
    def slb_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slb_spec", value)


@pulumi.input_type
class _GatewayState:
    def __init__(__self__, *,
                 backup_vswitch_id: Optional[pulumi.Input[str]] = None,
                 delete_slb: Optional[pulumi.Input[bool]] = None,
                 enterprise_security_group: Optional[pulumi.Input[bool]] = None,
                 gateway_name: Optional[pulumi.Input[str]] = None,
                 internet_slb_spec: Optional[pulumi.Input[str]] = None,
                 replica: Optional[pulumi.Input[int]] = None,
                 slb_lists: Optional[pulumi.Input[Sequence[pulumi.Input['GatewaySlbListArgs']]]] = None,
                 slb_spec: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Gateway resources.
        :param pulumi.Input[str] backup_vswitch_id: The backup vswitch id.
        :param pulumi.Input[bool] delete_slb: Whether to delete the SLB purchased on behalf of the gateway at the same time.
        :param pulumi.Input[bool] enterprise_security_group: Whether the enterprise security group type.
        :param pulumi.Input[str] gateway_name: The name of the Gateway .
        :param pulumi.Input[str] internet_slb_spec: Public network SLB specifications.
        :param pulumi.Input[int] replica: Number of Gateway Nodes.
        :param pulumi.Input[Sequence[pulumi.Input['GatewaySlbListArgs']]] slb_lists: A list of gateway Slb.
        :param pulumi.Input[str] slb_spec: Private network SLB specifications.
        :param pulumi.Input[str] spec: Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        :param pulumi.Input[str] status: The status of the gateway.
        :param pulumi.Input[str] vpc_id: The ID of the vpc.
        :param pulumi.Input[str] vswitch_id: The ID of the vswitch.
        """
        if backup_vswitch_id is not None:
            pulumi.set(__self__, "backup_vswitch_id", backup_vswitch_id)
        if delete_slb is not None:
            pulumi.set(__self__, "delete_slb", delete_slb)
        if enterprise_security_group is not None:
            pulumi.set(__self__, "enterprise_security_group", enterprise_security_group)
        if gateway_name is not None:
            pulumi.set(__self__, "gateway_name", gateway_name)
        if internet_slb_spec is not None:
            pulumi.set(__self__, "internet_slb_spec", internet_slb_spec)
        if replica is not None:
            pulumi.set(__self__, "replica", replica)
        if slb_lists is not None:
            pulumi.set(__self__, "slb_lists", slb_lists)
        if slb_spec is not None:
            pulumi.set(__self__, "slb_spec", slb_spec)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="backupVswitchId")
    def backup_vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The backup vswitch id.
        """
        return pulumi.get(self, "backup_vswitch_id")

    @backup_vswitch_id.setter
    def backup_vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_vswitch_id", value)

    @property
    @pulumi.getter(name="deleteSlb")
    def delete_slb(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to delete the SLB purchased on behalf of the gateway at the same time.
        """
        return pulumi.get(self, "delete_slb")

    @delete_slb.setter
    def delete_slb(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_slb", value)

    @property
    @pulumi.getter(name="enterpriseSecurityGroup")
    def enterprise_security_group(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the enterprise security group type.
        """
        return pulumi.get(self, "enterprise_security_group")

    @enterprise_security_group.setter
    def enterprise_security_group(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enterprise_security_group", value)

    @property
    @pulumi.getter(name="gatewayName")
    def gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Gateway .
        """
        return pulumi.get(self, "gateway_name")

    @gateway_name.setter
    def gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_name", value)

    @property
    @pulumi.getter(name="internetSlbSpec")
    def internet_slb_spec(self) -> Optional[pulumi.Input[str]]:
        """
        Public network SLB specifications.
        """
        return pulumi.get(self, "internet_slb_spec")

    @internet_slb_spec.setter
    def internet_slb_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_slb_spec", value)

    @property
    @pulumi.getter
    def replica(self) -> Optional[pulumi.Input[int]]:
        """
        Number of Gateway Nodes.
        """
        return pulumi.get(self, "replica")

    @replica.setter
    def replica(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "replica", value)

    @property
    @pulumi.getter(name="slbLists")
    def slb_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GatewaySlbListArgs']]]]:
        """
        A list of gateway Slb.
        """
        return pulumi.get(self, "slb_lists")

    @slb_lists.setter
    def slb_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GatewaySlbListArgs']]]]):
        pulumi.set(self, "slb_lists", value)

    @property
    @pulumi.getter(name="slbSpec")
    def slb_spec(self) -> Optional[pulumi.Input[str]]:
        """
        Private network SLB specifications.
        """
        return pulumi.get(self, "slb_spec")

    @slb_spec.setter
    def slb_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slb_spec", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input[str]]:
        """
        Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the gateway.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the vpc.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the vswitch.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class Gateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_vswitch_id: Optional[pulumi.Input[str]] = None,
                 delete_slb: Optional[pulumi.Input[bool]] = None,
                 enterprise_security_group: Optional[pulumi.Input[bool]] = None,
                 gateway_name: Optional[pulumi.Input[str]] = None,
                 internet_slb_spec: Optional[pulumi.Input[str]] = None,
                 replica: Optional[pulumi.Input[int]] = None,
                 slb_spec: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Microservice Engine (MSE) Gateway resource.

        For information about Microservice Engine (MSE) Gateway and how to use it, see [What is Gateway](https://help.aliyun.com/document_detail/347638.html).

        > **NOTE:** Available in v1.157.0+.

        ## Import

        Microservice Engine (MSE) Gateway can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:mse/gateway:Gateway example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_vswitch_id: The backup vswitch id.
        :param pulumi.Input[bool] delete_slb: Whether to delete the SLB purchased on behalf of the gateway at the same time.
        :param pulumi.Input[bool] enterprise_security_group: Whether the enterprise security group type.
        :param pulumi.Input[str] gateway_name: The name of the Gateway .
        :param pulumi.Input[str] internet_slb_spec: Public network SLB specifications.
        :param pulumi.Input[int] replica: Number of Gateway Nodes.
        :param pulumi.Input[str] slb_spec: Private network SLB specifications.
        :param pulumi.Input[str] spec: Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        :param pulumi.Input[str] vpc_id: The ID of the vpc.
        :param pulumi.Input[str] vswitch_id: The ID of the vswitch.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Microservice Engine (MSE) Gateway resource.

        For information about Microservice Engine (MSE) Gateway and how to use it, see [What is Gateway](https://help.aliyun.com/document_detail/347638.html).

        > **NOTE:** Available in v1.157.0+.

        ## Import

        Microservice Engine (MSE) Gateway can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:mse/gateway:Gateway example <id>
        ```

        :param str resource_name: The name of the resource.
        :param GatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_vswitch_id: Optional[pulumi.Input[str]] = None,
                 delete_slb: Optional[pulumi.Input[bool]] = None,
                 enterprise_security_group: Optional[pulumi.Input[bool]] = None,
                 gateway_name: Optional[pulumi.Input[str]] = None,
                 internet_slb_spec: Optional[pulumi.Input[str]] = None,
                 replica: Optional[pulumi.Input[int]] = None,
                 slb_spec: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewayArgs.__new__(GatewayArgs)

            __props__.__dict__["backup_vswitch_id"] = backup_vswitch_id
            __props__.__dict__["delete_slb"] = delete_slb
            __props__.__dict__["enterprise_security_group"] = enterprise_security_group
            __props__.__dict__["gateway_name"] = gateway_name
            __props__.__dict__["internet_slb_spec"] = internet_slb_spec
            if replica is None and not opts.urn:
                raise TypeError("Missing required property 'replica'")
            __props__.__dict__["replica"] = replica
            __props__.__dict__["slb_spec"] = slb_spec
            if spec is None and not opts.urn:
                raise TypeError("Missing required property 'spec'")
            __props__.__dict__["spec"] = spec
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["slb_lists"] = None
            __props__.__dict__["status"] = None
        super(Gateway, __self__).__init__(
            'alicloud:mse/gateway:Gateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_vswitch_id: Optional[pulumi.Input[str]] = None,
            delete_slb: Optional[pulumi.Input[bool]] = None,
            enterprise_security_group: Optional[pulumi.Input[bool]] = None,
            gateway_name: Optional[pulumi.Input[str]] = None,
            internet_slb_spec: Optional[pulumi.Input[str]] = None,
            replica: Optional[pulumi.Input[int]] = None,
            slb_lists: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GatewaySlbListArgs', 'GatewaySlbListArgsDict']]]]] = None,
            slb_spec: Optional[pulumi.Input[str]] = None,
            spec: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'Gateway':
        """
        Get an existing Gateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_vswitch_id: The backup vswitch id.
        :param pulumi.Input[bool] delete_slb: Whether to delete the SLB purchased on behalf of the gateway at the same time.
        :param pulumi.Input[bool] enterprise_security_group: Whether the enterprise security group type.
        :param pulumi.Input[str] gateway_name: The name of the Gateway .
        :param pulumi.Input[str] internet_slb_spec: Public network SLB specifications.
        :param pulumi.Input[int] replica: Number of Gateway Nodes.
        :param pulumi.Input[Sequence[pulumi.Input[Union['GatewaySlbListArgs', 'GatewaySlbListArgsDict']]]] slb_lists: A list of gateway Slb.
        :param pulumi.Input[str] slb_spec: Private network SLB specifications.
        :param pulumi.Input[str] spec: Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        :param pulumi.Input[str] status: The status of the gateway.
        :param pulumi.Input[str] vpc_id: The ID of the vpc.
        :param pulumi.Input[str] vswitch_id: The ID of the vswitch.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayState.__new__(_GatewayState)

        __props__.__dict__["backup_vswitch_id"] = backup_vswitch_id
        __props__.__dict__["delete_slb"] = delete_slb
        __props__.__dict__["enterprise_security_group"] = enterprise_security_group
        __props__.__dict__["gateway_name"] = gateway_name
        __props__.__dict__["internet_slb_spec"] = internet_slb_spec
        __props__.__dict__["replica"] = replica
        __props__.__dict__["slb_lists"] = slb_lists
        __props__.__dict__["slb_spec"] = slb_spec
        __props__.__dict__["spec"] = spec
        __props__.__dict__["status"] = status
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vswitch_id"] = vswitch_id
        return Gateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupVswitchId")
    def backup_vswitch_id(self) -> pulumi.Output[Optional[str]]:
        """
        The backup vswitch id.
        """
        return pulumi.get(self, "backup_vswitch_id")

    @property
    @pulumi.getter(name="deleteSlb")
    def delete_slb(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to delete the SLB purchased on behalf of the gateway at the same time.
        """
        return pulumi.get(self, "delete_slb")

    @property
    @pulumi.getter(name="enterpriseSecurityGroup")
    def enterprise_security_group(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the enterprise security group type.
        """
        return pulumi.get(self, "enterprise_security_group")

    @property
    @pulumi.getter(name="gatewayName")
    def gateway_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the Gateway .
        """
        return pulumi.get(self, "gateway_name")

    @property
    @pulumi.getter(name="internetSlbSpec")
    def internet_slb_spec(self) -> pulumi.Output[Optional[str]]:
        """
        Public network SLB specifications.
        """
        return pulumi.get(self, "internet_slb_spec")

    @property
    @pulumi.getter
    def replica(self) -> pulumi.Output[int]:
        """
        Number of Gateway Nodes.
        """
        return pulumi.get(self, "replica")

    @property
    @pulumi.getter(name="slbLists")
    def slb_lists(self) -> pulumi.Output[Sequence['outputs.GatewaySlbList']]:
        """
        A list of gateway Slb.
        """
        return pulumi.get(self, "slb_lists")

    @property
    @pulumi.getter(name="slbSpec")
    def slb_spec(self) -> pulumi.Output[Optional[str]]:
        """
        Private network SLB specifications.
        """
        return pulumi.get(self, "slb_spec")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[str]:
        """
        Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the gateway.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the vpc.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The ID of the vswitch.
        """
        return pulumi.get(self, "vswitch_id")

