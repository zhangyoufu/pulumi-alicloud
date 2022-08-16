// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./alias";
export * from "./customDomain";
export * from "./function";
export * from "./functionAsyncInvokeConfig";
export * from "./getCustomDomains";
export * from "./getFunctions";
export * from "./getService";
export * from "./getServices";
export * from "./getTriggers";
export * from "./getZones";
export * from "./layerVersion";
export * from "./service";
export * from "./trigger";

// Import resources to register:
import { Alias } from "./alias";
import { CustomDomain } from "./customDomain";
import { Function } from "./function";
import { FunctionAsyncInvokeConfig } from "./functionAsyncInvokeConfig";
import { LayerVersion } from "./layerVersion";
import { Service } from "./service";
import { Trigger } from "./trigger";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:fc/alias:Alias":
                return new Alias(name, <any>undefined, { urn })
            case "alicloud:fc/customDomain:CustomDomain":
                return new CustomDomain(name, <any>undefined, { urn })
            case "alicloud:fc/function:Function":
                return new Function(name, <any>undefined, { urn })
            case "alicloud:fc/functionAsyncInvokeConfig:FunctionAsyncInvokeConfig":
                return new FunctionAsyncInvokeConfig(name, <any>undefined, { urn })
            case "alicloud:fc/layerVersion:LayerVersion":
                return new LayerVersion(name, <any>undefined, { urn })
            case "alicloud:fc/service:Service":
                return new Service(name, <any>undefined, { urn })
            case "alicloud:fc/trigger:Trigger":
                return new Trigger(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "fc/alias", _module)
pulumi.runtime.registerResourceModule("alicloud", "fc/customDomain", _module)
pulumi.runtime.registerResourceModule("alicloud", "fc/function", _module)
pulumi.runtime.registerResourceModule("alicloud", "fc/functionAsyncInvokeConfig", _module)
pulumi.runtime.registerResourceModule("alicloud", "fc/layerVersion", _module)
pulumi.runtime.registerResourceModule("alicloud", "fc/service", _module)
pulumi.runtime.registerResourceModule("alicloud", "fc/trigger", _module)
