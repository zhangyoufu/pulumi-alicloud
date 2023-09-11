// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Container Registry Enterprise Edition instance can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cr/registryEnterpriseInstance:RegistryEnterpriseInstance default cri-test
 * ```
 */
export class RegistryEnterpriseInstance extends pulumi.CustomResource {
    /**
     * Get an existing RegistryEnterpriseInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryEnterpriseInstanceState, opts?: pulumi.CustomResourceOptions): RegistryEnterpriseInstance {
        return new RegistryEnterpriseInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cr/registryEnterpriseInstance:RegistryEnterpriseInstance';

    /**
     * Returns true if the given object is an instance of RegistryEnterpriseInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistryEnterpriseInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistryEnterpriseInstance.__pulumiType;
    }

    /**
     * Time of Container Registry Enterprise Edition instance creation.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * Name of your customized oss bucket. Use this bucket as instance storage if set.
     */
    public readonly customOssBucket!: pulumi.Output<string | undefined>;
    /**
     * Time of Container Registry Enterprise Edition instance expiration.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * Name of Container Registry Enterprise Edition instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
     */
    public readonly kmsEncryptedPassword!: pulumi.Output<string | undefined>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    public readonly kmsEncryptionContext!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
     */
    public readonly paymentType!: pulumi.Output<string | undefined>;
    /**
     * Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
     */
    public readonly renewPeriod!: pulumi.Output<number | undefined>;
    /**
     * Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
     */
    public readonly renewalStatus!: pulumi.Output<string | undefined>;
    /**
     * Status of Container Registry Enterprise Edition instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a RegistryEnterpriseInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryEnterpriseInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryEnterpriseInstanceArgs | RegistryEnterpriseInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegistryEnterpriseInstanceState | undefined;
            resourceInputs["createdTime"] = state ? state.createdTime : undefined;
            resourceInputs["customOssBucket"] = state ? state.customOssBucket : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["renewPeriod"] = state ? state.renewPeriod : undefined;
            resourceInputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as RegistryEnterpriseInstanceArgs | undefined;
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            resourceInputs["customOssBucket"] = args ? args.customOssBucket : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["renewPeriod"] = args ? args.renewPeriod : undefined;
            resourceInputs["renewalStatus"] = args ? args.renewalStatus : undefined;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(RegistryEnterpriseInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegistryEnterpriseInstance resources.
 */
export interface RegistryEnterpriseInstanceState {
    /**
     * Time of Container Registry Enterprise Edition instance creation.
     */
    createdTime?: pulumi.Input<string>;
    /**
     * Name of your customized oss bucket. Use this bucket as instance storage if set.
     */
    customOssBucket?: pulumi.Input<string>;
    /**
     * Time of Container Registry Enterprise Edition instance expiration.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
     */
    password?: pulumi.Input<string>;
    /**
     * Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
     */
    period?: pulumi.Input<number>;
    /**
     * Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * Status of Container Registry Enterprise Edition instance.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RegistryEnterpriseInstance resource.
 */
export interface RegistryEnterpriseInstanceArgs {
    /**
     * Name of your customized oss bucket. Use this bucket as instance storage if set.
     */
    customOssBucket?: pulumi.Input<string>;
    /**
     * Name of Container Registry Enterprise Edition instance.
     */
    instanceName: pulumi.Input<string>;
    /**
     * Type of Container Registry Enterprise Edition instance. Valid values: `Basic`, `Standard`, `Advanced`. **NOTE:** International Account doesn't supports `Standard`.
     */
    instanceType: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to an instance. If the `password` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The password of the Instance. The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.
     */
    password?: pulumi.Input<string>;
    /**
     * Subscription of Container Registry Enterprise Edition instance. Default value: `Subscription`. Valid values: `Subscription`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * Service time of Container Registry Enterprise Edition instance. Default value: `12`. Valid values: `1`, `2`, `3`, `6`, `12`, `24`, `36`, `48`, `60`. Unit: `month`.
     */
    period?: pulumi.Input<number>;
    /**
     * Renewal period of Container Registry Enterprise Edition instance. Unit: `month`.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * Renewal status of Container Registry Enterprise Edition instance. Valid values: `AutoRenewal`, `ManualRenewal`.
     */
    renewalStatus?: pulumi.Input<string>;
}
