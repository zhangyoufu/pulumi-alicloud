# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TransitRouterRouteTableAssociationArgs', 'TransitRouterRouteTableAssociation']

@pulumi.input_type
class TransitRouterRouteTableAssociationArgs:
    def __init__(__self__, *,
                 transit_router_attachment_id: pulumi.Input[str],
                 transit_router_route_table_id: pulumi.Input[str],
                 dry_run: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a TransitRouterRouteTableAssociation resource.
        :param pulumi.Input[str] transit_router_attachment_id: The ID the transit router attachment.
        :param pulumi.Input[str] transit_router_route_table_id: The ID of the transit router route table.
        :param pulumi.Input[bool] dry_run: The dry run.
               
               > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zone_id of zone_mapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
        """
        pulumi.set(__self__, "transit_router_attachment_id", transit_router_attachment_id)
        pulumi.set(__self__, "transit_router_route_table_id", transit_router_route_table_id)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> pulumi.Input[str]:
        """
        The ID the transit router attachment.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @transit_router_attachment_id.setter
    def transit_router_attachment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_router_attachment_id", value)

    @property
    @pulumi.getter(name="transitRouterRouteTableId")
    def transit_router_route_table_id(self) -> pulumi.Input[str]:
        """
        The ID of the transit router route table.
        """
        return pulumi.get(self, "transit_router_route_table_id")

    @transit_router_route_table_id.setter
    def transit_router_route_table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_router_route_table_id", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        The dry run.

        > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zone_id of zone_mapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)


