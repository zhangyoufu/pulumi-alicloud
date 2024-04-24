# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BucketVersioningInitArgs', 'BucketVersioning']

@pulumi.input_type
class BucketVersioningInitArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BucketVersioning resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] status: A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
        """
        pulumi.set(__self__, "bucket", bucket)
        if status is not None:
            pulumi.set(__self__, "status", status)

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

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _BucketVersioningState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BucketVersioning resources.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] status: A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
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
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class BucketVersioning(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        OSS Bucket Versioning can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:oss/bucketVersioning:BucketVersioning example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] status: A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketVersioningInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        OSS Bucket Versioning can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:oss/bucketVersioning:BucketVersioning example <id>
        ```

        :param str resource_name: The name of the resource.
        :param BucketVersioningInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketVersioningInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketVersioningInitArgs.__new__(BucketVersioningInitArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            __props__.__dict__["status"] = status
        super(BucketVersioning, __self__).__init__(
            'alicloud:oss/bucketVersioning:BucketVersioning',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'BucketVersioning':
        """
        Get an existing BucketVersioning resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] status: A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketVersioningState.__new__(_BucketVersioningState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["status"] = status
        return BucketVersioning(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        A bucket can be in one of the following versioning states: disabled, enabled, or suspended. By default, versioning is disabled for a bucket. Updating the value from Enabled or Suspended to Disabled will result in errors, because OSS does not support returning buckets to an unversioned state. .
        """
        return pulumi.get(self, "status")

