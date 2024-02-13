// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dataworks;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dataworks.FolderArgs;
import com.pulumi.alicloud.dataworks.inputs.FolderState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Data Works Folder resource.
 * 
 * For information about Data Works Folder and how to use it, see [What is Folder](https://help.aliyun.com/document_detail/173940.html).
 * 
 * &gt; **NOTE:** Available in v1.131.0+.
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
 * import com.pulumi.alicloud.dataworks.Folder;
 * import com.pulumi.alicloud.dataworks.FolderArgs;
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
 *         var example = new Folder(&#34;example&#34;, FolderArgs.builder()        
 *             .folderPath(&#34;Business Flow/tfTestAcc/folderDi/tftest1&#34;)
 *             .projectId(&#34;320687&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Data Works Folder can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dataworks/folder:Folder example &lt;folder_id&gt;:&lt;$.ProjectId&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dataworks/folder:Folder")
public class Folder extends com.pulumi.resources.CustomResource {
    @Export(name="folderId", refs={String.class}, tree="[0]")
    private Output<String> folderId;

    public Output<String> folderId() {
        return this.folderId;
    }
    /**
     * Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
     * 
     */
    @Export(name="folderPath", refs={String.class}, tree="[0]")
    private Output<String> folderPath;

    /**
     * @return Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
     * 
     */
    public Output<String> folderPath() {
        return this.folderPath;
    }
    /**
     * The ID of the project.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectId;

    /**
     * @return The ID of the project.
     * 
     */
    public Output<Optional<String>> projectId() {
        return Codegen.optional(this.projectId);
    }
    @Export(name="projectIdentifier", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectIdentifier;

    public Output<Optional<String>> projectIdentifier() {
        return Codegen.optional(this.projectIdentifier);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Folder(String name) {
        this(name, FolderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Folder(String name, FolderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Folder(String name, FolderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dataworks/folder:Folder", name, args == null ? FolderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Folder(String name, Output<String> id, @Nullable FolderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dataworks/folder:Folder", name, state, makeResourceOptions(options, id));
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
    public static Folder get(String name, Output<String> id, @Nullable FolderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Folder(name, id, state, options);
    }
}
