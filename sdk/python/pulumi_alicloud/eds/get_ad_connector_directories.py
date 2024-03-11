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
    'GetAdConnectorDirectoriesResult',
    'AwaitableGetAdConnectorDirectoriesResult',
    'get_ad_connector_directories',
    'get_ad_connector_directories_output',
]

@pulumi.output_type
class GetAdConnectorDirectoriesResult:
    """
    A collection of values returned by getAdConnectorDirectories.
    """
    def __init__(__self__, directories=None, id=None, ids=None, name_regex=None, names=None, output_file=None, status=None):
        if directories and not isinstance(directories, list):
            raise TypeError("Expected argument 'directories' to be a list")
        pulumi.set(__self__, "directories", directories)
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
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def directories(self) -> Sequence['outputs.GetAdConnectorDirectoriesDirectoryResult']:
        return pulumi.get(self, "directories")

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
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetAdConnectorDirectoriesResult(GetAdConnectorDirectoriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAdConnectorDirectoriesResult(
            directories=self.directories,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status)


def get_ad_connector_directories(ids: Optional[Sequence[str]] = None,
                                 name_regex: Optional[str] = None,
                                 output_file: Optional[str] = None,
                                 status: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAdConnectorDirectoriesResult:
    """
    This data source provides the Ecd Ad Connector Directories of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.174.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.eds.get_ad_connector_directories(ids=["example_id"])
    pulumi.export("ecdAdConnectorDirectoryId1", ids.directories[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Ad Connector Directory IDs.
    :param str name_regex: A regex string to filter results by Ad Connector Directory name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of directory.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:eds/getAdConnectorDirectories:getAdConnectorDirectories', __args__, opts=opts, typ=GetAdConnectorDirectoriesResult).value

    return AwaitableGetAdConnectorDirectoriesResult(
        directories=pulumi.get(__ret__, 'directories'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_ad_connector_directories)
def get_ad_connector_directories_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                        name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                        status: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAdConnectorDirectoriesResult]:
    """
    This data source provides the Ecd Ad Connector Directories of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.174.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.eds.get_ad_connector_directories(ids=["example_id"])
    pulumi.export("ecdAdConnectorDirectoryId1", ids.directories[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Ad Connector Directory IDs.
    :param str name_regex: A regex string to filter results by Ad Connector Directory name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of directory.
    """
    ...
