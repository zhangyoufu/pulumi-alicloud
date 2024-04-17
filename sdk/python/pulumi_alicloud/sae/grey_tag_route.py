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

__all__ = ['GreyTagRouteArgs', 'GreyTagRoute']

@pulumi.input_type
class GreyTagRouteArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 grey_tag_route_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 dubbo_rules: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleArgs']]]] = None,
                 sc_rules: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleArgs']]]] = None):
        """
        The set of arguments for constructing a GreyTagRoute resource.
        :param pulumi.Input[str] app_id: The ID  of the SAE Application.
        :param pulumi.Input[str] grey_tag_route_name: The name of GreyTagRoute.
        :param pulumi.Input[str] description: The description of GreyTagRoute.
        :param pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleArgs']]] dubbo_rules: The grayscale rule created for Dubbo Application. See `dubbo_rules` below.
        :param pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleArgs']]] sc_rules: The grayscale rule created for SpringCloud Application. See `sc_rules` below.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "grey_tag_route_name", grey_tag_route_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dubbo_rules is not None:
            pulumi.set(__self__, "dubbo_rules", dubbo_rules)
        if sc_rules is not None:
            pulumi.set(__self__, "sc_rules", sc_rules)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        The ID  of the SAE Application.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="greyTagRouteName")
    def grey_tag_route_name(self) -> pulumi.Input[str]:
        """
        The name of GreyTagRoute.
        """
        return pulumi.get(self, "grey_tag_route_name")

    @grey_tag_route_name.setter
    def grey_tag_route_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "grey_tag_route_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of GreyTagRoute.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dubboRules")
    def dubbo_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleArgs']]]]:
        """
        The grayscale rule created for Dubbo Application. See `dubbo_rules` below.
        """
        return pulumi.get(self, "dubbo_rules")

    @dubbo_rules.setter
    def dubbo_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleArgs']]]]):
        pulumi.set(self, "dubbo_rules", value)

    @property
    @pulumi.getter(name="scRules")
    def sc_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleArgs']]]]:
        """
        The grayscale rule created for SpringCloud Application. See `sc_rules` below.
        """
        return pulumi.get(self, "sc_rules")

    @sc_rules.setter
    def sc_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleArgs']]]]):
        pulumi.set(self, "sc_rules", value)


@pulumi.input_type
class _GreyTagRouteState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dubbo_rules: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleArgs']]]] = None,
                 grey_tag_route_name: Optional[pulumi.Input[str]] = None,
                 sc_rules: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleArgs']]]] = None):
        """
        Input properties used for looking up and filtering GreyTagRoute resources.
        :param pulumi.Input[str] app_id: The ID  of the SAE Application.
        :param pulumi.Input[str] description: The description of GreyTagRoute.
        :param pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleArgs']]] dubbo_rules: The grayscale rule created for Dubbo Application. See `dubbo_rules` below.
        :param pulumi.Input[str] grey_tag_route_name: The name of GreyTagRoute.
        :param pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleArgs']]] sc_rules: The grayscale rule created for SpringCloud Application. See `sc_rules` below.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dubbo_rules is not None:
            pulumi.set(__self__, "dubbo_rules", dubbo_rules)
        if grey_tag_route_name is not None:
            pulumi.set(__self__, "grey_tag_route_name", grey_tag_route_name)
        if sc_rules is not None:
            pulumi.set(__self__, "sc_rules", sc_rules)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID  of the SAE Application.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of GreyTagRoute.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dubboRules")
    def dubbo_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleArgs']]]]:
        """
        The grayscale rule created for Dubbo Application. See `dubbo_rules` below.
        """
        return pulumi.get(self, "dubbo_rules")

    @dubbo_rules.setter
    def dubbo_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteDubboRuleArgs']]]]):
        pulumi.set(self, "dubbo_rules", value)

    @property
    @pulumi.getter(name="greyTagRouteName")
    def grey_tag_route_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of GreyTagRoute.
        """
        return pulumi.get(self, "grey_tag_route_name")

    @grey_tag_route_name.setter
    def grey_tag_route_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grey_tag_route_name", value)

    @property
    @pulumi.getter(name="scRules")
    def sc_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleArgs']]]]:
        """
        The grayscale rule created for SpringCloud Application. See `sc_rules` below.
        """
        return pulumi.get(self, "sc_rules")

    @sc_rules.setter
    def sc_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GreyTagRouteScRuleArgs']]]]):
        pulumi.set(self, "sc_rules", value)


class GreyTagRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dubbo_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteDubboRuleArgs']]]]] = None,
                 grey_tag_route_name: Optional[pulumi.Input[str]] = None,
                 sc_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteScRuleArgs']]]]] = None,
                 __props__=None):
        """
        Provides a Serverless App Engine (SAE) GreyTagRoute resource.

        For information about Serverless App Engine (SAE) GreyTagRoute and how to use it, see [What is GreyTagRoute](https://www.alibabacloud.com/help/en/sae/latest/create-grey-tag-route).

        > **NOTE:** Available since v1.160.0.

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
        default_integer = random.index.Integer("default",
            max=99999,
            min=10000)
        default = alicloud.get_regions(current=True)
        default_get_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=default_network.id,
            zone_id=default_get_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default", vpc_id=default_network.id)
        default_namespace = alicloud.sae.Namespace("default",
            namespace_id=f"{default.regions[0].id}:example{default_integer['result']}",
            namespace_name=name,
            namespace_description=name,
            enable_micro_registration=False)
        default_application = alicloud.sae.Application("default",
            app_description=name,
            app_name=f"{name}-{default_integer['result']}",
            namespace_id=default_namespace.id,
            image_url=f"registry-vpc.{default.regions[0].id}.aliyuncs.com/sae-demo-image/consumer:1.0",
            package_type="Image",
            security_group_id=default_security_group.id,
            vpc_id=default_network.id,
            vswitch_id=default_switch.id,
            timezone="Asia/Beijing",
            replicas=5,
            cpu=500,
            memory=2048)
        default_grey_tag_route = alicloud.sae.GreyTagRoute("default",
            grey_tag_route_name=name,
            description=name,
            app_id=default_application.id,
            sc_rules=[alicloud.sae.GreyTagRouteScRuleArgs(
                items=[alicloud.sae.GreyTagRouteScRuleItemArgs(
                    type="param",
                    name="tfexample",
                    operator="rawvalue",
                    value="example",
                    cond="==",
                )],
                path="/tf/example",
                condition="AND",
            )],
            dubbo_rules=[alicloud.sae.GreyTagRouteDubboRuleArgs(
                items=[alicloud.sae.GreyTagRouteDubboRuleItemArgs(
                    cond="==",
                    expr=".key1",
                    index=1,
                    operator="rawvalue",
                    value="value1",
                )],
                condition="OR",
                group="DUBBO",
                method_name="example",
                service_name="com.example.service",
                version="1.0.0",
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Serverless App Engine (SAE) GreyTagRoute can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:sae/greyTagRoute:GreyTagRoute example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID  of the SAE Application.
        :param pulumi.Input[str] description: The description of GreyTagRoute.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteDubboRuleArgs']]]] dubbo_rules: The grayscale rule created for Dubbo Application. See `dubbo_rules` below.
        :param pulumi.Input[str] grey_tag_route_name: The name of GreyTagRoute.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteScRuleArgs']]]] sc_rules: The grayscale rule created for SpringCloud Application. See `sc_rules` below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GreyTagRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Serverless App Engine (SAE) GreyTagRoute resource.

        For information about Serverless App Engine (SAE) GreyTagRoute and how to use it, see [What is GreyTagRoute](https://www.alibabacloud.com/help/en/sae/latest/create-grey-tag-route).

        > **NOTE:** Available since v1.160.0.

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
        default_integer = random.index.Integer("default",
            max=99999,
            min=10000)
        default = alicloud.get_regions(current=True)
        default_get_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=default_network.id,
            zone_id=default_get_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default", vpc_id=default_network.id)
        default_namespace = alicloud.sae.Namespace("default",
            namespace_id=f"{default.regions[0].id}:example{default_integer['result']}",
            namespace_name=name,
            namespace_description=name,
            enable_micro_registration=False)
        default_application = alicloud.sae.Application("default",
            app_description=name,
            app_name=f"{name}-{default_integer['result']}",
            namespace_id=default_namespace.id,
            image_url=f"registry-vpc.{default.regions[0].id}.aliyuncs.com/sae-demo-image/consumer:1.0",
            package_type="Image",
            security_group_id=default_security_group.id,
            vpc_id=default_network.id,
            vswitch_id=default_switch.id,
            timezone="Asia/Beijing",
            replicas=5,
            cpu=500,
            memory=2048)
        default_grey_tag_route = alicloud.sae.GreyTagRoute("default",
            grey_tag_route_name=name,
            description=name,
            app_id=default_application.id,
            sc_rules=[alicloud.sae.GreyTagRouteScRuleArgs(
                items=[alicloud.sae.GreyTagRouteScRuleItemArgs(
                    type="param",
                    name="tfexample",
                    operator="rawvalue",
                    value="example",
                    cond="==",
                )],
                path="/tf/example",
                condition="AND",
            )],
            dubbo_rules=[alicloud.sae.GreyTagRouteDubboRuleArgs(
                items=[alicloud.sae.GreyTagRouteDubboRuleItemArgs(
                    cond="==",
                    expr=".key1",
                    index=1,
                    operator="rawvalue",
                    value="value1",
                )],
                condition="OR",
                group="DUBBO",
                method_name="example",
                service_name="com.example.service",
                version="1.0.0",
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Serverless App Engine (SAE) GreyTagRoute can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:sae/greyTagRoute:GreyTagRoute example <id>
        ```

        :param str resource_name: The name of the resource.
        :param GreyTagRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GreyTagRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dubbo_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteDubboRuleArgs']]]]] = None,
                 grey_tag_route_name: Optional[pulumi.Input[str]] = None,
                 sc_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteScRuleArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GreyTagRouteArgs.__new__(GreyTagRouteArgs)

            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            __props__.__dict__["description"] = description
            __props__.__dict__["dubbo_rules"] = dubbo_rules
            if grey_tag_route_name is None and not opts.urn:
                raise TypeError("Missing required property 'grey_tag_route_name'")
            __props__.__dict__["grey_tag_route_name"] = grey_tag_route_name
            __props__.__dict__["sc_rules"] = sc_rules
        super(GreyTagRoute, __self__).__init__(
            'alicloud:sae/greyTagRoute:GreyTagRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dubbo_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteDubboRuleArgs']]]]] = None,
            grey_tag_route_name: Optional[pulumi.Input[str]] = None,
            sc_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteScRuleArgs']]]]] = None) -> 'GreyTagRoute':
        """
        Get an existing GreyTagRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID  of the SAE Application.
        :param pulumi.Input[str] description: The description of GreyTagRoute.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteDubboRuleArgs']]]] dubbo_rules: The grayscale rule created for Dubbo Application. See `dubbo_rules` below.
        :param pulumi.Input[str] grey_tag_route_name: The name of GreyTagRoute.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GreyTagRouteScRuleArgs']]]] sc_rules: The grayscale rule created for SpringCloud Application. See `sc_rules` below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GreyTagRouteState.__new__(_GreyTagRouteState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["description"] = description
        __props__.__dict__["dubbo_rules"] = dubbo_rules
        __props__.__dict__["grey_tag_route_name"] = grey_tag_route_name
        __props__.__dict__["sc_rules"] = sc_rules
        return GreyTagRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        The ID  of the SAE Application.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of GreyTagRoute.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dubboRules")
    def dubbo_rules(self) -> pulumi.Output[Optional[Sequence['outputs.GreyTagRouteDubboRule']]]:
        """
        The grayscale rule created for Dubbo Application. See `dubbo_rules` below.
        """
        return pulumi.get(self, "dubbo_rules")

    @property
    @pulumi.getter(name="greyTagRouteName")
    def grey_tag_route_name(self) -> pulumi.Output[str]:
        """
        The name of GreyTagRoute.
        """
        return pulumi.get(self, "grey_tag_route_name")

    @property
    @pulumi.getter(name="scRules")
    def sc_rules(self) -> pulumi.Output[Optional[Sequence['outputs.GreyTagRouteScRule']]]:
        """
        The grayscale rule created for SpringCloud Application. See `sc_rules` below.
        """
        return pulumi.get(self, "sc_rules")

