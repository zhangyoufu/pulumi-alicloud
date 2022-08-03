# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['Ipv6GatewayArgs', 'Ipv6Gateway']

@pulumi.input_type
class Ipv6GatewayArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ipv6_gateway_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ipv6Gateway resource.
        :param pulumi.Input[str] vpc_id: The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
        :param pulumi.Input[str] description: The description of the IPv6 gateway. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        :param pulumi.Input[str] ipv6_gateway_name: The name of the IPv6 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] spec: The edition of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ipv6_gateway_name is not None:
            pulumi.set(__self__, "ipv6_gateway_name", ipv6_gateway_name)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the IPv6 gateway. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipv6GatewayName")
    def ipv6_gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IPv6 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "ipv6_gateway_name")

    @ipv6_gateway_name.setter
    def ipv6_gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_gateway_name", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input[str]]:
        """
        The edition of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spec", value)


@pulumi.input_type
class _Ipv6GatewayState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 ipv6_gateway_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ipv6Gateway resources.
        :param pulumi.Input[str] description: The description of the IPv6 gateway. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        :param pulumi.Input[str] ipv6_gateway_name: The name of the IPv6 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] spec: The edition of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
        :param pulumi.Input[str] status: The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
        :param pulumi.Input[str] vpc_id: The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ipv6_gateway_name is not None:
            pulumi.set(__self__, "ipv6_gateway_name", ipv6_gateway_name)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the IPv6 gateway. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipv6GatewayName")
    def ipv6_gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IPv6 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "ipv6_gateway_name")

    @ipv6_gateway_name.setter
    def ipv6_gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_gateway_name", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input[str]]:
        """
        The edition of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class Ipv6Gateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ipv6_gateway_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VPC Ipv6 Gateway resource.

        For information about VPC Ipv6 Gateway and how to use it, see [What is Ipv6 Gateway](https://www.alibabacloud.com/help/doc-detail/102214.htm).

        > **NOTE:** Available in v1.142.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.vpc.Network("default",
            vpc_name="example_value",
            enable_ipv6=True)
        example = alicloud.vpc.Ipv6Gateway("example",
            ipv6_gateway_name="example_value",
            vpc_id=default.id)
        ```

        ## Import

        VPC Ipv6 Gateway can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/ipv6Gateway:Ipv6Gateway example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the IPv6 gateway. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        :param pulumi.Input[str] ipv6_gateway_name: The name of the IPv6 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] spec: The edition of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
        :param pulumi.Input[str] vpc_id: The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ipv6GatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Ipv6 Gateway resource.

        For information about VPC Ipv6 Gateway and how to use it, see [What is Ipv6 Gateway](https://www.alibabacloud.com/help/doc-detail/102214.htm).

        > **NOTE:** Available in v1.142.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.vpc.Network("default",
            vpc_name="example_value",
            enable_ipv6=True)
        example = alicloud.vpc.Ipv6Gateway("example",
            ipv6_gateway_name="example_value",
            vpc_id=default.id)
        ```

        ## Import

        VPC Ipv6 Gateway can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/ipv6Gateway:Ipv6Gateway example <id>
        ```

        :param str resource_name: The name of the resource.
        :param Ipv6GatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ipv6GatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ipv6_gateway_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Ipv6GatewayArgs.__new__(Ipv6GatewayArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["ipv6_gateway_name"] = ipv6_gateway_name
            __props__.__dict__["spec"] = spec
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["status"] = None
        super(Ipv6Gateway, __self__).__init__(
            'alicloud:vpc/ipv6Gateway:Ipv6Gateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            ipv6_gateway_name: Optional[pulumi.Input[str]] = None,
            spec: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'Ipv6Gateway':
        """
        Get an existing Ipv6Gateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the IPv6 gateway. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        :param pulumi.Input[str] ipv6_gateway_name: The name of the IPv6 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] spec: The edition of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
        :param pulumi.Input[str] status: The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
        :param pulumi.Input[str] vpc_id: The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ipv6GatewayState.__new__(_Ipv6GatewayState)

        __props__.__dict__["description"] = description
        __props__.__dict__["ipv6_gateway_name"] = ipv6_gateway_name
        __props__.__dict__["spec"] = spec
        __props__.__dict__["status"] = status
        __props__.__dict__["vpc_id"] = vpc_id
        return Ipv6Gateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the IPv6 gateway. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipv6GatewayName")
    def ipv6_gateway_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the IPv6 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "ipv6_gateway_name")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[str]:
        """
        The edition of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
        """
        return pulumi.get(self, "vpc_id")

