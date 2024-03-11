// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cr;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cr.ChainArgs;
import com.pulumi.alicloud.cr.inputs.ChainState;
import com.pulumi.alicloud.cr.outputs.ChainChainConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CR Chain resource.
 * 
 * For information about CR Chain and how to use it, see [What is Chain](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createchain).
 * 
 * &gt; **NOTE:** Available since v1.161.0.
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
 * import com.pulumi.alicloud.cr.RegistryEnterpriseInstance;
 * import com.pulumi.alicloud.cr.RegistryEnterpriseInstanceArgs;
 * import com.pulumi.alicloud.cs.RegistryEnterpriseNamespace;
 * import com.pulumi.alicloud.cs.RegistryEnterpriseNamespaceArgs;
 * import com.pulumi.alicloud.cs.RegistryEnterpriseRepo;
 * import com.pulumi.alicloud.cs.RegistryEnterpriseRepoArgs;
 * import com.pulumi.alicloud.cr.Chain;
 * import com.pulumi.alicloud.cr.ChainArgs;
 * import com.pulumi.alicloud.cr.inputs.ChainChainConfigArgs;
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
 *         var defaultRegistryEnterpriseInstance = new RegistryEnterpriseInstance(&#34;defaultRegistryEnterpriseInstance&#34;, RegistryEnterpriseInstanceArgs.builder()        
 *             .paymentType(&#34;Subscription&#34;)
 *             .period(1)
 *             .renewPeriod(0)
 *             .renewalStatus(&#34;ManualRenewal&#34;)
 *             .instanceType(&#34;Advanced&#34;)
 *             .instanceName(name)
 *             .build());
 * 
 *         var defaultRegistryEnterpriseNamespace = new RegistryEnterpriseNamespace(&#34;defaultRegistryEnterpriseNamespace&#34;, RegistryEnterpriseNamespaceArgs.builder()        
 *             .instanceId(defaultRegistryEnterpriseInstance.id())
 *             .autoCreate(false)
 *             .defaultVisibility(&#34;PUBLIC&#34;)
 *             .build());
 * 
 *         var defaultRegistryEnterpriseRepo = new RegistryEnterpriseRepo(&#34;defaultRegistryEnterpriseRepo&#34;, RegistryEnterpriseRepoArgs.builder()        
 *             .instanceId(defaultRegistryEnterpriseInstance.id())
 *             .namespace(defaultRegistryEnterpriseNamespace.name())
 *             .summary(&#34;this is summary of my new repo&#34;)
 *             .repoType(&#34;PUBLIC&#34;)
 *             .detail(&#34;this is a public repo&#34;)
 *             .build());
 * 
 *         var defaultChain = new Chain(&#34;defaultChain&#34;, ChainArgs.builder()        
 *             .chainName(name)
 *             .description(name)
 *             .instanceId(defaultRegistryEnterpriseNamespace.instanceId())
 *             .repoName(defaultRegistryEnterpriseRepo.name())
 *             .repoNamespaceName(defaultRegistryEnterpriseNamespace.name())
 *             .chainConfigs(ChainChainConfigArgs.builder()
 *                 .routers(                
 *                     ChainChainConfigRouterArgs.builder()
 *                         .froms(ChainChainConfigRouterFromArgs.builder()
 *                             .nodeName(&#34;DOCKER_IMAGE_BUILD&#34;)
 *                             .build())
 *                         .tos(ChainChainConfigRouterToArgs.builder()
 *                             .nodeName(&#34;DOCKER_IMAGE_PUSH&#34;)
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigRouterArgs.builder()
 *                         .froms(ChainChainConfigRouterFromArgs.builder()
 *                             .nodeName(&#34;DOCKER_IMAGE_PUSH&#34;)
 *                             .build())
 *                         .tos(ChainChainConfigRouterToArgs.builder()
 *                             .nodeName(&#34;VULNERABILITY_SCANNING&#34;)
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigRouterArgs.builder()
 *                         .froms(ChainChainConfigRouterFromArgs.builder()
 *                             .nodeName(&#34;VULNERABILITY_SCANNING&#34;)
 *                             .build())
 *                         .tos(ChainChainConfigRouterToArgs.builder()
 *                             .nodeName(&#34;ACTIVATE_REPLICATION&#34;)
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigRouterArgs.builder()
 *                         .froms(ChainChainConfigRouterFromArgs.builder()
 *                             .nodeName(&#34;ACTIVATE_REPLICATION&#34;)
 *                             .build())
 *                         .tos(ChainChainConfigRouterToArgs.builder()
 *                             .nodeName(&#34;TRIGGER&#34;)
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigRouterArgs.builder()
 *                         .froms(ChainChainConfigRouterFromArgs.builder()
 *                             .nodeName(&#34;VULNERABILITY_SCANNING&#34;)
 *                             .build())
 *                         .tos(ChainChainConfigRouterToArgs.builder()
 *                             .nodeName(&#34;SNAPSHOT&#34;)
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigRouterArgs.builder()
 *                         .froms(ChainChainConfigRouterFromArgs.builder()
 *                             .nodeName(&#34;SNAPSHOT&#34;)
 *                             .build())
 *                         .tos(ChainChainConfigRouterToArgs.builder()
 *                             .nodeName(&#34;TRIGGER_SNAPSHOT&#34;)
 *                             .build())
 *                         .build())
 *                 .nodes(                
 *                     ChainChainConfigNodeArgs.builder()
 *                         .enable(true)
 *                         .nodeName(&#34;DOCKER_IMAGE_BUILD&#34;)
 *                         .nodeConfigs(ChainChainConfigNodeNodeConfigArgs.builder()
 *                             .denyPolicies()
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigNodeArgs.builder()
 *                         .enable(true)
 *                         .nodeName(&#34;DOCKER_IMAGE_PUSH&#34;)
 *                         .nodeConfigs(ChainChainConfigNodeNodeConfigArgs.builder()
 *                             .denyPolicies()
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigNodeArgs.builder()
 *                         .enable(true)
 *                         .nodeName(&#34;VULNERABILITY_SCANNING&#34;)
 *                         .nodeConfigs(ChainChainConfigNodeNodeConfigArgs.builder()
 *                             .denyPolicies(ChainChainConfigNodeNodeConfigDenyPolicyArgs.builder()
 *                                 .issueLevel(&#34;MEDIUM&#34;)
 *                                 .issueCount(1)
 *                                 .action(&#34;BLOCK_DELETE_TAG&#34;)
 *                                 .logic(&#34;AND&#34;)
 *                                 .build())
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigNodeArgs.builder()
 *                         .enable(true)
 *                         .nodeName(&#34;ACTIVATE_REPLICATION&#34;)
 *                         .nodeConfigs(ChainChainConfigNodeNodeConfigArgs.builder()
 *                             .denyPolicies()
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigNodeArgs.builder()
 *                         .enable(true)
 *                         .nodeName(&#34;TRIGGER&#34;)
 *                         .nodeConfigs(ChainChainConfigNodeNodeConfigArgs.builder()
 *                             .denyPolicies()
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigNodeArgs.builder()
 *                         .enable(false)
 *                         .nodeName(&#34;SNAPSHOT&#34;)
 *                         .nodeConfigs(ChainChainConfigNodeNodeConfigArgs.builder()
 *                             .denyPolicies()
 *                             .build())
 *                         .build(),
 *                     ChainChainConfigNodeArgs.builder()
 *                         .enable(false)
 *                         .nodeName(&#34;TRIGGER_SNAPSHOT&#34;)
 *                         .nodeConfigs(ChainChainConfigNodeNodeConfigArgs.builder()
 *                             .denyPolicies()
 *                             .build())
 *                         .build())
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
 * CR Chain can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cr/chain:Chain example &lt;instance_id&gt;:&lt;chain_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cr/chain:Chain")
public class Chain extends com.pulumi.resources.CustomResource {
    /**
     * The configuration of delivery chain. See `chain_config` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
     * 
     */
    @Export(name="chainConfigs", refs={List.class,ChainChainConfig.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ChainChainConfig>> chainConfigs;

    /**
     * @return The configuration of delivery chain. See `chain_config` below. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
     * 
     */
    public Output<Optional<List<ChainChainConfig>>> chainConfigs() {
        return Codegen.optional(this.chainConfigs);
    }
    /**
     * Delivery chain ID.
     * 
     */
    @Export(name="chainId", refs={String.class}, tree="[0]")
    private Output<String> chainId;

    /**
     * @return Delivery chain ID.
     * 
     */
    public Output<String> chainId() {
        return this.chainId;
    }
    /**
     * The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators &#34;_&#34;, &#34;-&#34;, &#34;.&#34; can be used, noted that the separator cannot be at the first or last position.
     * 
     */
    @Export(name="chainName", refs={String.class}, tree="[0]")
    private Output<String> chainName;

    /**
     * @return The name of delivery chain. The length of the name is 1-64 characters, lowercase English letters and numbers, and the separators &#34;_&#34;, &#34;-&#34;, &#34;.&#34; can be used, noted that the separator cannot be at the first or last position.
     * 
     */
    public Output<String> chainName() {
        return this.chainName;
    }
    /**
     * The description delivery chain.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description delivery chain.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ID of CR Enterprise Edition instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of CR Enterprise Edition instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
     * 
     */
    @Export(name="repoName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repoName;

    /**
     * @return The name of CR Enterprise Edition repository. **NOTE:** This parameter must specify a correct value, otherwise the created resource will be incorrect.
     * 
     */
    public Output<Optional<String>> repoName() {
        return Codegen.optional(this.repoName);
    }
    /**
     * The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
     * 
     */
    @Export(name="repoNamespaceName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repoNamespaceName;

    /**
     * @return The name of CR Enterprise Edition namespace. **NOTE:** This parameter must specify the correct value, otherwise the created resource will be incorrect.
     * 
     */
    public Output<Optional<String>> repoNamespaceName() {
        return Codegen.optional(this.repoNamespaceName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Chain(String name) {
        this(name, ChainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Chain(String name, ChainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Chain(String name, ChainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cr/chain:Chain", name, args == null ? ChainArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Chain(String name, Output<String> id, @Nullable ChainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cr/chain:Chain", name, state, makeResourceOptions(options, id));
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
    public static Chain get(String name, Output<String> id, @Nullable ChainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Chain(name, id, state, options);
    }
}
