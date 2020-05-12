# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class NetworkInterface(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
    """
    mac: pulumi.Output[str]
    """
    (Available in 1.54.0+) The MAC address of an ENI.
    """
    name: pulumi.Output[str]
    """
    Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
    """
    private_ip: pulumi.Output[str]
    """
    The primary private IP of the ENI.
    """
    private_ips: pulumi.Output[list]
    """
    List of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
    """
    private_ips_count: pulumi.Output[float]
    """
    Number of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
    """
    resource_group_id: pulumi.Output[str]
    """
    The Id of resource group which the network interface belongs.
    """
    security_groups: pulumi.Output[list]
    """
    A list of security group ids to associate with.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vswitch_id: pulumi.Output[str]
    """
    The VSwitch to create the ENI in.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, private_ip=None, private_ips=None, private_ips_count=None, resource_group_id=None, security_groups=None, tags=None, vswitch_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an ECS Elastic Network Interface resource.

        For information about Elastic Network Interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html).

        > **NOTE** Only one of private_ips or private_ips_count can be specified when assign private IPs. 

        ## Example Usage



        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "networkInterfaceName"
        vpc = alicloud.vpc.Network("vpc", cidr_block="192.168.0.0/24")
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        vswitch = alicloud.vpc.Switch("vswitch",
            availability_zone=default_zones.zones[0]["id"],
            cidr_block="192.168.0.0/24",
            vpc_id=vpc.id)
        group = alicloud.ecs.SecurityGroup("group", vpc_id=vpc.id)
        default_network_interface = alicloud.vpc.NetworkInterface("defaultNetworkInterface",
            private_ip="192.168.0.2",
            private_ips_count=3,
            security_groups=[group.id],
            vswitch_id=vswitch.id)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        :param pulumi.Input[str] name: Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
        :param pulumi.Input[str] private_ip: The primary private IP of the ENI.
        :param pulumi.Input[list] private_ips: List of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        :param pulumi.Input[float] private_ips_count: Number of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the network interface belongs.
        :param pulumi.Input[list] security_groups: A list of security group ids to associate with.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: The VSwitch to create the ENI in.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['name'] = name
            __props__['private_ip'] = private_ip
            __props__['private_ips'] = private_ips
            __props__['private_ips_count'] = private_ips_count
            __props__['resource_group_id'] = resource_group_id
            if security_groups is None:
                raise TypeError("Missing required property 'security_groups'")
            __props__['security_groups'] = security_groups
            __props__['tags'] = tags
            if vswitch_id is None:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__['vswitch_id'] = vswitch_id
            __props__['mac'] = None
        super(NetworkInterface, __self__).__init__(
            'alicloud:vpc/networkInterface:NetworkInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, mac=None, name=None, private_ip=None, private_ips=None, private_ips_count=None, resource_group_id=None, security_groups=None, tags=None, vswitch_id=None):
        """
        Get an existing NetworkInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        :param pulumi.Input[str] mac: (Available in 1.54.0+) The MAC address of an ENI.
        :param pulumi.Input[str] name: Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
        :param pulumi.Input[str] private_ip: The primary private IP of the ENI.
        :param pulumi.Input[list] private_ips: List of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        :param pulumi.Input[float] private_ips_count: Number of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the network interface belongs.
        :param pulumi.Input[list] security_groups: A list of security group ids to associate with.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: The VSwitch to create the ENI in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["mac"] = mac
        __props__["name"] = name
        __props__["private_ip"] = private_ip
        __props__["private_ips"] = private_ips
        __props__["private_ips_count"] = private_ips_count
        __props__["resource_group_id"] = resource_group_id
        __props__["security_groups"] = security_groups
        __props__["tags"] = tags
        __props__["vswitch_id"] = vswitch_id
        return NetworkInterface(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

