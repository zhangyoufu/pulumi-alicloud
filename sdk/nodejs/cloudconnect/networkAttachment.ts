// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Connect Network Attachment resource. This topic describes how to associate a Smart Access Gateway (SAG) instance with a network instance. You must associate an SAG instance with a network instance if you want to connect the SAG to Alibaba Cloud. You can connect an SAG to Alibaba Cloud through a leased line, the Internet, or the active and standby links.
 *
 * For information about Cloud Connect Network Attachment and how to use it, see [What is Cloud Connect Network Attachment](https://www.alibabacloud.com/help/doc-detail/124230.htm).
 *
 * > **NOTE:** Available in 1.64.0+
 *
 * > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ccn = new alicloud.cloudconnect.Network("ccn", {isDefault: "true"});
 * const _default = new alicloud.cloudconnect.NetworkAttachment("default", {
 *     ccnId: ccn.id,
 *     sagId: "sag-xxxxx",
 * }, {
 *     dependsOn: [ccn],
 * });
 * ```
 *
 * ## Import
 *
 * The Cloud Connect Network Attachment can be imported using the instance_id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cloudconnect/networkAttachment:NetworkAttachment example ccn-abc123456:sag-abc123456
 * ```
 */
export class NetworkAttachment extends pulumi.CustomResource {
    /**
     * Get an existing NetworkAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkAttachmentState, opts?: pulumi.CustomResourceOptions): NetworkAttachment {
        return new NetworkAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudconnect/networkAttachment:NetworkAttachment';

    /**
     * Returns true if the given object is an instance of NetworkAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkAttachment.__pulumiType;
    }

    /**
     * The ID of the CCN instance.
     */
    public readonly ccnId!: pulumi.Output<string>;
    /**
     * The ID of the Smart Access Gateway instance.
     */
    public readonly sagId!: pulumi.Output<string>;

    /**
     * Create a NetworkAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkAttachmentArgs | NetworkAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkAttachmentState | undefined;
            inputs["ccnId"] = state ? state.ccnId : undefined;
            inputs["sagId"] = state ? state.sagId : undefined;
        } else {
            const args = argsOrState as NetworkAttachmentArgs | undefined;
            if ((!args || args.ccnId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ccnId'");
            }
            if ((!args || args.sagId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sagId'");
            }
            inputs["ccnId"] = args ? args.ccnId : undefined;
            inputs["sagId"] = args ? args.sagId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NetworkAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkAttachment resources.
 */
export interface NetworkAttachmentState {
    /**
     * The ID of the CCN instance.
     */
    readonly ccnId?: pulumi.Input<string>;
    /**
     * The ID of the Smart Access Gateway instance.
     */
    readonly sagId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkAttachment resource.
 */
export interface NetworkAttachmentArgs {
    /**
     * The ID of the CCN instance.
     */
    readonly ccnId: pulumi.Input<string>;
    /**
     * The ID of the Smart Access Gateway instance.
     */
    readonly sagId: pulumi.Input<string>;
}
