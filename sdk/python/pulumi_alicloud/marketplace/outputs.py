# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetProductProductResult',
    'GetProductProductSkusResult',
    'GetProductProductSkusImageResult',
    'GetProductProductSkusPackageVersionResult',
    'GetProductsProductResult',
]

@pulumi.output_type
class GetProductProductResult(dict):
    def __init__(__self__, *,
                 code: str,
                 description: str,
                 name: str,
                 skuses: List['outputs.GetProductProductSkusResult']):
        """
        :param str code: The code of the product.
        :param str description: The description of the product.
        :param str name: The name of the product.
        :param List['GetProductProductSkusArgs'] skuses: A list of one element containing sku attributes of an object. Each element contains the following attributes:
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "skuses", skuses)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        The code of the product.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the product.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the product.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def skuses(self) -> List['outputs.GetProductProductSkusResult']:
        """
        A list of one element containing sku attributes of an object. Each element contains the following attributes:
        """
        return pulumi.get(self, "skuses")


@pulumi.output_type
class GetProductProductSkusResult(dict):
    def __init__(__self__, *,
                 images: List['outputs.GetProductProductSkusImageResult'],
                 package_versions: List['outputs.GetProductProductSkusPackageVersionResult'],
                 sku_code: str,
                 sku_name: str):
        """
        :param List['GetProductProductSkusImageArgs'] images: The list of custom ECS images, Each element contains the following attributes:
        :param List['GetProductProductSkusPackageVersionArgs'] package_versions: The list of package version details of this product sku, Each element contains the following attributes:
        :param str sku_code: The sku code of this product sku.
        :param str sku_name: The sku name of this product sku.
        """
        pulumi.set(__self__, "images", images)
        pulumi.set(__self__, "package_versions", package_versions)
        pulumi.set(__self__, "sku_code", sku_code)
        pulumi.set(__self__, "sku_name", sku_name)

    @property
    @pulumi.getter
    def images(self) -> List['outputs.GetProductProductSkusImageResult']:
        """
        The list of custom ECS images, Each element contains the following attributes:
        """
        return pulumi.get(self, "images")

    @property
    @pulumi.getter(name="packageVersions")
    def package_versions(self) -> List['outputs.GetProductProductSkusPackageVersionResult']:
        """
        The list of package version details of this product sku, Each element contains the following attributes:
        """
        return pulumi.get(self, "package_versions")

    @property
    @pulumi.getter(name="skuCode")
    def sku_code(self) -> str:
        """
        The sku code of this product sku.
        """
        return pulumi.get(self, "sku_code")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> str:
        """
        The sku name of this product sku.
        """
        return pulumi.get(self, "sku_name")


@pulumi.output_type
class GetProductProductSkusImageResult(dict):
    def __init__(__self__, *,
                 image_id: str,
                 image_name: str,
                 region_id: str):
        """
        :param str image_id: The Ecs image id.
        :param str image_name: The Ecs image display name.
        :param str region_id: The Ecs image region.
        """
        pulumi.set(__self__, "image_id", image_id)
        pulumi.set(__self__, "image_name", image_name)
        pulumi.set(__self__, "region_id", region_id)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        """
        The Ecs image id.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> str:
        """
        The Ecs image display name.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> str:
        """
        The Ecs image region.
        """
        return pulumi.get(self, "region_id")


@pulumi.output_type
class GetProductProductSkusPackageVersionResult(dict):
    def __init__(__self__, *,
                 package_name: str,
                 package_version: str):
        """
        :param str package_name: The package name of this product sku package.
        :param str package_version: The package version of this product sku package. Currently, the API products can return package_version, but others can not for ensure.
        """
        pulumi.set(__self__, "package_name", package_name)
        pulumi.set(__self__, "package_version", package_version)

    @property
    @pulumi.getter(name="packageName")
    def package_name(self) -> str:
        """
        The package name of this product sku package.
        """
        return pulumi.get(self, "package_name")

    @property
    @pulumi.getter(name="packageVersion")
    def package_version(self) -> str:
        """
        The package version of this product sku package. Currently, the API products can return package_version, but others can not for ensure.
        """
        return pulumi.get(self, "package_version")


@pulumi.output_type
class GetProductsProductResult(dict):
    def __init__(__self__, *,
                 category_id: float,
                 code: str,
                 delivery_date: str,
                 delivery_way: str,
                 image_url: str,
                 name: str,
                 operation_system: str,
                 score: str,
                 short_description: str,
                 suggested_price: str,
                 supplier_id: float,
                 supplier_name: str,
                 tags: str,
                 target_url: str,
                 warranty_date: str):
        """
        :param float category_id: The Category ID of products. For more information, see [DescribeProducts](https://help.aliyun.com/document_detail/89834.htm).
        :param str code: The code of the product.
        :param str delivery_date: The delivery date of the product.
        :param str delivery_way: The delivery way of the product.
        :param str image_url: The image URL of the product.
        :param str name: The name of the product.
        :param str operation_system: The operation system of the product.
        :param str score: The rating information of the product.
        :param str short_description: The short description of the product.
        :param str suggested_price: The suggested price of the product.
        :param float supplier_id: The supplier id of the product.
        :param str supplier_name: The supplier name of the product.
        :param str tags: The tags of the product.
        :param str target_url: The detail page URL of the product.
        :param str warranty_date: The warranty date of the product.
        """
        pulumi.set(__self__, "category_id", category_id)
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "delivery_date", delivery_date)
        pulumi.set(__self__, "delivery_way", delivery_way)
        pulumi.set(__self__, "image_url", image_url)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "operation_system", operation_system)
        pulumi.set(__self__, "score", score)
        pulumi.set(__self__, "short_description", short_description)
        pulumi.set(__self__, "suggested_price", suggested_price)
        pulumi.set(__self__, "supplier_id", supplier_id)
        pulumi.set(__self__, "supplier_name", supplier_name)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "target_url", target_url)
        pulumi.set(__self__, "warranty_date", warranty_date)

    @property
    @pulumi.getter(name="categoryId")
    def category_id(self) -> float:
        """
        The Category ID of products. For more information, see [DescribeProducts](https://help.aliyun.com/document_detail/89834.htm).
        """
        return pulumi.get(self, "category_id")

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        The code of the product.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter(name="deliveryDate")
    def delivery_date(self) -> str:
        """
        The delivery date of the product.
        """
        return pulumi.get(self, "delivery_date")

    @property
    @pulumi.getter(name="deliveryWay")
    def delivery_way(self) -> str:
        """
        The delivery way of the product.
        """
        return pulumi.get(self, "delivery_way")

    @property
    @pulumi.getter(name="imageUrl")
    def image_url(self) -> str:
        """
        The image URL of the product.
        """
        return pulumi.get(self, "image_url")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the product.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operationSystem")
    def operation_system(self) -> str:
        """
        The operation system of the product.
        """
        return pulumi.get(self, "operation_system")

    @property
    @pulumi.getter
    def score(self) -> str:
        """
        The rating information of the product.
        """
        return pulumi.get(self, "score")

    @property
    @pulumi.getter(name="shortDescription")
    def short_description(self) -> str:
        """
        The short description of the product.
        """
        return pulumi.get(self, "short_description")

    @property
    @pulumi.getter(name="suggestedPrice")
    def suggested_price(self) -> str:
        """
        The suggested price of the product.
        """
        return pulumi.get(self, "suggested_price")

    @property
    @pulumi.getter(name="supplierId")
    def supplier_id(self) -> float:
        """
        The supplier id of the product.
        """
        return pulumi.get(self, "supplier_id")

    @property
    @pulumi.getter(name="supplierName")
    def supplier_name(self) -> str:
        """
        The supplier name of the product.
        """
        return pulumi.get(self, "supplier_name")

    @property
    @pulumi.getter
    def tags(self) -> str:
        """
        The tags of the product.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetUrl")
    def target_url(self) -> str:
        """
        The detail page URL of the product.
        """
        return pulumi.get(self, "target_url")

    @property
    @pulumi.getter(name="warrantyDate")
    def warranty_date(self) -> str:
        """
        The warranty date of the product.
        """
        return pulumi.get(self, "warranty_date")


