// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.polardb.EndpointAddressArgs;
import com.pulumi.alicloud.polardb.inputs.EndpointAddressState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a PolarDB endpoint address resource to allocate an Internet endpoint address string for PolarDB instance.
 * 
 * &gt; **NOTE:** Available since v1.68.0. Each PolarDB instance will allocate a intranet connection string automatically and its prefix is Cluster ID.
 *  To avoid unnecessary conflict, please specified a internet connection prefix before applying the resource.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.polardb.PolardbFunctions;
 * import com.pulumi.alicloud.polardb.inputs.GetNodeClassesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.polardb.Cluster;
 * import com.pulumi.alicloud.polardb.ClusterArgs;
 * import com.pulumi.alicloud.polardb.inputs.GetEndpointsArgs;
 * import com.pulumi.alicloud.polardb.EndpointAddress;
 * import com.pulumi.alicloud.polardb.EndpointAddressArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var default = PolardbFunctions.getNodeClasses(GetNodeClassesArgs.builder()
 *             .dbType("MySQL")
 *             .dbVersion("8.0")
 *             .payType("PostPaid")
 *             .category("Normal")
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .vpcName("terraform-example")
 *             .cidrBlock("172.16.0.0/16")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock("172.16.0.0/24")
 *             .zoneId(default_.classes()[0].zoneId())
 *             .vswitchName("terraform-example")
 *             .build());
 * 
 *         var defaultCluster = new Cluster("defaultCluster", ClusterArgs.builder()
 *             .dbType("MySQL")
 *             .dbVersion("8.0")
 *             .dbNodeClass(default_.classes()[0].supportedEngines()[0].availableResources()[0].dbNodeClass())
 *             .payType("PostPaid")
 *             .vswitchId(defaultSwitch.id())
 *             .description("terraform-example")
 *             .build());
 * 
 *         final var defaultGetEndpoints = PolardbFunctions.getEndpoints(GetEndpointsArgs.builder()
 *             .dbClusterId(defaultCluster.id())
 *             .build());
 * 
 *         var defaultEndpointAddress = new EndpointAddress("defaultEndpointAddress", EndpointAddressArgs.builder()
 *             .dbClusterId(defaultCluster.id())
 *             .dbEndpointId(defaultGetEndpoints.applyValue(getEndpointsResult -> getEndpointsResult).applyValue(defaultGetEndpoints -> defaultGetEndpoints.applyValue(getEndpointsResult -> getEndpointsResult.endpoints()[0].dbEndpointId())))
 *             .connectionPrefix("polardbexample")
 *             .netType("Public")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * PolarDB endpoint address can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:polardb/endpointAddress:EndpointAddress example pc-abc123456:pe-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:polardb/endpointAddress:EndpointAddress")
public class EndpointAddress extends com.pulumi.resources.CustomResource {
    /**
     * Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
     * 
     */
    @Export(name="connectionPrefix", refs={String.class}, tree="[0]")
    private Output<String> connectionPrefix;

    /**
     * @return Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
     * 
     */
    public Output<String> connectionPrefix() {
        return this.connectionPrefix;
    }
    /**
     * Connection cluster or endpoint string.
     * 
     */
    @Export(name="connectionString", refs={String.class}, tree="[0]")
    private Output<String> connectionString;

    /**
     * @return Connection cluster or endpoint string.
     * 
     */
    public Output<String> connectionString() {
        return this.connectionString;
    }
    /**
     * The Id of cluster that can run database.
     * 
     */
    @Export(name="dbClusterId", refs={String.class}, tree="[0]")
    private Output<String> dbClusterId;

    /**
     * @return The Id of cluster that can run database.
     * 
     */
    public Output<String> dbClusterId() {
        return this.dbClusterId;
    }
    /**
     * The Id of endpoint that can run database.
     * 
     */
    @Export(name="dbEndpointId", refs={String.class}, tree="[0]")
    private Output<String> dbEndpointId;

    /**
     * @return The Id of endpoint that can run database.
     * 
     */
    public Output<String> dbEndpointId() {
        return this.dbEndpointId;
    }
    /**
     * The ip address of connection string.
     * 
     */
    @Export(name="ipAddress", refs={String.class}, tree="[0]")
    private Output<String> ipAddress;

    /**
     * @return The ip address of connection string.
     * 
     */
    public Output<String> ipAddress() {
        return this.ipAddress;
    }
    /**
     * Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
     * 
     */
    @Export(name="netType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> netType;

    /**
     * @return Internet connection net type. Valid value: `Public`. Default to `Public`. Currently supported only `Public`.
     * 
     */
    public Output<Optional<String>> netType() {
        return Codegen.optional(this.netType);
    }
    /**
     * Port of the specified endpoint. Valid values: 3000 to 5999.
     * 
     */
    @Export(name="port", refs={String.class}, tree="[0]")
    private Output<String> port;

    /**
     * @return Port of the specified endpoint. Valid values: 3000 to 5999.
     * 
     */
    public Output<String> port() {
        return this.port;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EndpointAddress(java.lang.String name) {
        this(name, EndpointAddressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EndpointAddress(java.lang.String name, EndpointAddressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EndpointAddress(java.lang.String name, EndpointAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/endpointAddress:EndpointAddress", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EndpointAddress(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/endpointAddress:EndpointAddress", name, state, makeResourceOptions(options, id), false);
    }

    private static EndpointAddressArgs makeArgs(EndpointAddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EndpointAddressArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static EndpointAddress get(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointAddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EndpointAddress(name, id, state, options);
    }
}
