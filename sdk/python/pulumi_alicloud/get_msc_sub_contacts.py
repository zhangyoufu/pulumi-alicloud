# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetMscSubContactsResult',
    'AwaitableGetMscSubContactsResult',
    'get_msc_sub_contacts',
    'get_msc_sub_contacts_output',
]

@pulumi.output_type
class GetMscSubContactsResult:
    """
    A collection of values returned by getMscSubContacts.
    """
    def __init__(__self__, contacts=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if contacts and not isinstance(contacts, list):
            raise TypeError("Expected argument 'contacts' to be a list")
        pulumi.set(__self__, "contacts", contacts)
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

    @property
    @pulumi.getter
    def contacts(self) -> Sequence['outputs.GetMscSubContactsContactResult']:
        return pulumi.get(self, "contacts")

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


class AwaitableGetMscSubContactsResult(GetMscSubContactsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMscSubContactsResult(
            contacts=self.contacts,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)


def get_msc_sub_contacts(ids: Optional[Sequence[str]] = None,
                         name_regex: Optional[str] = None,
                         output_file: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMscSubContactsResult:
    """
    This data source provides the Message Center Contacts of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.132.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.get_msc_sub_contacts()
    pulumi.export("mscSubContactId1", ids.contacts[0].id)
    name_regex = alicloud.get_msc_sub_contacts(name_regex="^my-Contact")
    pulumi.export("mscSubContactId2", name_regex.contacts[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Contact IDs.
    :param str name_regex: A regex string to filter results by Contact name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:index/getMscSubContacts:getMscSubContacts', __args__, opts=opts, typ=GetMscSubContactsResult).value

    return AwaitableGetMscSubContactsResult(
        contacts=pulumi.get(__ret__, 'contacts'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'))


@_utilities.lift_output_func(get_msc_sub_contacts)
def get_msc_sub_contacts_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMscSubContactsResult]:
    """
    This data source provides the Message Center Contacts of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.132.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.get_msc_sub_contacts()
    pulumi.export("mscSubContactId1", ids.contacts[0].id)
    name_regex = alicloud.get_msc_sub_contacts(name_regex="^my-Contact")
    pulumi.export("mscSubContactId2", name_regex.contacts[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] ids: A list of Contact IDs.
    :param str name_regex: A regex string to filter results by Contact name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    """
    ...
