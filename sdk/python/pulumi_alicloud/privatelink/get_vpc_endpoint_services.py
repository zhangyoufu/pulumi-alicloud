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

__all__ = [
    'GetVpcEndpointServicesResult',
    'AwaitableGetVpcEndpointServicesResult',
    'get_vpc_endpoint_services',
    'get_vpc_endpoint_services_output',
]

@pulumi.output_type
class GetVpcEndpointServicesResult:
    """
    A collection of values returned by getVpcEndpointServices.
    """
    def __init__(__self__, auto_accept_connection=None, id=None, ids=None, name_regex=None, names=None, output_file=None, service_business_status=None, services=None, status=None, vpc_endpoint_service_name=None):
        if auto_accept_connection and not isinstance(auto_accept_connection, bool):
            raise TypeError("Expected argument 'auto_accept_connection' to be a bool")
        pulumi.set(__self__, "auto_accept_connection", auto_accept_connection)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if service_business_status and not isinstance(service_business_status, str):
            raise TypeError("Expected argument 'service_business_status' to be a str")
        pulumi.set(__self__, "service_business_status", service_business_status)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vpc_endpoint_service_name and not isinstance(vpc_endpoint_service_name, str):
            raise TypeError("Expected argument 'vpc_endpoint_service_name' to be a str")
        pulumi.set(__self__, "vpc_endpoint_service_name", vpc_endpoint_service_name)

    @property
    @pulumi.getter(name="autoAcceptConnection")
    def auto_accept_connection(self) -> Optional[bool]:
        return pulumi.get(self, "auto_accept_connection")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="serviceBusinessStatus")
    def service_business_status(self) -> Optional[str]:
        return pulumi.get(self, "service_business_status")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetVpcEndpointServicesServiceResult']:
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcEndpointServiceName")
    def vpc_endpoint_service_name(self) -> Optional[str]:
        return pulumi.get(self, "vpc_endpoint_service_name")


class AwaitableGetVpcEndpointServicesResult(GetVpcEndpointServicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcEndpointServicesResult(
            auto_accept_connection=self.auto_accept_connection,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            service_business_status=self.service_business_status,
            services=self.services,
            status=self.status,
            vpc_endpoint_service_name=self.vpc_endpoint_service_name)


def get_vpc_endpoint_services(auto_accept_connection: Optional[bool] = None,
                              ids: Optional[Sequence[str]] = None,
                              name_regex: Optional[str] = None,
                              output_file: Optional[str] = None,
                              service_business_status: Optional[str] = None,
                              status: Optional[str] = None,
                              vpc_endpoint_service_name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcEndpointServicesResult:
    """
    This data source provides the Privatelink Vpc Endpoint Services of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.109.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.privatelink.get_vpc_endpoint_services(ids=["example_value"],
        name_regex="the_resource_name")
    pulumi.export("firstPrivatelinkVpcEndpointServiceId", example.services[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param bool auto_accept_connection: Whether to automatically accept terminal node connections..
    :param Sequence[str] ids: A list of Vpc Endpoint Service IDs.
    :param str name_regex: A regex string to filter results by Vpc Endpoint Service name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str service_business_status: The business status of the terminal node service..
    :param str status: The Status of Vpc Endpoint Service.
    :param str vpc_endpoint_service_name: The name of Vpc Endpoint Service.
    """
    __args__ = dict()
    __args__['autoAcceptConnection'] = auto_accept_connection
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['serviceBusinessStatus'] = service_business_status
    __args__['status'] = status
    __args__['vpcEndpointServiceName'] = vpc_endpoint_service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:privatelink/getVpcEndpointServices:getVpcEndpointServices', __args__, opts=opts, typ=GetVpcEndpointServicesResult).value

    return AwaitableGetVpcEndpointServicesResult(
        auto_accept_connection=pulumi.get(__ret__, 'auto_accept_connection'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        service_business_status=pulumi.get(__ret__, 'service_business_status'),
        services=pulumi.get(__ret__, 'services'),
        status=pulumi.get(__ret__, 'status'),
        vpc_endpoint_service_name=pulumi.get(__ret__, 'vpc_endpoint_service_name'))


@_utilities.lift_output_func(get_vpc_endpoint_services)
def get_vpc_endpoint_services_output(auto_accept_connection: Optional[pulumi.Input[Optional[bool]]] = None,
                                     ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                     name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                     service_business_status: Optional[pulumi.Input[Optional[str]]] = None,
                                     status: Optional[pulumi.Input[Optional[str]]] = None,
                                     vpc_endpoint_service_name: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcEndpointServicesResult]:
    """
    This data source provides the Privatelink Vpc Endpoint Services of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.109.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.privatelink.get_vpc_endpoint_services(ids=["example_value"],
        name_regex="the_resource_name")
    pulumi.export("firstPrivatelinkVpcEndpointServiceId", example.services[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param bool auto_accept_connection: Whether to automatically accept terminal node connections..
    :param Sequence[str] ids: A list of Vpc Endpoint Service IDs.
    :param str name_regex: A regex string to filter results by Vpc Endpoint Service name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str service_business_status: The business status of the terminal node service..
    :param str status: The Status of Vpc Endpoint Service.
    :param str vpc_endpoint_service_name: The name of Vpc Endpoint Service.
    """
    ...
