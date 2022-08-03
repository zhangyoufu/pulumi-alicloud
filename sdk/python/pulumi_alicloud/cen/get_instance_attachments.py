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
    'GetInstanceAttachmentsResult',
    'AwaitableGetInstanceAttachmentsResult',
    'get_instance_attachments',
    'get_instance_attachments_output',
]

@pulumi.output_type
class GetInstanceAttachmentsResult:
    """
    A collection of values returned by getInstanceAttachments.
    """
    def __init__(__self__, attachments=None, child_instance_region_id=None, child_instance_type=None, id=None, ids=None, instance_id=None, output_file=None, status=None):
        if attachments and not isinstance(attachments, list):
            raise TypeError("Expected argument 'attachments' to be a list")
        pulumi.set(__self__, "attachments", attachments)
        if child_instance_region_id and not isinstance(child_instance_region_id, str):
            raise TypeError("Expected argument 'child_instance_region_id' to be a str")
        pulumi.set(__self__, "child_instance_region_id", child_instance_region_id)
        if child_instance_type and not isinstance(child_instance_type, str):
            raise TypeError("Expected argument 'child_instance_type' to be a str")
        pulumi.set(__self__, "child_instance_type", child_instance_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def attachments(self) -> Sequence['outputs.GetInstanceAttachmentsAttachmentResult']:
        """
        A list of CEN Instance Attachments. Each element contains the following attributes:
        """
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter(name="childInstanceRegionId")
    def child_instance_region_id(self) -> Optional[str]:
        """
        The ID of the region to which the network belongs.
        """
        return pulumi.get(self, "child_instance_region_id")

    @property
    @pulumi.getter(name="childInstanceType")
    def child_instance_type(self) -> Optional[str]:
        """
        The type of the associated network.
        """
        return pulumi.get(self, "child_instance_type")

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
        """
        A list of CEN Instance Attachment IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The ID of the CEN instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the network.
        """
        return pulumi.get(self, "status")


class AwaitableGetInstanceAttachmentsResult(GetInstanceAttachmentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceAttachmentsResult(
            attachments=self.attachments,
            child_instance_region_id=self.child_instance_region_id,
            child_instance_type=self.child_instance_type,
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            output_file=self.output_file,
            status=self.status)


def get_instance_attachments(child_instance_region_id: Optional[str] = None,
                             child_instance_type: Optional[str] = None,
                             instance_id: Optional[str] = None,
                             output_file: Optional[str] = None,
                             status: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceAttachmentsResult:
    """
    This data source provides Cen Instance Attachments of the current Alibaba Cloud User.

    > **NOTE:** Available in v1.97.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cen.get_instance_attachments(instance_id="cen-o40h17ll9w********")
    pulumi.export("theFirstAttachmentedInstanceId", example.attachments[0].child_instance_id)
    ```


    :param str child_instance_region_id: The region to which the network to be queried belongs.
    :param str child_instance_type: The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
    :param str instance_id: The ID of the CEN instance.
    :param str status: The status of the Cen Child Instance Attachment. Valid value: `Attaching`, `Attached` and `Aetaching`.
    """
    __args__ = dict()
    __args__['childInstanceRegionId'] = child_instance_region_id
    __args__['childInstanceType'] = child_instance_type
    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cen/getInstanceAttachments:getInstanceAttachments', __args__, opts=opts, typ=GetInstanceAttachmentsResult).value

    return AwaitableGetInstanceAttachmentsResult(
        attachments=__ret__.attachments,
        child_instance_region_id=__ret__.child_instance_region_id,
        child_instance_type=__ret__.child_instance_type,
        id=__ret__.id,
        ids=__ret__.ids,
        instance_id=__ret__.instance_id,
        output_file=__ret__.output_file,
        status=__ret__.status)


@_utilities.lift_output_func(get_instance_attachments)
def get_instance_attachments_output(child_instance_region_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    child_instance_type: Optional[pulumi.Input[Optional[str]]] = None,
                                    instance_id: Optional[pulumi.Input[str]] = None,
                                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    status: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceAttachmentsResult]:
    """
    This data source provides Cen Instance Attachments of the current Alibaba Cloud User.

    > **NOTE:** Available in v1.97.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.cen.get_instance_attachments(instance_id="cen-o40h17ll9w********")
    pulumi.export("theFirstAttachmentedInstanceId", example.attachments[0].child_instance_id)
    ```


    :param str child_instance_region_id: The region to which the network to be queried belongs.
    :param str child_instance_type: The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
    :param str instance_id: The ID of the CEN instance.
    :param str status: The status of the Cen Child Instance Attachment. Valid value: `Attaching`, `Attached` and `Aetaching`.
    """
    ...
