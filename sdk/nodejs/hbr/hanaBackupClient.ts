// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Hybrid Backup Recovery (HBR) Hana Backup Client resource.
 *
 * For information about Hybrid Backup Recovery (HBR) Hana Backup Client and how to use it, see [What is Hana Backup Client](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/api-doc-hbr-2017-09-08-api-doc-createclients).
 *
 * > **NOTE:** Available in v1.198.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const exampleZones = alicloud.getZones({
 *     availableResourceCreation: "Instance",
 * });
 * const exampleInstanceTypes = exampleZones.then(exampleZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: exampleZones.zones?.[0]?.id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }));
 * const exampleImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_[0-9]+_[0-9]+_x64*",
 *     owners: "system",
 * });
 * const exampleNetwork = new alicloud.vpc.Network("exampleNetwork", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 * });
 * const exampleSwitch = new alicloud.vpc.Switch("exampleSwitch", {
 *     vswitchName: "terraform-example",
 *     cidrBlock: "172.17.3.0/24",
 *     vpcId: exampleNetwork.id,
 *     zoneId: exampleZones.then(exampleZones => exampleZones.zones?.[0]?.id),
 * });
 * const exampleSecurityGroup = new alicloud.ecs.SecurityGroup("exampleSecurityGroup", {vpcId: exampleNetwork.id});
 * const exampleInstance = new alicloud.ecs.Instance("exampleInstance", {
 *     imageId: exampleImages.then(exampleImages => exampleImages.images?.[0]?.id),
 *     instanceType: exampleInstanceTypes.then(exampleInstanceTypes => exampleInstanceTypes.instanceTypes?.[0]?.id),
 *     availabilityZone: exampleZones.then(exampleZones => exampleZones.zones?.[0]?.id),
 *     securityGroups: [exampleSecurityGroup.id],
 *     instanceName: "terraform-example",
 *     internetChargeType: "PayByBandwidth",
 *     vswitchId: exampleSwitch.id,
 * });
 * const exampleResourceGroups = alicloud.resourcemanager.getResourceGroups({
 *     status: "OK",
 * });
 * const exampleVault = new alicloud.hbr.Vault("exampleVault", {vaultName: "terraform-example"});
 * const exampleHanaInstance = new alicloud.hbr.HanaInstance("exampleHanaInstance", {
 *     alertSetting: "INHERITED",
 *     hanaName: "terraform-example",
 *     host: "1.1.1.1",
 *     instanceNumber: 1,
 *     password: "YouPassword123",
 *     resourceGroupId: exampleResourceGroups.then(exampleResourceGroups => exampleResourceGroups.groups?.[0]?.id),
 *     sid: "HXE",
 *     useSsl: false,
 *     userName: "admin",
 *     validateCertificate: false,
 *     vaultId: exampleVault.id,
 * });
 * const _default = new alicloud.hbr.HanaBackupClient("default", {
 *     vaultId: exampleVault.id,
 *     clientInfo: pulumi.interpolate`[ { "instanceId": "${exampleInstance.id}", "clusterId": "${exampleHanaInstance.hanaInstanceId}", "sourceTypes": [ "HANA" ]  }]`,
 *     alertSetting: "INHERITED",
 *     useHttps: true,
 * });
 * ```
 *
 * ## Import
 *
 * Hybrid Backup Recovery (HBR) Hana Backup Client can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:hbr/hanaBackupClient:HanaBackupClient example <vault_id>:<client_id>
 * ```
 */
export class HanaBackupClient extends pulumi.CustomResource {
    /**
     * Get an existing HanaBackupClient resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HanaBackupClientState, opts?: pulumi.CustomResourceOptions): HanaBackupClient {
        return new HanaBackupClient(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:hbr/hanaBackupClient:HanaBackupClient';

    /**
     * Returns true if the given object is an instance of HanaBackupClient.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HanaBackupClient {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HanaBackupClient.__pulumiType;
    }

    /**
     * The alert settings. Valid value: `INHERITED`.
     */
    public readonly alertSetting!: pulumi.Output<string>;
    /**
     * The ID of the backup client.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * The installation information of the HBR clients.
     */
    public readonly clientInfo!: pulumi.Output<string | undefined>;
    /**
     * The ID of the SAP HANA instance.
     */
    public /*out*/ readonly clusterId!: pulumi.Output<string>;
    /**
     * The ID of the instance.
     */
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    /**
     * The status of the Hana Backup Client.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
     */
    public readonly useHttps!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the backup vault.
     */
    public readonly vaultId!: pulumi.Output<string>;

    /**
     * Create a HanaBackupClient resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HanaBackupClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HanaBackupClientArgs | HanaBackupClientState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HanaBackupClientState | undefined;
            resourceInputs["alertSetting"] = state ? state.alertSetting : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientInfo"] = state ? state.clientInfo : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["useHttps"] = state ? state.useHttps : undefined;
            resourceInputs["vaultId"] = state ? state.vaultId : undefined;
        } else {
            const args = argsOrState as HanaBackupClientArgs | undefined;
            if ((!args || args.vaultId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vaultId'");
            }
            resourceInputs["alertSetting"] = args ? args.alertSetting : undefined;
            resourceInputs["clientInfo"] = args ? args.clientInfo : undefined;
            resourceInputs["useHttps"] = args ? args.useHttps : undefined;
            resourceInputs["vaultId"] = args ? args.vaultId : undefined;
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["clusterId"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HanaBackupClient.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HanaBackupClient resources.
 */
export interface HanaBackupClientState {
    /**
     * The alert settings. Valid value: `INHERITED`.
     */
    alertSetting?: pulumi.Input<string>;
    /**
     * The ID of the backup client.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The installation information of the HBR clients.
     */
    clientInfo?: pulumi.Input<string>;
    /**
     * The ID of the SAP HANA instance.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The status of the Hana Backup Client.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
     */
    useHttps?: pulumi.Input<boolean>;
    /**
     * The ID of the backup vault.
     */
    vaultId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HanaBackupClient resource.
 */
export interface HanaBackupClientArgs {
    /**
     * The alert settings. Valid value: `INHERITED`.
     */
    alertSetting?: pulumi.Input<string>;
    /**
     * The installation information of the HBR clients.
     */
    clientInfo?: pulumi.Input<string>;
    /**
     * Specifies whether to transmit data over HTTPS. Valid values: `true`, `false`.
     */
    useHttps?: pulumi.Input<boolean>;
    /**
     * The ID of the backup vault.
     */
    vaultId: pulumi.Input<string>;
}
