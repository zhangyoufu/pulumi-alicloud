// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Threat Detection Honeypot Node resource.
 *
 * For information about Threat Detection Honeypot Node and how to use it, see [What is Honeypot Node](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createhoneypotnode).
 *
 * > **NOTE:** Available since v1.195.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf_example";
 * const _default = new alicloud.threatdetection.HoneypotNode("default", {
 *     nodeName: name,
 *     availableProbeNum: 20,
 *     securityGroupProbeIpLists: ["0.0.0.0/0"],
 * });
 * ```
 *
 * ## Import
 *
 * Threat Detection Honeypot Node can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:threatdetection/honeypotNode:HoneypotNode example <id>
 * ```
 */
export class HoneypotNode extends pulumi.CustomResource {
    /**
     * Get an existing HoneypotNode resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HoneypotNodeState, opts?: pulumi.CustomResourceOptions): HoneypotNode {
        return new HoneypotNode(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:threatdetection/honeypotNode:HoneypotNode';

    /**
     * Returns true if the given object is an instance of HoneypotNode.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HoneypotNode {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HoneypotNode.__pulumiType;
    }

    /**
     * Whether to allow honeypot access to the external network. Value:-**true**: Allow-**false**: Disabled
     */
    public readonly allowHoneypotAccessInternet!: pulumi.Output<boolean | undefined>;
    /**
     * Number of probes available.
     */
    public readonly availableProbeNum!: pulumi.Output<number>;
    /**
     * The creation time of the resource
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Management node name.
     */
    public readonly nodeName!: pulumi.Output<string>;
    /**
     * Release the collection of network segments.
     */
    public readonly securityGroupProbeIpLists!: pulumi.Output<string[] | undefined>;
    /**
     * The status of the resource
     */
    public /*out*/ readonly status!: pulumi.Output<number>;

    /**
     * Create a HoneypotNode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HoneypotNodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HoneypotNodeArgs | HoneypotNodeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HoneypotNodeState | undefined;
            resourceInputs["allowHoneypotAccessInternet"] = state ? state.allowHoneypotAccessInternet : undefined;
            resourceInputs["availableProbeNum"] = state ? state.availableProbeNum : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["nodeName"] = state ? state.nodeName : undefined;
            resourceInputs["securityGroupProbeIpLists"] = state ? state.securityGroupProbeIpLists : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as HoneypotNodeArgs | undefined;
            if ((!args || args.availableProbeNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'availableProbeNum'");
            }
            if ((!args || args.nodeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeName'");
            }
            resourceInputs["allowHoneypotAccessInternet"] = args ? args.allowHoneypotAccessInternet : undefined;
            resourceInputs["availableProbeNum"] = args ? args.availableProbeNum : undefined;
            resourceInputs["nodeName"] = args ? args.nodeName : undefined;
            resourceInputs["securityGroupProbeIpLists"] = args ? args.securityGroupProbeIpLists : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HoneypotNode.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HoneypotNode resources.
 */
export interface HoneypotNodeState {
    /**
     * Whether to allow honeypot access to the external network. Value:-**true**: Allow-**false**: Disabled
     */
    allowHoneypotAccessInternet?: pulumi.Input<boolean>;
    /**
     * Number of probes available.
     */
    availableProbeNum?: pulumi.Input<number>;
    /**
     * The creation time of the resource
     */
    createTime?: pulumi.Input<string>;
    /**
     * Management node name.
     */
    nodeName?: pulumi.Input<string>;
    /**
     * Release the collection of network segments.
     */
    securityGroupProbeIpLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the resource
     */
    status?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HoneypotNode resource.
 */
export interface HoneypotNodeArgs {
    /**
     * Whether to allow honeypot access to the external network. Value:-**true**: Allow-**false**: Disabled
     */
    allowHoneypotAccessInternet?: pulumi.Input<boolean>;
    /**
     * Number of probes available.
     */
    availableProbeNum: pulumi.Input<number>;
    /**
     * Management node name.
     */
    nodeName: pulumi.Input<string>;
    /**
     * Release the collection of network segments.
     */
    securityGroupProbeIpLists?: pulumi.Input<pulumi.Input<string>[]>;
}
