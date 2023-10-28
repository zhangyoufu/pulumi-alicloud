# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WebLockConfigArgs', 'WebLockConfig']

@pulumi.input_type
class WebLockConfigArgs:
    def __init__(__self__, *,
                 defence_mode: pulumi.Input[str],
                 dir: pulumi.Input[str],
                 local_backup_dir: pulumi.Input[str],
                 mode: pulumi.Input[str],
                 uuid: pulumi.Input[str],
                 exclusive_dir: Optional[pulumi.Input[str]] = None,
                 exclusive_file: Optional[pulumi.Input[str]] = None,
                 exclusive_file_type: Optional[pulumi.Input[str]] = None,
                 inclusive_file_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebLockConfig resource.
        :param pulumi.Input[str] defence_mode: Protection mode. Value:-**block**: Intercept-**audit**: Alarm
        :param pulumi.Input[str] dir: Specify the protection directory.
        :param pulumi.Input[str] local_backup_dir: The local backup path is used to protect the safe backup of the Directory.
        :param pulumi.Input[str] mode: Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
        :param pulumi.Input[str] uuid: Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
        :param pulumi.Input[str] exclusive_dir: Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] exclusive_file: Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] exclusive_file_type: Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] inclusive_file_type: Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
        """
        pulumi.set(__self__, "defence_mode", defence_mode)
        pulumi.set(__self__, "dir", dir)
        pulumi.set(__self__, "local_backup_dir", local_backup_dir)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "uuid", uuid)
        if exclusive_dir is not None:
            pulumi.set(__self__, "exclusive_dir", exclusive_dir)
        if exclusive_file is not None:
            pulumi.set(__self__, "exclusive_file", exclusive_file)
        if exclusive_file_type is not None:
            pulumi.set(__self__, "exclusive_file_type", exclusive_file_type)
        if inclusive_file_type is not None:
            pulumi.set(__self__, "inclusive_file_type", inclusive_file_type)

    @property
    @pulumi.getter(name="defenceMode")
    def defence_mode(self) -> pulumi.Input[str]:
        """
        Protection mode. Value:-**block**: Intercept-**audit**: Alarm
        """
        return pulumi.get(self, "defence_mode")

    @defence_mode.setter
    def defence_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "defence_mode", value)

    @property
    @pulumi.getter
    def dir(self) -> pulumi.Input[str]:
        """
        Specify the protection directory.
        """
        return pulumi.get(self, "dir")

    @dir.setter
    def dir(self, value: pulumi.Input[str]):
        pulumi.set(self, "dir", value)

    @property
    @pulumi.getter(name="localBackupDir")
    def local_backup_dir(self) -> pulumi.Input[str]:
        """
        The local backup path is used to protect the safe backup of the Directory.
        """
        return pulumi.get(self, "local_backup_dir")

    @local_backup_dir.setter
    def local_backup_dir(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_backup_dir", value)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Input[str]:
        """
        Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: pulumi.Input[str]):
        pulumi.set(self, "uuid", value)

    @property
    @pulumi.getter(name="exclusiveDir")
    def exclusive_dir(self) -> Optional[pulumi.Input[str]]:
        """
        Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        """
        return pulumi.get(self, "exclusive_dir")

    @exclusive_dir.setter
    def exclusive_dir(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclusive_dir", value)

    @property
    @pulumi.getter(name="exclusiveFile")
    def exclusive_file(self) -> Optional[pulumi.Input[str]]:
        """
        Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        """
        return pulumi.get(self, "exclusive_file")

    @exclusive_file.setter
    def exclusive_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclusive_file", value)

    @property
    @pulumi.getter(name="exclusiveFileType")
    def exclusive_file_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        """
        return pulumi.get(self, "exclusive_file_type")

    @exclusive_file_type.setter
    def exclusive_file_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclusive_file_type", value)

    @property
    @pulumi.getter(name="inclusiveFileType")
    def inclusive_file_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
        """
        return pulumi.get(self, "inclusive_file_type")

    @inclusive_file_type.setter
    def inclusive_file_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inclusive_file_type", value)


@pulumi.input_type
class _WebLockConfigState:
    def __init__(__self__, *,
                 defence_mode: Optional[pulumi.Input[str]] = None,
                 dir: Optional[pulumi.Input[str]] = None,
                 exclusive_dir: Optional[pulumi.Input[str]] = None,
                 exclusive_file: Optional[pulumi.Input[str]] = None,
                 exclusive_file_type: Optional[pulumi.Input[str]] = None,
                 inclusive_file_type: Optional[pulumi.Input[str]] = None,
                 local_backup_dir: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebLockConfig resources.
        :param pulumi.Input[str] defence_mode: Protection mode. Value:-**block**: Intercept-**audit**: Alarm
        :param pulumi.Input[str] dir: Specify the protection directory.
        :param pulumi.Input[str] exclusive_dir: Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] exclusive_file: Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] exclusive_file_type: Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] inclusive_file_type: Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
        :param pulumi.Input[str] local_backup_dir: The local backup path is used to protect the safe backup of the Directory.
        :param pulumi.Input[str] mode: Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
        :param pulumi.Input[str] uuid: Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
        """
        if defence_mode is not None:
            pulumi.set(__self__, "defence_mode", defence_mode)
        if dir is not None:
            pulumi.set(__self__, "dir", dir)
        if exclusive_dir is not None:
            pulumi.set(__self__, "exclusive_dir", exclusive_dir)
        if exclusive_file is not None:
            pulumi.set(__self__, "exclusive_file", exclusive_file)
        if exclusive_file_type is not None:
            pulumi.set(__self__, "exclusive_file_type", exclusive_file_type)
        if inclusive_file_type is not None:
            pulumi.set(__self__, "inclusive_file_type", inclusive_file_type)
        if local_backup_dir is not None:
            pulumi.set(__self__, "local_backup_dir", local_backup_dir)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="defenceMode")
    def defence_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Protection mode. Value:-**block**: Intercept-**audit**: Alarm
        """
        return pulumi.get(self, "defence_mode")

    @defence_mode.setter
    def defence_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "defence_mode", value)

    @property
    @pulumi.getter
    def dir(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the protection directory.
        """
        return pulumi.get(self, "dir")

    @dir.setter
    def dir(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dir", value)

    @property
    @pulumi.getter(name="exclusiveDir")
    def exclusive_dir(self) -> Optional[pulumi.Input[str]]:
        """
        Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        """
        return pulumi.get(self, "exclusive_dir")

    @exclusive_dir.setter
    def exclusive_dir(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclusive_dir", value)

    @property
    @pulumi.getter(name="exclusiveFile")
    def exclusive_file(self) -> Optional[pulumi.Input[str]]:
        """
        Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        """
        return pulumi.get(self, "exclusive_file")

    @exclusive_file.setter
    def exclusive_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclusive_file", value)

    @property
    @pulumi.getter(name="exclusiveFileType")
    def exclusive_file_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        """
        return pulumi.get(self, "exclusive_file_type")

    @exclusive_file_type.setter
    def exclusive_file_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclusive_file_type", value)

    @property
    @pulumi.getter(name="inclusiveFileType")
    def inclusive_file_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
        """
        return pulumi.get(self, "inclusive_file_type")

    @inclusive_file_type.setter
    def inclusive_file_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inclusive_file_type", value)

    @property
    @pulumi.getter(name="localBackupDir")
    def local_backup_dir(self) -> Optional[pulumi.Input[str]]:
        """
        The local backup path is used to protect the safe backup of the Directory.
        """
        return pulumi.get(self, "local_backup_dir")

    @local_backup_dir.setter
    def local_backup_dir(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_backup_dir", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


class WebLockConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 defence_mode: Optional[pulumi.Input[str]] = None,
                 dir: Optional[pulumi.Input[str]] = None,
                 exclusive_dir: Optional[pulumi.Input[str]] = None,
                 exclusive_file: Optional[pulumi.Input[str]] = None,
                 exclusive_file_type: Optional[pulumi.Input[str]] = None,
                 inclusive_file_type: Optional[pulumi.Input[str]] = None,
                 local_backup_dir: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Threat Detection Web Lock Config resource.

        For information about Threat Detection Web Lock Config and how to use it, see [What is Web Lock Config](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifyweblockstart).

        > **NOTE:** Available in v1.195.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_assets = alicloud.threatdetection.get_assets(machine_types="ecs")
        default_web_lock_config = alicloud.threatdetection.WebLockConfig("defaultWebLockConfig",
            inclusive_file_type="php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg",
            uuid=default_assets.ids[0],
            mode="whitelist",
            local_backup_dir="/usr/local/aegis/bak",
            dir="/tmp/",
            defence_mode="audit")
        ```

        ## Import

        Threat Detection Web Lock Config can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:threatdetection/webLockConfig:WebLockConfig example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defence_mode: Protection mode. Value:-**block**: Intercept-**audit**: Alarm
        :param pulumi.Input[str] dir: Specify the protection directory.
        :param pulumi.Input[str] exclusive_dir: Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] exclusive_file: Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] exclusive_file_type: Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] inclusive_file_type: Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
        :param pulumi.Input[str] local_backup_dir: The local backup path is used to protect the safe backup of the Directory.
        :param pulumi.Input[str] mode: Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
        :param pulumi.Input[str] uuid: Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebLockConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Threat Detection Web Lock Config resource.

        For information about Threat Detection Web Lock Config and how to use it, see [What is Web Lock Config](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifyweblockstart).

        > **NOTE:** Available in v1.195.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_assets = alicloud.threatdetection.get_assets(machine_types="ecs")
        default_web_lock_config = alicloud.threatdetection.WebLockConfig("defaultWebLockConfig",
            inclusive_file_type="php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg",
            uuid=default_assets.ids[0],
            mode="whitelist",
            local_backup_dir="/usr/local/aegis/bak",
            dir="/tmp/",
            defence_mode="audit")
        ```

        ## Import

        Threat Detection Web Lock Config can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:threatdetection/webLockConfig:WebLockConfig example <id>
        ```

        :param str resource_name: The name of the resource.
        :param WebLockConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebLockConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 defence_mode: Optional[pulumi.Input[str]] = None,
                 dir: Optional[pulumi.Input[str]] = None,
                 exclusive_dir: Optional[pulumi.Input[str]] = None,
                 exclusive_file: Optional[pulumi.Input[str]] = None,
                 exclusive_file_type: Optional[pulumi.Input[str]] = None,
                 inclusive_file_type: Optional[pulumi.Input[str]] = None,
                 local_backup_dir: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebLockConfigArgs.__new__(WebLockConfigArgs)

            if defence_mode is None and not opts.urn:
                raise TypeError("Missing required property 'defence_mode'")
            __props__.__dict__["defence_mode"] = defence_mode
            if dir is None and not opts.urn:
                raise TypeError("Missing required property 'dir'")
            __props__.__dict__["dir"] = dir
            __props__.__dict__["exclusive_dir"] = exclusive_dir
            __props__.__dict__["exclusive_file"] = exclusive_file
            __props__.__dict__["exclusive_file_type"] = exclusive_file_type
            __props__.__dict__["inclusive_file_type"] = inclusive_file_type
            if local_backup_dir is None and not opts.urn:
                raise TypeError("Missing required property 'local_backup_dir'")
            __props__.__dict__["local_backup_dir"] = local_backup_dir
            if mode is None and not opts.urn:
                raise TypeError("Missing required property 'mode'")
            __props__.__dict__["mode"] = mode
            if uuid is None and not opts.urn:
                raise TypeError("Missing required property 'uuid'")
            __props__.__dict__["uuid"] = uuid
        super(WebLockConfig, __self__).__init__(
            'alicloud:threatdetection/webLockConfig:WebLockConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            defence_mode: Optional[pulumi.Input[str]] = None,
            dir: Optional[pulumi.Input[str]] = None,
            exclusive_dir: Optional[pulumi.Input[str]] = None,
            exclusive_file: Optional[pulumi.Input[str]] = None,
            exclusive_file_type: Optional[pulumi.Input[str]] = None,
            inclusive_file_type: Optional[pulumi.Input[str]] = None,
            local_backup_dir: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None) -> 'WebLockConfig':
        """
        Get an existing WebLockConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defence_mode: Protection mode. Value:-**block**: Intercept-**audit**: Alarm
        :param pulumi.Input[str] dir: Specify the protection directory.
        :param pulumi.Input[str] exclusive_dir: Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] exclusive_file: Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] exclusive_file_type: Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        :param pulumi.Input[str] inclusive_file_type: Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
        :param pulumi.Input[str] local_backup_dir: The local backup path is used to protect the safe backup of the Directory.
        :param pulumi.Input[str] mode: Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
        :param pulumi.Input[str] uuid: Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebLockConfigState.__new__(_WebLockConfigState)

        __props__.__dict__["defence_mode"] = defence_mode
        __props__.__dict__["dir"] = dir
        __props__.__dict__["exclusive_dir"] = exclusive_dir
        __props__.__dict__["exclusive_file"] = exclusive_file
        __props__.__dict__["exclusive_file_type"] = exclusive_file_type
        __props__.__dict__["inclusive_file_type"] = inclusive_file_type
        __props__.__dict__["local_backup_dir"] = local_backup_dir
        __props__.__dict__["mode"] = mode
        __props__.__dict__["uuid"] = uuid
        return WebLockConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defenceMode")
    def defence_mode(self) -> pulumi.Output[str]:
        """
        Protection mode. Value:-**block**: Intercept-**audit**: Alarm
        """
        return pulumi.get(self, "defence_mode")

    @property
    @pulumi.getter
    def dir(self) -> pulumi.Output[str]:
        """
        Specify the protection directory.
        """
        return pulumi.get(self, "dir")

    @property
    @pulumi.getter(name="exclusiveDir")
    def exclusive_dir(self) -> pulumi.Output[Optional[str]]:
        """
        Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        """
        return pulumi.get(self, "exclusive_dir")

    @property
    @pulumi.getter(name="exclusiveFile")
    def exclusive_file(self) -> pulumi.Output[Optional[str]]:
        """
        Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        """
        return pulumi.get(self, "exclusive_file")

    @property
    @pulumi.getter(name="exclusiveFileType")
    def exclusive_file_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
        """
        return pulumi.get(self, "exclusive_file_type")

    @property
    @pulumi.getter(name="inclusiveFileType")
    def inclusive_file_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
        """
        return pulumi.get(self, "inclusive_file_type")

    @property
    @pulumi.getter(name="localBackupDir")
    def local_backup_dir(self) -> pulumi.Output[str]:
        """
        The local backup path is used to protect the safe backup of the Directory.
        """
        return pulumi.get(self, "local_backup_dir")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
        """
        return pulumi.get(self, "uuid")

