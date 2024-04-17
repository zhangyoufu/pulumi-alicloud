// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Hybrid Backup Recovery (HBR) Replication Vault resource.
 *
 * For information about Hybrid Backup Recovery (HBR) Replication Vault and how to use it, see [What is Replication Vault](https://www.alibabacloud.com/help/en/doc-detail/345603.html).
 *
 * > **NOTE:** Available in v1.152.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const sourceRegion = config.get("sourceRegion") || "cn-hangzhou";
 * const default = alicloud.hbr.getReplicationVaultRegions({});
 * const defaultInteger = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const defaultVault = new alicloud.hbr.Vault("default", {vaultName: `terraform-example-${defaultInteger.result}`});
 * const defaultReplicationVault = new alicloud.hbr.ReplicationVault("default", {
 *     replicationSourceRegionId: sourceRegion,
 *     replicationSourceVaultId: defaultVault.id,
 *     vaultName: "terraform-example",
 *     vaultStorageClass: "STANDARD",
 *     description: "terraform-example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Hybrid Backup Recovery (HBR) Replication Vault can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:hbr/replicationVault:ReplicationVault example <id>
 * ```
 */
export class ReplicationVault extends pulumi.CustomResource {
    /**
     * Get an existing ReplicationVault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicationVaultState, opts?: pulumi.CustomResourceOptions): ReplicationVault {
        return new ReplicationVault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:hbr/replicationVault:ReplicationVault';

    /**
     * Returns true if the given object is an instance of ReplicationVault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicationVault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicationVault.__pulumiType;
    }

    /**
     * The description of the backup vault. The description must be 0 to 255 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the region where the source vault resides.
     */
    public readonly replicationSourceRegionId!: pulumi.Output<string>;
    /**
     * The ID of the source vault.
     */
    public readonly replicationSourceVaultId!: pulumi.Output<string>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The name of the backup vault. The name must be 1 to 64 characters in length.
     */
    public readonly vaultName!: pulumi.Output<string>;
    /**
     * The storage type of the backup vault. Valid values: `STANDARD`.
     */
    public readonly vaultStorageClass!: pulumi.Output<string>;

    /**
     * Create a ReplicationVault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicationVaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicationVaultArgs | ReplicationVaultState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReplicationVaultState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["replicationSourceRegionId"] = state ? state.replicationSourceRegionId : undefined;
            resourceInputs["replicationSourceVaultId"] = state ? state.replicationSourceVaultId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vaultName"] = state ? state.vaultName : undefined;
            resourceInputs["vaultStorageClass"] = state ? state.vaultStorageClass : undefined;
        } else {
            const args = argsOrState as ReplicationVaultArgs | undefined;
            if ((!args || args.replicationSourceRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicationSourceRegionId'");
            }
            if ((!args || args.replicationSourceVaultId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicationSourceVaultId'");
            }
            if ((!args || args.vaultName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vaultName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["replicationSourceRegionId"] = args ? args.replicationSourceRegionId : undefined;
            resourceInputs["replicationSourceVaultId"] = args ? args.replicationSourceVaultId : undefined;
            resourceInputs["vaultName"] = args ? args.vaultName : undefined;
            resourceInputs["vaultStorageClass"] = args ? args.vaultStorageClass : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReplicationVault.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicationVault resources.
 */
export interface ReplicationVaultState {
    /**
     * The description of the backup vault. The description must be 0 to 255 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the region where the source vault resides.
     */
    replicationSourceRegionId?: pulumi.Input<string>;
    /**
     * The ID of the source vault.
     */
    replicationSourceVaultId?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of the backup vault. The name must be 1 to 64 characters in length.
     */
    vaultName?: pulumi.Input<string>;
    /**
     * The storage type of the backup vault. Valid values: `STANDARD`.
     */
    vaultStorageClass?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplicationVault resource.
 */
export interface ReplicationVaultArgs {
    /**
     * The description of the backup vault. The description must be 0 to 255 characters in length.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the region where the source vault resides.
     */
    replicationSourceRegionId: pulumi.Input<string>;
    /**
     * The ID of the source vault.
     */
    replicationSourceVaultId: pulumi.Input<string>;
    /**
     * The name of the backup vault. The name must be 1 to 64 characters in length.
     */
    vaultName: pulumi.Input<string>;
    /**
     * The storage type of the backup vault. Valid values: `STANDARD`.
     */
    vaultStorageClass?: pulumi.Input<string>;
}
