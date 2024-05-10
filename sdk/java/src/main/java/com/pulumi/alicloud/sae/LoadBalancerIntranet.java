// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.sae.LoadBalancerIntranetArgs;
import com.pulumi.alicloud.sae.inputs.LoadBalancerIntranetState;
import com.pulumi.alicloud.sae.outputs.LoadBalancerIntranetIntranet;
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
 * For information about Serverless App Engine (SAE) Load Balancer Intranet Attachment and how to use it, see [alicloud.sae.LoadBalancerIntranet](https://www.alibabacloud.com/help/en/sae/latest/bindslb).
 * 
 * &gt; **NOTE:** Available since v1.165.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 * import com.pulumi.alicloud.sae.LoadBalancerIntranet;
 * import com.pulumi.alicloud.sae.LoadBalancerIntranetArgs;
 * import com.pulumi.alicloud.sae.inputs.LoadBalancerIntranetIntranetArgs;
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
 *         final var name = config.get("name").orElse("tf-example");
 *         final var default = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultInteger = new Integer("defaultInteger", IntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         final var defaultGetZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("VSwitch")
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock("10.4.0.0/16")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock("10.4.0.0/24")
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(defaultGetZones.applyValue(getZonesResult -> getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup("defaultSecurityGroup", SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultNamespace = new Namespace("defaultNamespace", NamespaceArgs.builder()        
 *             .namespaceId(String.format("%s:example%s", default_.regions()[0].id(),defaultInteger.result()))
 *             .namespaceName(name)
 *             .namespaceDescription(name)
 *             .enableMicroRegistration(false)
 *             .build());
 * 
 *         var defaultApplication = new Application("defaultApplication", ApplicationArgs.builder()        
 *             .appDescription(name)
 *             .appName(String.format("%s-%s", name,defaultInteger.result()))
 *             .namespaceId(defaultNamespace.id())
 *             .imageUrl("registry-vpc.cn-hangzhou.aliyuncs.com/lxepoo/apache-php5")
 *             .packageType("Image")
 *             .jdk("Open JDK 8")
 *             .securityGroupId(defaultSecurityGroup.id())
 *             .vpcId(defaultNetwork.id())
 *             .vswitchId(defaultSwitch.id())
 *             .timezone("Asia/Beijing")
 *             .replicas("5")
 *             .cpu("500")
 *             .memory("2048")
 *             .build());
 * 
 *         var defaultApplicationLoadBalancer = new ApplicationLoadBalancer("defaultApplicationLoadBalancer", ApplicationLoadBalancerArgs.builder()        
 *             .loadBalancerName(name)
 *             .vswitchId(defaultSwitch.id())
 *             .loadBalancerSpec("slb.s2.small")
 *             .addressType("intranet")
 *             .build());
 * 
 *         var defaultLoadBalancerIntranet = new LoadBalancerIntranet("defaultLoadBalancerIntranet", LoadBalancerIntranetArgs.builder()        
 *             .appId(defaultApplication.id())
 *             .intranetSlbId(defaultApplicationLoadBalancer.id())
 *             .intranets(LoadBalancerIntranetIntranetArgs.builder()
 *                 .protocol("TCP")
 *                 .port(80)
 *                 .targetPort(8080)
 *                 .build())
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
 * The resource can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:sae/loadBalancerIntranet:LoadBalancerIntranet example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:sae/loadBalancerIntranet:LoadBalancerIntranet")
public class LoadBalancerIntranet extends com.pulumi.resources.CustomResource {
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
     * Use designated private network SLBs that have been purchased to support non-shared instances.
     * 
     */
    @Export(name="intranetIp", refs={String.class}, tree="[0]")
    private Output<String> intranetIp;

    /**
     * @return Use designated private network SLBs that have been purchased to support non-shared instances.
     * 
     */
    public Output<String> intranetIp() {
        return this.intranetIp;
    }
    /**
     * The intranet SLB ID.
     * 
     */
    @Export(name="intranetSlbId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> intranetSlbId;

    /**
     * @return The intranet SLB ID.
     * 
     */
    public Output<Optional<String>> intranetSlbId() {
        return Codegen.optional(this.intranetSlbId);
    }
    /**
     * The bound private network SLB. See `intranet` below.
     * 
     */
    @Export(name="intranets", refs={List.class,LoadBalancerIntranetIntranet.class}, tree="[0,1]")
    private Output<List<LoadBalancerIntranetIntranet>> intranets;

    /**
     * @return The bound private network SLB. See `intranet` below.
     * 
     */
    public Output<List<LoadBalancerIntranetIntranet>> intranets() {
        return this.intranets;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancerIntranet(String name) {
        this(name, LoadBalancerIntranetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancerIntranet(String name, LoadBalancerIntranetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancerIntranet(String name, LoadBalancerIntranetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sae/loadBalancerIntranet:LoadBalancerIntranet", name, args == null ? LoadBalancerIntranetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoadBalancerIntranet(String name, Output<String> id, @Nullable LoadBalancerIntranetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sae/loadBalancerIntranet:LoadBalancerIntranet", name, state, makeResourceOptions(options, id));
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
    public static LoadBalancerIntranet get(String name, Output<String> id, @Nullable LoadBalancerIntranetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancerIntranet(name, id, state, options);
    }
}
