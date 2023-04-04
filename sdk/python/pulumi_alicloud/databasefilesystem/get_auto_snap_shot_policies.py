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
    'GetAutoSnapShotPoliciesResult',
    'AwaitableGetAutoSnapShotPoliciesResult',
    'get_auto_snap_shot_policies',
    'get_auto_snap_shot_policies_output',
]

@pulumi.output_type
class GetAutoSnapShotPoliciesResult:
    """
    A collection of values returned by getAutoSnapShotPolicies.
    """
    def __init__(__self__, auto_snap_shot_policies=None, id=None, ids=None, name_regex=None, names=None, output_file=None, page_number=None, page_size=None):
        if auto_snap_shot_policies and not isinstance(auto_snap_shot_policies, list):
            raise TypeError("Expected argument 'auto_snap_shot_policies' to be a list")
        pulumi.set(__self__, "auto_snap_shot_policies", auto_snap_shot_policies)
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

    @property
    @pulumi.getter(name="autoSnapShotPolicies")
    def auto_snap_shot_policies(self) -> Sequence['outputs.GetAutoSnapShotPoliciesAutoSnapShotPolicyResult']:
        """
        A list of Auto Snap Shot Policy Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "auto_snap_shot_policies")

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
        A list of Auto Snap Shot Policy IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of Auto Snap Shot Policy names.
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


class AwaitableGetAutoSnapShotPoliciesResult(GetAutoSnapShotPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAutoSnapShotPoliciesResult(
            auto_snap_shot_policies=self.auto_snap_shot_policies,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size)


def get_auto_snap_shot_policies(ids: Optional[Sequence[str]] = None,
                                name_regex: Optional[str] = None,
                                output_file: Optional[str] = None,
                                page_number: Optional[int] = None,
                                page_size: Optional[int] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAutoSnapShotPoliciesResult:
    """
    This data source provides Dbfs Auto Snap Shot Policy available to the user.[What is Auto Snap Shot Policy](https://help.aliyun.com/document_detail/469597.html)

    > **NOTE:** Available in 1.202.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.databasefilesystem.get_auto_snap_shot_policies(ids=[alicloud_dbfs_auto_snap_shot_policy["default"]["id"]])
    pulumi.export("alicloudDbfsAutoSnapShotPolicyExampleId", default.auto_snap_shot_policies[0].id)
    ```


    :param Sequence[str] ids: A list of Auto Snap Shot Policy IDs.
    :param str name_regex: A regex string to filter results by Auto Snap Shot Policy name.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:databasefilesystem/getAutoSnapShotPolicies:getAutoSnapShotPolicies', __args__, opts=opts, typ=GetAutoSnapShotPoliciesResult).value

    return AwaitableGetAutoSnapShotPoliciesResult(
        auto_snap_shot_policies=__ret__.auto_snap_shot_policies,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        page_number=__ret__.page_number,
        page_size=__ret__.page_size)


@_utilities.lift_output_func(get_auto_snap_shot_policies)
def get_auto_snap_shot_policies_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                       name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                       page_number: Optional[pulumi.Input[Optional[int]]] = None,
                                       page_size: Optional[pulumi.Input[Optional[int]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAutoSnapShotPoliciesResult]:
    """
    This data source provides Dbfs Auto Snap Shot Policy available to the user.[What is Auto Snap Shot Policy](https://help.aliyun.com/document_detail/469597.html)

    > **NOTE:** Available in 1.202.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default = alicloud.databasefilesystem.get_auto_snap_shot_policies(ids=[alicloud_dbfs_auto_snap_shot_policy["default"]["id"]])
    pulumi.export("alicloudDbfsAutoSnapShotPolicyExampleId", default.auto_snap_shot_policies[0].id)
    ```


    :param Sequence[str] ids: A list of Auto Snap Shot Policy IDs.
    :param str name_regex: A regex string to filter results by Auto Snap Shot Policy name.
    """
    ...
