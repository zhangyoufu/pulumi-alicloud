# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 abandon_able_check_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 account_name_prefix: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 payer_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] display_name: Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] abandon_able_check_ids: The IDs of the check items that you can choose to ignore for the member deletion. 
               If you want to delete the account, please use datasource `resourcemanager_get_account_deletion_check_task`
               to get check ids and set them.
        :param pulumi.Input[str] account_name_prefix: The name prefix of account.
        :param pulumi.Input[str] folder_id: The ID of the parent folder.
        :param pulumi.Input[str] payer_account_id: The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
               
               > **NOTE:** The member name must be unique within the resource directory.
        """
        pulumi.set(__self__, "display_name", display_name)
        if abandon_able_check_ids is not None:
            pulumi.set(__self__, "abandon_able_check_ids", abandon_able_check_ids)
        if account_name_prefix is not None:
            pulumi.set(__self__, "account_name_prefix", account_name_prefix)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if payer_account_id is not None:
            pulumi.set(__self__, "payer_account_id", payer_account_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="abandonAbleCheckIds")
    def abandon_able_check_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of the check items that you can choose to ignore for the member deletion. 
        If you want to delete the account, please use datasource `resourcemanager_get_account_deletion_check_task`
        to get check ids and set them.
        """
        return pulumi.get(self, "abandon_able_check_ids")

    @abandon_able_check_ids.setter
    def abandon_able_check_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "abandon_able_check_ids", value)

    @property
    @pulumi.getter(name="accountNamePrefix")
    def account_name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The name prefix of account.
        """
        return pulumi.get(self, "account_name_prefix")

    @account_name_prefix.setter
    def account_name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name_prefix", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the parent folder.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter(name="payerAccountId")
    def payer_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
        """
        return pulumi.get(self, "payer_account_id")

    @payer_account_id.setter
    def payer_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payer_account_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.

        > **NOTE:** The member name must be unique within the resource directory.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _AccountState:
    def __init__(__self__, *,
                 abandon_able_check_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 account_name_prefix: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 join_method: Optional[pulumi.Input[str]] = None,
                 join_time: Optional[pulumi.Input[str]] = None,
                 modify_time: Optional[pulumi.Input[str]] = None,
                 payer_account_id: Optional[pulumi.Input[str]] = None,
                 resource_directory_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Account resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] abandon_able_check_ids: The IDs of the check items that you can choose to ignore for the member deletion. 
               If you want to delete the account, please use datasource `resourcemanager_get_account_deletion_check_task`
               to get check ids and set them.
        :param pulumi.Input[str] account_name_prefix: The name prefix of account.
        :param pulumi.Input[str] display_name: Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
        :param pulumi.Input[str] folder_id: The ID of the parent folder.
        :param pulumi.Input[str] join_method: Ways for members to join the resource directory. Valid values: `invited`, `created`.
        :param pulumi.Input[str] join_time: The time when the member joined the resource directory.
        :param pulumi.Input[str] modify_time: The modification time of the invitation.
        :param pulumi.Input[str] payer_account_id: The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
        :param pulumi.Input[str] resource_directory_id: Resource directory ID.
        :param pulumi.Input[str] status: Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
               
               > **NOTE:** The member name must be unique within the resource directory.
        :param pulumi.Input[str] type: Member type. The value of `ResourceAccount` indicates the resource account.
        """
        if abandon_able_check_ids is not None:
            pulumi.set(__self__, "abandon_able_check_ids", abandon_able_check_ids)
        if account_name_prefix is not None:
            pulumi.set(__self__, "account_name_prefix", account_name_prefix)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if join_method is not None:
            pulumi.set(__self__, "join_method", join_method)
        if join_time is not None:
            pulumi.set(__self__, "join_time", join_time)
        if modify_time is not None:
            pulumi.set(__self__, "modify_time", modify_time)
        if payer_account_id is not None:
            pulumi.set(__self__, "payer_account_id", payer_account_id)
        if resource_directory_id is not None:
            pulumi.set(__self__, "resource_directory_id", resource_directory_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="abandonAbleCheckIds")
    def abandon_able_check_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of the check items that you can choose to ignore for the member deletion. 
        If you want to delete the account, please use datasource `resourcemanager_get_account_deletion_check_task`
        to get check ids and set them.
        """
        return pulumi.get(self, "abandon_able_check_ids")

    @abandon_able_check_ids.setter
    def abandon_able_check_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "abandon_able_check_ids", value)

    @property
    @pulumi.getter(name="accountNamePrefix")
    def account_name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The name prefix of account.
        """
        return pulumi.get(self, "account_name_prefix")

    @account_name_prefix.setter
    def account_name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name_prefix", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the parent folder.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter(name="joinMethod")
    def join_method(self) -> Optional[pulumi.Input[str]]:
        """
        Ways for members to join the resource directory. Valid values: `invited`, `created`.
        """
        return pulumi.get(self, "join_method")

    @join_method.setter
    def join_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "join_method", value)

    @property
    @pulumi.getter(name="joinTime")
    def join_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the member joined the resource directory.
        """
        return pulumi.get(self, "join_time")

    @join_time.setter
    def join_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "join_time", value)

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> Optional[pulumi.Input[str]]:
        """
        The modification time of the invitation.
        """
        return pulumi.get(self, "modify_time")

    @modify_time.setter
    def modify_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modify_time", value)

    @property
    @pulumi.getter(name="payerAccountId")
    def payer_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
        """
        return pulumi.get(self, "payer_account_id")

    @payer_account_id.setter
    def payer_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payer_account_id", value)

    @property
    @pulumi.getter(name="resourceDirectoryId")
    def resource_directory_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource directory ID.
        """
        return pulumi.get(self, "resource_directory_id")

    @resource_directory_id.setter
    def resource_directory_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_directory_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.

        > **NOTE:** The member name must be unique within the resource directory.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Member type. The value of `ResourceAccount` indicates the resource account.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 abandon_able_check_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 account_name_prefix: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 payer_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a Resource Manager Account resource. Member accounts are containers for resources in a resource directory. These accounts isolate resources and serve as organizational units in the resource directory. You can create member accounts in a folder and then manage them in a unified manner.
        For information about Resource Manager Account and how to use it, see [What is Resource Manager Account](https://www.alibabacloud.com/help/en/doc-detail/111231.htm).

        > **NOTE:** Available since v1.83.0.

        > **NOTE:** From version 1.188.0, the resource can be destroyed. The member deletion feature is in invitational preview. You can contact the service manager of Alibaba Cloud to apply for a trial. see [how to destroy it](https://www.alibabacloud.com/help/en/resource-management/latest/delete-account).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        display_name = config.get("displayName")
        if display_name is None:
            display_name = "EAccount"
        default = random.RandomInteger("default",
            min=10000,
            max=99999)
        example_folder = alicloud.resourcemanager.Folder("exampleFolder", folder_name=default.result.apply(lambda result: f"{name}-{result}"))
        example_account = alicloud.resourcemanager.Account("exampleAccount",
            display_name=default.result.apply(lambda result: f"{display_name}-{result}"),
            folder_id=example_folder.id)
        ```
        <!--End PulumiCodeChooser -->

        ### Deleting `resourcemanager.Account` or removing it from your configuration

        Deleting the resource manager account or removing it from your configuration will remove it from your state file and management,
        but may not destroy the account. If there are some dependent resource in the account,
        the deleting account will enter a silence period of 45 days. After the silence period ends,
        the system automatically starts to delete the member. [See More Details](https://www.alibabacloud.com/help/en/resource-management/latest/delete-resource-account).

        ## Import

        Resource Manager Account can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:resourcemanager/account:Account example 13148890145*****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] abandon_able_check_ids: The IDs of the check items that you can choose to ignore for the member deletion. 
               If you want to delete the account, please use datasource `resourcemanager_get_account_deletion_check_task`
               to get check ids and set them.
        :param pulumi.Input[str] account_name_prefix: The name prefix of account.
        :param pulumi.Input[str] display_name: Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
        :param pulumi.Input[str] folder_id: The ID of the parent folder.
        :param pulumi.Input[str] payer_account_id: The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
               
               > **NOTE:** The member name must be unique within the resource directory.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Resource Manager Account resource. Member accounts are containers for resources in a resource directory. These accounts isolate resources and serve as organizational units in the resource directory. You can create member accounts in a folder and then manage them in a unified manner.
        For information about Resource Manager Account and how to use it, see [What is Resource Manager Account](https://www.alibabacloud.com/help/en/doc-detail/111231.htm).

        > **NOTE:** Available since v1.83.0.

        > **NOTE:** From version 1.188.0, the resource can be destroyed. The member deletion feature is in invitational preview. You can contact the service manager of Alibaba Cloud to apply for a trial. see [how to destroy it](https://www.alibabacloud.com/help/en/resource-management/latest/delete-account).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        display_name = config.get("displayName")
        if display_name is None:
            display_name = "EAccount"
        default = random.RandomInteger("default",
            min=10000,
            max=99999)
        example_folder = alicloud.resourcemanager.Folder("exampleFolder", folder_name=default.result.apply(lambda result: f"{name}-{result}"))
        example_account = alicloud.resourcemanager.Account("exampleAccount",
            display_name=default.result.apply(lambda result: f"{display_name}-{result}"),
            folder_id=example_folder.id)
        ```
        <!--End PulumiCodeChooser -->

        ### Deleting `resourcemanager.Account` or removing it from your configuration

        Deleting the resource manager account or removing it from your configuration will remove it from your state file and management,
        but may not destroy the account. If there are some dependent resource in the account,
        the deleting account will enter a silence period of 45 days. After the silence period ends,
        the system automatically starts to delete the member. [See More Details](https://www.alibabacloud.com/help/en/resource-management/latest/delete-resource-account).

        ## Import

        Resource Manager Account can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:resourcemanager/account:Account example 13148890145*****
        ```

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 abandon_able_check_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 account_name_prefix: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 payer_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            __props__.__dict__["abandon_able_check_ids"] = abandon_able_check_ids
            __props__.__dict__["account_name_prefix"] = account_name_prefix
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["payer_account_id"] = payer_account_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["join_method"] = None
            __props__.__dict__["join_time"] = None
            __props__.__dict__["modify_time"] = None
            __props__.__dict__["resource_directory_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
        super(Account, __self__).__init__(
            'alicloud:resourcemanager/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            abandon_able_check_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            account_name_prefix: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            join_method: Optional[pulumi.Input[str]] = None,
            join_time: Optional[pulumi.Input[str]] = None,
            modify_time: Optional[pulumi.Input[str]] = None,
            payer_account_id: Optional[pulumi.Input[str]] = None,
            resource_directory_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] abandon_able_check_ids: The IDs of the check items that you can choose to ignore for the member deletion. 
               If you want to delete the account, please use datasource `resourcemanager_get_account_deletion_check_task`
               to get check ids and set them.
        :param pulumi.Input[str] account_name_prefix: The name prefix of account.
        :param pulumi.Input[str] display_name: Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
        :param pulumi.Input[str] folder_id: The ID of the parent folder.
        :param pulumi.Input[str] join_method: Ways for members to join the resource directory. Valid values: `invited`, `created`.
        :param pulumi.Input[str] join_time: The time when the member joined the resource directory.
        :param pulumi.Input[str] modify_time: The modification time of the invitation.
        :param pulumi.Input[str] payer_account_id: The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
        :param pulumi.Input[str] resource_directory_id: Resource directory ID.
        :param pulumi.Input[str] status: Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
               
               > **NOTE:** The member name must be unique within the resource directory.
        :param pulumi.Input[str] type: Member type. The value of `ResourceAccount` indicates the resource account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountState.__new__(_AccountState)

        __props__.__dict__["abandon_able_check_ids"] = abandon_able_check_ids
        __props__.__dict__["account_name_prefix"] = account_name_prefix
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["join_method"] = join_method
        __props__.__dict__["join_time"] = join_time
        __props__.__dict__["modify_time"] = modify_time
        __props__.__dict__["payer_account_id"] = payer_account_id
        __props__.__dict__["resource_directory_id"] = resource_directory_id
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="abandonAbleCheckIds")
    def abandon_able_check_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IDs of the check items that you can choose to ignore for the member deletion. 
        If you want to delete the account, please use datasource `resourcemanager_get_account_deletion_check_task`
        to get check ids and set them.
        """
        return pulumi.get(self, "abandon_able_check_ids")

    @property
    @pulumi.getter(name="accountNamePrefix")
    def account_name_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The name prefix of account.
        """
        return pulumi.get(self, "account_name_prefix")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        """
        The ID of the parent folder.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter(name="joinMethod")
    def join_method(self) -> pulumi.Output[str]:
        """
        Ways for members to join the resource directory. Valid values: `invited`, `created`.
        """
        return pulumi.get(self, "join_method")

    @property
    @pulumi.getter(name="joinTime")
    def join_time(self) -> pulumi.Output[str]:
        """
        The time when the member joined the resource directory.
        """
        return pulumi.get(self, "join_time")

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> pulumi.Output[str]:
        """
        The modification time of the invitation.
        """
        return pulumi.get(self, "modify_time")

    @property
    @pulumi.getter(name="payerAccountId")
    def payer_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
        """
        return pulumi.get(self, "payer_account_id")

    @property
    @pulumi.getter(name="resourceDirectoryId")
    def resource_directory_id(self) -> pulumi.Output[str]:
        """
        Resource directory ID.
        """
        return pulumi.get(self, "resource_directory_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.

        > **NOTE:** The member name must be unique within the resource directory.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Member type. The value of `ResourceAccount` indicates the resource account.
        """
        return pulumi.get(self, "type")

