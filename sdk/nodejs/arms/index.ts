// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./alertContact";
export * from "./alertContactGroup";
export * from "./dispatchRule";
export * from "./getAlertContactGroups";
export * from "./getAlertContacts";
export * from "./getDispatchRules";
export * from "./getPrometheusAlertRules";
export * from "./prometheusAlertRule";

// Import resources to register:
import { AlertContact } from "./alertContact";
import { AlertContactGroup } from "./alertContactGroup";
import { DispatchRule } from "./dispatchRule";
import { PrometheusAlertRule } from "./prometheusAlertRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:arms/alertContact:AlertContact":
                return new AlertContact(name, <any>undefined, { urn })
            case "alicloud:arms/alertContactGroup:AlertContactGroup":
                return new AlertContactGroup(name, <any>undefined, { urn })
            case "alicloud:arms/dispatchRule:DispatchRule":
                return new DispatchRule(name, <any>undefined, { urn })
            case "alicloud:arms/prometheusAlertRule:PrometheusAlertRule":
                return new PrometheusAlertRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "arms/alertContact", _module)
pulumi.runtime.registerResourceModule("alicloud", "arms/alertContactGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "arms/dispatchRule", _module)
pulumi.runtime.registerResourceModule("alicloud", "arms/prometheusAlertRule", _module)
