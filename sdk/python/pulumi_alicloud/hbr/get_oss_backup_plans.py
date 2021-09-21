# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetOssBackupPlansResult',
    'AwaitableGetOssBackupPlansResult',
    'get_oss_backup_plans',
]

@pulumi.output_type
class GetOssBackupPlansResult:
    """
    A collection of values returned by getOssBackupPlans.
    """
    def __init__(__self__, bucket=None, id=None, ids=None, name_regex=None, names=None, output_file=None, plans=None, vault_id=None):
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        pulumi.set(__self__, "bucket", bucket)
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
        if plans and not isinstance(plans, list):
            raise TypeError("Expected argument 'plans' to be a list")
        pulumi.set(__self__, "plans", plans)
        if vault_id and not isinstance(vault_id, str):
            raise TypeError("Expected argument 'vault_id' to be a str")
        pulumi.set(__self__, "vault_id", vault_id)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[str]:
        return pulumi.get(self, "bucket")

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
    @pulumi.getter
    def plans(self) -> Sequence['outputs.GetOssBackupPlansPlanResult']:
        return pulumi.get(self, "plans")

    @property
    @pulumi.getter(name="vaultId")
    def vault_id(self) -> Optional[str]:
        return pulumi.get(self, "vault_id")


class AwaitableGetOssBackupPlansResult(GetOssBackupPlansResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOssBackupPlansResult(
            bucket=self.bucket,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            plans=self.plans,
            vault_id=self.vault_id)


def get_oss_backup_plans(bucket: Optional[str] = None,
                         ids: Optional[Sequence[str]] = None,
                         name_regex: Optional[str] = None,
                         output_file: Optional[str] = None,
                         vault_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOssBackupPlansResult:
    """
    This data source provides the Hbr OssBackupPlans of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.131.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.hbr.get_oss_backup_plans(name_regex="^my-OssBackupPlan")
    pulumi.export("hbrOssBackupPlanId", ids.plans[0].id)
    ```


    :param str bucket: The name of OSS bucket.
    :param Sequence[str] ids: A list of OssBackupPlan IDs.
    :param str name_regex: A regex string to filter results by OssBackupPlan name.
    :param str vault_id: The ID of backup vault.
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['vaultId'] = vault_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:hbr/getOssBackupPlans:getOssBackupPlans', __args__, opts=opts, typ=GetOssBackupPlansResult).value

    return AwaitableGetOssBackupPlansResult(
        bucket=__ret__.bucket,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        plans=__ret__.plans,
        vault_id=__ret__.vault_id)
