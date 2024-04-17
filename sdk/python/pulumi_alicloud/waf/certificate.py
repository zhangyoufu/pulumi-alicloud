# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CertificateArgs', 'Certificate']

@pulumi.input_type
class CertificateArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 certificate: Optional[pulumi.Input[str]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Certificate resource.
        :param pulumi.Input[str] domain: The domain that you want to add to WAF.
        :param pulumi.Input[str] instance_id: The ID of the WAF instance.
        :param pulumi.Input[str] certificate: Certificate file content.
        :param pulumi.Input[str] certificate_id: The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
        :param pulumi.Input[str] certificate_name: Certificate file name.
        :param pulumi.Input[str] private_key: The private key.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "instance_id", instance_id)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if certificate_name is not None:
            pulumi.set(__self__, "certificate_name", certificate_name)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The domain that you want to add to WAF.
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
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate file content.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate file name.
        """
        return pulumi.get(self, "certificate_name")

    @certificate_name.setter
    def certificate_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_name", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The private key.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)


@pulumi.input_type
class _CertificateState:
    def __init__(__self__, *,
                 certificate: Optional[pulumi.Input[str]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Certificate resources.
        :param pulumi.Input[str] certificate: Certificate file content.
        :param pulumi.Input[str] certificate_id: The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
        :param pulumi.Input[str] certificate_name: Certificate file name.
        :param pulumi.Input[str] domain: The domain that you want to add to WAF.
        :param pulumi.Input[str] instance_id: The ID of the WAF instance.
        :param pulumi.Input[str] private_key: The private key.
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if certificate_name is not None:
            pulumi.set(__self__, "certificate_name", certificate_name)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate file content.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate file name.
        """
        return pulumi.get(self, "certificate_name")

    @certificate_name.setter
    def certificate_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_name", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain that you want to add to WAF.
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
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The private key.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)


class Certificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a WAF Certificate resource.

        For information about WAF Certificate and how to use it, see [What is Certificate](https://www.alibabacloud.com/help/doc-detail/28517.htm).

        > **NOTE:** Available in v1.135.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.waf.Certificate("default",
            certificate_name="your_certificate_name",
            instance_id="your_instance_id",
            domain="your_domain_name",
            private_key="your_private_key",
            certificate="your_certificate")
        default2 = alicloud.waf.Certificate("default2",
            instance_id="your_instance_id",
            domain="your_domain_name",
            certificate_id="your_certificate_id")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        WAF Certificate can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:waf/certificate:Certificate example <instance_id>:<domain>:<certificate_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: Certificate file content.
        :param pulumi.Input[str] certificate_id: The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
        :param pulumi.Input[str] certificate_name: Certificate file name.
        :param pulumi.Input[str] domain: The domain that you want to add to WAF.
        :param pulumi.Input[str] instance_id: The ID of the WAF instance.
        :param pulumi.Input[str] private_key: The private key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF Certificate resource.

        For information about WAF Certificate and how to use it, see [What is Certificate](https://www.alibabacloud.com/help/doc-detail/28517.htm).

        > **NOTE:** Available in v1.135.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.waf.Certificate("default",
            certificate_name="your_certificate_name",
            instance_id="your_instance_id",
            domain="your_domain_name",
            private_key="your_private_key",
            certificate="your_certificate")
        default2 = alicloud.waf.Certificate("default2",
            instance_id="your_instance_id",
            domain="your_domain_name",
            certificate_id="your_certificate_id")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        WAF Certificate can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:waf/certificate:Certificate example <instance_id>:<domain>:<certificate_id>
        ```

        :param str resource_name: The name of the resource.
        :param CertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CertificateArgs.__new__(CertificateArgs)

            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["certificate_id"] = certificate_id
            __props__.__dict__["certificate_name"] = certificate_name
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["private_key"] = private_key
        super(Certificate, __self__).__init__(
            'alicloud:waf/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            certificate_id: Optional[pulumi.Input[str]] = None,
            certificate_name: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: Certificate file content.
        :param pulumi.Input[str] certificate_id: The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
        :param pulumi.Input[str] certificate_name: Certificate file name.
        :param pulumi.Input[str] domain: The domain that you want to add to WAF.
        :param pulumi.Input[str] instance_id: The ID of the WAF instance.
        :param pulumi.Input[str] private_key: The private key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertificateState.__new__(_CertificateState)

        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["certificate_id"] = certificate_id
        __props__.__dict__["certificate_name"] = certificate_name
        __props__.__dict__["domain"] = domain
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["private_key"] = private_key
        return Certificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional[str]]:
        """
        Certificate file content.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> pulumi.Output[str]:
        """
        The certificate id is automatically generated when you upload your certificate content.**NOTE:** you can also use Certificate ID saved in SSL.
        """
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> pulumi.Output[str]:
        """
        Certificate file name.
        """
        return pulumi.get(self, "certificate_name")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The domain that you want to add to WAF.
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
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[Optional[str]]:
        """
        The private key.
        """
        return pulumi.get(self, "private_key")

