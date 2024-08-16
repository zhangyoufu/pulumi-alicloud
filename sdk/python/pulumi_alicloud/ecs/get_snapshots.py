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
    'GetSnapshotsResult',
    'AwaitableGetSnapshotsResult',
    'get_snapshots',
    'get_snapshots_output',
]

@pulumi.output_type
class GetSnapshotsResult:
    """
    A collection of values returned by getSnapshots.
    """
    def __init__(__self__, category=None, dry_run=None, encrypted=None, id=None, ids=None, kms_key_id=None, name_regex=None, names=None, output_file=None, resource_group_id=None, snapshot_link_id=None, snapshot_name=None, snapshot_type=None, snapshots=None, source_disk_type=None, status=None, tags=None, type=None, usage=None):
        if category and not isinstance(category, str):
            raise TypeError("Expected argument 'category' to be a str")
        pulumi.set(__self__, "category", category)
        if dry_run and not isinstance(dry_run, bool):
            raise TypeError("Expected argument 'dry_run' to be a bool")
        pulumi.set(__self__, "dry_run", dry_run)
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        pulumi.set(__self__, "encrypted", encrypted)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if snapshot_link_id and not isinstance(snapshot_link_id, str):
            raise TypeError("Expected argument 'snapshot_link_id' to be a str")
        pulumi.set(__self__, "snapshot_link_id", snapshot_link_id)
        if snapshot_name and not isinstance(snapshot_name, str):
            raise TypeError("Expected argument 'snapshot_name' to be a str")
        pulumi.set(__self__, "snapshot_name", snapshot_name)
        if snapshot_type and not isinstance(snapshot_type, str):
            raise TypeError("Expected argument 'snapshot_type' to be a str")
        pulumi.set(__self__, "snapshot_type", snapshot_type)
        if snapshots and not isinstance(snapshots, list):
            raise TypeError("Expected argument 'snapshots' to be a list")
        pulumi.set(__self__, "snapshots", snapshots)
        if source_disk_type and not isinstance(source_disk_type, str):
            raise TypeError("Expected argument 'source_disk_type' to be a str")
        pulumi.set(__self__, "source_disk_type", source_disk_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if usage and not isinstance(usage, str):
            raise TypeError("Expected argument 'usage' to be a str")
        pulumi.set(__self__, "usage", usage)

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[bool]:
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter
    def encrypted(self) -> Optional[bool]:
        """
        Whether the snapshot is encrypted or not.
        """
        return pulumi.get(self, "encrypted")

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
        A list of snapshot IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of snapshots names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="snapshotLinkId")
    def snapshot_link_id(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_link_id")

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_name")

    @property
    @pulumi.getter(name="snapshotType")
    def snapshot_type(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_type")

    @property
    @pulumi.getter
    def snapshots(self) -> Sequence['outputs.GetSnapshotsSnapshotResult']:
        """
        A list of snapshots. Each element contains the following attributes:
        """
        return pulumi.get(self, "snapshots")

    @property
    @pulumi.getter(name="sourceDiskType")
    def source_disk_type(self) -> Optional[str]:
        """
        Source disk attribute. Value range: `System`,`Data`.
        """
        return pulumi.get(self, "source_disk_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The snapshot status. Value range: `progressing`, `accomplished` and `failed`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags assigned to the snapshot.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def usage(self) -> Optional[str]:
        """
        Whether the snapshots are used to create resources or not. Value range: `image`, `disk`, `image_disk` and `none`.
        """
        return pulumi.get(self, "usage")


class AwaitableGetSnapshotsResult(GetSnapshotsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotsResult(
            category=self.category,
            dry_run=self.dry_run,
            encrypted=self.encrypted,
            id=self.id,
            ids=self.ids,
            kms_key_id=self.kms_key_id,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            resource_group_id=self.resource_group_id,
            snapshot_link_id=self.snapshot_link_id,
            snapshot_name=self.snapshot_name,
            snapshot_type=self.snapshot_type,
            snapshots=self.snapshots,
            source_disk_type=self.source_disk_type,
            status=self.status,
            tags=self.tags,
            type=self.type,
            usage=self.usage)


def get_snapshots(category: Optional[str] = None,
                  dry_run: Optional[bool] = None,
                  encrypted: Optional[bool] = None,
                  ids: Optional[Sequence[str]] = None,
                  kms_key_id: Optional[str] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  resource_group_id: Optional[str] = None,
                  snapshot_link_id: Optional[str] = None,
                  snapshot_name: Optional[str] = None,
                  snapshot_type: Optional[str] = None,
                  source_disk_type: Optional[str] = None,
                  status: Optional[str] = None,
                  tags: Optional[Mapping[str, str]] = None,
                  type: Optional[str] = None,
                  usage: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSnapshotsResult:
    """
    > **DEPRECATED:** This datasource has been renamed to ecs_get_ecs_snapshots from version 1.120.0.

    Use this data source to get a list of snapshot according to the specified filters in an Alibaba Cloud account.

    For information about snapshot and how to use it, see [Snapshot](https://www.alibabacloud.com/help/doc-detail/25460.html).

    > **NOTE:**  Available in 1.40.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    snapshots = alicloud.ecs.get_snapshots(ids=["s-123456890abcdef"],
        name_regex="tf-testAcc-snapshot")
    ```

    ## Argument Reference

    The following arguments are supported:

    * `instance_id` - (Optional) The specified instance ID.
    * `disk_id` - (Optional) The specified disk ID.
    * `encrypted` - (Optional) Queries the encrypted snapshots. Optional values: `true`: Encrypted snapshots. `false`: No encryption attribute limit. Default value: `false`.
    * `ids` - (Optional)  A list of snapshot IDs.
    * `name_regex` - (Optional) A regex string to filter results by snapshot name.
    * `status` - (Optional) The specified snapshot status. Default value: `all`. Optional values:
      * progressing: The snapshots are being created.
      * accomplished: The snapshots are ready to use.
      * failed: The snapshot creation failed.
      * all: All status.
    * `type` - (Optional) The snapshot category. Default value: `all`. Optional values:
      * auto: Auto snapshots.
      * user: Manual snapshots.
      * all: Auto and manual snapshots.
    * `source_disk_type` - (Optional) The type of source disk:
      * System: The snapshots are created for system disks.
      * Data: The snapshots are created for data disks.
    * `usage` - (Optional) The usage of the snapshot:
      * image: The snapshots are used to create custom images.
      * disk: The snapshots are used to CreateDisk.
      * mage_disk: The snapshots are used to create custom images and data disks.
      * none: The snapshots are not used yet.
    * `tags` - (Optional) A map of tags assigned to snapshots.
    * `output_file` - (Optional) The name of output file that saves the filter results.


    :param bool encrypted: Whether the snapshot is encrypted or not.
    :param Sequence[str] ids: A list of snapshot IDs.
    :param str source_disk_type: Source disk attribute. Value range: `System`,`Data`.
    :param str status: The snapshot status. Value range: `progressing`, `accomplished` and `failed`.
    :param Mapping[str, str] tags: A map of tags assigned to the snapshot.
    :param str usage: Whether the snapshots are used to create resources or not. Value range: `image`, `disk`, `image_disk` and `none`.
    """
    __args__ = dict()
    __args__['category'] = category
    __args__['dryRun'] = dry_run
    __args__['encrypted'] = encrypted
    __args__['ids'] = ids
    __args__['kmsKeyId'] = kms_key_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['resourceGroupId'] = resource_group_id
    __args__['snapshotLinkId'] = snapshot_link_id
    __args__['snapshotName'] = snapshot_name
    __args__['snapshotType'] = snapshot_type
    __args__['sourceDiskType'] = source_disk_type
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['type'] = type
    __args__['usage'] = usage
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ecs/getSnapshots:getSnapshots', __args__, opts=opts, typ=GetSnapshotsResult).value

    return AwaitableGetSnapshotsResult(
        category=pulumi.get(__ret__, 'category'),
        dry_run=pulumi.get(__ret__, 'dry_run'),
        encrypted=pulumi.get(__ret__, 'encrypted'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        kms_key_id=pulumi.get(__ret__, 'kms_key_id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        resource_group_id=pulumi.get(__ret__, 'resource_group_id'),
        snapshot_link_id=pulumi.get(__ret__, 'snapshot_link_id'),
        snapshot_name=pulumi.get(__ret__, 'snapshot_name'),
        snapshot_type=pulumi.get(__ret__, 'snapshot_type'),
        snapshots=pulumi.get(__ret__, 'snapshots'),
        source_disk_type=pulumi.get(__ret__, 'source_disk_type'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        usage=pulumi.get(__ret__, 'usage'))


@_utilities.lift_output_func(get_snapshots)
def get_snapshots_output(category: Optional[pulumi.Input[Optional[str]]] = None,
                         dry_run: Optional[pulumi.Input[Optional[bool]]] = None,
                         encrypted: Optional[pulumi.Input[Optional[bool]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         kms_key_id: Optional[pulumi.Input[Optional[str]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                         snapshot_link_id: Optional[pulumi.Input[Optional[str]]] = None,
                         snapshot_name: Optional[pulumi.Input[Optional[str]]] = None,
                         snapshot_type: Optional[pulumi.Input[Optional[str]]] = None,
                         source_disk_type: Optional[pulumi.Input[Optional[str]]] = None,
                         status: Optional[pulumi.Input[Optional[str]]] = None,
                         tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                         type: Optional[pulumi.Input[Optional[str]]] = None,
                         usage: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSnapshotsResult]:
    """
    > **DEPRECATED:** This datasource has been renamed to ecs_get_ecs_snapshots from version 1.120.0.

    Use this data source to get a list of snapshot according to the specified filters in an Alibaba Cloud account.

    For information about snapshot and how to use it, see [Snapshot](https://www.alibabacloud.com/help/doc-detail/25460.html).

    > **NOTE:**  Available in 1.40.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    snapshots = alicloud.ecs.get_snapshots(ids=["s-123456890abcdef"],
        name_regex="tf-testAcc-snapshot")
    ```

    ## Argument Reference

    The following arguments are supported:

    * `instance_id` - (Optional) The specified instance ID.
    * `disk_id` - (Optional) The specified disk ID.
    * `encrypted` - (Optional) Queries the encrypted snapshots. Optional values: `true`: Encrypted snapshots. `false`: No encryption attribute limit. Default value: `false`.
    * `ids` - (Optional)  A list of snapshot IDs.
    * `name_regex` - (Optional) A regex string to filter results by snapshot name.
    * `status` - (Optional) The specified snapshot status. Default value: `all`. Optional values:
      * progressing: The snapshots are being created.
      * accomplished: The snapshots are ready to use.
      * failed: The snapshot creation failed.
      * all: All status.
    * `type` - (Optional) The snapshot category. Default value: `all`. Optional values:
      * auto: Auto snapshots.
      * user: Manual snapshots.
      * all: Auto and manual snapshots.
    * `source_disk_type` - (Optional) The type of source disk:
      * System: The snapshots are created for system disks.
      * Data: The snapshots are created for data disks.
    * `usage` - (Optional) The usage of the snapshot:
      * image: The snapshots are used to create custom images.
      * disk: The snapshots are used to CreateDisk.
      * mage_disk: The snapshots are used to create custom images and data disks.
      * none: The snapshots are not used yet.
    * `tags` - (Optional) A map of tags assigned to snapshots.
    * `output_file` - (Optional) The name of output file that saves the filter results.


    :param bool encrypted: Whether the snapshot is encrypted or not.
    :param Sequence[str] ids: A list of snapshot IDs.
    :param str source_disk_type: Source disk attribute. Value range: `System`,`Data`.
    :param str status: The snapshot status. Value range: `progressing`, `accomplished` and `failed`.
    :param Mapping[str, str] tags: A map of tags assigned to the snapshot.
    :param str usage: Whether the snapshots are used to create resources or not. Value range: `image`, `disk`, `image_disk` and `none`.
    """
    ...
