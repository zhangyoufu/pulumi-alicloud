// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cassandra.inputs;

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
     * The list of Cassandra cluster ids.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return The list of Cassandra cluster ids.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to apply to the cluster name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to apply to the cluster name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The name of file that can save the collection of clusters after running `pulumi preview`.
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return The name of file that can save the collection of clusters after running `pulumi preview`.
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,String> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Map<String,String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetClustersPlainArgs() {}

    private GetClustersPlainArgs(GetClustersPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
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
         * @param ids The list of Cassandra cluster ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids The list of Cassandra cluster ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to apply to the cluster name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param outputFile The name of file that can save the collection of clusters after running `pulumi preview`.
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
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
