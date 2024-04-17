// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cr;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cr.ChartRepositoryArgs;
import com.pulumi.alicloud.cr.inputs.ChartRepositoryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CR Chart Repository resource.
 * 
 * For information about CR Chart Repository and how to use it, see [What is Chart Repository](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createchartrepository).
 * 
 * &gt; **NOTE:** Available since v1.149.0.
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
 * import com.pulumi.alicloud.cr.ChartNamespace;
 * import com.pulumi.alicloud.cr.ChartNamespaceArgs;
 * import com.pulumi.alicloud.cr.ChartRepository;
 * import com.pulumi.alicloud.cr.ChartRepositoryArgs;
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
 *         var example = new RegistryEnterpriseInstance(&#34;example&#34;, RegistryEnterpriseInstanceArgs.builder()        
 *             .paymentType(&#34;Subscription&#34;)
 *             .period(1)
 *             .renewPeriod(0)
 *             .renewalStatus(&#34;ManualRenewal&#34;)
 *             .instanceType(&#34;Advanced&#34;)
 *             .instanceName(name)
 *             .build());
 * 
 *         var exampleChartNamespace = new ChartNamespace(&#34;exampleChartNamespace&#34;, ChartNamespaceArgs.builder()        
 *             .instanceId(example.id())
 *             .namespaceName(name)
 *             .build());
 * 
 *         var exampleChartRepository = new ChartRepository(&#34;exampleChartRepository&#34;, ChartRepositoryArgs.builder()        
 *             .repoNamespaceName(exampleChartNamespace.namespaceName())
 *             .instanceId(exampleChartNamespace.instanceId())
 *             .repoName(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * CR Chart Repository can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cr/chartRepository:ChartRepository example &lt;instance_id&gt;:&lt;repo_namespace_name&gt;:&lt;repo_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cr/chartRepository:ChartRepository")
public class ChartRepository extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Container Registry instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the Container Registry instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The name of the repository that you want to create.
     * 
     */
    @Export(name="repoName", refs={String.class}, tree="[0]")
    private Output<String> repoName;

    /**
     * @return The name of the repository that you want to create.
     * 
     */
    public Output<String> repoName() {
        return this.repoName;
    }
    /**
     * The namespace to which the repository belongs.
     * 
     */
    @Export(name="repoNamespaceName", refs={String.class}, tree="[0]")
    private Output<String> repoNamespaceName;

    /**
     * @return The namespace to which the repository belongs.
     * 
     */
    public Output<String> repoNamespaceName() {
        return this.repoNamespaceName;
    }
    /**
     * The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
     * 
     */
    @Export(name="repoType", refs={String.class}, tree="[0]")
    private Output<String> repoType;

    /**
     * @return The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
     * 
     */
    public Output<String> repoType() {
        return this.repoType;
    }
    /**
     * The summary about the repository.
     * 
     */
    @Export(name="summary", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> summary;

    /**
     * @return The summary about the repository.
     * 
     */
    public Output<Optional<String>> summary() {
        return Codegen.optional(this.summary);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ChartRepository(String name) {
        this(name, ChartRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ChartRepository(String name, ChartRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ChartRepository(String name, ChartRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cr/chartRepository:ChartRepository", name, args == null ? ChartRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ChartRepository(String name, Output<String> id, @Nullable ChartRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cr/chartRepository:ChartRepository", name, state, makeResourceOptions(options, id));
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
    public static ChartRepository get(String name, Output<String> id, @Nullable ChartRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ChartRepository(name, id, state, options);
    }
}
