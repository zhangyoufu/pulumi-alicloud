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
from ._inputs import *

__all__ = ['ShardingNetworkPublicAddressArgs', 'ShardingNetworkPublicAddress']

@pulumi.input_type
class ShardingNetworkPublicAddressArgs:
    def __init__(__self__, *,
                 db_instance_id: pulumi.Input[str],
                 node_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ShardingNetworkPublicAddress resource.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[str] node_id: The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
        """
        pulumi.set(__self__, "db_instance_id", db_instance_id)
        pulumi.set(__self__, "node_id", node_id)

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
    @pulumi.getter(name="nodeId")
    def node_id(self) -> pulumi.Input[str]:
        """
        The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_id", value)


@pulumi.input_type
class _ShardingNetworkPublicAddressState:
    def __init__(__self__, *,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 network_addresses: Optional[pulumi.Input[Sequence[pulumi.Input['ShardingNetworkPublicAddressNetworkAddressArgs']]]] = None,
                 node_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ShardingNetworkPublicAddress resources.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[Sequence[pulumi.Input['ShardingNetworkPublicAddressNetworkAddressArgs']]] network_addresses: The endpoint of the instance.
        :param pulumi.Input[str] node_id: The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
        """
        if db_instance_id is not None:
            pulumi.set(__self__, "db_instance_id", db_instance_id)
        if network_addresses is not None:
            pulumi.set(__self__, "network_addresses", network_addresses)
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)

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
    @pulumi.getter(name="networkAddresses")
    def network_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ShardingNetworkPublicAddressNetworkAddressArgs']]]]:
        """
        The endpoint of the instance.
        """
        return pulumi.get(self, "network_addresses")

    @network_addresses.setter
    def network_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ShardingNetworkPublicAddressNetworkAddressArgs']]]]):
        pulumi.set(self, "network_addresses", value)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)


class ShardingNetworkPublicAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 node_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a MongoDB Sharding Network Public Address resource.

        For information about MongoDB Sharding Network Public Address and how to use it, see [What is Sharding Network Public Address](https://www.alibabacloud.com/help/doc-detail/67602.html).

        > **NOTE:** Available since v1.149.0.

        > **NOTE:** This operation supports sharded cluster instances only.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.mongodb.get_zones()
        index = len(default.zones) - 1
        zone_id = default.zones[index].id
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name=name,
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=zone_id)
        default_sharding_instance = alicloud.mongodb.ShardingInstance("default",
            zone_id=zone_id,
            vswitch_id=default_switch.id,
            engine_version="4.2",
            name=name,
            shard_lists=[
                {
                    "node_class": "dds.shard.mid",
                    "node_storage": 10,
                },
                {
                    "node_class": "dds.shard.standard",
                    "node_storage": 20,
                    "readonly_replicas": 1,
                },
            ],
            mongo_lists=[
                {
                    "node_class": "dds.mongos.mid",
                },
                {
                    "node_class": "dds.mongos.mid",
                },
            ])
        example = alicloud.mongodb.ShardingNetworkPublicAddress("example",
            db_instance_id=default_sharding_instance.id,
            node_id=default_sharding_instance.mongo_lists[0].node_id)
        ```

        ## Import

        MongoDB Sharding Network Public Address can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:mongodb/shardingNetworkPublicAddress:ShardingNetworkPublicAddress example <db_instance_id>:<node_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[str] node_id: The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ShardingNetworkPublicAddressArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a MongoDB Sharding Network Public Address resource.

        For information about MongoDB Sharding Network Public Address and how to use it, see [What is Sharding Network Public Address](https://www.alibabacloud.com/help/doc-detail/67602.html).

        > **NOTE:** Available since v1.149.0.

        > **NOTE:** This operation supports sharded cluster instances only.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.mongodb.get_zones()
        index = len(default.zones) - 1
        zone_id = default.zones[index].id
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name=name,
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=zone_id)
        default_sharding_instance = alicloud.mongodb.ShardingInstance("default",
            zone_id=zone_id,
            vswitch_id=default_switch.id,
            engine_version="4.2",
            name=name,
            shard_lists=[
                {
                    "node_class": "dds.shard.mid",
                    "node_storage": 10,
                },
                {
                    "node_class": "dds.shard.standard",
                    "node_storage": 20,
                    "readonly_replicas": 1,
                },
            ],
            mongo_lists=[
                {
                    "node_class": "dds.mongos.mid",
                },
                {
                    "node_class": "dds.mongos.mid",
                },
            ])
        example = alicloud.mongodb.ShardingNetworkPublicAddress("example",
            db_instance_id=default_sharding_instance.id,
            node_id=default_sharding_instance.mongo_lists[0].node_id)
        ```

        ## Import

        MongoDB Sharding Network Public Address can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:mongodb/shardingNetworkPublicAddress:ShardingNetworkPublicAddress example <db_instance_id>:<node_id>
        ```

        :param str resource_name: The name of the resource.
        :param ShardingNetworkPublicAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ShardingNetworkPublicAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 node_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ShardingNetworkPublicAddressArgs.__new__(ShardingNetworkPublicAddressArgs)

            if db_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_instance_id'")
            __props__.__dict__["db_instance_id"] = db_instance_id
            if node_id is None and not opts.urn:
                raise TypeError("Missing required property 'node_id'")
            __props__.__dict__["node_id"] = node_id
            __props__.__dict__["network_addresses"] = None
        super(ShardingNetworkPublicAddress, __self__).__init__(
            'alicloud:mongodb/shardingNetworkPublicAddress:ShardingNetworkPublicAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            db_instance_id: Optional[pulumi.Input[str]] = None,
            network_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ShardingNetworkPublicAddressNetworkAddressArgs', 'ShardingNetworkPublicAddressNetworkAddressArgsDict']]]]] = None,
            node_id: Optional[pulumi.Input[str]] = None) -> 'ShardingNetworkPublicAddress':
        """
        Get an existing ShardingNetworkPublicAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: The ID of the instance.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ShardingNetworkPublicAddressNetworkAddressArgs', 'ShardingNetworkPublicAddressNetworkAddressArgsDict']]]] network_addresses: The endpoint of the instance.
        :param pulumi.Input[str] node_id: The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ShardingNetworkPublicAddressState.__new__(_ShardingNetworkPublicAddressState)

        __props__.__dict__["db_instance_id"] = db_instance_id
        __props__.__dict__["network_addresses"] = network_addresses
        __props__.__dict__["node_id"] = node_id
        return ShardingNetworkPublicAddress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "db_instance_id")

    @property
    @pulumi.getter(name="networkAddresses")
    def network_addresses(self) -> pulumi.Output[Sequence['outputs.ShardingNetworkPublicAddressNetworkAddress']]:
        """
        The endpoint of the instance.
        """
        return pulumi.get(self, "network_addresses")

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> pulumi.Output[str]:
        """
        The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
        """
        return pulumi.get(self, "node_id")

