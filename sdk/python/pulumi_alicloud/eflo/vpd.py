# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpdArgs', 'Vpd']

@pulumi.input_type
class VpdArgs:
    def __init__(__self__, *,
                 cidr: pulumi.Input[str],
                 vpd_name: pulumi.Input[str],
                 resource_group_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Vpd resource.
        :param pulumi.Input[str] cidr: CIDR network segment.
        :param pulumi.Input[str] vpd_name: The Name of the VPD.
        :param pulumi.Input[str] resource_group_id: The Resource group id.
        """
        pulumi.set(__self__, "cidr", cidr)
        pulumi.set(__self__, "vpd_name", vpd_name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Input[str]:
        """
        CIDR network segment.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="vpdName")
    def vpd_name(self) -> pulumi.Input[str]:
        """
        The Name of the VPD.
        """
        return pulumi.get(self, "vpd_name")

    @vpd_name.setter
    def vpd_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpd_name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Resource group id.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)


@pulumi.input_type
class _VpdState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 gmt_modified: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpd_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Vpd resources.
        :param pulumi.Input[str] cidr: CIDR network segment.
        :param pulumi.Input[str] create_time: The creation time of the resource
        :param pulumi.Input[str] gmt_modified: Modification time
        :param pulumi.Input[str] resource_group_id: The Resource group id.
        :param pulumi.Input[str] status: The Vpd status.
        :param pulumi.Input[str] vpd_name: The Name of the VPD.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if gmt_modified is not None:
            pulumi.set(__self__, "gmt_modified", gmt_modified)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpd_name is not None:
            pulumi.set(__self__, "vpd_name", vpd_name)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        CIDR network segment.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the resource
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="gmtModified")
    def gmt_modified(self) -> Optional[pulumi.Input[str]]:
        """
        Modification time
        """
        return pulumi.get(self, "gmt_modified")

    @gmt_modified.setter
    def gmt_modified(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gmt_modified", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Resource group id.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The Vpd status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpdName")
    def vpd_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the VPD.
        """
        return pulumi.get(self, "vpd_name")

    @vpd_name.setter
    def vpd_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpd_name", value)


class Vpd(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 vpd_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Eflo Vpd resource.

        For information about Eflo Vpd and how to use it, see [What is Vpd](https://www.alibabacloud.com/help/en/pai/user-guide/overview-of-intelligent-computing-lingjun).

        > **NOTE:** Available since v1.201.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_vpd = alicloud.eflo.Vpd("defaultVpd",
            cidr="10.0.0.0/8",
            vpd_name=name,
            resource_group_id=default_resource_groups.groups[0].id)
        ```

        ## Import

        Eflo Vpd can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:eflo/vpd:Vpd example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: CIDR network segment.
        :param pulumi.Input[str] resource_group_id: The Resource group id.
        :param pulumi.Input[str] vpd_name: The Name of the VPD.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpdArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Eflo Vpd resource.

        For information about Eflo Vpd and how to use it, see [What is Vpd](https://www.alibabacloud.com/help/en/pai/user-guide/overview-of-intelligent-computing-lingjun).

        > **NOTE:** Available since v1.201.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_vpd = alicloud.eflo.Vpd("defaultVpd",
            cidr="10.0.0.0/8",
            vpd_name=name,
            resource_group_id=default_resource_groups.groups[0].id)
        ```

        ## Import

        Eflo Vpd can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:eflo/vpd:Vpd example <id>
        ```

        :param str resource_name: The name of the resource.
        :param VpdArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpdArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 vpd_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpdArgs.__new__(VpdArgs)

            if cidr is None and not opts.urn:
                raise TypeError("Missing required property 'cidr'")
            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["resource_group_id"] = resource_group_id
            if vpd_name is None and not opts.urn:
                raise TypeError("Missing required property 'vpd_name'")
            __props__.__dict__["vpd_name"] = vpd_name
            __props__.__dict__["create_time"] = None
            __props__.__dict__["gmt_modified"] = None
            __props__.__dict__["status"] = None
        super(Vpd, __self__).__init__(
            'alicloud:eflo/vpd:Vpd',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            gmt_modified: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vpd_name: Optional[pulumi.Input[str]] = None) -> 'Vpd':
        """
        Get an existing Vpd resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: CIDR network segment.
        :param pulumi.Input[str] create_time: The creation time of the resource
        :param pulumi.Input[str] gmt_modified: Modification time
        :param pulumi.Input[str] resource_group_id: The Resource group id.
        :param pulumi.Input[str] status: The Vpd status.
        :param pulumi.Input[str] vpd_name: The Name of the VPD.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpdState.__new__(_VpdState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["gmt_modified"] = gmt_modified
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["status"] = status
        __props__.__dict__["vpd_name"] = vpd_name
        return Vpd(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        CIDR network segment.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the resource
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="gmtModified")
    def gmt_modified(self) -> pulumi.Output[str]:
        """
        Modification time
        """
        return pulumi.get(self, "gmt_modified")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Resource group id.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The Vpd status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpdName")
    def vpd_name(self) -> pulumi.Output[str]:
        """
        The Name of the VPD.
        """
        return pulumi.get(self, "vpd_name")

