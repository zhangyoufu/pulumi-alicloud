# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MasterSlaveServerGroupArgs', 'MasterSlaveServerGroup']

@pulumi.input_type
class MasterSlaveServerGroupArgs:
    def __init__(__self__, *,
                 load_balancer_id: pulumi.Input[str],
                 delete_protection_validation: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input['MasterSlaveServerGroupServerArgs']]]] = None):
        """
        The set of arguments for constructing a MasterSlaveServerGroup resource.
        :param pulumi.Input[str] load_balancer_id: The Load Balancer ID which is used to launch a new master slave server group.
        :param pulumi.Input[bool] delete_protection_validation: Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        :param pulumi.Input[str] name: Name of the master slave server group.
        :param pulumi.Input[Sequence[pulumi.Input['MasterSlaveServerGroupServerArgs']]] servers: A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        """
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if delete_protection_validation is not None:
            pulumi.set(__self__, "delete_protection_validation", delete_protection_validation)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if servers is not None:
            pulumi.set(__self__, "servers", servers)

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> pulumi.Input[str]:
        """
        The Load Balancer ID which is used to launch a new master slave server group.
        """
        return pulumi.get(self, "load_balancer_id")

    @load_balancer_id.setter
    def load_balancer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "load_balancer_id", value)

    @property
    @pulumi.getter(name="deleteProtectionValidation")
    def delete_protection_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        """
        return pulumi.get(self, "delete_protection_validation")

    @delete_protection_validation.setter
    def delete_protection_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_protection_validation", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the master slave server group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MasterSlaveServerGroupServerArgs']]]]:
        """
        A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        """
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MasterSlaveServerGroupServerArgs']]]]):
        pulumi.set(self, "servers", value)


@pulumi.input_type
class _MasterSlaveServerGroupState:
    def __init__(__self__, *,
                 delete_protection_validation: Optional[pulumi.Input[bool]] = None,
                 load_balancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input['MasterSlaveServerGroupServerArgs']]]] = None):
        """
        Input properties used for looking up and filtering MasterSlaveServerGroup resources.
        :param pulumi.Input[bool] delete_protection_validation: Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        :param pulumi.Input[str] load_balancer_id: The Load Balancer ID which is used to launch a new master slave server group.
        :param pulumi.Input[str] name: Name of the master slave server group.
        :param pulumi.Input[Sequence[pulumi.Input['MasterSlaveServerGroupServerArgs']]] servers: A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        """
        if delete_protection_validation is not None:
            pulumi.set(__self__, "delete_protection_validation", delete_protection_validation)
        if load_balancer_id is not None:
            pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if servers is not None:
            pulumi.set(__self__, "servers", servers)

    @property
    @pulumi.getter(name="deleteProtectionValidation")
    def delete_protection_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        """
        return pulumi.get(self, "delete_protection_validation")

    @delete_protection_validation.setter
    def delete_protection_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_protection_validation", value)

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Load Balancer ID which is used to launch a new master slave server group.
        """
        return pulumi.get(self, "load_balancer_id")

    @load_balancer_id.setter
    def load_balancer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_balancer_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the master slave server group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MasterSlaveServerGroupServerArgs']]]]:
        """
        A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        """
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MasterSlaveServerGroupServerArgs']]]]):
        pulumi.set(self, "servers", value)


class MasterSlaveServerGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_protection_validation: Optional[pulumi.Input[bool]] = None,
                 load_balancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MasterSlaveServerGroupServerArgs']]]]] = None,
                 __props__=None):
        """
        A master slave server group contains two ECS instances. The master slave server group can help you to define multiple listening dimension.

        > **NOTE:** One ECS instance can be added into multiple master slave server groups.

        > **NOTE:** One master slave server group can only add two ECS instances, which are master server and slave server.

        > **NOTE:** One master slave server group can be attached with tcp/udp listeners in one load balancer.

        > **NOTE:** One Classic and Internet load balancer, its master slave server group can add Classic and VPC ECS instances.

        > **NOTE:** One Classic and Intranet load balancer, its master slave server group can only add Classic ECS instances.

        > **NOTE:** One VPC load balancer, its master slave server group can only add the same VPC ECS instances.

        > **NOTE:** Available in 1.54.0+

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        default_instance_types = alicloud.ecs.get_instance_types(availability_zone=default_zones.zones[0].id,
            eni_amount=2)
        image = alicloud.ecs.get_images(name_regex="^ubuntu_18.*64",
            most_recent=True,
            owners="system")
        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testAccSlbMasterSlaveServerGroupVpc"
        number = config.get("number")
        if number is None:
            number = "1"
        main_network = alicloud.vpc.Network("mainNetwork", cidr_block="172.16.0.0/16")
        main_switch = alicloud.vpc.Switch("mainSwitch",
            vpc_id=main_network.id,
            cidr_block="172.16.0.0/16",
            zone_id=default_zones.zones[0].id,
            vswitch_name=name)
        group_security_group = alicloud.ecs.SecurityGroup("groupSecurityGroup", vpc_id=main_network.id)
        instance_instance = []
        for range in [{"value": i} for i in range(0, 2)]:
            instance_instance.append(alicloud.ecs.Instance(f"instanceInstance-{range['value']}",
                image_id=image.images[0].id,
                instance_type=default_instance_types.instance_types[0].id,
                instance_name=name,
                security_groups=[group_security_group.id],
                internet_charge_type="PayByTraffic",
                internet_max_bandwidth_out=10,
                availability_zone=default_zones.zones[0].id,
                instance_charge_type="PostPaid",
                system_disk_category="cloud_efficiency",
                vswitch_id=main_switch.id))
        instance_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("instanceApplicationLoadBalancer",
            load_balancer_name=name,
            vswitch_id=main_switch.id,
            load_balancer_spec="slb.s2.small")
        default_network_interface = []
        for range in [{"value": i} for i in range(0, number)]:
            default_network_interface.append(alicloud.vpc.NetworkInterface(f"defaultNetworkInterface-{range['value']}",
                vswitch_id=main_switch.id,
                security_groups=[group_security_group.id]))
        default_network_interface_attachment = []
        for range in [{"value": i} for i in range(0, number)]:
            default_network_interface_attachment.append(alicloud.vpc.NetworkInterfaceAttachment(f"defaultNetworkInterfaceAttachment-{range['value']}",
                instance_id=instance_instance[0].id,
                network_interface_id=[__item.id for __item in default_network_interface][range["index"]]))
        group_master_slave_server_group = alicloud.slb.MasterSlaveServerGroup("groupMasterSlaveServerGroup",
            load_balancer_id=instance_application_load_balancer.id,
            servers=[
                alicloud.slb.MasterSlaveServerGroupServerArgs(
                    server_id=instance_instance[0].id,
                    port=100,
                    weight=100,
                    server_type="Master",
                ),
                alicloud.slb.MasterSlaveServerGroupServerArgs(
                    server_id=instance_instance[1].id,
                    port=100,
                    weight=100,
                    server_type="Slave",
                ),
            ])
        tcp = alicloud.slb.Listener("tcp",
            load_balancer_id=instance_application_load_balancer.id,
            master_slave_server_group_id=group_master_slave_server_group.id,
            frontend_port=22,
            protocol="tcp",
            bandwidth=10,
            health_check_type="tcp",
            persistence_timeout=3600,
            healthy_threshold=8,
            unhealthy_threshold=8,
            health_check_timeout=8,
            health_check_interval=5,
            health_check_http_code="http_2xx",
            health_check_connect_port=20,
            health_check_uri="/console",
            established_timeout=600)
        ```
        ## Block servers

        The servers mapping supports the following:

        * `server_ids` - (Required) A list backend server ID (ECS instance ID).
        * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
        * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
        * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
        * `server_type` - (Optional) The server type of the backend server. Valid value Master, Slave.
        * `is_backup` - (Removed from v1.63.0) Determine if the server is executing. Valid value 0, 1.

        ## Import

        Load balancer master slave server group can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_protection_validation: Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        :param pulumi.Input[str] load_balancer_id: The Load Balancer ID which is used to launch a new master slave server group.
        :param pulumi.Input[str] name: Name of the master slave server group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MasterSlaveServerGroupServerArgs']]]] servers: A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MasterSlaveServerGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A master slave server group contains two ECS instances. The master slave server group can help you to define multiple listening dimension.

        > **NOTE:** One ECS instance can be added into multiple master slave server groups.

        > **NOTE:** One master slave server group can only add two ECS instances, which are master server and slave server.

        > **NOTE:** One master slave server group can be attached with tcp/udp listeners in one load balancer.

        > **NOTE:** One Classic and Internet load balancer, its master slave server group can add Classic and VPC ECS instances.

        > **NOTE:** One Classic and Intranet load balancer, its master slave server group can only add Classic ECS instances.

        > **NOTE:** One VPC load balancer, its master slave server group can only add the same VPC ECS instances.

        > **NOTE:** Available in 1.54.0+

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        default_instance_types = alicloud.ecs.get_instance_types(availability_zone=default_zones.zones[0].id,
            eni_amount=2)
        image = alicloud.ecs.get_images(name_regex="^ubuntu_18.*64",
            most_recent=True,
            owners="system")
        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-testAccSlbMasterSlaveServerGroupVpc"
        number = config.get("number")
        if number is None:
            number = "1"
        main_network = alicloud.vpc.Network("mainNetwork", cidr_block="172.16.0.0/16")
        main_switch = alicloud.vpc.Switch("mainSwitch",
            vpc_id=main_network.id,
            cidr_block="172.16.0.0/16",
            zone_id=default_zones.zones[0].id,
            vswitch_name=name)
        group_security_group = alicloud.ecs.SecurityGroup("groupSecurityGroup", vpc_id=main_network.id)
        instance_instance = []
        for range in [{"value": i} for i in range(0, 2)]:
            instance_instance.append(alicloud.ecs.Instance(f"instanceInstance-{range['value']}",
                image_id=image.images[0].id,
                instance_type=default_instance_types.instance_types[0].id,
                instance_name=name,
                security_groups=[group_security_group.id],
                internet_charge_type="PayByTraffic",
                internet_max_bandwidth_out=10,
                availability_zone=default_zones.zones[0].id,
                instance_charge_type="PostPaid",
                system_disk_category="cloud_efficiency",
                vswitch_id=main_switch.id))
        instance_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("instanceApplicationLoadBalancer",
            load_balancer_name=name,
            vswitch_id=main_switch.id,
            load_balancer_spec="slb.s2.small")
        default_network_interface = []
        for range in [{"value": i} for i in range(0, number)]:
            default_network_interface.append(alicloud.vpc.NetworkInterface(f"defaultNetworkInterface-{range['value']}",
                vswitch_id=main_switch.id,
                security_groups=[group_security_group.id]))
        default_network_interface_attachment = []
        for range in [{"value": i} for i in range(0, number)]:
            default_network_interface_attachment.append(alicloud.vpc.NetworkInterfaceAttachment(f"defaultNetworkInterfaceAttachment-{range['value']}",
                instance_id=instance_instance[0].id,
                network_interface_id=[__item.id for __item in default_network_interface][range["index"]]))
        group_master_slave_server_group = alicloud.slb.MasterSlaveServerGroup("groupMasterSlaveServerGroup",
            load_balancer_id=instance_application_load_balancer.id,
            servers=[
                alicloud.slb.MasterSlaveServerGroupServerArgs(
                    server_id=instance_instance[0].id,
                    port=100,
                    weight=100,
                    server_type="Master",
                ),
                alicloud.slb.MasterSlaveServerGroupServerArgs(
                    server_id=instance_instance[1].id,
                    port=100,
                    weight=100,
                    server_type="Slave",
                ),
            ])
        tcp = alicloud.slb.Listener("tcp",
            load_balancer_id=instance_application_load_balancer.id,
            master_slave_server_group_id=group_master_slave_server_group.id,
            frontend_port=22,
            protocol="tcp",
            bandwidth=10,
            health_check_type="tcp",
            persistence_timeout=3600,
            healthy_threshold=8,
            unhealthy_threshold=8,
            health_check_timeout=8,
            health_check_interval=5,
            health_check_http_code="http_2xx",
            health_check_connect_port=20,
            health_check_uri="/console",
            established_timeout=600)
        ```
        ## Block servers

        The servers mapping supports the following:

        * `server_ids` - (Required) A list backend server ID (ECS instance ID).
        * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
        * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
        * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
        * `server_type` - (Optional) The server type of the backend server. Valid value Master, Slave.
        * `is_backup` - (Removed from v1.63.0) Determine if the server is executing. Valid value 0, 1.

        ## Import

        Load balancer master slave server group can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup example abc123456
        ```

        :param str resource_name: The name of the resource.
        :param MasterSlaveServerGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MasterSlaveServerGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_protection_validation: Optional[pulumi.Input[bool]] = None,
                 load_balancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MasterSlaveServerGroupServerArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MasterSlaveServerGroupArgs.__new__(MasterSlaveServerGroupArgs)

            __props__.__dict__["delete_protection_validation"] = delete_protection_validation
            if load_balancer_id is None and not opts.urn:
                raise TypeError("Missing required property 'load_balancer_id'")
            __props__.__dict__["load_balancer_id"] = load_balancer_id
            __props__.__dict__["name"] = name
            __props__.__dict__["servers"] = servers
        super(MasterSlaveServerGroup, __self__).__init__(
            'alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delete_protection_validation: Optional[pulumi.Input[bool]] = None,
            load_balancer_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MasterSlaveServerGroupServerArgs']]]]] = None) -> 'MasterSlaveServerGroup':
        """
        Get an existing MasterSlaveServerGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_protection_validation: Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        :param pulumi.Input[str] load_balancer_id: The Load Balancer ID which is used to launch a new master slave server group.
        :param pulumi.Input[str] name: Name of the master slave server group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MasterSlaveServerGroupServerArgs']]]] servers: A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MasterSlaveServerGroupState.__new__(_MasterSlaveServerGroupState)

        __props__.__dict__["delete_protection_validation"] = delete_protection_validation
        __props__.__dict__["load_balancer_id"] = load_balancer_id
        __props__.__dict__["name"] = name
        __props__.__dict__["servers"] = servers
        return MasterSlaveServerGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deleteProtectionValidation")
    def delete_protection_validation(self) -> pulumi.Output[Optional[bool]]:
        """
        Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        """
        return pulumi.get(self, "delete_protection_validation")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> pulumi.Output[str]:
        """
        The Load Balancer ID which is used to launch a new master slave server group.
        """
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the master slave server group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def servers(self) -> pulumi.Output[Optional[Sequence['outputs.MasterSlaveServerGroupServer']]]:
        """
        A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
        """
        return pulumi.get(self, "servers")

