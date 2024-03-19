// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mse;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mse.EngineNamespaceArgs;
import com.pulumi.alicloud.mse.inputs.EngineNamespaceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Microservice Engine (MSE) Engine Namespace resource.
 * 
 * For information about Microservice Engine (MSE) Engine Namespace and how to use it, see [What is Engine Namespace](https://www.alibabacloud.com/help/en/mse/developer-reference/api-mse-2019-05-31-createenginenamespace).
 * 
 * &gt; **NOTE:** Available in v1.166.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.mse.Cluster;
 * import com.pulumi.alicloud.mse.ClusterArgs;
 * import com.pulumi.alicloud.mse.EngineNamespace;
 * import com.pulumi.alicloud.mse.EngineNamespaceArgs;
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
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         final var exampleZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var exampleNetwork = new Network(&#34;exampleNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .build());
 * 
 *         var exampleSwitch = new Switch(&#34;exampleSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .vpcId(exampleNetwork.id())
 *             .zoneId(exampleZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var default_ = new Cluster(&#34;default&#34;, ClusterArgs.builder()        
 *             .connectionType(&#34;slb&#34;)
 *             .netType(&#34;privatenet&#34;)
 *             .vswitchId(exampleSwitch.id())
 *             .clusterSpecification(&#34;MSE_SC_1_2_60_c&#34;)
 *             .clusterVersion(&#34;NACOS_2_0_0&#34;)
 *             .instanceCount(&#34;1&#34;)
 *             .pubNetworkFlow(&#34;1&#34;)
 *             .clusterAliasName(name)
 *             .mseVersion(&#34;mse_dev&#34;)
 *             .clusterType(&#34;Nacos-Ans&#34;)
 *             .build());
 * 
 *         var exampleEngineNamespace = new EngineNamespace(&#34;exampleEngineNamespace&#34;, EngineNamespaceArgs.builder()        
 *             .clusterId(default_.id())
 *             .namespaceShowName(name)
 *             .namespaceId(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Microservice Engine (MSE) Engine Namespace can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:mse/engineNamespace:EngineNamespace example &lt;cluster_id&gt;:&lt;namespace_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:mse/engineNamespace:EngineNamespace")
public class EngineNamespace extends com.pulumi.resources.CustomResource {
    /**
     * The language type of the returned information. Valid values: `zh`, `en`.
     * 
     */
    @Export(name="acceptLanguage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> acceptLanguage;

    /**
     * @return The language type of the returned information. Valid values: `zh`, `en`.
     * 
     */
    public Output<Optional<String>> acceptLanguage() {
        return Codegen.optional(this.acceptLanguage);
    }
    /**
     * The id of the cluster.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return The id of the cluster.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The id of Namespace.
     * 
     */
    @Export(name="namespaceId", refs={String.class}, tree="[0]")
    private Output<String> namespaceId;

    /**
     * @return The id of Namespace.
     * 
     */
    public Output<String> namespaceId() {
        return this.namespaceId;
    }
    /**
     * The name of the Engine Namespace.
     * 
     */
    @Export(name="namespaceShowName", refs={String.class}, tree="[0]")
    private Output<String> namespaceShowName;

    /**
     * @return The name of the Engine Namespace.
     * 
     */
    public Output<String> namespaceShowName() {
        return this.namespaceShowName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EngineNamespace(String name) {
        this(name, EngineNamespaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EngineNamespace(String name, EngineNamespaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EngineNamespace(String name, EngineNamespaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mse/engineNamespace:EngineNamespace", name, args == null ? EngineNamespaceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EngineNamespace(String name, Output<String> id, @Nullable EngineNamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mse/engineNamespace:EngineNamespace", name, state, makeResourceOptions(options, id));
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
    public static EngineNamespace get(String name, Output<String> id, @Nullable EngineNamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EngineNamespace(name, id, state, options);
    }
}
