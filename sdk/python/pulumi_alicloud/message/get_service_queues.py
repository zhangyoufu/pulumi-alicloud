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
    'GetServiceQueuesResult',
    'AwaitableGetServiceQueuesResult',
    'get_service_queues',
    'get_service_queues_output',
]

@pulumi.output_type
class GetServiceQueuesResult:
    """
    A collection of values returned by getServiceQueues.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, names=None, output_file=None, page_number=None, page_size=None, queue_name=None, queues=None):
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
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)
        if queue_name and not isinstance(queue_name, str):
            raise TypeError("Expected argument 'queue_name' to be a str")
        pulumi.set(__self__, "queue_name", queue_name)
        if queues and not isinstance(queues, list):
            raise TypeError("Expected argument 'queues' to be a list")
        pulumi.set(__self__, "queues", queues)

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
        """
        A list of Queue names.
        """
        return pulumi.get(self, "names")

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
    @pulumi.getter(name="queueName")
    def queue_name(self) -> Optional[str]:
        """
        The name of the queue.
        """
        return pulumi.get(self, "queue_name")

    @property
    @pulumi.getter
    def queues(self) -> Sequence['outputs.GetServiceQueuesQueueResult']:
        """
        A list of Queues. Each element contains the following attributes:
        """
        return pulumi.get(self, "queues")


class AwaitableGetServiceQueuesResult(GetServiceQueuesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceQueuesResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size,
            queue_name=self.queue_name,
            queues=self.queues)


def get_service_queues(ids: Optional[Sequence[str]] = None,
                       name_regex: Optional[str] = None,
                       output_file: Optional[str] = None,
                       page_number: Optional[int] = None,
                       page_size: Optional[int] = None,
                       queue_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceQueuesResult:
    """
    This data source provides the Message Notification Service Queues of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.188.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.message.get_service_queues(ids=["example_id"])
    pulumi.export("queueId1", ids.queues[0].id)
    name = alicloud.message.get_service_queues(queue_name="tf-example")
    pulumi.export("queueId2", name.queues[0].id)
    ```


    :param Sequence[str] ids: A list of Queue IDs. Its element value is same as Queue Name.
    :param str name_regex: A regex string to filter results by Queue name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str queue_name: The name of the queue.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    __args__['queueName'] = queue_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:message/getServiceQueues:getServiceQueues', __args__, opts=opts, typ=GetServiceQueuesResult).value

    return AwaitableGetServiceQueuesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        page_number=__ret__.page_number,
        page_size=__ret__.page_size,
        queue_name=__ret__.queue_name,
        queues=__ret__.queues)


@_utilities.lift_output_func(get_service_queues)
def get_service_queues_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              page_number: Optional[pulumi.Input[Optional[int]]] = None,
                              page_size: Optional[pulumi.Input[Optional[int]]] = None,
                              queue_name: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceQueuesResult]:
    """
    This data source provides the Message Notification Service Queues of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.188.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.message.get_service_queues(ids=["example_id"])
    pulumi.export("queueId1", ids.queues[0].id)
    name = alicloud.message.get_service_queues(queue_name="tf-example")
    pulumi.export("queueId2", name.queues[0].id)
    ```


    :param Sequence[str] ids: A list of Queue IDs. Its element value is same as Queue Name.
    :param str name_regex: A regex string to filter results by Queue name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str queue_name: The name of the queue.
    """
    ...
