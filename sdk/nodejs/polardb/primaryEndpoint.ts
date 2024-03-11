// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a PolarDB endpoint resource to manage primary endpoint of PolarDB cluster.
 *
 * > **NOTE:** Available since v1.217.0
 *
 * > **NOTE:** The default primary endpoint can not be created or deleted manually.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultNodeClasses = alicloud.polardb.getNodeClasses({
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     payType: "PostPaid",
 *     category: "Normal",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: "terraform-example",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultNodeClasses.then(defaultNodeClasses => defaultNodeClasses.classes?.[0]?.zoneId),
 *     vswitchName: "terraform-example",
 * });
 * const defaultCluster = new alicloud.polardb.Cluster("defaultCluster", {
 *     dbType: "MySQL",
 *     dbVersion: "8.0",
 *     dbNodeClass: defaultNodeClasses.then(defaultNodeClasses => defaultNodeClasses.classes?.[0]?.supportedEngines?.[0]?.availableResources?.[0]?.dbNodeClass),
 *     payType: "PostPaid",
 *     vswitchId: defaultSwitch.id,
 *     description: "terraform-example",
 * });
 * const defaultPrimaryEndpoint = new alicloud.polardb.PrimaryEndpoint("defaultPrimaryEndpoint", {dbClusterId: defaultCluster.id});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * PolarDB endpoint can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:polardb/primaryEndpoint:PrimaryEndpoint example pc-abc123456:pe-abc123456
 * ```
 */
export class PrimaryEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing PrimaryEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrimaryEndpointState, opts?: pulumi.CustomResourceOptions): PrimaryEndpoint {
        return new PrimaryEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:polardb/primaryEndpoint:PrimaryEndpoint';

    /**
     * Returns true if the given object is an instance of PrimaryEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrimaryEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrimaryEndpoint.__pulumiType;
    }

    /**
     * Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
     */
    public readonly connectionPrefix!: pulumi.Output<string>;
    /**
     * The Id of cluster that can run database.
     */
    public readonly dbClusterId!: pulumi.Output<string>;
    /**
     * The name of the endpoint.
     */
    public readonly dbEndpointDescription!: pulumi.Output<string | undefined>;
    /**
     * The ID of the cluster endpoint.
     */
    public /*out*/ readonly dbEndpointId!: pulumi.Output<string>;
    /**
     * Type of endpoint.
     */
    public /*out*/ readonly endpointType!: pulumi.Output<string>;
    /**
     * The network type of the endpoint address.
     */
    public readonly netType!: pulumi.Output<string | undefined>;
    /**
     * Port of the specified endpoint. Valid values: 3000 to 5999.
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
     * **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
     * For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
     */
    public readonly sslAutoRotate!: pulumi.Output<string | undefined>;
    /**
     * The specifies SSL certificate download link.
     */
    public /*out*/ readonly sslCertificateUrl!: pulumi.Output<string>;
    /**
     * The SSL connection string.
     */
    public /*out*/ readonly sslConnectionString!: pulumi.Output<string>;
    /**
     * Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
     */
    public readonly sslEnabled!: pulumi.Output<string | undefined>;
    /**
     * The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    public /*out*/ readonly sslExpireTime!: pulumi.Output<string>;

    /**
     * Create a PrimaryEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrimaryEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrimaryEndpointArgs | PrimaryEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrimaryEndpointState | undefined;
            resourceInputs["connectionPrefix"] = state ? state.connectionPrefix : undefined;
            resourceInputs["dbClusterId"] = state ? state.dbClusterId : undefined;
            resourceInputs["dbEndpointDescription"] = state ? state.dbEndpointDescription : undefined;
            resourceInputs["dbEndpointId"] = state ? state.dbEndpointId : undefined;
            resourceInputs["endpointType"] = state ? state.endpointType : undefined;
            resourceInputs["netType"] = state ? state.netType : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["sslAutoRotate"] = state ? state.sslAutoRotate : undefined;
            resourceInputs["sslCertificateUrl"] = state ? state.sslCertificateUrl : undefined;
            resourceInputs["sslConnectionString"] = state ? state.sslConnectionString : undefined;
            resourceInputs["sslEnabled"] = state ? state.sslEnabled : undefined;
            resourceInputs["sslExpireTime"] = state ? state.sslExpireTime : undefined;
        } else {
            const args = argsOrState as PrimaryEndpointArgs | undefined;
            if ((!args || args.dbClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbClusterId'");
            }
            resourceInputs["connectionPrefix"] = args ? args.connectionPrefix : undefined;
            resourceInputs["dbClusterId"] = args ? args.dbClusterId : undefined;
            resourceInputs["dbEndpointDescription"] = args ? args.dbEndpointDescription : undefined;
            resourceInputs["netType"] = args ? args.netType : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["sslAutoRotate"] = args ? args.sslAutoRotate : undefined;
            resourceInputs["sslEnabled"] = args ? args.sslEnabled : undefined;
            resourceInputs["dbEndpointId"] = undefined /*out*/;
            resourceInputs["endpointType"] = undefined /*out*/;
            resourceInputs["sslCertificateUrl"] = undefined /*out*/;
            resourceInputs["sslConnectionString"] = undefined /*out*/;
            resourceInputs["sslExpireTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrimaryEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrimaryEndpoint resources.
 */
export interface PrimaryEndpointState {
    /**
     * Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
     */
    connectionPrefix?: pulumi.Input<string>;
    /**
     * The Id of cluster that can run database.
     */
    dbClusterId?: pulumi.Input<string>;
    /**
     * The name of the endpoint.
     */
    dbEndpointDescription?: pulumi.Input<string>;
    /**
     * The ID of the cluster endpoint.
     */
    dbEndpointId?: pulumi.Input<string>;
    /**
     * Type of endpoint.
     */
    endpointType?: pulumi.Input<string>;
    /**
     * The network type of the endpoint address.
     */
    netType?: pulumi.Input<string>;
    /**
     * Port of the specified endpoint. Valid values: 3000 to 5999.
     */
    port?: pulumi.Input<string>;
    /**
     * Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
     * **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
     * For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
     */
    sslAutoRotate?: pulumi.Input<string>;
    /**
     * The specifies SSL certificate download link.
     */
    sslCertificateUrl?: pulumi.Input<string>;
    /**
     * The SSL connection string.
     */
    sslConnectionString?: pulumi.Input<string>;
    /**
     * Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
     */
    sslEnabled?: pulumi.Input<string>;
    /**
     * The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     */
    sslExpireTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrimaryEndpoint resource.
 */
export interface PrimaryEndpointArgs {
    /**
     * Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
     */
    connectionPrefix?: pulumi.Input<string>;
    /**
     * The Id of cluster that can run database.
     */
    dbClusterId: pulumi.Input<string>;
    /**
     * The name of the endpoint.
     */
    dbEndpointDescription?: pulumi.Input<string>;
    /**
     * The network type of the endpoint address.
     */
    netType?: pulumi.Input<string>;
    /**
     * Port of the specified endpoint. Valid values: 3000 to 5999.
     */
    port?: pulumi.Input<string>;
    /**
     * Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
     * **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
     * For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
     */
    sslAutoRotate?: pulumi.Input<string>;
    /**
     * Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
     */
    sslEnabled?: pulumi.Input<string>;
}
