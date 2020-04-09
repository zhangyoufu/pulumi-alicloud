# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AccountPrivilege(pulumi.CustomResource):
    account_name: pulumi.Output[str]
    """
    A specified account name.
    """
    db_names: pulumi.Output[list]
    """
    List of specified database name.
    """
    instance_id: pulumi.Output[str]
    """
    The Id of instance in which account belongs.
    """
    privilege: pulumi.Output[str]
    """
    The privilege of one account access database. Valid values: 
    - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
    - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
    - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
    - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
    - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, db_names=None, instance_id=None, privilege=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an RDS account privilege resource and used to grant several database some access privilege. A database can be granted by multiple account.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/db_account_privilege.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: A specified account name.
        :param pulumi.Input[list] db_names: List of specified database name.
        :param pulumi.Input[str] instance_id: The Id of instance in which account belongs.
        :param pulumi.Input[str] privilege: The privilege of one account access database. Valid values: 
               - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
               - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
               - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
               - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
               - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if db_names is None:
                raise TypeError("Missing required property 'db_names'")
            __props__['db_names'] = db_names
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['privilege'] = privilege
        super(AccountPrivilege, __self__).__init__(
            'alicloud:rds/accountPrivilege:AccountPrivilege',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_name=None, db_names=None, instance_id=None, privilege=None):
        """
        Get an existing AccountPrivilege resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: A specified account name.
        :param pulumi.Input[list] db_names: List of specified database name.
        :param pulumi.Input[str] instance_id: The Id of instance in which account belongs.
        :param pulumi.Input[str] privilege: The privilege of one account access database. Valid values: 
               - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
               - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
               - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
               - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
               - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["db_names"] = db_names
        __props__["instance_id"] = instance_id
        __props__["privilege"] = privilege
        return AccountPrivilege(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

