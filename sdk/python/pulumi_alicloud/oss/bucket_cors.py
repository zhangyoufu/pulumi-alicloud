# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BucketCorsArgs', 'BucketCors']

@pulumi.input_type
class BucketCorsArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 cors_rules: pulumi.Input[Sequence[pulumi.Input['BucketCorsCorsRuleArgs']]],
                 response_vary: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a BucketCors resource.
        :param pulumi.Input[str] bucket: The name of the Bucket.
        :param pulumi.Input[Sequence[pulumi.Input['BucketCorsCorsRuleArgs']]] cors_rules: The Cross-Origin Resource Sharing (CORS) configuration of the Bucket. See `cors_rule` below.
        :param pulumi.Input[bool] response_vary: Specifies whether to return the Vary: Origin header. Valid values: true: returns the Vary: Origin header, regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds. false: does not return the Vary: Origin header. This element is valid only when at least one CORS rule is configured.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "cors_rules", cors_rules)
        if response_vary is not None:
            pulumi.set(__self__, "response_vary", response_vary)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        The name of the Bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="corsRules")
    def cors_rules(self) -> pulumi.Input[Sequence[pulumi.Input['BucketCorsCorsRuleArgs']]]:
        """
        The Cross-Origin Resource Sharing (CORS) configuration of the Bucket. See `cors_rule` below.
        """
        return pulumi.get(self, "cors_rules")

    @cors_rules.setter
    def cors_rules(self, value: pulumi.Input[Sequence[pulumi.Input['BucketCorsCorsRuleArgs']]]):
        pulumi.set(self, "cors_rules", value)

    @property
    @pulumi.getter(name="responseVary")
    def response_vary(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to return the Vary: Origin header. Valid values: true: returns the Vary: Origin header, regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds. false: does not return the Vary: Origin header. This element is valid only when at least one CORS rule is configured.
        """
        return pulumi.get(self, "response_vary")

    @response_vary.setter
    def response_vary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "response_vary", value)


@pulumi.input_type
class _BucketCorsState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 cors_rules: Optional[pulumi.Input[Sequence[pulumi.Input['BucketCorsCorsRuleArgs']]]] = None,
                 response_vary: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering BucketCors resources.
        :param pulumi.Input[str] bucket: The name of the Bucket.
        :param pulumi.Input[Sequence[pulumi.Input['BucketCorsCorsRuleArgs']]] cors_rules: The Cross-Origin Resource Sharing (CORS) configuration of the Bucket. See `cors_rule` below.
        :param pulumi.Input[bool] response_vary: Specifies whether to return the Vary: Origin header. Valid values: true: returns the Vary: Origin header, regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds. false: does not return the Vary: Origin header. This element is valid only when at least one CORS rule is configured.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if cors_rules is not None:
            pulumi.set(__self__, "cors_rules", cors_rules)
        if response_vary is not None:
            pulumi.set(__self__, "response_vary", response_vary)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="corsRules")
    def cors_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketCorsCorsRuleArgs']]]]:
        """
        The Cross-Origin Resource Sharing (CORS) configuration of the Bucket. See `cors_rule` below.
        """
        return pulumi.get(self, "cors_rules")

    @cors_rules.setter
    def cors_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketCorsCorsRuleArgs']]]]):
        pulumi.set(self, "cors_rules", value)

    @property
    @pulumi.getter(name="responseVary")
    def response_vary(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to return the Vary: Origin header. Valid values: true: returns the Vary: Origin header, regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds. false: does not return the Vary: Origin header. This element is valid only when at least one CORS rule is configured.
        """
        return pulumi.get(self, "response_vary")

    @response_vary.setter
    def response_vary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "response_vary", value)


