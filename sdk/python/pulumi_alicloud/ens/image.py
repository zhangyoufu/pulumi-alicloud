# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ImageArgs', 'Image']

@pulumi.input_type
class ImageArgs:
    def __init__(__self__, *,
                 image_name: pulumi.Input[str],
                 delete_after_image_upload: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Image resource.
        :param pulumi.Input[str] image_name: Image Name.
        :param pulumi.Input[str] delete_after_image_upload: Whether the instance is automatically released after the image is packaged and uploaded successfully, only the build machine is supported.  Optional values: true: When the instance is released, the image is released together with the instance. false: When the instance is released, the image is retained and is not released together with the instance. Empty means false by default.
        :param pulumi.Input[str] instance_id: The ID of the instance corresponding to the image.
        """
        pulumi.set(__self__, "image_name", image_name)
        if delete_after_image_upload is not None:
            pulumi.set(__self__, "delete_after_image_upload", delete_after_image_upload)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> pulumi.Input[str]:
        """
        Image Name.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_name", value)

    @property
    @pulumi.getter(name="deleteAfterImageUpload")
    def delete_after_image_upload(self) -> Optional[pulumi.Input[str]]:
        """
        Whether the instance is automatically released after the image is packaged and uploaded successfully, only the build machine is supported.  Optional values: true: When the instance is released, the image is released together with the instance. false: When the instance is released, the image is retained and is not released together with the instance. Empty means false by default.
        """
        return pulumi.get(self, "delete_after_image_upload")

    @delete_after_image_upload.setter
    def delete_after_image_upload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_after_image_upload", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance corresponding to the image.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _ImageState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 delete_after_image_upload: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Image resources.
        :param pulumi.Input[str] create_time: Image creation time.
        :param pulumi.Input[str] delete_after_image_upload: Whether the instance is automatically released after the image is packaged and uploaded successfully, only the build machine is supported.  Optional values: true: When the instance is released, the image is released together with the instance. false: When the instance is released, the image is retained and is not released together with the instance. Empty means false by default.
        :param pulumi.Input[str] image_name: Image Name.
        :param pulumi.Input[str] instance_id: The ID of the instance corresponding to the image.
        :param pulumi.Input[str] status: Mirror Status  Optional values: Creating: Creating Packing: Packing Uploading: Uploading Pack_failed: Packing failed Upload_failed: Upload failed Available: Only images in the Available state can be used and operated. Unavailable: Not available Copying: Copying.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if delete_after_image_upload is not None:
            pulumi.set(__self__, "delete_after_image_upload", delete_after_image_upload)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Image creation time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="deleteAfterImageUpload")
    def delete_after_image_upload(self) -> Optional[pulumi.Input[str]]:
        """
        Whether the instance is automatically released after the image is packaged and uploaded successfully, only the build machine is supported.  Optional values: true: When the instance is released, the image is released together with the instance. false: When the instance is released, the image is retained and is not released together with the instance. Empty means false by default.
        """
        return pulumi.get(self, "delete_after_image_upload")

    @delete_after_image_upload.setter
    def delete_after_image_upload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_after_image_upload", value)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[pulumi.Input[str]]:
        """
        Image Name.
        """
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance corresponding to the image.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Mirror Status  Optional values: Creating: Creating Packing: Packing Uploading: Uploading Pack_failed: Packing failed Upload_failed: Upload failed Available: Only images in the Available state can be used and operated. Unavailable: Not available Copying: Copying.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Image(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_after_image_upload: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a ENS Image resource.

        For information about ENS Image and how to use it, see [What is Image](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.216.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_instance = alicloud.ens.Instance("defaultInstance",
            system_disk=alicloud.ens.InstanceSystemDiskArgs(
                size=20,
            ),
            schedule_area_level="Region",
            image_id="centos_6_08_64_20G_alibase_20171208",
            payment_type="PayAsYouGo",
            password="12345678ABCabc",
            amount=1,
            internet_max_bandwidth_out=10,
            public_ip_identification=True,
            ens_region_id="cn-chenzhou-telecom_unicom_cmcc",
            period_unit="Month",
            instance_type="ens.sn1.stiny",
            status="Stopped")
        default_image = alicloud.ens.Image("defaultImage",
            image_name=name,
            instance_id=default_instance.id,
            delete_after_image_upload="false")
        ```

        ## Import

        ENS Image can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ens/image:Image example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] delete_after_image_upload: Whether the instance is automatically released after the image is packaged and uploaded successfully, only the build machine is supported.  Optional values: true: When the instance is released, the image is released together with the instance. false: When the instance is released, the image is retained and is not released together with the instance. Empty means false by default.
        :param pulumi.Input[str] image_name: Image Name.
        :param pulumi.Input[str] instance_id: The ID of the instance corresponding to the image.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a ENS Image resource.

        For information about ENS Image and how to use it, see [What is Image](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.216.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default_instance = alicloud.ens.Instance("defaultInstance",
            system_disk=alicloud.ens.InstanceSystemDiskArgs(
                size=20,
            ),
            schedule_area_level="Region",
            image_id="centos_6_08_64_20G_alibase_20171208",
            payment_type="PayAsYouGo",
            password="12345678ABCabc",
            amount=1,
            internet_max_bandwidth_out=10,
            public_ip_identification=True,
            ens_region_id="cn-chenzhou-telecom_unicom_cmcc",
            period_unit="Month",
            instance_type="ens.sn1.stiny",
            status="Stopped")
        default_image = alicloud.ens.Image("defaultImage",
            image_name=name,
            instance_id=default_instance.id,
            delete_after_image_upload="false")
        ```

        ## Import

        ENS Image can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ens/image:Image example <id>
        ```

        :param str resource_name: The name of the resource.
        :param ImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_after_image_upload: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageArgs.__new__(ImageArgs)

            __props__.__dict__["delete_after_image_upload"] = delete_after_image_upload
            if image_name is None and not opts.urn:
                raise TypeError("Missing required property 'image_name'")
            __props__.__dict__["image_name"] = image_name
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["status"] = None
        super(Image, __self__).__init__(
            'alicloud:ens/image:Image',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            delete_after_image_upload: Optional[pulumi.Input[str]] = None,
            image_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Image':
        """
        Get an existing Image resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Image creation time.
        :param pulumi.Input[str] delete_after_image_upload: Whether the instance is automatically released after the image is packaged and uploaded successfully, only the build machine is supported.  Optional values: true: When the instance is released, the image is released together with the instance. false: When the instance is released, the image is retained and is not released together with the instance. Empty means false by default.
        :param pulumi.Input[str] image_name: Image Name.
        :param pulumi.Input[str] instance_id: The ID of the instance corresponding to the image.
        :param pulumi.Input[str] status: Mirror Status  Optional values: Creating: Creating Packing: Packing Uploading: Uploading Pack_failed: Packing failed Upload_failed: Upload failed Available: Only images in the Available state can be used and operated. Unavailable: Not available Copying: Copying.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImageState.__new__(_ImageState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["delete_after_image_upload"] = delete_after_image_upload
        __props__.__dict__["image_name"] = image_name
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["status"] = status
        return Image(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Image creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteAfterImageUpload")
    def delete_after_image_upload(self) -> pulumi.Output[Optional[str]]:
        """
        Whether the instance is automatically released after the image is packaged and uploaded successfully, only the build machine is supported.  Optional values: true: When the instance is released, the image is released together with the instance. false: When the instance is released, the image is retained and is not released together with the instance. Empty means false by default.
        """
        return pulumi.get(self, "delete_after_image_upload")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> pulumi.Output[str]:
        """
        Image Name.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the instance corresponding to the image.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Mirror Status  Optional values: Creating: Creating Packing: Packing Uploading: Uploading Pack_failed: Packing failed Upload_failed: Upload failed Available: Only images in the Available state can be used and operated. Unavailable: Not available Copying: Copying.
        """
        return pulumi.get(self, "status")

