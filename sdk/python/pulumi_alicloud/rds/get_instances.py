# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInstancesResult:
    """
    A collection of values returned by getInstances.
    """
    def __init__(__self__, connection_mode=None, db_type=None, engine=None, id=None, ids=None, instances=None, name_regex=None, names=None, output_file=None, status=None, tags=None, vpc_id=None, vswitch_id=None):
        if connection_mode and not isinstance(connection_mode, str):
            raise TypeError("Expected argument 'connection_mode' to be a str")
        __self__.connection_mode = connection_mode
        """
        `Standard` for standard access mode and `Safe` for high security access mode.
        """
        if db_type and not isinstance(db_type, str):
            raise TypeError("Expected argument 'db_type' to be a str")
        __self__.db_type = db_type
        """
        `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
        """
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        __self__.engine = engine
        """
        Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of RDS instance IDs. 
        """
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        __self__.instances = instances
        """
        A list of RDS instances. Each element contains the following attributes:
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of RDS instance names. 
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        Status of the instance.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id
        """
        ID of the VPC the instance belongs to.
        """
        if vswitch_id and not isinstance(vswitch_id, str):
            raise TypeError("Expected argument 'vswitch_id' to be a str")
        __self__.vswitch_id = vswitch_id
        """
        ID of the VSwitch the instance belongs to.
        """
class AwaitableGetInstancesResult(GetInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesResult(
            connection_mode=self.connection_mode,
            db_type=self.db_type,
            engine=self.engine,
            id=self.id,
            ids=self.ids,
            instances=self.instances,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status,
            tags=self.tags,
            vpc_id=self.vpc_id,
            vswitch_id=self.vswitch_id)

def get_instances(connection_mode=None,db_type=None,engine=None,ids=None,name_regex=None,output_file=None,status=None,tags=None,vpc_id=None,vswitch_id=None,opts=None):
    """
    The `rds.getInstances` data source provides a collection of RDS instances available in Alibaba Cloud account.
    Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.




    :param str connection_mode: `Standard` for standard access mode and `Safe` for high security access mode.
    :param str db_type: `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
    :param str engine: Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
    :param list ids: A list of RDS instance IDs. 
    :param str name_regex: A regex string to filter results by instance name.
    :param str status: Status of the instance.
    :param dict tags: A map of tags assigned to the DB instances. 
           Note: Before 1.60.0, the value's format is a `json` string which including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `"{\"key1\":\"value1\"}"`
    :param str vpc_id: Used to retrieve instances belong to specified VPC.
    :param str vswitch_id: Used to retrieve instances belong to specified `vswitch` resources.
    """
    __args__ = dict()


    __args__['connectionMode'] = connection_mode
    __args__['dbType'] = db_type
    __args__['engine'] = engine
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    __args__['vswitchId'] = vswitch_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:rds/getInstances:getInstances', __args__, opts=opts).value

    return AwaitableGetInstancesResult(
        connection_mode=__ret__.get('connectionMode'),
        db_type=__ret__.get('dbType'),
        engine=__ret__.get('engine'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        instances=__ret__.get('instances'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        output_file=__ret__.get('outputFile'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'),
        vswitch_id=__ret__.get('vswitchId'))
