# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EcsDedicatedHostClusterArgs', 'EcsDedicatedHostCluster']

@pulumi.input_type
class EcsDedicatedHostClusterArgs:
    def __init__(__self__, *,
                 zone_id: pulumi.Input[str],
                 dedicated_host_cluster_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a EcsDedicatedHostCluster resource.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the dedicated host cluster.
        :param pulumi.Input[str] dedicated_host_cluster_name: The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
        :param pulumi.Input[str] description: The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "zone_id", zone_id)
        if dedicated_host_cluster_name is not None:
            pulumi.set(__self__, "dedicated_host_cluster_name", dedicated_host_cluster_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The ID of the zone in which to create the dedicated host cluster.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter(name="dedicatedHostClusterName")
    def dedicated_host_cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
        """
        return pulumi.get(self, "dedicated_host_cluster_name")

    @dedicated_host_cluster_name.setter
    def dedicated_host_cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dedicated_host_cluster_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EcsDedicatedHostClusterState:
    def __init__(__self__, *,
                 dedicated_host_cluster_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EcsDedicatedHostCluster resources.
        :param pulumi.Input[str] dedicated_host_cluster_name: The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
        :param pulumi.Input[str] description: The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the dedicated host cluster.
        """
        if dedicated_host_cluster_name is not None:
            pulumi.set(__self__, "dedicated_host_cluster_name", dedicated_host_cluster_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="dedicatedHostClusterName")
    def dedicated_host_cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
        """
        return pulumi.get(self, "dedicated_host_cluster_name")

    @dedicated_host_cluster_name.setter
    def dedicated_host_cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dedicated_host_cluster_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the zone in which to create the dedicated host cluster.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class EcsDedicatedHostCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_host_cluster_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ECS Dedicated Host Cluster resource.

        For information about ECS Dedicated Host Cluster and how to use it, see [What is Dedicated Host Cluster](https://www.alibabacloud.com/help/en/doc-detail/184667.html).

        > **NOTE:** Available since v1.146.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_zones = alicloud.get_zones()
        example_ecs_dedicated_host_cluster = alicloud.ecs.EcsDedicatedHostCluster("exampleEcsDedicatedHostCluster",
            dedicated_host_cluster_name="example_value",
            description="example_value",
            zone_id=example_zones.zones[0].id,
            tags={
                "Create": "TF",
                "For": "DDH_Cluster_Test",
            })
        ```

        ## Import

        ECS Dedicated Host Cluster can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsDedicatedHostCluster:EcsDedicatedHostCluster example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dedicated_host_cluster_name: The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
        :param pulumi.Input[str] description: The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the dedicated host cluster.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EcsDedicatedHostClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ECS Dedicated Host Cluster resource.

        For information about ECS Dedicated Host Cluster and how to use it, see [What is Dedicated Host Cluster](https://www.alibabacloud.com/help/en/doc-detail/184667.html).

        > **NOTE:** Available since v1.146.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_zones = alicloud.get_zones()
        example_ecs_dedicated_host_cluster = alicloud.ecs.EcsDedicatedHostCluster("exampleEcsDedicatedHostCluster",
            dedicated_host_cluster_name="example_value",
            description="example_value",
            zone_id=example_zones.zones[0].id,
            tags={
                "Create": "TF",
                "For": "DDH_Cluster_Test",
            })
        ```

        ## Import

        ECS Dedicated Host Cluster can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/ecsDedicatedHostCluster:EcsDedicatedHostCluster example <id>
        ```

        :param str resource_name: The name of the resource.
        :param EcsDedicatedHostClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EcsDedicatedHostClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_host_cluster_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EcsDedicatedHostClusterArgs.__new__(EcsDedicatedHostClusterArgs)

            __props__.__dict__["dedicated_host_cluster_name"] = dedicated_host_cluster_name
            __props__.__dict__["description"] = description
            __props__.__dict__["dry_run"] = dry_run
            __props__.__dict__["tags"] = tags
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(EcsDedicatedHostCluster, __self__).__init__(
            'alicloud:ecs/ecsDedicatedHostCluster:EcsDedicatedHostCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dedicated_host_cluster_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'EcsDedicatedHostCluster':
        """
        Get an existing EcsDedicatedHostCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dedicated_host_cluster_name: The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
        :param pulumi.Input[str] description: The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] dry_run: The dry run.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the dedicated host cluster.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EcsDedicatedHostClusterState.__new__(_EcsDedicatedHostClusterState)

        __props__.__dict__["dedicated_host_cluster_name"] = dedicated_host_cluster_name
        __props__.__dict__["description"] = description
        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["tags"] = tags
        __props__.__dict__["zone_id"] = zone_id
        return EcsDedicatedHostCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dedicatedHostClusterName")
    def dedicated_host_cluster_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the dedicated host cluster. The name must be `2` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot contain `http://` or `https://`.
        """
        return pulumi.get(self, "dedicated_host_cluster_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the dedicated host cluster. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        The dry run.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The ID of the zone in which to create the dedicated host cluster.
        """
        return pulumi.get(self, "zone_id")

