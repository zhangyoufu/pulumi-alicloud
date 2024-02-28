# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecurityPolicyArgs', 'SecurityPolicy']

@pulumi.input_type
class SecurityPolicyArgs:
    def __init__(__self__, *,
                 ciphers: pulumi.Input[Sequence[pulumi.Input[str]]],
                 tls_versions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_policy_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a SecurityPolicy resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ciphers: The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
               - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
               - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
               - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tls_versions: The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] security_policy_name: The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the resource.
        """
        pulumi.set(__self__, "ciphers", ciphers)
        pulumi.set(__self__, "tls_versions", tls_versions)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if security_policy_name is not None:
            pulumi.set(__self__, "security_policy_name", security_policy_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def ciphers(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
        - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
        - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
        - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`.
        """
        return pulumi.get(self, "ciphers")

    @ciphers.setter
    def ciphers(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "ciphers", value)

    @property
    @pulumi.getter(name="tlsVersions")
    def tls_versions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        """
        return pulumi.get(self, "tls_versions")

    @tls_versions.setter
    def tls_versions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "tls_versions", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="securityPolicyName")
    def security_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "security_policy_name")

    @security_policy_name.setter
    def security_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tag of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SecurityPolicyState:
    def __init__(__self__, *,
                 ciphers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_policy_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 tls_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SecurityPolicy resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ciphers: The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
               - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
               - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
               - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] security_policy_name: The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tls_versions: The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        """
        if ciphers is not None:
            pulumi.set(__self__, "ciphers", ciphers)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if security_policy_name is not None:
            pulumi.set(__self__, "security_policy_name", security_policy_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tls_versions is not None:
            pulumi.set(__self__, "tls_versions", tls_versions)

    @property
    @pulumi.getter
    def ciphers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
        - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
        - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
        - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`.
        """
        return pulumi.get(self, "ciphers")

    @ciphers.setter
    def ciphers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ciphers", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

    @property
    @pulumi.getter(name="securityPolicyName")
    def security_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "security_policy_name")

    @security_policy_name.setter
    def security_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy_name", value)

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

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The tag of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tlsVersions")
    def tls_versions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        """
        return pulumi.get(self, "tls_versions")

    @tls_versions.setter
    def tls_versions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tls_versions", value)


class SecurityPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ciphers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_policy_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 tls_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a NLB Security Policy resource.

        For information about NLB Security Policy and how to use it, see [What is Security Policy](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createsecuritypolicy-nlb).

        > **NOTE:** Available since v1.187.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_security_policy = alicloud.nlb.SecurityPolicy("defaultSecurityPolicy",
            resource_group_id=default_resource_groups.ids[0],
            security_policy_name=name,
            ciphers=[
                "ECDHE-RSA-AES128-SHA",
                "ECDHE-ECDSA-AES128-SHA",
            ],
            tls_versions=[
                "TLSv1.0",
                "TLSv1.1",
                "TLSv1.2",
            ],
            tags={
                "Created": "TF",
                "For": "example",
            })
        ```

        ## Import

        NLB Security Policy can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nlb/securityPolicy:SecurityPolicy example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ciphers: The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
               - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
               - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
               - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] security_policy_name: The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tls_versions: The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a NLB Security Policy resource.

        For information about NLB Security Policy and how to use it, see [What is Security Policy](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createsecuritypolicy-nlb).

        > **NOTE:** Available since v1.187.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        default_resource_groups = alicloud.resourcemanager.get_resource_groups()
        default_security_policy = alicloud.nlb.SecurityPolicy("defaultSecurityPolicy",
            resource_group_id=default_resource_groups.ids[0],
            security_policy_name=name,
            ciphers=[
                "ECDHE-RSA-AES128-SHA",
                "ECDHE-ECDSA-AES128-SHA",
            ],
            tls_versions=[
                "TLSv1.0",
                "TLSv1.1",
                "TLSv1.2",
            ],
            tags={
                "Created": "TF",
                "For": "example",
            })
        ```

        ## Import

        NLB Security Policy can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:nlb/securityPolicy:SecurityPolicy example <id>
        ```

        :param str resource_name: The name of the resource.
        :param SecurityPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ciphers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 security_policy_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 tls_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityPolicyArgs.__new__(SecurityPolicyArgs)

            if ciphers is None and not opts.urn:
                raise TypeError("Missing required property 'ciphers'")
            __props__.__dict__["ciphers"] = ciphers
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["security_policy_name"] = security_policy_name
            __props__.__dict__["tags"] = tags
            if tls_versions is None and not opts.urn:
                raise TypeError("Missing required property 'tls_versions'")
            __props__.__dict__["tls_versions"] = tls_versions
            __props__.__dict__["status"] = None
        super(SecurityPolicy, __self__).__init__(
            'alicloud:nlb/securityPolicy:SecurityPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ciphers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            security_policy_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            tls_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SecurityPolicy':
        """
        Get an existing SecurityPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ciphers: The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
               - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
               - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
               - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`.
        :param pulumi.Input[str] resource_group_id: The ID of the resource group.
        :param pulumi.Input[str] security_policy_name: The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        :param pulumi.Input[str] status: The status of the resource.
        :param pulumi.Input[Mapping[str, Any]] tags: The tag of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tls_versions: The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityPolicyState.__new__(_SecurityPolicyState)

        __props__.__dict__["ciphers"] = ciphers
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["security_policy_name"] = security_policy_name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tls_versions"] = tls_versions
        return SecurityPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ciphers(self) -> pulumi.Output[Sequence[str]]:
        """
        The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
        - TLS 1.0 and TLS 1.1 support the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA`, `DES-CBC3-SHA`
        - TLS 1.2 supports the following cipher suites: `ECDHE-ECDSA-AES128-SHA`, `ECDHE-ECDSA-AES256-SHA`, `ECDHE-RSA-AES128-SHA`, `ECDHE-RSA-AES256-SHA`, `AES128-SHA`, `AES256-SHA, DES-CBC3-SHA`, `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES256-SHA384`, `AES128-GCM-SHA256`, `AES256-GCM-SHA384`, `AES128-SHA256`, `AES256-SHA256`
        - TLS 1.3 supports the following cipher suites: `TLS_AES_128_GCM_SHA256`, `TLS_AES_256_GCM_SHA384`, `TLS_CHACHA20_POLY1305_SHA256`, `TLS_AES_128_CCM_SHA256`, `TLS_AES_128_CCM_8_SHA256`.
        """
        return pulumi.get(self, "ciphers")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="securityPolicyName")
    def security_policy_name(self) -> pulumi.Output[str]:
        """
        The name of the security policy. The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
        """
        return pulumi.get(self, "security_policy_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The tag of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tlsVersions")
    def tls_versions(self) -> pulumi.Output[Sequence[str]]:
        """
        The supported versions of the Transport Layer Security (TLS) protocol. Valid values: `TLSv1.0`, `TLSv1.1`, `TLSv1.2`, and `TLSv1.3`. You can specify at most four TLS versions.
        """
        return pulumi.get(self, "tls_versions")

