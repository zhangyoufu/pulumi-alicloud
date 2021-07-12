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
    'GetReceiversResult',
    'AwaitableGetReceiversResult',
    'get_receivers',
]

@pulumi.output_type
class GetReceiversResult:
    """
    A collection of values returned by getReceivers.
    """
    def __init__(__self__, id=None, ids=None, key_word=None, name_regex=None, names=None, output_file=None, receiverses=None, status=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if key_word and not isinstance(key_word, str):
            raise TypeError("Expected argument 'key_word' to be a str")
        pulumi.set(__self__, "key_word", key_word)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if receiverses and not isinstance(receiverses, list):
            raise TypeError("Expected argument 'receiverses' to be a list")
        pulumi.set(__self__, "receiverses", receiverses)
        if status and not isinstance(status, int):
            raise TypeError("Expected argument 'status' to be a int")
        pulumi.set(__self__, "status", status)

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
    @pulumi.getter(name="keyWord")
    def key_word(self) -> Optional[str]:
        return pulumi.get(self, "key_word")

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
    def receiverses(self) -> Sequence['outputs.GetReceiversReceiverseResult']:
        return pulumi.get(self, "receiverses")

    @property
    @pulumi.getter
    def status(self) -> Optional[int]:
        return pulumi.get(self, "status")


class AwaitableGetReceiversResult(GetReceiversResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReceiversResult(
            id=self.id,
            ids=self.ids,
            key_word=self.key_word,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            receiverses=self.receiverses,
            status=self.status)


def get_receivers(ids: Optional[Sequence[str]] = None,
                  key_word: Optional[str] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  status: Optional[int] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReceiversResult:
    """
    This data source provides the Direct Mail Receiverses of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.125.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.directmail.get_receivers(ids=["ca73b1e4fb0df7c935a5097a****"],
        name_regex="the_resource_name")
    pulumi.export("firstDirectMailReceiversId", example.receiverses[0].id)
    ```


    :param Sequence[str] ids: A list of Receivers IDs.
    :param str key_word: The key word.
    :param str name_regex: A regex string to filter results by Receivers name.
    :param int status: The status of the resource.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['keyWord'] = key_word
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:directmail/getReceivers:getReceivers', __args__, opts=opts, typ=GetReceiversResult).value

    return AwaitableGetReceiversResult(
        id=__ret__.id,
        ids=__ret__.ids,
        key_word=__ret__.key_word,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        receiverses=__ret__.receiverses,
        status=__ret__.status)
