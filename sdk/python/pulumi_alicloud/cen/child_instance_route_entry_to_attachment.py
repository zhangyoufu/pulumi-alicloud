# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ChildInstanceRouteEntryToAttachmentArgs', 'ChildInstanceRouteEntryToAttachment']

@pulumi.input_type
class ChildInstanceRouteEntryToAttachmentArgs:
    def __init__(__self__, *,
                 cen_id: pulumi.Input[str],
                 child_instance_route_table_id: pulumi.Input[str],
                 destination_cidr_block: pulumi.Input[str],
                 transit_router_attachment_id: pulumi.Input[str],
                 dry_run: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ChildInstanceRouteEntryToAttachment resource.
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] child_instance_route_table_id: The first ID of the resource
        :param pulumi.Input[str] destination_cidr_block: DestinationCidrBlock
        :param pulumi.Input[str] transit_router_attachment_id: TransitRouterAttachmentId
        :param pulumi.Input[bool] dry_run: Whether to perform pre-check on this request, including permission and instance status verification.
        """
        pulumi.set(__self__, "cen_id", cen_id)
        pulumi.set(__self__, "child_instance_route_table_id", child_instance_route_table_id)
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        pulumi.set(__self__, "transit_router_attachment_id", transit_router_attachment_id)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Input[str]:
        """
        The ID of the CEN instance.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="childInstanceRouteTableId")
    def child_instance_route_table_id(self) -> pulumi.Input[str]:
        """
        The first ID of the resource
        """
        return pulumi.get(self, "child_instance_route_table_id")

    @child_instance_route_table_id.setter
    def child_instance_route_table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "child_instance_route_table_id", value)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Input[str]:
        """
        DestinationCidrBlock
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> pulumi.Input[str]:
        """
        TransitRouterAttachmentId
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @transit_router_attachment_id.setter
    def transit_router_attachment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_router_attachment_id", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to perform pre-check on this request, including permission and instance status verification.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)


@pulumi.input_type
class _ChildInstanceRouteEntryToAttachmentState:
    def __init__(__self__, *,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 child_instance_route_table_id: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 service_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ChildInstanceRouteEntryToAttachment resources.
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] child_instance_route_table_id: The first ID of the resource
        :param pulumi.Input[str] destination_cidr_block: DestinationCidrBlock
        :param pulumi.Input[bool] dry_run: Whether to perform pre-check on this request, including permission and instance status verification.
        :param pulumi.Input[str] service_type: ServiceType
        :param pulumi.Input[str] status: The status of the resource
        :param pulumi.Input[str] transit_router_attachment_id: TransitRouterAttachmentId
        """
        if cen_id is not None:
            pulumi.set(__self__, "cen_id", cen_id)
        if child_instance_route_table_id is not None:
            pulumi.set(__self__, "child_instance_route_table_id", child_instance_route_table_id)
        if destination_cidr_block is not None:
            pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if service_type is not None:
            pulumi.set(__self__, "service_type", service_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if transit_router_attachment_id is not None:
            pulumi.set(__self__, "transit_router_attachment_id", transit_router_attachment_id)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the CEN instance.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="childInstanceRouteTableId")
    def child_instance_route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The first ID of the resource
        """
        return pulumi.get(self, "child_instance_route_table_id")

    @child_instance_route_table_id.setter
    def child_instance_route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "child_instance_route_table_id", value)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        DestinationCidrBlock
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to perform pre-check on this request, including permission and instance status verification.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> Optional[pulumi.Input[str]]:
        """
        ServiceType
        """
        return pulumi.get(self, "service_type")

    @service_type.setter
    def service_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        TransitRouterAttachmentId
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @transit_router_attachment_id.setter
    def transit_router_attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_attachment_id", value)


class ChildInstanceRouteEntryToAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 child_instance_route_table_id: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cen Child Instance Route Entry To Attachment resource.

        For information about Cen Child Instance Route Entry To Attachment and how to use it, see [What is Child Instance Route Entry To Attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createcenchildinstancerouteentrytoattachment).

        > **NOTE:** Available since v1.195.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.cen.get_transit_router_available_resources()
        master_zone = default.resources[0].master_zones[0]
        slave_zone = default.resources[0].slave_zones[1]
        example = alicloud.vpc.Network("example",
            vpc_name=name,
            cidr_block="192.168.0.0/16")
        example_master = alicloud.vpc.Switch("example_master",
            vswitch_name=name,
            cidr_block="192.168.1.0/24",
            vpc_id=example.id,
            zone_id=master_zone)
        example_slave = alicloud.vpc.Switch("example_slave",
            vswitch_name=name,
            cidr_block="192.168.2.0/24",
            vpc_id=example.id,
            zone_id=slave_zone)
        example_instance = alicloud.cen.Instance("example",
            cen_instance_name=name,
            protection_level="REDUCED")
        example_transit_router = alicloud.cen.TransitRouter("example",
            transit_router_name=name,
            cen_id=example_instance.id)
        example_transit_router_vpc_attachment = alicloud.cen.TransitRouterVpcAttachment("example",
            cen_id=example_instance.id,
            transit_router_id=example_transit_router.transit_router_id,
            vpc_id=example.id,
            zone_mappings=[
                alicloud.cen.TransitRouterVpcAttachmentZoneMappingArgs(
                    zone_id=master_zone,
                    vswitch_id=example_master.id,
                ),
                alicloud.cen.TransitRouterVpcAttachmentZoneMappingArgs(
                    zone_id=slave_zone,
                    vswitch_id=example_slave.id,
                ),
            ],
            transit_router_attachment_name=name,
            transit_router_attachment_description=name)
        example_route_table = alicloud.vpc.RouteTable("example",
            vpc_id=example.id,
            route_table_name=name,
            description=name)
        example_child_instance_route_entry_to_attachment = alicloud.cen.ChildInstanceRouteEntryToAttachment("example",
            transit_router_attachment_id=example_transit_router_vpc_attachment.transit_router_attachment_id,
            cen_id=example_instance.id,
            destination_cidr_block="10.0.0.0/24",
            child_instance_route_table_id=example_route_table.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Cen Child Instance Route Entry To Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment example <cen_id>:<child_instance_route_table_id>:<transit_router_attachment_id>:<destination_cidr_block>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] child_instance_route_table_id: The first ID of the resource
        :param pulumi.Input[str] destination_cidr_block: DestinationCidrBlock
        :param pulumi.Input[bool] dry_run: Whether to perform pre-check on this request, including permission and instance status verification.
        :param pulumi.Input[str] transit_router_attachment_id: TransitRouterAttachmentId
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ChildInstanceRouteEntryToAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cen Child Instance Route Entry To Attachment resource.

        For information about Cen Child Instance Route Entry To Attachment and how to use it, see [What is Child Instance Route Entry To Attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createcenchildinstancerouteentrytoattachment).

        > **NOTE:** Available since v1.195.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.cen.get_transit_router_available_resources()
        master_zone = default.resources[0].master_zones[0]
        slave_zone = default.resources[0].slave_zones[1]
        example = alicloud.vpc.Network("example",
            vpc_name=name,
            cidr_block="192.168.0.0/16")
        example_master = alicloud.vpc.Switch("example_master",
            vswitch_name=name,
            cidr_block="192.168.1.0/24",
            vpc_id=example.id,
            zone_id=master_zone)
        example_slave = alicloud.vpc.Switch("example_slave",
            vswitch_name=name,
            cidr_block="192.168.2.0/24",
            vpc_id=example.id,
            zone_id=slave_zone)
        example_instance = alicloud.cen.Instance("example",
            cen_instance_name=name,
            protection_level="REDUCED")
        example_transit_router = alicloud.cen.TransitRouter("example",
            transit_router_name=name,
            cen_id=example_instance.id)
        example_transit_router_vpc_attachment = alicloud.cen.TransitRouterVpcAttachment("example",
            cen_id=example_instance.id,
            transit_router_id=example_transit_router.transit_router_id,
            vpc_id=example.id,
            zone_mappings=[
                alicloud.cen.TransitRouterVpcAttachmentZoneMappingArgs(
                    zone_id=master_zone,
                    vswitch_id=example_master.id,
                ),
                alicloud.cen.TransitRouterVpcAttachmentZoneMappingArgs(
                    zone_id=slave_zone,
                    vswitch_id=example_slave.id,
                ),
            ],
            transit_router_attachment_name=name,
            transit_router_attachment_description=name)
        example_route_table = alicloud.vpc.RouteTable("example",
            vpc_id=example.id,
            route_table_name=name,
            description=name)
        example_child_instance_route_entry_to_attachment = alicloud.cen.ChildInstanceRouteEntryToAttachment("example",
            transit_router_attachment_id=example_transit_router_vpc_attachment.transit_router_attachment_id,
            cen_id=example_instance.id,
            destination_cidr_block="10.0.0.0/24",
            child_instance_route_table_id=example_route_table.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Cen Child Instance Route Entry To Attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment example <cen_id>:<child_instance_route_table_id>:<transit_router_attachment_id>:<destination_cidr_block>
        ```

        :param str resource_name: The name of the resource.
        :param ChildInstanceRouteEntryToAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ChildInstanceRouteEntryToAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 child_instance_route_table_id: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ChildInstanceRouteEntryToAttachmentArgs.__new__(ChildInstanceRouteEntryToAttachmentArgs)

            if cen_id is None and not opts.urn:
                raise TypeError("Missing required property 'cen_id'")
            __props__.__dict__["cen_id"] = cen_id
            if child_instance_route_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'child_instance_route_table_id'")
            __props__.__dict__["child_instance_route_table_id"] = child_instance_route_table_id
            if destination_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__.__dict__["destination_cidr_block"] = destination_cidr_block
            __props__.__dict__["dry_run"] = dry_run
            if transit_router_attachment_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_attachment_id'")
            __props__.__dict__["transit_router_attachment_id"] = transit_router_attachment_id
            __props__.__dict__["service_type"] = None
            __props__.__dict__["status"] = None
        super(ChildInstanceRouteEntryToAttachment, __self__).__init__(
            'alicloud:cen/childInstanceRouteEntryToAttachment:ChildInstanceRouteEntryToAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cen_id: Optional[pulumi.Input[str]] = None,
            child_instance_route_table_id: Optional[pulumi.Input[str]] = None,
            destination_cidr_block: Optional[pulumi.Input[str]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            service_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            transit_router_attachment_id: Optional[pulumi.Input[str]] = None) -> 'ChildInstanceRouteEntryToAttachment':
        """
        Get an existing ChildInstanceRouteEntryToAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] child_instance_route_table_id: The first ID of the resource
        :param pulumi.Input[str] destination_cidr_block: DestinationCidrBlock
        :param pulumi.Input[bool] dry_run: Whether to perform pre-check on this request, including permission and instance status verification.
        :param pulumi.Input[str] service_type: ServiceType
        :param pulumi.Input[str] status: The status of the resource
        :param pulumi.Input[str] transit_router_attachment_id: TransitRouterAttachmentId
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ChildInstanceRouteEntryToAttachmentState.__new__(_ChildInstanceRouteEntryToAttachmentState)

        __props__.__dict__["cen_id"] = cen_id
        __props__.__dict__["child_instance_route_table_id"] = child_instance_route_table_id
        __props__.__dict__["destination_cidr_block"] = destination_cidr_block
        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["service_type"] = service_type
        __props__.__dict__["status"] = status
        __props__.__dict__["transit_router_attachment_id"] = transit_router_attachment_id
        return ChildInstanceRouteEntryToAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Output[str]:
        """
        The ID of the CEN instance.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter(name="childInstanceRouteTableId")
    def child_instance_route_table_id(self) -> pulumi.Output[str]:
        """
        The first ID of the resource
        """
        return pulumi.get(self, "child_instance_route_table_id")

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[str]:
        """
        DestinationCidrBlock
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to perform pre-check on this request, including permission and instance status verification.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> pulumi.Output[str]:
        """
        ServiceType
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> pulumi.Output[str]:
        """
        TransitRouterAttachmentId
        """
        return pulumi.get(self, "transit_router_attachment_id")

