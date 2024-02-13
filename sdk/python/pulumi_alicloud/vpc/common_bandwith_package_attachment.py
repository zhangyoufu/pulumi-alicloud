# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CommonBandwithPackageAttachmentArgs', 'CommonBandwithPackageAttachment']

@pulumi.input_type
class CommonBandwithPackageAttachmentArgs:
    def __init__(__self__, *,
                 bandwidth_package_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 bandwidth_package_bandwidth: Optional[pulumi.Input[str]] = None,
                 cancel_common_bandwidth_package_ip_bandwidth: Optional[pulumi.Input[bool]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CommonBandwithPackageAttachment resource.
        :param pulumi.Input[str] bandwidth_package_id: The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        :param pulumi.Input[str] instance_id: The instance_id of the common bandwidth package attachment, the field can't be changed.
        :param pulumi.Input[str] bandwidth_package_bandwidth: The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        :param pulumi.Input[bool] cancel_common_bandwidth_package_ip_bandwidth: Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        :param pulumi.Input[str] ip_type: IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        """
        pulumi.set(__self__, "bandwidth_package_id", bandwidth_package_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if bandwidth_package_bandwidth is not None:
            pulumi.set(__self__, "bandwidth_package_bandwidth", bandwidth_package_bandwidth)
        if cancel_common_bandwidth_package_ip_bandwidth is not None:
            pulumi.set(__self__, "cancel_common_bandwidth_package_ip_bandwidth", cancel_common_bandwidth_package_ip_bandwidth)
        if ip_type is not None:
            pulumi.set(__self__, "ip_type", ip_type)

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> pulumi.Input[str]:
        """
        The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @bandwidth_package_id.setter
    def bandwidth_package_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bandwidth_package_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The instance_id of the common bandwidth package attachment, the field can't be changed.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="bandwidthPackageBandwidth")
    def bandwidth_package_bandwidth(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        """
        return pulumi.get(self, "bandwidth_package_bandwidth")

    @bandwidth_package_bandwidth.setter
    def bandwidth_package_bandwidth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_bandwidth", value)

    @property
    @pulumi.getter(name="cancelCommonBandwidthPackageIpBandwidth")
    def cancel_common_bandwidth_package_ip_bandwidth(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        """
        return pulumi.get(self, "cancel_common_bandwidth_package_ip_bandwidth")

    @cancel_common_bandwidth_package_ip_bandwidth.setter
    def cancel_common_bandwidth_package_ip_bandwidth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cancel_common_bandwidth_package_ip_bandwidth", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[pulumi.Input[str]]:
        """
        IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_type", value)


@pulumi.input_type
class _CommonBandwithPackageAttachmentState:
    def __init__(__self__, *,
                 bandwidth_package_bandwidth: Optional[pulumi.Input[str]] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 cancel_common_bandwidth_package_ip_bandwidth: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CommonBandwithPackageAttachment resources.
        :param pulumi.Input[str] bandwidth_package_bandwidth: The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        :param pulumi.Input[str] bandwidth_package_id: The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        :param pulumi.Input[bool] cancel_common_bandwidth_package_ip_bandwidth: Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        :param pulumi.Input[str] instance_id: The instance_id of the common bandwidth package attachment, the field can't be changed.
        :param pulumi.Input[str] ip_type: IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        :param pulumi.Input[str] status: The status of the Internet Shared Bandwidth instance.
        """
        if bandwidth_package_bandwidth is not None:
            pulumi.set(__self__, "bandwidth_package_bandwidth", bandwidth_package_bandwidth)
        if bandwidth_package_id is not None:
            pulumi.set(__self__, "bandwidth_package_id", bandwidth_package_id)
        if cancel_common_bandwidth_package_ip_bandwidth is not None:
            pulumi.set(__self__, "cancel_common_bandwidth_package_ip_bandwidth", cancel_common_bandwidth_package_ip_bandwidth)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if ip_type is not None:
            pulumi.set(__self__, "ip_type", ip_type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="bandwidthPackageBandwidth")
    def bandwidth_package_bandwidth(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        """
        return pulumi.get(self, "bandwidth_package_bandwidth")

    @bandwidth_package_bandwidth.setter
    def bandwidth_package_bandwidth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_bandwidth", value)

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> Optional[pulumi.Input[str]]:
        """
        The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @bandwidth_package_id.setter
    def bandwidth_package_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_id", value)

    @property
    @pulumi.getter(name="cancelCommonBandwidthPackageIpBandwidth")
    def cancel_common_bandwidth_package_ip_bandwidth(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        """
        return pulumi.get(self, "cancel_common_bandwidth_package_ip_bandwidth")

    @cancel_common_bandwidth_package_ip_bandwidth.setter
    def cancel_common_bandwidth_package_ip_bandwidth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cancel_common_bandwidth_package_ip_bandwidth", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance_id of the common bandwidth package attachment, the field can't be changed.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[pulumi.Input[str]]:
        """
        IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the Internet Shared Bandwidth instance.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class CommonBandwithPackageAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_package_bandwidth: Optional[pulumi.Input[str]] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 cancel_common_bandwidth_package_ip_bandwidth: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_common_bandwith_package = alicloud.vpc.CommonBandwithPackage("defaultCommonBandwithPackage",
            bandwidth="3",
            internet_charge_type="PayByBandwidth")
        default_eip_address = alicloud.ecs.EipAddress("defaultEipAddress",
            bandwidth="3",
            internet_charge_type="PayByTraffic")
        default_common_bandwith_package_attachment = alicloud.vpc.CommonBandwithPackageAttachment("defaultCommonBandwithPackageAttachment",
            bandwidth_package_id=default_common_bandwith_package.id,
            instance_id=default_eip_address.id,
            bandwidth_package_bandwidth="2",
            ip_type="EIP")
        ```

        ## Import

        cbwp Common Bandwidth Package Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment example <bandwidth_package_id>:<instance_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_package_bandwidth: The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        :param pulumi.Input[str] bandwidth_package_id: The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        :param pulumi.Input[bool] cancel_common_bandwidth_package_ip_bandwidth: Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        :param pulumi.Input[str] instance_id: The instance_id of the common bandwidth package attachment, the field can't be changed.
        :param pulumi.Input[str] ip_type: IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CommonBandwithPackageAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_common_bandwith_package = alicloud.vpc.CommonBandwithPackage("defaultCommonBandwithPackage",
            bandwidth="3",
            internet_charge_type="PayByBandwidth")
        default_eip_address = alicloud.ecs.EipAddress("defaultEipAddress",
            bandwidth="3",
            internet_charge_type="PayByTraffic")
        default_common_bandwith_package_attachment = alicloud.vpc.CommonBandwithPackageAttachment("defaultCommonBandwithPackageAttachment",
            bandwidth_package_id=default_common_bandwith_package.id,
            instance_id=default_eip_address.id,
            bandwidth_package_bandwidth="2",
            ip_type="EIP")
        ```

        ## Import

        cbwp Common Bandwidth Package Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment example <bandwidth_package_id>:<instance_id>
        ```

        :param str resource_name: The name of the resource.
        :param CommonBandwithPackageAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CommonBandwithPackageAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_package_bandwidth: Optional[pulumi.Input[str]] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 cancel_common_bandwidth_package_ip_bandwidth: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CommonBandwithPackageAttachmentArgs.__new__(CommonBandwithPackageAttachmentArgs)

            __props__.__dict__["bandwidth_package_bandwidth"] = bandwidth_package_bandwidth
            if bandwidth_package_id is None and not opts.urn:
                raise TypeError("Missing required property 'bandwidth_package_id'")
            __props__.__dict__["bandwidth_package_id"] = bandwidth_package_id
            __props__.__dict__["cancel_common_bandwidth_package_ip_bandwidth"] = cancel_common_bandwidth_package_ip_bandwidth
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["ip_type"] = ip_type
            __props__.__dict__["status"] = None
        super(CommonBandwithPackageAttachment, __self__).__init__(
            'alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth_package_bandwidth: Optional[pulumi.Input[str]] = None,
            bandwidth_package_id: Optional[pulumi.Input[str]] = None,
            cancel_common_bandwidth_package_ip_bandwidth: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            ip_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'CommonBandwithPackageAttachment':
        """
        Get an existing CommonBandwithPackageAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_package_bandwidth: The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        :param pulumi.Input[str] bandwidth_package_id: The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        :param pulumi.Input[bool] cancel_common_bandwidth_package_ip_bandwidth: Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        :param pulumi.Input[str] instance_id: The instance_id of the common bandwidth package attachment, the field can't be changed.
        :param pulumi.Input[str] ip_type: IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        :param pulumi.Input[str] status: The status of the Internet Shared Bandwidth instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CommonBandwithPackageAttachmentState.__new__(_CommonBandwithPackageAttachmentState)

        __props__.__dict__["bandwidth_package_bandwidth"] = bandwidth_package_bandwidth
        __props__.__dict__["bandwidth_package_id"] = bandwidth_package_id
        __props__.__dict__["cancel_common_bandwidth_package_ip_bandwidth"] = cancel_common_bandwidth_package_ip_bandwidth
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["ip_type"] = ip_type
        __props__.__dict__["status"] = status
        return CommonBandwithPackageAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bandwidthPackageBandwidth")
    def bandwidth_package_bandwidth(self) -> pulumi.Output[str]:
        """
        The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
        """
        return pulumi.get(self, "bandwidth_package_bandwidth")

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> pulumi.Output[str]:
        """
        The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @property
    @pulumi.getter(name="cancelCommonBandwidthPackageIpBandwidth")
    def cancel_common_bandwidth_package_ip_bandwidth(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
        """
        return pulumi.get(self, "cancel_common_bandwidth_package_ip_bandwidth")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The instance_id of the common bandwidth package attachment, the field can't be changed.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Output[Optional[str]]:
        """
        IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
        """
        return pulumi.get(self, "ip_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the Internet Shared Bandwidth instance.
        """
        return pulumi.get(self, "status")

