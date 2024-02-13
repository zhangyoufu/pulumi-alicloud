# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TransitRouterCidrArgs', 'TransitRouterCidr']

@pulumi.input_type
class TransitRouterCidrArgs:
    def __init__(__self__, *,
                 cidr: pulumi.Input[str],
                 transit_router_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 publish_cidr_route: Optional[pulumi.Input[bool]] = None,
                 transit_router_cidr_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TransitRouterCidr resource.
        :param pulumi.Input[str] cidr: The cidr of the transit router.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        :param pulumi.Input[str] description: The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] publish_cidr_route: Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        :param pulumi.Input[str] transit_router_cidr_name: The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        pulumi.set(__self__, "cidr", cidr)
        pulumi.set(__self__, "transit_router_id", transit_router_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if publish_cidr_route is not None:
            pulumi.set(__self__, "publish_cidr_route", publish_cidr_route)
        if transit_router_cidr_name is not None:
            pulumi.set(__self__, "transit_router_cidr_name", transit_router_cidr_name)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Input[str]:
        """
        The cidr of the transit router.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr", value)

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
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="publishCidrRoute")
    def publish_cidr_route(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        """
        return pulumi.get(self, "publish_cidr_route")

    @publish_cidr_route.setter
    def publish_cidr_route(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "publish_cidr_route", value)

    @property
    @pulumi.getter(name="transitRouterCidrName")
    def transit_router_cidr_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "transit_router_cidr_name")

    @transit_router_cidr_name.setter
    def transit_router_cidr_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_cidr_name", value)


@pulumi.input_type
class _TransitRouterCidrState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 publish_cidr_route: Optional[pulumi.Input[bool]] = None,
                 transit_router_cidr_id: Optional[pulumi.Input[str]] = None,
                 transit_router_cidr_name: Optional[pulumi.Input[str]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TransitRouterCidr resources.
        :param pulumi.Input[str] cidr: The cidr of the transit router.
        :param pulumi.Input[str] description: The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] publish_cidr_route: Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        :param pulumi.Input[str] transit_router_cidr_id: The ID of the transit router cidr.
        :param pulumi.Input[str] transit_router_cidr_name: The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if publish_cidr_route is not None:
            pulumi.set(__self__, "publish_cidr_route", publish_cidr_route)
        if transit_router_cidr_id is not None:
            pulumi.set(__self__, "transit_router_cidr_id", transit_router_cidr_id)
        if transit_router_cidr_name is not None:
            pulumi.set(__self__, "transit_router_cidr_name", transit_router_cidr_name)
        if transit_router_id is not None:
            pulumi.set(__self__, "transit_router_id", transit_router_id)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The cidr of the transit router.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="publishCidrRoute")
    def publish_cidr_route(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        """
        return pulumi.get(self, "publish_cidr_route")

    @publish_cidr_route.setter
    def publish_cidr_route(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "publish_cidr_route", value)

    @property
    @pulumi.getter(name="transitRouterCidrId")
    def transit_router_cidr_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the transit router cidr.
        """
        return pulumi.get(self, "transit_router_cidr_id")

    @transit_router_cidr_id.setter
    def transit_router_cidr_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_cidr_id", value)

    @property
    @pulumi.getter(name="transitRouterCidrName")
    def transit_router_cidr_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "transit_router_cidr_name")

    @transit_router_cidr_name.setter
    def transit_router_cidr_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_cidr_name", value)

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


