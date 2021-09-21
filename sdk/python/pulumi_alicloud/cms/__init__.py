# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .alarm import *
from .alarm_contact import *
from .alarm_contact_group import *
from .get_alarm_contact_groups import *
from .get_alarm_contacts import *
from .get_group_metric_rules import *
from .get_metric_rule_templates import *
from .get_monitor_group_instances import *
from .get_monitor_groups import *
from .get_service import *
from .group_metric_rule import *
from .metric_rule_template import *
from .monitor_group import *
from .monitor_group_instances import *
from .site_monitor import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "alicloud:cms/alarm:Alarm":
                return Alarm(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cms/alarmContact:AlarmContact":
                return AlarmContact(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cms/alarmContactGroup:AlarmContactGroup":
                return AlarmContactGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cms/groupMetricRule:GroupMetricRule":
                return GroupMetricRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cms/metricRuleTemplate:MetricRuleTemplate":
                return MetricRuleTemplate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cms/monitorGroup:MonitorGroup":
                return MonitorGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cms/monitorGroupInstances:MonitorGroupInstances":
                return MonitorGroupInstances(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:cms/siteMonitor:SiteMonitor":
                return SiteMonitor(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("alicloud", "cms/alarm", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cms/alarmContact", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cms/alarmContactGroup", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cms/groupMetricRule", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cms/metricRuleTemplate", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cms/monitorGroup", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cms/monitorGroupInstances", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "cms/siteMonitor", _module_instance)

_register_module()
