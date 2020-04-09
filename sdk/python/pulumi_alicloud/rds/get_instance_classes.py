# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInstanceClassesResult:
    """
    A collection of values returned by getInstanceClasses.
    """
    def __init__(__self__, category=None, db_instance_class=None, engine=None, engine_version=None, id=None, ids=None, instance_charge_type=None, instance_classes=None, multi_zone=None, output_file=None, sorted_by=None, storage_type=None, zone_id=None):
        if category and not isinstance(category, str):
            raise TypeError("Expected argument 'category' to be a str")
        __self__.category = category
        if db_instance_class and not isinstance(db_instance_class, str):
            raise TypeError("Expected argument 'db_instance_class' to be a str")
        __self__.db_instance_class = db_instance_class
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        __self__.engine = engine
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        __self__.engine_version = engine_version
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        (Available in 1.60.0+) A list of Rds instance class codes.
        """
        if instance_charge_type and not isinstance(instance_charge_type, str):
            raise TypeError("Expected argument 'instance_charge_type' to be a str")
        __self__.instance_charge_type = instance_charge_type
        if instance_classes and not isinstance(instance_classes, list):
            raise TypeError("Expected argument 'instance_classes' to be a list")
        __self__.instance_classes = instance_classes
        """
        A list of Rds available resource. Each element contains the following attributes:
        """
        if multi_zone and not isinstance(multi_zone, bool):
            raise TypeError("Expected argument 'multi_zone' to be a bool")
        __self__.multi_zone = multi_zone
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if sorted_by and not isinstance(sorted_by, str):
            raise TypeError("Expected argument 'sorted_by' to be a str")
        __self__.sorted_by = sorted_by
        if storage_type and not isinstance(storage_type, str):
            raise TypeError("Expected argument 'storage_type' to be a str")
        __self__.storage_type = storage_type
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        __self__.zone_id = zone_id
class AwaitableGetInstanceClassesResult(GetInstanceClassesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceClassesResult(
            category=self.category,
            db_instance_class=self.db_instance_class,
            engine=self.engine,
            engine_version=self.engine_version,
            id=self.id,
            ids=self.ids,
            instance_charge_type=self.instance_charge_type,
            instance_classes=self.instance_classes,
            multi_zone=self.multi_zone,
            output_file=self.output_file,
            sorted_by=self.sorted_by,
            storage_type=self.storage_type,
            zone_id=self.zone_id)

def get_instance_classes(category=None,db_instance_class=None,engine=None,engine_version=None,instance_charge_type=None,multi_zone=None,output_file=None,sorted_by=None,storage_type=None,zone_id=None,opts=None):
    """
    This data source provides the RDS instance classes resource available info of Alibaba Cloud.

    > **NOTE:** Available in v1.46.0+



    > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/db_instance_classes.html.markdown.


    :param str category: DB Instance category. the value like [`Basic`, `HighAvailability`, `Finance`], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).
    :param str db_instance_class: The DB instance class type by the user.
    :param str engine: Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
    :param str engine_version: Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
    :param str instance_charge_type: Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
    :param bool multi_zone: Whether to show multi available zone. Default false to not show multi availability zone.
    :param str storage_type: The DB instance storage space required by the user. Valid values: `cloud_ssd` and `local_ssd`.
    :param str zone_id: The Zone to launch the DB instance.
    """
    __args__ = dict()


    __args__['category'] = category
    __args__['dbInstanceClass'] = db_instance_class
    __args__['engine'] = engine
    __args__['engineVersion'] = engine_version
    __args__['instanceChargeType'] = instance_charge_type
    __args__['multiZone'] = multi_zone
    __args__['outputFile'] = output_file
    __args__['sortedBy'] = sorted_by
    __args__['storageType'] = storage_type
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:rds/getInstanceClasses:getInstanceClasses', __args__, opts=opts).value

    return AwaitableGetInstanceClassesResult(
        category=__ret__.get('category'),
        db_instance_class=__ret__.get('dbInstanceClass'),
        engine=__ret__.get('engine'),
        engine_version=__ret__.get('engineVersion'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        instance_charge_type=__ret__.get('instanceChargeType'),
        instance_classes=__ret__.get('instanceClasses'),
        multi_zone=__ret__.get('multiZone'),
        output_file=__ret__.get('outputFile'),
        sorted_by=__ret__.get('sortedBy'),
        storage_type=__ret__.get('storageType'),
        zone_id=__ret__.get('zoneId'))
