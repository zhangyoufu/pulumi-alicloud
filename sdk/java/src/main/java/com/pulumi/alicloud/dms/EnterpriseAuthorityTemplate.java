// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dms.EnterpriseAuthorityTemplateArgs;
import com.pulumi.alicloud.dms.inputs.EnterpriseAuthorityTemplateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DMS Enterprise Authority Template resource.
 * 
 * For information about DMS Enterprise Authority Template and how to use it, see [What is Authority Template](https://www.alibabacloud.com/help/en/dms/developer-reference/api-dms-enterprise-2018-11-01-createauthoritytemplate).
 * 
 * &gt; **NOTE:** Available since v1.212.0.
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
 * import com.pulumi.alicloud.dms.DmsFunctions;
 * import com.pulumi.alicloud.dms.inputs.GetUserTenantsArgs;
 * import com.pulumi.alicloud.dms.EnterpriseAuthorityTemplate;
 * import com.pulumi.alicloud.dms.EnterpriseAuthorityTemplateArgs;
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
 *         final var name = config.get("name").orElse("terraform-example");
 *         final var default = DmsFunctions.getUserTenants(GetUserTenantsArgs.builder()
 *             .status("ACTIVE")
 *             .build());
 * 
 *         var defaultEnterpriseAuthorityTemplate = new EnterpriseAuthorityTemplate("defaultEnterpriseAuthorityTemplate", EnterpriseAuthorityTemplateArgs.builder()
 *             .tid(default_.ids()[0])
 *             .authorityTemplateName(name)
 *             .description(name)
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
 * DMS Enterprise Authority Template can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dms/enterpriseAuthorityTemplate:EnterpriseAuthorityTemplate example &lt;tid&gt;:&lt;authority_template_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dms/enterpriseAuthorityTemplate:EnterpriseAuthorityTemplate")
public class EnterpriseAuthorityTemplate extends com.pulumi.resources.CustomResource {
    /**
     * Permission template ID.
     * 
     */
    @Export(name="authorityTemplateId", refs={Integer.class}, tree="[0]")
    private Output<Integer> authorityTemplateId;

    /**
     * @return Permission template ID.
     * 
     */
    public Output<Integer> authorityTemplateId() {
        return this.authorityTemplateId;
    }
    /**
     * Permission Template name.
     * 
     */
    @Export(name="authorityTemplateName", refs={String.class}, tree="[0]")
    private Output<String> authorityTemplateName;

    /**
     * @return Permission Template name.
     * 
     */
    public Output<String> authorityTemplateName() {
        return this.authorityTemplateName;
    }
    /**
     * The creation time of the resource.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Permission template description information.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Permission template description information.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Tenant ID.
     * 
     */
    @Export(name="tid", refs={Integer.class}, tree="[0]")
    private Output<Integer> tid;

    /**
     * @return Tenant ID.
     * 
     */
    public Output<Integer> tid() {
        return this.tid;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EnterpriseAuthorityTemplate(java.lang.String name) {
        this(name, EnterpriseAuthorityTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnterpriseAuthorityTemplate(java.lang.String name, EnterpriseAuthorityTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnterpriseAuthorityTemplate(java.lang.String name, EnterpriseAuthorityTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dms/enterpriseAuthorityTemplate:EnterpriseAuthorityTemplate", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EnterpriseAuthorityTemplate(java.lang.String name, Output<java.lang.String> id, @Nullable EnterpriseAuthorityTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dms/enterpriseAuthorityTemplate:EnterpriseAuthorityTemplate", name, state, makeResourceOptions(options, id), false);
    }

    private static EnterpriseAuthorityTemplateArgs makeArgs(EnterpriseAuthorityTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EnterpriseAuthorityTemplateArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static EnterpriseAuthorityTemplate get(java.lang.String name, Output<java.lang.String> id, @Nullable EnterpriseAuthorityTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnterpriseAuthorityTemplate(name, id, state, options);
    }
}
