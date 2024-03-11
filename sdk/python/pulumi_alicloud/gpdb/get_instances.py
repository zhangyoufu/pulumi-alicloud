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
    def __init__(__self__, availability_zone=None, db_instance_categories=None, db_instance_modes=None, description=None, enable_details=None, id=None, ids=None, instance_network_type=None, instances=None, name_regex=None, names=None, output_file=None, resource_group_id=None, status=None, tags=None, vswitch_id=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if db_instance_categories and not isinstance(db_instance_categories, str):
            raise TypeError("Expected argument 'db_instance_categories' to be a str")
        pulumi.set(__self__, "db_instance_categories", db_instance_categories)
        if db_instance_modes and not isinstance(db_instance_modes, str):
            raise TypeError("Expected argument 'db_instance_modes' to be a str")
        pulumi.set(__self__, "db_instance_modes", db_instance_modes)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_network_type and not isinstance(instance_network_type, str):
            raise TypeError("Expected argument 'instance_network_type' to be a str")
        pulumi.set(__self__, "instance_network_type", instance_network_type)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vswitch_id and not isinstance(vswitch_id, str):
            raise TypeError("Expected argument 'vswitch_id' to be a str")
        pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="dbInstanceCategories")
    def db_instance_categories(self) -> Optional[str]:
        return pulumi.get(self, "db_instance_categories")

    @property
    @pulumi.getter(name="dbInstanceModes")
    def db_instance_modes(self) -> Optional[str]:
        return pulumi.get(self, "db_instance_modes")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

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
    @pulumi.getter(name="instanceNetworkType")
    def instance_network_type(self) -> Optional[str]:
        return pulumi.get(self, "instance_network_type")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetInstancesInstanceResult']:
        return pulumi.get(self, "instances")

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
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[str]:
        return pulumi.get(self, "vswitch_id")


class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            availability_zone=self.availability_zone,
            db_instance_categories=self.db_instance_categories,
            db_instance_modes=self.db_instance_modes,
            description=self.description,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            instance_network_type=self.instance_network_type,
            instances=self.instances,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            resource_group_id=self.resource_group_id,
            status=self.status,
            tags=self.tags,
            vswitch_id=self.vswitch_id)


def get_instances(availability_zone: Optional[str] = None,
                  db_instance_categories: Optional[str] = None,
                  db_instance_modes: Optional[str] = None,
                  description: Optional[str] = None,
                  enable_details: Optional[bool] = None,
                  ids: Optional[Sequence[str]] = None,
                  instance_network_type: Optional[str] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  resource_group_id: Optional[str] = None,
                  status: Optional[str] = None,
                  tags: Optional[Mapping[str, Any]] = None,
                  vswitch_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstancesResult:
    """
    This data source provides the AnalyticDB for PostgreSQL instances of the current Alibaba Cloud user.

    > **NOTE:**  Available in 1.47.0+

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.gpdb.get_instances()
    pulumi.export("gpdbDbInstanceId1", ids.instances[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str availability_zone: Instance availability zone.
    :param str db_instance_categories: The db instance categories.
    :param str db_instance_modes: The db instance modes.
    :param str description: The description of the instance.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: The ids list of AnalyticDB for PostgreSQL instances.
    :param str instance_network_type: The network type of the instance.
    :param str name_regex: A regex string to apply to the instance name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The ID of the enterprise resource group to which the instance belongs.
    :param str status: The status of the instance. Valid values: `Creating`, `DBInstanceClassChanging`, `DBInstanceNetTypeChanging`, `Deleting`, `EngineVersionUpgrading`, `GuardDBInstanceCreating`, `GuardSwitching`, `Importing`, `ImportingFromOtherInstance`, `Rebooting`, `Restoring`, `Running`, `Transfering`, `TransferingToOtherInstance`.
    :param Mapping[str, Any] tags: The tags of the instance.
    :param str vswitch_id: The vswitch id.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['dbInstanceCategories'] = db_instance_categories
    __args__['dbInstanceModes'] = db_instance_modes
    __args__['description'] = description
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['instanceNetworkType'] = instance_network_type
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['resourceGroupId'] = resource_group_id
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['vswitchId'] = vswitch_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:gpdb/getInstances:getInstances', __args__, opts=opts, typ=GetInstancesResult).value

    return AwaitableGetInstancesResult(
        availability_zone=pulumi.get(__ret__, 'availability_zone'),
        db_instance_categories=pulumi.get(__ret__, 'db_instance_categories'),
        db_instance_modes=pulumi.get(__ret__, 'db_instance_modes'),
        description=pulumi.get(__ret__, 'description'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instance_network_type=pulumi.get(__ret__, 'instance_network_type'),
        instances=pulumi.get(__ret__, 'instances'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        resource_group_id=pulumi.get(__ret__, 'resource_group_id'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        vswitch_id=pulumi.get(__ret__, 'vswitch_id'))


@_utilities.lift_output_func(get_instances)
def get_instances_output(availability_zone: Optional[pulumi.Input[Optional[str]]] = None,
                         db_instance_categories: Optional[pulumi.Input[Optional[str]]] = None,
                         db_instance_modes: Optional[pulumi.Input[Optional[str]]] = None,
                         description: Optional[pulumi.Input[Optional[str]]] = None,
                         enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         instance_network_type: Optional[pulumi.Input[Optional[str]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                         status: Optional[pulumi.Input[Optional[str]]] = None,
                         tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                         vswitch_id: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstancesResult]:
    """
    This data source provides the AnalyticDB for PostgreSQL instances of the current Alibaba Cloud user.

    > **NOTE:**  Available in 1.47.0+

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.gpdb.get_instances()
    pulumi.export("gpdbDbInstanceId1", ids.instances[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str availability_zone: Instance availability zone.
    :param str db_instance_categories: The db instance categories.
    :param str db_instance_modes: The db instance modes.
    :param str description: The description of the instance.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: The ids list of AnalyticDB for PostgreSQL instances.
    :param str instance_network_type: The network type of the instance.
    :param str name_regex: A regex string to apply to the instance name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The ID of the enterprise resource group to which the instance belongs.
    :param str status: The status of the instance. Valid values: `Creating`, `DBInstanceClassChanging`, `DBInstanceNetTypeChanging`, `Deleting`, `EngineVersionUpgrading`, `GuardDBInstanceCreating`, `GuardSwitching`, `Importing`, `ImportingFromOtherInstance`, `Rebooting`, `Restoring`, `Running`, `Transfering`, `TransferingToOtherInstance`.
    :param Mapping[str, Any] tags: The tags of the instance.
    :param str vswitch_id: The vswitch id.
    """
    ...
