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
    'GetApisResult',
    'AwaitableGetApisResult',
    'get_apis',
    'get_apis_output',
]

@pulumi.output_type
class GetApisResult:
    """
    A collection of values returned by getApis.
    """
    def __init__(__self__, api_id=None, apis=None, group_id=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if api_id and not isinstance(api_id, str):
            raise TypeError("Expected argument 'api_id' to be a str")
        pulumi.set(__self__, "api_id", api_id)
        if apis and not isinstance(apis, list):
            raise TypeError("Expected argument 'apis' to be a list")
        pulumi.set(__self__, "apis", apis)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
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
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[str]:
        warnings.warn("""Field 'api_id' has been deprecated from provider version 1.52.2. New field 'ids' replaces it.""", DeprecationWarning)
        pulumi.log.warn("""api_id is deprecated: Field 'api_id' has been deprecated from provider version 1.52.2. New field 'ids' replaces it.""")

        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter
    def apis(self) -> Sequence['outputs.GetApisApiResult']:
        """
        A list of apis. Each element contains the following attributes:
        """
        return pulumi.get(self, "apis")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[str]:
        """
        The group id that the apis belong to.
        """
        return pulumi.get(self, "group_id")

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
        A list of api IDs.
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
        A list of api names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")


class AwaitableGetApisResult(GetApisResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApisResult(
            api_id=self.api_id,
            apis=self.apis,
            group_id=self.group_id,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)


def get_apis(api_id: Optional[str] = None,
             group_id: Optional[str] = None,
             ids: Optional[Sequence[str]] = None,
             name_regex: Optional[str] = None,
             output_file: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApisResult:
    """
    This data source provides the apis of the current Alibaba Cloud user.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    data_apigatway_apis = alicloud.apigateway.get_apis(output_file="output_ApiGatawayApis")
    pulumi.export("firstApiId", data_apigatway["apis"][0]["id"])
    ```
    <!--End PulumiCodeChooser -->


    :param str api_id: (It has been deprecated from version 1.52.2, and use field 'ids' to replace.) ID of the specified API.
    :param str group_id: ID of the specified group.
    :param Sequence[str] ids: A list of api IDs.
    :param str name_regex: A regex string to filter api gateway apis by name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['groupId'] = group_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:apigateway/getApis:getApis', __args__, opts=opts, typ=GetApisResult).value

    return AwaitableGetApisResult(
        api_id=pulumi.get(__ret__, 'api_id'),
        apis=pulumi.get(__ret__, 'apis'),
        group_id=pulumi.get(__ret__, 'group_id'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'))


@_utilities.lift_output_func(get_apis)
def get_apis_output(api_id: Optional[pulumi.Input[Optional[str]]] = None,
                    group_id: Optional[pulumi.Input[Optional[str]]] = None,
                    ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApisResult]:
    """
    This data source provides the apis of the current Alibaba Cloud user.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    data_apigatway_apis = alicloud.apigateway.get_apis(output_file="output_ApiGatawayApis")
    pulumi.export("firstApiId", data_apigatway["apis"][0]["id"])
    ```
    <!--End PulumiCodeChooser -->


    :param str api_id: (It has been deprecated from version 1.52.2, and use field 'ids' to replace.) ID of the specified API.
    :param str group_id: ID of the specified group.
    :param Sequence[str] ids: A list of api IDs.
    :param str name_regex: A regex string to filter api gateway apis by name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
