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
    'GetServerGroupsResult',
    'AwaitableGetServerGroupsResult',
    'get_server_groups',
    'get_server_groups_output',
]

@pulumi.output_type
class GetServerGroupsResult:
    """
    A collection of values returned by getServerGroups.
    """
    def __init__(__self__, id=None, ids=None, load_balancer_id=None, name_regex=None, names=None, output_file=None, slb_server_groups=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if load_balancer_id and not isinstance(load_balancer_id, str):
            raise TypeError("Expected argument 'load_balancer_id' to be a str")
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if slb_server_groups and not isinstance(slb_server_groups, list):
            raise TypeError("Expected argument 'slb_server_groups' to be a list")
        pulumi.set(__self__, "slb_server_groups", slb_server_groups)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        """
        A list of SLB VServer groups IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of SLB VServer groups names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="slbServerGroups")
    def slb_server_groups(self) -> Sequence['outputs.GetServerGroupsSlbServerGroupResult']:
        """
        A list of SLB VServer groups. Each element contains the following attributes:
        """
        return pulumi.get(self, "slb_server_groups")


class AwaitableGetServerGroupsResult(GetServerGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerGroupsResult(
            id=self.id,
            ids=self.ids,
            load_balancer_id=self.load_balancer_id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            slb_server_groups=self.slb_server_groups)


def get_server_groups(ids: Optional[Sequence[str]] = None,
                      load_balancer_id: Optional[str] = None,
                      name_regex: Optional[str] = None,
                      output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerGroupsResult:
    """
    This data source provides the VServer groups related to a server load balancer.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "slbservergroups"
    default_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
        available_resource_creation="VSwitch")
    default_network = alicloud.vpc.Network("defaultNetwork",
        vpc_name=name,
        cidr_block="172.16.0.0/16")
    default_switch = alicloud.vpc.Switch("defaultSwitch",
        vpc_id=default_network.id,
        cidr_block="172.16.0.0/16",
        zone_id=default_zones.zones[0].id,
        vswitch_name=name)
    default_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer",
        load_balancer_name=name,
        vswitch_id=default_switch.id)
    default_server_group = alicloud.slb.ServerGroup("defaultServerGroup", load_balancer_id=default_application_load_balancer.id)
    sample_ds = alicloud.slb.get_server_groups_output(load_balancer_id=default_application_load_balancer.id)
    pulumi.export("firstSlbServerGroupId", sample_ds.slb_server_groups[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of VServer group IDs to filter results.
    :param str load_balancer_id: ID of the SLB.
    :param str name_regex: A regex string to filter results by VServer group name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['loadBalancerId'] = load_balancer_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:slb/getServerGroups:getServerGroups', __args__, opts=opts, typ=GetServerGroupsResult).value

    return AwaitableGetServerGroupsResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        load_balancer_id=pulumi.get(__ret__, 'load_balancer_id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        slb_server_groups=pulumi.get(__ret__, 'slb_server_groups'))


@_utilities.lift_output_func(get_server_groups)
def get_server_groups_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             load_balancer_id: Optional[pulumi.Input[str]] = None,
                             name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                             output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerGroupsResult]:
    """
    This data source provides the VServer groups related to a server load balancer.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "slbservergroups"
    default_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
        available_resource_creation="VSwitch")
    default_network = alicloud.vpc.Network("defaultNetwork",
        vpc_name=name,
        cidr_block="172.16.0.0/16")
    default_switch = alicloud.vpc.Switch("defaultSwitch",
        vpc_id=default_network.id,
        cidr_block="172.16.0.0/16",
        zone_id=default_zones.zones[0].id,
        vswitch_name=name)
    default_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer",
        load_balancer_name=name,
        vswitch_id=default_switch.id)
    default_server_group = alicloud.slb.ServerGroup("defaultServerGroup", load_balancer_id=default_application_load_balancer.id)
    sample_ds = alicloud.slb.get_server_groups_output(load_balancer_id=default_application_load_balancer.id)
    pulumi.export("firstSlbServerGroupId", sample_ds.slb_server_groups[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of VServer group IDs to filter results.
    :param str load_balancer_id: ID of the SLB.
    :param str name_regex: A regex string to filter results by VServer group name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
