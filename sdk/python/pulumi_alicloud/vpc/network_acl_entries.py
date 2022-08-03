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

__all__ = ['NetworkAclEntriesArgs', 'NetworkAclEntries']

@pulumi.input_type
class NetworkAclEntriesArgs:
    def __init__(__self__, *,
                 network_acl_id: pulumi.Input[str],
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesEgressArgs']]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesIngressArgs']]]] = None):
        """
        The set of arguments for constructing a NetworkAclEntries resource.
        :param pulumi.Input[str] network_acl_id: The id of the network acl, the field can't be changed.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesEgressArgs']]] egresses: List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesIngressArgs']]] ingresses: List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        """
        pulumi.set(__self__, "network_acl_id", network_acl_id)
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)

    @property
    @pulumi.getter(name="networkAclId")
    def network_acl_id(self) -> pulumi.Input[str]:
        """
        The id of the network acl, the field can't be changed.
        """
        return pulumi.get(self, "network_acl_id")

    @network_acl_id.setter
    def network_acl_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_acl_id", value)

    @property
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesEgressArgs']]]]:
        """
        List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        """
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesEgressArgs']]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesIngressArgs']]]]:
        """
        List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        """
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesIngressArgs']]]]):
        pulumi.set(self, "ingresses", value)


@pulumi.input_type
class _NetworkAclEntriesState:
    def __init__(__self__, *,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesEgressArgs']]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesIngressArgs']]]] = None,
                 network_acl_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkAclEntries resources.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesEgressArgs']]] egresses: List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesIngressArgs']]] ingresses: List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        :param pulumi.Input[str] network_acl_id: The id of the network acl, the field can't be changed.
        """
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)
        if network_acl_id is not None:
            pulumi.set(__self__, "network_acl_id", network_acl_id)

    @property
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesEgressArgs']]]]:
        """
        List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        """
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesEgressArgs']]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesIngressArgs']]]]:
        """
        List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        """
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclEntriesIngressArgs']]]]):
        pulumi.set(self, "ingresses", value)

    @property
    @pulumi.getter(name="networkAclId")
    def network_acl_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the network acl, the field can't be changed.
        """
        return pulumi.get(self, "network_acl_id")

    @network_acl_id.setter
    def network_acl_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_acl_id", value)


class NetworkAclEntries(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesEgressArgs']]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesIngressArgs']]]]] = None,
                 network_acl_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "NetworkAclEntries"
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/12")
        default_network_acl = alicloud.vpc.NetworkAcl("defaultNetworkAcl", vpc_id=default_network.id)
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/21",
            zone_id=default_zones.zones[0].id)
        default_network_acl_attachment = alicloud.vpc.NetworkAclAttachment("defaultNetworkAclAttachment",
            network_acl_id=default_network_acl.id,
            resources=[alicloud.vpc.NetworkAclAttachmentResourceArgs(
                resource_id=default_switch.id,
                resource_type="VSwitch",
            )])
        default_network_acl_entries = alicloud.vpc.NetworkAclEntries("defaultNetworkAclEntries",
            network_acl_id=default_network_acl.id,
            ingresses=[alicloud.vpc.NetworkAclEntriesIngressArgs(
                protocol="all",
                port="-1/-1",
                source_cidr_ip="0.0.0.0/32",
                name=name,
                entry_type="custom",
                policy="accept",
                description=name,
            )],
            egresses=[alicloud.vpc.NetworkAclEntriesEgressArgs(
                protocol="all",
                port="-1/-1",
                destination_cidr_ip="0.0.0.0/32",
                name=name,
                entry_type="custom",
                policy="accept",
                description=name,
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesEgressArgs']]]] egresses: List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesIngressArgs']]]] ingresses: List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        :param pulumi.Input[str] network_acl_id: The id of the network acl, the field can't be changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkAclEntriesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "NetworkAclEntries"
        default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/12")
        default_network_acl = alicloud.vpc.NetworkAcl("defaultNetworkAcl", vpc_id=default_network.id)
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/21",
            zone_id=default_zones.zones[0].id)
        default_network_acl_attachment = alicloud.vpc.NetworkAclAttachment("defaultNetworkAclAttachment",
            network_acl_id=default_network_acl.id,
            resources=[alicloud.vpc.NetworkAclAttachmentResourceArgs(
                resource_id=default_switch.id,
                resource_type="VSwitch",
            )])
        default_network_acl_entries = alicloud.vpc.NetworkAclEntries("defaultNetworkAclEntries",
            network_acl_id=default_network_acl.id,
            ingresses=[alicloud.vpc.NetworkAclEntriesIngressArgs(
                protocol="all",
                port="-1/-1",
                source_cidr_ip="0.0.0.0/32",
                name=name,
                entry_type="custom",
                policy="accept",
                description=name,
            )],
            egresses=[alicloud.vpc.NetworkAclEntriesEgressArgs(
                protocol="all",
                port="-1/-1",
                destination_cidr_ip="0.0.0.0/32",
                name=name,
                entry_type="custom",
                policy="accept",
                description=name,
            )])
        ```

        :param str resource_name: The name of the resource.
        :param NetworkAclEntriesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkAclEntriesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesEgressArgs']]]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesIngressArgs']]]]] = None,
                 network_acl_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkAclEntriesArgs.__new__(NetworkAclEntriesArgs)

            __props__.__dict__["egresses"] = egresses
            __props__.__dict__["ingresses"] = ingresses
            if network_acl_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_acl_id'")
            __props__.__dict__["network_acl_id"] = network_acl_id
        super(NetworkAclEntries, __self__).__init__(
            'alicloud:vpc/networkAclEntries:NetworkAclEntries',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            egresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesEgressArgs']]]]] = None,
            ingresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesIngressArgs']]]]] = None,
            network_acl_id: Optional[pulumi.Input[str]] = None) -> 'NetworkAclEntries':
        """
        Get an existing NetworkAclEntries resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesEgressArgs']]]] egresses: List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAclEntriesIngressArgs']]]] ingresses: List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        :param pulumi.Input[str] network_acl_id: The id of the network acl, the field can't be changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkAclEntriesState.__new__(_NetworkAclEntriesState)

        __props__.__dict__["egresses"] = egresses
        __props__.__dict__["ingresses"] = ingresses
        __props__.__dict__["network_acl_id"] = network_acl_id
        return NetworkAclEntries(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def egresses(self) -> pulumi.Output[Optional[Sequence['outputs.NetworkAclEntriesEgress']]]:
        """
        List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        """
        return pulumi.get(self, "egresses")

    @property
    @pulumi.getter
    def ingresses(self) -> pulumi.Output[Optional[Sequence['outputs.NetworkAclEntriesIngress']]]:
        """
        List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        """
        return pulumi.get(self, "ingresses")

    @property
    @pulumi.getter(name="networkAclId")
    def network_acl_id(self) -> pulumi.Output[str]:
        """
        The id of the network acl, the field can't be changed.
        """
        return pulumi.get(self, "network_acl_id")

