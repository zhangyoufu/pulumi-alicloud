// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Sag Qos resource. Smart Access Gateway (SAG) supports quintuple-based QoS functions to differentiate traffic of different services and ensure high-priority traffic bandwidth.
 *
 * For information about Sag Qos and how to use it, see [What is Qos](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/createqos).
 *
 * > **NOTE:** Available since v1.60.0.
 *
 * > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
 * const _default = new alicloud.rocketmq.Qos("default", {});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * The Sag Qos can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:rocketmq/qos:Qos example qos-abc123456
 * ```
 */
export class Qos extends pulumi.CustomResource {
    /**
     * Get an existing Qos resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QosState, opts?: pulumi.CustomResourceOptions): Qos {
        return new Qos(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rocketmq/qos:Qos';

    /**
     * Returns true if the given object is an instance of Qos.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Qos {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Qos.__pulumiType;
    }

    /**
     * The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Qos resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: QosArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QosArgs | QosState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QosState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as QosArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Qos.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Qos resources.
 */
export interface QosState {
    /**
     * The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Qos resource.
 */
export interface QosArgs {
    /**
     * The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
     */
    name?: pulumi.Input<string>;
}
