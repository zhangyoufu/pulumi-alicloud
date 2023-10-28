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
    'GetZnodesResult',
    'AwaitableGetZnodesResult',
    'get_znodes',
    'get_znodes_output',
]

@pulumi.output_type
class GetZnodesResult:
    """
    A collection of values returned by getZnodes.
    """
    def __init__(__self__, accept_language=None, cluster_id=None, id=None, ids=None, name_regex=None, names=None, output_file=None, path=None, znodes=None):
        if accept_language and not isinstance(accept_language, str):
            raise TypeError("Expected argument 'accept_language' to be a str")
        pulumi.set(__self__, "accept_language", accept_language)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
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
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if znodes and not isinstance(znodes, list):
            raise TypeError("Expected argument 'znodes' to be a list")
        pulumi.set(__self__, "znodes", znodes)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> Optional[str]:
        return pulumi.get(self, "accept_language")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

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
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def znodes(self) -> Sequence['outputs.GetZnodesZnodeResult']:
        return pulumi.get(self, "znodes")


class AwaitableGetZnodesResult(GetZnodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZnodesResult(
            accept_language=self.accept_language,
            cluster_id=self.cluster_id,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            path=self.path,
            znodes=self.znodes)


def get_znodes(accept_language: Optional[str] = None,
               cluster_id: Optional[str] = None,
               ids: Optional[Sequence[str]] = None,
               name_regex: Optional[str] = None,
               output_file: Optional[str] = None,
               path: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZnodesResult:
    """
    This data source provides the Mse Znodes of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.162.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.mse.get_znodes(cluster_id="example_value",
        path="/",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("mseZnodeId1", ids.znodes[0].id)
    name_regex = alicloud.mse.get_znodes(path="/",
        cluster_id="example_value",
        name_regex="^my-Znode")
    pulumi.export("mseZnodeId2", name_regex.znodes[0].id)
    ```


    :param str accept_language: The language type of the returned information. Valid values: `zh` or `en`.
    :param str cluster_id: The ID of the Cluster.
    :param Sequence[str] ids: A list of Znode IDs.
    :param str name_regex: A regex string to filter results by Znode name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str path: The Node path.
    """
    __args__ = dict()
    __args__['acceptLanguage'] = accept_language
    __args__['clusterId'] = cluster_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['path'] = path
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:mse/getZnodes:getZnodes', __args__, opts=opts, typ=GetZnodesResult).value

    return AwaitableGetZnodesResult(
        accept_language=pulumi.get(__ret__, 'accept_language'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        path=pulumi.get(__ret__, 'path'),
        znodes=pulumi.get(__ret__, 'znodes'))


@_utilities.lift_output_func(get_znodes)
def get_znodes_output(accept_language: Optional[pulumi.Input[Optional[str]]] = None,
                      cluster_id: Optional[pulumi.Input[str]] = None,
                      ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      path: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZnodesResult]:
    """
    This data source provides the Mse Znodes of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.162.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.mse.get_znodes(cluster_id="example_value",
        path="/",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("mseZnodeId1", ids.znodes[0].id)
    name_regex = alicloud.mse.get_znodes(path="/",
        cluster_id="example_value",
        name_regex="^my-Znode")
    pulumi.export("mseZnodeId2", name_regex.znodes[0].id)
    ```


    :param str accept_language: The language type of the returned information. Valid values: `zh` or `en`.
    :param str cluster_id: The ID of the Cluster.
    :param Sequence[str] ids: A list of Znode IDs.
    :param str name_regex: A regex string to filter results by Znode name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str path: The Node path.
    """
    ...