class TransitRouterCidr(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 publish_cidr_route: Optional[pulumi.Input[bool]] = None,
                 transit_router_cidr_name: Optional[pulumi.Input[str]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud Enterprise Network (CEN) Transit Router Cidr resource.

        For information about Cloud Enterprise Network (CEN) Transit Router Cidr and how to use it, see [What is Transit Router Cidr](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/createtransitroutercidr).

        > **NOTE:** Available since v1.193.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_instance = alicloud.cen.Instance("exampleInstance",
            cen_instance_name="tf_example",
            description="an example for cen")
        example_transit_router = alicloud.cen.TransitRouter("exampleTransitRouter",
            transit_router_name="tf_example",
            cen_id=example_instance.id)
        example_transit_router_cidr = alicloud.cen.TransitRouterCidr("exampleTransitRouterCidr",
            transit_router_id=example_transit_router.transit_router_id,
            cidr="192.168.0.0/16",
            transit_router_cidr_name="tf_example",
            description="tf_example",
            publish_cidr_route=True)
        ```

        ## Import

        Cloud Enterprise Network (CEN) Transit Router Cidr can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/transitRouterCidr:TransitRouterCidr default <transit_router_id>:<transit_router_cidr_id>.
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: The cidr of the transit router.
        :param pulumi.Input[str] description: The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] publish_cidr_route: Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        :param pulumi.Input[str] transit_router_cidr_name: The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransitRouterCidrArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Enterprise Network (CEN) Transit Router Cidr resource.

        For information about Cloud Enterprise Network (CEN) Transit Router Cidr and how to use it, see [What is Transit Router Cidr](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/createtransitroutercidr).

        > **NOTE:** Available since v1.193.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_instance = alicloud.cen.Instance("exampleInstance",
            cen_instance_name="tf_example",
            description="an example for cen")
        example_transit_router = alicloud.cen.TransitRouter("exampleTransitRouter",
            transit_router_name="tf_example",
            cen_id=example_instance.id)
        example_transit_router_cidr = alicloud.cen.TransitRouterCidr("exampleTransitRouterCidr",
            transit_router_id=example_transit_router.transit_router_id,
            cidr="192.168.0.0/16",
            transit_router_cidr_name="tf_example",
            description="tf_example",
            publish_cidr_route=True)
        ```

        ## Import

        Cloud Enterprise Network (CEN) Transit Router Cidr can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/transitRouterCidr:TransitRouterCidr default <transit_router_id>:<transit_router_cidr_id>.
        ```

        :param str resource_name: The name of the resource.
        :param TransitRouterCidrArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransitRouterCidrArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 publish_cidr_route: Optional[pulumi.Input[bool]] = None,
                 transit_router_cidr_name: Optional[pulumi.Input[str]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransitRouterCidrArgs.__new__(TransitRouterCidrArgs)

            if cidr is None and not opts.urn:
                raise TypeError("Missing required property 'cidr'")
            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["description"] = description
            __props__.__dict__["publish_cidr_route"] = publish_cidr_route
            __props__.__dict__["transit_router_cidr_name"] = transit_router_cidr_name
            if transit_router_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_id'")
            __props__.__dict__["transit_router_id"] = transit_router_id
            __props__.__dict__["transit_router_cidr_id"] = None
        super(TransitRouterCidr, __self__).__init__(
            'alicloud:cen/transitRouterCidr:TransitRouterCidr',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            publish_cidr_route: Optional[pulumi.Input[bool]] = None,
            transit_router_cidr_id: Optional[pulumi.Input[str]] = None,
            transit_router_cidr_name: Optional[pulumi.Input[str]] = None,
            transit_router_id: Optional[pulumi.Input[str]] = None) -> 'TransitRouterCidr':
        """
        Get an existing TransitRouterCidr resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: The cidr of the transit router.
        :param pulumi.Input[str] description: The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] publish_cidr_route: Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        :param pulumi.Input[str] transit_router_cidr_id: The ID of the transit router cidr.
        :param pulumi.Input[str] transit_router_cidr_name: The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransitRouterCidrState.__new__(_TransitRouterCidrState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["description"] = description
        __props__.__dict__["publish_cidr_route"] = publish_cidr_route
        __props__.__dict__["transit_router_cidr_id"] = transit_router_cidr_id
        __props__.__dict__["transit_router_cidr_name"] = transit_router_cidr_name
        __props__.__dict__["transit_router_id"] = transit_router_id
        return TransitRouterCidr(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        The cidr of the transit router.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="publishCidrRoute")
    def publish_cidr_route(self) -> pulumi.Output[bool]:
        """
        Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        """
        return pulumi.get(self, "publish_cidr_route")

    @property
    @pulumi.getter(name="transitRouterCidrId")
    def transit_router_cidr_id(self) -> pulumi.Output[str]:
        """
        The ID of the transit router cidr.
        """
        return pulumi.get(self, "transit_router_cidr_id")

    @property
    @pulumi.getter(name="transitRouterCidrName")
    def transit_router_cidr_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "transit_router_cidr_name")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> pulumi.Output[str]:
        """
        The ID of the transit router.
        """
        return pulumi.get(self, "transit_router_id")

