# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RemoteWriteArgs', 'RemoteWrite']

@pulumi.input_type
class RemoteWriteArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 remote_write_yaml: pulumi.Input[str]):
        """
        The set of arguments for constructing a RemoteWrite resource.
        :param pulumi.Input[str] cluster_id: The ID of the Prometheus instance.
        :param pulumi.Input[str] remote_write_yaml: The details of the Remote Write configuration item. Specify the value in the YAML format.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "remote_write_yaml", remote_write_yaml)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The ID of the Prometheus instance.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="remoteWriteYaml")
    def remote_write_yaml(self) -> pulumi.Input[str]:
        """
        The details of the Remote Write configuration item. Specify the value in the YAML format.
        """
        return pulumi.get(self, "remote_write_yaml")

    @remote_write_yaml.setter
    def remote_write_yaml(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_write_yaml", value)


@pulumi.input_type
class _RemoteWriteState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 remote_write_name: Optional[pulumi.Input[str]] = None,
                 remote_write_yaml: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RemoteWrite resources.
        :param pulumi.Input[str] cluster_id: The ID of the Prometheus instance.
        :param pulumi.Input[str] remote_write_name: The name of the Remote Write configuration item.
        :param pulumi.Input[str] remote_write_yaml: The details of the Remote Write configuration item. Specify the value in the YAML format.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if remote_write_name is not None:
            pulumi.set(__self__, "remote_write_name", remote_write_name)
        if remote_write_yaml is not None:
            pulumi.set(__self__, "remote_write_yaml", remote_write_yaml)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Prometheus instance.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="remoteWriteName")
    def remote_write_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Remote Write configuration item.
        """
        return pulumi.get(self, "remote_write_name")

    @remote_write_name.setter
    def remote_write_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_write_name", value)

    @property
    @pulumi.getter(name="remoteWriteYaml")
    def remote_write_yaml(self) -> Optional[pulumi.Input[str]]:
        """
        The details of the Remote Write configuration item. Specify the value in the YAML format.
        """
        return pulumi.get(self, "remote_write_yaml")

    @remote_write_yaml.setter
    def remote_write_yaml(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_write_yaml", value)


class RemoteWrite(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 remote_write_yaml: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Application Real-Time Monitoring Service (ARMS) Remote Write resource.

        For information about Application Real-Time Monitoring Service (ARMS) Remote Write and how to use it, see [What is Remote Write](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-addprometheusremotewrite).

        > **NOTE:** Available since v1.204.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=default_network.id,
            zone_id=default.zones[len(default.zones) - 1].id)
        default_security_group = alicloud.ecs.SecurityGroup("default",
            name=name,
            vpc_id=default_network.id)
        default_get_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_prometheus = alicloud.arms.Prometheus("default",
            cluster_type="ecs",
            grafana_instance_id="free",
            vpc_id=default_network.id,
            vswitch_id=default_switch.id,
            security_group_id=default_security_group.id,
            cluster_name=default_network.id.apply(lambda id: f"{name}-{id}"),
            resource_group_id=default_get_resource_groups.groups[0].id,
            tags={
                "Created": "TF",
                "For": "Prometheus",
            })
        default_remote_write = alicloud.arms.RemoteWrite("default",
            cluster_id=default_prometheus.id,
            remote_write_yaml=\"\"\"remote_write:
        - name: ArmsRemoteWrite
          url: http://47.96.227.137:8080/prometheus/xxx/yyy/cn-hangzhou/api/v3/write
          basic_auth: {username: 666, password: '******'}
          write_relabel_configs:
          - source_labels: [instance_id]
            separator: ;
            regex: si-6e2ca86444db4e55a7c1
            replacement: $1
            action: keep
        \"\"\")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Application Real-Time Monitoring Service (ARMS) Remote Write can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:arms/remoteWrite:RemoteWrite example <cluster_id>:<remote_write_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The ID of the Prometheus instance.
        :param pulumi.Input[str] remote_write_yaml: The details of the Remote Write configuration item. Specify the value in the YAML format.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RemoteWriteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Application Real-Time Monitoring Service (ARMS) Remote Write resource.

        For information about Application Real-Time Monitoring Service (ARMS) Remote Write and how to use it, see [What is Remote Write](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-addprometheusremotewrite).

        > **NOTE:** Available since v1.204.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=default_network.id,
            zone_id=default.zones[len(default.zones) - 1].id)
        default_security_group = alicloud.ecs.SecurityGroup("default",
            name=name,
            vpc_id=default_network.id)
        default_get_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_prometheus = alicloud.arms.Prometheus("default",
            cluster_type="ecs",
            grafana_instance_id="free",
            vpc_id=default_network.id,
            vswitch_id=default_switch.id,
            security_group_id=default_security_group.id,
            cluster_name=default_network.id.apply(lambda id: f"{name}-{id}"),
            resource_group_id=default_get_resource_groups.groups[0].id,
            tags={
                "Created": "TF",
                "For": "Prometheus",
            })
        default_remote_write = alicloud.arms.RemoteWrite("default",
            cluster_id=default_prometheus.id,
            remote_write_yaml=\"\"\"remote_write:
        - name: ArmsRemoteWrite
          url: http://47.96.227.137:8080/prometheus/xxx/yyy/cn-hangzhou/api/v3/write
          basic_auth: {username: 666, password: '******'}
          write_relabel_configs:
          - source_labels: [instance_id]
            separator: ;
            regex: si-6e2ca86444db4e55a7c1
            replacement: $1
            action: keep
        \"\"\")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Application Real-Time Monitoring Service (ARMS) Remote Write can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:arms/remoteWrite:RemoteWrite example <cluster_id>:<remote_write_name>
        ```

        :param str resource_name: The name of the resource.
        :param RemoteWriteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RemoteWriteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 remote_write_yaml: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RemoteWriteArgs.__new__(RemoteWriteArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if remote_write_yaml is None and not opts.urn:
                raise TypeError("Missing required property 'remote_write_yaml'")
            __props__.__dict__["remote_write_yaml"] = remote_write_yaml
            __props__.__dict__["remote_write_name"] = None
        super(RemoteWrite, __self__).__init__(
            'alicloud:arms/remoteWrite:RemoteWrite',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            remote_write_name: Optional[pulumi.Input[str]] = None,
            remote_write_yaml: Optional[pulumi.Input[str]] = None) -> 'RemoteWrite':
        """
        Get an existing RemoteWrite resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The ID of the Prometheus instance.
        :param pulumi.Input[str] remote_write_name: The name of the Remote Write configuration item.
        :param pulumi.Input[str] remote_write_yaml: The details of the Remote Write configuration item. Specify the value in the YAML format.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RemoteWriteState.__new__(_RemoteWriteState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["remote_write_name"] = remote_write_name
        __props__.__dict__["remote_write_yaml"] = remote_write_yaml
        return RemoteWrite(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The ID of the Prometheus instance.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="remoteWriteName")
    def remote_write_name(self) -> pulumi.Output[str]:
        """
        The name of the Remote Write configuration item.
        """
        return pulumi.get(self, "remote_write_name")

    @property
    @pulumi.getter(name="remoteWriteYaml")
    def remote_write_yaml(self) -> pulumi.Output[str]:
        """
        The details of the Remote Write configuration item. Specify the value in the YAML format.
        """
        return pulumi.get(self, "remote_write_yaml")

