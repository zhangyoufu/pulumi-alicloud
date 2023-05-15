// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVpcFlowLogsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVpcFlowLogsPlainArgs Empty = new GetVpcFlowLogsPlainArgs();

    /**
     * The Description of flow log.
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return The Description of flow log.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The flow log name.
     * 
     */
    @Import(name="flowLogName")
    private @Nullable String flowLogName;

    /**
     * @return The flow log name.
     * 
     */
    public Optional<String> flowLogName() {
        return Optional.ofNullable(this.flowLogName);
    }

    /**
     * A list of Flow Log IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Flow Log IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The log store name.
     * 
     */
    @Import(name="logStoreName")
    private @Nullable String logStoreName;

    /**
     * @return The log store name.
     * 
     */
    public Optional<String> logStoreName() {
        return Optional.ofNullable(this.logStoreName);
    }

    /**
     * A regex string to filter results by Flow Log name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Flow Log name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
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
     * The project name.
     * 
     */
    @Import(name="projectName")
    private @Nullable String projectName;

    /**
     * @return The project name.
     * 
     */
    public Optional<String> projectName() {
        return Optional.ofNullable(this.projectName);
    }

    /**
     * The resource id.
     * 
     */
    @Import(name="resourceId")
    private @Nullable String resourceId;

    /**
     * @return The resource id.
     * 
     */
    public Optional<String> resourceId() {
        return Optional.ofNullable(this.resourceId);
    }

    /**
     * The resource type.
     * 
     */
    @Import(name="resourceType")
    private @Nullable String resourceType;

    /**
     * @return The resource type.
     * 
     */
    public Optional<String> resourceType() {
        return Optional.ofNullable(this.resourceType);
    }

    /**
     * The status of flow log.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of flow log.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The traffic type.
     * 
     */
    @Import(name="trafficType")
    private @Nullable String trafficType;

    /**
     * @return The traffic type.
     * 
     */
    public Optional<String> trafficType() {
        return Optional.ofNullable(this.trafficType);
    }

    private GetVpcFlowLogsPlainArgs() {}

    private GetVpcFlowLogsPlainArgs(GetVpcFlowLogsPlainArgs $) {
        this.description = $.description;
        this.flowLogName = $.flowLogName;
        this.ids = $.ids;
        this.logStoreName = $.logStoreName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.projectName = $.projectName;
        this.resourceId = $.resourceId;
        this.resourceType = $.resourceType;
        this.status = $.status;
        this.trafficType = $.trafficType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVpcFlowLogsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVpcFlowLogsPlainArgs $;

        public Builder() {
            $ = new GetVpcFlowLogsPlainArgs();
        }

        public Builder(GetVpcFlowLogsPlainArgs defaults) {
            $ = new GetVpcFlowLogsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The Description of flow log.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param flowLogName The flow log name.
         * 
         * @return builder
         * 
         */
        public Builder flowLogName(@Nullable String flowLogName) {
            $.flowLogName = flowLogName;
            return this;
        }

        /**
         * @param ids A list of Flow Log IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Flow Log IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param logStoreName The log store name.
         * 
         * @return builder
         * 
         */
        public Builder logStoreName(@Nullable String logStoreName) {
            $.logStoreName = logStoreName;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Flow Log name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
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

        /**
         * @param projectName The project name.
         * 
         * @return builder
         * 
         */
        public Builder projectName(@Nullable String projectName) {
            $.projectName = projectName;
            return this;
        }

        /**
         * @param resourceId The resource id.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(@Nullable String resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceType The resource type.
         * 
         * @return builder
         * 
         */
        public Builder resourceType(@Nullable String resourceType) {
            $.resourceType = resourceType;
            return this;
        }

        /**
         * @param status The status of flow log.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param trafficType The traffic type.
         * 
         * @return builder
         * 
         */
        public Builder trafficType(@Nullable String trafficType) {
            $.trafficType = trafficType;
            return this;
        }

        public GetVpcFlowLogsPlainArgs build() {
            return $;
        }
    }

}
