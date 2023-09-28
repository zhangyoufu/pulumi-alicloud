// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.nlb.ServerGroupArgs;
import com.pulumi.alicloud.nlb.inputs.ServerGroupState;
import com.pulumi.alicloud.nlb.outputs.ServerGroupHealthCheck;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a NLB Server Group resource.
 * 
 * For information about NLB Server Group and how to use it, see [What is Server Group](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createservergroup-nlb).
 * 
 * &gt; **NOTE:** Available since v1.186.0.
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
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.nlb.ServerGroup;
 * import com.pulumi.alicloud.nlb.ServerGroupArgs;
 * import com.pulumi.alicloud.nlb.inputs.ServerGroupHealthCheckArgs;
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
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.4.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultServerGroup = new ServerGroup(&#34;defaultServerGroup&#34;, ServerGroupArgs.builder()        
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .serverGroupName(name)
 *             .serverGroupType(&#34;Instance&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .scheduler(&#34;Wrr&#34;)
 *             .protocol(&#34;TCP&#34;)
 *             .connectionDrain(true)
 *             .connectionDrainTimeout(60)
 *             .addressIpVersion(&#34;Ipv4&#34;)
 *             .healthCheck(ServerGroupHealthCheckArgs.builder()
 *                 .healthCheckEnabled(true)
 *                 .healthCheckType(&#34;TCP&#34;)
 *                 .healthCheckConnectPort(0)
 *                 .healthyThreshold(2)
 *                 .unhealthyThreshold(2)
 *                 .healthCheckConnectTimeout(5)
 *                 .healthCheckInterval(10)
 *                 .httpCheckMethod(&#34;GET&#34;)
 *                 .healthCheckHttpCodes(                
 *                     &#34;http_2xx&#34;,
 *                     &#34;http_3xx&#34;,
 *                     &#34;http_4xx&#34;)
 *                 .build())
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;TF&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;example&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * NLB Server Group can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:nlb/serverGroup:ServerGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:nlb/serverGroup:ServerGroup")
public class ServerGroup extends com.pulumi.resources.CustomResource {
    /**
     * The protocol version. Valid values: `Ipv4` (default), `DualStack`.
     * 
     */
    @Export(name="addressIpVersion", type=String.class, parameters={})
    private Output<String> addressIpVersion;

    /**
     * @return The protocol version. Valid values: `Ipv4` (default), `DualStack`.
     * 
     */
    public Output<String> addressIpVersion() {
        return this.addressIpVersion;
    }
    /**
     * Specifies whether to enable connection draining.
     * 
     */
    @Export(name="connectionDrain", type=Boolean.class, parameters={})
    private Output<Boolean> connectionDrain;

    /**
     * @return Specifies whether to enable connection draining.
     * 
     */
    public Output<Boolean> connectionDrain() {
        return this.connectionDrain;
    }
    /**
     * The timeout period of connection draining. Unit: seconds. Valid values: 10 to 900.
     * 
     */
    @Export(name="connectionDrainTimeout", type=Integer.class, parameters={})
    private Output<Integer> connectionDrainTimeout;

    /**
     * @return The timeout period of connection draining. Unit: seconds. Valid values: 10 to 900.
     * 
     */
    public Output<Integer> connectionDrainTimeout() {
        return this.connectionDrainTimeout;
    }
    /**
     * HealthCheck. See `health_check` below.
     * 
     */
    @Export(name="healthCheck", type=ServerGroupHealthCheck.class, parameters={})
    private Output<ServerGroupHealthCheck> healthCheck;

    /**
     * @return HealthCheck. See `health_check` below.
     * 
     */
    public Output<ServerGroupHealthCheck> healthCheck() {
        return this.healthCheck;
    }
    /**
     * Indicates whether client address retention is enabled.
     * 
     */
    @Export(name="preserveClientIpEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> preserveClientIpEnabled;

    /**
     * @return Indicates whether client address retention is enabled.
     * 
     */
    public Output<Boolean> preserveClientIpEnabled() {
        return this.preserveClientIpEnabled;
    }
    /**
     * The backend protocol. Valid values: `TCP` (default), `UDP`, and `TCPSSL`.
     * 
     */
    @Export(name="protocol", type=String.class, parameters={})
    private Output<String> protocol;

    /**
     * @return The backend protocol. Valid values: `TCP` (default), `UDP`, and `TCPSSL`.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The ID of the resource group to which the security group belongs.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the security group belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The routing algorithm. Valid values:
     * 
     */
    @Export(name="scheduler", type=String.class, parameters={})
    private Output<String> scheduler;

    /**
     * @return The routing algorithm. Valid values:
     * 
     */
    public Output<String> scheduler() {
        return this.scheduler;
    }
    /**
     * The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    @Export(name="serverGroupName", type=String.class, parameters={})
    private Output<String> serverGroupName;

    /**
     * @return The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    public Output<String> serverGroupName() {
        return this.serverGroupName;
    }
    /**
     * The type of the server group. Valid values:
     * 
     */
    @Export(name="serverGroupType", type=String.class, parameters={})
    private Output<String> serverGroupType;

    /**
     * @return The type of the server group. Valid values:
     * 
     */
    public Output<String> serverGroupType() {
        return this.serverGroupType;
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The id of the vpc.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The id of the vpc.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServerGroup(String name) {
        this(name, ServerGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServerGroup(String name, ServerGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServerGroup(String name, ServerGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nlb/serverGroup:ServerGroup", name, args == null ? ServerGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServerGroup(String name, Output<String> id, @Nullable ServerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nlb/serverGroup:ServerGroup", name, state, makeResourceOptions(options, id));
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
    public static ServerGroup get(String name, Output<String> id, @Nullable ServerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServerGroup(name, id, state, options);
    }
}
