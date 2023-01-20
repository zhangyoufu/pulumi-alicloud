// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DMS Enterprise Proxy Access resource.
 *
 * For information about DMS Enterprise Proxy Access and how to use it, see [What is Proxy Access](https://www.alibabacloud.com/help/zh/data-management-service/latest/api-doc-dms-enterprise-2018-11-01-api-doc-createproxyaccess).
 *
 * > **NOTE:** Available in v1.195.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.dms.EnterpriseProxyAccess("default", {
 *     indepAccount: "dmstest",
 *     indepPassword: "PASSWORD-DEMO",
 *     proxyId: "1881",
 *     userId: "104442",
 * });
 * ```
 *
 * ## Import
 *
 * DMS Enterprise Proxy Access can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:dms/enterpriseProxyAccess:EnterpriseProxyAccess example <id>
 * ```
 */
export class EnterpriseProxyAccess extends pulumi.CustomResource {
    /**
     * Get an existing EnterpriseProxyAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnterpriseProxyAccessState, opts?: pulumi.CustomResourceOptions): EnterpriseProxyAccess {
        return new EnterpriseProxyAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dms/enterpriseProxyAccess:EnterpriseProxyAccess';

    /**
     * Returns true if the given object is an instance of EnterpriseProxyAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnterpriseProxyAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnterpriseProxyAccess.__pulumiType;
    }

    /**
     * The authorized account of the security agent.
     */
    public /*out*/ readonly accessId!: pulumi.Output<string>;
    /**
     * Secure access agent authorization password.
     */
    public /*out*/ readonly accessSecret!: pulumi.Output<string>;
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Database account.
     */
    public readonly indepAccount!: pulumi.Output<string | undefined>;
    /**
     * Database password.
     */
    public readonly indepPassword!: pulumi.Output<string | undefined>;
    /**
     * The ID of the instance.
     */
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    /**
     * The source information of the security access agent permission is enabled, and the return value is as follows:
     * * **Owner Authorization**: The UID of the owner in parentheses.
     * * **Work Order Authorization**: The ticket number in parentheses is the number of the user to apply for permission.
     */
    public /*out*/ readonly originInfo!: pulumi.Output<string>;
    /**
     * Security Protection authorization ID. After the target user is authorized by the security protection agent, the system automatically generates a security protection authorization ID, which is globally unique.
     */
    public readonly proxyAccessId!: pulumi.Output<string>;
    /**
     * The ID of the security agent.
     */
    public readonly proxyId!: pulumi.Output<string>;
    /**
     * The user ID.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * User nickname.
     */
    public /*out*/ readonly userName!: pulumi.Output<string>;
    /**
     * User UID.
     */
    public /*out*/ readonly userUid!: pulumi.Output<string>;

    /**
     * Create a EnterpriseProxyAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnterpriseProxyAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnterpriseProxyAccessArgs | EnterpriseProxyAccessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EnterpriseProxyAccessState | undefined;
            resourceInputs["accessId"] = state ? state.accessId : undefined;
            resourceInputs["accessSecret"] = state ? state.accessSecret : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["indepAccount"] = state ? state.indepAccount : undefined;
            resourceInputs["indepPassword"] = state ? state.indepPassword : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["originInfo"] = state ? state.originInfo : undefined;
            resourceInputs["proxyAccessId"] = state ? state.proxyAccessId : undefined;
            resourceInputs["proxyId"] = state ? state.proxyId : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["userUid"] = state ? state.userUid : undefined;
        } else {
            const args = argsOrState as EnterpriseProxyAccessArgs | undefined;
            if ((!args || args.proxyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'proxyId'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["indepAccount"] = args ? args.indepAccount : undefined;
            resourceInputs["indepPassword"] = args?.indepPassword ? pulumi.secret(args.indepPassword) : undefined;
            resourceInputs["proxyAccessId"] = args ? args.proxyAccessId : undefined;
            resourceInputs["proxyId"] = args ? args.proxyId : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["accessId"] = undefined /*out*/;
            resourceInputs["accessSecret"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["originInfo"] = undefined /*out*/;
            resourceInputs["userName"] = undefined /*out*/;
            resourceInputs["userUid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessSecret", "indepPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(EnterpriseProxyAccess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EnterpriseProxyAccess resources.
 */
export interface EnterpriseProxyAccessState {
    /**
     * The authorized account of the security agent.
     */
    accessId?: pulumi.Input<string>;
    /**
     * Secure access agent authorization password.
     */
    accessSecret?: pulumi.Input<string>;
    createTime?: pulumi.Input<string>;
    /**
     * Database account.
     */
    indepAccount?: pulumi.Input<string>;
    /**
     * Database password.
     */
    indepPassword?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The source information of the security access agent permission is enabled, and the return value is as follows:
     * * **Owner Authorization**: The UID of the owner in parentheses.
     * * **Work Order Authorization**: The ticket number in parentheses is the number of the user to apply for permission.
     */
    originInfo?: pulumi.Input<string>;
    /**
     * Security Protection authorization ID. After the target user is authorized by the security protection agent, the system automatically generates a security protection authorization ID, which is globally unique.
     */
    proxyAccessId?: pulumi.Input<string>;
    /**
     * The ID of the security agent.
     */
    proxyId?: pulumi.Input<string>;
    /**
     * The user ID.
     */
    userId?: pulumi.Input<string>;
    /**
     * User nickname.
     */
    userName?: pulumi.Input<string>;
    /**
     * User UID.
     */
    userUid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EnterpriseProxyAccess resource.
 */
export interface EnterpriseProxyAccessArgs {
    /**
     * Database account.
     */
    indepAccount?: pulumi.Input<string>;
    /**
     * Database password.
     */
    indepPassword?: pulumi.Input<string>;
    /**
     * Security Protection authorization ID. After the target user is authorized by the security protection agent, the system automatically generates a security protection authorization ID, which is globally unique.
     */
    proxyAccessId?: pulumi.Input<string>;
    /**
     * The ID of the security agent.
     */
    proxyId: pulumi.Input<string>;
    /**
     * The user ID.
     */
    userId: pulumi.Input<string>;
}
