// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.sae.IngressArgs;
import com.pulumi.alicloud.sae.inputs.IngressState;
import com.pulumi.alicloud.sae.outputs.IngressDefaultRule;
import com.pulumi.alicloud.sae.outputs.IngressRule;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Serverless App Engine (SAE) Ingress resource.
 * 
 * For information about Serverless App Engine (SAE) Ingress and how to use it, see [What is Ingress](https://www.alibabacloud.com/help/en/sae/latest/createingress).
 * 
 * &gt; **NOTE:** Available since v1.137.0.
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
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
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
 * import com.pulumi.alicloud.sae.Ingress;
 * import com.pulumi.alicloud.sae.IngressArgs;
 * import com.pulumi.alicloud.sae.inputs.IngressRuleArgs;
 * import com.pulumi.alicloud.sae.inputs.IngressDefaultRuleArgs;
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
 *         final var defaultRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         var defaultRandomInteger = new RandomInteger(&#34;defaultRandomInteger&#34;, RandomIntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
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
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultNamespace = new Namespace(&#34;defaultNamespace&#34;, NamespaceArgs.builder()        
 *             .namespaceId(defaultRandomInteger.result().applyValue(result -&gt; String.format(&#34;%s:example%s&#34;, defaultRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].id()),result)))
 *             .namespaceName(name)
 *             .namespaceDescription(name)
 *             .enableMicroRegistration(false)
 *             .build());
 * 
 *         var defaultApplication = new Application(&#34;defaultApplication&#34;, ApplicationArgs.builder()        
 *             .appDescription(name)
 *             .appName(defaultRandomInteger.result().applyValue(result -&gt; String.format(&#34;%s-%s&#34;, name,result)))
 *             .namespaceId(defaultNamespace.id())
 *             .imageUrl(String.format(&#34;registry-vpc.%s.aliyuncs.com/sae-demo-image/consumer:1.0&#34;, defaultRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].id())))
 *             .packageType(&#34;Image&#34;)
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
 *             .addressType(&#34;intranet&#34;)
 *             .build());
 * 
 *         var defaultIngress = new Ingress(&#34;defaultIngress&#34;, IngressArgs.builder()        
 *             .slbId(defaultApplicationLoadBalancer.id())
 *             .namespaceId(defaultNamespace.id())
 *             .listenerPort(&#34;80&#34;)
 *             .rules(IngressRuleArgs.builder()
 *                 .appId(defaultApplication.id())
 *                 .containerPort(&#34;443&#34;)
 *                 .domain(&#34;www.alicloud.com&#34;)
 *                 .appName(defaultApplication.appName())
 *                 .path(&#34;/&#34;)
 *                 .build())
 *             .defaultRule(IngressDefaultRuleArgs.builder()
 *                 .appId(defaultApplication.id())
 *                 .containerPort(&#34;443&#34;)
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
 * Serverless App Engine (SAE) Ingress can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:sae/ingress:Ingress example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:sae/ingress:Ingress")
public class Ingress extends com.pulumi.resources.CustomResource {
    /**
     * The certificate ID of the HTTPS listener. The `cert_id` takes effect only when `load_balance_type` is set to `clb`.
     * 
     */
    @Export(name="certId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> certId;

    /**
     * @return The certificate ID of the HTTPS listener. The `cert_id` takes effect only when `load_balance_type` is set to `clb`.
     * 
     */
    public Output<Optional<String>> certId() {
        return Codegen.optional(this.certId);
    }
    /**
     * The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `cert_ids` takes effect only when `load_balance_type` is set to `alb`.
     * 
     */
    @Export(name="certIds", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> certIds;

    /**
     * @return The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `cert_ids` takes effect only when `load_balance_type` is set to `alb`.
     * 
     */
    public Output<Optional<String>> certIds() {
        return Codegen.optional(this.certIds);
    }
    /**
     * Default Rule. See `default_rule` below.
     * 
     */
    @Export(name="defaultRule", refs={IngressDefaultRule.class}, tree="[0]")
    private Output</* @Nullable */ IngressDefaultRule> defaultRule;

    /**
     * @return Default Rule. See `default_rule` below.
     * 
     */
    public Output<Optional<IngressDefaultRule>> defaultRule() {
        return Codegen.optional(this.defaultRule);
    }
    /**
     * Description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * SLB listening port.
     * 
     */
    @Export(name="listenerPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> listenerPort;

    /**
     * @return SLB listening port.
     * 
     */
    public Output<Integer> listenerPort() {
        return this.listenerPort;
    }
    /**
     * The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
     * 
     */
    @Export(name="listenerProtocol", refs={String.class}, tree="[0]")
    private Output<String> listenerProtocol;

    /**
     * @return The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
     * 
     */
    public Output<String> listenerProtocol() {
        return this.listenerProtocol;
    }
    /**
     * The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
     * 
     */
    @Export(name="loadBalanceType", refs={String.class}, tree="[0]")
    private Output<String> loadBalanceType;

    /**
     * @return The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
     * 
     */
    public Output<String> loadBalanceType() {
        return this.loadBalanceType;
    }
    /**
     * The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
     * 
     */
    @Export(name="namespaceId", refs={String.class}, tree="[0]")
    private Output<String> namespaceId;

    /**
     * @return The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
     * 
     */
    public Output<String> namespaceId() {
        return this.namespaceId;
    }
    /**
     * Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
     * 
     */
    @Export(name="rules", refs={List.class,IngressRule.class}, tree="[0,1]")
    private Output<List<IngressRule>> rules;

    /**
     * @return Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
     * 
     */
    public Output<List<IngressRule>> rules() {
        return this.rules;
    }
    /**
     * SLB ID.
     * 
     */
    @Export(name="slbId", refs={String.class}, tree="[0]")
    private Output<String> slbId;

    /**
     * @return SLB ID.
     * 
     */
    public Output<String> slbId() {
        return this.slbId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ingress(String name) {
        this(name, IngressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ingress(String name, IngressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ingress(String name, IngressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sae/ingress:Ingress", name, args == null ? IngressArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ingress(String name, Output<String> id, @Nullable IngressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sae/ingress:Ingress", name, state, makeResourceOptions(options, id));
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
    public static Ingress get(String name, Output<String> id, @Nullable IngressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ingress(name, id, state, options);
    }
}
