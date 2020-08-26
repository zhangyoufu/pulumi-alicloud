# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Role']


class Role(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assume_role_policy_document: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 max_session_duration: Optional[pulumi.Input[float]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Resource Manager role resource. Members are resource containers in the resource directory, which can physically isolate resources to form an independent resource grouping unit. You can create members in the resource folder to manage them in a unified manner.
        For information about Resource Manager role and how to use it, see [What is Resource Manager role](https://www.alibabacloud.com/help/en/doc-detail/111231.htm).

        > **NOTE:** Available in v1.82.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Add a Resource Manager role.
        example = alicloud.resourcemanager.Role("example",
            assume_role_policy_document=\"\"\"     {
                  "Statement": [
                       {
                            "Action": "sts:AssumeRole",
                            "Effect": "Allow",
                            "Principal": {
                                "RAM":[
                                        "acs:ram::103755469187****:root"，
                                        "acs:ram::104408977069****:root"
                                ]
                            }
                        }
                  ],
                  "Version": "1"
             }
        	 
        \"\"\",
            role_name="testrd")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] assume_role_policy_document: The content of the permissions strategy that plays a role.
        :param pulumi.Input[str] description: The description of the Resource Manager role.
        :param pulumi.Input[float] max_session_duration: Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
        :param pulumi.Input[str] role_name: Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-".
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if assume_role_policy_document is None:
                raise TypeError("Missing required property 'assume_role_policy_document'")
            __props__['assume_role_policy_document'] = assume_role_policy_document
            __props__['description'] = description
            __props__['max_session_duration'] = max_session_duration
            if role_name is None:
                raise TypeError("Missing required property 'role_name'")
            __props__['role_name'] = role_name
            __props__['arn'] = None
            __props__['create_date'] = None
            __props__['role_id'] = None
            __props__['update_date'] = None
        super(Role, __self__).__init__(
            'alicloud:resourcemanager/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            assume_role_policy_document: Optional[pulumi.Input[str]] = None,
            create_date: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            max_session_duration: Optional[pulumi.Input[float]] = None,
            role_id: Optional[pulumi.Input[str]] = None,
            role_name: Optional[pulumi.Input[str]] = None,
            update_date: Optional[pulumi.Input[str]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The resource descriptor of the role.
        :param pulumi.Input[str] assume_role_policy_document: The content of the permissions strategy that plays a role.
        :param pulumi.Input[str] create_date: Role creation time.
        :param pulumi.Input[str] description: The description of the Resource Manager role.
        :param pulumi.Input[float] max_session_duration: Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
        :param pulumi.Input[str] role_name: Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-".
        :param pulumi.Input[str] update_date: Role update time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["assume_role_policy_document"] = assume_role_policy_document
        __props__["create_date"] = create_date
        __props__["description"] = description
        __props__["max_session_duration"] = max_session_duration
        __props__["role_id"] = role_id
        __props__["role_name"] = role_name
        __props__["update_date"] = update_date
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The resource descriptor of the role.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="assumeRolePolicyDocument")
    def assume_role_policy_document(self) -> str:
        """
        The content of the permissions strategy that plays a role.
        """
        return pulumi.get(self, "assume_role_policy_document")

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> str:
        """
        Role creation time.
        """
        return pulumi.get(self, "create_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the Resource Manager role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> Optional[float]:
        """
        Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
        """
        return pulumi.get(self, "max_session_duration")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> str:
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> str:
        """
        Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-".
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter(name="updateDate")
    def update_date(self) -> str:
        """
        Role update time.
        """
        return pulumi.get(self, "update_date")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

