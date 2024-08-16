// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ros.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetStacksPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetStacksPlainArgs Empty = new GetStacksPlainArgs();

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
     * A list of Stack IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Stack IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Stack name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Stack name.
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
     * ParentStackId.
     * 
     */
    @Import(name="parentStackId")
    private @Nullable String parentStackId;

    /**
     * @return ParentStackId.
     * 
     */
    public Optional<String> parentStackId() {
        return Optional.ofNullable(this.parentStackId);
    }

    /**
     * The show nested stack.
     * 
     */
    @Import(name="showNestedStack")
    private @Nullable Boolean showNestedStack;

    /**
     * @return The show nested stack.
     * 
     */
    public Optional<Boolean> showNestedStack() {
        return Optional.ofNullable(this.showNestedStack);
    }

    /**
     * StackName.
     * 
     */
    @Import(name="stackName")
    private @Nullable String stackName;

    /**
     * @return StackName.
     * 
     */
    public Optional<String> stackName() {
        return Optional.ofNullable(this.stackName);
    }

    /**
     * The status of Stack. Valid Values: `CREATE_COMPLETE`, `CREATE_FAILED`, `CREATE_IN_PROGRESS`, `DELETE_COMPLETE`, `DELETE_FAILED`, `DELETE_IN_PROGRESS`, `ROLLBACK_COMPLETE`, `ROLLBACK_FAILED`, `ROLLBACK_IN_PROGRESS`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of Stack. Valid Values: `CREATE_COMPLETE`, `CREATE_FAILED`, `CREATE_IN_PROGRESS`, `DELETE_COMPLETE`, `DELETE_FAILED`, `DELETE_IN_PROGRESS`, `ROLLBACK_COMPLETE`, `ROLLBACK_FAILED`, `ROLLBACK_IN_PROGRESS`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{&#34;key1&#34;:&#34;value1&#34;}`.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,String> tags;

    /**
     * @return Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{&#34;key1&#34;:&#34;value1&#34;}`.
     * 
     */
    public Optional<Map<String,String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetStacksPlainArgs() {}

    private GetStacksPlainArgs(GetStacksPlainArgs $) {
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.parentStackId = $.parentStackId;
        this.showNestedStack = $.showNestedStack;
        this.stackName = $.stackName;
        this.status = $.status;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetStacksPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetStacksPlainArgs $;

        public Builder() {
            $ = new GetStacksPlainArgs();
        }

        public Builder(GetStacksPlainArgs defaults) {
            $ = new GetStacksPlainArgs(Objects.requireNonNull(defaults));
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
         * @param ids A list of Stack IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Stack IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Stack name.
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
         * @param parentStackId ParentStackId.
         * 
         * @return builder
         * 
         */
        public Builder parentStackId(@Nullable String parentStackId) {
            $.parentStackId = parentStackId;
            return this;
        }

        /**
         * @param showNestedStack The show nested stack.
         * 
         * @return builder
         * 
         */
        public Builder showNestedStack(@Nullable Boolean showNestedStack) {
            $.showNestedStack = showNestedStack;
            return this;
        }

        /**
         * @param stackName StackName.
         * 
         * @return builder
         * 
         */
        public Builder stackName(@Nullable String stackName) {
            $.stackName = stackName;
            return this;
        }

        /**
         * @param status The status of Stack. Valid Values: `CREATE_COMPLETE`, `CREATE_FAILED`, `CREATE_IN_PROGRESS`, `DELETE_COMPLETE`, `DELETE_FAILED`, `DELETE_IN_PROGRESS`, `ROLLBACK_COMPLETE`, `ROLLBACK_FAILED`, `ROLLBACK_IN_PROGRESS`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags Query the instance bound to the tag. The format of the incoming value is `json` string, including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `{&#34;key1&#34;:&#34;value1&#34;}`.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,String> tags) {
            $.tags = tags;
            return this;
        }

        public GetStacksPlainArgs build() {
            return $;
        }
    }

}
