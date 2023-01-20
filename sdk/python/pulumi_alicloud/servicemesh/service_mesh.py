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

__all__ = ['ServiceMeshArgs', 'ServiceMesh']

@pulumi.input_type
class ServiceMeshArgs:
    def __init__(__self__, *,
                 network: pulumi.Input['ServiceMeshNetworkArgs'],
                 cluster_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_spec: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 extra_configuration: Optional[pulumi.Input['ServiceMeshExtraConfigurationArgs']] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 load_balancer: Optional[pulumi.Input['ServiceMeshLoadBalancerArgs']] = None,
                 mesh_config: Optional[pulumi.Input['ServiceMeshMeshConfigArgs']] = None,
                 service_mesh_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceMesh resource.
        :param pulumi.Input['ServiceMeshNetworkArgs'] network: The network configuration of the Service grid. See the following `Block network`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cluster_ids: The array of the cluster ids.
        :param pulumi.Input[str] cluster_spec: The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`.
        :param pulumi.Input[str] edition: The type  of the resource. Valid values: `Default` and `Pro`. `Default`:the standard. `Pro`:the Pro version.
        :param pulumi.Input['ServiceMeshExtraConfigurationArgs'] extra_configuration: The configurations of additional features for the ASM instance. See the following `Block extra_configuration`.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input['ServiceMeshLoadBalancerArgs'] load_balancer: The configuration of the Load Balancer. See the following `Block load_balancer`.
        :param pulumi.Input['ServiceMeshMeshConfigArgs'] mesh_config: The configuration of the Service grid. See the following `Block mesh_config`.
        :param pulumi.Input[str] service_mesh_name: The name of the resource.
        :param pulumi.Input[str] version: The version of the resource. you can look up the version using `servicemesh.get_versions`. **Note:** The `version` supports updating from v1.170.0, the relevant version can be obtained via `istio_operator_version` in `servicemesh.get_service_meshes`.
        """
        pulumi.set(__self__, "network", network)
        if cluster_ids is not None:
            pulumi.set(__self__, "cluster_ids", cluster_ids)
        if cluster_spec is not None:
            pulumi.set(__self__, "cluster_spec", cluster_spec)
        if edition is not None:
            pulumi.set(__self__, "edition", edition)
        if extra_configuration is not None:
            pulumi.set(__self__, "extra_configuration", extra_configuration)
        if force is not None:
            pulumi.set(__self__, "force", force)
        if load_balancer is not None:
            pulumi.set(__self__, "load_balancer", load_balancer)
        if mesh_config is not None:
            pulumi.set(__self__, "mesh_config", mesh_config)
        if service_mesh_name is not None:
            pulumi.set(__self__, "service_mesh_name", service_mesh_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def network(self) -> pulumi.Input['ServiceMeshNetworkArgs']:
        """
        The network configuration of the Service grid. See the following `Block network`.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: pulumi.Input['ServiceMeshNetworkArgs']):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="clusterIds")
    def cluster_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The array of the cluster ids.
        """
        return pulumi.get(self, "cluster_ids")

    @cluster_ids.setter
    def cluster_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cluster_ids", value)

    @property
    @pulumi.getter(name="clusterSpec")
    def cluster_spec(self) -> Optional[pulumi.Input[str]]:
        """
        The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`.
        """
        return pulumi.get(self, "cluster_spec")

    @cluster_spec.setter
    def cluster_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_spec", value)

    @property
    @pulumi.getter
    def edition(self) -> Optional[pulumi.Input[str]]:
        """
        The type  of the resource. Valid values: `Default` and `Pro`. `Default`:the standard. `Pro`:the Pro version.
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter(name="extraConfiguration")
    def extra_configuration(self) -> Optional[pulumi.Input['ServiceMeshExtraConfigurationArgs']]:
        """
        The configurations of additional features for the ASM instance. See the following `Block extra_configuration`.
        """
        return pulumi.get(self, "extra_configuration")

    @extra_configuration.setter
    def extra_configuration(self, value: Optional[pulumi.Input['ServiceMeshExtraConfigurationArgs']]):
        pulumi.set(self, "extra_configuration", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[bool]]:
        """
        This parameter is used for resource destroy. Default value is `false`.
        """
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> Optional[pulumi.Input['ServiceMeshLoadBalancerArgs']]:
        """
        The configuration of the Load Balancer. See the following `Block load_balancer`.
        """
        return pulumi.get(self, "load_balancer")

    @load_balancer.setter
    def load_balancer(self, value: Optional[pulumi.Input['ServiceMeshLoadBalancerArgs']]):
        pulumi.set(self, "load_balancer", value)

    @property
    @pulumi.getter(name="meshConfig")
    def mesh_config(self) -> Optional[pulumi.Input['ServiceMeshMeshConfigArgs']]:
        """
        The configuration of the Service grid. See the following `Block mesh_config`.
        """
        return pulumi.get(self, "mesh_config")

    @mesh_config.setter
    def mesh_config(self, value: Optional[pulumi.Input['ServiceMeshMeshConfigArgs']]):
        pulumi.set(self, "mesh_config", value)

    @property
    @pulumi.getter(name="serviceMeshName")
    def service_mesh_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "service_mesh_name")

    @service_mesh_name.setter
    def service_mesh_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_mesh_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the resource. you can look up the version using `servicemesh.get_versions`. **Note:** The `version` supports updating from v1.170.0, the relevant version can be obtained via `istio_operator_version` in `servicemesh.get_service_meshes`.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _ServiceMeshState:
    def __init__(__self__, *,
                 cluster_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_spec: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 extra_configuration: Optional[pulumi.Input['ServiceMeshExtraConfigurationArgs']] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 load_balancer: Optional[pulumi.Input['ServiceMeshLoadBalancerArgs']] = None,
                 mesh_config: Optional[pulumi.Input['ServiceMeshMeshConfigArgs']] = None,
                 network: Optional[pulumi.Input['ServiceMeshNetworkArgs']] = None,
                 service_mesh_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceMesh resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cluster_ids: The array of the cluster ids.
        :param pulumi.Input[str] cluster_spec: The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`.
        :param pulumi.Input[str] edition: The type  of the resource. Valid values: `Default` and `Pro`. `Default`:the standard. `Pro`:the Pro version.
        :param pulumi.Input['ServiceMeshExtraConfigurationArgs'] extra_configuration: The configurations of additional features for the ASM instance. See the following `Block extra_configuration`.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input['ServiceMeshLoadBalancerArgs'] load_balancer: The configuration of the Load Balancer. See the following `Block load_balancer`.
        :param pulumi.Input['ServiceMeshMeshConfigArgs'] mesh_config: The configuration of the Service grid. See the following `Block mesh_config`.
        :param pulumi.Input['ServiceMeshNetworkArgs'] network: The network configuration of the Service grid. See the following `Block network`.
        :param pulumi.Input[str] service_mesh_name: The name of the resource.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `running` or `initial`.
        :param pulumi.Input[str] version: The version of the resource. you can look up the version using `servicemesh.get_versions`. **Note:** The `version` supports updating from v1.170.0, the relevant version can be obtained via `istio_operator_version` in `servicemesh.get_service_meshes`.
        """
        if cluster_ids is not None:
            pulumi.set(__self__, "cluster_ids", cluster_ids)
        if cluster_spec is not None:
            pulumi.set(__self__, "cluster_spec", cluster_spec)
        if edition is not None:
            pulumi.set(__self__, "edition", edition)
        if extra_configuration is not None:
            pulumi.set(__self__, "extra_configuration", extra_configuration)
        if force is not None:
            pulumi.set(__self__, "force", force)
        if load_balancer is not None:
            pulumi.set(__self__, "load_balancer", load_balancer)
        if mesh_config is not None:
            pulumi.set(__self__, "mesh_config", mesh_config)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if service_mesh_name is not None:
            pulumi.set(__self__, "service_mesh_name", service_mesh_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="clusterIds")
    def cluster_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The array of the cluster ids.
        """
        return pulumi.get(self, "cluster_ids")

    @cluster_ids.setter
    def cluster_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cluster_ids", value)

    @property
    @pulumi.getter(name="clusterSpec")
    def cluster_spec(self) -> Optional[pulumi.Input[str]]:
        """
        The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`.
        """
        return pulumi.get(self, "cluster_spec")

    @cluster_spec.setter
    def cluster_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_spec", value)

    @property
    @pulumi.getter
    def edition(self) -> Optional[pulumi.Input[str]]:
        """
        The type  of the resource. Valid values: `Default` and `Pro`. `Default`:the standard. `Pro`:the Pro version.
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter(name="extraConfiguration")
    def extra_configuration(self) -> Optional[pulumi.Input['ServiceMeshExtraConfigurationArgs']]:
        """
        The configurations of additional features for the ASM instance. See the following `Block extra_configuration`.
        """
        return pulumi.get(self, "extra_configuration")

    @extra_configuration.setter
    def extra_configuration(self, value: Optional[pulumi.Input['ServiceMeshExtraConfigurationArgs']]):
        pulumi.set(self, "extra_configuration", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[bool]]:
        """
        This parameter is used for resource destroy. Default value is `false`.
        """
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> Optional[pulumi.Input['ServiceMeshLoadBalancerArgs']]:
        """
        The configuration of the Load Balancer. See the following `Block load_balancer`.
        """
        return pulumi.get(self, "load_balancer")

    @load_balancer.setter
    def load_balancer(self, value: Optional[pulumi.Input['ServiceMeshLoadBalancerArgs']]):
        pulumi.set(self, "load_balancer", value)

    @property
    @pulumi.getter(name="meshConfig")
    def mesh_config(self) -> Optional[pulumi.Input['ServiceMeshMeshConfigArgs']]:
        """
        The configuration of the Service grid. See the following `Block mesh_config`.
        """
        return pulumi.get(self, "mesh_config")

    @mesh_config.setter
    def mesh_config(self, value: Optional[pulumi.Input['ServiceMeshMeshConfigArgs']]):
        pulumi.set(self, "mesh_config", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input['ServiceMeshNetworkArgs']]:
        """
        The network configuration of the Service grid. See the following `Block network`.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input['ServiceMeshNetworkArgs']]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="serviceMeshName")
    def service_mesh_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "service_mesh_name")

    @service_mesh_name.setter
    def service_mesh_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_mesh_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource. Valid values: `running` or `initial`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the resource. you can look up the version using `servicemesh.get_versions`. **Note:** The `version` supports updating from v1.170.0, the relevant version can be obtained via `istio_operator_version` in `servicemesh.get_service_meshes`.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class ServiceMesh(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_spec: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 extra_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceMeshExtraConfigurationArgs']]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 load_balancer: Optional[pulumi.Input[pulumi.InputType['ServiceMeshLoadBalancerArgs']]] = None,
                 mesh_config: Optional[pulumi.Input[pulumi.InputType['ServiceMeshMeshConfigArgs']]] = None,
                 network: Optional[pulumi.Input[pulumi.InputType['ServiceMeshNetworkArgs']]] = None,
                 service_mesh_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Service Mesh Service Mesh resource.

        For information about Service Mesh Service Mesh and how to use it, see [What is Service Mesh](https://help.aliyun.com/document_detail/171559.html).

        > **NOTE:** Available in v1.138.0+.

        ## Import

        Service Mesh Service Mesh can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:servicemesh/serviceMesh:ServiceMesh example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cluster_ids: The array of the cluster ids.
        :param pulumi.Input[str] cluster_spec: The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`.
        :param pulumi.Input[str] edition: The type  of the resource. Valid values: `Default` and `Pro`. `Default`:the standard. `Pro`:the Pro version.
        :param pulumi.Input[pulumi.InputType['ServiceMeshExtraConfigurationArgs']] extra_configuration: The configurations of additional features for the ASM instance. See the following `Block extra_configuration`.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input[pulumi.InputType['ServiceMeshLoadBalancerArgs']] load_balancer: The configuration of the Load Balancer. See the following `Block load_balancer`.
        :param pulumi.Input[pulumi.InputType['ServiceMeshMeshConfigArgs']] mesh_config: The configuration of the Service grid. See the following `Block mesh_config`.
        :param pulumi.Input[pulumi.InputType['ServiceMeshNetworkArgs']] network: The network configuration of the Service grid. See the following `Block network`.
        :param pulumi.Input[str] service_mesh_name: The name of the resource.
        :param pulumi.Input[str] version: The version of the resource. you can look up the version using `servicemesh.get_versions`. **Note:** The `version` supports updating from v1.170.0, the relevant version can be obtained via `istio_operator_version` in `servicemesh.get_service_meshes`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceMeshArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Service Mesh Service Mesh resource.

        For information about Service Mesh Service Mesh and how to use it, see [What is Service Mesh](https://help.aliyun.com/document_detail/171559.html).

        > **NOTE:** Available in v1.138.0+.

        ## Import

        Service Mesh Service Mesh can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:servicemesh/serviceMesh:ServiceMesh example <id>
        ```

        :param str resource_name: The name of the resource.
        :param ServiceMeshArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceMeshArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_spec: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 extra_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceMeshExtraConfigurationArgs']]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 load_balancer: Optional[pulumi.Input[pulumi.InputType['ServiceMeshLoadBalancerArgs']]] = None,
                 mesh_config: Optional[pulumi.Input[pulumi.InputType['ServiceMeshMeshConfigArgs']]] = None,
                 network: Optional[pulumi.Input[pulumi.InputType['ServiceMeshNetworkArgs']]] = None,
                 service_mesh_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceMeshArgs.__new__(ServiceMeshArgs)

            __props__.__dict__["cluster_ids"] = cluster_ids
            __props__.__dict__["cluster_spec"] = cluster_spec
            __props__.__dict__["edition"] = edition
            __props__.__dict__["extra_configuration"] = extra_configuration
            __props__.__dict__["force"] = force
            __props__.__dict__["load_balancer"] = load_balancer
            __props__.__dict__["mesh_config"] = mesh_config
            if network is None and not opts.urn:
                raise TypeError("Missing required property 'network'")
            __props__.__dict__["network"] = network
            __props__.__dict__["service_mesh_name"] = service_mesh_name
            __props__.__dict__["version"] = version
            __props__.__dict__["status"] = None
        super(ServiceMesh, __self__).__init__(
            'alicloud:servicemesh/serviceMesh:ServiceMesh',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            cluster_spec: Optional[pulumi.Input[str]] = None,
            edition: Optional[pulumi.Input[str]] = None,
            extra_configuration: Optional[pulumi.Input[pulumi.InputType['ServiceMeshExtraConfigurationArgs']]] = None,
            force: Optional[pulumi.Input[bool]] = None,
            load_balancer: Optional[pulumi.Input[pulumi.InputType['ServiceMeshLoadBalancerArgs']]] = None,
            mesh_config: Optional[pulumi.Input[pulumi.InputType['ServiceMeshMeshConfigArgs']]] = None,
            network: Optional[pulumi.Input[pulumi.InputType['ServiceMeshNetworkArgs']]] = None,
            service_mesh_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'ServiceMesh':
        """
        Get an existing ServiceMesh resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cluster_ids: The array of the cluster ids.
        :param pulumi.Input[str] cluster_spec: The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`.
        :param pulumi.Input[str] edition: The type  of the resource. Valid values: `Default` and `Pro`. `Default`:the standard. `Pro`:the Pro version.
        :param pulumi.Input[pulumi.InputType['ServiceMeshExtraConfigurationArgs']] extra_configuration: The configurations of additional features for the ASM instance. See the following `Block extra_configuration`.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input[pulumi.InputType['ServiceMeshLoadBalancerArgs']] load_balancer: The configuration of the Load Balancer. See the following `Block load_balancer`.
        :param pulumi.Input[pulumi.InputType['ServiceMeshMeshConfigArgs']] mesh_config: The configuration of the Service grid. See the following `Block mesh_config`.
        :param pulumi.Input[pulumi.InputType['ServiceMeshNetworkArgs']] network: The network configuration of the Service grid. See the following `Block network`.
        :param pulumi.Input[str] service_mesh_name: The name of the resource.
        :param pulumi.Input[str] status: The status of the resource. Valid values: `running` or `initial`.
        :param pulumi.Input[str] version: The version of the resource. you can look up the version using `servicemesh.get_versions`. **Note:** The `version` supports updating from v1.170.0, the relevant version can be obtained via `istio_operator_version` in `servicemesh.get_service_meshes`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceMeshState.__new__(_ServiceMeshState)

        __props__.__dict__["cluster_ids"] = cluster_ids
        __props__.__dict__["cluster_spec"] = cluster_spec
        __props__.__dict__["edition"] = edition
        __props__.__dict__["extra_configuration"] = extra_configuration
        __props__.__dict__["force"] = force
        __props__.__dict__["load_balancer"] = load_balancer
        __props__.__dict__["mesh_config"] = mesh_config
        __props__.__dict__["network"] = network
        __props__.__dict__["service_mesh_name"] = service_mesh_name
        __props__.__dict__["status"] = status
        __props__.__dict__["version"] = version
        return ServiceMesh(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterIds")
    def cluster_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The array of the cluster ids.
        """
        return pulumi.get(self, "cluster_ids")

    @property
    @pulumi.getter(name="clusterSpec")
    def cluster_spec(self) -> pulumi.Output[str]:
        """
        The service mesh instance specification. Valid values: `standard`,`enterprise`,`ultimate`.
        """
        return pulumi.get(self, "cluster_spec")

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Output[str]:
        """
        The type  of the resource. Valid values: `Default` and `Pro`. `Default`:the standard. `Pro`:the Pro version.
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="extraConfiguration")
    def extra_configuration(self) -> pulumi.Output['outputs.ServiceMeshExtraConfiguration']:
        """
        The configurations of additional features for the ASM instance. See the following `Block extra_configuration`.
        """
        return pulumi.get(self, "extra_configuration")

    @property
    @pulumi.getter
    def force(self) -> pulumi.Output[Optional[bool]]:
        """
        This parameter is used for resource destroy. Default value is `false`.
        """
        return pulumi.get(self, "force")

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> pulumi.Output['outputs.ServiceMeshLoadBalancer']:
        """
        The configuration of the Load Balancer. See the following `Block load_balancer`.
        """
        return pulumi.get(self, "load_balancer")

    @property
    @pulumi.getter(name="meshConfig")
    def mesh_config(self) -> pulumi.Output['outputs.ServiceMeshMeshConfig']:
        """
        The configuration of the Service grid. See the following `Block mesh_config`.
        """
        return pulumi.get(self, "mesh_config")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output['outputs.ServiceMeshNetwork']:
        """
        The network configuration of the Service grid. See the following `Block network`.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="serviceMeshName")
    def service_mesh_name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "service_mesh_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource. Valid values: `running` or `initial`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The version of the resource. you can look up the version using `servicemesh.get_versions`. **Note:** The `version` supports updating from v1.170.0, the relevant version can be obtained via `istio_operator_version` in `servicemesh.get_service_meshes`.
        """
        return pulumi.get(self, "version")

