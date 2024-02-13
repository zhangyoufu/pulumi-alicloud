# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FolderArgs', 'Folder']

@pulumi.input_type
class FolderArgs:
    def __init__(__self__, *,
                 folder_name: pulumi.Input[str],
                 parent_folder_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Folder resource.
        :param pulumi.Input[str] folder_name: The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        :param pulumi.Input[str] parent_folder_id: The ID of the parent folder. If not set, the system default value will be used.
        """
        pulumi.set(__self__, "folder_name", folder_name)
        if parent_folder_id is not None:
            pulumi.set(__self__, "parent_folder_id", parent_folder_id)

    @property
    @pulumi.getter(name="folderName")
    def folder_name(self) -> pulumi.Input[str]:
        """
        The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        """
        return pulumi.get(self, "folder_name")

    @folder_name.setter
    def folder_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "folder_name", value)

    @property
    @pulumi.getter(name="parentFolderId")
    def parent_folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the parent folder. If not set, the system default value will be used.
        """
        return pulumi.get(self, "parent_folder_id")

    @parent_folder_id.setter
    def parent_folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_folder_id", value)


@pulumi.input_type
class _FolderState:
    def __init__(__self__, *,
                 folder_name: Optional[pulumi.Input[str]] = None,
                 parent_folder_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Folder resources.
        :param pulumi.Input[str] folder_name: The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        :param pulumi.Input[str] parent_folder_id: The ID of the parent folder. If not set, the system default value will be used.
        """
        if folder_name is not None:
            pulumi.set(__self__, "folder_name", folder_name)
        if parent_folder_id is not None:
            pulumi.set(__self__, "parent_folder_id", parent_folder_id)

    @property
    @pulumi.getter(name="folderName")
    def folder_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        """
        return pulumi.get(self, "folder_name")

    @folder_name.setter
    def folder_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_name", value)

    @property
    @pulumi.getter(name="parentFolderId")
    def parent_folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the parent folder. If not set, the system default value will be used.
        """
        return pulumi.get(self, "parent_folder_id")

    @parent_folder_id.setter
    def parent_folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_folder_id", value)


class Folder(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder_name: Optional[pulumi.Input[str]] = None,
                 parent_folder_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Resource Manager Folder resource. A folder is an organizational unit in a resource directory. You can use folders to build an organizational structure for resources.
        For information about Resource Manager Foler and how to use it, see [What is Resource Manager Folder](https://www.alibabacloud.com/help/en/doc-detail/111221.htm).

        > **NOTE:** Available since v1.82.0.

        > **NOTE:** A maximum of five levels of folders can be created under the root folder.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        example = alicloud.resourcemanager.Folder("example", folder_name=name)
        ```

        ## Import

        Resource Manager Folder can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:resourcemanager/folder:Folder example fd-u8B321****	
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder_name: The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        :param pulumi.Input[str] parent_folder_id: The ID of the parent folder. If not set, the system default value will be used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FolderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Resource Manager Folder resource. A folder is an organizational unit in a resource directory. You can use folders to build an organizational structure for resources.
        For information about Resource Manager Foler and how to use it, see [What is Resource Manager Folder](https://www.alibabacloud.com/help/en/doc-detail/111221.htm).

        > **NOTE:** Available since v1.82.0.

        > **NOTE:** A maximum of five levels of folders can be created under the root folder.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        example = alicloud.resourcemanager.Folder("example", folder_name=name)
        ```

        ## Import

        Resource Manager Folder can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:resourcemanager/folder:Folder example fd-u8B321****	
        ```

        :param str resource_name: The name of the resource.
        :param FolderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FolderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder_name: Optional[pulumi.Input[str]] = None,
                 parent_folder_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FolderArgs.__new__(FolderArgs)

            if folder_name is None and not opts.urn:
                raise TypeError("Missing required property 'folder_name'")
            __props__.__dict__["folder_name"] = folder_name
            __props__.__dict__["parent_folder_id"] = parent_folder_id
        super(Folder, __self__).__init__(
            'alicloud:resourcemanager/folder:Folder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            folder_name: Optional[pulumi.Input[str]] = None,
            parent_folder_id: Optional[pulumi.Input[str]] = None) -> 'Folder':
        """
        Get an existing Folder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder_name: The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        :param pulumi.Input[str] parent_folder_id: The ID of the parent folder. If not set, the system default value will be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FolderState.__new__(_FolderState)

        __props__.__dict__["folder_name"] = folder_name
        __props__.__dict__["parent_folder_id"] = parent_folder_id
        return Folder(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="folderName")
    def folder_name(self) -> pulumi.Output[str]:
        """
        The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
        """
        return pulumi.get(self, "folder_name")

    @property
    @pulumi.getter(name="parentFolderId")
    def parent_folder_id(self) -> pulumi.Output[str]:
        """
        The ID of the parent folder. If not set, the system default value will be used.
        """
        return pulumi.get(self, "parent_folder_id")

