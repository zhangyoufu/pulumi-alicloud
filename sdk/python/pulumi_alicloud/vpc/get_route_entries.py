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

__all__ = [
    'GetRouteEntriesResult',
    'AwaitableGetRouteEntriesResult',
    'get_route_entries',
    'get_route_entries_output',
]

@pulumi.output_type
class GetRouteEntriesResult:
    """
    A collection of values returned by getRouteEntries.
    """
    def __init__(__self__, cidr_block=None, entries=None, id=None, instance_id=None, output_file=None, route_table_id=None, type=None):
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError("Expected argument 'cidr_block' to be a str")
        pulumi.set(__self__, "cidr_block", cidr_block)
        if entries and not isinstance(entries, list):
            raise TypeError("Expected argument 'entries' to be a list")
        pulumi.set(__self__, "entries", entries)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if route_table_id and not isinstance(route_table_id, str):
            raise TypeError("Expected argument 'route_table_id' to be a str")
        pulumi.set(__self__, "route_table_id", route_table_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[str]:
        """
        The destination CIDR block of the route entry.
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter
    def entries(self) -> Sequence['outputs.GetRouteEntriesEntryResult']:
        """
        A list of Route Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        The instance ID of the next hop.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> str:
        """
        The ID of the router table to which the route entry belongs.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the route entry.
        """
        return pulumi.get(self, "type")


class AwaitableGetRouteEntriesResult(GetRouteEntriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteEntriesResult(
            cidr_block=self.cidr_block,
            entries=self.entries,
            id=self.id,
            instance_id=self.instance_id,
            output_file=self.output_file,
            route_table_id=self.route_table_id,
            type=self.type)


def get_route_entries(cidr_block: Optional[str] = None,
                      instance_id: Optional[str] = None,
                      output_file: Optional[str] = None,
                      route_table_id: Optional[str] = None,
                      type: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteEntriesResult:
    """
    This data source provides a list of Route Entries owned by an Alibaba Cloud account.

    > **NOTE:** Available in 1.37.0+.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.get_zones(available_resource_creation="VSwitch")
    default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default.zones[0].id,
        cpu_core_count=1,
        memory_size=2)
    default_get_images = alicloud.ecs.get_images(name_regex="^ubuntu_18.*64",
        most_recent=True,
        owners="system")
    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "tf-testAccRouteEntryConfig"
    foo_network = alicloud.vpc.Network("foo",
        name=name,
        cidr_block="10.1.0.0/21")
    foo_switch = alicloud.vpc.Switch("foo",
        vpc_id=foo_network.id,
        cidr_block="10.1.1.0/24",
        availability_zone=default.zones[0].id,
        vswitch_name=name)
    tf_test_foo = alicloud.ecs.SecurityGroup("tf_test_foo",
        name=name,
        description="foo",
        vpc_id=foo_network.id)
    foo_instance = alicloud.ecs.Instance("foo",
        security_groups=[tf_test_foo.id],
        vswitch_id=foo_switch.id,
        allocate_public_ip=True,
        instance_charge_type="PostPaid",
        instance_type=default_get_instance_types.instance_types[0].id,
        internet_charge_type="PayByTraffic",
        internet_max_bandwidth_out=5,
        system_disk_category="cloud_efficiency",
        image_id=default_get_images.images[0].id,
        instance_name=name)
    foo_route_entry = alicloud.vpc.RouteEntry("foo",
        route_table_id=foo_network.route_table_id,
        destination_cidrblock="172.11.1.1/32",
        nexthop_type="Instance",
        nexthop_id=foo_instance.id)
    ingress = alicloud.ecs.SecurityGroupRule("ingress",
        type="ingress",
        ip_protocol="tcp",
        nic_type="intranet",
        policy="accept",
        port_range="22/22",
        priority=1,
        security_group_id=tf_test_foo.id,
        cidr_ip="0.0.0.0/0")
    foo = alicloud.vpc.get_route_entries_output(route_table_id=foo_route_entry.route_table_id)
    ```
    <!--End PulumiCodeChooser -->


    :param str cidr_block: The destination CIDR block of the route entry.
    :param str instance_id: The instance ID of the next hop.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str route_table_id: The ID of the router table to which the route entry belongs.
    :param str type: The type of the route entry.
    """
    __args__ = dict()
    __args__['cidrBlock'] = cidr_block
    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    __args__['routeTableId'] = route_table_id
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:vpc/getRouteEntries:getRouteEntries', __args__, opts=opts, typ=GetRouteEntriesResult).value

    return AwaitableGetRouteEntriesResult(
        cidr_block=pulumi.get(__ret__, 'cidr_block'),
        entries=pulumi.get(__ret__, 'entries'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        route_table_id=pulumi.get(__ret__, 'route_table_id'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_route_entries)
def get_route_entries_output(cidr_block: Optional[pulumi.Input[Optional[str]]] = None,
                             instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                             output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             route_table_id: Optional[pulumi.Input[str]] = None,
                             type: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouteEntriesResult]:
    """
    This data source provides a list of Route Entries owned by an Alibaba Cloud account.

    > **NOTE:** Available in 1.37.0+.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.get_zones(available_resource_creation="VSwitch")
    default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default.zones[0].id,
        cpu_core_count=1,
        memory_size=2)
    default_get_images = alicloud.ecs.get_images(name_regex="^ubuntu_18.*64",
        most_recent=True,
        owners="system")
    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "tf-testAccRouteEntryConfig"
    foo_network = alicloud.vpc.Network("foo",
        name=name,
        cidr_block="10.1.0.0/21")
    foo_switch = alicloud.vpc.Switch("foo",
        vpc_id=foo_network.id,
        cidr_block="10.1.1.0/24",
        availability_zone=default.zones[0].id,
        vswitch_name=name)
    tf_test_foo = alicloud.ecs.SecurityGroup("tf_test_foo",
        name=name,
        description="foo",
        vpc_id=foo_network.id)
    foo_instance = alicloud.ecs.Instance("foo",
        security_groups=[tf_test_foo.id],
        vswitch_id=foo_switch.id,
        allocate_public_ip=True,
        instance_charge_type="PostPaid",
        instance_type=default_get_instance_types.instance_types[0].id,
        internet_charge_type="PayByTraffic",
        internet_max_bandwidth_out=5,
        system_disk_category="cloud_efficiency",
        image_id=default_get_images.images[0].id,
        instance_name=name)
    foo_route_entry = alicloud.vpc.RouteEntry("foo",
        route_table_id=foo_network.route_table_id,
        destination_cidrblock="172.11.1.1/32",
        nexthop_type="Instance",
        nexthop_id=foo_instance.id)
    ingress = alicloud.ecs.SecurityGroupRule("ingress",
        type="ingress",
        ip_protocol="tcp",
        nic_type="intranet",
        policy="accept",
        port_range="22/22",
        priority=1,
        security_group_id=tf_test_foo.id,
        cidr_ip="0.0.0.0/0")
    foo = alicloud.vpc.get_route_entries_output(route_table_id=foo_route_entry.route_table_id)
    ```
    <!--End PulumiCodeChooser -->


    :param str cidr_block: The destination CIDR block of the route entry.
    :param str instance_id: The instance ID of the next hop.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str route_table_id: The ID of the router table to which the route entry belongs.
    :param str type: The type of the route entry.
    """
    ...
