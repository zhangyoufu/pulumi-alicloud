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
    'GetZoneRecordsResult',
    'AwaitableGetZoneRecordsResult',
    'get_zone_records',
    'get_zone_records_output',
]

@pulumi.output_type
class GetZoneRecordsResult:
    """
    A collection of values returned by getZoneRecords.
    """
    def __init__(__self__, id=None, ids=None, keyword=None, lang=None, output_file=None, records=None, search_mode=None, status=None, tag=None, user_client_ip=None, zone_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if keyword and not isinstance(keyword, str):
            raise TypeError("Expected argument 'keyword' to be a str")
        pulumi.set(__self__, "keyword", keyword)
        if lang and not isinstance(lang, str):
            raise TypeError("Expected argument 'lang' to be a str")
        pulumi.set(__self__, "lang", lang)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if records and not isinstance(records, list):
            raise TypeError("Expected argument 'records' to be a list")
        pulumi.set(__self__, "records", records)
        if search_mode and not isinstance(search_mode, str):
            raise TypeError("Expected argument 'search_mode' to be a str")
        pulumi.set(__self__, "search_mode", search_mode)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if user_client_ip and not isinstance(user_client_ip, str):
            raise TypeError("Expected argument 'user_client_ip' to be a str")
        pulumi.set(__self__, "user_client_ip", user_client_ip)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

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
        """
        A list of Private Zone Record IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def keyword(self) -> Optional[str]:
        return pulumi.get(self, "keyword")

    @property
    @pulumi.getter
    def lang(self) -> Optional[str]:
        return pulumi.get(self, "lang")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def records(self) -> Sequence['outputs.GetZoneRecordsRecordResult']:
        """
        A list of zone records. Each element contains the following attributes:
        """
        return pulumi.get(self, "records")

    @property
    @pulumi.getter(name="searchMode")
    def search_mode(self) -> Optional[str]:
        return pulumi.get(self, "search_mode")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Status of the Private Zone Record.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="userClientIp")
    def user_client_ip(self) -> Optional[str]:
        return pulumi.get(self, "user_client_ip")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        return pulumi.get(self, "zone_id")


class AwaitableGetZoneRecordsResult(GetZoneRecordsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneRecordsResult(
            id=self.id,
            ids=self.ids,
            keyword=self.keyword,
            lang=self.lang,
            output_file=self.output_file,
            records=self.records,
            search_mode=self.search_mode,
            status=self.status,
            tag=self.tag,
            user_client_ip=self.user_client_ip,
            zone_id=self.zone_id)


def get_zone_records(ids: Optional[Sequence[str]] = None,
                     keyword: Optional[str] = None,
                     lang: Optional[str] = None,
                     output_file: Optional[str] = None,
                     search_mode: Optional[str] = None,
                     status: Optional[str] = None,
                     tag: Optional[str] = None,
                     user_client_ip: Optional[str] = None,
                     zone_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZoneRecordsResult:
    """
    This data source provides Private Zone Records resource information owned by an Alibaba Cloud account.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    records_ds = alicloud.pvtz.get_zone_records(zone_id=basic["id"],
        keyword=foo["value"])
    pulumi.export("firstRecordId", records_ds.records[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Private Zone Record IDs.
    :param str keyword: Keyword for record rr and value.
    :param str lang: User language.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str search_mode: Search mode. Value: 
           - LIKE: fuzzy search.
           - EXACT: precise search. It is not filled in by default.
    :param str status: Resolve record status. Value:
           - ENABLE: enable resolution.
           - DISABLE: pause parsing.
    :param str tag: It is not filled in by default, and queries the current zone resolution records. Fill in "ecs" to query the host name record list under the VPC associated with the current zone.
    :param str user_client_ip: User ip.
    :param str zone_id: ID of the Private Zone.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['keyword'] = keyword
    __args__['lang'] = lang
    __args__['outputFile'] = output_file
    __args__['searchMode'] = search_mode
    __args__['status'] = status
    __args__['tag'] = tag
    __args__['userClientIp'] = user_client_ip
    __args__['zoneId'] = zone_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:pvtz/getZoneRecords:getZoneRecords', __args__, opts=opts, typ=GetZoneRecordsResult).value

    return AwaitableGetZoneRecordsResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        keyword=pulumi.get(__ret__, 'keyword'),
        lang=pulumi.get(__ret__, 'lang'),
        output_file=pulumi.get(__ret__, 'output_file'),
        records=pulumi.get(__ret__, 'records'),
        search_mode=pulumi.get(__ret__, 'search_mode'),
        status=pulumi.get(__ret__, 'status'),
        tag=pulumi.get(__ret__, 'tag'),
        user_client_ip=pulumi.get(__ret__, 'user_client_ip'),
        zone_id=pulumi.get(__ret__, 'zone_id'))


@_utilities.lift_output_func(get_zone_records)
def get_zone_records_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            keyword: Optional[pulumi.Input[Optional[str]]] = None,
                            lang: Optional[pulumi.Input[Optional[str]]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            search_mode: Optional[pulumi.Input[Optional[str]]] = None,
                            status: Optional[pulumi.Input[Optional[str]]] = None,
                            tag: Optional[pulumi.Input[Optional[str]]] = None,
                            user_client_ip: Optional[pulumi.Input[Optional[str]]] = None,
                            zone_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZoneRecordsResult]:
    """
    This data source provides Private Zone Records resource information owned by an Alibaba Cloud account.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    records_ds = alicloud.pvtz.get_zone_records(zone_id=basic["id"],
        keyword=foo["value"])
    pulumi.export("firstRecordId", records_ds.records[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Private Zone Record IDs.
    :param str keyword: Keyword for record rr and value.
    :param str lang: User language.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str search_mode: Search mode. Value: 
           - LIKE: fuzzy search.
           - EXACT: precise search. It is not filled in by default.
    :param str status: Resolve record status. Value:
           - ENABLE: enable resolution.
           - DISABLE: pause parsing.
    :param str tag: It is not filled in by default, and queries the current zone resolution records. Fill in "ecs" to query the host name record list under the VPC associated with the current zone.
    :param str user_client_ip: User ip.
    :param str zone_id: ID of the Private Zone.
    """
    ...
