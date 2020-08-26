# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ImageExport']


class ImageExport(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 oss_bucket: Optional[pulumi.Input[str]] = None,
                 oss_prefix: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Export a custom image to the OSS bucket in the same region as the custom image.

        > **NOTE:** If you create an ECS instance using a mirror image and create a system disk snapshot again, exporting a custom image created from the system disk snapshot is not supported.

        > **NOTE:** Support for exporting custom images that include data disk snapshot information in the image. The number of data disks cannot exceed 4 and the maximum capacity of a single data disk cannot exceed 500 GiB.

        > **NOTE:** Before exporting the image, you must authorize the cloud server ECS official service account to write OSS permissions through RAM.

        > **NOTE:** Available in 1.68.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.ecs.ImageExport("default",
            image_id="m-bp1gxy***",
            oss_bucket="ecsimageexportconfig",
            oss_prefix="ecsExport")
        ```
        ## Attributes Reference0

         The following attributes are exported:

        * `id` - ID of the image.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] image_id: The source image ID.
        :param pulumi.Input[str] oss_bucket: Save the exported OSS bucket.
        :param pulumi.Input[str] oss_prefix: The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
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

            if image_id is None:
                raise TypeError("Missing required property 'image_id'")
            __props__['image_id'] = image_id
            if oss_bucket is None:
                raise TypeError("Missing required property 'oss_bucket'")
            __props__['oss_bucket'] = oss_bucket
            __props__['oss_prefix'] = oss_prefix
        super(ImageExport, __self__).__init__(
            'alicloud:ecs/imageExport:ImageExport',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            oss_bucket: Optional[pulumi.Input[str]] = None,
            oss_prefix: Optional[pulumi.Input[str]] = None) -> 'ImageExport':
        """
        Get an existing ImageExport resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] image_id: The source image ID.
        :param pulumi.Input[str] oss_bucket: Save the exported OSS bucket.
        :param pulumi.Input[str] oss_prefix: The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["image_id"] = image_id
        __props__["oss_bucket"] = oss_bucket
        __props__["oss_prefix"] = oss_prefix
        return ImageExport(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        """
        The source image ID.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="ossBucket")
    def oss_bucket(self) -> str:
        """
        Save the exported OSS bucket.
        """
        return pulumi.get(self, "oss_bucket")

    @property
    @pulumi.getter(name="ossPrefix")
    def oss_prefix(self) -> Optional[str]:
        """
        The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        """
        return pulumi.get(self, "oss_prefix")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

