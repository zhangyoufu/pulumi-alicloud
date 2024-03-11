# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ImageExportArgs', 'ImageExport']

@pulumi.input_type
class ImageExportArgs:
    def __init__(__self__, *,
                 image_id: pulumi.Input[str],
                 oss_bucket: pulumi.Input[str],
                 oss_prefix: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ImageExport resource.
        :param pulumi.Input[str] image_id: The source image ID.
        :param pulumi.Input[str] oss_bucket: Save the exported OSS bucket.
        :param pulumi.Input[str] oss_prefix: The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        """
        pulumi.set(__self__, "image_id", image_id)
        pulumi.set(__self__, "oss_bucket", oss_bucket)
        if oss_prefix is not None:
            pulumi.set(__self__, "oss_prefix", oss_prefix)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Input[str]:
        """
        The source image ID.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="ossBucket")
    def oss_bucket(self) -> pulumi.Input[str]:
        """
        Save the exported OSS bucket.
        """
        return pulumi.get(self, "oss_bucket")

    @oss_bucket.setter
    def oss_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "oss_bucket", value)

    @property
    @pulumi.getter(name="ossPrefix")
    def oss_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        """
        return pulumi.get(self, "oss_prefix")

    @oss_prefix.setter
    def oss_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oss_prefix", value)


@pulumi.input_type
class _ImageExportState:
    def __init__(__self__, *,
                 image_id: Optional[pulumi.Input[str]] = None,
                 oss_bucket: Optional[pulumi.Input[str]] = None,
                 oss_prefix: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ImageExport resources.
        :param pulumi.Input[str] image_id: The source image ID.
        :param pulumi.Input[str] oss_bucket: Save the exported OSS bucket.
        :param pulumi.Input[str] oss_prefix: The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        """
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if oss_bucket is not None:
            pulumi.set(__self__, "oss_bucket", oss_bucket)
        if oss_prefix is not None:
            pulumi.set(__self__, "oss_prefix", oss_prefix)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The source image ID.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="ossBucket")
    def oss_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        Save the exported OSS bucket.
        """
        return pulumi.get(self, "oss_bucket")

    @oss_bucket.setter
    def oss_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oss_bucket", value)

    @property
    @pulumi.getter(name="ossPrefix")
    def oss_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        """
        return pulumi.get(self, "oss_prefix")

    @oss_prefix.setter
    def oss_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "oss_prefix", value)


class ImageExport(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 oss_bucket: Optional[pulumi.Input[str]] = None,
                 oss_prefix: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Export a custom image to the OSS bucket in the same region as the custom image.

        > **NOTE:** If you create an ECS instance using a mirror image and create a system disk snapshot again, exporting a custom image created from the system disk snapshot is not supported.

        > **NOTE:** Support for exporting custom images that include data disk snapshot information in the image. The number of data disks cannot exceed 4 and the maximum capacity of a single data disk cannot exceed 500 GiB.

        > **NOTE:** Before exporting the image, you must authorize the cloud server ECS official service account to write OSS permissions through RAM.

        > **NOTE:** Available since v1.68.0+.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        default_zones = alicloud.get_zones(available_resource_creation="Instance")
        default_instance_types = alicloud.ecs.get_instance_types(instance_type_family="ecs.sn1ne")
        default_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_instance = alicloud.ecs.Instance("defaultInstance",
            availability_zone=default_zones.zones[0].id,
            instance_name="terraform-example",
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id,
            instance_type=default_instance_types.ids[0],
            image_id=default_images.ids[0],
            internet_max_bandwidth_out=10)
        default_random_integer = random.RandomInteger("defaultRandomInteger",
            max=99999,
            min=10000)
        default_image = alicloud.ecs.Image("defaultImage",
            instance_id=default_instance.id,
            image_name=default_random_integer.result.apply(lambda result: f"terraform-example-{result}"),
            description="terraform-example")
        default_bucket = alicloud.oss.Bucket("defaultBucket", bucket=default_random_integer.result.apply(lambda result: f"example-value-{result}"))
        default_image_export = alicloud.ecs.ImageExport("defaultImageExport",
            image_id=default_image.id,
            oss_bucket=default_bucket.id,
            oss_prefix="ecsExport")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] image_id: The source image ID.
        :param pulumi.Input[str] oss_bucket: Save the exported OSS bucket.
        :param pulumi.Input[str] oss_prefix: The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageExportArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Export a custom image to the OSS bucket in the same region as the custom image.

        > **NOTE:** If you create an ECS instance using a mirror image and create a system disk snapshot again, exporting a custom image created from the system disk snapshot is not supported.

        > **NOTE:** Support for exporting custom images that include data disk snapshot information in the image. The number of data disks cannot exceed 4 and the maximum capacity of a single data disk cannot exceed 500 GiB.

        > **NOTE:** Before exporting the image, you must authorize the cloud server ECS official service account to write OSS permissions through RAM.

        > **NOTE:** Available since v1.68.0+.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        default_zones = alicloud.get_zones(available_resource_creation="Instance")
        default_instance_types = alicloud.ecs.get_instance_types(instance_type_family="ecs.sn1ne")
        default_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_instance = alicloud.ecs.Instance("defaultInstance",
            availability_zone=default_zones.zones[0].id,
            instance_name="terraform-example",
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id,
            instance_type=default_instance_types.ids[0],
            image_id=default_images.ids[0],
            internet_max_bandwidth_out=10)
        default_random_integer = random.RandomInteger("defaultRandomInteger",
            max=99999,
            min=10000)
        default_image = alicloud.ecs.Image("defaultImage",
            instance_id=default_instance.id,
            image_name=default_random_integer.result.apply(lambda result: f"terraform-example-{result}"),
            description="terraform-example")
        default_bucket = alicloud.oss.Bucket("defaultBucket", bucket=default_random_integer.result.apply(lambda result: f"example-value-{result}"))
        default_image_export = alicloud.ecs.ImageExport("defaultImageExport",
            image_id=default_image.id,
            oss_bucket=default_bucket.id,
            oss_prefix="ecsExport")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param ImageExportArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageExportArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 oss_bucket: Optional[pulumi.Input[str]] = None,
                 oss_prefix: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageExportArgs.__new__(ImageExportArgs)

            if image_id is None and not opts.urn:
                raise TypeError("Missing required property 'image_id'")
            __props__.__dict__["image_id"] = image_id
            if oss_bucket is None and not opts.urn:
                raise TypeError("Missing required property 'oss_bucket'")
            __props__.__dict__["oss_bucket"] = oss_bucket
            __props__.__dict__["oss_prefix"] = oss_prefix
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

        __props__ = _ImageExportState.__new__(_ImageExportState)

        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["oss_bucket"] = oss_bucket
        __props__.__dict__["oss_prefix"] = oss_prefix
        return ImageExport(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        """
        The source image ID.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="ossBucket")
    def oss_bucket(self) -> pulumi.Output[str]:
        """
        Save the exported OSS bucket.
        """
        return pulumi.get(self, "oss_bucket")

    @property
    @pulumi.getter(name="ossPrefix")
    def oss_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
        """
        return pulumi.get(self, "oss_prefix")

