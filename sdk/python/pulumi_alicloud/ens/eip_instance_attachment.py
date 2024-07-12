# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EipInstanceAttachmentArgs', 'EipInstanceAttachment']

@pulumi.input_type
class EipInstanceAttachmentArgs:
    def __init__(__self__, *,
                 allocation_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 instance_type: Optional[pulumi.Input[str]] = None,
                 standby: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a EipInstanceAttachment resource.
        :param pulumi.Input[str] allocation_id: The first ID of the resource
        :param pulumi.Input[str] instance_id: Instance ID
        :param pulumi.Input[str] instance_type: The type of the EIP instance. Value:
               - `Nat`:NAT gateway.
               - `SlbInstance`: Server Load Balancer (ELB).
               - `NetworkInterface`: Secondary ENI.
               - `EnsInstance` (default): The ENS instance.
        :param pulumi.Input[bool] standby: Indicates whether the EIP is a backup EIP. Value:
               - true: Spare.
               - false: not standby.
        """
        pulumi.set(__self__, "allocation_id", allocation_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if standby is not None:
            pulumi.set(__self__, "standby", standby)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> pulumi.Input[str]:
        """
        The first ID of the resource
        """
        return pulumi.get(self, "allocation_id")

    @allocation_id.setter
    def allocation_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "allocation_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the EIP instance. Value:
        - `Nat`:NAT gateway.
        - `SlbInstance`: Server Load Balancer (ELB).
        - `NetworkInterface`: Secondary ENI.
        - `EnsInstance` (default): The ENS instance.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter
    def standby(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the EIP is a backup EIP. Value:
        - true: Spare.
        - false: not standby.
        """
        return pulumi.get(self, "standby")

    @standby.setter
    def standby(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "standby", value)


@pulumi.input_type
class _EipInstanceAttachmentState:
    def __init__(__self__, *,
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 standby: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EipInstanceAttachment resources.
        :param pulumi.Input[str] allocation_id: The first ID of the resource
        :param pulumi.Input[str] instance_id: Instance ID
        :param pulumi.Input[str] instance_type: The type of the EIP instance. Value:
               - `Nat`:NAT gateway.
               - `SlbInstance`: Server Load Balancer (ELB).
               - `NetworkInterface`: Secondary ENI.
               - `EnsInstance` (default): The ENS instance.
        :param pulumi.Input[bool] standby: Indicates whether the EIP is a backup EIP. Value:
               - true: Spare.
               - false: not standby.
        :param pulumi.Input[str] status: The status of the EIP.
        """
        if allocation_id is not None:
            pulumi.set(__self__, "allocation_id", allocation_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if standby is not None:
            pulumi.set(__self__, "standby", standby)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The first ID of the resource
        """
        return pulumi.get(self, "allocation_id")

    @allocation_id.setter
    def allocation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allocation_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the EIP instance. Value:
        - `Nat`:NAT gateway.
        - `SlbInstance`: Server Load Balancer (ELB).
        - `NetworkInterface`: Secondary ENI.
        - `EnsInstance` (default): The ENS instance.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter
    def standby(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the EIP is a backup EIP. Value:
        - true: Spare.
        - false: not standby.
        """
        return pulumi.get(self, "standby")

    @standby.setter
    def standby(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "standby", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the EIP.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class EipInstanceAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 standby: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a Ens Eip Instance Attachment resource.

        Bind an EIP to an instance.

        For information about Ens Eip Instance Attachment and how to use it, see [What is Eip Instance Attachment](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.227.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        ens_region_id = config.get("ensRegionId")
        if ens_region_id is None:
            ens_region_id = "cn-chenzhou-telecom_unicom_cmcc"
        default_x_kjq1_w = alicloud.ens.Instance("defaultXKjq1W",
            system_disk=alicloud.ens.InstanceSystemDiskArgs(
                size=20,
                category="cloud_efficiency",
            ),
            scheduling_strategy="Concentrate",
            schedule_area_level="Region",
            image_id="centos_6_08_64_20G_alibase_20171208",
            payment_type="Subscription",
            instance_type="ens.sn1.stiny",
            password="12345678abcABC",
            status="Running",
            amount=1,
            internet_charge_type="95BandwidthByMonth",
            instance_name=name,
            auto_use_coupon="true",
            instance_charge_strategy="PriceHighPriority",
            ens_region_id=ens_region_id,
            period_unit="Month")
        defaults_gs_n4e = alicloud.ens.Eip("defaultsGsN4e",
            bandwidth=5,
            eip_name=name,
            ens_region_id=ens_region_id,
            internet_charge_type="95BandwidthByMonth",
            payment_type="PayAsYouGo")
        default = alicloud.ens.EipInstanceAttachment("default",
            instance_id=default_x_kjq1_w.id,
            allocation_id=defaults_gs_n4e.id,
            instance_type="EnsInstance",
            standby=False)
        ```

        ## Import

        Ens Eip Instance Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ens/eipInstanceAttachment:EipInstanceAttachment example <allocation_id>:<instance_id>:<instance_type>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_id: The first ID of the resource
        :param pulumi.Input[str] instance_id: Instance ID
        :param pulumi.Input[str] instance_type: The type of the EIP instance. Value:
               - `Nat`:NAT gateway.
               - `SlbInstance`: Server Load Balancer (ELB).
               - `NetworkInterface`: Secondary ENI.
               - `EnsInstance` (default): The ENS instance.
        :param pulumi.Input[bool] standby: Indicates whether the EIP is a backup EIP. Value:
               - true: Spare.
               - false: not standby.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EipInstanceAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Ens Eip Instance Attachment resource.

        Bind an EIP to an instance.

        For information about Ens Eip Instance Attachment and how to use it, see [What is Eip Instance Attachment](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.227.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        ens_region_id = config.get("ensRegionId")
        if ens_region_id is None:
            ens_region_id = "cn-chenzhou-telecom_unicom_cmcc"
        default_x_kjq1_w = alicloud.ens.Instance("defaultXKjq1W",
            system_disk=alicloud.ens.InstanceSystemDiskArgs(
                size=20,
                category="cloud_efficiency",
            ),
            scheduling_strategy="Concentrate",
            schedule_area_level="Region",
            image_id="centos_6_08_64_20G_alibase_20171208",
            payment_type="Subscription",
            instance_type="ens.sn1.stiny",
            password="12345678abcABC",
            status="Running",
            amount=1,
            internet_charge_type="95BandwidthByMonth",
            instance_name=name,
            auto_use_coupon="true",
            instance_charge_strategy="PriceHighPriority",
            ens_region_id=ens_region_id,
            period_unit="Month")
        defaults_gs_n4e = alicloud.ens.Eip("defaultsGsN4e",
            bandwidth=5,
            eip_name=name,
            ens_region_id=ens_region_id,
            internet_charge_type="95BandwidthByMonth",
            payment_type="PayAsYouGo")
        default = alicloud.ens.EipInstanceAttachment("default",
            instance_id=default_x_kjq1_w.id,
            allocation_id=defaults_gs_n4e.id,
            instance_type="EnsInstance",
            standby=False)
        ```

        ## Import

        Ens Eip Instance Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ens/eipInstanceAttachment:EipInstanceAttachment example <allocation_id>:<instance_id>:<instance_type>
        ```

        :param str resource_name: The name of the resource.
        :param EipInstanceAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EipInstanceAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 standby: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EipInstanceAttachmentArgs.__new__(EipInstanceAttachmentArgs)

            if allocation_id is None and not opts.urn:
                raise TypeError("Missing required property 'allocation_id'")
            __props__.__dict__["allocation_id"] = allocation_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["standby"] = standby
            __props__.__dict__["status"] = None
        super(EipInstanceAttachment, __self__).__init__(
            'alicloud:ens/eipInstanceAttachment:EipInstanceAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            standby: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'EipInstanceAttachment':
        """
        Get an existing EipInstanceAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_id: The first ID of the resource
        :param pulumi.Input[str] instance_id: Instance ID
        :param pulumi.Input[str] instance_type: The type of the EIP instance. Value:
               - `Nat`:NAT gateway.
               - `SlbInstance`: Server Load Balancer (ELB).
               - `NetworkInterface`: Secondary ENI.
               - `EnsInstance` (default): The ENS instance.
        :param pulumi.Input[bool] standby: Indicates whether the EIP is a backup EIP. Value:
               - true: Spare.
               - false: not standby.
        :param pulumi.Input[str] status: The status of the EIP.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EipInstanceAttachmentState.__new__(_EipInstanceAttachmentState)

        __props__.__dict__["allocation_id"] = allocation_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instance_type"] = instance_type
        __props__.__dict__["standby"] = standby
        __props__.__dict__["status"] = status
        return EipInstanceAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> pulumi.Output[str]:
        """
        The first ID of the resource
        """
        return pulumi.get(self, "allocation_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        The type of the EIP instance. Value:
        - `Nat`:NAT gateway.
        - `SlbInstance`: Server Load Balancer (ELB).
        - `NetworkInterface`: Secondary ENI.
        - `EnsInstance` (default): The ENS instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def standby(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether the EIP is a backup EIP. Value:
        - true: Spare.
        - false: not standby.
        """
        return pulumi.get(self, "standby")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the EIP.
        """
        return pulumi.get(self, "status")

