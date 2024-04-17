// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ApsaraDB for MyBase Dedicated Host Account resource.
 *
 * For information about ApsaraDB for MyBase Dedicated Host Account and how to use it, see [What is Dedicated Host Account](https://www.alibabacloud.com/help/en/apsaradb-for-mybase/latest/creatededicatedhostaccount).
 *
 * > **NOTE:** Available since v1.148.0.
 *
 * > **NOTE:** Each Dedicated host can have only one account. Before you create an account for a host, make sure that the existing account is deleted.
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
 * const name = config.get("name") || "tf_example";
 * const default = alicloud.cddc.getZones({});
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: _default.then(_default => _default.ids?.[0]),
 * });
 * const defaultDedicatedHostGroup = new alicloud.cddc.DedicatedHostGroup("default", {
 *     engine: "MySQL",
 *     vpcId: defaultNetwork.id,
 *     cpuAllocationRatio: 101,
 *     memAllocationRatio: 50,
 *     diskAllocationRatio: 200,
 *     allocationPolicy: "Evenly",
 *     hostReplacePolicy: "Manual",
 *     dedicatedHostGroupDesc: name,
 *     openPermission: true,
 * });
 * const defaultGetHostEcsLevelInfos = _default.then(_default => alicloud.cddc.getHostEcsLevelInfos({
 *     dbType: "mysql",
 *     zoneId: _default.ids?.[0],
 *     storageType: "cloud_essd",
 * }));
 * const defaultDedicatedHost = new alicloud.cddc.DedicatedHost("default", {
 *     hostName: name,
 *     dedicatedHostGroupId: defaultDedicatedHostGroup.id,
 *     hostClass: defaultGetHostEcsLevelInfos.then(defaultGetHostEcsLevelInfos => defaultGetHostEcsLevelInfos.infos?.[0]?.resClassCode),
 *     zoneId: _default.then(_default => _default.ids?.[0]),
 *     vswitchId: defaultSwitch.id,
 *     paymentType: "Subscription",
 *     tags: {
 *         Created: "TF",
 *         For: "CDDC_DEDICATED",
 *     },
 * });
 * const defaultDedicatedHostAccount = new alicloud.cddc.DedicatedHostAccount("default", {
 *     accountName: name,
 *     accountPassword: "Password1234",
 *     dedicatedHostId: defaultDedicatedHost.dedicatedHostId,
 *     accountType: "Normal",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ApsaraDB for MyBase Dedicated Host Account can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount example <dedicated_host_id>:<account_name>
 * ```
 */
export class DedicatedHostAccount extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedHostAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DedicatedHostAccountState, opts?: pulumi.CustomResourceOptions): DedicatedHostAccount {
        return new DedicatedHostAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount';

    /**
     * Returns true if the given object is an instance of DedicatedHostAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedHostAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedHostAccount.__pulumiType;
    }

    /**
     * The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&*()_+-=`.
     */
    public readonly accountPassword!: pulumi.Output<string>;
    /**
     * The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
     */
    public readonly accountType!: pulumi.Output<string | undefined>;
    /**
     * The ID of Dedicated the host.
     */
    public readonly dedicatedHostId!: pulumi.Output<string>;

    /**
     * Create a DedicatedHostAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedHostAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DedicatedHostAccountArgs | DedicatedHostAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DedicatedHostAccountState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["accountType"] = state ? state.accountType : undefined;
            resourceInputs["dedicatedHostId"] = state ? state.dedicatedHostId : undefined;
        } else {
            const args = argsOrState as DedicatedHostAccountArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.accountPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountPassword'");
            }
            if ((!args || args.dedicatedHostId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dedicatedHostId'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args?.accountPassword ? pulumi.secret(args.accountPassword) : undefined;
            resourceInputs["accountType"] = args ? args.accountType : undefined;
            resourceInputs["dedicatedHostId"] = args ? args.dedicatedHostId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accountPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DedicatedHostAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DedicatedHostAccount resources.
 */
export interface DedicatedHostAccountState {
    /**
     * The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&*()_+-=`.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
     */
    accountType?: pulumi.Input<string>;
    /**
     * The ID of Dedicated the host.
     */
    dedicatedHostId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DedicatedHostAccount resource.
 */
export interface DedicatedHostAccountArgs {
    /**
     * The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
     */
    accountName: pulumi.Input<string>;
    /**
     * The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&*()_+-=`.
     */
    accountPassword: pulumi.Input<string>;
    /**
     * The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
     */
    accountType?: pulumi.Input<string>;
    /**
     * The ID of Dedicated the host.
     */
    dedicatedHostId: pulumi.Input<string>;
}
