// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CEN child instance grant resource, which allow you to authorize a VPC or VBR to a CEN of a different account.
 * 
 * For more information about how to use it, see [Attach a network in a different account](https://www.alibabacloud.com/help/doc-detail/73645.htm). 
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testAccCenInstanceGrantBasic";
 * 
 * // Create a new instance-grant and use it to grant one child instance of account1 to a new CEN of account 2.
 * const account1 = new alicloud.Provider("account1", {
 *     accessKey: "access123",
 *     secretKey: "secret123",
 * });
 * const account2 = new alicloud.Provider("account2", {
 *     accessKey: "access456",
 *     secretKey: "secret456",
 * });
 * const cen = new alicloud.cen.Instance("cen", {}, {provider: account2});
 * const vpc = new alicloud.vpc.Network("vpc", {
 *     cidrBlock: "192.168.0.0/16",
 * }, {provider: account1});
 * const fooInstanceGrant = new alicloud.cen.InstanceGrant("foo", {
 *     cenId: cen.id,
 *     cenOwnerId: "uid2",
 *     childInstanceId: vpc.id,
 * }, {provider: account1});
 * const fooInstanceAttachment = new alicloud.cen.InstanceAttachment("foo", {
 *     childInstanceId: vpc.id,
 *     childInstanceOwnerId: "uid1",
 *     childInstanceRegionId: "cn-qingdao",
 *     instanceId: cen.id,
 * }, {provider: account2,dependsOn: [fooInstanceGrant]});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cen_instance_grant.html.markdown.
 */
export class InstanceGrant extends pulumi.CustomResource {
    /**
     * Get an existing InstanceGrant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceGrantState, opts?: pulumi.CustomResourceOptions): InstanceGrant {
        return new InstanceGrant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/instanceGrant:InstanceGrant';

    /**
     * Returns true if the given object is an instance of InstanceGrant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceGrant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceGrant.__pulumiType;
    }

    /**
     * The ID of the CEN.
     */
    public readonly cenId!: pulumi.Output<string>;
    /**
     * The owner UID of the  CEN which the child instance granted to.
     */
    public readonly cenOwnerId!: pulumi.Output<string>;
    /**
     * The ID of the child instance to grant.
     */
    public readonly childInstanceId!: pulumi.Output<string>;

    /**
     * Create a InstanceGrant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceGrantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceGrantArgs | InstanceGrantState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceGrantState | undefined;
            inputs["cenId"] = state ? state.cenId : undefined;
            inputs["cenOwnerId"] = state ? state.cenOwnerId : undefined;
            inputs["childInstanceId"] = state ? state.childInstanceId : undefined;
        } else {
            const args = argsOrState as InstanceGrantArgs | undefined;
            if (!args || args.cenId === undefined) {
                throw new Error("Missing required property 'cenId'");
            }
            if (!args || args.cenOwnerId === undefined) {
                throw new Error("Missing required property 'cenOwnerId'");
            }
            if (!args || args.childInstanceId === undefined) {
                throw new Error("Missing required property 'childInstanceId'");
            }
            inputs["cenId"] = args ? args.cenId : undefined;
            inputs["cenOwnerId"] = args ? args.cenOwnerId : undefined;
            inputs["childInstanceId"] = args ? args.childInstanceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceGrant.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceGrant resources.
 */
export interface InstanceGrantState {
    /**
     * The ID of the CEN.
     */
    readonly cenId?: pulumi.Input<string>;
    /**
     * The owner UID of the  CEN which the child instance granted to.
     */
    readonly cenOwnerId?: pulumi.Input<string>;
    /**
     * The ID of the child instance to grant.
     */
    readonly childInstanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceGrant resource.
 */
export interface InstanceGrantArgs {
    /**
     * The ID of the CEN.
     */
    readonly cenId: pulumi.Input<string>;
    /**
     * The owner UID of the  CEN which the child instance granted to.
     */
    readonly cenOwnerId: pulumi.Input<string>;
    /**
     * The ID of the child instance to grant.
     */
    readonly childInstanceId: pulumi.Input<string>;
}
