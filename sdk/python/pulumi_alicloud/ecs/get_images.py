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
    'GetImagesResult',
    'AwaitableGetImagesResult',
    'get_images',
    'get_images_output',
]

@pulumi.output_type
class GetImagesResult:
    """
    A collection of values returned by getImages.
    """
    def __init__(__self__, action_type=None, architecture=None, dry_run=None, id=None, ids=None, image_family=None, image_id=None, image_name=None, image_owner_id=None, images=None, instance_type=None, is_support_cloud_init=None, is_support_io_optimized=None, most_recent=None, name_regex=None, os_type=None, output_file=None, owners=None, resource_group_id=None, snapshot_id=None, status=None, tags=None, usage=None):
        if action_type and not isinstance(action_type, str):
            raise TypeError("Expected argument 'action_type' to be a str")
        pulumi.set(__self__, "action_type", action_type)
        if architecture and not isinstance(architecture, str):
            raise TypeError("Expected argument 'architecture' to be a str")
        pulumi.set(__self__, "architecture", architecture)
        if dry_run and not isinstance(dry_run, bool):
            raise TypeError("Expected argument 'dry_run' to be a bool")
        pulumi.set(__self__, "dry_run", dry_run)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if image_family and not isinstance(image_family, str):
            raise TypeError("Expected argument 'image_family' to be a str")
        pulumi.set(__self__, "image_family", image_family)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if image_name and not isinstance(image_name, str):
            raise TypeError("Expected argument 'image_name' to be a str")
        pulumi.set(__self__, "image_name", image_name)
        if image_owner_id and not isinstance(image_owner_id, str):
            raise TypeError("Expected argument 'image_owner_id' to be a str")
        pulumi.set(__self__, "image_owner_id", image_owner_id)
        if images and not isinstance(images, list):
            raise TypeError("Expected argument 'images' to be a list")
        pulumi.set(__self__, "images", images)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if is_support_cloud_init and not isinstance(is_support_cloud_init, bool):
            raise TypeError("Expected argument 'is_support_cloud_init' to be a bool")
        pulumi.set(__self__, "is_support_cloud_init", is_support_cloud_init)
        if is_support_io_optimized and not isinstance(is_support_io_optimized, bool):
            raise TypeError("Expected argument 'is_support_io_optimized' to be a bool")
        pulumi.set(__self__, "is_support_io_optimized", is_support_io_optimized)
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        pulumi.set(__self__, "most_recent", most_recent)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        pulumi.set(__self__, "os_type", os_type)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if owners and not isinstance(owners, str):
            raise TypeError("Expected argument 'owners' to be a str")
        pulumi.set(__self__, "owners", owners)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if usage and not isinstance(usage, str):
            raise TypeError("Expected argument 'usage' to be a str")
        pulumi.set(__self__, "usage", usage)

    @property
    @pulumi.getter(name="actionType")
    def action_type(self) -> Optional[str]:
        return pulumi.get(self, "action_type")

    @property
    @pulumi.getter
    def architecture(self) -> Optional[str]:
        """
        Platform type of the image system: i386 or x86_64.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[bool]:
        return pulumi.get(self, "dry_run")

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
        A list of image IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="imageFamily")
    def image_family(self) -> Optional[str]:
        return pulumi.get(self, "image_family")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[str]:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[str]:
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="imageOwnerId")
    def image_owner_id(self) -> Optional[str]:
        return pulumi.get(self, "image_owner_id")

    @property
    @pulumi.getter
    def images(self) -> Sequence['outputs.GetImagesImageResult']:
        """
        A list of images. Each element contains the following attributes:
        """
        return pulumi.get(self, "images")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="isSupportCloudInit")
    def is_support_cloud_init(self) -> Optional[bool]:
        return pulumi.get(self, "is_support_cloud_init")

    @property
    @pulumi.getter(name="isSupportIoOptimized")
    def is_support_io_optimized(self) -> Optional[bool]:
        return pulumi.get(self, "is_support_io_optimized")

    @property
    @pulumi.getter(name="mostRecent")
    def most_recent(self) -> Optional[bool]:
        return pulumi.get(self, "most_recent")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> Optional[str]:
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def owners(self) -> Optional[str]:
        return pulumi.get(self, "owners")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[str]:
        """
        Snapshot ID.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Status of the image. Possible values: `UnAvailable`, `Available`, `Creating` and `CreateFailed`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def usage(self) -> Optional[str]:
        return pulumi.get(self, "usage")


