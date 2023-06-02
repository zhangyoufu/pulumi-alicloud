# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EngineNamespaceArgs', 'EngineNamespace']

@pulumi.input_type
class EngineNamespaceArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 namespace_id: pulumi.Input[str],
                 namespace_show_name: pulumi.Input[str],
                 accept_language: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EngineNamespace resource.
        :param pulumi.Input[str] cluster_id: The id of the cluster.
        :param pulumi.Input[str] namespace_id: The id of Namespace.
        :param pulumi.Input[str] namespace_show_name: The name of the Engine Namespace.
        :param pulumi.Input[str] accept_language: The language type of the returned information. Valid values: `zh`, `en`.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "namespace_id", namespace_id)
        pulumi.set(__self__, "namespace_show_name", namespace_show_name)
        if accept_language is not None:
            pulumi.set(__self__, "accept_language", accept_language)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The id of the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Input[str]:
        """
        The id of Namespace.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="namespaceShowName")
    def namespace_show_name(self) -> pulumi.Input[str]:
        """
        The name of the Engine Namespace.
        """
        return pulumi.get(self, "namespace_show_name")

    @namespace_show_name.setter
    def namespace_show_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_show_name", value)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> Optional[pulumi.Input[str]]:
        """
        The language type of the returned information. Valid values: `zh`, `en`.
        """
        return pulumi.get(self, "accept_language")

    @accept_language.setter
    def accept_language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accept_language", value)


@pulumi.input_type
class _EngineNamespaceState:
    def __init__(__self__, *,
                 accept_language: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 namespace_show_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EngineNamespace resources.
        :param pulumi.Input[str] accept_language: The language type of the returned information. Valid values: `zh`, `en`.
        :param pulumi.Input[str] cluster_id: The id of the cluster.
        :param pulumi.Input[str] namespace_id: The id of Namespace.
        :param pulumi.Input[str] namespace_show_name: The name of the Engine Namespace.
        """
        if accept_language is not None:
            pulumi.set(__self__, "accept_language", accept_language)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if namespace_show_name is not None:
            pulumi.set(__self__, "namespace_show_name", namespace_show_name)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> Optional[pulumi.Input[str]]:
        """
        The language type of the returned information. Valid values: `zh`, `en`.
        """
        return pulumi.get(self, "accept_language")

    @accept_language.setter
    def accept_language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accept_language", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of Namespace.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="namespaceShowName")
    def namespace_show_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Engine Namespace.
        """
        return pulumi.get(self, "namespace_show_name")

    @namespace_show_name.setter
    def namespace_show_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_show_name", value)


class EngineNamespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_language: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 namespace_show_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Microservice Engine (MSE) Engine Namespace resource.

        For information about Microservice Engine (MSE) Engine Namespace and how to use it, see [What is Engine Namespace](https://www.alibabacloud.com/help/zh/microservices-engine/latest/api-doc-mse-2019-05-31-api-doc-createenginenamespace).

        > **NOTE:** Available in v1.166.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        example_switch = alicloud.vpc.Switch("exampleSwitch",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=example_network.id,
            zone_id=example_zones.zones[0].id)
        example_cluster = alicloud.mse.Cluster("exampleCluster",
            cluster_specification="MSE_SC_1_2_60_c",
            cluster_type="Nacos-Ans",
            cluster_version="NACOS_2_0_0",
            instance_count=1,
            net_type="privatenet",
            pub_network_flow="1",
            connection_type="slb",
            cluster_alias_name="terraform-example",
            mse_version="mse_dev",
            vswitch_id=example_switch.id,
            vpc_id=example_network.id)
        example_engine_namespace = alicloud.mse.EngineNamespace("exampleEngineNamespace",
            cluster_id=example_cluster.cluster_id,
            namespace_show_name="terraform-example",
            namespace_id="terraform-example")
        ```

        ## Import

        Microservice Engine (MSE) Engine Namespace can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:mse/engineNamespace:EngineNamespace example <cluster_id>:<namespace_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accept_language: The language type of the returned information. Valid values: `zh`, `en`.
        :param pulumi.Input[str] cluster_id: The id of the cluster.
        :param pulumi.Input[str] namespace_id: The id of Namespace.
        :param pulumi.Input[str] namespace_show_name: The name of the Engine Namespace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EngineNamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Microservice Engine (MSE) Engine Namespace resource.

        For information about Microservice Engine (MSE) Engine Namespace and how to use it, see [What is Engine Namespace](https://www.alibabacloud.com/help/zh/microservices-engine/latest/api-doc-mse-2019-05-31-api-doc-createenginenamespace).

        > **NOTE:** Available in v1.166.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        example_switch = alicloud.vpc.Switch("exampleSwitch",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=example_network.id,
            zone_id=example_zones.zones[0].id)
        example_cluster = alicloud.mse.Cluster("exampleCluster",
            cluster_specification="MSE_SC_1_2_60_c",
            cluster_type="Nacos-Ans",
            cluster_version="NACOS_2_0_0",
            instance_count=1,
            net_type="privatenet",
            pub_network_flow="1",
            connection_type="slb",
            cluster_alias_name="terraform-example",
            mse_version="mse_dev",
            vswitch_id=example_switch.id,
            vpc_id=example_network.id)
        example_engine_namespace = alicloud.mse.EngineNamespace("exampleEngineNamespace",
            cluster_id=example_cluster.cluster_id,
            namespace_show_name="terraform-example",
            namespace_id="terraform-example")
        ```

        ## Import

        Microservice Engine (MSE) Engine Namespace can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:mse/engineNamespace:EngineNamespace example <cluster_id>:<namespace_id>
        ```

        :param str resource_name: The name of the resource.
        :param EngineNamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EngineNamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_language: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 namespace_show_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EngineNamespaceArgs.__new__(EngineNamespaceArgs)

            __props__.__dict__["accept_language"] = accept_language
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if namespace_id is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_id'")
            __props__.__dict__["namespace_id"] = namespace_id
            if namespace_show_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_show_name'")
            __props__.__dict__["namespace_show_name"] = namespace_show_name
        super(EngineNamespace, __self__).__init__(
            'alicloud:mse/engineNamespace:EngineNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accept_language: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            namespace_show_name: Optional[pulumi.Input[str]] = None) -> 'EngineNamespace':
        """
        Get an existing EngineNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accept_language: The language type of the returned information. Valid values: `zh`, `en`.
        :param pulumi.Input[str] cluster_id: The id of the cluster.
        :param pulumi.Input[str] namespace_id: The id of Namespace.
        :param pulumi.Input[str] namespace_show_name: The name of the Engine Namespace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EngineNamespaceState.__new__(_EngineNamespaceState)

        __props__.__dict__["accept_language"] = accept_language
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["namespace_show_name"] = namespace_show_name
        return EngineNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> pulumi.Output[Optional[str]]:
        """
        The language type of the returned information. Valid values: `zh`, `en`.
        """
        return pulumi.get(self, "accept_language")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The id of the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[str]:
        """
        The id of Namespace.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="namespaceShowName")
    def namespace_show_name(self) -> pulumi.Output[str]:
        """
        The name of the Engine Namespace.
        """
        return pulumi.get(self, "namespace_show_name")

