# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LoadBalancerArgs', 'LoadBalancer']

@pulumi.input_type
class LoadBalancerArgs:
    def __init__(__self__, *,
                 ens_region_id: pulumi.Input[str],
                 load_balancer_spec: pulumi.Input[str],
                 network_id: pulumi.Input[str],
                 payment_type: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str],
                 load_balancer_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LoadBalancer resource.
        :param pulumi.Input[str] ens_region_id: The ID of the ENS node.
        :param pulumi.Input[str] load_balancer_spec: Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
        :param pulumi.Input[str] network_id: The network ID of the created edge load balancing (ELB) instance.
        :param pulumi.Input[str] payment_type: Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
        :param pulumi.Input[str] vswitch_id: The ID of the vSwitch to which the VPC instance belongs.
        :param pulumi.Input[str] load_balancer_name: Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
        """
        pulumi.set(__self__, "ens_region_id", ens_region_id)
        pulumi.set(__self__, "load_balancer_spec", load_balancer_spec)
        pulumi.set(__self__, "network_id", network_id)
        pulumi.set(__self__, "payment_type", payment_type)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if load_balancer_name is not None:
            pulumi.set(__self__, "load_balancer_name", load_balancer_name)

    @property
    @pulumi.getter(name="ensRegionId")
    def ens_region_id(self) -> pulumi.Input[str]:
        """
        The ID of the ENS node.
        """
        return pulumi.get(self, "ens_region_id")

    @ens_region_id.setter
    def ens_region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ens_region_id", value)

    @property
    @pulumi.getter(name="loadBalancerSpec")
    def load_balancer_spec(self) -> pulumi.Input[str]:
        """
        Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
        """
        return pulumi.get(self, "load_balancer_spec")

    @load_balancer_spec.setter
    def load_balancer_spec(self, value: pulumi.Input[str]):
        pulumi.set(self, "load_balancer_spec", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[str]:
        """
        The network ID of the created edge load balancing (ELB) instance.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> pulumi.Input[str]:
        """
        Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
        """
        return pulumi.get(self, "payment_type")

    @payment_type.setter
    def payment_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "payment_type", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        The ID of the vSwitch to which the VPC instance belongs.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
        """
        return pulumi.get(self, "load_balancer_name")

    @load_balancer_name.setter
    def load_balancer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_balancer_name", value)


@pulumi.input_type
class _LoadBalancerState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 ens_region_id: Optional[pulumi.Input[str]] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 load_balancer_spec: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LoadBalancer resources.
        :param pulumi.Input[str] create_time: The creation Time (UTC) of the load balancing instance.
        :param pulumi.Input[str] ens_region_id: The ID of the ENS node.
        :param pulumi.Input[str] load_balancer_name: Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
        :param pulumi.Input[str] load_balancer_spec: Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
        :param pulumi.Input[str] network_id: The network ID of the created edge load balancing (ELB) instance.
        :param pulumi.Input[str] payment_type: Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
        :param pulumi.Input[str] status: The status of the SLB instance.
        :param pulumi.Input[str] vswitch_id: The ID of the vSwitch to which the VPC instance belongs.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if ens_region_id is not None:
            pulumi.set(__self__, "ens_region_id", ens_region_id)
        if load_balancer_name is not None:
            pulumi.set(__self__, "load_balancer_name", load_balancer_name)
        if load_balancer_spec is not None:
            pulumi.set(__self__, "load_balancer_spec", load_balancer_spec)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if payment_type is not None:
            pulumi.set(__self__, "payment_type", payment_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation Time (UTC) of the load balancing instance.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="ensRegionId")
    def ens_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the ENS node.
        """
        return pulumi.get(self, "ens_region_id")

    @ens_region_id.setter
    def ens_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ens_region_id", value)

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
        """
        return pulumi.get(self, "load_balancer_name")

    @load_balancer_name.setter
    def load_balancer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_balancer_name", value)

    @property
    @pulumi.getter(name="loadBalancerSpec")
    def load_balancer_spec(self) -> Optional[pulumi.Input[str]]:
        """
        Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
        """
        return pulumi.get(self, "load_balancer_spec")

    @load_balancer_spec.setter
    def load_balancer_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_balancer_spec", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The network ID of the created edge load balancing (ELB) instance.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> Optional[pulumi.Input[str]]:
        """
        Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
        """
        return pulumi.get(self, "payment_type")

    @payment_type.setter
    def payment_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the SLB instance.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the vSwitch to which the VPC instance belongs.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class LoadBalancer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ens_region_id: Optional[pulumi.Input[str]] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 load_balancer_spec: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ENS Load Balancer resource. Load balancing. When you use it for the first time, please contact the product classmates to add a resource whitelist.

        For information about ENS Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/en/ens/developer-reference/api-createloadbalancer).

        > **NOTE:** Available since v1.213.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        network = alicloud.ens.Network("network",
            network_name=name,
            description=name,
            cidr_block="192.168.2.0/24",
            ens_region_id="cn-chenzhou-telecom_unicom_cmcc")
        switch = alicloud.ens.Vswitch("switch",
            description=name,
            cidr_block="192.168.2.0/24",
            vswitch_name=name,
            ens_region_id="cn-chenzhou-telecom_unicom_cmcc",
            network_id=network.id)
        default = alicloud.ens.LoadBalancer("default",
            load_balancer_name=name,
            payment_type="PayAsYouGo",
            ens_region_id="cn-chenzhou-telecom_unicom_cmcc",
            load_balancer_spec="elb.s1.small",
            vswitch_id=switch.id,
            network_id=network.id)
        ```

        ## Import

        ENS Load Balancer can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ens/loadBalancer:LoadBalancer example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ens_region_id: The ID of the ENS node.
        :param pulumi.Input[str] load_balancer_name: Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
        :param pulumi.Input[str] load_balancer_spec: Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
        :param pulumi.Input[str] network_id: The network ID of the created edge load balancing (ELB) instance.
        :param pulumi.Input[str] payment_type: Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
        :param pulumi.Input[str] vswitch_id: The ID of the vSwitch to which the VPC instance belongs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoadBalancerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ENS Load Balancer resource. Load balancing. When you use it for the first time, please contact the product classmates to add a resource whitelist.

        For information about ENS Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/en/ens/developer-reference/api-createloadbalancer).

        > **NOTE:** Available since v1.213.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        network = alicloud.ens.Network("network",
            network_name=name,
            description=name,
            cidr_block="192.168.2.0/24",
            ens_region_id="cn-chenzhou-telecom_unicom_cmcc")
        switch = alicloud.ens.Vswitch("switch",
            description=name,
            cidr_block="192.168.2.0/24",
            vswitch_name=name,
            ens_region_id="cn-chenzhou-telecom_unicom_cmcc",
            network_id=network.id)
        default = alicloud.ens.LoadBalancer("default",
            load_balancer_name=name,
            payment_type="PayAsYouGo",
            ens_region_id="cn-chenzhou-telecom_unicom_cmcc",
            load_balancer_spec="elb.s1.small",
            vswitch_id=switch.id,
            network_id=network.id)
        ```

        ## Import

        ENS Load Balancer can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ens/loadBalancer:LoadBalancer example <id>
        ```

        :param str resource_name: The name of the resource.
        :param LoadBalancerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoadBalancerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ens_region_id: Optional[pulumi.Input[str]] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 load_balancer_spec: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 payment_type: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoadBalancerArgs.__new__(LoadBalancerArgs)

            if ens_region_id is None and not opts.urn:
                raise TypeError("Missing required property 'ens_region_id'")
            __props__.__dict__["ens_region_id"] = ens_region_id
            __props__.__dict__["load_balancer_name"] = load_balancer_name
            if load_balancer_spec is None and not opts.urn:
                raise TypeError("Missing required property 'load_balancer_spec'")
            __props__.__dict__["load_balancer_spec"] = load_balancer_spec
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            if payment_type is None and not opts.urn:
                raise TypeError("Missing required property 'payment_type'")
            __props__.__dict__["payment_type"] = payment_type
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["status"] = None
        super(LoadBalancer, __self__).__init__(
            'alicloud:ens/loadBalancer:LoadBalancer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            ens_region_id: Optional[pulumi.Input[str]] = None,
            load_balancer_name: Optional[pulumi.Input[str]] = None,
            load_balancer_spec: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            payment_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'LoadBalancer':
        """
        Get an existing LoadBalancer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The creation Time (UTC) of the load balancing instance.
        :param pulumi.Input[str] ens_region_id: The ID of the ENS node.
        :param pulumi.Input[str] load_balancer_name: Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
        :param pulumi.Input[str] load_balancer_spec: Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
        :param pulumi.Input[str] network_id: The network ID of the created edge load balancing (ELB) instance.
        :param pulumi.Input[str] payment_type: Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
        :param pulumi.Input[str] status: The status of the SLB instance.
        :param pulumi.Input[str] vswitch_id: The ID of the vSwitch to which the VPC instance belongs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LoadBalancerState.__new__(_LoadBalancerState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["ens_region_id"] = ens_region_id
        __props__.__dict__["load_balancer_name"] = load_balancer_name
        __props__.__dict__["load_balancer_spec"] = load_balancer_spec
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["payment_type"] = payment_type
        __props__.__dict__["status"] = status
        __props__.__dict__["vswitch_id"] = vswitch_id
        return LoadBalancer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation Time (UTC) of the load balancing instance.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="ensRegionId")
    def ens_region_id(self) -> pulumi.Output[str]:
        """
        The ID of the ENS node.
        """
        return pulumi.get(self, "ens_region_id")

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the Server Load Balancer instanceRules:The length is 1~80 English or Chinese characters. When this parameter is not specified, the system randomly assigns an instance nameCannot start with `http://` and `https`.
        """
        return pulumi.get(self, "load_balancer_name")

    @property
    @pulumi.getter(name="loadBalancerSpec")
    def load_balancer_spec(self) -> pulumi.Output[str]:
        """
        Specifications of the Server Load Balancer instance. Valid values: elb.s1.small,elb.s3.medium,elb.s2.small,elb.s2.medium,elb.s3.small.
        """
        return pulumi.get(self, "load_balancer_spec")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        The network ID of the created edge load balancing (ELB) instance.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="paymentType")
    def payment_type(self) -> pulumi.Output[str]:
        """
        Server Load Balancer Instance Payment Type. Valid value: PayAsYouGo.
        """
        return pulumi.get(self, "payment_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the SLB instance.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The ID of the vSwitch to which the VPC instance belongs.
        """
        return pulumi.get(self, "vswitch_id")

