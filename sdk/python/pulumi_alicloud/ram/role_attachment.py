# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['RoleAttachment']


class RoleAttachment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a RAM role attachment resource to bind role for several ECS instances.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        default_instance_types = alicloud.ecs.get_instance_types(availability_zone=default_zones.zones[0].id,
            cpu_core_count=2,
            memory_size=4)
        default_images = alicloud.ecs.get_images(most_recent=True,
            name_regex="^ubuntu_18.*64",
            owners="system")
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            availability_zone=default_zones.zones[0].id,
            cidr_block="172.16.0.0/24",
            vpc_id=default_network.id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_security_group_rule = alicloud.ecs.SecurityGroupRule("defaultSecurityGroupRule",
            cidr_ip="172.16.0.0/24",
            ip_protocol="tcp",
            nic_type="intranet",
            policy="accept",
            port_range="22/22",
            priority=1,
            security_group_id=default_security_group.id,
            type="ingress")
        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "ecsInstanceVPCExample"
        foo = alicloud.ecs.Instance("foo",
            image_id=default_images.images[0].id,
            instance_name=name,
            instance_type=default_instance_types.instance_types[0].id,
            internet_charge_type="PayByTraffic",
            internet_max_bandwidth_out=5,
            security_groups=[default_security_group.id],
            system_disk_category="cloud_efficiency",
            vswitch_id=default_switch.id)
        role = alicloud.ram.Role("role",
            description="this is a test",
            document=\"\"\"  {
            "Statement": [
              {
                "Action": "sts:AssumeRole",
                "Effect": "Allow",
                "Principal": {
                  "Service": [
                    "ecs.aliyuncs.com"
                  ]
                }
              }
            ],
            "Version": "1"
          }
          
        \"\"\",
            force=True)
        attach = alicloud.ram.RoleAttachment("attach",
            instance_ids=[[__item.id for __item in [foo]]],
            role_name=role.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] instance_ids: The list of ECS instance's IDs.
        :param pulumi.Input[str] role_name: The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
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

            if instance_ids is None:
                raise TypeError("Missing required property 'instance_ids'")
            __props__['instance_ids'] = instance_ids
            if role_name is None:
                raise TypeError("Missing required property 'role_name'")
            __props__['role_name'] = role_name
        super(RoleAttachment, __self__).__init__(
            'alicloud:ram/roleAttachment:RoleAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            role_name: Optional[pulumi.Input[str]] = None) -> 'RoleAttachment':
        """
        Get an existing RoleAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] instance_ids: The list of ECS instance's IDs.
        :param pulumi.Input[str] role_name: The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["instance_ids"] = instance_ids
        __props__["role_name"] = role_name
        return RoleAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> List[str]:
        """
        The list of ECS instance's IDs.
        """
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> str:
        """
        The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        """
        return pulumi.get(self, "role_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

