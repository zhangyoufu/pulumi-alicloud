// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.sae.LoadBalancerInternetArgs;
import com.pulumi.alicloud.sae.inputs.LoadBalancerInternetState;
import com.pulumi.alicloud.sae.outputs.LoadBalancerInternetInternet;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Alicloud Serverless App Engine (SAE) Application Load Balancer Attachment resource.
 * 
 * For information about Serverless App Engine (SAE) Load Balancer Internet Attachment and how to use it, see [alicloud.sae.LoadBalancerInternet](https://www.alibabacloud.com/help/en/sae/latest/bindslb).
 * 
 * &gt; **NOTE:** Available since v1.164.0.
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
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.sae.Namespace;
 * import com.pulumi.alicloud.sae.NamespaceArgs;
 * import com.pulumi.alicloud.sae.Application;
 * import com.pulumi.alicloud.sae.ApplicationArgs;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancer;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancerArgs;
 * import com.pulumi.alicloud.sae.LoadBalancerInternet;
 * import com.pulumi.alicloud.sae.LoadBalancerInternetArgs;
 * import com.pulumi.alicloud.sae.inputs.LoadBalancerInternetInternetArgs;
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
 *         final var default = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultInteger = new Integer(&#34;defaultInteger&#34;, IntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         final var defaultGetZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.4.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;10.4.0.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(defaultGetZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultNamespace = new Namespace(&#34;defaultNamespace&#34;, NamespaceArgs.builder()        
 *             .namespaceId(String.format(&#34;%s:example%s&#34;, default_.regions()[0].id(),defaultInteger.result()))
 *             .namespaceName(name)
 *             .namespaceDescription(name)
 *             .enableMicroRegistration(false)
 *             .build());
 * 
 *         var defaultApplication = new Application(&#34;defaultApplication&#34;, ApplicationArgs.builder()        
 *             .appDescription(name)
 *             .appName(String.format(&#34;%s-%s&#34;, name,defaultInteger.result()))
 *             .namespaceId(defaultNamespace.id())
 *             .imageUrl(&#34;registry-vpc.cn-hangzhou.aliyuncs.com/lxepoo/apache-php5&#34;)
 *             .packageType(&#34;Image&#34;)
 *             .jdk(&#34;Open JDK 8&#34;)
 *             .securityGroupId(defaultSecurityGroup.id())
 *             .vpcId(defaultNetwork.id())
 *             .vswitchId(defaultSwitch.id())
 *             .timezone(&#34;Asia/Beijing&#34;)
 *             .replicas(&#34;5&#34;)
 *             .cpu(&#34;500&#34;)
 *             .memory(&#34;2048&#34;)
 *             .build());
 * 
 *         var defaultApplicationLoadBalancer = new ApplicationLoadBalancer(&#34;defaultApplicationLoadBalancer&#34;, ApplicationLoadBalancerArgs.builder()        
 *             .loadBalancerName(name)
 *             .vswitchId(defaultSwitch.id())
 *             .loadBalancerSpec(&#34;slb.s2.small&#34;)
 *             .addressType(&#34;internet&#34;)
 *             .build());
 * 
 *         var defaultLoadBalancerInternet = new LoadBalancerInternet(&#34;defaultLoadBalancerInternet&#34;, LoadBalancerInternetArgs.builder()        
 *             .appId(defaultApplication.id())
 *             .internetSlbId(defaultApplicationLoadBalancer.id())
 *             .internets(LoadBalancerInternetInternetArgs.builder()
 *                 .protocol(&#34;TCP&#34;)
 *                 .port(80)
 *                 .targetPort(8080)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * The resource can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:sae/loadBalancerInternet:LoadBalancerInternet example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:sae/loadBalancerInternet:LoadBalancerInternet")
public class LoadBalancerInternet extends com.pulumi.resources.CustomResource {
    /**
     * The target application ID that needs to be bound to the SLB.
     * 
     */
    @Export(name="appId", refs={String.class}, tree="[0]")
    private Output<String> appId;

    /**
     * @return The target application ID that needs to be bound to the SLB.
     * 
     */
    public Output<String> appId() {
        return this.appId;
    }
    /**
     * Use designated public network SLBs that have been purchased to support non-shared instances.
     * 
     */
    @Export(name="internetIp", refs={String.class}, tree="[0]")
    private Output<String> internetIp;

    /**
     * @return Use designated public network SLBs that have been purchased to support non-shared instances.
     * 
     */
    public Output<String> internetIp() {
        return this.internetIp;
    }
    /**
     * The internet SLB ID.
     * 
     */
    @Export(name="internetSlbId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> internetSlbId;

    /**
     * @return The internet SLB ID.
     * 
     */
    public Output<Optional<String>> internetSlbId() {
        return Codegen.optional(this.internetSlbId);
    }
    /**
     * The bound private network SLB. See `internet` below.
     * 
     */
    @Export(name="internets", refs={List.class,LoadBalancerInternetInternet.class}, tree="[0,1]")
    private Output<List<LoadBalancerInternetInternet>> internets;

    /**
     * @return The bound private network SLB. See `internet` below.
     * 
     */
    public Output<List<LoadBalancerInternetInternet>> internets() {
        return this.internets;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancerInternet(String name) {
        this(name, LoadBalancerInternetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancerInternet(String name, LoadBalancerInternetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancerInternet(String name, LoadBalancerInternetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sae/loadBalancerInternet:LoadBalancerInternet", name, args == null ? LoadBalancerInternetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoadBalancerInternet(String name, Output<String> id, @Nullable LoadBalancerInternetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sae/loadBalancerInternet:LoadBalancerInternet", name, state, makeResourceOptions(options, id));
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
    public static LoadBalancerInternet get(String name, Output<String> id, @Nullable LoadBalancerInternetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancerInternet(name, id, state, options);
    }
}
