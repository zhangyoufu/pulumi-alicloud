# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class BackupPolicy(pulumi.CustomResource):
    backup_periods: pulumi.Output[list]
    """
    Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
    """
    backup_time: pulumi.Output[str]
    """
    Backup time, in the format of HH:mmZ- HH:mm Z
    """
    instance_id: pulumi.Output[str]
    """
    The id of ApsaraDB for Redis or Memcache intance.
    """
    def __init__(__self__, resource_name, opts=None, backup_periods=None, backup_time=None, instance_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a backup policy for ApsaraDB Redis / Memcache instance resource. 



        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/kvstore_backup_policy.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] backup_periods: Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
        :param pulumi.Input[str] backup_time: Backup time, in the format of HH:mmZ- HH:mm Z
        :param pulumi.Input[str] instance_id: The id of ApsaraDB for Redis or Memcache intance.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['backup_periods'] = backup_periods
            __props__['backup_time'] = backup_time
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
        super(BackupPolicy, __self__).__init__(
            'alicloud:kvstore/backupPolicy:BackupPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backup_periods=None, backup_time=None, instance_id=None):
        """
        Get an existing BackupPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] backup_periods: Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
        :param pulumi.Input[str] backup_time: Backup time, in the format of HH:mmZ- HH:mm Z
        :param pulumi.Input[str] instance_id: The id of ApsaraDB for Redis or Memcache intance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backup_periods"] = backup_periods
        __props__["backup_time"] = backup_time
        __props__["instance_id"] = instance_id
        return BackupPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

