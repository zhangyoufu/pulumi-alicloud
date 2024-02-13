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

__all__ = ['CustomLineArgs', 'CustomLine']

@pulumi.input_type
class CustomLineArgs:
    def __init__(__self__, *,
                 custom_line_name: pulumi.Input[str],
                 domain_name: pulumi.Input[str],
                 ip_segment_lists: pulumi.Input[Sequence[pulumi.Input['CustomLineIpSegmentListArgs']]],
                 lang: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CustomLine resource.
        :param pulumi.Input[str] custom_line_name: The name of the Custom Line.
        :param pulumi.Input[str] domain_name: The Domain name.
        :param pulumi.Input[Sequence[pulumi.Input['CustomLineIpSegmentListArgs']]] ip_segment_lists: The IP segment list. See `ip_segment_list` below for details.
        :param pulumi.Input[str] lang: The lang.
        """
        pulumi.set(__self__, "custom_line_name", custom_line_name)
        pulumi.set(__self__, "domain_name", domain_name)
        pulumi.set(__self__, "ip_segment_lists", ip_segment_lists)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)

    @property
    @pulumi.getter(name="customLineName")
    def custom_line_name(self) -> pulumi.Input[str]:
        """
        The name of the Custom Line.
        """
        return pulumi.get(self, "custom_line_name")

    @custom_line_name.setter
    def custom_line_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "custom_line_name", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        The Domain name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="ipSegmentLists")
    def ip_segment_lists(self) -> pulumi.Input[Sequence[pulumi.Input['CustomLineIpSegmentListArgs']]]:
        """
        The IP segment list. See `ip_segment_list` below for details.
        """
        return pulumi.get(self, "ip_segment_lists")

    @ip_segment_lists.setter
    def ip_segment_lists(self, value: pulumi.Input[Sequence[pulumi.Input['CustomLineIpSegmentListArgs']]]):
        pulumi.set(self, "ip_segment_lists", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The lang.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)


@pulumi.input_type
class _CustomLineState:
    def __init__(__self__, *,
                 custom_line_name: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 ip_segment_lists: Optional[pulumi.Input[Sequence[pulumi.Input['CustomLineIpSegmentListArgs']]]] = None,
                 lang: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomLine resources.
        :param pulumi.Input[str] custom_line_name: The name of the Custom Line.
        :param pulumi.Input[str] domain_name: The Domain name.
        :param pulumi.Input[Sequence[pulumi.Input['CustomLineIpSegmentListArgs']]] ip_segment_lists: The IP segment list. See `ip_segment_list` below for details.
        :param pulumi.Input[str] lang: The lang.
        """
        if custom_line_name is not None:
            pulumi.set(__self__, "custom_line_name", custom_line_name)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if ip_segment_lists is not None:
            pulumi.set(__self__, "ip_segment_lists", ip_segment_lists)
        if lang is not None:
            pulumi.set(__self__, "lang", lang)

    @property
    @pulumi.getter(name="customLineName")
    def custom_line_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Custom Line.
        """
        return pulumi.get(self, "custom_line_name")

    @custom_line_name.setter
    def custom_line_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_line_name", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Domain name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="ipSegmentLists")
    def ip_segment_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomLineIpSegmentListArgs']]]]:
        """
        The IP segment list. See `ip_segment_list` below for details.
        """
        return pulumi.get(self, "ip_segment_lists")

    @ip_segment_lists.setter
    def ip_segment_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomLineIpSegmentListArgs']]]]):
        pulumi.set(self, "ip_segment_lists", value)

    @property
    @pulumi.getter
    def lang(self) -> Optional[pulumi.Input[str]]:
        """
        The lang.
        """
        return pulumi.get(self, "lang")

    @lang.setter
    def lang(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lang", value)


class CustomLine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_line_name: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 ip_segment_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomLineIpSegmentListArgs']]]]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Alidns Custom Line resource.

        For information about Alidns Custom Line and how to use it, see [What is Custom Line](https://www.alibabacloud.com/help/en/doc-detail/145059.html).

        > **NOTE:** Available since v1.151.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.dns.CustomLine("default",
            custom_line_name="tf-example",
            domain_name="alicloud-provider.com",
            ip_segment_lists=[alicloud.dns.CustomLineIpSegmentListArgs(
                end_ip="192.0.2.125",
                start_ip="192.0.2.123",
            )])
        ```

        ## Import

        Alidns Custom Line can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dns/customLine:CustomLine example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] custom_line_name: The name of the Custom Line.
        :param pulumi.Input[str] domain_name: The Domain name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomLineIpSegmentListArgs']]]] ip_segment_lists: The IP segment list. See `ip_segment_list` below for details.
        :param pulumi.Input[str] lang: The lang.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomLineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Alidns Custom Line resource.

        For information about Alidns Custom Line and how to use it, see [What is Custom Line](https://www.alibabacloud.com/help/en/doc-detail/145059.html).

        > **NOTE:** Available since v1.151.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.dns.CustomLine("default",
            custom_line_name="tf-example",
            domain_name="alicloud-provider.com",
            ip_segment_lists=[alicloud.dns.CustomLineIpSegmentListArgs(
                end_ip="192.0.2.125",
                start_ip="192.0.2.123",
            )])
        ```

        ## Import

        Alidns Custom Line can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dns/customLine:CustomLine example <id>
        ```

        :param str resource_name: The name of the resource.
        :param CustomLineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomLineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_line_name: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 ip_segment_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomLineIpSegmentListArgs']]]]] = None,
                 lang: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomLineArgs.__new__(CustomLineArgs)

            if custom_line_name is None and not opts.urn:
                raise TypeError("Missing required property 'custom_line_name'")
            __props__.__dict__["custom_line_name"] = custom_line_name
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            if ip_segment_lists is None and not opts.urn:
                raise TypeError("Missing required property 'ip_segment_lists'")
            __props__.__dict__["ip_segment_lists"] = ip_segment_lists
            __props__.__dict__["lang"] = lang
        super(CustomLine, __self__).__init__(
            'alicloud:dns/customLine:CustomLine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_line_name: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            ip_segment_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomLineIpSegmentListArgs']]]]] = None,
            lang: Optional[pulumi.Input[str]] = None) -> 'CustomLine':
        """
        Get an existing CustomLine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] custom_line_name: The name of the Custom Line.
        :param pulumi.Input[str] domain_name: The Domain name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomLineIpSegmentListArgs']]]] ip_segment_lists: The IP segment list. See `ip_segment_list` below for details.
        :param pulumi.Input[str] lang: The lang.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomLineState.__new__(_CustomLineState)

        __props__.__dict__["custom_line_name"] = custom_line_name
        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["ip_segment_lists"] = ip_segment_lists
        __props__.__dict__["lang"] = lang
        return CustomLine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customLineName")
    def custom_line_name(self) -> pulumi.Output[str]:
        """
        The name of the Custom Line.
        """
        return pulumi.get(self, "custom_line_name")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        The Domain name.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="ipSegmentLists")
    def ip_segment_lists(self) -> pulumi.Output[Sequence['outputs.CustomLineIpSegmentList']]:
        """
        The IP segment list. See `ip_segment_list` below for details.
        """
        return pulumi.get(self, "ip_segment_lists")

    @property
    @pulumi.getter
    def lang(self) -> pulumi.Output[Optional[str]]:
        """
        The lang.
        """
        return pulumi.get(self, "lang")

