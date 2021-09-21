// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a CDN Accelerated Domain resource.
 *
 * For information about domain config and how to use it, see [Batch set config](https://www.alibabacloud.com/help/zh/doc-detail/90915.htm)
 *
 * > **NOTE:** Available in v1.34.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a new Domain config.
 * const domain = new alicloud.cdn.DomainNew("domain", {
 *     domainName: "mycdndomain.xiaozhu.com",
 *     cdnType: "web",
 *     scope: "overseas",
 *     sources: [{
 *         content: "1.1.1.1",
 *         type: "ipaddr",
 *         priority: "20",
 *         port: 80,
 *         weight: "15",
 *     }],
 * });
 * const config = new alicloud.cdn.DomainConfig("config", {
 *     domainName: domain.domainName,
 *     functionName: "ip_allow_list_set",
 *     functionArgs: [{
 *         argName: "ip_list",
 *         argValue: "110.110.110.110",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * CDN domain config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cdn/domainConfig:DomainConfig example <domain_name>:<function_name>:<config_id>
 * ```
 *
 * ```sh
 *  $ pulumi import alicloud:cdn/domainConfig:DomainConfig example <domain_name>:<function_name>
 * ```
 */
export class DomainConfig extends pulumi.CustomResource {
    /**
     * Get an existing DomainConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainConfigState, opts?: pulumi.CustomResourceOptions): DomainConfig {
        return new DomainConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cdn/domainConfig:DomainConfig';

    /**
     * Returns true if the given object is an instance of DomainConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainConfig.__pulumiType;
    }

    /**
     * (Available in 1.132.0+) The ID of the domain config function.
     */
    public /*out*/ readonly configId!: pulumi.Output<string>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The args of the domain config.
     */
    public readonly functionArgs!: pulumi.Output<outputs.cdn.DomainConfigFunctionArg[]>;
    /**
     * The name of the domain config.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * (Available in 1.132.0+) The Status of the function. Valid values: `success`, `testing`, `failed`, and `configuring`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a DomainConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainConfigArgs | DomainConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainConfigState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["functionArgs"] = state ? state.functionArgs : undefined;
            inputs["functionName"] = state ? state.functionName : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as DomainConfigArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.functionArgs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionArgs'");
            }
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["functionArgs"] = args ? args.functionArgs : undefined;
            inputs["functionName"] = args ? args.functionName : undefined;
            inputs["configId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DomainConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainConfig resources.
 */
export interface DomainConfigState {
    /**
     * (Available in 1.132.0+) The ID of the domain config function.
     */
    readonly configId?: pulumi.Input<string>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * The args of the domain config.
     */
    readonly functionArgs?: pulumi.Input<pulumi.Input<inputs.cdn.DomainConfigFunctionArg>[]>;
    /**
     * The name of the domain config.
     */
    readonly functionName?: pulumi.Input<string>;
    /**
     * (Available in 1.132.0+) The Status of the function. Valid values: `success`, `testing`, `failed`, and `configuring`.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainConfig resource.
 */
export interface DomainConfigArgs {
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    readonly domainName: pulumi.Input<string>;
    /**
     * The args of the domain config.
     */
    readonly functionArgs: pulumi.Input<pulumi.Input<inputs.cdn.DomainConfigFunctionArg>[]>;
    /**
     * The name of the domain config.
     */
    readonly functionName: pulumi.Input<string>;
}
