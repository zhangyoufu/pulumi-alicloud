// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./aggregateCompliancePack";
export * from "./aggregateConfigRule";
export * from "./aggregator";
export * from "./compliancePack";
export * from "./configurationRecorder";
export * from "./deliveryChannel";
export * from "./getAggregateCompliancePacks";
export * from "./getAggregateConfigRules";
export * from "./getAggregators";
export * from "./getCompliancePacks";
export * from "./getConfigurationRecorders";
export * from "./getDeliveryChannels";
export * from "./getRules";
export * from "./rule";

// Import resources to register:
import { AggregateCompliancePack } from "./aggregateCompliancePack";
import { AggregateConfigRule } from "./aggregateConfigRule";
import { Aggregator } from "./aggregator";
import { CompliancePack } from "./compliancePack";
import { ConfigurationRecorder } from "./configurationRecorder";
import { DeliveryChannel } from "./deliveryChannel";
import { Rule } from "./rule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:cfg/aggregateCompliancePack:AggregateCompliancePack":
                return new AggregateCompliancePack(name, <any>undefined, { urn })
            case "alicloud:cfg/aggregateConfigRule:AggregateConfigRule":
                return new AggregateConfigRule(name, <any>undefined, { urn })
            case "alicloud:cfg/aggregator:Aggregator":
                return new Aggregator(name, <any>undefined, { urn })
            case "alicloud:cfg/compliancePack:CompliancePack":
                return new CompliancePack(name, <any>undefined, { urn })
            case "alicloud:cfg/configurationRecorder:ConfigurationRecorder":
                return new ConfigurationRecorder(name, <any>undefined, { urn })
            case "alicloud:cfg/deliveryChannel:DeliveryChannel":
                return new DeliveryChannel(name, <any>undefined, { urn })
            case "alicloud:cfg/rule:Rule":
                return new Rule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "cfg/aggregateCompliancePack", _module)
pulumi.runtime.registerResourceModule("alicloud", "cfg/aggregateConfigRule", _module)
pulumi.runtime.registerResourceModule("alicloud", "cfg/aggregator", _module)
pulumi.runtime.registerResourceModule("alicloud", "cfg/compliancePack", _module)
pulumi.runtime.registerResourceModule("alicloud", "cfg/configurationRecorder", _module)
pulumi.runtime.registerResourceModule("alicloud", "cfg/deliveryChannel", _module)
pulumi.runtime.registerResourceModule("alicloud", "cfg/rule", _module)
