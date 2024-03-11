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
    'GetCustomRoutingEndpointGroupDestinationsResult',
    'AwaitableGetCustomRoutingEndpointGroupDestinationsResult',
    'get_custom_routing_endpoint_group_destinations',
    'get_custom_routing_endpoint_group_destinations_output',
]

@pulumi.output_type
class GetCustomRoutingEndpointGroupDestinationsResult:
    """
    A collection of values returned by getCustomRoutingEndpointGroupDestinations.
    """
    def __init__(__self__, accelerator_id=None, custom_routing_endpoint_group_destinations=None, endpoint_group_id=None, from_port=None, id=None, ids=None, listener_id=None, output_file=None, page_number=None, page_size=None, protocols=None, to_port=None):
        if accelerator_id and not isinstance(accelerator_id, str):
            raise TypeError("Expected argument 'accelerator_id' to be a str")
        pulumi.set(__self__, "accelerator_id", accelerator_id)
        if custom_routing_endpoint_group_destinations and not isinstance(custom_routing_endpoint_group_destinations, list):
            raise TypeError("Expected argument 'custom_routing_endpoint_group_destinations' to be a list")
        pulumi.set(__self__, "custom_routing_endpoint_group_destinations", custom_routing_endpoint_group_destinations)
        if endpoint_group_id and not isinstance(endpoint_group_id, str):
            raise TypeError("Expected argument 'endpoint_group_id' to be a str")
        pulumi.set(__self__, "endpoint_group_id", endpoint_group_id)
        if from_port and not isinstance(from_port, int):
            raise TypeError("Expected argument 'from_port' to be a int")
        pulumi.set(__self__, "from_port", from_port)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if listener_id and not isinstance(listener_id, str):
            raise TypeError("Expected argument 'listener_id' to be a str")
        pulumi.set(__self__, "listener_id", listener_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)
        if protocols and not isinstance(protocols, list):
            raise TypeError("Expected argument 'protocols' to be a list")
        pulumi.set(__self__, "protocols", protocols)
        if to_port and not isinstance(to_port, int):
            raise TypeError("Expected argument 'to_port' to be a int")
        pulumi.set(__self__, "to_port", to_port)

    @property
    @pulumi.getter(name="acceleratorId")
    def accelerator_id(self) -> str:
        """
        The ID of the GA instance.
        """
        return pulumi.get(self, "accelerator_id")

    @property
    @pulumi.getter(name="customRoutingEndpointGroupDestinations")
    def custom_routing_endpoint_group_destinations(self) -> Sequence['outputs.GetCustomRoutingEndpointGroupDestinationsCustomRoutingEndpointGroupDestinationResult']:
        """
        A list of Custom Routing Endpoint Group Destinations. Each element contains the following attributes:
        """
        return pulumi.get(self, "custom_routing_endpoint_group_destinations")

    @property
    @pulumi.getter(name="endpointGroupId")
    def endpoint_group_id(self) -> Optional[str]:
        """
        The ID of the Custom Routing Endpoint Group.
        """
        return pulumi.get(self, "endpoint_group_id")

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> Optional[int]:
        """
        The start port of the backend service port range of the endpoint group.
        """
        return pulumi.get(self, "from_port")

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
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[str]:
        """
        The ID of the listener.
        """
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="pageNumber")
    def page_number(self) -> Optional[int]:
        return pulumi.get(self, "page_number")

    @property
    @pulumi.getter(name="pageSize")
    def page_size(self) -> Optional[int]:
        return pulumi.get(self, "page_size")

    @property
    @pulumi.getter
    def protocols(self) -> Optional[Sequence[str]]:
        """
        The backend service protocol of the endpoint group.
        """
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> Optional[int]:
        """
        The end port of the backend service port range of the endpoint group.
        """
        return pulumi.get(self, "to_port")


