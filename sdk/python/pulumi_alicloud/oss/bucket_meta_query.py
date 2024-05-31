# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BucketMetaQueryArgs', 'BucketMetaQuery']

@pulumi.input_type
class BucketMetaQueryArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str]):
        """
        The set of arguments for constructing a BucketMetaQuery resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        """
        pulumi.set(__self__, "bucket", bucket)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)


@pulumi.input_type
class _BucketMetaQueryState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BucketMetaQuery resources.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] create_time: The creation time of the metadata index database. The format is mm:ss + TIMEZONE in the YYYY-MM-DDTHH format of RFC 3339. Where YYYY-MM-DD indicates the year, month and day, T indicates the beginning of the time element, HH:mm:ss indicates the hour, minute and second, and TIMEZONE indicates the time zone.
        :param pulumi.Input[str] status: The status of the resource.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the metadata index database. The format is mm:ss + TIMEZONE in the YYYY-MM-DDTHH format of RFC 3339. Where YYYY-MM-DD indicates the year, month and day, T indicates the beginning of the time element, HH:mm:ss indicates the hour, minute and second, and TIMEZONE indicates the time zone.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class BucketMetaQuery(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a OSS Bucket Meta Query resource. Enables the metadata management feature for a bucket.

        For information about OSS Bucket Meta Query and how to use it, see [What is Bucket Meta Query](https://www.alibabacloud.com/help/en/oss/developer-reference/openmetaquery).

        > **NOTE:** Available since v1.224.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = random.index.Integer("default",
            min=10000,
            max=99999)
        create_bucket = alicloud.oss.Bucket("CreateBucket",
            storage_class="Standard",
            bucket=f"{name}-{default['result']}")
        default_bucket_meta_query = alicloud.oss.BucketMetaQuery("default", bucket=create_bucket.bucket)
        ```

        ## Import

        OSS Bucket Meta Query can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:oss/bucketMetaQuery:BucketMetaQuery example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketMetaQueryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a OSS Bucket Meta Query resource. Enables the metadata management feature for a bucket.

        For information about OSS Bucket Meta Query and how to use it, see [What is Bucket Meta Query](https://www.alibabacloud.com/help/en/oss/developer-reference/openmetaquery).

        > **NOTE:** Available since v1.224.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = random.index.Integer("default",
            min=10000,
            max=99999)
        create_bucket = alicloud.oss.Bucket("CreateBucket",
            storage_class="Standard",
            bucket=f"{name}-{default['result']}")
        default_bucket_meta_query = alicloud.oss.BucketMetaQuery("default", bucket=create_bucket.bucket)
        ```

        ## Import

        OSS Bucket Meta Query can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:oss/bucketMetaQuery:BucketMetaQuery example <id>
        ```

        :param str resource_name: The name of the resource.
        :param BucketMetaQueryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketMetaQueryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketMetaQueryArgs.__new__(BucketMetaQueryArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            __props__.__dict__["create_time"] = None
            __props__.__dict__["status"] = None
        super(BucketMetaQuery, __self__).__init__(
            'alicloud:oss/bucketMetaQuery:BucketMetaQuery',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'BucketMetaQuery':
        """
        Get an existing BucketMetaQuery resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] create_time: The creation time of the metadata index database. The format is mm:ss + TIMEZONE in the YYYY-MM-DDTHH format of RFC 3339. Where YYYY-MM-DD indicates the year, month and day, T indicates the beginning of the time element, HH:mm:ss indicates the hour, minute and second, and TIMEZONE indicates the time zone.
        :param pulumi.Input[str] status: The status of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketMetaQueryState.__new__(_BucketMetaQueryState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["status"] = status
        return BucketMetaQuery(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of the metadata index database. The format is mm:ss + TIMEZONE in the YYYY-MM-DDTHH format of RFC 3339. Where YYYY-MM-DD indicates the year, month and day, T indicates the beginning of the time element, HH:mm:ss indicates the hour, minute and second, and TIMEZONE indicates the time zone.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource.
        """
        return pulumi.get(self, "status")

