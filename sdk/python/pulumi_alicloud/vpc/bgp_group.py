# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BgpGroupArgs', 'BgpGroup']

@pulumi.input_type
class BgpGroupArgs:
    def __init__(__self__, *,
                 peer_asn: pulumi.Input[int],
                 router_id: pulumi.Input[str],
                 auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_fake_asn: Optional[pulumi.Input[bool]] = None,
                 local_asn: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a BgpGroup resource.
        :param pulumi.Input[int] peer_asn: The AS number of the BGP peer.
        :param pulumi.Input[str] router_id: The ID of the VBR.
        :param pulumi.Input[str] auth_key: The authentication key of the BGP group.
        :param pulumi.Input[str] bgp_group_name: The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] description: The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] is_fake_asn: The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
        :param pulumi.Input[int] local_asn: The AS number on the Alibaba Cloud side.
        """
        pulumi.set(__self__, "peer_asn", peer_asn)
        pulumi.set(__self__, "router_id", router_id)
        if auth_key is not None:
            pulumi.set(__self__, "auth_key", auth_key)
        if bgp_group_name is not None:
            pulumi.set(__self__, "bgp_group_name", bgp_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_fake_asn is not None:
            pulumi.set(__self__, "is_fake_asn", is_fake_asn)
        if local_asn is not None:
            pulumi.set(__self__, "local_asn", local_asn)

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> pulumi.Input[int]:
        """
        The AS number of the BGP peer.
        """
        return pulumi.get(self, "peer_asn")

    @peer_asn.setter
    def peer_asn(self, value: pulumi.Input[int]):
        pulumi.set(self, "peer_asn", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Input[str]:
        """
        The ID of the VBR.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "router_id", value)

    @property
    @pulumi.getter(name="authKey")
    def auth_key(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication key of the BGP group.
        """
        return pulumi.get(self, "auth_key")

    @auth_key.setter
    def auth_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_key", value)

    @property
    @pulumi.getter(name="bgpGroupName")
    def bgp_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "bgp_group_name")

    @bgp_group_name.setter
    def bgp_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_group_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isFakeAsn")
    def is_fake_asn(self) -> Optional[pulumi.Input[bool]]:
        """
        The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
        """
        return pulumi.get(self, "is_fake_asn")

    @is_fake_asn.setter
    def is_fake_asn(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_fake_asn", value)

    @property
    @pulumi.getter(name="localAsn")
    def local_asn(self) -> Optional[pulumi.Input[int]]:
        """
        The AS number on the Alibaba Cloud side.
        """
        return pulumi.get(self, "local_asn")

    @local_asn.setter
    def local_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "local_asn", value)


@pulumi.input_type
class _BgpGroupState:
    def __init__(__self__, *,
                 auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_fake_asn: Optional[pulumi.Input[bool]] = None,
                 local_asn: Optional[pulumi.Input[int]] = None,
                 peer_asn: Optional[pulumi.Input[int]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BgpGroup resources.
        :param pulumi.Input[str] auth_key: The authentication key of the BGP group.
        :param pulumi.Input[str] bgp_group_name: The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] description: The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] is_fake_asn: The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
        :param pulumi.Input[int] local_asn: The AS number on the Alibaba Cloud side.
        :param pulumi.Input[int] peer_asn: The AS number of the BGP peer.
        :param pulumi.Input[str] router_id: The ID of the VBR.
        :param pulumi.Input[str] status: The status of the resource.
        """
        if auth_key is not None:
            pulumi.set(__self__, "auth_key", auth_key)
        if bgp_group_name is not None:
            pulumi.set(__self__, "bgp_group_name", bgp_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_fake_asn is not None:
            pulumi.set(__self__, "is_fake_asn", is_fake_asn)
        if local_asn is not None:
            pulumi.set(__self__, "local_asn", local_asn)
        if peer_asn is not None:
            pulumi.set(__self__, "peer_asn", peer_asn)
        if router_id is not None:
            pulumi.set(__self__, "router_id", router_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="authKey")
    def auth_key(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication key of the BGP group.
        """
        return pulumi.get(self, "auth_key")

    @auth_key.setter
    def auth_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_key", value)

    @property
    @pulumi.getter(name="bgpGroupName")
    def bgp_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "bgp_group_name")

    @bgp_group_name.setter
    def bgp_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_group_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isFakeAsn")
    def is_fake_asn(self) -> Optional[pulumi.Input[bool]]:
        """
        The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
        """
        return pulumi.get(self, "is_fake_asn")

    @is_fake_asn.setter
    def is_fake_asn(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_fake_asn", value)

    @property
    @pulumi.getter(name="localAsn")
    def local_asn(self) -> Optional[pulumi.Input[int]]:
        """
        The AS number on the Alibaba Cloud side.
        """
        return pulumi.get(self, "local_asn")

    @local_asn.setter
    def local_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "local_asn", value)

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> Optional[pulumi.Input[int]]:
        """
        The AS number of the BGP peer.
        """
        return pulumi.get(self, "peer_asn")

    @peer_asn.setter
    def peer_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "peer_asn", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VBR.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "router_id", value)

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


class BgpGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_fake_asn: Optional[pulumi.Input[bool]] = None,
                 local_asn: Optional[pulumi.Input[int]] = None,
                 peer_asn: Optional[pulumi.Input[int]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VPC Bgp Group resource.

        For information about VPC Bgp Group and how to use it, see [What is Bgp Group](https://www.alibabacloud.com/help/en/doc-detail/91267.html).

        > **NOTE:** Available since v1.152.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        example_physical_connections = alicloud.expressconnect.get_physical_connections(name_regex="^preserved-NODELETING")
        vlan_id = random.RandomInteger("vlanId",
            max=2999,
            min=1)
        example_virtual_border_router = alicloud.expressconnect.VirtualBorderRouter("exampleVirtualBorderRouter",
            local_gateway_ip="10.0.0.1",
            peer_gateway_ip="10.0.0.2",
            peering_subnet_mask="255.255.255.252",
            physical_connection_id=example_physical_connections.connections[0].id,
            virtual_border_router_name=name,
            vlan_id=vlan_id.id,
            min_rx_interval=1000,
            min_tx_interval=1000,
            detect_multiplier=10)
        example_bgp_group = alicloud.vpc.BgpGroup("exampleBgpGroup",
            auth_key="YourPassword+12345678",
            bgp_group_name=name,
            description=name,
            peer_asn=1111,
            router_id=example_virtual_border_router.id,
            is_fake_asn=True)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        VPC Bgp Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpc/bgpGroup:BgpGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_key: The authentication key of the BGP group.
        :param pulumi.Input[str] bgp_group_name: The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] description: The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] is_fake_asn: The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
        :param pulumi.Input[int] local_asn: The AS number on the Alibaba Cloud side.
        :param pulumi.Input[int] peer_asn: The AS number of the BGP peer.
        :param pulumi.Input[str] router_id: The ID of the VBR.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BgpGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Bgp Group resource.

        For information about VPC Bgp Group and how to use it, see [What is Bgp Group](https://www.alibabacloud.com/help/en/doc-detail/91267.html).

        > **NOTE:** Available since v1.152.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        example_physical_connections = alicloud.expressconnect.get_physical_connections(name_regex="^preserved-NODELETING")
        vlan_id = random.RandomInteger("vlanId",
            max=2999,
            min=1)
        example_virtual_border_router = alicloud.expressconnect.VirtualBorderRouter("exampleVirtualBorderRouter",
            local_gateway_ip="10.0.0.1",
            peer_gateway_ip="10.0.0.2",
            peering_subnet_mask="255.255.255.252",
            physical_connection_id=example_physical_connections.connections[0].id,
            virtual_border_router_name=name,
            vlan_id=vlan_id.id,
            min_rx_interval=1000,
            min_tx_interval=1000,
            detect_multiplier=10)
        example_bgp_group = alicloud.vpc.BgpGroup("exampleBgpGroup",
            auth_key="YourPassword+12345678",
            bgp_group_name=name,
            description=name,
            peer_asn=1111,
            router_id=example_virtual_border_router.id,
            is_fake_asn=True)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        VPC Bgp Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:vpc/bgpGroup:BgpGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param BgpGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_fake_asn: Optional[pulumi.Input[bool]] = None,
                 local_asn: Optional[pulumi.Input[int]] = None,
                 peer_asn: Optional[pulumi.Input[int]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpGroupArgs.__new__(BgpGroupArgs)

            __props__.__dict__["auth_key"] = auth_key
            __props__.__dict__["bgp_group_name"] = bgp_group_name
            __props__.__dict__["description"] = description
            __props__.__dict__["is_fake_asn"] = is_fake_asn
            __props__.__dict__["local_asn"] = local_asn
            if peer_asn is None and not opts.urn:
                raise TypeError("Missing required property 'peer_asn'")
            __props__.__dict__["peer_asn"] = peer_asn
            if router_id is None and not opts.urn:
                raise TypeError("Missing required property 'router_id'")
            __props__.__dict__["router_id"] = router_id
            __props__.__dict__["status"] = None
        super(BgpGroup, __self__).__init__(
            'alicloud:vpc/bgpGroup:BgpGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_key: Optional[pulumi.Input[str]] = None,
            bgp_group_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            is_fake_asn: Optional[pulumi.Input[bool]] = None,
            local_asn: Optional[pulumi.Input[int]] = None,
            peer_asn: Optional[pulumi.Input[int]] = None,
            router_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'BgpGroup':
        """
        Get an existing BgpGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_key: The authentication key of the BGP group.
        :param pulumi.Input[str] bgp_group_name: The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] description: The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[bool] is_fake_asn: The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
        :param pulumi.Input[int] local_asn: The AS number on the Alibaba Cloud side.
        :param pulumi.Input[int] peer_asn: The AS number of the BGP peer.
        :param pulumi.Input[str] router_id: The ID of the VBR.
        :param pulumi.Input[str] status: The status of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpGroupState.__new__(_BgpGroupState)

        __props__.__dict__["auth_key"] = auth_key
        __props__.__dict__["bgp_group_name"] = bgp_group_name
        __props__.__dict__["description"] = description
        __props__.__dict__["is_fake_asn"] = is_fake_asn
        __props__.__dict__["local_asn"] = local_asn
        __props__.__dict__["peer_asn"] = peer_asn
        __props__.__dict__["router_id"] = router_id
        __props__.__dict__["status"] = status
        return BgpGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authKey")
    def auth_key(self) -> pulumi.Output[Optional[str]]:
        """
        The authentication key of the BGP group.
        """
        return pulumi.get(self, "auth_key")

    @property
    @pulumi.getter(name="bgpGroupName")
    def bgp_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the BGP group. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "bgp_group_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the BGP group. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isFakeAsn")
    def is_fake_asn(self) -> pulumi.Output[bool]:
        """
        The is fake asn. A router that runs BGP typically belongs to only one AS. In some cases, for example, the AS needs to be migrated or is merged with another AS, a new AS number replaces the original one.
        """
        return pulumi.get(self, "is_fake_asn")

    @property
    @pulumi.getter(name="localAsn")
    def local_asn(self) -> pulumi.Output[int]:
        """
        The AS number on the Alibaba Cloud side.
        """
        return pulumi.get(self, "local_asn")

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> pulumi.Output[int]:
        """
        The AS number of the BGP peer.
        """
        return pulumi.get(self, "peer_asn")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Output[str]:
        """
        The ID of the VBR.
        """
        return pulumi.get(self, "router_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

