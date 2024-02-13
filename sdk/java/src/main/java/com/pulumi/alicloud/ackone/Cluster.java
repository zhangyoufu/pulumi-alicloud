// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ackone;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ackone.ClusterArgs;
import com.pulumi.alicloud.ackone.inputs.ClusterState;
import com.pulumi.alicloud.ackone.outputs.ClusterNetwork;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Ack One Cluster resource. Fleet Manager Cluster.
 * 
 * For information about Ack One Cluster and how to use it, see [What is Cluster](https://www.alibabacloud.com/help/en/ack/distributed-cloud-container-platform-for-kubernetes/developer-reference/api-adcp-2022-01-01-createhubcluster).
 * 
 * &gt; **NOTE:** Available since v1.212.0.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ackone.Cluster;
 * import com.pulumi.alicloud.ackone.ClusterArgs;
 * import com.pulumi.alicloud.ackone.inputs.ClusterNetworkArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var defaultVpc = new Network(&#34;defaultVpc&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .vpcName(name)
 *             .build());
 * 
 *         var defaultyVSwitch = new Switch(&#34;defaultyVSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultVpc.id())
 *             .cidrBlock(&#34;172.16.2.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultCluster = new Cluster(&#34;defaultCluster&#34;, ClusterArgs.builder()        
 *             .network(ClusterNetworkArgs.builder()
 *                 .vpcId(defaultVpc.id())
 *                 .vswitches(defaultyVSwitch.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Ack One Cluster can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ackone/cluster:Cluster example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ackone/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * Cluster name.
     * 
     */
    @Export(name="clusterName", refs={String.class}, tree="[0]")
    private Output<String> clusterName;

    /**
     * @return Cluster name.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }
    /**
     * Cluster creation time.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Cluster creation time.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Cluster network information. See `network` below.
     * 
     */
    @Export(name="network", refs={ClusterNetwork.class}, tree="[0]")
    private Output<ClusterNetwork> network;

    /**
     * @return Cluster network information. See `network` below.
     * 
     */
    public Output<ClusterNetwork> network() {
        return this.network;
    }
    /**
     * Cluster attributes. Valid values: &#39;Default&#39;, &#39;XFlow&#39;.
     * 
     */
    @Export(name="profile", refs={String.class}, tree="[0]")
    private Output<String> profile;

    /**
     * @return Cluster attributes. Valid values: &#39;Default&#39;, &#39;XFlow&#39;.
     * 
     */
    public Output<String> profile() {
        return this.profile;
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ackone/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ackone/cluster:Cluster", name, state, makeResourceOptions(options, id));
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
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
