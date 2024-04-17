# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['Ipv4GatewayArgs', 'Ipv4Gateway']

@pulumi.input_type
class Ipv4GatewayArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ipv4_gateway_description: Optional[pulumi.Input[str]] = None,
                 ipv4_gateway_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Ipv4Gateway resource.
        :param pulumi.Input[str] vpc_id: The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        :param pulumi.Input[bool] dry_run: Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        :param pulumi.Input[bool] enabled: Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        :param pulumi.Input[str] ipv4_gateway_description: The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        :param pulumi.Input[str] ipv4_gateway_name: The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group to which the instance belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: The tags of the current resource.
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if ipv4_gateway_description is not None:
            pulumi.set(__self__, "ipv4_gateway_description", ipv4_gateway_description)
        if ipv4_gateway_name is not None:
            pulumi.set(__self__, "ipv4_gateway_name", ipv4_gateway_name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="ipv4GatewayDescription")
    def ipv4_gateway_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        """
        return pulumi.get(self, "ipv4_gateway_description")

    @ipv4_gateway_description.setter
    def ipv4_gateway_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_gateway_description", value)

    @property
    @pulumi.getter(name="ipv4GatewayName")
    def ipv4_gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        """
        return pulumi.get(self, "ipv4_gateway_name")

    @ipv4_gateway_name.setter
    def ipv4_gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_gateway_name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group to which the instance belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tags of the current resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _Ipv4GatewayState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ipv4_gateway_description: Optional[pulumi.Input[str]] = None,
                 ipv4_gateway_id: Optional[pulumi.Input[str]] = None,
                 ipv4_gateway_name: Optional[pulumi.Input[str]] = None,
                 ipv4_gateway_route_table_id: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ipv4Gateway resources.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[bool] dry_run: Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        :param pulumi.Input[bool] enabled: Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        :param pulumi.Input[str] ipv4_gateway_description: The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        :param pulumi.Input[str] ipv4_gateway_id: Resource primary key field.
        :param pulumi.Input[str] ipv4_gateway_name: The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        :param pulumi.Input[str] ipv4_gateway_route_table_id: ID of the route table associated with IPv4 Gateway.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group to which the instance belongs.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[Mapping[str, Any]] tags: The tags of the current resource.
        :param pulumi.Input[str] vpc_id: The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if ipv4_gateway_description is not None:
            pulumi.set(__self__, "ipv4_gateway_description", ipv4_gateway_description)
        if ipv4_gateway_id is not None:
            pulumi.set(__self__, "ipv4_gateway_id", ipv4_gateway_id)
        if ipv4_gateway_name is not None:
            pulumi.set(__self__, "ipv4_gateway_name", ipv4_gateway_name)
        if ipv4_gateway_route_table_id is not None:
            pulumi.set(__self__, "ipv4_gateway_route_table_id", ipv4_gateway_route_table_id)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

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
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="ipv4GatewayDescription")
    def ipv4_gateway_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        """
        return pulumi.get(self, "ipv4_gateway_description")

    @ipv4_gateway_description.setter
    def ipv4_gateway_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_gateway_description", value)

    @property
    @pulumi.getter(name="ipv4GatewayId")
    def ipv4_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource primary key field.
        """
        return pulumi.get(self, "ipv4_gateway_id")

    @ipv4_gateway_id.setter
    def ipv4_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_gateway_id", value)

    @property
    @pulumi.getter(name="ipv4GatewayName")
    def ipv4_gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        """
        return pulumi.get(self, "ipv4_gateway_name")

    @ipv4_gateway_name.setter
    def ipv4_gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_gateway_name", value)

    @property
    @pulumi.getter(name="ipv4GatewayRouteTableId")
    def ipv4_gateway_route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the route table associated with IPv4 Gateway.
        """
        return pulumi.get(self, "ipv4_gateway_route_table_id")

    @ipv4_gateway_route_table_id.setter
    def ipv4_gateway_route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_gateway_route_table_id", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group to which the instance belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

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
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tags of the current resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class Ipv4Gateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ipv4_gateway_description: Optional[pulumi.Input[str]] = None,
                 ipv4_gateway_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Vpc Ipv4 Gateway resource.

        For information about Vpc Ipv4 Gateway and how to use it, see [What is Ipv4 Gateway](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createipv4gateway).

        > **NOTE:** Available in v1.181.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testacc-example"
        default = alicloud.resourcemanager.ResourceGroup("default",
            display_name="tf-testAcc-rg665",
            resource_group_name=name)
        modify = alicloud.resourcemanager.ResourceGroup("modify",
            display_name="tf-testAcc-rg298",
            resource_group_name=f"{name}1")
        default_network = alicloud.vpc.Network("default",
            vpc_name=f"{name}2",
            cidr_block="10.0.0.0/8")
        default_ipv4_gateway = alicloud.vpc.Ipv4Gateway("default",
            ipv4_gateway_name=name,
            ipv4_gateway_description="tf-testAcc-Ipv4Gateway",
            resource_group_id=default.id,
            vpc_id=default_network.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Vpc Ipv4 Gateway can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpc/ipv4Gateway:Ipv4Gateway example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dry_run: Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        :param pulumi.Input[bool] enabled: Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        :param pulumi.Input[str] ipv4_gateway_description: The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        :param pulumi.Input[str] ipv4_gateway_name: The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group to which the instance belongs.
        :param pulumi.Input[Mapping[str, Any]] tags: The tags of the current resource.
        :param pulumi.Input[str] vpc_id: The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ipv4GatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Vpc Ipv4 Gateway resource.

        For information about Vpc Ipv4 Gateway and how to use it, see [What is Ipv4 Gateway](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createipv4gateway).

        > **NOTE:** Available in v1.181.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testacc-example"
        default = alicloud.resourcemanager.ResourceGroup("default",
            display_name="tf-testAcc-rg665",
            resource_group_name=name)
        modify = alicloud.resourcemanager.ResourceGroup("modify",
            display_name="tf-testAcc-rg298",
            resource_group_name=f"{name}1")
        default_network = alicloud.vpc.Network("default",
            vpc_name=f"{name}2",
            cidr_block="10.0.0.0/8")
        default_ipv4_gateway = alicloud.vpc.Ipv4Gateway("default",
            ipv4_gateway_name=name,
            ipv4_gateway_description="tf-testAcc-Ipv4Gateway",
            resource_group_id=default.id,
            vpc_id=default_network.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Vpc Ipv4 Gateway can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpc/ipv4Gateway:Ipv4Gateway example <id>
        ```

        :param str resource_name: The name of the resource.
        :param Ipv4GatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ipv4GatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ipv4_gateway_description: Optional[pulumi.Input[str]] = None,
                 ipv4_gateway_name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Ipv4GatewayArgs.__new__(Ipv4GatewayArgs)

            __props__.__dict__["dry_run"] = dry_run
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["ipv4_gateway_description"] = ipv4_gateway_description
            __props__.__dict__["ipv4_gateway_name"] = ipv4_gateway_name
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["ipv4_gateway_id"] = None
            __props__.__dict__["ipv4_gateway_route_table_id"] = None
            __props__.__dict__["status"] = None
        super(Ipv4Gateway, __self__).__init__(
            'alicloud:vpc/ipv4Gateway:Ipv4Gateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            ipv4_gateway_description: Optional[pulumi.Input[str]] = None,
            ipv4_gateway_id: Optional[pulumi.Input[str]] = None,
            ipv4_gateway_name: Optional[pulumi.Input[str]] = None,
            ipv4_gateway_route_table_id: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'Ipv4Gateway':
        """
        Get an existing Ipv4Gateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The creation time of the resource.
        :param pulumi.Input[bool] dry_run: Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        :param pulumi.Input[bool] enabled: Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        :param pulumi.Input[str] ipv4_gateway_description: The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        :param pulumi.Input[str] ipv4_gateway_id: Resource primary key field.
        :param pulumi.Input[str] ipv4_gateway_name: The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        :param pulumi.Input[str] ipv4_gateway_route_table_id: ID of the route table associated with IPv4 Gateway.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group to which the instance belongs.
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[Mapping[str, Any]] tags: The tags of the current resource.
        :param pulumi.Input[str] vpc_id: The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ipv4GatewayState.__new__(_Ipv4GatewayState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["ipv4_gateway_description"] = ipv4_gateway_description
        __props__.__dict__["ipv4_gateway_id"] = ipv4_gateway_id
        __props__.__dict__["ipv4_gateway_name"] = ipv4_gateway_name
        __props__.__dict__["ipv4_gateway_route_table_id"] = ipv4_gateway_route_table_id
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        return Ipv4Gateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="ipv4GatewayDescription")
    def ipv4_gateway_description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        """
        return pulumi.get(self, "ipv4_gateway_description")

    @property
    @pulumi.getter(name="ipv4GatewayId")
    def ipv4_gateway_id(self) -> pulumi.Output[str]:
        """
        Resource primary key field.
        """
        return pulumi.get(self, "ipv4_gateway_id")

    @property
    @pulumi.getter(name="ipv4GatewayName")
    def ipv4_gateway_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        """
        return pulumi.get(self, "ipv4_gateway_name")

    @property
    @pulumi.getter(name="ipv4GatewayRouteTableId")
    def ipv4_gateway_route_table_id(self) -> pulumi.Output[str]:
        """
        ID of the route table associated with IPv4 Gateway.
        """
        return pulumi.get(self, "ipv4_gateway_route_table_id")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource group to which the instance belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The tags of the current resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        """
        return pulumi.get(self, "vpc_id")

