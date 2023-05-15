// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHoneypotPresetsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetHoneypotPresetsPlainArgs Empty = new GetHoneypotPresetsPlainArgs();

    @Import(name="currentPage")
    private @Nullable Integer currentPage;

    public Optional<Integer> currentPage() {
        return Optional.ofNullable(this.currentPage);
    }

    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * Honeypot mirror name
     * 
     */
    @Import(name="honeypotImageName")
    private @Nullable String honeypotImageName;

    /**
     * @return Honeypot mirror name
     * 
     */
    public Optional<String> honeypotImageName() {
        return Optional.ofNullable(this.honeypotImageName);
    }

    /**
     * A list of Honeypot Preset IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Honeypot Preset IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    @Import(name="lang")
    private @Nullable String lang;

    public Optional<String> lang() {
        return Optional.ofNullable(this.lang);
    }

    /**
     * Unique id of management node
     * 
     */
    @Import(name="nodeId")
    private @Nullable String nodeId;

    /**
     * @return Unique id of management node
     * 
     */
    public Optional<String> nodeId() {
        return Optional.ofNullable(this.nodeId);
    }

    @Import(name="nodeName")
    private @Nullable String nodeName;

    public Optional<String> nodeName() {
        return Optional.ofNullable(this.nodeName);
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
     * Honeypot template custom name
     * 
     */
    @Import(name="presetName")
    private @Nullable String presetName;

    /**
     * @return Honeypot template custom name
     * 
     */
    public Optional<String> presetName() {
        return Optional.ofNullable(this.presetName);
    }

    private GetHoneypotPresetsPlainArgs() {}

    private GetHoneypotPresetsPlainArgs(GetHoneypotPresetsPlainArgs $) {
        this.currentPage = $.currentPage;
        this.enableDetails = $.enableDetails;
        this.honeypotImageName = $.honeypotImageName;
        this.ids = $.ids;
        this.lang = $.lang;
        this.nodeId = $.nodeId;
        this.nodeName = $.nodeName;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.presetName = $.presetName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHoneypotPresetsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHoneypotPresetsPlainArgs $;

        public Builder() {
            $ = new GetHoneypotPresetsPlainArgs();
        }

        public Builder(GetHoneypotPresetsPlainArgs defaults) {
            $ = new GetHoneypotPresetsPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder currentPage(@Nullable Integer currentPage) {
            $.currentPage = currentPage;
            return this;
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param honeypotImageName Honeypot mirror name
         * 
         * @return builder
         * 
         */
        public Builder honeypotImageName(@Nullable String honeypotImageName) {
            $.honeypotImageName = honeypotImageName;
            return this;
        }

        /**
         * @param ids A list of Honeypot Preset IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Honeypot Preset IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        public Builder lang(@Nullable String lang) {
            $.lang = lang;
            return this;
        }

        /**
         * @param nodeId Unique id of management node
         * 
         * @return builder
         * 
         */
        public Builder nodeId(@Nullable String nodeId) {
            $.nodeId = nodeId;
            return this;
        }

        public Builder nodeName(@Nullable String nodeName) {
            $.nodeName = nodeName;
            return this;
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

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param presetName Honeypot template custom name
         * 
         * @return builder
         * 
         */
        public Builder presetName(@Nullable String presetName) {
            $.presetName = presetName;
            return this;
        }

        public GetHoneypotPresetsPlainArgs build() {
            return $;
        }
    }

}
