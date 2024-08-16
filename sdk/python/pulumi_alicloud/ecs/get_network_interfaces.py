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
    'GetNetworkInterfacesResult',
    'AwaitableGetNetworkInterfacesResult',
    'get_network_interfaces',
    'get_network_interfaces_output',
]

@pulumi.output_type
class GetNetworkInterfacesResult:
    """
    A collection of values returned by getNetworkInterfaces.
    """
    def __init__(__self__, id=None, ids=None, instance_id=None, interfaces=None, name=None, name_regex=None, names=None, network_interface_name=None, output_file=None, primary_ip_address=None, private_ip=None, resource_group_id=None, security_group_id=None, service_managed=None, status=None, tags=None, type=None, vpc_id=None, vswitch_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if network_interface_name and not isinstance(network_interface_name, str):
            raise TypeError("Expected argument 'network_interface_name' to be a str")
        pulumi.set(__self__, "network_interface_name", network_interface_name)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if primary_ip_address and not isinstance(primary_ip_address, str):
            raise TypeError("Expected argument 'primary_ip_address' to be a str")
        pulumi.set(__self__, "primary_ip_address", primary_ip_address)
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        pulumi.set(__self__, "private_ip", private_ip)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if service_managed and not isinstance(service_managed, bool):
            raise TypeError("Expected argument 'service_managed' to be a bool")
        pulumi.set(__self__, "service_managed", service_managed)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_id and not isinstance(vswitch_id, str):
            raise TypeError("Expected argument 'vswitch_id' to be a str")
        pulumi.set(__self__, "vswitch_id", vswitch_id)

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
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        ID of the instance that the ENI is attached to.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def interfaces(self) -> Sequence['outputs.GetNetworkInterfacesInterfaceResult']:
        """
        A list of ENIs. Each element contains the following attributes:
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter
    @_utilities.deprecated("""Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead""")
    def name(self) -> Optional[str]:
        """
        Name of the ENI.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="networkInterfaceName")
    def network_interface_name(self) -> Optional[str]:
        return pulumi.get(self, "network_interface_name")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="primaryIpAddress")
    def primary_ip_address(self) -> Optional[str]:
        return pulumi.get(self, "primary_ip_address")

    @property
    @pulumi.getter(name="privateIp")
    @_utilities.deprecated("""Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead""")
    def private_ip(self) -> Optional[str]:
        """
        Primary private IP of the ENI.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        """
        The Id of resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[str]:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="serviceManaged")
    def service_managed(self) -> Optional[bool]:
        return pulumi.get(self, "service_managed")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Current status of the ENI.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags assigned to the ENI.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        ID of the VPC that the ENI belongs to.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[str]:
        """
        ID of the vSwitch that the ENI is linked to.
        """
        return pulumi.get(self, "vswitch_id")