class AwaitableGetImagesResult(GetImagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImagesResult(
            action_type=self.action_type,
            architecture=self.architecture,
            dry_run=self.dry_run,
            id=self.id,
            ids=self.ids,
            image_family=self.image_family,
            image_id=self.image_id,
            image_name=self.image_name,
            image_owner_id=self.image_owner_id,
            images=self.images,
            instance_type=self.instance_type,
            is_support_cloud_init=self.is_support_cloud_init,
            is_support_io_optimized=self.is_support_io_optimized,
            most_recent=self.most_recent,
            name_regex=self.name_regex,
            os_type=self.os_type,
            output_file=self.output_file,
            owners=self.owners,
            resource_group_id=self.resource_group_id,
            snapshot_id=self.snapshot_id,
            status=self.status,
            tags=self.tags,
            usage=self.usage)


def get_images(action_type: Optional[str] = None,
               architecture: Optional[str] = None,
               dry_run: Optional[bool] = None,
               image_family: Optional[str] = None,
               image_id: Optional[str] = None,
               image_name: Optional[str] = None,
               image_owner_id: Optional[str] = None,
               instance_type: Optional[str] = None,
               is_support_cloud_init: Optional[bool] = None,
               is_support_io_optimized: Optional[bool] = None,
               most_recent: Optional[bool] = None,
               name_regex: Optional[str] = None,
               os_type: Optional[str] = None,
               output_file: Optional[str] = None,
               owners: Optional[str] = None,
               resource_group_id: Optional[str] = None,
               snapshot_id: Optional[str] = None,
               status: Optional[str] = None,
               tags: Optional[Mapping[str, Any]] = None,
               usage: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImagesResult:
    """
    This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud,
    other public images and the ones available on the image market.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    images_ds = alicloud.ecs.get_images(name_regex="^centos_6",
        owners="system")
    pulumi.export("firstImageId", images_ds.images[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str action_type: The scenario in which the image will be used. Default value: `CreateEcs`. Valid values:
    :param str architecture: The image architecture. Valid values: `i386` and `x86_64`.
    :param bool dry_run: Specifies whether the image is running on an ECS instance. Default value: `false`. Valid values:
    :param str image_family: The name of the image family. You can set this parameter to query images of the specified image family. This parameter is empty by default.
    :param str image_id: The ID of the image.
    :param str image_name: The name of the image.
    :param str image_owner_id: The ID of the Alibaba Cloud account to which the image belongs. This parameter takes effect only when you query shared images or community images.
    :param str instance_type: The instance type for which the image can be used.
    :param bool is_support_cloud_init: Specifies whether the image supports cloud-init.
    :param bool is_support_io_optimized: Specifies whether the image can be used on I/O optimized instances.
    :param bool most_recent: If more than one result are returned, select the most recent one.
    :param str name_regex: A regex string to filter resulting images by name.
    :param str os_type: The operating system type of the image. Valid values: `windows` and `linux`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
           
           > **NOTE:** At least one of the `name_regex`, `most_recent` and `owners` must be set.
    :param str owners: Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
    :param str resource_group_id: The ID of the resource group to which the custom image belongs.
    :param str snapshot_id: The ID of the snapshot used to create the custom image.
    :param str status: The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values:
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    :param str usage: Specifies whether to check the validity of the request without actually making the request. Valid values:
    """
    __args__ = dict()
    __args__['actionType'] = action_type
    __args__['architecture'] = architecture
    __args__['dryRun'] = dry_run
    __args__['imageFamily'] = image_family
    __args__['imageId'] = image_id
    __args__['imageName'] = image_name
    __args__['imageOwnerId'] = image_owner_id
    __args__['instanceType'] = instance_type
    __args__['isSupportCloudInit'] = is_support_cloud_init
    __args__['isSupportIoOptimized'] = is_support_io_optimized
    __args__['mostRecent'] = most_recent
    __args__['nameRegex'] = name_regex
    __args__['osType'] = os_type
    __args__['outputFile'] = output_file
    __args__['owners'] = owners
    __args__['resourceGroupId'] = resource_group_id
    __args__['snapshotId'] = snapshot_id
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['usage'] = usage
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getImages:getImages', __args__, opts=opts, typ=GetImagesResult).value

    return AwaitableGetImagesResult(
        action_type=pulumi.get(__ret__, 'action_type'),
        architecture=pulumi.get(__ret__, 'architecture'),
        dry_run=pulumi.get(__ret__, 'dry_run'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        image_family=pulumi.get(__ret__, 'image_family'),
        image_id=pulumi.get(__ret__, 'image_id'),
        image_name=pulumi.get(__ret__, 'image_name'),
        image_owner_id=pulumi.get(__ret__, 'image_owner_id'),
        images=pulumi.get(__ret__, 'images'),
        instance_type=pulumi.get(__ret__, 'instance_type'),
        is_support_cloud_init=pulumi.get(__ret__, 'is_support_cloud_init'),
        is_support_io_optimized=pulumi.get(__ret__, 'is_support_io_optimized'),
        most_recent=pulumi.get(__ret__, 'most_recent'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        os_type=pulumi.get(__ret__, 'os_type'),
        output_file=pulumi.get(__ret__, 'output_file'),
        owners=pulumi.get(__ret__, 'owners'),
        resource_group_id=pulumi.get(__ret__, 'resource_group_id'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        usage=pulumi.get(__ret__, 'usage'))


@_utilities.lift_output_func(get_images)
def get_images_output(action_type: Optional[pulumi.Input[Optional[str]]] = None,
                      architecture: Optional[pulumi.Input[Optional[str]]] = None,
                      dry_run: Optional[pulumi.Input[Optional[bool]]] = None,
                      image_family: Optional[pulumi.Input[Optional[str]]] = None,
                      image_id: Optional[pulumi.Input[Optional[str]]] = None,
                      image_name: Optional[pulumi.Input[Optional[str]]] = None,
                      image_owner_id: Optional[pulumi.Input[Optional[str]]] = None,
                      instance_type: Optional[pulumi.Input[Optional[str]]] = None,
                      is_support_cloud_init: Optional[pulumi.Input[Optional[bool]]] = None,
                      is_support_io_optimized: Optional[pulumi.Input[Optional[bool]]] = None,
                      most_recent: Optional[pulumi.Input[Optional[bool]]] = None,
                      name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                      os_type: Optional[pulumi.Input[Optional[str]]] = None,
                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      owners: Optional[pulumi.Input[Optional[str]]] = None,
                      resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                      snapshot_id: Optional[pulumi.Input[Optional[str]]] = None,
                      status: Optional[pulumi.Input[Optional[str]]] = None,
                      tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                      usage: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImagesResult]:
    """
    This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud,
    other public images and the ones available on the image market.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    images_ds = alicloud.ecs.get_images(name_regex="^centos_6",
        owners="system")
    pulumi.export("firstImageId", images_ds.images[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str action_type: The scenario in which the image will be used. Default value: `CreateEcs`. Valid values:
    :param str architecture: The image architecture. Valid values: `i386` and `x86_64`.
    :param bool dry_run: Specifies whether the image is running on an ECS instance. Default value: `false`. Valid values:
    :param str image_family: The name of the image family. You can set this parameter to query images of the specified image family. This parameter is empty by default.
    :param str image_id: The ID of the image.
    :param str image_name: The name of the image.
    :param str image_owner_id: The ID of the Alibaba Cloud account to which the image belongs. This parameter takes effect only when you query shared images or community images.
    :param str instance_type: The instance type for which the image can be used.
    :param bool is_support_cloud_init: Specifies whether the image supports cloud-init.
    :param bool is_support_io_optimized: Specifies whether the image can be used on I/O optimized instances.
    :param bool most_recent: If more than one result are returned, select the most recent one.
    :param str name_regex: A regex string to filter resulting images by name.
    :param str os_type: The operating system type of the image. Valid values: `windows` and `linux`.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
           
           > **NOTE:** At least one of the `name_regex`, `most_recent` and `owners` must be set.
    :param str owners: Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
    :param str resource_group_id: The ID of the resource group to which the custom image belongs.
    :param str snapshot_id: The ID of the snapshot used to create the custom image.
    :param str status: The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values:
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    :param str usage: Specifies whether to check the validity of the request without actually making the request. Valid values:
    """
    ...
