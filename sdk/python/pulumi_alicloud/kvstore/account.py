# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Account(pulumi.CustomResource):
    account_name: pulumi.Output[str]
    """
    Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
    """
    account_password: pulumi.Output[str]
    """
    Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
    """
    account_privilege: pulumi.Output[str]
    """
    The privilege of account access database. Valid values: 
    - RoleReadOnly: This value is only for Redis and Memcache
    - RoleReadWrite: This value is only for Redis and Memcache
    - RoleRepl: This value supports instance to read, write, and open SYNC / PSYNC commands.
    Only for Redis which engine version is 4.0 and architecture type is standard
    """
    account_type: pulumi.Output[str]
    """
    Privilege type of account.
    - Normal: Common privilege.
    Default to Normal.
    """
    description: pulumi.Output[str]
    """
    Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
    """
    instance_id: pulumi.Output[str]
    """
    The Id of instance in which account belongs. (The engine version of instance must be 4.0 or 4.0+)
    """
    kms_encrypted_password: pulumi.Output[str]
    """
    An KMS encrypts password used to a KVStore account. If the `account_password` is filled in, this field will be ignored.
    """
    kms_encryption_context: pulumi.Output[dict]
    """
    An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a KVStore account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, account_password=None, account_privilege=None, account_type=None, description=None, instance_id=None, kms_encrypted_password=None, kms_encryption_context=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a kvstore account resource and used to manage databases.

        > **NOTE:** Available in 1.66.0+



        > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/kvstore_account.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
        :param pulumi.Input[str] account_password: Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
        :param pulumi.Input[str] account_privilege: The privilege of account access database. Valid values: 
               - RoleReadOnly: This value is only for Redis and Memcache
               - RoleReadWrite: This value is only for Redis and Memcache
               - RoleRepl: This value supports instance to read, write, and open SYNC / PSYNC commands.
               Only for Redis which engine version is 4.0 and architecture type is standard
        :param pulumi.Input[str] account_type: Privilege type of account.
               - Normal: Common privilege.
               Default to Normal.
        :param pulumi.Input[str] description: Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        :param pulumi.Input[str] instance_id: The Id of instance in which account belongs. (The engine version of instance must be 4.0 or 4.0+)
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to a KVStore account. If the `account_password` is filled in, this field will be ignored.
        :param pulumi.Input[dict] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a KVStore account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
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
            __props__['account_password'] = account_password
            __props__['account_privilege'] = account_privilege
            __props__['account_type'] = account_type
            __props__['description'] = description
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['kms_encrypted_password'] = kms_encrypted_password
            __props__['kms_encryption_context'] = kms_encryption_context
        super(Account, __self__).__init__(
            'alicloud:kvstore/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_name=None, account_password=None, account_privilege=None, account_type=None, description=None, instance_id=None, kms_encrypted_password=None, kms_encryption_context=None):
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
        :param pulumi.Input[str] account_password: Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `account_password` and `kms_encrypted_password` fields.
        :param pulumi.Input[str] account_privilege: The privilege of account access database. Valid values: 
               - RoleReadOnly: This value is only for Redis and Memcache
               - RoleReadWrite: This value is only for Redis and Memcache
               - RoleRepl: This value supports instance to read, write, and open SYNC / PSYNC commands.
               Only for Redis which engine version is 4.0 and architecture type is standard
        :param pulumi.Input[str] account_type: Privilege type of account.
               - Normal: Common privilege.
               Default to Normal.
        :param pulumi.Input[str] description: Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
        :param pulumi.Input[str] instance_id: The Id of instance in which account belongs. (The engine version of instance must be 4.0 or 4.0+)
        :param pulumi.Input[str] kms_encrypted_password: An KMS encrypts password used to a KVStore account. If the `account_password` is filled in, this field will be ignored.
        :param pulumi.Input[dict] kms_encryption_context: An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a KVStore account with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["account_password"] = account_password
        __props__["account_privilege"] = account_privilege
        __props__["account_type"] = account_type
        __props__["description"] = description
        __props__["instance_id"] = instance_id
        __props__["kms_encrypted_password"] = kms_encrypted_password
        __props__["kms_encryption_context"] = kms_encryption_context
        return Account(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

