// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetClustersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetClustersPlainArgs Empty = new GetClustersPlainArgs();

    /**
     * Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
     * 
     */
    @Import(name="dbType")
    private @Nullable String dbType;

    /**
     * @return Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
     * 
     */
    public Optional<String> dbType() {
        return Optional.ofNullable(this.dbType);
    }

    /**
     * A regex string to filter results by cluster description.
     * 
     */
    @Import(name="descriptionRegex")
    private @Nullable String descriptionRegex;

    /**
     * @return A regex string to filter results by cluster description.
     * 
     */
    public Optional<String> descriptionRegex() {
        return Optional.ofNullable(this.descriptionRegex);
    }

    /**
     * A list of PolarDB cluster IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of PolarDB cluster IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * status of the cluster.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return status of the cluster.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,String> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    public Optional<Map<String,String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetClustersPlainArgs() {}

    private GetClustersPlainArgs(GetClustersPlainArgs $) {
        this.dbType = $.dbType;
        this.descriptionRegex = $.descriptionRegex;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.status = $.status;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetClustersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetClustersPlainArgs $;

        public Builder() {
            $ = new GetClustersPlainArgs();
        }

        public Builder(GetClustersPlainArgs defaults) {
            $ = new GetClustersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbType Database type. Options are `MySQL`, `Oracle` and `PostgreSQL`. If no value is specified, all types are returned.
         * 
         * @return builder
         * 
         */
        public Builder dbType(@Nullable String dbType) {
            $.dbType = dbType;
            return this;
        }

        /**
         * @param descriptionRegex A regex string to filter results by cluster description.
         * 
         * @return builder
         * 
         */
        public Builder descriptionRegex(@Nullable String descriptionRegex) {
            $.descriptionRegex = descriptionRegex;
            return this;
        }

        /**
         * @param ids A list of PolarDB cluster IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of PolarDB cluster IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param status status of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
         * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,String> tags) {
            $.tags = tags;
            return this;
        }

        public GetClustersPlainArgs build() {
            return $;
        }
    }

}
