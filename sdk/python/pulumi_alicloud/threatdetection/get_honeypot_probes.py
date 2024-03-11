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
    'GetHoneypotProbesResult',
    'AwaitableGetHoneypotProbesResult',
    'get_honeypot_probes',
    'get_honeypot_probes_output',
]

@pulumi.output_type
class GetHoneypotProbesResult:
    """
    A collection of values returned by getHoneypotProbes.
    """
    def __init__(__self__, display_name=None, enable_details=None, id=None, ids=None, name_regex=None, output_file=None, probe_status=None, probe_type=None, probes=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if probe_status and not isinstance(probe_status, str):
            raise TypeError("Expected argument 'probe_status' to be a str")
        pulumi.set(__self__, "probe_status", probe_status)
        if probe_type and not isinstance(probe_type, str):
            raise TypeError("Expected argument 'probe_type' to be a str")
        pulumi.set(__self__, "probe_type", probe_type)
        if probes and not isinstance(probes, list):
            raise TypeError("Expected argument 'probes' to be a list")
        pulumi.set(__self__, "probes", probes)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        Probe name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

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
        A list of Honeypot Probe IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="probeStatus")
    def probe_status(self) -> Optional[str]:
        return pulumi.get(self, "probe_status")

    @property
    @pulumi.getter(name="probeType")
    def probe_type(self) -> Optional[str]:
        """
        Probe type, support `host_probe` and `vpc_black_hole_probe`.
        """
        return pulumi.get(self, "probe_type")

    @property
    @pulumi.getter
    def probes(self) -> Sequence['outputs.GetHoneypotProbesProbeResult']:
        """
        A list of Honeypot Probe Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "probes")


class AwaitableGetHoneypotProbesResult(GetHoneypotProbesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHoneypotProbesResult(
            display_name=self.display_name,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            probe_status=self.probe_status,
            probe_type=self.probe_type,
            probes=self.probes)


def get_honeypot_probes(display_name: Optional[str] = None,
                        enable_details: Optional[bool] = None,
                        ids: Optional[Sequence[str]] = None,
                        name_regex: Optional[str] = None,
                        output_file: Optional[str] = None,
                        probe_status: Optional[str] = None,
                        probe_type: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHoneypotProbesResult:
    """
    This data source provides Threat Detection Honeypot Probe available to the user.[What is Honeypot Probe](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createhoneypotprobe)

    > **NOTE:** Available in 1.195.0+

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "tf-testAccThreatDetectionHoneypotProbe"
    default_honeypot_probe = alicloud.threatdetection.HoneypotProbe("defaultHoneypotProbe",
        uuid="e52c7872-29d1-4aa1-9908-0299abd53606",
        probe_type="host_probe",
        control_node_id="e1397077-4941-4b14-b533-ca2bdebd00a3",
        ping=True,
        honeypot_bind_lists=[alicloud.threatdetection.HoneypotProbeHoneypotBindListArgs(
            bind_port_lists=[alicloud.threatdetection.HoneypotProbeHoneypotBindListBindPortListArgs(
                start_port=80,
                end_port=80,
            )],
            honeypot_id="4925bf9784de992ecd017ad051528a03b3927ef814eeff76c2ebb3ab9a84bf05",
        )],
        display_name=name,
        arp=True)
    default_honeypot_probes = alicloud.threatdetection.get_honeypot_probes_output(ids=[default_honeypot_probe.id],
        display_name=name,
        probe_type="host_probe",
        enable_details=True)
    pulumi.export("alicloudThreatDetectionHoneypotProbeExampleId", default_honeypot_probes.probes[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str display_name: Probe name
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Honeypot Probe IDs.
    :param str name_regex: A regex string to filter results by display name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str probe_type: Probe type
    """
    __args__ = dict()
    __args__['displayName'] = display_name
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['probeStatus'] = probe_status
    __args__['probeType'] = probe_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:threatdetection/getHoneypotProbes:getHoneypotProbes', __args__, opts=opts, typ=GetHoneypotProbesResult).value

    return AwaitableGetHoneypotProbesResult(
        display_name=pulumi.get(__ret__, 'display_name'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        probe_status=pulumi.get(__ret__, 'probe_status'),
        probe_type=pulumi.get(__ret__, 'probe_type'),
        probes=pulumi.get(__ret__, 'probes'))


@_utilities.lift_output_func(get_honeypot_probes)
def get_honeypot_probes_output(display_name: Optional[pulumi.Input[Optional[str]]] = None,
                               enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                               ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                               output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               probe_status: Optional[pulumi.Input[Optional[str]]] = None,
                               probe_type: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHoneypotProbesResult]:
    """
    This data source provides Threat Detection Honeypot Probe available to the user.[What is Honeypot Probe](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createhoneypotprobe)

    > **NOTE:** Available in 1.195.0+

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "tf-testAccThreatDetectionHoneypotProbe"
    default_honeypot_probe = alicloud.threatdetection.HoneypotProbe("defaultHoneypotProbe",
        uuid="e52c7872-29d1-4aa1-9908-0299abd53606",
        probe_type="host_probe",
        control_node_id="e1397077-4941-4b14-b533-ca2bdebd00a3",
        ping=True,
        honeypot_bind_lists=[alicloud.threatdetection.HoneypotProbeHoneypotBindListArgs(
            bind_port_lists=[alicloud.threatdetection.HoneypotProbeHoneypotBindListBindPortListArgs(
                start_port=80,
                end_port=80,
            )],
            honeypot_id="4925bf9784de992ecd017ad051528a03b3927ef814eeff76c2ebb3ab9a84bf05",
        )],
        display_name=name,
        arp=True)
    default_honeypot_probes = alicloud.threatdetection.get_honeypot_probes_output(ids=[default_honeypot_probe.id],
        display_name=name,
        probe_type="host_probe",
        enable_details=True)
    pulumi.export("alicloudThreatDetectionHoneypotProbeExampleId", default_honeypot_probes.probes[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str display_name: Probe name
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Honeypot Probe IDs.
    :param str name_regex: A regex string to filter results by display name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str probe_type: Probe type
    """
    ...
