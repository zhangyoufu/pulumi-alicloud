# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BasicEndpointGroupArgs', 'BasicEndpointGroup']

@pulumi.input_type
class BasicEndpointGroupArgs:
    def __init__(__self__, *,
                 accelerator_id: pulumi.Input[str],
                 endpoint_group_region: pulumi.Input[str],
                 basic_endpoint_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint_address: Optional[pulumi.Input[str]] = None,
                 endpoint_sub_address: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BasicEndpointGroup resource.
        :param pulumi.Input[str] accelerator_id: The ID of the basic GA instance.
        :param pulumi.Input[str] endpoint_group_region: The ID of the region where you want to create the endpoint group.
        :param pulumi.Input[str] basic_endpoint_group_name: The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        :param pulumi.Input[str] description: The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        :param pulumi.Input[str] endpoint_address: The address of the endpoint.
        :param pulumi.Input[str] endpoint_sub_address: The sub address of the endpoint.
        :param pulumi.Input[str] endpoint_type: The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        """
        pulumi.set(__self__, "accelerator_id", accelerator_id)
        pulumi.set(__self__, "endpoint_group_region", endpoint_group_region)
        if basic_endpoint_group_name is not None:
            pulumi.set(__self__, "basic_endpoint_group_name", basic_endpoint_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if endpoint_address is not None:
            pulumi.set(__self__, "endpoint_address", endpoint_address)
        if endpoint_sub_address is not None:
            pulumi.set(__self__, "endpoint_sub_address", endpoint_sub_address)
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)

    @property
    @pulumi.getter(name="acceleratorId")
    def accelerator_id(self) -> pulumi.Input[str]:
        """
        The ID of the basic GA instance.
        """
        return pulumi.get(self, "accelerator_id")

    @accelerator_id.setter
    def accelerator_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "accelerator_id", value)

    @property
    @pulumi.getter(name="endpointGroupRegion")
    def endpoint_group_region(self) -> pulumi.Input[str]:
        """
        The ID of the region where you want to create the endpoint group.
        """
        return pulumi.get(self, "endpoint_group_region")

    @endpoint_group_region.setter
    def endpoint_group_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_group_region", value)

    @property
    @pulumi.getter(name="basicEndpointGroupName")
    def basic_endpoint_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        """
        return pulumi.get(self, "basic_endpoint_group_name")

    @basic_endpoint_group_name.setter
    def basic_endpoint_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_endpoint_group_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endpointAddress")
    def endpoint_address(self) -> Optional[pulumi.Input[str]]:
        """
        The address of the endpoint.
        """
        return pulumi.get(self, "endpoint_address")

    @endpoint_address.setter
    def endpoint_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_address", value)

    @property
    @pulumi.getter(name="endpointSubAddress")
    def endpoint_sub_address(self) -> Optional[pulumi.Input[str]]:
        """
        The sub address of the endpoint.
        """
        return pulumi.get(self, "endpoint_sub_address")

    @endpoint_sub_address.setter
    def endpoint_sub_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_sub_address", value)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_type", value)


