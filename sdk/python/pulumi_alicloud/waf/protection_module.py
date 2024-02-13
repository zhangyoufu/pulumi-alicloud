# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProtectionModuleArgs', 'ProtectionModule']

@pulumi.input_type
class ProtectionModuleArgs:
    def __init__(__self__, *,
                 defense_type: pulumi.Input[str],
                 domain: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 mode: pulumi.Input[int],
                 status: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ProtectionModule resource.
        :param pulumi.Input[str] defense_type: The Protection Module. Valid values: `ac_cc`, `antifraud`, `dld`, `normalized`, `waf`.
        :param pulumi.Input[str] domain: The domain name that is added to WAF.
        :param pulumi.Input[str] instance_id: The ID of the WAF instance.
        :param pulumi.Input[int] mode: The protection mode of the specified protection module. **NOTE:** The value of the Mode parameter varies based on the value of the `defense_type` parameter.
               * The `defense_type` is `waf`. `0`: block mode. `1`: warn mode.
               * The `defense_type` is `dld`. `0`: warn mode. `1`: block mode.
               * The `defense_type` is `ac_cc`. `0`: prevention mode. `1`: protection-emergency mode.
               * The `defense_type` is `antifraud`. `0`: warn mode. `1`: block mode. `2`: strict interception mode.
               * The `defense_type` is `normalized`. `0`: warn mode. `1`: block mode.
        :param pulumi.Input[int] status: The status of the resource. Valid values: `0`, `1`.
        """
        pulumi.set(__self__, "defense_type", defense_type)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "mode", mode)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="defenseType")
    def defense_type(self) -> pulumi.Input[str]:
        """
        The Protection Module. Valid values: `ac_cc`, `antifraud`, `dld`, `normalized`, `waf`.
        """
        return pulumi.get(self, "defense_type")

    @defense_type.setter
    def defense_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "defense_type", value)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The domain name that is added to WAF.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the WAF instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[int]:
        """
        The protection mode of the specified protection module. **NOTE:** The value of the Mode parameter varies based on the value of the `defense_type` parameter.
        * The `defense_type` is `waf`. `0`: block mode. `1`: warn mode.
        * The `defense_type` is `dld`. `0`: warn mode. `1`: block mode.
        * The `defense_type` is `ac_cc`. `0`: prevention mode. `1`: protection-emergency mode.
        * The `defense_type` is `antifraud`. `0`: warn mode. `1`: block mode. `2`: strict interception mode.
        * The `defense_type` is `normalized`. `0`: warn mode. `1`: block mode.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[int]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of the resource. Valid values: `0`, `1`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _ProtectionModuleState:
    def __init__(__self__, *,
                 defense_type: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ProtectionModule resources.
        :param pulumi.Input[str] defense_type: The Protection Module. Valid values: `ac_cc`, `antifraud`, `dld`, `normalized`, `waf`.
        :param pulumi.Input[str] domain: The domain name that is added to WAF.
        :param pulumi.Input[str] instance_id: The ID of the WAF instance.
        :param pulumi.Input[int] mode: The protection mode of the specified protection module. **NOTE:** The value of the Mode parameter varies based on the value of the `defense_type` parameter.
               * The `defense_type` is `waf`. `0`: block mode. `1`: warn mode.
               * The `defense_type` is `dld`. `0`: warn mode. `1`: block mode.
               * The `defense_type` is `ac_cc`. `0`: prevention mode. `1`: protection-emergency mode.
               * The `defense_type` is `antifraud`. `0`: warn mode. `1`: block mode. `2`: strict interception mode.
               * The `defense_type` is `normalized`. `0`: warn mode. `1`: block mode.
        :param pulumi.Input[int] status: The status of the resource. Valid values: `0`, `1`.
        """
        if defense_type is not None:
            pulumi.set(__self__, "defense_type", defense_type)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="defenseType")
    def defense_type(self) -> Optional[pulumi.Input[str]]:
        """
        The Protection Module. Valid values: `ac_cc`, `antifraud`, `dld`, `normalized`, `waf`.
        """
        return pulumi.get(self, "defense_type")

    @defense_type.setter
    def defense_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "defense_type", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain name that is added to WAF.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the WAF instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[int]]:
        """
        The protection mode of the specified protection module. **NOTE:** The value of the Mode parameter varies based on the value of the `defense_type` parameter.
        * The `defense_type` is `waf`. `0`: block mode. `1`: warn mode.
        * The `defense_type` is `dld`. `0`: warn mode. `1`: block mode.
        * The `defense_type` is `ac_cc`. `0`: prevention mode. `1`: protection-emergency mode.
        * The `defense_type` is `antifraud`. `0`: warn mode. `1`: block mode. `2`: strict interception mode.
        * The `defense_type` is `normalized`. `0`: warn mode. `1`: block mode.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of the resource. Valid values: `0`, `1`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)


class ProtectionModule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 defense_type: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a Web Application Firewall(WAF) Protection Module resource.

        For information about Web Application Firewall(WAF) Protection Module and how to use it, see [What is Protection Module](https://www.alibabacloud.com/help/en/doc-detail/160775.htm).

        > **NOTE:** Available in v1.141.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_instances = alicloud.waf.get_instances()
        default_domain = alicloud.waf.Domain("defaultDomain",
            domain_name="you domain",
            instance_id=default_instances.ids[0],
            is_access_product="On",
            source_ips=["1.1.1.1"],
            cluster_type="PhysicalCluster",
            http2_ports=["443"],
            http_ports=["80"],
            https_ports=["443"],
            http_to_user_ip="Off",
            https_redirect="Off",
            load_balancing="IpHash",
            log_headers=[alicloud.waf.DomainLogHeaderArgs(
                key="foo",
                value="http",
            )])
        default_protection_module = alicloud.waf.ProtectionModule("defaultProtectionModule",
            instance_id=default_instances.ids[0],
            domain=default_domain.domain_name,
            defense_type="ac_cc",
            mode=0,
            status=0)
        ```

        ## Import

        Web Application Firewall(WAF) Protection Module can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:waf/protectionModule:ProtectionModule example <instance_id>:<domain>:<defense_type>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defense_type: The Protection Module. Valid values: `ac_cc`, `antifraud`, `dld`, `normalized`, `waf`.
        :param pulumi.Input[str] domain: The domain name that is added to WAF.
        :param pulumi.Input[str] instance_id: The ID of the WAF instance.
        :param pulumi.Input[int] mode: The protection mode of the specified protection module. **NOTE:** The value of the Mode parameter varies based on the value of the `defense_type` parameter.
               * The `defense_type` is `waf`. `0`: block mode. `1`: warn mode.
               * The `defense_type` is `dld`. `0`: warn mode. `1`: block mode.
               * The `defense_type` is `ac_cc`. `0`: prevention mode. `1`: protection-emergency mode.
               * The `defense_type` is `antifraud`. `0`: warn mode. `1`: block mode. `2`: strict interception mode.
               * The `defense_type` is `normalized`. `0`: warn mode. `1`: block mode.
        :param pulumi.Input[int] status: The status of the resource. Valid values: `0`, `1`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProtectionModuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Web Application Firewall(WAF) Protection Module resource.

        For information about Web Application Firewall(WAF) Protection Module and how to use it, see [What is Protection Module](https://www.alibabacloud.com/help/en/doc-detail/160775.htm).

        > **NOTE:** Available in v1.141.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_instances = alicloud.waf.get_instances()
        default_domain = alicloud.waf.Domain("defaultDomain",
            domain_name="you domain",
            instance_id=default_instances.ids[0],
            is_access_product="On",
            source_ips=["1.1.1.1"],
            cluster_type="PhysicalCluster",
            http2_ports=["443"],
            http_ports=["80"],
            https_ports=["443"],
            http_to_user_ip="Off",
            https_redirect="Off",
            load_balancing="IpHash",
            log_headers=[alicloud.waf.DomainLogHeaderArgs(
                key="foo",
                value="http",
            )])
        default_protection_module = alicloud.waf.ProtectionModule("defaultProtectionModule",
            instance_id=default_instances.ids[0],
            domain=default_domain.domain_name,
            defense_type="ac_cc",
            mode=0,
            status=0)
        ```

        ## Import

        Web Application Firewall(WAF) Protection Module can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:waf/protectionModule:ProtectionModule example <instance_id>:<domain>:<defense_type>
        ```

        :param str resource_name: The name of the resource.
        :param ProtectionModuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProtectionModuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 defense_type: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProtectionModuleArgs.__new__(ProtectionModuleArgs)

            if defense_type is None and not opts.urn:
                raise TypeError("Missing required property 'defense_type'")
            __props__.__dict__["defense_type"] = defense_type
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if mode is None and not opts.urn:
                raise TypeError("Missing required property 'mode'")
            __props__.__dict__["mode"] = mode
            __props__.__dict__["status"] = status
        super(ProtectionModule, __self__).__init__(
            'alicloud:waf/protectionModule:ProtectionModule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            defense_type: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[int]] = None) -> 'ProtectionModule':
        """
        Get an existing ProtectionModule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defense_type: The Protection Module. Valid values: `ac_cc`, `antifraud`, `dld`, `normalized`, `waf`.
        :param pulumi.Input[str] domain: The domain name that is added to WAF.
        :param pulumi.Input[str] instance_id: The ID of the WAF instance.
        :param pulumi.Input[int] mode: The protection mode of the specified protection module. **NOTE:** The value of the Mode parameter varies based on the value of the `defense_type` parameter.
               * The `defense_type` is `waf`. `0`: block mode. `1`: warn mode.
               * The `defense_type` is `dld`. `0`: warn mode. `1`: block mode.
               * The `defense_type` is `ac_cc`. `0`: prevention mode. `1`: protection-emergency mode.
               * The `defense_type` is `antifraud`. `0`: warn mode. `1`: block mode. `2`: strict interception mode.
               * The `defense_type` is `normalized`. `0`: warn mode. `1`: block mode.
        :param pulumi.Input[int] status: The status of the resource. Valid values: `0`, `1`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProtectionModuleState.__new__(_ProtectionModuleState)

        __props__.__dict__["defense_type"] = defense_type
        __props__.__dict__["domain"] = domain
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["mode"] = mode
        __props__.__dict__["status"] = status
        return ProtectionModule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defenseType")
    def defense_type(self) -> pulumi.Output[str]:
        """
        The Protection Module. Valid values: `ac_cc`, `antifraud`, `dld`, `normalized`, `waf`.
        """
        return pulumi.get(self, "defense_type")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The domain name that is added to WAF.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the WAF instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[int]:
        """
        The protection mode of the specified protection module. **NOTE:** The value of the Mode parameter varies based on the value of the `defense_type` parameter.
        * The `defense_type` is `waf`. `0`: block mode. `1`: warn mode.
        * The `defense_type` is `dld`. `0`: warn mode. `1`: block mode.
        * The `defense_type` is `ac_cc`. `0`: prevention mode. `1`: protection-emergency mode.
        * The `defense_type` is `antifraud`. `0`: warn mode. `1`: block mode. `2`: strict interception mode.
        * The `defense_type` is `normalized`. `0`: warn mode. `1`: block mode.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[int]]:
        """
        The status of the resource. Valid values: `0`, `1`.
        """
        return pulumi.get(self, "status")

