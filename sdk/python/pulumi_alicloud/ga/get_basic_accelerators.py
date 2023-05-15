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
    'GetBasicAcceleratorsResult',
    'AwaitableGetBasicAcceleratorsResult',
    'get_basic_accelerators',
    'get_basic_accelerators_output',
]

@pulumi.output_type
class GetBasicAcceleratorsResult:
    """
    A collection of values returned by getBasicAccelerators.
    """
    def __init__(__self__, accelerator_id=None, accelerators=None, id=None, ids=None, name_regex=None, names=None, output_file=None, page_number=None, page_size=None, status=None):
        if accelerator_id and not isinstance(accelerator_id, str):
            raise TypeError("Expected argument 'accelerator_id' to be a str")
        pulumi.set(__self__, "accelerator_id", accelerator_id)
        if accelerators and not isinstance(accelerators, list):
            raise TypeError("Expected argument 'accelerators' to be a list")
        pulumi.set(__self__, "accelerators", accelerators)
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
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="acceleratorId")
    def accelerator_id(self) -> Optional[str]:
        return pulumi.get(self, "accelerator_id")

    @property
    @pulumi.getter
    def accelerators(self) -> Sequence['outputs.GetBasicAcceleratorsAcceleratorResult']:
        """
        A list of Global Accelerator Basic Accelerators. Each element contains the following attributes:
        """
        return pulumi.get(self, "accelerators")

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
        """
        A list of Global Accelerator Basic Accelerator names.
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

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the Global Accelerator Basic Accelerator instance.
        """
        return pulumi.get(self, "status")


class AwaitableGetBasicAcceleratorsResult(GetBasicAcceleratorsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBasicAcceleratorsResult(
            accelerator_id=self.accelerator_id,
            accelerators=self.accelerators,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size,
            status=self.status)


def get_basic_accelerators(accelerator_id: Optional[str] = None,
                           ids: Optional[Sequence[str]] = None,
                           name_regex: Optional[str] = None,
                           output_file: Optional[str] = None,
                           page_number: Optional[int] = None,
                           page_size: Optional[int] = None,
                           status: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBasicAcceleratorsResult:
    """
    This data source provides the Global Accelerator (GA) Basic Accelerators of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.194.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ga.get_basic_accelerators(ids=["example_id"])
    pulumi.export("gaBasicAcceleratorId1", ids.accelerators[0].id)
    name_regex = alicloud.ga.get_basic_accelerators(name_regex="tf-example")
    pulumi.export("gaBasicAcceleratorId2", name_regex.accelerators[0].id)
    ```


    :param str accelerator_id: The ID of the Global Accelerator Basic Accelerator instance.
    :param Sequence[str] ids: A list of Global Accelerator Basic Accelerator IDs.
    :param str name_regex: A regex string to filter results by Global Accelerator Basic Accelerator name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the Global Accelerator Basic Accelerator instance. Valid Value: `init`, `active`, `configuring`, `binding`, `unbinding`, `deleting`, `finacialLocked`.
    """
    __args__ = dict()
    __args__['acceleratorId'] = accelerator_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:ga/getBasicAccelerators:getBasicAccelerators', __args__, opts=opts, typ=GetBasicAcceleratorsResult).value

    return AwaitableGetBasicAcceleratorsResult(
        accelerator_id=__ret__.accelerator_id,
        accelerators=__ret__.accelerators,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        page_number=__ret__.page_number,
        page_size=__ret__.page_size,
        status=__ret__.status)


@_utilities.lift_output_func(get_basic_accelerators)
def get_basic_accelerators_output(accelerator_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  page_number: Optional[pulumi.Input[Optional[int]]] = None,
                                  page_size: Optional[pulumi.Input[Optional[int]]] = None,
                                  status: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBasicAcceleratorsResult]:
    """
    This data source provides the Global Accelerator (GA) Basic Accelerators of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.194.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.ga.get_basic_accelerators(ids=["example_id"])
    pulumi.export("gaBasicAcceleratorId1", ids.accelerators[0].id)
    name_regex = alicloud.ga.get_basic_accelerators(name_regex="tf-example")
    pulumi.export("gaBasicAcceleratorId2", name_regex.accelerators[0].id)
    ```


    :param str accelerator_id: The ID of the Global Accelerator Basic Accelerator instance.
    :param Sequence[str] ids: A list of Global Accelerator Basic Accelerator IDs.
    :param str name_regex: A regex string to filter results by Global Accelerator Basic Accelerator name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the Global Accelerator Basic Accelerator instance. Valid Value: `init`, `active`, `configuring`, `binding`, `unbinding`, `deleting`, `finacialLocked`.
    """
    ...
