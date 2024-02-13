// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provide RDS cluster instance endpoint connection resources, see [What is RDS DB Instance Endpoint](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-createdbinstanceendpoint).
 *
 * > **NOTE:** Available since v1.203.0.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const defaultZones = alicloud.rds.getZones({
 *     engine: "MySQL",
 *     engineVersion: "8.0",
 *     instanceChargeType: "PostPaid",
 *     category: "cluster",
 *     dbInstanceStorageType: "cloud_essd",
 * });
 * const defaultInstanceClasses = defaultZones.then(defaultZones => alicloud.rds.getInstanceClasses({
 *     zoneId: defaultZones.ids?.[0],
 *     engine: "MySQL",
 *     engineVersion: "8.0",
 *     category: "cluster",
 *     dbInstanceStorageType: "cloud_essd",
 *     instanceChargeType: "PostPaid",
 * }));
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.ids?.[0]),
 *     vswitchName: name,
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultInstance = new alicloud.rds.Instance("defaultInstance", {
 *     engine: "MySQL",
 *     engineVersion: "8.0",
 *     instanceType: defaultInstanceClasses.then(defaultInstanceClasses => defaultInstanceClasses.instanceClasses?.[0]?.instanceClass),
 *     instanceStorage: defaultInstanceClasses.then(defaultInstanceClasses => defaultInstanceClasses.instanceClasses?.[0]?.storageRange?.min),
 *     instanceChargeType: "Postpaid",
 *     instanceName: name,
 *     vswitchId: defaultSwitch.id,
 *     monitoringPeriod: 60,
 *     dbInstanceStorageType: "cloud_essd",
 *     securityGroupIds: [defaultSecurityGroup.id],
 *     zoneId: defaultZones.then(defaultZones => defaultZones.ids?.[0]),
 *     zoneIdSlaveA: defaultZones.then(defaultZones => defaultZones.ids?.[0]),
 * });
 * const defaultDbNode = new alicloud.rds.DbNode("defaultDbNode", {
 *     dbInstanceId: defaultInstance.id,
 *     classCode: defaultInstance.instanceType,
 *     zoneId: defaultSwitch.zoneId,
 * });
 * const defaultDbInstanceEndpoint = new alicloud.rds.DbInstanceEndpoint("defaultDbInstanceEndpoint", {
 *     dbInstanceId: defaultDbNode.dbInstanceId,
 *     vpcId: defaultNetwork.id,
 *     vswitchId: defaultInstance.vswitchId,
 *     connectionStringPrefix: "example",
 *     port: "3306",
 *     dbInstanceEndpointDescription: name,
 *     nodeItems: [{
 *         nodeId: defaultDbNode.nodeId,
 *         weight: 25,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * RDS database endpoint feature can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint example <db_instance_id>:<db_instance_endpoint_id>
 * ```
 */
export class DbInstanceEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing DbInstanceEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DbInstanceEndpointState, opts?: pulumi.CustomResourceOptions): DbInstanceEndpoint {
        return new DbInstanceEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rds/dbInstanceEndpoint:DbInstanceEndpoint';

    /**
     * Returns true if the given object is an instance of DbInstanceEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DbInstanceEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DbInstanceEndpoint.__pulumiType;
    }

    /**
     * The internal endpoint.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * The IP address of the internal endpoint.
     */
    public readonly connectionStringPrefix!: pulumi.Output<string>;
    /**
     * The user-defined description of the endpoint.
     */
    public readonly dbInstanceEndpointDescription!: pulumi.Output<string>;
    /**
     * The Endpoint ID of the instance.
     */
    public /*out*/ readonly dbInstanceEndpointId!: pulumi.Output<string>;
    /**
     * The type of the endpoint.
     */
    public /*out*/ readonly dbInstanceEndpointType!: pulumi.Output<string>;
    /**
     * The ID of the instance.
     */
    public readonly dbInstanceId!: pulumi.Output<string>;
    /**
     * The type of the IP address.
     */
    public /*out*/ readonly ipType!: pulumi.Output<string>;
    /**
     * The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight). See `nodeItems` below.
     */
    public readonly nodeItems!: pulumi.Output<outputs.rds.DbInstanceEndpointNodeItem[]>;
    /**
     * The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * The IP address of the internal endpoint.
     */
    public /*out*/ readonly privateIpAddress!: pulumi.Output<string>;
    /**
     * The virtual private cloud (VPC) ID of the internal endpoint.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The vSwitch ID of the internal endpoint.
     */
    public readonly vswitchId!: pulumi.Output<string>;

    /**
     * Create a DbInstanceEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DbInstanceEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DbInstanceEndpointArgs | DbInstanceEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DbInstanceEndpointState | undefined;
            resourceInputs["connectionString"] = state ? state.connectionString : undefined;
            resourceInputs["connectionStringPrefix"] = state ? state.connectionStringPrefix : undefined;
            resourceInputs["dbInstanceEndpointDescription"] = state ? state.dbInstanceEndpointDescription : undefined;
            resourceInputs["dbInstanceEndpointId"] = state ? state.dbInstanceEndpointId : undefined;
            resourceInputs["dbInstanceEndpointType"] = state ? state.dbInstanceEndpointType : undefined;
            resourceInputs["dbInstanceId"] = state ? state.dbInstanceId : undefined;
            resourceInputs["ipType"] = state ? state.ipType : undefined;
            resourceInputs["nodeItems"] = state ? state.nodeItems : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as DbInstanceEndpointArgs | undefined;
            if ((!args || args.connectionStringPrefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionStringPrefix'");
            }
            if ((!args || args.dbInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceId'");
            }
            if ((!args || args.nodeItems === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeItems'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["connectionStringPrefix"] = args ? args.connectionStringPrefix : undefined;
            resourceInputs["dbInstanceEndpointDescription"] = args ? args.dbInstanceEndpointDescription : undefined;
            resourceInputs["dbInstanceId"] = args ? args.dbInstanceId : undefined;
            resourceInputs["nodeItems"] = args ? args.nodeItems : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["connectionString"] = undefined /*out*/;
            resourceInputs["dbInstanceEndpointId"] = undefined /*out*/;
            resourceInputs["dbInstanceEndpointType"] = undefined /*out*/;
            resourceInputs["ipType"] = undefined /*out*/;
            resourceInputs["privateIpAddress"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DbInstanceEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DbInstanceEndpoint resources.
 */
export interface DbInstanceEndpointState {
    /**
     * The internal endpoint.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * The IP address of the internal endpoint.
     */
    connectionStringPrefix?: pulumi.Input<string>;
    /**
     * The user-defined description of the endpoint.
     */
    dbInstanceEndpointDescription?: pulumi.Input<string>;
    /**
     * The Endpoint ID of the instance.
     */
    dbInstanceEndpointId?: pulumi.Input<string>;
    /**
     * The type of the endpoint.
     */
    dbInstanceEndpointType?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    dbInstanceId?: pulumi.Input<string>;
    /**
     * The type of the IP address.
     */
    ipType?: pulumi.Input<string>;
    /**
     * The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight). See `nodeItems` below.
     */
    nodeItems?: pulumi.Input<pulumi.Input<inputs.rds.DbInstanceEndpointNodeItem>[]>;
    /**
     * The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
     */
    port?: pulumi.Input<string>;
    /**
     * The IP address of the internal endpoint.
     */
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The virtual private cloud (VPC) ID of the internal endpoint.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The vSwitch ID of the internal endpoint.
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DbInstanceEndpoint resource.
 */
export interface DbInstanceEndpointArgs {
    /**
     * The IP address of the internal endpoint.
     */
    connectionStringPrefix: pulumi.Input<string>;
    /**
     * The user-defined description of the endpoint.
     */
    dbInstanceEndpointDescription?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    dbInstanceId: pulumi.Input<string>;
    /**
     * The information about the node that is configured for the endpoint.  It contains two sub-fields(node_id and weight). See `nodeItems` below.
     */
    nodeItems: pulumi.Input<pulumi.Input<inputs.rds.DbInstanceEndpointNodeItem>[]>;
    /**
     * The port number of the internal endpoint. You can specify the port number for the internal endpoint.Valid values: 3000 to 5999.
     */
    port: pulumi.Input<string>;
    /**
     * The virtual private cloud (VPC) ID of the internal endpoint.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The vSwitch ID of the internal endpoint.
     */
    vswitchId: pulumi.Input<string>;
}
