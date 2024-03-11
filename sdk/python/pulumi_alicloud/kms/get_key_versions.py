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
    'GetKeyVersionsResult',
    'AwaitableGetKeyVersionsResult',
    'get_key_versions',
    'get_key_versions_output',
]

@pulumi.output_type
class GetKeyVersionsResult:
    """
    A collection of values returned by getKeyVersions.
    """
    def __init__(__self__, id=None, ids=None, key_id=None, output_file=None, versions=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if key_id and not isinstance(key_id, str):
            raise TypeError("Expected argument 'key_id' to be a str")
        pulumi.set(__self__, "key_id", key_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        pulumi.set(__self__, "versions", versions)

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
        A list of KMS KeyVersion IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        """
        ID of the key.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def versions(self) -> Sequence['outputs.GetKeyVersionsVersionResult']:
        """
        A list of KMS KeyVersions. Each element contains the following attributes:
        """
        return pulumi.get(self, "versions")


class AwaitableGetKeyVersionsResult(GetKeyVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyVersionsResult(
            id=self.id,
            ids=self.ids,
            key_id=self.key_id,
            output_file=self.output_file,
            versions=self.versions)


def get_key_versions(ids: Optional[Sequence[str]] = None,
                     key_id: Optional[str] = None,
                     output_file: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeyVersionsResult:
    """
    This data source provides a list of KMS KeyVersions in an Alibaba Cloud account according to the specified filters.

    > NOTE: Available in v1.85.0+

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    alicloud_kms_key_versions_ds = alicloud.kms.get_key_versions(ids=["d89e8a53-b708-41aa-8c67-6873axxx"],
        key_id="08438c-b4d5-4d05-928c-07b7xxxx")
    pulumi.export("allVersions", alicloud_kms_key_versions_ds.versions)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of KMS KeyVersion IDs.
    :param str key_id: The id of kms key.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['keyId'] = key_id
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:kms/getKeyVersions:getKeyVersions', __args__, opts=opts, typ=GetKeyVersionsResult).value

    return AwaitableGetKeyVersionsResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        key_id=pulumi.get(__ret__, 'key_id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        versions=pulumi.get(__ret__, 'versions'))


@_utilities.lift_output_func(get_key_versions)
def get_key_versions_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            key_id: Optional[pulumi.Input[str]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKeyVersionsResult]:
    """
    This data source provides a list of KMS KeyVersions in an Alibaba Cloud account according to the specified filters.

    > NOTE: Available in v1.85.0+

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    alicloud_kms_key_versions_ds = alicloud.kms.get_key_versions(ids=["d89e8a53-b708-41aa-8c67-6873axxx"],
        key_id="08438c-b4d5-4d05-928c-07b7xxxx")
    pulumi.export("allVersions", alicloud_kms_key_versions_ds.versions)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of KMS KeyVersion IDs.
    :param str key_id: The id of kms key.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
