# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RouteTableArgs', 'RouteTable']

@pulumi.input_type
class RouteTableArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 associate_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a RouteTable resource.
        :param pulumi.Input[str] vpc_id: The ID of VPC.
               
               The following arguments will be discarded. Please use new fields as soon as possible:
        :param pulumi.Input[str] associate_type: The type of cloud resource that is bound to the routing table. Value:
               - **VSwitch**: switch.
               - **Gateway**:IPv4 Gateway.
        :param pulumi.Input[str] description: Description of the routing table.
        :param pulumi.Input[str] name: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        :param pulumi.Input[str] route_table_name: The name of the routing table.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag.
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        if associate_type is not None:
            pulumi.set(__self__, "associate_type", associate_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if route_table_name is not None:
            pulumi.set(__self__, "route_table_name", route_table_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of VPC.

        The following arguments will be discarded. Please use new fields as soon as possible:
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="associateType")
    def associate_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of cloud resource that is bound to the routing table. Value:
        - **VSwitch**: switch.
        - **Gateway**:IPv4 Gateway.
        """
        return pulumi.get(self, "associate_type")

    @associate_type.setter
    def associate_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associate_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the routing table.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        """
        warnings.warn("""Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""", DeprecationWarning)
        pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""")

        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the routing table.
        """
        return pulumi.get(self, "route_table_name")

    @route_table_name.setter
    def route_table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tag.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _RouteTableState:
    def __init__(__self__, *,
                 associate_type: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouteTable resources.
        :param pulumi.Input[str] associate_type: The type of cloud resource that is bound to the routing table. Value:
               - **VSwitch**: switch.
               - **Gateway**:IPv4 Gateway.
        :param pulumi.Input[str] create_time: The creation time of the routing table.
        :param pulumi.Input[str] description: Description of the routing table.
        :param pulumi.Input[str] name: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[str] route_table_name: The name of the routing table.
        :param pulumi.Input[str] status: Routing table state.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag.
        :param pulumi.Input[str] vpc_id: The ID of VPC.
               
               The following arguments will be discarded. Please use new fields as soon as possible:
        """
        if associate_type is not None:
            pulumi.set(__self__, "associate_type", associate_type)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if route_table_name is not None:
            pulumi.set(__self__, "route_table_name", route_table_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="associateType")
    def associate_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of cloud resource that is bound to the routing table. Value:
        - **VSwitch**: switch.
        - **Gateway**:IPv4 Gateway.
        """
        return pulumi.get(self, "associate_type")

    @associate_type.setter
    def associate_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associate_type", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the routing table.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the routing table.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        """
        warnings.warn("""Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""", DeprecationWarning)
        pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""")

        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the routing table.
        """
        return pulumi.get(self, "route_table_name")

    @route_table_name.setter
    def route_table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Routing table state.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tag.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of VPC.

        The following arguments will be discarded. Please use new fields as soon as possible:
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class RouteTable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associate_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VPC Route Table resource. Currently, customized route tables are available in most regions apart from China (Beijing), China (Hangzhou), and China (Shenzhen) regions.

        For information about VPC Route Table and how to use it, see [What is Route Table](https://www.alibabacloud.com/help/doc-detail/87057.htm).

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_vpc = alicloud.vpc.Network("defaultVpc", vpc_name=name)
        default = alicloud.vpc.RouteTable("default",
            description="test-description",
            vpc_id=default_vpc.id,
            route_table_name=name,
            associate_type="VSwitch")
        ```

        ## Import

        VPC Route Table can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpc/routeTable:RouteTable example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] associate_type: The type of cloud resource that is bound to the routing table. Value:
               - **VSwitch**: switch.
               - **Gateway**:IPv4 Gateway.
        :param pulumi.Input[str] description: Description of the routing table.
        :param pulumi.Input[str] name: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        :param pulumi.Input[str] route_table_name: The name of the routing table.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag.
        :param pulumi.Input[str] vpc_id: The ID of VPC.
               
               The following arguments will be discarded. Please use new fields as soon as possible:
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteTableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Route Table resource. Currently, customized route tables are available in most regions apart from China (Beijing), China (Hangzhou), and China (Shenzhen) regions.

        For information about VPC Route Table and how to use it, see [What is Route Table](https://www.alibabacloud.com/help/doc-detail/87057.htm).

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_vpc = alicloud.vpc.Network("defaultVpc", vpc_name=name)
        default = alicloud.vpc.RouteTable("default",
            description="test-description",
            vpc_id=default_vpc.id,
            route_table_name=name,
            associate_type="VSwitch")
        ```

        ## Import

        VPC Route Table can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpc/routeTable:RouteTable example <id>
        ```

        :param str resource_name: The name of the resource.
        :param RouteTableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteTableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associate_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouteTableArgs.__new__(RouteTableArgs)

            __props__.__dict__["associate_type"] = associate_type
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["route_table_name"] = route_table_name
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["resource_group_id"] = None
            __props__.__dict__["status"] = None
        super(RouteTable, __self__).__init__(
            'alicloud:vpc/routeTable:RouteTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            associate_type: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            route_table_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'RouteTable':
        """
        Get an existing RouteTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] associate_type: The type of cloud resource that is bound to the routing table. Value:
               - **VSwitch**: switch.
               - **Gateway**:IPv4 Gateway.
        :param pulumi.Input[str] create_time: The creation time of the routing table.
        :param pulumi.Input[str] description: Description of the routing table.
        :param pulumi.Input[str] name: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        :param pulumi.Input[str] resource_group_id: Resource group ID.
        :param pulumi.Input[str] route_table_name: The name of the routing table.
        :param pulumi.Input[str] status: Routing table state.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag.
        :param pulumi.Input[str] vpc_id: The ID of VPC.
               
               The following arguments will be discarded. Please use new fields as soon as possible:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouteTableState.__new__(_RouteTableState)

        __props__.__dict__["associate_type"] = associate_type
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["route_table_name"] = route_table_name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        return RouteTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="associateType")
    def associate_type(self) -> pulumi.Output[str]:
        """
        The type of cloud resource that is bound to the routing table. Value:
        - **VSwitch**: switch.
        - **Gateway**:IPv4 Gateway.
        """
        return pulumi.get(self, "associate_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the routing table.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the routing table.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.
        """
        warnings.warn("""Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""", DeprecationWarning)
        pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.1. New field 'route_table_name' instead.""")

        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        Resource group ID.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> pulumi.Output[str]:
        """
        The name of the routing table.
        """
        return pulumi.get(self, "route_table_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Routing table state.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The tag.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of VPC.

        The following arguments will be discarded. Please use new fields as soon as possible:
        """
        return pulumi.get(self, "vpc_id")

