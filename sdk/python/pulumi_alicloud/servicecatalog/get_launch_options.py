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

__all__ = [
    'GetLaunchOptionsResult',
    'AwaitableGetLaunchOptionsResult',
    'get_launch_options',
    'get_launch_options_output',
]

@pulumi.output_type
class GetLaunchOptionsResult:
    """
    A collection of values returned by getLaunchOptions.
    """
    def __init__(__self__, id=None, ids=None, launch_options=None, name_regex=None, options=None, output_file=None, product_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if launch_options and not isinstance(launch_options, list):
            raise TypeError("Expected argument 'launch_options' to be a list")
        pulumi.set(__self__, "launch_options", launch_options)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if options and not isinstance(options, list):
            raise TypeError("Expected argument 'options' to be a list")
        if options is not None:
            warnings.warn("""Field 'options' has been deprecated from provider version 1.197.0.""", DeprecationWarning)
            pulumi.log.warn("""options is deprecated: Field 'options' has been deprecated from provider version 1.197.0.""")

        pulumi.set(__self__, "options", options)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if product_id and not isinstance(product_id, str):
            raise TypeError("Expected argument 'product_id' to be a str")
        pulumi.set(__self__, "product_id", product_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="launchOptions")
    def launch_options(self) -> Sequence['outputs.GetLaunchOptionsLaunchOptionResult']:
        """
        A list of Launch Option Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "launch_options")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def options(self) -> Sequence['outputs.GetLaunchOptionsOptionResult']:
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> str:
        return pulumi.get(self, "product_id")


class AwaitableGetLaunchOptionsResult(GetLaunchOptionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLaunchOptionsResult(
            id=self.id,
            ids=self.ids,
            launch_options=self.launch_options,
            name_regex=self.name_regex,
            options=self.options,
            output_file=self.output_file,
            product_id=self.product_id)


def get_launch_options(ids: Optional[Sequence[str]] = None,
                       name_regex: Optional[str] = None,
                       output_file: Optional[str] = None,
                       product_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLaunchOptionsResult:
    """
    This data source provides Service Catalog Launch Option available to the user.[What is Launch Option](https://www.alibabacloud.com/help/en/servicecatalog/latest/api-doc-servicecatalog-2021-09-01-api-doc-listlaunchoptions)

    > **NOTE:** Available in 1.196.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_end_user_products = alicloud.servicecatalog.get_end_user_products(name_regex="ram模板创建")
    default_launch_options = alicloud.servicecatalog.get_launch_options(product_id="data.alicloud_service_catalog_end_user_products.default.end_user_products.0.id")
    pulumi.export("alicloudServiceCatalogLaunchOptionExampleId", default_launch_options.launch_options[0].id)
    ```


    :param str name_regex: A regex string to filter results by portfolio name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str product_id: Product ID.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['productId'] = product_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:servicecatalog/getLaunchOptions:getLaunchOptions', __args__, opts=opts, typ=GetLaunchOptionsResult).value

    return AwaitableGetLaunchOptionsResult(
        id=__ret__.id,
        ids=__ret__.ids,
        launch_options=__ret__.launch_options,
        name_regex=__ret__.name_regex,
        options=__ret__.options,
        output_file=__ret__.output_file,
        product_id=__ret__.product_id)


@_utilities.lift_output_func(get_launch_options)
def get_launch_options_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              product_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLaunchOptionsResult]:
    """
    This data source provides Service Catalog Launch Option available to the user.[What is Launch Option](https://www.alibabacloud.com/help/en/servicecatalog/latest/api-doc-servicecatalog-2021-09-01-api-doc-listlaunchoptions)

    > **NOTE:** Available in 1.196.0+

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_end_user_products = alicloud.servicecatalog.get_end_user_products(name_regex="ram模板创建")
    default_launch_options = alicloud.servicecatalog.get_launch_options(product_id="data.alicloud_service_catalog_end_user_products.default.end_user_products.0.id")
    pulumi.export("alicloudServiceCatalogLaunchOptionExampleId", default_launch_options.launch_options[0].id)
    ```


    :param str name_regex: A regex string to filter results by portfolio name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str product_id: Product ID.
    """
    ...