class AwaitableGetCustomRoutingEndpointGroupDestinationsResult(GetCustomRoutingEndpointGroupDestinationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomRoutingEndpointGroupDestinationsResult(
            accelerator_id=self.accelerator_id,
            custom_routing_endpoint_group_destinations=self.custom_routing_endpoint_group_destinations,
            endpoint_group_id=self.endpoint_group_id,
            from_port=self.from_port,
            id=self.id,
            ids=self.ids,
            listener_id=self.listener_id,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size,
            protocols=self.protocols,
            to_port=self.to_port)


def get_custom_routing_endpoint_group_destinations(accelerator_id: Optional[str] = None,
                                                   endpoint_group_id: Optional[str] = None,
                                                   from_port: Optional[int] = None,
                                                   ids: Optional[Sequence[str]] = None,
                                                   listener_id: Optional[str] = None,
                                                   output_file: Optional[str] = None,
                                                   page_number: Optional[int] = None,
                                                   page_size: Optional[int] = None,
                                                   protocols: Optional[Sequence[str]] = None,
                                                   to_port: Optional[int] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomRoutingEndpointGroupDestinationsResult:
    """
    This data source provides the Global Accelerator (GA) Custom Routing Endpoint Group Destinations of the current Alibaba Cloud user.

    > **NOTE:** Available in 1.197.0+

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ga.get_custom_routing_endpoint_group_destinations(ids=["example_id"],
        accelerator_id="your_accelerator_id")
    pulumi.export("gaCustomRoutingEndpointGroupDestinationsId1", ids.custom_routing_endpoint_group_destinations[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str accelerator_id: The ID of the GA instance.
    :param str endpoint_group_id: The ID of the endpoint group.
    :param int from_port: The start port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
    :param Sequence[str] ids: A list of Custom Routing Endpoint Group Destination IDs.
    :param str listener_id: The ID of the listener.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param Sequence[str] protocols: The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
    :param int to_port: The end port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
    """
    __args__ = dict()
    __args__['acceleratorId'] = accelerator_id
    __args__['endpointGroupId'] = endpoint_group_id
    __args__['fromPort'] = from_port
    __args__['ids'] = ids
    __args__['listenerId'] = listener_id
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    __args__['protocols'] = protocols
    __args__['toPort'] = to_port
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ga/getCustomRoutingEndpointGroupDestinations:getCustomRoutingEndpointGroupDestinations', __args__, opts=opts, typ=GetCustomRoutingEndpointGroupDestinationsResult).value

    return AwaitableGetCustomRoutingEndpointGroupDestinationsResult(
        accelerator_id=pulumi.get(__ret__, 'accelerator_id'),
        custom_routing_endpoint_group_destinations=pulumi.get(__ret__, 'custom_routing_endpoint_group_destinations'),
        endpoint_group_id=pulumi.get(__ret__, 'endpoint_group_id'),
        from_port=pulumi.get(__ret__, 'from_port'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        listener_id=pulumi.get(__ret__, 'listener_id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        page_number=pulumi.get(__ret__, 'page_number'),
        page_size=pulumi.get(__ret__, 'page_size'),
        protocols=pulumi.get(__ret__, 'protocols'),
        to_port=pulumi.get(__ret__, 'to_port'))


@_utilities.lift_output_func(get_custom_routing_endpoint_group_destinations)
def get_custom_routing_endpoint_group_destinations_output(accelerator_id: Optional[pulumi.Input[str]] = None,
                                                          endpoint_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                          from_port: Optional[pulumi.Input[Optional[int]]] = None,
                                                          ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                                          listener_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                          output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                          page_number: Optional[pulumi.Input[Optional[int]]] = None,
                                                          page_size: Optional[pulumi.Input[Optional[int]]] = None,
                                                          protocols: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                                          to_port: Optional[pulumi.Input[Optional[int]]] = None,
                                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCustomRoutingEndpointGroupDestinationsResult]:
    """
    This data source provides the Global Accelerator (GA) Custom Routing Endpoint Group Destinations of the current Alibaba Cloud user.

    > **NOTE:** Available in 1.197.0+

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ga.get_custom_routing_endpoint_group_destinations(ids=["example_id"],
        accelerator_id="your_accelerator_id")
    pulumi.export("gaCustomRoutingEndpointGroupDestinationsId1", ids.custom_routing_endpoint_group_destinations[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str accelerator_id: The ID of the GA instance.
    :param str endpoint_group_id: The ID of the endpoint group.
    :param int from_port: The start port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
    :param Sequence[str] ids: A list of Custom Routing Endpoint Group Destination IDs.
    :param str listener_id: The ID of the listener.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param Sequence[str] protocols: The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
    :param int to_port: The end port of the backend service port range of the endpoint group. The `from_port` value must be smaller than or equal to the `to_port` value. Valid values: `1` to `65499`.
    """
    ...
