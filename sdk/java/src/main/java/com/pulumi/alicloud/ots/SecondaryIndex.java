// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ots.SecondaryIndexArgs;
import com.pulumi.alicloud.ots.inputs.SecondaryIndexState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an OTS secondary index resource.
 * 
 * For information about OTS secondary index and how to use it, see [Secondary index overview](https://www.alibabacloud.com/help/en/tablestore/latest/secondary-index-overview).
 * 
 * &gt; **NOTE:** Available since v1.187.0.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * OTS secondary index can be imported using id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ots/secondaryIndex:SecondaryIndex index1 &lt;instance_name&gt;:&lt;table_name&gt;:&lt;index_name&gt;:&lt;index_type&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ots/secondaryIndex:SecondaryIndex")
public class SecondaryIndex extends com.pulumi.resources.CustomResource {
    /**
     * A list of defined column for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    @Export(name="definedColumns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> definedColumns;

    /**
     * @return A list of defined column for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    public Output<Optional<List<String>>> definedColumns() {
        return Codegen.optional(this.definedColumns);
    }
    /**
     * whether the index contains data that already exists in the data table. When include_base_data is set to true, it means that stock data is included.
     * 
     */
    @Export(name="includeBaseData", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> includeBaseData;

    /**
     * @return whether the index contains data that already exists in the data table. When include_base_data is set to true, it means that stock data is included.
     * 
     */
    public Output<Boolean> includeBaseData() {
        return this.includeBaseData;
    }
    /**
     * The index name of the OTS Table. If changed, a new index would be created.
     * 
     */
    @Export(name="indexName", refs={String.class}, tree="[0]")
    private Output<String> indexName;

    /**
     * @return The index name of the OTS Table. If changed, a new index would be created.
     * 
     */
    public Output<String> indexName() {
        return this.indexName;
    }
    /**
     * The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
     * 
     */
    @Export(name="indexType", refs={String.class}, tree="[0]")
    private Output<String> indexType;

    /**
     * @return The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
     * 
     */
    public Output<String> indexType() {
        return this.indexType;
    }
    /**
     * The name of the OTS instance in which table will located.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output<String> instanceName;

    /**
     * @return The name of the OTS instance in which table will located.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }
    /**
     * A list of primary keys for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    @Export(name="primaryKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> primaryKeys;

    /**
     * @return A list of primary keys for index, referenced from Table&#39;s primary keys or predefined columns.
     * 
     */
    public Output<List<String>> primaryKeys() {
        return this.primaryKeys;
    }
    /**
     * The name of the OTS table. If changed, a new table would be created.
     * 
     */
    @Export(name="tableName", refs={String.class}, tree="[0]")
    private Output<String> tableName;

    /**
     * @return The name of the OTS table. If changed, a new table would be created.
     * 
     */
    public Output<String> tableName() {
        return this.tableName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecondaryIndex(String name) {
        this(name, SecondaryIndexArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecondaryIndex(String name, SecondaryIndexArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecondaryIndex(String name, SecondaryIndexArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ots/secondaryIndex:SecondaryIndex", name, args == null ? SecondaryIndexArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecondaryIndex(String name, Output<String> id, @Nullable SecondaryIndexState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ots/secondaryIndex:SecondaryIndex", name, state, makeResourceOptions(options, id));
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
    public static SecondaryIndex get(String name, Output<String> id, @Nullable SecondaryIndexState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecondaryIndex(name, id, state, options);
    }
}
