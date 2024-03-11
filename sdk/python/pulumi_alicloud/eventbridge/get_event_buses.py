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
    'GetEventBusesResult',
    'AwaitableGetEventBusesResult',
    'get_event_buses',
    'get_event_buses_output',
]

@pulumi.output_type
class GetEventBusesResult:
    """
    A collection of values returned by getEventBuses.
    """
    def __init__(__self__, buses=None, event_bus_type=None, id=None, ids=None, name_prefix=None, name_regex=None, names=None, output_file=None):
        if buses and not isinstance(buses, list):
            raise TypeError("Expected argument 'buses' to be a list")
        pulumi.set(__self__, "buses", buses)
        if event_bus_type and not isinstance(event_bus_type, str):
            raise TypeError("Expected argument 'event_bus_type' to be a str")
        pulumi.set(__self__, "event_bus_type", event_bus_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_prefix and not isinstance(name_prefix, str):
            raise TypeError("Expected argument 'name_prefix' to be a str")
        pulumi.set(__self__, "name_prefix", name_prefix)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)

    @property
    @pulumi.getter
    def buses(self) -> Sequence['outputs.GetEventBusesBusResult']:
        return pulumi.get(self, "buses")

    @property
    @pulumi.getter(name="eventBusType")
    def event_bus_type(self) -> Optional[str]:
        return pulumi.get(self, "event_bus_type")

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
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[str]:
        return pulumi.get(self, "name_prefix")

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


class AwaitableGetEventBusesResult(GetEventBusesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventBusesResult(
            buses=self.buses,
            event_bus_type=self.event_bus_type,
            id=self.id,
            ids=self.ids,
            name_prefix=self.name_prefix,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)


def get_event_buses(event_bus_type: Optional[str] = None,
                    ids: Optional[Sequence[str]] = None,
                    name_prefix: Optional[str] = None,
                    name_regex: Optional[str] = None,
                    output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventBusesResult:
    """
    This data source provides the Event Bridge Event Buses of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.129.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.eventbridge.get_event_buses()
    pulumi.export("eventBridgeEventBusId1", ids.buses[0].id)
    name_regex = alicloud.eventbridge.get_event_buses(name_regex="^my-EventBus")
    pulumi.export("eventBridgeEventBusId2", name_regex.buses[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str event_bus_type: The event bus type.
    :param Sequence[str] ids: A list of Event Bus IDs. Its element value is same as Event Bus Name.
    :param str name_prefix: The name prefix.
    :param str name_regex: A regex string to filter results by Event Bus name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['eventBusType'] = event_bus_type
    __args__['ids'] = ids
    __args__['namePrefix'] = name_prefix
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:eventbridge/getEventBuses:getEventBuses', __args__, opts=opts, typ=GetEventBusesResult).value

    return AwaitableGetEventBusesResult(
        buses=pulumi.get(__ret__, 'buses'),
        event_bus_type=pulumi.get(__ret__, 'event_bus_type'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_prefix=pulumi.get(__ret__, 'name_prefix'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'))


@_utilities.lift_output_func(get_event_buses)
def get_event_buses_output(event_bus_type: Optional[pulumi.Input[Optional[str]]] = None,
                           ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           name_prefix: Optional[pulumi.Input[Optional[str]]] = None,
                           name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEventBusesResult]:
    """
    This data source provides the Event Bridge Event Buses of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.129.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.eventbridge.get_event_buses()
    pulumi.export("eventBridgeEventBusId1", ids.buses[0].id)
    name_regex = alicloud.eventbridge.get_event_buses(name_regex="^my-EventBus")
    pulumi.export("eventBridgeEventBusId2", name_regex.buses[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str event_bus_type: The event bus type.
    :param Sequence[str] ids: A list of Event Bus IDs. Its element value is same as Event Bus Name.
    :param str name_prefix: The name prefix.
    :param str name_regex: A regex string to filter results by Event Bus name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
