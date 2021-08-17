// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Redis And Memcache (KVStore) Audit Log Config resource.
 *
 * > **NOTE:** Available in v1.130.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.kvstore.AuditLogConfig("example", {
 *     dbAudit: true,
 *     instanceId: "r-abc123455",
 *     retention: 1,
 * });
 * ```
 *
 * ## Import
 *
 * Redis And Memcache (KVStore) Audit Log Config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:kvstore/auditLogConfig:AuditLogConfig example <instance_id>
 * ```
 */
export class AuditLogConfig extends pulumi.CustomResource {
    /**
     * Get an existing AuditLogConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuditLogConfigState, opts?: pulumi.CustomResourceOptions): AuditLogConfig {
        return new AuditLogConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:kvstore/auditLogConfig:AuditLogConfig';

    /**
     * Returns true if the given object is an instance of AuditLogConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuditLogConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuditLogConfig.__pulumiType;
    }

    /**
     * Instance Creation Time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Indicates Whether to Enable the Audit Log.  Valid value: 
     * * true: Default Value, Open.
     * * false: Closed.
     */
    public readonly dbAudit!: pulumi.Output<boolean | undefined>;
    /**
     * Instance ID, Call the Describeinstances Get.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Audit Log Retention Period Value: 1~365.
     */
    public readonly retention!: pulumi.Output<number | undefined>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a AuditLogConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuditLogConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuditLogConfigArgs | AuditLogConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuditLogConfigState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["dbAudit"] = state ? state.dbAudit : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["retention"] = state ? state.retention : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AuditLogConfigArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["dbAudit"] = args ? args.dbAudit : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["retention"] = args ? args.retention : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuditLogConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuditLogConfig resources.
 */
export interface AuditLogConfigState {
    /**
     * Instance Creation Time.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * Indicates Whether to Enable the Audit Log.  Valid value: 
     * * true: Default Value, Open.
     * * false: Closed.
     */
    readonly dbAudit?: pulumi.Input<boolean>;
    /**
     * Instance ID, Call the Describeinstances Get.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * Audit Log Retention Period Value: 1~365.
     */
    readonly retention?: pulumi.Input<number>;
    /**
     * The status of the resource.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuditLogConfig resource.
 */
export interface AuditLogConfigArgs {
    /**
     * Indicates Whether to Enable the Audit Log.  Valid value: 
     * * true: Default Value, Open.
     * * false: Closed.
     */
    readonly dbAudit?: pulumi.Input<boolean>;
    /**
     * Instance ID, Call the Describeinstances Get.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * Audit Log Retention Period Value: 1~365.
     */
    readonly retention?: pulumi.Input<number>;
}