@pulumi.input_type
class _TransitRouterRouteTableAssociationState:
    def __init__(__self__, *,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                 transit_router_route_table_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TransitRouterRouteTableAssociation resources.
        :param pulumi.Input[bool] dry_run: The dry run.
               
               > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zone_id of zone_mapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
        :param pulumi.Input[str] status: The associating status of the network.
        :param pulumi.Input[str] transit_router_attachment_id: The ID the transit router attachment.
        :param pulumi.Input[str] transit_router_route_table_id: The ID of the transit router route table.
        """
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if transit_router_attachment_id is not None:
            pulumi.set(__self__, "transit_router_attachment_id", transit_router_attachment_id)
        if transit_router_route_table_id is not None:
            pulumi.set(__self__, "transit_router_route_table_id", transit_router_route_table_id)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        The dry run.

        > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zone_id of zone_mapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The associating status of the network.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID the transit router attachment.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @transit_router_attachment_id.setter
    def transit_router_attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_attachment_id", value)

    @property
    @pulumi.getter(name="transitRouterRouteTableId")
    def transit_router_route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the transit router route table.
        """
        return pulumi.get(self, "transit_router_route_table_id")

    @transit_router_route_table_id.setter
    def transit_router_route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_route_table_id", value)


class TransitRouterRouteTableAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                 transit_router_route_table_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a CEN transit router route table association resource.[What is Cen Transit Router Route Table Association](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutetableaggregation)

        > **NOTE:** Available since v1.126.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        default = alicloud.cen.get_transit_router_available_resources()
        master_zone = default.resources[0].master_zones[0]
        slave_zone = default.resources[0].slave_zones[1]
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name=name,
            cidr_block="192.168.0.0/16")
        example_master = alicloud.vpc.Switch("exampleMaster",
            vswitch_name=name,
            cidr_block="192.168.1.0/24",
            vpc_id=example_network.id,
            zone_id=master_zone)
        example_slave = alicloud.vpc.Switch("exampleSlave",
            vswitch_name=name,
            cidr_block="192.168.2.0/24",
            vpc_id=example_network.id,
            zone_id=slave_zone)
        example_instance = alicloud.cen.Instance("exampleInstance",
            cen_instance_name=name,
            protection_level="REDUCED")
        example_transit_router = alicloud.cen.TransitRouter("exampleTransitRouter",
            transit_router_name=name,
            cen_id=example_instance.id)
        example_transit_router_vpc_attachment = alicloud.cen.TransitRouterVpcAttachment("exampleTransitRouterVpcAttachment",
            cen_id=example_instance.id,
            transit_router_id=example_transit_router.transit_router_id,
            vpc_id=example_network.id,
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
        example_transit_router_route_table = alicloud.cen.TransitRouterRouteTable("exampleTransitRouterRouteTable", transit_router_id=example_transit_router.transit_router_id)
        example_transit_router_route_table_association = alicloud.cen.TransitRouterRouteTableAssociation("exampleTransitRouterRouteTableAssociation",
            transit_router_route_table_id=example_transit_router_route_table.transit_router_route_table_id,
            transit_router_attachment_id=example_transit_router_vpc_attachment.transit_router_attachment_id)
        ```

        ## Import

        CEN transit router route table association can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/transitRouterRouteTableAssociation:TransitRouterRouteTableAssociation default tr-********:tr-attach-********
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dry_run: The dry run.
               
               > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zone_id of zone_mapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
        :param pulumi.Input[str] transit_router_attachment_id: The ID the transit router attachment.
        :param pulumi.Input[str] transit_router_route_table_id: The ID of the transit router route table.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransitRouterRouteTableAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CEN transit router route table association resource.[What is Cen Transit Router Route Table Association](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutetableaggregation)

        > **NOTE:** Available since v1.126.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        default = alicloud.cen.get_transit_router_available_resources()
        master_zone = default.resources[0].master_zones[0]
        slave_zone = default.resources[0].slave_zones[1]
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name=name,
            cidr_block="192.168.0.0/16")
        example_master = alicloud.vpc.Switch("exampleMaster",
            vswitch_name=name,
            cidr_block="192.168.1.0/24",
            vpc_id=example_network.id,
            zone_id=master_zone)
        example_slave = alicloud.vpc.Switch("exampleSlave",
            vswitch_name=name,
            cidr_block="192.168.2.0/24",
            vpc_id=example_network.id,
            zone_id=slave_zone)
        example_instance = alicloud.cen.Instance("exampleInstance",
            cen_instance_name=name,
            protection_level="REDUCED")
        example_transit_router = alicloud.cen.TransitRouter("exampleTransitRouter",
            transit_router_name=name,
            cen_id=example_instance.id)
        example_transit_router_vpc_attachment = alicloud.cen.TransitRouterVpcAttachment("exampleTransitRouterVpcAttachment",
            cen_id=example_instance.id,
            transit_router_id=example_transit_router.transit_router_id,
            vpc_id=example_network.id,
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
        example_transit_router_route_table = alicloud.cen.TransitRouterRouteTable("exampleTransitRouterRouteTable", transit_router_id=example_transit_router.transit_router_id)
        example_transit_router_route_table_association = alicloud.cen.TransitRouterRouteTableAssociation("exampleTransitRouterRouteTableAssociation",
            transit_router_route_table_id=example_transit_router_route_table.transit_router_route_table_id,
            transit_router_attachment_id=example_transit_router_vpc_attachment.transit_router_attachment_id)
        ```

        ## Import

        CEN transit router route table association can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/transitRouterRouteTableAssociation:TransitRouterRouteTableAssociation default tr-********:tr-attach-********
        ```

        :param str resource_name: The name of the resource.
        :param TransitRouterRouteTableAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransitRouterRouteTableAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                 transit_router_route_table_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransitRouterRouteTableAssociationArgs.__new__(TransitRouterRouteTableAssociationArgs)

            __props__.__dict__["dry_run"] = dry_run
            if transit_router_attachment_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_attachment_id'")
            __props__.__dict__["transit_router_attachment_id"] = transit_router_attachment_id
            if transit_router_route_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_route_table_id'")
            __props__.__dict__["transit_router_route_table_id"] = transit_router_route_table_id
            __props__.__dict__["status"] = None
        super(TransitRouterRouteTableAssociation, __self__).__init__(
            'alicloud:cen/transitRouterRouteTableAssociation:TransitRouterRouteTableAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None,
            transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
            transit_router_route_table_id: Optional[pulumi.Input[str]] = None) -> 'TransitRouterRouteTableAssociation':
        """
        Get an existing TransitRouterRouteTableAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dry_run: The dry run.
               
               > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zone_id of zone_mapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
        :param pulumi.Input[str] status: The associating status of the network.
        :param pulumi.Input[str] transit_router_attachment_id: The ID the transit router attachment.
        :param pulumi.Input[str] transit_router_route_table_id: The ID of the transit router route table.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransitRouterRouteTableAssociationState.__new__(_TransitRouterRouteTableAssociationState)

        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["status"] = status
        __props__.__dict__["transit_router_attachment_id"] = transit_router_attachment_id
        __props__.__dict__["transit_router_route_table_id"] = transit_router_route_table_id
        return TransitRouterRouteTableAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        The dry run.

        > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zone_id of zone_mapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The associating status of the network.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> pulumi.Output[str]:
        """
        The ID the transit router attachment.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @property
    @pulumi.getter(name="transitRouterRouteTableId")
    def transit_router_route_table_id(self) -> pulumi.Output[str]:
        """
        The ID of the transit router route table.
        """
        return pulumi.get(self, "transit_router_route_table_id")

