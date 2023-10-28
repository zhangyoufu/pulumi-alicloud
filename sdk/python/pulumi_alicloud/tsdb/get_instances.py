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
    'GetInstancesResult',
    'AwaitableGetInstancesResult',
    'get_instances',
    'get_instances_output',
]

@pulumi.output_type
class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, app_key=None, enable_details=None, engine_type=None, id=None, ids=None, instances=None, output_file=None, query_str=None, status=None, status_list=None):
        if app_key and not isinstance(app_key, str):
            raise TypeError("Expected argument 'app_key' to be a str")
        pulumi.set(__self__, "app_key", app_key)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if engine_type and not isinstance(engine_type, str):
            raise TypeError("Expected argument 'engine_type' to be a str")
        pulumi.set(__self__, "engine_type", engine_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if query_str and not isinstance(query_str, str):
            raise TypeError("Expected argument 'query_str' to be a str")
        pulumi.set(__self__, "query_str", query_str)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_list and not isinstance(status_list, str):
            raise TypeError("Expected argument 'status_list' to be a str")
        pulumi.set(__self__, "status_list", status_list)

    @property
    @pulumi.getter(name="appKey")
    def app_key(self) -> Optional[str]:
        return pulumi.get(self, "app_key")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> Optional[str]:
        return pulumi.get(self, "engine_type")

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
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetInstancesInstanceResult']:
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="queryStr")
    def query_str(self) -> Optional[str]:
        return pulumi.get(self, "query_str")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusList")
    def status_list(self) -> Optional[str]:
        return pulumi.get(self, "status_list")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            app_key=self.app_key,
            enable_details=self.enable_details,
            engine_type=self.engine_type,
            id=self.id,
            ids=self.ids,
            instances=self.instances,
            output_file=self.output_file,
            query_str=self.query_str,
            status=self.status,
            status_list=self.status_list)


def get_instances(app_key: Optional[str] = None,
                  enable_details: Optional[bool] = None,
                  engine_type: Optional[str] = None,
                  ids: Optional[Sequence[str]] = None,
                  output_file: Optional[str] = None,
                  query_str: Optional[str] = None,
                  status: Optional[str] = None,
                  status_list: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    This data source provides the Time Series Database (TSDB) Instances of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.112.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.tsdb.get_instances(ids=["example_value"])
    pulumi.export("firstTsdbInstanceId", example.instances[0].id)
    ```


    :param str app_key: The app key.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param str engine_type: The engine type of instance. Enumerative: `tsdb_tsdb` refers to TSDB, `tsdb_influxdb` refers to TSDB for InfluxDB️.
    :param Sequence[str] ids: A list of Instance IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str query_str: The query str.
    :param str status: Instance status, enumerative: ACTIVATION,DELETED, CREATING,CLASS_CHANGING,LOCKED.
    :param str status_list: The status list.
    """
    __args__ = dict()
    __args__['appKey'] = app_key
    __args__['enableDetails'] = enable_details
    __args__['engineType'] = engine_type
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['queryStr'] = query_str
    __args__['status'] = status
    __args__['statusList'] = status_list
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:tsdb/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        app_key=pulumi.get(__ret__, 'app_key'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        engine_type=pulumi.get(__ret__, 'engine_type'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instances=pulumi.get(__ret__, 'instances'),
        output_file=pulumi.get(__ret__, 'output_file'),
        query_str=pulumi.get(__ret__, 'query_str'),
        status=pulumi.get(__ret__, 'status'),
        status_list=pulumi.get(__ret__, 'status_list'))


@_utilities.lift_output_func(get_instances)
def get_instances_output(app_key: Optional[pulumi.Input[Optional[str]]] = None,
                         enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                         engine_type: Optional[pulumi.Input[Optional[str]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         query_str: Optional[pulumi.Input[Optional[str]]] = None,
                         status: Optional[pulumi.Input[Optional[str]]] = None,
                         status_list: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancesResult]:
    """
    This data source provides the Time Series Database (TSDB) Instances of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.112.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.tsdb.get_instances(ids=["example_value"])
    pulumi.export("firstTsdbInstanceId", example.instances[0].id)
    ```


    :param str app_key: The app key.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param str engine_type: The engine type of instance. Enumerative: `tsdb_tsdb` refers to TSDB, `tsdb_influxdb` refers to TSDB for InfluxDB️.
    :param Sequence[str] ids: A list of Instance IDs.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str query_str: The query str.
    :param str status: Instance status, enumerative: ACTIVATION,DELETED, CREATING,CLASS_CHANGING,LOCKED.
    :param str status_list: The status list.
    """
    ...
