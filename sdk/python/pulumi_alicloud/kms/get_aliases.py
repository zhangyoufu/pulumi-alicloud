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
    'GetAliasesResult',
    'AwaitableGetAliasesResult',
    'get_aliases',
    'get_aliases_output',
]

@pulumi.output_type
class GetAliasesResult:
    """
    A collection of values returned by getAliases.
    """
    def __init__(__self__, aliases=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if aliases and not isinstance(aliases, list):
            raise TypeError("Expected argument 'aliases' to be a list")
        pulumi.set(__self__, "aliases", aliases)
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

    @property
    @pulumi.getter
    def aliases(self) -> Sequence['outputs.GetAliasesAliasResult']:
        """
        A list of KMS User alias. Each element contains the following attributes:
        """
        return pulumi.get(self, "aliases")

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
        A list of kms aliases IDs. The value is same as KMS alias_name.
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
        A list of KMS alias name.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetAliasesResult(GetAliasesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAliasesResult(
            aliases=self.aliases,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)


def get_aliases(ids: Optional[Sequence[str]] = None,
                name_regex: Optional[str] = None,
                output_file: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAliasesResult:
    """
    This data source provides a list of KMS aliases in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in v1.79.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    kms_aliases = alicloud.kms.get_aliases(ids=["d89e8a53-b708-41aa-8c67-6873axxx"],
        name_regex="alias/tf-testKmsAlias_123")
    pulumi.export("firstKeyId", data["alicloud_kms_keys"]["kms_keys_ds"]["keys"][0]["id"])
    ```


    :param Sequence[str] ids: A list of KMS aliases IDs. The value is same as KMS alias_name.
    :param str name_regex: A regex string to filter the results by the KMS alias name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:kms/getAliases:getAliases', __args__, opts=opts, typ=GetAliasesResult).value

    return AwaitableGetAliasesResult(
        aliases=pulumi.get(__ret__, 'aliases'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'))


@_utilities.lift_output_func(get_aliases)
def get_aliases_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAliasesResult]:
    """
    This data source provides a list of KMS aliases in an Alibaba Cloud account according to the specified filters.

    > **NOTE:** Available in v1.79.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    kms_aliases = alicloud.kms.get_aliases(ids=["d89e8a53-b708-41aa-8c67-6873axxx"],
        name_regex="alias/tf-testKmsAlias_123")
    pulumi.export("firstKeyId", data["alicloud_kms_keys"]["kms_keys_ds"]["keys"][0]["id"])
    ```


    :param Sequence[str] ids: A list of KMS aliases IDs. The value is same as KMS alias_name.
    :param str name_regex: A regex string to filter the results by the KMS alias name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
