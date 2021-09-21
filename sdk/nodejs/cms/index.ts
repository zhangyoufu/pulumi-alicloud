// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./alarm";
export * from "./alarmContact";
export * from "./alarmContactGroup";
export * from "./getAlarmContactGroups";
export * from "./getAlarmContacts";
export * from "./getGroupMetricRules";
export * from "./getMetricRuleTemplates";
export * from "./getMonitorGroupInstances";
export * from "./getMonitorGroups";
export * from "./getService";
export * from "./groupMetricRule";
export * from "./metricRuleTemplate";
export * from "./monitorGroup";
export * from "./monitorGroupInstances";
export * from "./siteMonitor";

// Import resources to register:
import { Alarm } from "./alarm";
import { AlarmContact } from "./alarmContact";
import { AlarmContactGroup } from "./alarmContactGroup";
import { GroupMetricRule } from "./groupMetricRule";
import { MetricRuleTemplate } from "./metricRuleTemplate";
import { MonitorGroup } from "./monitorGroup";
import { MonitorGroupInstances } from "./monitorGroupInstances";
import { SiteMonitor } from "./siteMonitor";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:cms/alarm:Alarm":
                return new Alarm(name, <any>undefined, { urn })
            case "alicloud:cms/alarmContact:AlarmContact":
                return new AlarmContact(name, <any>undefined, { urn })
            case "alicloud:cms/alarmContactGroup:AlarmContactGroup":
                return new AlarmContactGroup(name, <any>undefined, { urn })
            case "alicloud:cms/groupMetricRule:GroupMetricRule":
                return new GroupMetricRule(name, <any>undefined, { urn })
            case "alicloud:cms/metricRuleTemplate:MetricRuleTemplate":
                return new MetricRuleTemplate(name, <any>undefined, { urn })
            case "alicloud:cms/monitorGroup:MonitorGroup":
                return new MonitorGroup(name, <any>undefined, { urn })
            case "alicloud:cms/monitorGroupInstances:MonitorGroupInstances":
                return new MonitorGroupInstances(name, <any>undefined, { urn })
            case "alicloud:cms/siteMonitor:SiteMonitor":
                return new SiteMonitor(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "cms/alarm", _module)
pulumi.runtime.registerResourceModule("alicloud", "cms/alarmContact", _module)
pulumi.runtime.registerResourceModule("alicloud", "cms/alarmContactGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "cms/groupMetricRule", _module)
pulumi.runtime.registerResourceModule("alicloud", "cms/metricRuleTemplate", _module)
pulumi.runtime.registerResourceModule("alicloud", "cms/monitorGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "cms/monitorGroupInstances", _module)
pulumi.runtime.registerResourceModule("alicloud", "cms/siteMonitor", _module)
