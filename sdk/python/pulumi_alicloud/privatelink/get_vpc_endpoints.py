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
    'GetVpcEndpointsResult',
    'AwaitableGetVpcEndpointsResult',
    'get_vpc_endpoints',
    'get_vpc_endpoints_output',
]

@pulumi.output_type
class GetVpcEndpointsResult:
    """
    A collection of values returned by getVpcEndpoints.
    """
    def __init__(__self__, connection_status=None, enable_details=None, endpoints=None, id=None, ids=None, name_regex=None, names=None, output_file=None, service_name=None, status=None, vpc_endpoint_name=None, vpc_id=None):
        if connection_status and not isinstance(connection_status, str):
            raise TypeError("Expected argument 'connection_status' to be a str")
        pulumi.set(__self__, "connection_status", connection_status)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if endpoints and not isinstance(endpoints, list):
            raise TypeError("Expected argument 'endpoints' to be a list")
        pulumi.set(__self__, "endpoints", endpoints)
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
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vpc_endpoint_name and not isinstance(vpc_endpoint_name, str):
            raise TypeError("Expected argument 'vpc_endpoint_name' to be a str")
        pulumi.set(__self__, "vpc_endpoint_name", vpc_endpoint_name)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="connectionStatus")
    def connection_status(self) -> Optional[str]:
        return pulumi.get(self, "connection_status")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter
    def endpoints(self) -> Sequence['outputs.GetVpcEndpointsEndpointResult']:
        return pulumi.get(self, "endpoints")

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
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[str]:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpcEndpointName")
    def vpc_endpoint_name(self) -> Optional[str]:
        return pulumi.get(self, "vpc_endpoint_name")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableGetVpcEndpointsResult(GetVpcEndpointsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcEndpointsResult(
            connection_status=self.connection_status,
            enable_details=self.enable_details,
            endpoints=self.endpoints,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            service_name=self.service_name,
            status=self.status,
            vpc_endpoint_name=self.vpc_endpoint_name,
            vpc_id=self.vpc_id)


def get_vpc_endpoints(connection_status: Optional[str] = None,
                      enable_details: Optional[bool] = None,
                      ids: Optional[Sequence[str]] = None,
                      name_regex: Optional[str] = None,
                      output_file: Optional[str] = None,
                      service_name: Optional[str] = None,
                      status: Optional[str] = None,
                      vpc_endpoint_name: Optional[str] = None,
                      vpc_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcEndpointsResult:
    """
    This data source provides the Privatelink Vpc Endpoints of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.109.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.privatelink.get_vpc_endpoints(ids=["example_value"],
        name_regex="the_resource_name")
    pulumi.export("firstPrivatelinkVpcEndpointId", example.endpoints[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str connection_status: The status of Connection.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Vpc Endpoint IDs.
    :param str name_regex: A regex string to filter results by Vpc Endpoint name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str service_name: The name of the terminal node service associated with the terminal node.
    :param str status: The status of Vpc Endpoint.
    :param str vpc_endpoint_name: The name of Vpc Endpoint.
    :param str vpc_id: The private network to which the terminal node belongs.
    """
    __args__ = dict()
    __args__['connectionStatus'] = connection_status
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['serviceName'] = service_name
    __args__['status'] = status
    __args__['vpcEndpointName'] = vpc_endpoint_name
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:privatelink/getVpcEndpoints:getVpcEndpoints', __args__, opts=opts, typ=GetVpcEndpointsResult).value

    return AwaitableGetVpcEndpointsResult(
        connection_status=pulumi.get(__ret__, 'connection_status'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        endpoints=pulumi.get(__ret__, 'endpoints'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        service_name=pulumi.get(__ret__, 'service_name'),
        status=pulumi.get(__ret__, 'status'),
        vpc_endpoint_name=pulumi.get(__ret__, 'vpc_endpoint_name'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))


@_utilities.lift_output_func(get_vpc_endpoints)
def get_vpc_endpoints_output(connection_status: Optional[pulumi.Input[Optional[str]]] = None,
                             enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                             ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                             output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             service_name: Optional[pulumi.Input[Optional[str]]] = None,
                             status: Optional[pulumi.Input[Optional[str]]] = None,
                             vpc_endpoint_name: Optional[pulumi.Input[Optional[str]]] = None,
                             vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcEndpointsResult]:
    """
    This data source provides the Privatelink Vpc Endpoints of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.109.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.privatelink.get_vpc_endpoints(ids=["example_value"],
        name_regex="the_resource_name")
    pulumi.export("firstPrivatelinkVpcEndpointId", example.endpoints[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str connection_status: The status of Connection.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Vpc Endpoint IDs.
    :param str name_regex: A regex string to filter results by Vpc Endpoint name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str service_name: The name of the terminal node service associated with the terminal node.
    :param str status: The status of Vpc Endpoint.
    :param str vpc_endpoint_name: The name of Vpc Endpoint.
    :param str vpc_id: The private network to which the terminal node belongs.
    """
    ...
