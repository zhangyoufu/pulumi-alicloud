// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPluginsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPluginsPlainArgs Empty = new GetPluginsPlainArgs();

    /**
     * A list of Plugin IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Plugin IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Plugin name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Plugin name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="pageNumber")
    private @Nullable Integer pageNumber;

    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Integer pageSize;

    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    /**
     * The name of the plug-in that you want to create.
     * 
     */
    @Import(name="pluginName")
    private @Nullable String pluginName;

    /**
     * @return The name of the plug-in that you want to create.
     * 
     */
    public Optional<String> pluginName() {
        return Optional.ofNullable(this.pluginName);
    }

    /**
     * The type of the plug-in.
     * 
     */
    @Import(name="pluginType")
    private @Nullable String pluginType;

    /**
     * @return The type of the plug-in.
     * 
     */
    public Optional<String> pluginType() {
        return Optional.ofNullable(this.pluginType);
    }

    /**
     * The tag of the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,Object> tags;

    /**
     * @return The tag of the resource.
     * 
     */
    public Optional<Map<String,Object>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetPluginsPlainArgs() {}

    private GetPluginsPlainArgs(GetPluginsPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.pluginName = $.pluginName;
        this.pluginType = $.pluginType;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPluginsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPluginsPlainArgs $;

        public Builder() {
            $ = new GetPluginsPlainArgs();
        }

        public Builder(GetPluginsPlainArgs defaults) {
            $ = new GetPluginsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Plugin IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Plugin IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Plugin name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param pluginName The name of the plug-in that you want to create.
         * 
         * @return builder
         * 
         */
        public Builder pluginName(@Nullable String pluginName) {
            $.pluginName = pluginName;
            return this;
        }

        /**
         * @param pluginType The type of the plug-in.
         * 
         * @return builder
         * 
         */
        public Builder pluginType(@Nullable String pluginType) {
            $.pluginType = pluginType;
            return this;
        }

        /**
         * @param tags The tag of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,Object> tags) {
            $.tags = tags;
            return this;
        }

        public GetPluginsPlainArgs build() {
            return $;
        }
    }

}
