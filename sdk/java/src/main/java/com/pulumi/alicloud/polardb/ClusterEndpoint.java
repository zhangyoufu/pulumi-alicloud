// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.polardb.ClusterEndpointArgs;
import com.pulumi.alicloud.polardb.inputs.ClusterEndpointState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a PolarDB endpoint resource to manage cluster endpoint of PolarDB cluster.
 * 
 * &gt; **NOTE:** Available since v1.217.0
 * 
 * &gt; **NOTE:** The default cluster endpoint can not be created or deleted manually.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
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
 * import com.pulumi.alicloud.polardb.ClusterEndpoint;
 * import com.pulumi.alicloud.polardb.ClusterEndpointArgs;
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
 *             .payType(&#34;PostPaid&#34;)
 *             .category(&#34;Normal&#34;)
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
 *         var defaultClusterEndpoint = new ClusterEndpoint(&#34;defaultClusterEndpoint&#34;, ClusterEndpointArgs.builder()        
 *             .dbClusterId(defaultCluster.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * PolarDB endpoint can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:polardb/clusterEndpoint:ClusterEndpoint example pc-abc123456:pe-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:polardb/clusterEndpoint:ClusterEndpoint")
public class ClusterEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
     * 
     */
    @Export(name="autoAddNewNodes", refs={String.class}, tree="[0]")
    private Output<String> autoAddNewNodes;

    /**
     * @return Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
     * 
     */
    public Output<String> autoAddNewNodes() {
        return this.autoAddNewNodes;
    }
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
     * The name of the endpoint.
     * 
     */
    @Export(name="dbEndpointDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dbEndpointDescription;

    /**
     * @return The name of the endpoint.
     * 
     */
    public Output<Optional<String>> dbEndpointDescription() {
        return Codegen.optional(this.dbEndpointDescription);
    }
    /**
     * The ID of the cluster endpoint.
     * 
     */
    @Export(name="dbEndpointId", refs={String.class}, tree="[0]")
    private Output<String> dbEndpointId;

    /**
     * @return The ID of the cluster endpoint.
     * 
     */
    public Output<String> dbEndpointId() {
        return this.dbEndpointId;
    }
    /**
     * The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
     * 
     */
    @Export(name="endpointConfig", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> endpointConfig;

    /**
     * @return The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
     * 
     */
    public Output<Map<String,Object>> endpointConfig() {
        return this.endpointConfig;
    }
    /**
     * Type of endpoint.
     * 
     */
    @Export(name="endpointType", refs={String.class}, tree="[0]")
    private Output<String> endpointType;

    /**
     * @return Type of endpoint.
     * 
     */
    public Output<String> endpointType() {
        return this.endpointType;
    }
    /**
     * The network type of the endpoint address.
     * 
     */
    @Export(name="netType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> netType;

    /**
     * @return The network type of the endpoint address.
     * 
     */
    public Output<Optional<String>> netType() {
        return Codegen.optional(this.netType);
    }
    /**
     * Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
     * 
     */
    @Export(name="nodes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> nodes;

    /**
     * @return Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
     * 
     */
    public Output<List<String>> nodes() {
        return this.nodes;
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
     * Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
     * 
     */
    @Export(name="readWriteMode", refs={String.class}, tree="[0]")
    private Output<String> readWriteMode;

    /**
     * @return Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
     * 
     */
    public Output<String> readWriteMode() {
        return this.readWriteMode;
    }
    /**
     * Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
     * **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
     * For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
     * 
     */
    @Export(name="sslAutoRotate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sslAutoRotate;

    /**
     * @return Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
     * **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
     * For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
     * 
     */
    public Output<Optional<String>> sslAutoRotate() {
        return Codegen.optional(this.sslAutoRotate);
    }
    /**
     * The specifies SSL certificate download link.
     * 
     */
    @Export(name="sslCertificateUrl", refs={String.class}, tree="[0]")
    private Output<String> sslCertificateUrl;

    /**
     * @return The specifies SSL certificate download link.
     * 
     */
    public Output<String> sslCertificateUrl() {
        return this.sslCertificateUrl;
    }
    /**
     * The SSL connection string.
     * 
     */
    @Export(name="sslConnectionString", refs={String.class}, tree="[0]")
    private Output<String> sslConnectionString;

    /**
     * @return The SSL connection string.
     * 
     */
    public Output<String> sslConnectionString() {
        return this.sslConnectionString;
    }
    /**
     * Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
     * 
     */
    @Export(name="sslEnabled", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sslEnabled;

    /**
     * @return Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
     * 
     */
    public Output<Optional<String>> sslEnabled() {
        return Codegen.optional(this.sslEnabled);
    }
    /**
     * The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    @Export(name="sslExpireTime", refs={String.class}, tree="[0]")
    private Output<String> sslExpireTime;

    /**
     * @return The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    public Output<String> sslExpireTime() {
        return this.sslExpireTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClusterEndpoint(String name) {
        this(name, ClusterEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClusterEndpoint(String name, ClusterEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClusterEndpoint(String name, ClusterEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/clusterEndpoint:ClusterEndpoint", name, args == null ? ClusterEndpointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClusterEndpoint(String name, Output<String> id, @Nullable ClusterEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/clusterEndpoint:ClusterEndpoint", name, state, makeResourceOptions(options, id));
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
    public static ClusterEndpoint get(String name, Output<String> id, @Nullable ClusterEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClusterEndpoint(name, id, state, options);
    }
}
