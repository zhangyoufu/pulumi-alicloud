# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['RouteMap']


class RouteMap(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 as_path_match_mode: Optional[pulumi.Input[str]] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_region_id: Optional[pulumi.Input[str]] = None,
                 cidr_match_mode: Optional[pulumi.Input[str]] = None,
                 community_match_mode: Optional[pulumi.Input[str]] = None,
                 community_operate_mode: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_child_instance_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 destination_cidr_blocks: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 destination_instance_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 destination_instance_ids_reverse_match: Optional[pulumi.Input[bool]] = None,
                 destination_route_table_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 map_result: Optional[pulumi.Input[str]] = None,
                 match_asns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 match_community_sets: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 next_priority: Optional[pulumi.Input[float]] = None,
                 operate_community_sets: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 preference: Optional[pulumi.Input[float]] = None,
                 prepend_as_paths: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 route_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 source_child_instance_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 source_instance_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 source_instance_ids_reverse_match: Optional[pulumi.Input[bool]] = None,
                 source_region_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 source_route_table_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 transmit_direction: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This topic provides an overview of the route map function of Cloud Enterprise Networks (CENs).
        You can use the route map function to filter routes and modify route attributes.
        By doing so, you can manage the communication between networks attached to a CEN.

        For information about CEN Route Map and how to use it, see [Manage CEN Route Map](https://www.alibabacloud.com/help/doc-detail/124157.htm).

        > **NOTE:** Available in 1.82.0+

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_pulumi as pulumi

        default_instance = alicloud.cen.Instance("defaultInstance")
        vpc00_region = pulumi.providers.Alicloud("vpc00Region", region="cn-hangzhou")
        vpc01_region = pulumi.providers.Alicloud("vpc01Region", region="cn-shanghai")
        vpc00 = alicloud.vpc.Network("vpc00", cidr_block="172.16.0.0/12",
        opts=ResourceOptions(provider="alicloud.vpc00_region"))
        vpc01 = alicloud.vpc.Network("vpc01", cidr_block="172.16.0.0/12",
        opts=ResourceOptions(provider="alicloud.vpc01_region"))
        default00 = alicloud.cen.InstanceAttachment("default00",
            child_instance_id=vpc00.id,
            child_instance_region_id="cn-hangzhou",
            instance_id=default_instance.id)
        default01 = alicloud.cen.InstanceAttachment("default01",
            child_instance_id=vpc01.id,
            child_instance_region_id="cn-shanghai",
            instance_id=default_instance.id)
        default_route_map = alicloud.cen.RouteMap("defaultRouteMap",
            as_path_match_mode="Include",
            cen_id=alicloud_cen_instance["cen"]["id"],
            cen_region_id="cn-hangzhou",
            cidr_match_mode="Include",
            community_match_mode="Include",
            community_operate_mode="Additive",
            description="test-desc",
            destination_child_instance_types=["VPC"],
            destination_cidr_blocks=[vpc01.cidr_block],
            destination_instance_ids=[vpc01.id],
            destination_instance_ids_reverse_match=False,
            destination_route_table_ids=[vpc01.route_table_id],
            map_result="Permit",
            match_asns=["65501"],
            match_community_sets=["65501:1"],
            next_priority=1,
            operate_community_sets=["65501:1"],
            preference=20,
            prepend_as_paths=["65501"],
            priority=1,
            route_types=["System"],
            source_child_instance_types=["VPC"],
            source_instance_ids=[vpc00.id],
            source_instance_ids_reverse_match=False,
            source_region_ids=["cn-hangzhou"],
            source_route_table_ids=[vpc00.route_table_id],
            transmit_direction="RegionIn",
            opts=ResourceOptions(depends_on=[
                    "alicloud_cen_instance_attachment.default00",
                    "alicloud_cen_instance_attachment.default01",
                ]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] as_path_match_mode: A match statement. It indicates the mode in which the AS path attribute is matched. Valid values: ["Include", "Complete"].
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] cen_region_id: The ID of the region to which the CEN instance belongs.
        :param pulumi.Input[str] cidr_match_mode: A match statement. It indicates the mode in which the prefix attribute is matched. Valid values: ["Include", "Complete"].
        :param pulumi.Input[str] community_match_mode: A match statement. It indicates the mode in which the community attribute is matched. Valid values: ["Include", "Complete"].
        :param pulumi.Input[str] community_operate_mode: An action statement. It indicates the mode in which the community attribute is operated. Valid values: ["Additive", "Replace"].
        :param pulumi.Input[str] description: The description of the route map.
        :param pulumi.Input[List[pulumi.Input[str]]] destination_child_instance_types: A match statement that indicates the list of destination instance types. Valid values: ["VPC", "VBR", "CCN"].
        :param pulumi.Input[List[pulumi.Input[str]]] destination_cidr_blocks: A match statement that indicates the prefix list. The prefix is in the CIDR format. You can enter a maximum of 32 CIDR blocks.
        :param pulumi.Input[List[pulumi.Input[str]]] destination_instance_ids: A match statement that indicates the list of IDs of the destination instances.
        :param pulumi.Input[bool] destination_instance_ids_reverse_match: Indicates whether to enable the reverse match method for the DestinationInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        :param pulumi.Input[List[pulumi.Input[str]]] destination_route_table_ids: A match statement that indicates the list of IDs of the destination route tables. You can enter a maximum of 32 route table IDs.
        :param pulumi.Input[str] map_result: The action that is performed to a route if the route matches all the match conditions. Valid values: ["Permit", "Deny"].
        :param pulumi.Input[List[pulumi.Input[str]]] match_asns: A match statement that indicates the AS path list. The AS path is a well-known mandatory attribute, which describes the numbers of the ASs that a BGP route passes through during transmission.
        :param pulumi.Input[List[pulumi.Input[str]]] match_community_sets: A match statement that indicates the community set. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        :param pulumi.Input[float] next_priority: The priority of the next route map that is associated with the current route map. Value range: 1 to 100.
        :param pulumi.Input[List[pulumi.Input[str]]] operate_community_sets: An action statement that operates the community attribute. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        :param pulumi.Input[float] preference: An action statement that modifies the priority of the route. Value range: 1 to 100. The default priority of a route is 50. A lower value indicates a higher preference.
        :param pulumi.Input[List[pulumi.Input[str]]] prepend_as_paths: An action statement that indicates an AS path is prepended when the regional gateway receives or advertises a route.
        :param pulumi.Input[float] priority: The priority of the route map. Value range: 1 to 100. A lower value indicates a higher priority.
        :param pulumi.Input[List[pulumi.Input[str]]] route_types: A match statement that indicates the list of route types. Valid values: ["System", "Custom", "BGP"].
        :param pulumi.Input[List[pulumi.Input[str]]] source_child_instance_types: A match statement that indicates the list of source instance types. Valid values: ["VPC", "VBR", "CCN"].
        :param pulumi.Input[List[pulumi.Input[str]]] source_instance_ids: A match statement that indicates the list of IDs of the source instances.
        :param pulumi.Input[bool] source_instance_ids_reverse_match: Indicates whether to enable the reverse match method for the SourceInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        :param pulumi.Input[List[pulumi.Input[str]]] source_region_ids: A match statement that indicates the list of IDs of the source regions. You can enter a maximum of 32 region IDs.
        :param pulumi.Input[List[pulumi.Input[str]]] source_route_table_ids: A match statement that indicates the list of IDs of the source route tables. You can enter a maximum of 32 route table IDs.
        :param pulumi.Input[str] transmit_direction: The direction in which the route map is applied. Valid values: ["RegionIn", "RegionOut"].
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['as_path_match_mode'] = as_path_match_mode
            if cen_id is None:
                raise TypeError("Missing required property 'cen_id'")
            __props__['cen_id'] = cen_id
            if cen_region_id is None:
                raise TypeError("Missing required property 'cen_region_id'")
            __props__['cen_region_id'] = cen_region_id
            __props__['cidr_match_mode'] = cidr_match_mode
            __props__['community_match_mode'] = community_match_mode
            __props__['community_operate_mode'] = community_operate_mode
            __props__['description'] = description
            __props__['destination_child_instance_types'] = destination_child_instance_types
            __props__['destination_cidr_blocks'] = destination_cidr_blocks
            __props__['destination_instance_ids'] = destination_instance_ids
            __props__['destination_instance_ids_reverse_match'] = destination_instance_ids_reverse_match
            __props__['destination_route_table_ids'] = destination_route_table_ids
            if map_result is None:
                raise TypeError("Missing required property 'map_result'")
            __props__['map_result'] = map_result
            __props__['match_asns'] = match_asns
            __props__['match_community_sets'] = match_community_sets
            __props__['next_priority'] = next_priority
            __props__['operate_community_sets'] = operate_community_sets
            __props__['preference'] = preference
            __props__['prepend_as_paths'] = prepend_as_paths
            if priority is None:
                raise TypeError("Missing required property 'priority'")
            __props__['priority'] = priority
            __props__['route_types'] = route_types
            __props__['source_child_instance_types'] = source_child_instance_types
            __props__['source_instance_ids'] = source_instance_ids
            __props__['source_instance_ids_reverse_match'] = source_instance_ids_reverse_match
            __props__['source_region_ids'] = source_region_ids
            __props__['source_route_table_ids'] = source_route_table_ids
            if transmit_direction is None:
                raise TypeError("Missing required property 'transmit_direction'")
            __props__['transmit_direction'] = transmit_direction
            __props__['route_map_id'] = None
            __props__['status'] = None
        super(RouteMap, __self__).__init__(
            'alicloud:cen/routeMap:RouteMap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            as_path_match_mode: Optional[pulumi.Input[str]] = None,
            cen_id: Optional[pulumi.Input[str]] = None,
            cen_region_id: Optional[pulumi.Input[str]] = None,
            cidr_match_mode: Optional[pulumi.Input[str]] = None,
            community_match_mode: Optional[pulumi.Input[str]] = None,
            community_operate_mode: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            destination_child_instance_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            destination_cidr_blocks: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            destination_instance_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            destination_instance_ids_reverse_match: Optional[pulumi.Input[bool]] = None,
            destination_route_table_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            map_result: Optional[pulumi.Input[str]] = None,
            match_asns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            match_community_sets: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            next_priority: Optional[pulumi.Input[float]] = None,
            operate_community_sets: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            preference: Optional[pulumi.Input[float]] = None,
            prepend_as_paths: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            priority: Optional[pulumi.Input[float]] = None,
            route_map_id: Optional[pulumi.Input[str]] = None,
            route_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            source_child_instance_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            source_instance_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            source_instance_ids_reverse_match: Optional[pulumi.Input[bool]] = None,
            source_region_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            source_route_table_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            transmit_direction: Optional[pulumi.Input[str]] = None) -> 'RouteMap':
        """
        Get an existing RouteMap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] as_path_match_mode: A match statement. It indicates the mode in which the AS path attribute is matched. Valid values: ["Include", "Complete"].
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] cen_region_id: The ID of the region to which the CEN instance belongs.
        :param pulumi.Input[str] cidr_match_mode: A match statement. It indicates the mode in which the prefix attribute is matched. Valid values: ["Include", "Complete"].
        :param pulumi.Input[str] community_match_mode: A match statement. It indicates the mode in which the community attribute is matched. Valid values: ["Include", "Complete"].
        :param pulumi.Input[str] community_operate_mode: An action statement. It indicates the mode in which the community attribute is operated. Valid values: ["Additive", "Replace"].
        :param pulumi.Input[str] description: The description of the route map.
        :param pulumi.Input[List[pulumi.Input[str]]] destination_child_instance_types: A match statement that indicates the list of destination instance types. Valid values: ["VPC", "VBR", "CCN"].
        :param pulumi.Input[List[pulumi.Input[str]]] destination_cidr_blocks: A match statement that indicates the prefix list. The prefix is in the CIDR format. You can enter a maximum of 32 CIDR blocks.
        :param pulumi.Input[List[pulumi.Input[str]]] destination_instance_ids: A match statement that indicates the list of IDs of the destination instances.
        :param pulumi.Input[bool] destination_instance_ids_reverse_match: Indicates whether to enable the reverse match method for the DestinationInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        :param pulumi.Input[List[pulumi.Input[str]]] destination_route_table_ids: A match statement that indicates the list of IDs of the destination route tables. You can enter a maximum of 32 route table IDs.
        :param pulumi.Input[str] map_result: The action that is performed to a route if the route matches all the match conditions. Valid values: ["Permit", "Deny"].
        :param pulumi.Input[List[pulumi.Input[str]]] match_asns: A match statement that indicates the AS path list. The AS path is a well-known mandatory attribute, which describes the numbers of the ASs that a BGP route passes through during transmission.
        :param pulumi.Input[List[pulumi.Input[str]]] match_community_sets: A match statement that indicates the community set. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        :param pulumi.Input[float] next_priority: The priority of the next route map that is associated with the current route map. Value range: 1 to 100.
        :param pulumi.Input[List[pulumi.Input[str]]] operate_community_sets: An action statement that operates the community attribute. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        :param pulumi.Input[float] preference: An action statement that modifies the priority of the route. Value range: 1 to 100. The default priority of a route is 50. A lower value indicates a higher preference.
        :param pulumi.Input[List[pulumi.Input[str]]] prepend_as_paths: An action statement that indicates an AS path is prepended when the regional gateway receives or advertises a route.
        :param pulumi.Input[float] priority: The priority of the route map. Value range: 1 to 100. A lower value indicates a higher priority.
        :param pulumi.Input[List[pulumi.Input[str]]] route_types: A match statement that indicates the list of route types. Valid values: ["System", "Custom", "BGP"].
        :param pulumi.Input[List[pulumi.Input[str]]] source_child_instance_types: A match statement that indicates the list of source instance types. Valid values: ["VPC", "VBR", "CCN"].
        :param pulumi.Input[List[pulumi.Input[str]]] source_instance_ids: A match statement that indicates the list of IDs of the source instances.
        :param pulumi.Input[bool] source_instance_ids_reverse_match: Indicates whether to enable the reverse match method for the SourceInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        :param pulumi.Input[List[pulumi.Input[str]]] source_region_ids: A match statement that indicates the list of IDs of the source regions. You can enter a maximum of 32 region IDs.
        :param pulumi.Input[List[pulumi.Input[str]]] source_route_table_ids: A match statement that indicates the list of IDs of the source route tables. You can enter a maximum of 32 route table IDs.
        :param pulumi.Input[str] status: (Computed) The status of route map. Valid values: ["Creating", "Active", "Deleting"].
        :param pulumi.Input[str] transmit_direction: The direction in which the route map is applied. Valid values: ["RegionIn", "RegionOut"].
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["as_path_match_mode"] = as_path_match_mode
        __props__["cen_id"] = cen_id
        __props__["cen_region_id"] = cen_region_id
        __props__["cidr_match_mode"] = cidr_match_mode
        __props__["community_match_mode"] = community_match_mode
        __props__["community_operate_mode"] = community_operate_mode
        __props__["description"] = description
        __props__["destination_child_instance_types"] = destination_child_instance_types
        __props__["destination_cidr_blocks"] = destination_cidr_blocks
        __props__["destination_instance_ids"] = destination_instance_ids
        __props__["destination_instance_ids_reverse_match"] = destination_instance_ids_reverse_match
        __props__["destination_route_table_ids"] = destination_route_table_ids
        __props__["map_result"] = map_result
        __props__["match_asns"] = match_asns
        __props__["match_community_sets"] = match_community_sets
        __props__["next_priority"] = next_priority
        __props__["operate_community_sets"] = operate_community_sets
        __props__["preference"] = preference
        __props__["prepend_as_paths"] = prepend_as_paths
        __props__["priority"] = priority
        __props__["route_map_id"] = route_map_id
        __props__["route_types"] = route_types
        __props__["source_child_instance_types"] = source_child_instance_types
        __props__["source_instance_ids"] = source_instance_ids
        __props__["source_instance_ids_reverse_match"] = source_instance_ids_reverse_match
        __props__["source_region_ids"] = source_region_ids
        __props__["source_route_table_ids"] = source_route_table_ids
        __props__["status"] = status
        __props__["transmit_direction"] = transmit_direction
        return RouteMap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="asPathMatchMode")
    def as_path_match_mode(self) -> Optional[str]:
        """
        A match statement. It indicates the mode in which the AS path attribute is matched. Valid values: ["Include", "Complete"].
        """
        return pulumi.get(self, "as_path_match_mode")

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> str:
        """
        The ID of the CEN instance.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter(name="cenRegionId")
    def cen_region_id(self) -> str:
        """
        The ID of the region to which the CEN instance belongs.
        """
        return pulumi.get(self, "cen_region_id")

    @property
    @pulumi.getter(name="cidrMatchMode")
    def cidr_match_mode(self) -> Optional[str]:
        """
        A match statement. It indicates the mode in which the prefix attribute is matched. Valid values: ["Include", "Complete"].
        """
        return pulumi.get(self, "cidr_match_mode")

    @property
    @pulumi.getter(name="communityMatchMode")
    def community_match_mode(self) -> Optional[str]:
        """
        A match statement. It indicates the mode in which the community attribute is matched. Valid values: ["Include", "Complete"].
        """
        return pulumi.get(self, "community_match_mode")

    @property
    @pulumi.getter(name="communityOperateMode")
    def community_operate_mode(self) -> Optional[str]:
        """
        An action statement. It indicates the mode in which the community attribute is operated. Valid values: ["Additive", "Replace"].
        """
        return pulumi.get(self, "community_operate_mode")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the route map.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationChildInstanceTypes")
    def destination_child_instance_types(self) -> Optional[List[str]]:
        """
        A match statement that indicates the list of destination instance types. Valid values: ["VPC", "VBR", "CCN"].
        """
        return pulumi.get(self, "destination_child_instance_types")

    @property
    @pulumi.getter(name="destinationCidrBlocks")
    def destination_cidr_blocks(self) -> Optional[List[str]]:
        """
        A match statement that indicates the prefix list. The prefix is in the CIDR format. You can enter a maximum of 32 CIDR blocks.
        """
        return pulumi.get(self, "destination_cidr_blocks")

    @property
    @pulumi.getter(name="destinationInstanceIds")
    def destination_instance_ids(self) -> Optional[List[str]]:
        """
        A match statement that indicates the list of IDs of the destination instances.
        """
        return pulumi.get(self, "destination_instance_ids")

    @property
    @pulumi.getter(name="destinationInstanceIdsReverseMatch")
    def destination_instance_ids_reverse_match(self) -> Optional[bool]:
        """
        Indicates whether to enable the reverse match method for the DestinationInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        """
        return pulumi.get(self, "destination_instance_ids_reverse_match")

    @property
    @pulumi.getter(name="destinationRouteTableIds")
    def destination_route_table_ids(self) -> Optional[List[str]]:
        """
        A match statement that indicates the list of IDs of the destination route tables. You can enter a maximum of 32 route table IDs.
        """
        return pulumi.get(self, "destination_route_table_ids")

    @property
    @pulumi.getter(name="mapResult")
    def map_result(self) -> str:
        """
        The action that is performed to a route if the route matches all the match conditions. Valid values: ["Permit", "Deny"].
        """
        return pulumi.get(self, "map_result")

    @property
    @pulumi.getter(name="matchAsns")
    def match_asns(self) -> Optional[List[str]]:
        """
        A match statement that indicates the AS path list. The AS path is a well-known mandatory attribute, which describes the numbers of the ASs that a BGP route passes through during transmission.
        """
        return pulumi.get(self, "match_asns")

    @property
    @pulumi.getter(name="matchCommunitySets")
    def match_community_sets(self) -> Optional[List[str]]:
        """
        A match statement that indicates the community set. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        """
        return pulumi.get(self, "match_community_sets")

    @property
    @pulumi.getter(name="nextPriority")
    def next_priority(self) -> Optional[float]:
        """
        The priority of the next route map that is associated with the current route map. Value range: 1 to 100.
        """
        return pulumi.get(self, "next_priority")

    @property
    @pulumi.getter(name="operateCommunitySets")
    def operate_community_sets(self) -> Optional[List[str]]:
        """
        An action statement that operates the community attribute. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.
        """
        return pulumi.get(self, "operate_community_sets")

    @property
    @pulumi.getter
    def preference(self) -> Optional[float]:
        """
        An action statement that modifies the priority of the route. Value range: 1 to 100. The default priority of a route is 50. A lower value indicates a higher preference.
        """
        return pulumi.get(self, "preference")

    @property
    @pulumi.getter(name="prependAsPaths")
    def prepend_as_paths(self) -> Optional[List[str]]:
        """
        An action statement that indicates an AS path is prepended when the regional gateway receives or advertises a route.
        """
        return pulumi.get(self, "prepend_as_paths")

    @property
    @pulumi.getter
    def priority(self) -> float:
        """
        The priority of the route map. Value range: 1 to 100. A lower value indicates a higher priority.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="routeMapId")
    def route_map_id(self) -> str:
        return pulumi.get(self, "route_map_id")

    @property
    @pulumi.getter(name="routeTypes")
    def route_types(self) -> Optional[List[str]]:
        """
        A match statement that indicates the list of route types. Valid values: ["System", "Custom", "BGP"].
        """
        return pulumi.get(self, "route_types")

    @property
    @pulumi.getter(name="sourceChildInstanceTypes")
    def source_child_instance_types(self) -> Optional[List[str]]:
        """
        A match statement that indicates the list of source instance types. Valid values: ["VPC", "VBR", "CCN"].
        """
        return pulumi.get(self, "source_child_instance_types")

    @property
    @pulumi.getter(name="sourceInstanceIds")
    def source_instance_ids(self) -> Optional[List[str]]:
        """
        A match statement that indicates the list of IDs of the source instances.
        """
        return pulumi.get(self, "source_instance_ids")

    @property
    @pulumi.getter(name="sourceInstanceIdsReverseMatch")
    def source_instance_ids_reverse_match(self) -> Optional[bool]:
        """
        Indicates whether to enable the reverse match method for the SourceInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".
        """
        return pulumi.get(self, "source_instance_ids_reverse_match")

    @property
    @pulumi.getter(name="sourceRegionIds")
    def source_region_ids(self) -> Optional[List[str]]:
        """
        A match statement that indicates the list of IDs of the source regions. You can enter a maximum of 32 region IDs.
        """
        return pulumi.get(self, "source_region_ids")

    @property
    @pulumi.getter(name="sourceRouteTableIds")
    def source_route_table_ids(self) -> Optional[List[str]]:
        """
        A match statement that indicates the list of IDs of the source route tables. You can enter a maximum of 32 route table IDs.
        """
        return pulumi.get(self, "source_route_table_ids")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        (Computed) The status of route map. Valid values: ["Creating", "Active", "Deleting"].
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transmitDirection")
    def transmit_direction(self) -> str:
        """
        The direction in which the route map is applied. Valid values: ["RegionIn", "RegionOut"].
        """
        return pulumi.get(self, "transmit_direction")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

