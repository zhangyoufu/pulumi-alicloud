# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TransitRouterMulticastDomainArgs', 'TransitRouterMulticastDomain']

@pulumi.input_type
class TransitRouterMulticastDomainArgs:
    def __init__(__self__, *,
                 transit_router_id: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 transit_router_multicast_domain_description: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TransitRouterMulticastDomain resource.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] transit_router_multicast_domain_description: The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        :param pulumi.Input[str] transit_router_multicast_domain_name: The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        pulumi.set(__self__, "transit_router_id", transit_router_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if transit_router_multicast_domain_description is not None:
            pulumi.set(__self__, "transit_router_multicast_domain_description", transit_router_multicast_domain_description)
        if transit_router_multicast_domain_name is not None:
            pulumi.set(__self__, "transit_router_multicast_domain_name", transit_router_multicast_domain_name)

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> pulumi.Input[str]:
        """
        The ID of the transit router.
        """
        return pulumi.get(self, "transit_router_id")

    @transit_router_id.setter
    def transit_router_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_router_id", value)

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
    @pulumi.getter(name="transitRouterMulticastDomainDescription")
    def transit_router_multicast_domain_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "transit_router_multicast_domain_description")

    @transit_router_multicast_domain_description.setter
    def transit_router_multicast_domain_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_multicast_domain_description", value)

    @property
    @pulumi.getter(name="transitRouterMulticastDomainName")
    def transit_router_multicast_domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "transit_router_multicast_domain_name")

    @transit_router_multicast_domain_name.setter
    def transit_router_multicast_domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_multicast_domain_name", value)


@pulumi.input_type
class _TransitRouterMulticastDomainState:
    def __init__(__self__, *,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_description: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TransitRouterMulticastDomain resources.
        :param pulumi.Input[str] status: The status of the Transit Router Multicast Domain.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        :param pulumi.Input[str] transit_router_multicast_domain_description: The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        :param pulumi.Input[str] transit_router_multicast_domain_name: The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if transit_router_id is not None:
            pulumi.set(__self__, "transit_router_id", transit_router_id)
        if transit_router_multicast_domain_description is not None:
            pulumi.set(__self__, "transit_router_multicast_domain_description", transit_router_multicast_domain_description)
        if transit_router_multicast_domain_name is not None:
            pulumi.set(__self__, "transit_router_multicast_domain_name", transit_router_multicast_domain_name)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the Transit Router Multicast Domain.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the transit router.
        """
        return pulumi.get(self, "transit_router_id")

    @transit_router_id.setter
    def transit_router_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_id", value)

    @property
    @pulumi.getter(name="transitRouterMulticastDomainDescription")
    def transit_router_multicast_domain_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "transit_router_multicast_domain_description")

    @transit_router_multicast_domain_description.setter
    def transit_router_multicast_domain_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_multicast_domain_description", value)

    @property
    @pulumi.getter(name="transitRouterMulticastDomainName")
    def transit_router_multicast_domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "transit_router_multicast_domain_name")

    @transit_router_multicast_domain_name.setter
    def transit_router_multicast_domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_multicast_domain_name", value)


class TransitRouterMulticastDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_description: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud Enterprise Network (CEN) Transit Router Multicast Domain resource.

        For information about Cloud Enterprise Network (CEN) Transit Router Multicast Domain and how to use it, see [What is Transit Router Multicast Domain](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutermulticastdomain).

        > **NOTE:** Available since v1.195.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_instance = alicloud.cen.Instance("exampleInstance",
            cen_instance_name="tf_example",
            description="an example for cen")
        example_transit_router = alicloud.cen.TransitRouter("exampleTransitRouter",
            transit_router_name="tf_example",
            cen_id=example_instance.id,
            support_multicast=True)
        example_transit_router_multicast_domain = alicloud.cen.TransitRouterMulticastDomain("exampleTransitRouterMulticastDomain",
            transit_router_id=example_transit_router.transit_router_id,
            transit_router_multicast_domain_name="tf_example",
            transit_router_multicast_domain_description="tf_example")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Cloud Enterprise Network (CEN) Transit Router Multicast Domain can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        :param pulumi.Input[str] transit_router_multicast_domain_description: The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        :param pulumi.Input[str] transit_router_multicast_domain_name: The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransitRouterMulticastDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Enterprise Network (CEN) Transit Router Multicast Domain resource.

        For information about Cloud Enterprise Network (CEN) Transit Router Multicast Domain and how to use it, see [What is Transit Router Multicast Domain](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutermulticastdomain).

        > **NOTE:** Available since v1.195.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_instance = alicloud.cen.Instance("exampleInstance",
            cen_instance_name="tf_example",
            description="an example for cen")
        example_transit_router = alicloud.cen.TransitRouter("exampleTransitRouter",
            transit_router_name="tf_example",
            cen_id=example_instance.id,
            support_multicast=True)
        example_transit_router_multicast_domain = alicloud.cen.TransitRouterMulticastDomain("exampleTransitRouterMulticastDomain",
            transit_router_id=example_transit_router.transit_router_id,
            transit_router_multicast_domain_name="tf_example",
            transit_router_multicast_domain_description="tf_example")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Cloud Enterprise Network (CEN) Transit Router Multicast Domain can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain example <id>
        ```

        :param str resource_name: The name of the resource.
        :param TransitRouterMulticastDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransitRouterMulticastDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_description: Optional[pulumi.Input[str]] = None,
                 transit_router_multicast_domain_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransitRouterMulticastDomainArgs.__new__(TransitRouterMulticastDomainArgs)

            __props__.__dict__["tags"] = tags
            if transit_router_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_id'")
            __props__.__dict__["transit_router_id"] = transit_router_id
            __props__.__dict__["transit_router_multicast_domain_description"] = transit_router_multicast_domain_description
            __props__.__dict__["transit_router_multicast_domain_name"] = transit_router_multicast_domain_name
            __props__.__dict__["status"] = None
        super(TransitRouterMulticastDomain, __self__).__init__(
            'alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            transit_router_id: Optional[pulumi.Input[str]] = None,
            transit_router_multicast_domain_description: Optional[pulumi.Input[str]] = None,
            transit_router_multicast_domain_name: Optional[pulumi.Input[str]] = None) -> 'TransitRouterMulticastDomain':
        """
        Get an existing TransitRouterMulticastDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] status: The status of the Transit Router Multicast Domain.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        :param pulumi.Input[str] transit_router_multicast_domain_description: The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        :param pulumi.Input[str] transit_router_multicast_domain_name: The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransitRouterMulticastDomainState.__new__(_TransitRouterMulticastDomainState)

        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["transit_router_id"] = transit_router_id
        __props__.__dict__["transit_router_multicast_domain_description"] = transit_router_multicast_domain_description
        __props__.__dict__["transit_router_multicast_domain_name"] = transit_router_multicast_domain_name
        return TransitRouterMulticastDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the Transit Router Multicast Domain.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> pulumi.Output[str]:
        """
        The ID of the transit router.
        """
        return pulumi.get(self, "transit_router_id")

    @property
    @pulumi.getter(name="transitRouterMulticastDomainDescription")
    def transit_router_multicast_domain_description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "transit_router_multicast_domain_description")

    @property
    @pulumi.getter(name="transitRouterMulticastDomainName")
    def transit_router_multicast_domain_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "transit_router_multicast_domain_name")