@pulumi.input_type
class _BasicEndpointGroupState:
    def __init__(__self__, *,
                 accelerator_id: Optional[pulumi.Input[str]] = None,
                 basic_endpoint_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint_address: Optional[pulumi.Input[str]] = None,
                 endpoint_group_region: Optional[pulumi.Input[str]] = None,
                 endpoint_sub_address: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BasicEndpointGroup resources.
        :param pulumi.Input[str] accelerator_id: The ID of the basic GA instance.
        :param pulumi.Input[str] basic_endpoint_group_name: The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        :param pulumi.Input[str] description: The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        :param pulumi.Input[str] endpoint_address: The address of the endpoint.
        :param pulumi.Input[str] endpoint_group_region: The ID of the region where you want to create the endpoint group.
        :param pulumi.Input[str] endpoint_sub_address: The sub address of the endpoint.
        :param pulumi.Input[str] endpoint_type: The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        :param pulumi.Input[str] status: The status of the Basic Endpoint Group.
        """
        if accelerator_id is not None:
            pulumi.set(__self__, "accelerator_id", accelerator_id)
        if basic_endpoint_group_name is not None:
            pulumi.set(__self__, "basic_endpoint_group_name", basic_endpoint_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if endpoint_address is not None:
            pulumi.set(__self__, "endpoint_address", endpoint_address)
        if endpoint_group_region is not None:
            pulumi.set(__self__, "endpoint_group_region", endpoint_group_region)
        if endpoint_sub_address is not None:
            pulumi.set(__self__, "endpoint_sub_address", endpoint_sub_address)
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="acceleratorId")
    def accelerator_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the basic GA instance.
        """
        return pulumi.get(self, "accelerator_id")

    @accelerator_id.setter
    def accelerator_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accelerator_id", value)

    @property
    @pulumi.getter(name="basicEndpointGroupName")
    def basic_endpoint_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        """
        return pulumi.get(self, "basic_endpoint_group_name")

    @basic_endpoint_group_name.setter
    def basic_endpoint_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_endpoint_group_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endpointAddress")
    def endpoint_address(self) -> Optional[pulumi.Input[str]]:
        """
        The address of the endpoint.
        """
        return pulumi.get(self, "endpoint_address")

    @endpoint_address.setter
    def endpoint_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_address", value)

    @property
    @pulumi.getter(name="endpointGroupRegion")
    def endpoint_group_region(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the region where you want to create the endpoint group.
        """
        return pulumi.get(self, "endpoint_group_region")

    @endpoint_group_region.setter
    def endpoint_group_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_group_region", value)

    @property
    @pulumi.getter(name="endpointSubAddress")
    def endpoint_sub_address(self) -> Optional[pulumi.Input[str]]:
        """
        The sub address of the endpoint.
        """
        return pulumi.get(self, "endpoint_sub_address")

    @endpoint_sub_address.setter
    def endpoint_sub_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_sub_address", value)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the Basic Endpoint Group.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class BasicEndpointGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerator_id: Optional[pulumi.Input[str]] = None,
                 basic_endpoint_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint_address: Optional[pulumi.Input[str]] = None,
                 endpoint_group_region: Optional[pulumi.Input[str]] = None,
                 endpoint_sub_address: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Global Accelerator (GA) Basic Endpoint Group resource.

        For information about Global Accelerator (GA) Basic Endpoint Group and how to use it, see [What is Basic Endpoint Group](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicendpointgroup).

        > **NOTE:** Available since v1.194.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        region = config.get("region")
        if region is None:
            region = "cn-hangzhou"
        endpoint_group_region = config.get("endpointGroupRegion")
        if endpoint_group_region is None:
            endpoint_group_region = "cn-beijing"
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].id)
        default_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer",
            load_balancer_name="terraform-example",
            vswitch_id=default_switch.id,
            load_balancer_spec="slb.s2.small",
            address_type="intranet")
        default_basic_accelerator = alicloud.ga.BasicAccelerator("defaultBasicAccelerator",
            duration=1,
            basic_accelerator_name="terraform-example",
            description="terraform-example",
            bandwidth_billing_type="CDT",
            auto_use_coupon="true",
            auto_pay=True)
        default_basic_endpoint_group = alicloud.ga.BasicEndpointGroup("defaultBasicEndpointGroup",
            accelerator_id=default_basic_accelerator.id,
            endpoint_group_region=endpoint_group_region,
            endpoint_type="SLB",
            endpoint_address=default_application_load_balancer.id,
            endpoint_sub_address="192.168.0.1",
            basic_endpoint_group_name="terraform-example",
            description="terraform-example")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Global Accelerator (GA) Basic Endpoint Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ga/basicEndpointGroup:BasicEndpointGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_id: The ID of the basic GA instance.
        :param pulumi.Input[str] basic_endpoint_group_name: The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        :param pulumi.Input[str] description: The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        :param pulumi.Input[str] endpoint_address: The address of the endpoint.
        :param pulumi.Input[str] endpoint_group_region: The ID of the region where you want to create the endpoint group.
        :param pulumi.Input[str] endpoint_sub_address: The sub address of the endpoint.
        :param pulumi.Input[str] endpoint_type: The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BasicEndpointGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Global Accelerator (GA) Basic Endpoint Group resource.

        For information about Global Accelerator (GA) Basic Endpoint Group and how to use it, see [What is Basic Endpoint Group](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicendpointgroup).

        > **NOTE:** Available since v1.194.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        region = config.get("region")
        if region is None:
            region = "cn-hangzhou"
        endpoint_group_region = config.get("endpointGroupRegion")
        if endpoint_group_region is None:
            endpoint_group_region = "cn-beijing"
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].id)
        default_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer",
            load_balancer_name="terraform-example",
            vswitch_id=default_switch.id,
            load_balancer_spec="slb.s2.small",
            address_type="intranet")
        default_basic_accelerator = alicloud.ga.BasicAccelerator("defaultBasicAccelerator",
            duration=1,
            basic_accelerator_name="terraform-example",
            description="terraform-example",
            bandwidth_billing_type="CDT",
            auto_use_coupon="true",
            auto_pay=True)
        default_basic_endpoint_group = alicloud.ga.BasicEndpointGroup("defaultBasicEndpointGroup",
            accelerator_id=default_basic_accelerator.id,
            endpoint_group_region=endpoint_group_region,
            endpoint_type="SLB",
            endpoint_address=default_application_load_balancer.id,
            endpoint_sub_address="192.168.0.1",
            basic_endpoint_group_name="terraform-example",
            description="terraform-example")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Global Accelerator (GA) Basic Endpoint Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ga/basicEndpointGroup:BasicEndpointGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param BasicEndpointGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BasicEndpointGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerator_id: Optional[pulumi.Input[str]] = None,
                 basic_endpoint_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint_address: Optional[pulumi.Input[str]] = None,
                 endpoint_group_region: Optional[pulumi.Input[str]] = None,
                 endpoint_sub_address: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BasicEndpointGroupArgs.__new__(BasicEndpointGroupArgs)

            if accelerator_id is None and not opts.urn:
                raise TypeError("Missing required property 'accelerator_id'")
            __props__.__dict__["accelerator_id"] = accelerator_id
            __props__.__dict__["basic_endpoint_group_name"] = basic_endpoint_group_name
            __props__.__dict__["description"] = description
            __props__.__dict__["endpoint_address"] = endpoint_address
            if endpoint_group_region is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_group_region'")
            __props__.__dict__["endpoint_group_region"] = endpoint_group_region
            __props__.__dict__["endpoint_sub_address"] = endpoint_sub_address
            __props__.__dict__["endpoint_type"] = endpoint_type
            __props__.__dict__["status"] = None
        super(BasicEndpointGroup, __self__).__init__(
            'alicloud:ga/basicEndpointGroup:BasicEndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accelerator_id: Optional[pulumi.Input[str]] = None,
            basic_endpoint_group_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            endpoint_address: Optional[pulumi.Input[str]] = None,
            endpoint_group_region: Optional[pulumi.Input[str]] = None,
            endpoint_sub_address: Optional[pulumi.Input[str]] = None,
            endpoint_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'BasicEndpointGroup':
        """
        Get an existing BasicEndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_id: The ID of the basic GA instance.
        :param pulumi.Input[str] basic_endpoint_group_name: The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        :param pulumi.Input[str] description: The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        :param pulumi.Input[str] endpoint_address: The address of the endpoint.
        :param pulumi.Input[str] endpoint_group_region: The ID of the region where you want to create the endpoint group.
        :param pulumi.Input[str] endpoint_sub_address: The sub address of the endpoint.
        :param pulumi.Input[str] endpoint_type: The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        :param pulumi.Input[str] status: The status of the Basic Endpoint Group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BasicEndpointGroupState.__new__(_BasicEndpointGroupState)

        __props__.__dict__["accelerator_id"] = accelerator_id
        __props__.__dict__["basic_endpoint_group_name"] = basic_endpoint_group_name
        __props__.__dict__["description"] = description
        __props__.__dict__["endpoint_address"] = endpoint_address
        __props__.__dict__["endpoint_group_region"] = endpoint_group_region
        __props__.__dict__["endpoint_sub_address"] = endpoint_sub_address
        __props__.__dict__["endpoint_type"] = endpoint_type
        __props__.__dict__["status"] = status
        return BasicEndpointGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceleratorId")
    def accelerator_id(self) -> pulumi.Output[str]:
        """
        The ID of the basic GA instance.
        """
        return pulumi.get(self, "accelerator_id")

    @property
    @pulumi.getter(name="basicEndpointGroupName")
    def basic_endpoint_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the endpoint group. The `basic_endpoint_group_name` must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
        """
        return pulumi.get(self, "basic_endpoint_group_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the endpoint group. The `description` cannot exceed 256 characters in length and cannot contain http:// or https://.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endpointAddress")
    def endpoint_address(self) -> pulumi.Output[str]:
        """
        The address of the endpoint.
        """
        return pulumi.get(self, "endpoint_address")

    @property
    @pulumi.getter(name="endpointGroupRegion")
    def endpoint_group_region(self) -> pulumi.Output[str]:
        """
        The ID of the region where you want to create the endpoint group.
        """
        return pulumi.get(self, "endpoint_group_region")

    @property
    @pulumi.getter(name="endpointSubAddress")
    def endpoint_sub_address(self) -> pulumi.Output[str]:
        """
        The sub address of the endpoint.
        """
        return pulumi.get(self, "endpoint_sub_address")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> pulumi.Output[str]:
        """
        The type of the endpoint. Valid values: `ENI`, `SLB` and `ECS`.
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the Basic Endpoint Group.
        """
        return pulumi.get(self, "status")

