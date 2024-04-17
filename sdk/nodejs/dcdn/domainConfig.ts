// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a DCDN Accelerated Domain resource.
 *
 * For information about domain config and how to use it, see [Batch set config](https://www.alibabacloud.com/help/en/doc-detail/130632.htm).
 *
 * > **NOTE:** Available since v1.131.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const domainName = config.get("domainName") || "alibaba-example.com";
 * const example = new alicloud.dcdn.Domain("example", {
 *     domainName: domainName,
 *     scope: "overseas",
 *     status: "online",
 *     sources: [{
 *         content: "1.1.1.1",
 *         type: "ipaddr",
 *         priority: "20",
 *         port: 80,
 *         weight: "10",
 *     }],
 * });
 * const ipAllowListSet = new alicloud.dcdn.DomainConfig("ip_allow_list_set", {
 *     domainName: example.domainName,
 *     functionName: "ip_allow_list_set",
 *     functionArgs: [{
 *         argName: "ip_list",
 *         argValue: "192.168.0.1",
 *     }],
 * });
 * const refererWhiteListSet = new alicloud.dcdn.DomainConfig("referer_white_list_set", {
 *     domainName: example.domainName,
 *     functionName: "referer_white_list_set",
 *     functionArgs: [{
 *         argName: "refer_domain_allow_list",
 *         argValue: "110.110.110.110",
 *     }],
 * });
 * const filetypeBasedTtlSet = new alicloud.dcdn.DomainConfig("filetype_based_ttl_set", {
 *     domainName: example.domainName,
 *     functionName: "filetype_based_ttl_set",
 *     functionArgs: [
 *         {
 *             argName: "ttl",
 *             argValue: "300",
 *         },
 *         {
 *             argName: "file_type",
 *             argValue: "jpg",
 *         },
 *         {
 *             argName: "weight",
 *             argValue: "1",
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * DCDN domain config can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:dcdn/domainConfig:DomainConfig example <domain_name>:<function_name>:<config_id>
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
    public static readonly __pulumiType = 'alicloud:dcdn/domainConfig:DomainConfig';

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
     * The ID of the configuration.
     */
    public /*out*/ readonly configId!: pulumi.Output<string>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The args of the domain config. See `functionArgs` below.
     */
    public readonly functionArgs!: pulumi.Output<outputs.dcdn.DomainConfigFunctionArg[]>;
    /**
     * The name of the domain config.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
     */
    public readonly parentId!: pulumi.Output<string>;
    /**
     * The status of the Config.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainConfigState | undefined;
            resourceInputs["configId"] = state ? state.configId : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["functionArgs"] = state ? state.functionArgs : undefined;
            resourceInputs["functionName"] = state ? state.functionName : undefined;
            resourceInputs["parentId"] = state ? state.parentId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
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
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["functionArgs"] = args ? args.functionArgs : undefined;
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["parentId"] = args ? args.parentId : undefined;
            resourceInputs["configId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainConfig resources.
 */
export interface DomainConfigState {
    /**
     * The ID of the configuration.
     */
    configId?: pulumi.Input<string>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The args of the domain config. See `functionArgs` below.
     */
    functionArgs?: pulumi.Input<pulumi.Input<inputs.dcdn.DomainConfigFunctionArg>[]>;
    /**
     * The name of the domain config.
     */
    functionName?: pulumi.Input<string>;
    /**
     * By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
     */
    parentId?: pulumi.Input<string>;
    /**
     * The status of the Config.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainConfig resource.
 */
export interface DomainConfigArgs {
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    domainName: pulumi.Input<string>;
    /**
     * The args of the domain config. See `functionArgs` below.
     */
    functionArgs: pulumi.Input<pulumi.Input<inputs.dcdn.DomainConfigFunctionArg>[]>;
    /**
     * The name of the domain config.
     */
    functionName: pulumi.Input<string>;
    /**
     * By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
     */
    parentId?: pulumi.Input<string>;
}
