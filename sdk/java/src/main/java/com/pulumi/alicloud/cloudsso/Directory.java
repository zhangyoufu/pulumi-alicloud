// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudsso;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudsso.DirectoryArgs;
import com.pulumi.alicloud.cloudsso.inputs.DirectoryState;
import com.pulumi.alicloud.cloudsso.outputs.DirectorySamlIdentityProviderConfiguration;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud SSO Directory resource.
 * 
 * For information about Cloud SSO Directory and how to use it, see [What is Directory](https://www.alibabacloud.com/help/en/cloudsso/latest/api-cloudsso-2021-05-15-createdirectory).
 * 
 * &gt; **NOTE:** Available since v1.135.0.
 * 
 * &gt; **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
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
 * import com.pulumi.alicloud.cloudsso.CloudssoFunctions;
 * import com.pulumi.alicloud.cloudsso.inputs.GetDirectoriesArgs;
 * import com.pulumi.alicloud.cloudsso.Directory;
 * import com.pulumi.alicloud.cloudsso.DirectoryArgs;
 * import com.pulumi.codegen.internal.KeyedValue;
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
 *         final var default = CloudssoFunctions.getDirectories();
 * 
 *         for (var i = 0; i < default_.ids().length() > 0 ? 0 : 1; i++) {
 *             new Directory("defaultDirectory-" + i, DirectoryArgs.builder()
 *                 .directoryName(name)
 *                 .build());
 * 
 *         
 * }
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Cloud SSO Directory can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cloudsso/directory:Directory example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudsso/directory:Directory")
public class Directory extends com.pulumi.resources.CustomResource {
    /**
     * The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
     * 
     */
    @Export(name="directoryName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> directoryName;

    /**
     * @return The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
     * 
     */
    public Output<Optional<String>> directoryName() {
        return Codegen.optional(this.directoryName);
    }
    /**
     * The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
     * 
     */
    @Export(name="mfaAuthenticationStatus", refs={String.class}, tree="[0]")
    private Output<String> mfaAuthenticationStatus;

    /**
     * @return The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
     * 
     */
    public Output<String> mfaAuthenticationStatus() {
        return this.mfaAuthenticationStatus;
    }
    /**
     * The saml identity provider configuration. See `saml_identity_provider_configuration` below.
     * 
     * &gt; **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
     * 
     */
    @Export(name="samlIdentityProviderConfiguration", refs={DirectorySamlIdentityProviderConfiguration.class}, tree="[0]")
    private Output<DirectorySamlIdentityProviderConfiguration> samlIdentityProviderConfiguration;

    /**
     * @return The saml identity provider configuration. See `saml_identity_provider_configuration` below.
     * 
     * &gt; **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
     * 
     */
    public Output<DirectorySamlIdentityProviderConfiguration> samlIdentityProviderConfiguration() {
        return this.samlIdentityProviderConfiguration;
    }
    /**
     * The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     * 
     */
    @Export(name="scimSynchronizationStatus", refs={String.class}, tree="[0]")
    private Output<String> scimSynchronizationStatus;

    /**
     * @return The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     * 
     */
    public Output<String> scimSynchronizationStatus() {
        return this.scimSynchronizationStatus;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Directory(java.lang.String name) {
        this(name, DirectoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Directory(java.lang.String name, @Nullable DirectoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Directory(java.lang.String name, @Nullable DirectoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudsso/directory:Directory", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Directory(java.lang.String name, Output<java.lang.String> id, @Nullable DirectoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudsso/directory:Directory", name, state, makeResourceOptions(options, id), false);
    }

    private static DirectoryArgs makeArgs(@Nullable DirectoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DirectoryArgs.Empty : args;
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
    public static Directory get(java.lang.String name, Output<java.lang.String> id, @Nullable DirectoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Directory(name, id, state, options);
    }
}
