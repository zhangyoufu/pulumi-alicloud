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
    'GetAcceleratorSpareIpAttachmentsResult',
    'AwaitableGetAcceleratorSpareIpAttachmentsResult',
    'get_accelerator_spare_ip_attachments',
    'get_accelerator_spare_ip_attachments_output',
]

@pulumi.output_type
class GetAcceleratorSpareIpAttachmentsResult:
    """
    A collection of values returned by getAcceleratorSpareIpAttachments.
    """
    def __init__(__self__, accelerator_id=None, attachments=None, id=None, ids=None, output_file=None, status=None):
        if accelerator_id and not isinstance(accelerator_id, str):
            raise TypeError("Expected argument 'accelerator_id' to be a str")
        pulumi.set(__self__, "accelerator_id", accelerator_id)
        if attachments and not isinstance(attachments, list):
            raise TypeError("Expected argument 'attachments' to be a list")
        pulumi.set(__self__, "attachments", attachments)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="acceleratorId")
    def accelerator_id(self) -> str:
        return pulumi.get(self, "accelerator_id")

    @property
    @pulumi.getter
    def attachments(self) -> Sequence['outputs.GetAcceleratorSpareIpAttachmentsAttachmentResult']:
        return pulumi.get(self, "attachments")

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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetAcceleratorSpareIpAttachmentsResult(GetAcceleratorSpareIpAttachmentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAcceleratorSpareIpAttachmentsResult(
            accelerator_id=self.accelerator_id,
            attachments=self.attachments,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            status=self.status)


def get_accelerator_spare_ip_attachments(accelerator_id: Optional[str] = None,
                                         ids: Optional[Sequence[str]] = None,
                                         output_file: Optional[str] = None,
                                         status: Optional[str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAcceleratorSpareIpAttachmentsResult:
    """
    This data source provides the Ga Accelerator Spare Ip Attachments of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.167.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ga.get_accelerator_spare_ip_attachments(accelerator_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("gaAcceleratorSpareIpAttachmentId1", ids.attachments[0].id)
    ```


    :param str accelerator_id: The ID of the global acceleration instance.
    :param Sequence[str] ids: A list of Accelerator Spare Ip Attachment IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the standby CNAME IP address. Valid values: `active`, `inuse`.
    """
    __args__ = dict()
    __args__['acceleratorId'] = accelerator_id
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ga/getAcceleratorSpareIpAttachments:getAcceleratorSpareIpAttachments', __args__, opts=opts, typ=GetAcceleratorSpareIpAttachmentsResult).value

    return AwaitableGetAcceleratorSpareIpAttachmentsResult(
        accelerator_id=__ret__.accelerator_id,
        attachments=__ret__.attachments,
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        status=__ret__.status)


@_utilities.lift_output_func(get_accelerator_spare_ip_attachments)
def get_accelerator_spare_ip_attachments_output(accelerator_id: Optional[pulumi.Input[str]] = None,
                                                ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                                status: Optional[pulumi.Input[Optional[str]]] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAcceleratorSpareIpAttachmentsResult]:
    """
    This data source provides the Ga Accelerator Spare Ip Attachments of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.167.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ga.get_accelerator_spare_ip_attachments(accelerator_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("gaAcceleratorSpareIpAttachmentId1", ids.attachments[0].id)
    ```


    :param str accelerator_id: The ID of the global acceleration instance.
    :param Sequence[str] ids: A list of Accelerator Spare Ip Attachment IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the standby CNAME IP address. Valid values: `active`, `inuse`.
    """
    ...
