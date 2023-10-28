# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetMetaTagsTagResult',
]

@pulumi.output_type
class GetMetaTagsTagResult(dict):
    def __init__(__self__, *,
                 category: str,
                 key_name: str,
                 value_name: str):
        """
        :param str category: The type of the resource tags.
        :param str key_name: The name of the key.
        :param str value_name: The name of the value.
        """
        pulumi.set(__self__, "category", category)
        pulumi.set(__self__, "key_name", key_name)
        pulumi.set(__self__, "value_name", value_name)

    @property
    @pulumi.getter
    def category(self) -> str:
        """
        The type of the resource tags.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> str:
        """
        The name of the key.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="valueName")
    def value_name(self) -> str:
        """
        The name of the value.
        """
        return pulumi.get(self, "value_name")


