// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an ONS instance resource.
 * 
 * For more information about how to use it, see [RocketMQ Instance Management API](https://www.alibabacloud.com/help/doc-detail/106354.html). 
 * 
 * > **NOTE:** The number of instances in the same region cannot exceed 8. At present, the resource does not support region "mq-internet-access" and "ap-southeast-5". 
 * 
 * > **NOTE:** Available in 1.51.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const example = new alicloud.rocketmq.Instance("example", {
 *     remark: "tf-example-ons-instance-remark",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ons_instance.html.markdown.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rocketmq/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
     */
    public /*out*/ readonly instanceStatus!: pulumi.Output<number>;
    /**
     * The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.
     */
    public /*out*/ readonly instanceType!: pulumi.Output<number>;
    /**
     * Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Platinum edition instance expiration time.
     */
    public /*out*/ readonly releaseTime!: pulumi.Output<string>;
    /**
     * This attribute is a concise description of instance. The length cannot exceed 128.
     */
    public readonly remark!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["instanceStatus"] = state ? state.instanceStatus : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["releaseTime"] = state ? state.releaseTime : undefined;
            inputs["remark"] = state ? state.remark : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["remark"] = args ? args.remark : undefined;
            inputs["instanceStatus"] = undefined /*out*/;
            inputs["instanceType"] = undefined /*out*/;
            inputs["releaseTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
     */
    readonly instanceStatus?: pulumi.Input<number>;
    /**
     * The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.
     */
    readonly instanceType?: pulumi.Input<number>;
    /**
     * Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Platinum edition instance expiration time.
     */
    readonly releaseTime?: pulumi.Input<string>;
    /**
     * This attribute is a concise description of instance. The length cannot exceed 128.
     */
    readonly remark?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * This attribute is a concise description of instance. The length cannot exceed 128.
     */
    readonly remark?: pulumi.Input<string>;
}