class BucketCors(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 cors_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorsCorsRuleArgs']]]]] = None,
                 response_vary: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a OSS Bucket Cors resource. Cross-Origin Resource Sharing (CORS) allows web applications to access resources in other regions.

        For information about OSS Bucket Cors and how to use it, see [What is Bucket Cors](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.223.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        create_bucket = alicloud.oss.Bucket("CreateBucket",
            storage_class="Standard",
            bucket=name)
        default = alicloud.oss.BucketCors("default",
            bucket=create_bucket.bucket,
            response_vary=True,
            cors_rules=[alicloud.oss.BucketCorsCorsRuleArgs(
                allowed_methods=["GET"],
                allowed_origins=["*"],
                allowed_headers=[
                    "x-oss-test",
                    "x-oss-abc",
                ],
                expose_headers=["x-oss-request-id"],
                max_age_seconds=1000,
            )])
        ```

        ## Import

        OSS Bucket Cors can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:oss/bucketCors:BucketCors example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the Bucket.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorsCorsRuleArgs']]]] cors_rules: The Cross-Origin Resource Sharing (CORS) configuration of the Bucket. See `cors_rule` below.
        :param pulumi.Input[bool] response_vary: Specifies whether to return the Vary: Origin header. Valid values: true: returns the Vary: Origin header, regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds. false: does not return the Vary: Origin header. This element is valid only when at least one CORS rule is configured.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketCorsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a OSS Bucket Cors resource. Cross-Origin Resource Sharing (CORS) allows web applications to access resources in other regions.

        For information about OSS Bucket Cors and how to use it, see [What is Bucket Cors](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.223.0.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        create_bucket = alicloud.oss.Bucket("CreateBucket",
            storage_class="Standard",
            bucket=name)
        default = alicloud.oss.BucketCors("default",
            bucket=create_bucket.bucket,
            response_vary=True,
            cors_rules=[alicloud.oss.BucketCorsCorsRuleArgs(
                allowed_methods=["GET"],
                allowed_origins=["*"],
                allowed_headers=[
                    "x-oss-test",
                    "x-oss-abc",
                ],
                expose_headers=["x-oss-request-id"],
                max_age_seconds=1000,
            )])
        ```

        ## Import

        OSS Bucket Cors can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:oss/bucketCors:BucketCors example <id>
        ```

        :param str resource_name: The name of the resource.
        :param BucketCorsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketCorsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 cors_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorsCorsRuleArgs']]]]] = None,
                 response_vary: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketCorsArgs.__new__(BucketCorsArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if cors_rules is None and not opts.urn:
                raise TypeError("Missing required property 'cors_rules'")
            __props__.__dict__["cors_rules"] = cors_rules
            __props__.__dict__["response_vary"] = response_vary
        super(BucketCors, __self__).__init__(
            'alicloud:oss/bucketCors:BucketCors',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            cors_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorsCorsRuleArgs']]]]] = None,
            response_vary: Optional[pulumi.Input[bool]] = None) -> 'BucketCors':
        """
        Get an existing BucketCors resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the Bucket.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketCorsCorsRuleArgs']]]] cors_rules: The Cross-Origin Resource Sharing (CORS) configuration of the Bucket. See `cors_rule` below.
        :param pulumi.Input[bool] response_vary: Specifies whether to return the Vary: Origin header. Valid values: true: returns the Vary: Origin header, regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds. false: does not return the Vary: Origin header. This element is valid only when at least one CORS rule is configured.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketCorsState.__new__(_BucketCorsState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["cors_rules"] = cors_rules
        __props__.__dict__["response_vary"] = response_vary
        return BucketCors(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the Bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="corsRules")
    def cors_rules(self) -> pulumi.Output[Sequence['outputs.BucketCorsCorsRule']]:
        """
        The Cross-Origin Resource Sharing (CORS) configuration of the Bucket. See `cors_rule` below.
        """
        return pulumi.get(self, "cors_rules")

    @property
    @pulumi.getter(name="responseVary")
    def response_vary(self) -> pulumi.Output[bool]:
        """
        Specifies whether to return the Vary: Origin header. Valid values: true: returns the Vary: Origin header, regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds. false: does not return the Vary: Origin header. This element is valid only when at least one CORS rule is configured.
        """
        return pulumi.get(self, "response_vary")

