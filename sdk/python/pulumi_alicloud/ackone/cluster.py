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

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 network: pulumi.Input['ClusterNetworkArgs'],
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input['ClusterNetworkArgs'] network: Cluster network information. See `network` below.
        :param pulumi.Input[str] cluster_name: Cluster name.
        :param pulumi.Input[str] profile: Cluster attributes. Valid values: 'Default', 'XFlow'.
        """
        pulumi.set(__self__, "network", network)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)

    @property
    @pulumi.getter
    def network(self) -> pulumi.Input['ClusterNetworkArgs']:
        """
        Cluster network information. See `network` below.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: pulumi.Input['ClusterNetworkArgs']):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster attributes. Valid values: 'Default', 'XFlow'.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)


@pulumi.input_type
class _ClusterState:
    def __init__(__self__, *,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input['ClusterNetworkArgs']] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[str] cluster_name: Cluster name.
        :param pulumi.Input[str] create_time: Cluster creation time.
        :param pulumi.Input['ClusterNetworkArgs'] network: Cluster network information. See `network` below.
        :param pulumi.Input[str] profile: Cluster attributes. Valid values: 'Default', 'XFlow'.
        :param pulumi.Input[str] status: The status of the resource.
        """
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster creation time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input['ClusterNetworkArgs']]:
        """
        Cluster network information. See `network` below.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input['ClusterNetworkArgs']]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster attributes. Valid values: 'Default', 'XFlow'.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[pulumi.InputType['ClusterNetworkArgs']]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Ack One Cluster resource. Fleet Manager Cluster.

        For information about Ack One Cluster and how to use it, see [What is Cluster](https://www.alibabacloud.com/help/en/ack/distributed-cloud-container-platform-for-kubernetes/developer-reference/api-adcp-2022-01-01-createhubcluster).

        > **NOTE:** Available since v1.212.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_vpc = alicloud.vpc.Network("defaultVpc",
            cidr_block="172.16.0.0/12",
            vpc_name=name)
        defaulty_v_switch = alicloud.vpc.Switch("defaultyVSwitch",
            vpc_id=default_vpc.id,
            cidr_block="172.16.2.0/24",
            zone_id=default_zones.zones[0].id,
            vswitch_name=name)
        default_cluster = alicloud.ackone.Cluster("defaultCluster", network=alicloud.ackone.ClusterNetworkArgs(
            vpc_id=default_vpc.id,
            vswitches=[defaulty_v_switch.id],
        ))
        ```

        ## Import

        Ack One Cluster can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ackone/cluster:Cluster example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Cluster name.
        :param pulumi.Input[pulumi.InputType['ClusterNetworkArgs']] network: Cluster network information. See `network` below.
        :param pulumi.Input[str] profile: Cluster attributes. Valid values: 'Default', 'XFlow'.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Ack One Cluster resource. Fleet Manager Cluster.

        For information about Ack One Cluster and how to use it, see [What is Cluster](https://www.alibabacloud.com/help/en/ack/distributed-cloud-container-platform-for-kubernetes/developer-reference/api-adcp-2022-01-01-createhubcluster).

        > **NOTE:** Available since v1.212.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_vpc = alicloud.vpc.Network("defaultVpc",
            cidr_block="172.16.0.0/12",
            vpc_name=name)
        defaulty_v_switch = alicloud.vpc.Switch("defaultyVSwitch",
            vpc_id=default_vpc.id,
            cidr_block="172.16.2.0/24",
            zone_id=default_zones.zones[0].id,
            vswitch_name=name)
        default_cluster = alicloud.ackone.Cluster("defaultCluster", network=alicloud.ackone.ClusterNetworkArgs(
            vpc_id=default_vpc.id,
            vswitches=[defaulty_v_switch.id],
        ))
        ```

        ## Import

        Ack One Cluster can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ackone/cluster:Cluster example <id>
        ```

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[pulumi.InputType['ClusterNetworkArgs']]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            __props__.__dict__["cluster_name"] = cluster_name
            if network is None and not opts.urn:
                raise TypeError("Missing required property 'network'")
            __props__.__dict__["network"] = network
            __props__.__dict__["profile"] = profile
            __props__.__dict__["create_time"] = None
            __props__.__dict__["status"] = None
        super(Cluster, __self__).__init__(
            'alicloud:ackone/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[pulumi.InputType['ClusterNetworkArgs']]] = None,
            profile: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Cluster name.
        :param pulumi.Input[str] create_time: Cluster creation time.
        :param pulumi.Input[pulumi.InputType['ClusterNetworkArgs']] network: Cluster network information. See `network` below.
        :param pulumi.Input[str] profile: Cluster attributes. Valid values: 'Default', 'XFlow'.
        :param pulumi.Input[str] status: The status of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["network"] = network
        __props__.__dict__["profile"] = profile
        __props__.__dict__["status"] = status
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        Cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Cluster creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output['outputs.ClusterNetwork']:
        """
        Cluster network information. See `network` below.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def profile(self) -> pulumi.Output[str]:
        """
        Cluster attributes. Valid values: 'Default', 'XFlow'.
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