class AwaitableGetNetworkInterfacesResult(GetNetworkInterfacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkInterfacesResult(
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            interfaces=self.interfaces,
            name=self.name,
            name_regex=self.name_regex,
            names=self.names,
            network_interface_name=self.network_interface_name,
            output_file=self.output_file,
            primary_ip_address=self.primary_ip_address,
            private_ip=self.private_ip,
            resource_group_id=self.resource_group_id,
            security_group_id=self.security_group_id,
            service_managed=self.service_managed,
            status=self.status,
            tags=self.tags,
            type=self.type,
            vpc_id=self.vpc_id,
            vswitch_id=self.vswitch_id)


def get_network_interfaces(ids: Optional[Sequence[str]] = None,
                           instance_id: Optional[str] = None,
                           name: Optional[str] = None,
                           name_regex: Optional[str] = None,
                           network_interface_name: Optional[str] = None,
                           output_file: Optional[str] = None,
                           primary_ip_address: Optional[str] = None,
                           private_ip: Optional[str] = None,
                           resource_group_id: Optional[str] = None,
                           security_group_id: Optional[str] = None,
                           service_managed: Optional[bool] = None,
                           status: Optional[str] = None,
                           tags: Optional[Mapping[str, str]] = None,
                           type: Optional[str] = None,
                           vpc_id: Optional[str] = None,
                           vswitch_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkInterfacesResult:
    """
    > **DEPRECATED:** This datasource has been renamed to ecs_get_ecs_network_interfaces from version 1.123.1.

    Use this data source to get a list of elastic network interfaces according to the specified filters in an Alibaba Cloud account.

    For information about elastic network interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "networkInterfacesName"
    vpc = alicloud.vpc.Network("vpc",
        vpc_name=name,
        cidr_block="192.168.0.0/24")
    default = alicloud.get_zones(available_resource_creation="VSwitch")
    vswitch = alicloud.vpc.Switch("vswitch",
        vswitch_name=name,
        cidr_block="192.168.0.0/24",
        availability_zone=default.zones[0].id,
        vpc_id=vpc.id)
    group = alicloud.ecs.SecurityGroup("group",
        name=name,
        vpc_id=vpc.id)
    interface = alicloud.vpc.NetworkInterface("interface",
        name=f"{name}%d",
        vswitch_id=vswitch.id,
        security_groups=[group.id],
        description="Basic test",
        private_ip="192.168.0.2",
        tags={
            "TF-VER": "0.11.3",
        })
    instance = alicloud.ecs.Instance("instance",
        availability_zone=default.zones[0].id,
        security_groups=[group.id],
        instance_type="ecs.e3.xlarge",
        system_disk_category="cloud_efficiency",
        image_id="centos_7_04_64_20G_alibase_201701015.vhd",
        instance_name=name,
        vswitch_id=vswitch.id,
        internet_max_bandwidth_out=10)
    attachment = alicloud.vpc.NetworkInterfaceAttachment("attachment",
        instance_id=instance.id,
        network_interface_id=interface.id)
    default_get_network_interfaces = alicloud.ecs.get_network_interfaces_output(ids=[attachment.network_interface_id],
        name_regex=name,
        tags={
            "TF-VER": "0.11.3",
        },
        vpc_id=vpc.id,
        vswitch_id=vswitch.id,
        private_ip="192.168.0.2",
        security_group_id=group.id,
        type="Secondary",
        instance_id=instance.id)
    pulumi.export("eni0Name", default_get_network_interfaces.interfaces[0].name)
    ```

    ## Argument Reference

    The following arguments are supported:

    * `ids` - (Optional)  A list of ENI IDs.
    * `name_regex` - (Optional) A regex string to filter results by ENI name.
    * `vpc_id` - (Optional) The VPC ID linked to ENIs.
    * `vswitch_id` - (Optional) The vSwitch ID linked to ENIs.
    * `private_ip` - (Optional) The primary private IP address of the ENI.
    * `security_group_id` - (Optional) The security group ID linked to ENIs.
    * `name` - (Optional) The name of the ENIs.
    * `type` - (Optional) The type of ENIs, Only support for "Primary" or "Secondary".
    * `instance_id` - (Optional) The ECS instance ID that the ENI is attached to.
    * `tags` - (Optional) A map of tags assigned to ENIs.
    * `output_file` - (Optional) The name of output file that saves the filter results.
    * `resource_group_id` - (Optional, ForceNew, Available in 1.57.0+) The Id of resource group which the network interface belongs.


    :param str instance_id: ID of the instance that the ENI is attached to.
    :param str name: Name of the ENI.
    :param str private_ip: Primary private IP of the ENI.
    :param str resource_group_id: The Id of resource group.
    :param str status: Current status of the ENI.
    :param Mapping[str, str] tags: A map of tags assigned to the ENI.
    :param str vpc_id: ID of the VPC that the ENI belongs to.
    :param str vswitch_id: ID of the vSwitch that the ENI is linked to.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['networkInterfaceName'] = network_interface_name
    __args__['outputFile'] = output_file
    __args__['primaryIpAddress'] = primary_ip_address
    __args__['privateIp'] = private_ip
    __args__['resourceGroupId'] = resource_group_id
    __args__['securityGroupId'] = security_group_id
    __args__['serviceManaged'] = service_managed
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['type'] = type
    __args__['vpcId'] = vpc_id
    __args__['vswitchId'] = vswitch_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getNetworkInterfaces:getNetworkInterfaces', __args__, opts=opts, typ=GetNetworkInterfacesResult).value

    return AwaitableGetNetworkInterfacesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        interfaces=pulumi.get(__ret__, 'interfaces'),
        name=pulumi.get(__ret__, 'name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        network_interface_name=pulumi.get(__ret__, 'network_interface_name'),
        output_file=pulumi.get(__ret__, 'output_file'),
        primary_ip_address=pulumi.get(__ret__, 'primary_ip_address'),
        private_ip=pulumi.get(__ret__, 'private_ip'),
        resource_group_id=pulumi.get(__ret__, 'resource_group_id'),
        security_group_id=pulumi.get(__ret__, 'security_group_id'),
        service_managed=pulumi.get(__ret__, 'service_managed'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'),
        vswitch_id=pulumi.get(__ret__, 'vswitch_id'))


@_utilities.lift_output_func(get_network_interfaces)
def get_network_interfaces_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  name: Optional[pulumi.Input[Optional[str]]] = None,
                                  name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                  network_interface_name: Optional[pulumi.Input[Optional[str]]] = None,
                                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  primary_ip_address: Optional[pulumi.Input[Optional[str]]] = None,
                                  private_ip: Optional[pulumi.Input[Optional[str]]] = None,
                                  resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  security_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  service_managed: Optional[pulumi.Input[Optional[bool]]] = None,
                                  status: Optional[pulumi.Input[Optional[str]]] = None,
                                  tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                                  type: Optional[pulumi.Input[Optional[str]]] = None,
                                  vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  vswitch_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkInterfacesResult]:
    """
    > **DEPRECATED:** This datasource has been renamed to ecs_get_ecs_network_interfaces from version 1.123.1.

    Use this data source to get a list of elastic network interfaces according to the specified filters in an Alibaba Cloud account.

    For information about elastic network interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "networkInterfacesName"
    vpc = alicloud.vpc.Network("vpc",
        vpc_name=name,
        cidr_block="192.168.0.0/24")
    default = alicloud.get_zones(available_resource_creation="VSwitch")
    vswitch = alicloud.vpc.Switch("vswitch",
        vswitch_name=name,
        cidr_block="192.168.0.0/24",
        availability_zone=default.zones[0].id,
        vpc_id=vpc.id)
    group = alicloud.ecs.SecurityGroup("group",
        name=name,
        vpc_id=vpc.id)
    interface = alicloud.vpc.NetworkInterface("interface",
        name=f"{name}%d",
        vswitch_id=vswitch.id,
        security_groups=[group.id],
        description="Basic test",
        private_ip="192.168.0.2",
        tags={
            "TF-VER": "0.11.3",
        })
    instance = alicloud.ecs.Instance("instance",
        availability_zone=default.zones[0].id,
        security_groups=[group.id],
        instance_type="ecs.e3.xlarge",
        system_disk_category="cloud_efficiency",
        image_id="centos_7_04_64_20G_alibase_201701015.vhd",
        instance_name=name,
        vswitch_id=vswitch.id,
        internet_max_bandwidth_out=10)
    attachment = alicloud.vpc.NetworkInterfaceAttachment("attachment",
        instance_id=instance.id,
        network_interface_id=interface.id)
    default_get_network_interfaces = alicloud.ecs.get_network_interfaces_output(ids=[attachment.network_interface_id],
        name_regex=name,
        tags={
            "TF-VER": "0.11.3",
        },
        vpc_id=vpc.id,
        vswitch_id=vswitch.id,
        private_ip="192.168.0.2",
        security_group_id=group.id,
        type="Secondary",
        instance_id=instance.id)
    pulumi.export("eni0Name", default_get_network_interfaces.interfaces[0].name)
    ```

    ## Argument Reference

    The following arguments are supported:

    * `ids` - (Optional)  A list of ENI IDs.
    * `name_regex` - (Optional) A regex string to filter results by ENI name.
    * `vpc_id` - (Optional) The VPC ID linked to ENIs.
    * `vswitch_id` - (Optional) The vSwitch ID linked to ENIs.
    * `private_ip` - (Optional) The primary private IP address of the ENI.
    * `security_group_id` - (Optional) The security group ID linked to ENIs.
    * `name` - (Optional) The name of the ENIs.
    * `type` - (Optional) The type of ENIs, Only support for "Primary" or "Secondary".
    * `instance_id` - (Optional) The ECS instance ID that the ENI is attached to.
    * `tags` - (Optional) A map of tags assigned to ENIs.
    * `output_file` - (Optional) The name of output file that saves the filter results.
    * `resource_group_id` - (Optional, ForceNew, Available in 1.57.0+) The Id of resource group which the network interface belongs.


    :param str instance_id: ID of the instance that the ENI is attached to.
    :param str name: Name of the ENI.
    :param str private_ip: Primary private IP of the ENI.
    :param str resource_group_id: The Id of resource group.
    :param str status: Current status of the ENI.
    :param Mapping[str, str] tags: A map of tags assigned to the ENI.
    :param str vpc_id: ID of the VPC that the ENI belongs to.
    :param str vswitch_id: ID of the vSwitch that the ENI is linked to.
    """
    ...
