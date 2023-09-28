// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Threat Detection Honeypot Preset resource.
 *
 * For information about Threat Detection Honeypot Preset and how to use it, see [What is Honeypot Preset](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createhoneypotpreset).
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
 * const name = config.get("name") || "tfexample";
 * const defaultHoneypotImages = alicloud.threatdetection.getHoneypotImages({
 *     nameRegex: "^ruoyi",
 * });
 * const defaultHoneypotNode = new alicloud.threatdetection.HoneypotNode("defaultHoneypotNode", {
 *     nodeName: name,
 *     availableProbeNum: 20,
 *     securityGroupProbeIpLists: ["0.0.0.0/0"],
 * });
 * const defaultHoneypotPreset = new alicloud.threatdetection.HoneypotPreset("defaultHoneypotPreset", {
 *     presetName: name,
 *     nodeId: defaultHoneypotNode.id,
 *     honeypotImageName: defaultHoneypotImages.then(defaultHoneypotImages => defaultHoneypotImages.images?.[0]?.honeypotImageName),
 *     meta: {
 *         portraitOption: true,
 *         burp: "open",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Threat Detection Honeypot Preset can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:threatdetection/honeypotPreset:HoneypotPreset example <id>
 * ```
 */
export class HoneypotPreset extends pulumi.CustomResource {
    /**
     * Get an existing HoneypotPreset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HoneypotPresetState, opts?: pulumi.CustomResourceOptions): HoneypotPreset {
        return new HoneypotPreset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:threatdetection/honeypotPreset:HoneypotPreset';

    /**
     * Returns true if the given object is an instance of HoneypotPreset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HoneypotPreset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HoneypotPreset.__pulumiType;
    }

    /**
     * Honeypot mirror name
     */
    public readonly honeypotImageName!: pulumi.Output<string>;
    /**
     * Unique ID of honeypot Template
     */
    public /*out*/ readonly honeypotPresetId!: pulumi.Output<string>;
    /**
     * Honeypot template custom parameters. See `meta` below.
     */
    public readonly meta!: pulumi.Output<outputs.threatdetection.HoneypotPresetMeta>;
    /**
     * Unique id of management node
     */
    public readonly nodeId!: pulumi.Output<string>;
    /**
     * Honeypot template custom name
     */
    public readonly presetName!: pulumi.Output<string>;

    /**
     * Create a HoneypotPreset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HoneypotPresetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HoneypotPresetArgs | HoneypotPresetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HoneypotPresetState | undefined;
            resourceInputs["honeypotImageName"] = state ? state.honeypotImageName : undefined;
            resourceInputs["honeypotPresetId"] = state ? state.honeypotPresetId : undefined;
            resourceInputs["meta"] = state ? state.meta : undefined;
            resourceInputs["nodeId"] = state ? state.nodeId : undefined;
            resourceInputs["presetName"] = state ? state.presetName : undefined;
        } else {
            const args = argsOrState as HoneypotPresetArgs | undefined;
            if ((!args || args.honeypotImageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'honeypotImageName'");
            }
            if ((!args || args.meta === undefined) && !opts.urn) {
                throw new Error("Missing required property 'meta'");
            }
            if ((!args || args.nodeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeId'");
            }
            if ((!args || args.presetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'presetName'");
            }
            resourceInputs["honeypotImageName"] = args ? args.honeypotImageName : undefined;
            resourceInputs["meta"] = args ? args.meta : undefined;
            resourceInputs["nodeId"] = args ? args.nodeId : undefined;
            resourceInputs["presetName"] = args ? args.presetName : undefined;
            resourceInputs["honeypotPresetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HoneypotPreset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HoneypotPreset resources.
 */
export interface HoneypotPresetState {
    /**
     * Honeypot mirror name
     */
    honeypotImageName?: pulumi.Input<string>;
    /**
     * Unique ID of honeypot Template
     */
    honeypotPresetId?: pulumi.Input<string>;
    /**
     * Honeypot template custom parameters. See `meta` below.
     */
    meta?: pulumi.Input<inputs.threatdetection.HoneypotPresetMeta>;
    /**
     * Unique id of management node
     */
    nodeId?: pulumi.Input<string>;
    /**
     * Honeypot template custom name
     */
    presetName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HoneypotPreset resource.
 */
export interface HoneypotPresetArgs {
    /**
     * Honeypot mirror name
     */
    honeypotImageName: pulumi.Input<string>;
    /**
     * Honeypot template custom parameters. See `meta` below.
     */
    meta: pulumi.Input<inputs.threatdetection.HoneypotPresetMeta>;
    /**
     * Unique id of management node
     */
    nodeId: pulumi.Input<string>;
    /**
     * Honeypot template custom name
     */
    presetName: pulumi.Input<string>;
}
