// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.polardb.GlobalDatabaseNetworkArgs;
import com.pulumi.alicloud.polardb.inputs.GlobalDatabaseNetworkState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a PolarDB Global Database Network resource.
 * 
 * For information about PolarDB Global Database Network and how to use it, see [What is Global Database Network](https://www.alibabacloud.com/help/en/polardb/api-polardb-2017-08-01-createglobaldatabasenetwork).
 * 
 * &gt; **NOTE:** Available since v1.181.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
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
 * import com.pulumi.alicloud.polardb.GlobalDatabaseNetwork;
 * import com.pulumi.alicloud.polardb.GlobalDatabaseNetworkArgs;
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
 *         final var defaultNodeClasses = PolardbFunctions.getNodeClasses(GetNodeClassesArgs.builder()
 *             .dbType(&#34;MySQL&#34;)
 *             .dbVersion(&#34;8.0&#34;)
 *             .category(&#34;Normal&#34;)
 *             .payType(&#34;PostPaid&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultNodeClasses.applyValue(getNodeClassesResult -&gt; getNodeClassesResult.classes()[0].zoneId()))
 *             .vswitchName(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         var defaultCluster = new Cluster(&#34;defaultCluster&#34;, ClusterArgs.builder()        
 *             .dbType(&#34;MySQL&#34;)
 *             .dbVersion(&#34;8.0&#34;)
 *             .dbNodeClass(defaultNodeClasses.applyValue(getNodeClassesResult -&gt; getNodeClassesResult.classes()[0].supportedEngines()[0].availableResources()[0].dbNodeClass()))
 *             .payType(&#34;PostPaid&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .description(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         var defaultGlobalDatabaseNetwork = new GlobalDatabaseNetwork(&#34;defaultGlobalDatabaseNetwork&#34;, GlobalDatabaseNetworkArgs.builder()        
 *             .dbClusterId(defaultCluster.id())
 *             .description(&#34;terraform-example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * PolarDB Global Database Network can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:polardb/globalDatabaseNetwork:GlobalDatabaseNetwork example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:polardb/globalDatabaseNetwork:GlobalDatabaseNetwork")
public class GlobalDatabaseNetwork extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the primary cluster.
     * 
     */
    @Export(name="dbClusterId", refs={String.class}, tree="[0]")
    private Output<String> dbClusterId;

    /**
     * @return The ID of the primary cluster.
     * 
     */
    public Output<String> dbClusterId() {
        return this.dbClusterId;
    }
    /**
     * The description of the Global Database Network.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the Global Database Network.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The status of the Global Database Network.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Global Database Network.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GlobalDatabaseNetwork(String name) {
        this(name, GlobalDatabaseNetworkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GlobalDatabaseNetwork(String name, GlobalDatabaseNetworkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GlobalDatabaseNetwork(String name, GlobalDatabaseNetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/globalDatabaseNetwork:GlobalDatabaseNetwork", name, args == null ? GlobalDatabaseNetworkArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GlobalDatabaseNetwork(String name, Output<String> id, @Nullable GlobalDatabaseNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/globalDatabaseNetwork:GlobalDatabaseNetwork", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static GlobalDatabaseNetwork get(String name, Output<String> id, @Nullable GlobalDatabaseNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GlobalDatabaseNetwork(name, id, state, options);
    }
}
