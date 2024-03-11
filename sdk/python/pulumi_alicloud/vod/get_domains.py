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
    'GetDomainsResult',
    'AwaitableGetDomainsResult',
    'get_domains',
    'get_domains_output',
]

@pulumi.output_type
class GetDomainsResult:
    """
    A collection of values returned by getDomains.
    """
    def __init__(__self__, domain_search_type=None, domains=None, id=None, ids=None, name_regex=None, names=None, output_file=None, status=None, tags=None):
        if domain_search_type and not isinstance(domain_search_type, str):
            raise TypeError("Expected argument 'domain_search_type' to be a str")
        pulumi.set(__self__, "domain_search_type", domain_search_type)
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="domainSearchType")
    def domain_search_type(self) -> Optional[str]:
        return pulumi.get(self, "domain_search_type")

    @property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.GetDomainsDomainResult']:
        return pulumi.get(self, "domains")

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
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableGetDomainsResult(GetDomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainsResult(
            domain_search_type=self.domain_search_type,
            domains=self.domains,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status,
            tags=self.tags)


def get_domains(domain_search_type: Optional[str] = None,
                ids: Optional[Sequence[str]] = None,
                name_regex: Optional[str] = None,
                output_file: Optional[str] = None,
                status: Optional[str] = None,
                tags: Optional[Mapping[str, Any]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainsResult:
    """
    This data source provides the Vod Domains of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.136.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_domain = alicloud.vod.Domain("defaultDomain",
        domain_name="your_domain_name",
        scope="domestic",
        sources=[alicloud.vod.DomainSourceArgs(
            source_type="domain",
            source_content="your_source_content",
            source_port="80",
        )],
        tags={
            "key1": "value1",
            "key2": "value2",
        })
    default_domains = alicloud.vod.get_domains_output(ids=[default_domain.id],
        tags={
            "key1": "value1",
            "key2": "value2",
        })
    pulumi.export("vodDomain", default_domains.domains[0])
    ```
    <!--End PulumiCodeChooser -->


    :param str domain_search_type: The search method. Valid values:
    :param Sequence[str] ids: A list of Domain IDs. Its element value is same as Domain Name.
    :param str name_regex: A regex string to filter results by Domain name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    """
    __args__ = dict()
    __args__['domainSearchType'] = domain_search_type
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:vod/getDomains:getDomains', __args__, opts=opts, typ=GetDomainsResult).value

    return AwaitableGetDomainsResult(
        domain_search_type=pulumi.get(__ret__, 'domain_search_type'),
        domains=pulumi.get(__ret__, 'domains'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_domains)
def get_domains_output(domain_search_type: Optional[pulumi.Input[Optional[str]]] = None,
                       ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       status: Optional[pulumi.Input[Optional[str]]] = None,
                       tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainsResult]:
    """
    This data source provides the Vod Domains of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.136.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    default_domain = alicloud.vod.Domain("defaultDomain",
        domain_name="your_domain_name",
        scope="domestic",
        sources=[alicloud.vod.DomainSourceArgs(
            source_type="domain",
            source_content="your_source_content",
            source_port="80",
        )],
        tags={
            "key1": "value1",
            "key2": "value2",
        })
    default_domains = alicloud.vod.get_domains_output(ids=[default_domain.id],
        tags={
            "key1": "value1",
            "key2": "value2",
        })
    pulumi.export("vodDomain", default_domains.domains[0])
    ```
    <!--End PulumiCodeChooser -->


    :param str domain_search_type: The search method. Valid values:
    :param Sequence[str] ids: A list of Domain IDs. Its element value is same as Domain Name.
    :param str name_regex: A regex string to filter results by Domain name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str status: The status of the resource.
    :param Mapping[str, Any] tags: A mapping of tags to assign to the resource.
    """
    ...
