# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WafDomainArgs', 'WafDomain']

@pulumi.input_type
class WafDomainArgs:
    def __init__(__self__, *,
                 domain_name: pulumi.Input[str],
                 client_ip_tag: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WafDomain resource.
        :param pulumi.Input[str] domain_name: The accelerated domain name.
        :param pulumi.Input[str] client_ip_tag: The client ip tag.
        """
        pulumi.set(__self__, "domain_name", domain_name)
        if client_ip_tag is not None:
            pulumi.set(__self__, "client_ip_tag", client_ip_tag)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        The accelerated domain name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="clientIpTag")
    def client_ip_tag(self) -> Optional[pulumi.Input[str]]:
        """
        The client ip tag.
        """
        return pulumi.get(self, "client_ip_tag")

    @client_ip_tag.setter
    def client_ip_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_ip_tag", value)


@pulumi.input_type
class _WafDomainState:
    def __init__(__self__, *,
                 client_ip_tag: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WafDomain resources.
        :param pulumi.Input[str] client_ip_tag: The client ip tag.
        :param pulumi.Input[str] domain_name: The accelerated domain name.
        """
        if client_ip_tag is not None:
            pulumi.set(__self__, "client_ip_tag", client_ip_tag)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)

    @property
    @pulumi.getter(name="clientIpTag")
    def client_ip_tag(self) -> Optional[pulumi.Input[str]]:
        """
        The client ip tag.
        """
        return pulumi.get(self, "client_ip_tag")

    @client_ip_tag.setter
    def client_ip_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_ip_tag", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The accelerated domain name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)


class WafDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_ip_tag: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a DCDN Waf Domain resource.

        For information about DCDN Waf Domain and how to use it, see [What is Waf Domain](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchsetdcdnwafdomainconfigs).

        > **NOTE:** Available since v1.185.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        domain_name = config.get("domainName")
        if domain_name is None:
            domain_name = "tf-example.com"
        default = random.RandomInteger("default",
            min=10000,
            max=99999)
        example_domain = alicloud.dcdn.Domain("exampleDomain",
            domain_name=default.result.apply(lambda result: f"{domain_name}-{result}"),
            scope="overseas",
            sources=[alicloud.dcdn.DomainSourceArgs(
                content="1.1.1.1",
                port=80,
                priority="20",
                type="ipaddr",
                weight="10",
            )])
        example_waf_domain = alicloud.dcdn.WafDomain("exampleWafDomain",
            domain_name=example_domain.domain_name,
            client_ip_tag="X-Forwarded-For")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        DCDN Waf Domain can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dcdn/wafDomain:WafDomain example <domain_name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_ip_tag: The client ip tag.
        :param pulumi.Input[str] domain_name: The accelerated domain name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WafDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DCDN Waf Domain resource.

        For information about DCDN Waf Domain and how to use it, see [What is Waf Domain](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchsetdcdnwafdomainconfigs).

        > **NOTE:** Available since v1.185.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        domain_name = config.get("domainName")
        if domain_name is None:
            domain_name = "tf-example.com"
        default = random.RandomInteger("default",
            min=10000,
            max=99999)
        example_domain = alicloud.dcdn.Domain("exampleDomain",
            domain_name=default.result.apply(lambda result: f"{domain_name}-{result}"),
            scope="overseas",
            sources=[alicloud.dcdn.DomainSourceArgs(
                content="1.1.1.1",
                port=80,
                priority="20",
                type="ipaddr",
                weight="10",
            )])
        example_waf_domain = alicloud.dcdn.WafDomain("exampleWafDomain",
            domain_name=example_domain.domain_name,
            client_ip_tag="X-Forwarded-For")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        DCDN Waf Domain can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:dcdn/wafDomain:WafDomain example <domain_name>
        ```

        :param str resource_name: The name of the resource.
        :param WafDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WafDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_ip_tag: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WafDomainArgs.__new__(WafDomainArgs)

            __props__.__dict__["client_ip_tag"] = client_ip_tag
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
        super(WafDomain, __self__).__init__(
            'alicloud:dcdn/wafDomain:WafDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_ip_tag: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None) -> 'WafDomain':
        """
        Get an existing WafDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_ip_tag: The client ip tag.
        :param pulumi.Input[str] domain_name: The accelerated domain name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WafDomainState.__new__(_WafDomainState)

        __props__.__dict__["client_ip_tag"] = client_ip_tag
        __props__.__dict__["domain_name"] = domain_name
        return WafDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientIpTag")
    def client_ip_tag(self) -> pulumi.Output[Optional[str]]:
        """
        The client ip tag.
        """
        return pulumi.get(self, "client_ip_tag")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        The accelerated domain name.
        """
        return pulumi.get(self, "domain_name")

