# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DbInstanceEndpointAddressArgs', 'DbInstanceEndpointAddress']

@pulumi.input_type
class DbInstanceEndpointAddressArgs:
    def __init__(__self__, *,
                 connection_string_prefix: pulumi.Input[str],
                 db_instance_endpoint_id: pulumi.Input[str],
                 db_instance_id: pulumi.Input[str],
                 port: pulumi.Input[str]):
        """
        The set of arguments for constructing a DbInstanceEndpointAddress resource.
        :param pulumi.Input[str] connection_string_prefix: The prefix of the public endpoint.
        :param pulumi.Input[str] db_instance_endpoint_id: The Endpoint ID of the instance.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[str] port: The port number of the public endpoint.
        """
        pulumi.set(__self__, "connection_string_prefix", connection_string_prefix)
        pulumi.set(__self__, "db_instance_endpoint_id", db_instance_endpoint_id)
        pulumi.set(__self__, "db_instance_id", db_instance_id)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="connectionStringPrefix")
    def connection_string_prefix(self) -> pulumi.Input[str]:
        """
        The prefix of the public endpoint.
        """
        return pulumi.get(self, "connection_string_prefix")

    @connection_string_prefix.setter
    def connection_string_prefix(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_string_prefix", value)

    @property
    @pulumi.getter(name="dbInstanceEndpointId")
    def db_instance_endpoint_id(self) -> pulumi.Input[str]:
        """
        The Endpoint ID of the instance.
        """
        return pulumi.get(self, "db_instance_endpoint_id")

    @db_instance_endpoint_id.setter
    def db_instance_endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_instance_endpoint_id", value)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[str]:
        """
        The port number of the public endpoint.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[str]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class _DbInstanceEndpointAddressState:
    def __init__(__self__, *,
                 connection_string: Optional[pulumi.Input[str]] = None,
                 connection_string_prefix: Optional[pulumi.Input[str]] = None,
                 db_instance_endpoint_id: Optional[pulumi.Input[str]] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ip_type: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DbInstanceEndpointAddress resources.
        :param pulumi.Input[str] connection_string: The endpoint of the instance.
        :param pulumi.Input[str] connection_string_prefix: The prefix of the public endpoint.
        :param pulumi.Input[str] db_instance_endpoint_id: The Endpoint ID of the instance.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[str] ip_address: The IP address of the endpoint.
        :param pulumi.Input[str] ip_type: The type of the IP address.
        :param pulumi.Input[str] port: The port number of the public endpoint.
        """
        if connection_string is not None:
            pulumi.set(__self__, "connection_string", connection_string)
        if connection_string_prefix is not None:
            pulumi.set(__self__, "connection_string_prefix", connection_string_prefix)
        if db_instance_endpoint_id is not None:
            pulumi.set(__self__, "db_instance_endpoint_id", db_instance_endpoint_id)
        if db_instance_id is not None:
            pulumi.set(__self__, "db_instance_id", db_instance_id)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if ip_type is not None:
            pulumi.set(__self__, "ip_type", ip_type)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint of the instance.
        """
        return pulumi.get(self, "connection_string")

    @connection_string.setter
    def connection_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string", value)

    @property
    @pulumi.getter(name="connectionStringPrefix")
    def connection_string_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The prefix of the public endpoint.
        """
        return pulumi.get(self, "connection_string_prefix")

    @connection_string_prefix.setter
    def connection_string_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_string_prefix", value)

    @property
    @pulumi.getter(name="dbInstanceEndpointId")
    def db_instance_endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Endpoint ID of the instance.
        """
        return pulumi.get(self, "db_instance_endpoint_id")

    @db_instance_endpoint_id.setter
    def db_instance_endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_endpoint_id", value)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the endpoint.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the IP address.
        """
        return pulumi.get(self, "ip_type")

    @ip_type.setter
    def ip_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_type", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        The port number of the public endpoint.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)


class DbInstanceEndpointAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_string_prefix: Optional[pulumi.Input[str]] = None,
                 db_instance_endpoint_id: Optional[pulumi.Input[str]] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provide RDS cluster instance endpoint public connection resources.

        Information about RDS MySQL cluster endpoint address and how to use it, see [What is RDS DB Instance Endpoint Address](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-createdbinstanceendpointaddress).

        > **NOTE:** Available since v1.204.0.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_zones = alicloud.rds.get_zones(engine="MySQL",
            engine_version="8.0",
            instance_charge_type="PostPaid",
            category="cluster",
            db_instance_storage_type="cloud_essd")
        default_instance_classes = alicloud.rds.get_instance_classes(zone_id=default_zones.ids[0],
            engine="MySQL",
            engine_version="8.0",
            category="cluster",
            db_instance_storage_type="cloud_essd",
            instance_charge_type="PostPaid")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default_zones.ids[0],
            vswitch_name=name)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_instance = alicloud.rds.Instance("defaultInstance",
            engine="MySQL",
            engine_version="8.0",
            instance_type=default_instance_classes.instance_classes[0].instance_class,
            instance_storage=default_instance_classes.instance_classes[0].storage_range.min,
            instance_charge_type="Postpaid",
            instance_name=name,
            vswitch_id=default_switch.id,
            monitoring_period=60,
            db_instance_storage_type="cloud_essd",
            security_group_ids=[default_security_group.id],
            zone_id=default_zones.ids[0],
            zone_id_slave_a=default_zones.ids[0])
        default_db_node = alicloud.rds.DbNode("defaultDbNode",
            db_instance_id=default_instance.id,
            class_code=default_instance.instance_type,
            zone_id=default_switch.zone_id)
        default_db_instance_endpoint = alicloud.rds.DbInstanceEndpoint("defaultDbInstanceEndpoint",
            db_instance_id=default_db_node.db_instance_id,
            vpc_id=default_network.id,
            vswitch_id=default_instance.vswitch_id,
            connection_string_prefix="example",
            port="3306",
            db_instance_endpoint_description=name,
            node_items=[alicloud.rds.DbInstanceEndpointNodeItemArgs(
                node_id=default_db_node.node_id,
                weight=25,
            )])
        default_db_instance_endpoint_address = alicloud.rds.DbInstanceEndpointAddress("defaultDbInstanceEndpointAddress",
            db_instance_id=default_instance.id,
            db_instance_endpoint_id=default_db_instance_endpoint.db_instance_endpoint_id,
            connection_string_prefix="tf-example-prefix",
            port="3306")
        ```

        ## Import

        RDS database endpoint public address feature can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:rds/dbInstanceEndpointAddress:DbInstanceEndpointAddress example <db_instance_id>:<db_instance_endpoint_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_string_prefix: The prefix of the public endpoint.
        :param pulumi.Input[str] db_instance_endpoint_id: The Endpoint ID of the instance.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[str] port: The port number of the public endpoint.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DbInstanceEndpointAddressArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide RDS cluster instance endpoint public connection resources.

        Information about RDS MySQL cluster endpoint address and how to use it, see [What is RDS DB Instance Endpoint Address](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-createdbinstanceendpointaddress).

        > **NOTE:** Available since v1.204.0.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_zones = alicloud.rds.get_zones(engine="MySQL",
            engine_version="8.0",
            instance_charge_type="PostPaid",
            category="cluster",
            db_instance_storage_type="cloud_essd")
        default_instance_classes = alicloud.rds.get_instance_classes(zone_id=default_zones.ids[0],
            engine="MySQL",
            engine_version="8.0",
            category="cluster",
            db_instance_storage_type="cloud_essd",
            instance_charge_type="PostPaid")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name=name,
            cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            zone_id=default_zones.ids[0],
            vswitch_name=name)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_instance = alicloud.rds.Instance("defaultInstance",
            engine="MySQL",
            engine_version="8.0",
            instance_type=default_instance_classes.instance_classes[0].instance_class,
            instance_storage=default_instance_classes.instance_classes[0].storage_range.min,
            instance_charge_type="Postpaid",
            instance_name=name,
            vswitch_id=default_switch.id,
            monitoring_period=60,
            db_instance_storage_type="cloud_essd",
            security_group_ids=[default_security_group.id],
            zone_id=default_zones.ids[0],
            zone_id_slave_a=default_zones.ids[0])
        default_db_node = alicloud.rds.DbNode("defaultDbNode",
            db_instance_id=default_instance.id,
            class_code=default_instance.instance_type,
            zone_id=default_switch.zone_id)
        default_db_instance_endpoint = alicloud.rds.DbInstanceEndpoint("defaultDbInstanceEndpoint",
            db_instance_id=default_db_node.db_instance_id,
            vpc_id=default_network.id,
            vswitch_id=default_instance.vswitch_id,
            connection_string_prefix="example",
            port="3306",
            db_instance_endpoint_description=name,
            node_items=[alicloud.rds.DbInstanceEndpointNodeItemArgs(
                node_id=default_db_node.node_id,
                weight=25,
            )])
        default_db_instance_endpoint_address = alicloud.rds.DbInstanceEndpointAddress("defaultDbInstanceEndpointAddress",
            db_instance_id=default_instance.id,
            db_instance_endpoint_id=default_db_instance_endpoint.db_instance_endpoint_id,
            connection_string_prefix="tf-example-prefix",
            port="3306")
        ```

        ## Import

        RDS database endpoint public address feature can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:rds/dbInstanceEndpointAddress:DbInstanceEndpointAddress example <db_instance_id>:<db_instance_endpoint_id>
        ```

        :param str resource_name: The name of the resource.
        :param DbInstanceEndpointAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DbInstanceEndpointAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_string_prefix: Optional[pulumi.Input[str]] = None,
                 db_instance_endpoint_id: Optional[pulumi.Input[str]] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DbInstanceEndpointAddressArgs.__new__(DbInstanceEndpointAddressArgs)

            if connection_string_prefix is None and not opts.urn:
                raise TypeError("Missing required property 'connection_string_prefix'")
            __props__.__dict__["connection_string_prefix"] = connection_string_prefix
            if db_instance_endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_instance_endpoint_id'")
            __props__.__dict__["db_instance_endpoint_id"] = db_instance_endpoint_id
            if db_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_instance_id'")
            __props__.__dict__["db_instance_id"] = db_instance_id
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            __props__.__dict__["connection_string"] = None
            __props__.__dict__["ip_address"] = None
            __props__.__dict__["ip_type"] = None
        super(DbInstanceEndpointAddress, __self__).__init__(
            'alicloud:rds/dbInstanceEndpointAddress:DbInstanceEndpointAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_string: Optional[pulumi.Input[str]] = None,
            connection_string_prefix: Optional[pulumi.Input[str]] = None,
            db_instance_endpoint_id: Optional[pulumi.Input[str]] = None,
            db_instance_id: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            ip_type: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[str]] = None) -> 'DbInstanceEndpointAddress':
        """
        Get an existing DbInstanceEndpointAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_string: The endpoint of the instance.
        :param pulumi.Input[str] connection_string_prefix: The prefix of the public endpoint.
        :param pulumi.Input[str] db_instance_endpoint_id: The Endpoint ID of the instance.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[str] ip_address: The IP address of the endpoint.
        :param pulumi.Input[str] ip_type: The type of the IP address.
        :param pulumi.Input[str] port: The port number of the public endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DbInstanceEndpointAddressState.__new__(_DbInstanceEndpointAddressState)

        __props__.__dict__["connection_string"] = connection_string
        __props__.__dict__["connection_string_prefix"] = connection_string_prefix
        __props__.__dict__["db_instance_endpoint_id"] = db_instance_endpoint_id
        __props__.__dict__["db_instance_id"] = db_instance_id
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["ip_type"] = ip_type
        __props__.__dict__["port"] = port
        return DbInstanceEndpointAddress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> pulumi.Output[str]:
        """
        The endpoint of the instance.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter(name="connectionStringPrefix")
    def connection_string_prefix(self) -> pulumi.Output[str]:
        """
        The prefix of the public endpoint.
        """
        return pulumi.get(self, "connection_string_prefix")

    @property
    @pulumi.getter(name="dbInstanceEndpointId")
    def db_instance_endpoint_id(self) -> pulumi.Output[str]:
        """
        The Endpoint ID of the instance.
        """
        return pulumi.get(self, "db_instance_endpoint_id")

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "db_instance_id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The IP address of the endpoint.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> pulumi.Output[str]:
        """
        The type of the IP address.
        """
        return pulumi.get(self, "ip_type")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[str]:
        """
        The port number of the public endpoint.
        """
        return pulumi.get(self, "port")

