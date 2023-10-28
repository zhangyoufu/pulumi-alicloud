# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DomainExtensionArgs', 'DomainExtension']

@pulumi.input_type
class DomainExtensionArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 frontend_port: pulumi.Input[int],
                 load_balancer_id: pulumi.Input[str],
                 server_certificate_id: pulumi.Input[str],
                 delete_protection_validation: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a DomainExtension resource.
        :param pulumi.Input[str] domain: The domain name.
        :param pulumi.Input[int] frontend_port: The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        :param pulumi.Input[str] load_balancer_id: The ID of the SLB instance.
        :param pulumi.Input[str] server_certificate_id: The ID of the certificate used by the domain name.
        :param pulumi.Input[bool] delete_protection_validation: Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "frontend_port", frontend_port)
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        pulumi.set(__self__, "server_certificate_id", server_certificate_id)
        if delete_protection_validation is not None:
            pulumi.set(__self__, "delete_protection_validation", delete_protection_validation)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The domain name.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> pulumi.Input[int]:
        """
        The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        """
        return pulumi.get(self, "frontend_port")

    @frontend_port.setter
    def frontend_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "frontend_port", value)

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> pulumi.Input[str]:
        """
        The ID of the SLB instance.
        """
        return pulumi.get(self, "load_balancer_id")

    @load_balancer_id.setter
    def load_balancer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "load_balancer_id", value)

    @property
    @pulumi.getter(name="serverCertificateId")
    def server_certificate_id(self) -> pulumi.Input[str]:
        """
        The ID of the certificate used by the domain name.
        """
        return pulumi.get(self, "server_certificate_id")

    @server_certificate_id.setter
    def server_certificate_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_certificate_id", value)

    @property
    @pulumi.getter(name="deleteProtectionValidation")
    def delete_protection_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        """
        return pulumi.get(self, "delete_protection_validation")

    @delete_protection_validation.setter
    def delete_protection_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_protection_validation", value)


@pulumi.input_type
class _DomainExtensionState:
    def __init__(__self__, *,
                 delete_protection_validation: Optional[pulumi.Input[bool]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 frontend_port: Optional[pulumi.Input[int]] = None,
                 load_balancer_id: Optional[pulumi.Input[str]] = None,
                 server_certificate_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DomainExtension resources.
        :param pulumi.Input[bool] delete_protection_validation: Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        :param pulumi.Input[str] domain: The domain name.
        :param pulumi.Input[int] frontend_port: The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        :param pulumi.Input[str] load_balancer_id: The ID of the SLB instance.
        :param pulumi.Input[str] server_certificate_id: The ID of the certificate used by the domain name.
        """
        if delete_protection_validation is not None:
            pulumi.set(__self__, "delete_protection_validation", delete_protection_validation)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if frontend_port is not None:
            pulumi.set(__self__, "frontend_port", frontend_port)
        if load_balancer_id is not None:
            pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if server_certificate_id is not None:
            pulumi.set(__self__, "server_certificate_id", server_certificate_id)

    @property
    @pulumi.getter(name="deleteProtectionValidation")
    def delete_protection_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        """
        return pulumi.get(self, "delete_protection_validation")

    @delete_protection_validation.setter
    def delete_protection_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_protection_validation", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain name.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> Optional[pulumi.Input[int]]:
        """
        The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        """
        return pulumi.get(self, "frontend_port")

    @frontend_port.setter
    def frontend_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "frontend_port", value)

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the SLB instance.
        """
        return pulumi.get(self, "load_balancer_id")

    @load_balancer_id.setter
    def load_balancer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_balancer_id", value)

    @property
    @pulumi.getter(name="serverCertificateId")
    def server_certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the certificate used by the domain name.
        """
        return pulumi.get(self, "server_certificate_id")

    @server_certificate_id.setter
    def server_certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_certificate_id", value)


class DomainExtension(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_protection_validation: Optional[pulumi.Input[bool]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 frontend_port: Optional[pulumi.Input[int]] = None,
                 load_balancer_id: Optional[pulumi.Input[str]] = None,
                 server_certificate_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        HTTPS listeners of guaranteed-performance SLB support configuring multiple certificates, allowing you to forward requests with different domain names to different backend servers.
        Please refer to the [documentation](https://www.alibabacloud.com/help/doc-detail/85956.htm?spm=a2c63.p38356.b99.40.1c881563Co8p6w) for details.

        > **NOTE:** Available in 1.60.0+

        > **NOTE:** The instance with shared loadBalancerSpec doesn't support domainExtension.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        slb_domain_extension_name = config.get("slbDomainExtensionName")
        if slb_domain_extension_name is None:
            slb_domain_extension_name = "forDomainExtension"
        domain_extension_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        domain_extension_network = alicloud.vpc.Network("domainExtensionNetwork", vpc_name=slb_domain_extension_name)
        domain_extension_switch = alicloud.vpc.Switch("domainExtensionSwitch",
            vpc_id=domain_extension_network.id,
            cidr_block="172.16.0.0/21",
            zone_id=domain_extension_zones.zones[0].id,
            vswitch_name=slb_domain_extension_name)
        instance = alicloud.slb.ApplicationLoadBalancer("instance",
            load_balancer_name=slb_domain_extension_name,
            address_type="intranet",
            load_balancer_spec="slb.s2.small",
            vswitch_id=domain_extension_switch.id)
        domain_extension_server_certificate = alicloud.slb.ServerCertificate("domainExtensionServerCertificate",
            server_certificate=\"\"\"-----BEGIN CERTIFICATE-----
        MIIDdjCCAl4CCQCcm+erkcKN7DANBgkqhkiG9w0BAQsFADB9MQswCQYDVQQGEwJj
        bjELMAkGA1UECAwCYmoxEDAOBgNVBAcMB2JlaWppbmcxDzANBgNVBAoMBmFsaXl1
        bjELMAkGA1UECwwCc2MxFTATBgNVBAMMDHd3dy50ZXN0LmNvbTEaMBgGCSqGSIb3
        DQEJARYLMTIzQDEyMy5jb20wHhcNMTkwNDI2MDM0ODAxWhcNMjQwNDI1MDM0ODAx
        WjB9MQswCQYDVQQGEwJjbjELMAkGA1UECAwCYmoxEDAOBgNVBAcMB2JlaWppbmcx
        DzANBgNVBAoMBmFsaXl1bjELMAkGA1UECwwCc2MxFTATBgNVBAMMDHd3dy50ZXN0
        LmNvbTEaMBgGCSqGSIb3DQEJARYLMTIzQDEyMy5jb20wggEiMA0GCSqGSIb3DQEB
        AQUAA4IBDwAwggEKAoIBAQDKMKF5qmN/uoMjdH3D8aPRcUOA0s8rZpYhG8zbkF1j
        8gHYoB/FDvM7G7dfVsyjbMwLOxKvAhWvHHSpEz/t7gB+QdwrAMiMJwGmtCnXrh2E
        WiXgalMe1y4a/T5R7q+m4T1zFATf+kbnHWfkSGF4W7b6UBoaH+9StQ95CnqzNf/2
        p/Of7+S0XzCxFXw8GIVzZk0xFe6lHJzaq06f3mvzrD+4rpO56tTUvrgTY/n61gsF
        ZP7f0CJ2JQh6eNRFOEUSfxKu/Dy/+IsQxorCJY2Q59ZAf3rXrqDN104jw9PlwnLl
        qfZz3RMODN6BWjxE8rvRtT0qMfuAfv1gjBdWZN0hUYBRAgMBAAEwDQYJKoZIhvcN
        AQELBQADggEBAABzo82TxGp5poVkd5pLWj5ACgcBv8Cs6oH9D+4Jz9BmyuBUsQXh
        2aG0hQAe1mU61C9konsl/GTW8umJQ4M4lYEztXXwMf5PlBMGwebM0ZbSGg6jKtZg
        WCgJ3eP/FMmyXGL5Jji5+e09eObhUDVle4tdi0On97zBoz85W02rgWFAqZJwiEAP
        t+c7jX7uOSBq2/38iGStlrX5yB1at/gJXXiA5CL5OtlR3Okvb0/QH37efO1Nu39m
        lFi0ODPAVyXjVypAiLguDxPn6AtDTdk9Iw9B19OD4NrzNRWgSSX5vuxo/VcRcgWk
        3gEe9Ca0ZKN20q9XgthAiFFjl1S9ZgdA6Zc=
        -----END CERTIFICATE-----\"\"\",
            private_key=\"\"\"-----BEGIN RSA PRIVATE KEY-----
        MIIEowIBAAKCAQEAyjCheapjf7qDI3R9w/Gj0XFDgNLPK2aWIRvM25BdY/IB2KAf
        xQ7zOxu3X1bMo2zMCzsSrwIVrxx0qRM/7e4AfkHcKwDIjCcBprQp164dhFol4GpT
        HtcuGv0+Ue6vpuE9cxQE3/pG5x1n5EhheFu2+lAaGh/vUrUPeQp6szX/9qfzn+/k
        tF8wsRV8PBiFc2ZNMRXupRyc2qtOn95r86w/uK6TuerU1L64E2P5+tYLBWT+39Ai
        diUIenjURThFEn8Srvw8v/iLEMaKwiWNkOfWQH96166gzddOI8PT5cJy5an2c90T
        DgzegVo8RPK70bU9KjH7gH79YIwXVmTdIVGAUQIDAQABAoIBAE1J4a/8biR5S3/W
        G+03BYQeY8tuyjqw8FqfoeOcf9agwAvqybouSNQjeCk9qOQfxq/UWQQFK/zQR9gJ
        v7pX7GBXFK5rkj3g+0SaQhRsPmRFgY0Tl8qGPt2aSKRRNVv5ZeADmwlzRn86QmiF
        Mp0rkfqFfDTYWEepZszCML0ouzuxsW/9tq7rvtSjsgATNt31B3vFa3D3JBi31jUh
        5nfR9A3bATze7mQw3byEDiVl5ASRDgYyur403P1fDnMy9DBHZ8NaPOsFF6OBpJal
        BJsG5z00hll5PFN2jfmBQKlvAeU7wfwqdaSnGHOfqf2DeTTaFjIQ4gUhRn/m6pLo
        6kXttLECgYEA9sng0Qz/TcPFfM4tQ1gyvB1cKnnGIwg1FP8sfUjbbEgjaHhA224S
        k3BxtX2Kq6fhTXuwusAFc6OVMAZ76FgrQ5K4Ci7+DTsrF28z4b8td+p+lO/DxgP9
        lTgN+ddsiTOV4fUef9Z3yY0Zr0CnBUMbQYRaV2UIbCdiB0G4V/bt9TsCgYEA0bya
        Oo9wGI0RJV0bYP7qwO74Ra1/i1viWbRlS7jU37Q+AZstrlKcQ5CTPzOjKFKMiUzl
        4miWacZ0/q2n+Mvd7NbXGXTLijahnyOYKaHJYyh4oBymfkgAifRstE0Ki9gdvArb
        /I+emC0GvLSyfGN8UUeDJs4NmqdEXGqjo2JOV+MCgYALFv1MR5o9Y1u/hQBRs2fs
        PiGDIx+9OUQxYloccyaxEfjNXAIGGkcpavchIbgWiJ++PJ2vdquIC8TLeK8evL+M
        9M3iX0Q5UfxYvD2HmnCvn9D6Xl/cyRcfGnq+TGjrLW9BzSMGuZt+aiHKV0xqFx7l
        bc4leTvMqGRmURS4lzcQOwKBgQCDzA/i4sYfN25h21tcHXSpnsG3D2rJyQi5NCo/
        ZjunA92/JqOTGuiFcLGHEszhhtY3ZXJET1LNz18vtzKJnpqrvOnYXlOVW/U+SqDQ
        8JDb1c/PVZGuY1KrXkR9HLiW3kz5IJ3S3PFdUVYdeTN8BQxXCyg4V12nJJtJs912
        y0zN3wKBgGDS6YttCN6aI4EOABYE8fI1EYQ7vhfiYsaWGWSR1l6bQey7KR6M1ACz
        ZzMASNyytVt12yXE4/Emv6/pYqigbDLfL1zQJSLJ3EHJYTh2RxjR+AaGDudYFG/T
        liQ9YXhV5Iu2x1pNwrtFnssDdaaGpfA7l3xC00BL7Z+SAJyI4QKA
        -----END RSA PRIVATE KEY-----\"\"\")
        https = alicloud.slb.Listener("https",
            load_balancer_id=instance.id,
            backend_port=80,
            frontend_port=443,
            protocol="https",
            sticky_session="on",
            sticky_session_type="insert",
            cookie="testslblistenercookie",
            cookie_timeout=86400,
            health_check="on",
            health_check_uri="/cons",
            health_check_connect_port=20,
            healthy_threshold=8,
            unhealthy_threshold=8,
            health_check_timeout=8,
            health_check_interval=5,
            health_check_http_code="http_2xx,http_3xx",
            bandwidth=10,
            server_certificate_id=domain_extension_server_certificate.id)
        example1 = alicloud.slb.DomainExtension("example1",
            load_balancer_id=instance.id,
            frontend_port=https.frontend_port,
            domain="www.test.com",
            server_certificate_id=domain_extension_server_certificate.id)
        ```

        ## Import

        Load balancer domain_extension can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:slb/domainExtension:DomainExtension example de-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_protection_validation: Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        :param pulumi.Input[str] domain: The domain name.
        :param pulumi.Input[int] frontend_port: The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        :param pulumi.Input[str] load_balancer_id: The ID of the SLB instance.
        :param pulumi.Input[str] server_certificate_id: The ID of the certificate used by the domain name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainExtensionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        HTTPS listeners of guaranteed-performance SLB support configuring multiple certificates, allowing you to forward requests with different domain names to different backend servers.
        Please refer to the [documentation](https://www.alibabacloud.com/help/doc-detail/85956.htm?spm=a2c63.p38356.b99.40.1c881563Co8p6w) for details.

        > **NOTE:** Available in 1.60.0+

        > **NOTE:** The instance with shared loadBalancerSpec doesn't support domainExtension.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        slb_domain_extension_name = config.get("slbDomainExtensionName")
        if slb_domain_extension_name is None:
            slb_domain_extension_name = "forDomainExtension"
        domain_extension_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        domain_extension_network = alicloud.vpc.Network("domainExtensionNetwork", vpc_name=slb_domain_extension_name)
        domain_extension_switch = alicloud.vpc.Switch("domainExtensionSwitch",
            vpc_id=domain_extension_network.id,
            cidr_block="172.16.0.0/21",
            zone_id=domain_extension_zones.zones[0].id,
            vswitch_name=slb_domain_extension_name)
        instance = alicloud.slb.ApplicationLoadBalancer("instance",
            load_balancer_name=slb_domain_extension_name,
            address_type="intranet",
            load_balancer_spec="slb.s2.small",
            vswitch_id=domain_extension_switch.id)
        domain_extension_server_certificate = alicloud.slb.ServerCertificate("domainExtensionServerCertificate",
            server_certificate=\"\"\"-----BEGIN CERTIFICATE-----
        MIIDdjCCAl4CCQCcm+erkcKN7DANBgkqhkiG9w0BAQsFADB9MQswCQYDVQQGEwJj
        bjELMAkGA1UECAwCYmoxEDAOBgNVBAcMB2JlaWppbmcxDzANBgNVBAoMBmFsaXl1
        bjELMAkGA1UECwwCc2MxFTATBgNVBAMMDHd3dy50ZXN0LmNvbTEaMBgGCSqGSIb3
        DQEJARYLMTIzQDEyMy5jb20wHhcNMTkwNDI2MDM0ODAxWhcNMjQwNDI1MDM0ODAx
        WjB9MQswCQYDVQQGEwJjbjELMAkGA1UECAwCYmoxEDAOBgNVBAcMB2JlaWppbmcx
        DzANBgNVBAoMBmFsaXl1bjELMAkGA1UECwwCc2MxFTATBgNVBAMMDHd3dy50ZXN0
        LmNvbTEaMBgGCSqGSIb3DQEJARYLMTIzQDEyMy5jb20wggEiMA0GCSqGSIb3DQEB
        AQUAA4IBDwAwggEKAoIBAQDKMKF5qmN/uoMjdH3D8aPRcUOA0s8rZpYhG8zbkF1j
        8gHYoB/FDvM7G7dfVsyjbMwLOxKvAhWvHHSpEz/t7gB+QdwrAMiMJwGmtCnXrh2E
        WiXgalMe1y4a/T5R7q+m4T1zFATf+kbnHWfkSGF4W7b6UBoaH+9StQ95CnqzNf/2
        p/Of7+S0XzCxFXw8GIVzZk0xFe6lHJzaq06f3mvzrD+4rpO56tTUvrgTY/n61gsF
        ZP7f0CJ2JQh6eNRFOEUSfxKu/Dy/+IsQxorCJY2Q59ZAf3rXrqDN104jw9PlwnLl
        qfZz3RMODN6BWjxE8rvRtT0qMfuAfv1gjBdWZN0hUYBRAgMBAAEwDQYJKoZIhvcN
        AQELBQADggEBAABzo82TxGp5poVkd5pLWj5ACgcBv8Cs6oH9D+4Jz9BmyuBUsQXh
        2aG0hQAe1mU61C9konsl/GTW8umJQ4M4lYEztXXwMf5PlBMGwebM0ZbSGg6jKtZg
        WCgJ3eP/FMmyXGL5Jji5+e09eObhUDVle4tdi0On97zBoz85W02rgWFAqZJwiEAP
        t+c7jX7uOSBq2/38iGStlrX5yB1at/gJXXiA5CL5OtlR3Okvb0/QH37efO1Nu39m
        lFi0ODPAVyXjVypAiLguDxPn6AtDTdk9Iw9B19OD4NrzNRWgSSX5vuxo/VcRcgWk
        3gEe9Ca0ZKN20q9XgthAiFFjl1S9ZgdA6Zc=
        -----END CERTIFICATE-----\"\"\",
            private_key=\"\"\"-----BEGIN RSA PRIVATE KEY-----
        MIIEowIBAAKCAQEAyjCheapjf7qDI3R9w/Gj0XFDgNLPK2aWIRvM25BdY/IB2KAf
        xQ7zOxu3X1bMo2zMCzsSrwIVrxx0qRM/7e4AfkHcKwDIjCcBprQp164dhFol4GpT
        HtcuGv0+Ue6vpuE9cxQE3/pG5x1n5EhheFu2+lAaGh/vUrUPeQp6szX/9qfzn+/k
        tF8wsRV8PBiFc2ZNMRXupRyc2qtOn95r86w/uK6TuerU1L64E2P5+tYLBWT+39Ai
        diUIenjURThFEn8Srvw8v/iLEMaKwiWNkOfWQH96166gzddOI8PT5cJy5an2c90T
        DgzegVo8RPK70bU9KjH7gH79YIwXVmTdIVGAUQIDAQABAoIBAE1J4a/8biR5S3/W
        G+03BYQeY8tuyjqw8FqfoeOcf9agwAvqybouSNQjeCk9qOQfxq/UWQQFK/zQR9gJ
        v7pX7GBXFK5rkj3g+0SaQhRsPmRFgY0Tl8qGPt2aSKRRNVv5ZeADmwlzRn86QmiF
        Mp0rkfqFfDTYWEepZszCML0ouzuxsW/9tq7rvtSjsgATNt31B3vFa3D3JBi31jUh
        5nfR9A3bATze7mQw3byEDiVl5ASRDgYyur403P1fDnMy9DBHZ8NaPOsFF6OBpJal
        BJsG5z00hll5PFN2jfmBQKlvAeU7wfwqdaSnGHOfqf2DeTTaFjIQ4gUhRn/m6pLo
        6kXttLECgYEA9sng0Qz/TcPFfM4tQ1gyvB1cKnnGIwg1FP8sfUjbbEgjaHhA224S
        k3BxtX2Kq6fhTXuwusAFc6OVMAZ76FgrQ5K4Ci7+DTsrF28z4b8td+p+lO/DxgP9
        lTgN+ddsiTOV4fUef9Z3yY0Zr0CnBUMbQYRaV2UIbCdiB0G4V/bt9TsCgYEA0bya
        Oo9wGI0RJV0bYP7qwO74Ra1/i1viWbRlS7jU37Q+AZstrlKcQ5CTPzOjKFKMiUzl
        4miWacZ0/q2n+Mvd7NbXGXTLijahnyOYKaHJYyh4oBymfkgAifRstE0Ki9gdvArb
        /I+emC0GvLSyfGN8UUeDJs4NmqdEXGqjo2JOV+MCgYALFv1MR5o9Y1u/hQBRs2fs
        PiGDIx+9OUQxYloccyaxEfjNXAIGGkcpavchIbgWiJ++PJ2vdquIC8TLeK8evL+M
        9M3iX0Q5UfxYvD2HmnCvn9D6Xl/cyRcfGnq+TGjrLW9BzSMGuZt+aiHKV0xqFx7l
        bc4leTvMqGRmURS4lzcQOwKBgQCDzA/i4sYfN25h21tcHXSpnsG3D2rJyQi5NCo/
        ZjunA92/JqOTGuiFcLGHEszhhtY3ZXJET1LNz18vtzKJnpqrvOnYXlOVW/U+SqDQ
        8JDb1c/PVZGuY1KrXkR9HLiW3kz5IJ3S3PFdUVYdeTN8BQxXCyg4V12nJJtJs912
        y0zN3wKBgGDS6YttCN6aI4EOABYE8fI1EYQ7vhfiYsaWGWSR1l6bQey7KR6M1ACz
        ZzMASNyytVt12yXE4/Emv6/pYqigbDLfL1zQJSLJ3EHJYTh2RxjR+AaGDudYFG/T
        liQ9YXhV5Iu2x1pNwrtFnssDdaaGpfA7l3xC00BL7Z+SAJyI4QKA
        -----END RSA PRIVATE KEY-----\"\"\")
        https = alicloud.slb.Listener("https",
            load_balancer_id=instance.id,
            backend_port=80,
            frontend_port=443,
            protocol="https",
            sticky_session="on",
            sticky_session_type="insert",
            cookie="testslblistenercookie",
            cookie_timeout=86400,
            health_check="on",
            health_check_uri="/cons",
            health_check_connect_port=20,
            healthy_threshold=8,
            unhealthy_threshold=8,
            health_check_timeout=8,
            health_check_interval=5,
            health_check_http_code="http_2xx,http_3xx",
            bandwidth=10,
            server_certificate_id=domain_extension_server_certificate.id)
        example1 = alicloud.slb.DomainExtension("example1",
            load_balancer_id=instance.id,
            frontend_port=https.frontend_port,
            domain="www.test.com",
            server_certificate_id=domain_extension_server_certificate.id)
        ```

        ## Import

        Load balancer domain_extension can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:slb/domainExtension:DomainExtension example de-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param DomainExtensionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainExtensionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_protection_validation: Optional[pulumi.Input[bool]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 frontend_port: Optional[pulumi.Input[int]] = None,
                 load_balancer_id: Optional[pulumi.Input[str]] = None,
                 server_certificate_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainExtensionArgs.__new__(DomainExtensionArgs)

            __props__.__dict__["delete_protection_validation"] = delete_protection_validation
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if frontend_port is None and not opts.urn:
                raise TypeError("Missing required property 'frontend_port'")
            __props__.__dict__["frontend_port"] = frontend_port
            if load_balancer_id is None and not opts.urn:
                raise TypeError("Missing required property 'load_balancer_id'")
            __props__.__dict__["load_balancer_id"] = load_balancer_id
            if server_certificate_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_certificate_id'")
            __props__.__dict__["server_certificate_id"] = server_certificate_id
        super(DomainExtension, __self__).__init__(
            'alicloud:slb/domainExtension:DomainExtension',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delete_protection_validation: Optional[pulumi.Input[bool]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            frontend_port: Optional[pulumi.Input[int]] = None,
            load_balancer_id: Optional[pulumi.Input[str]] = None,
            server_certificate_id: Optional[pulumi.Input[str]] = None) -> 'DomainExtension':
        """
        Get an existing DomainExtension resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_protection_validation: Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        :param pulumi.Input[str] domain: The domain name.
        :param pulumi.Input[int] frontend_port: The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        :param pulumi.Input[str] load_balancer_id: The ID of the SLB instance.
        :param pulumi.Input[str] server_certificate_id: The ID of the certificate used by the domain name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainExtensionState.__new__(_DomainExtensionState)

        __props__.__dict__["delete_protection_validation"] = delete_protection_validation
        __props__.__dict__["domain"] = domain
        __props__.__dict__["frontend_port"] = frontend_port
        __props__.__dict__["load_balancer_id"] = load_balancer_id
        __props__.__dict__["server_certificate_id"] = server_certificate_id
        return DomainExtension(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deleteProtectionValidation")
    def delete_protection_validation(self) -> pulumi.Output[Optional[bool]]:
        """
        Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        """
        return pulumi.get(self, "delete_protection_validation")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The domain name.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> pulumi.Output[int]:
        """
        The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
        """
        return pulumi.get(self, "frontend_port")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> pulumi.Output[str]:
        """
        The ID of the SLB instance.
        """
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="serverCertificateId")
    def server_certificate_id(self) -> pulumi.Output[str]:
        """
        The ID of the certificate used by the domain name.
        """
        return pulumi.get(self, "server_certificate_id")

