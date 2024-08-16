// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetExecutionsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetExecutionsPlainArgs Empty = new GetExecutionsPlainArgs();

    /**
     * The category of template. Valid: `AlarmTrigger`, `EventTrigger`, `Other` and `TimerTrigger`.
     * 
     */
    @Import(name="category")
    private @Nullable String category;

    /**
     * @return The category of template. Valid: `AlarmTrigger`, `EventTrigger`, `Other` and `TimerTrigger`.
     * 
     */
    public Optional<String> category() {
        return Optional.ofNullable(this.category);
    }

    /**
     * The time when the execution was ended.
     * 
     */
    @Import(name="endDate")
    private @Nullable String endDate;

    /**
     * @return The time when the execution was ended.
     * 
     */
    public Optional<String> endDate() {
        return Optional.ofNullable(this.endDate);
    }

    /**
     * Execution whose end time is less than or equal to the specified time.
     * 
     */
    @Import(name="endDateAfter")
    private @Nullable String endDateAfter;

    /**
     * @return Execution whose end time is less than or equal to the specified time.
     * 
     */
    public Optional<String> endDateAfter() {
        return Optional.ofNullable(this.endDateAfter);
    }

    /**
     * The user who execute the template.
     * 
     */
    @Import(name="executedBy")
    private @Nullable String executedBy;

    /**
     * @return The user who execute the template.
     * 
     */
    public Optional<String> executedBy() {
        return Optional.ofNullable(this.executedBy);
    }

    /**
     * A list of OOS Execution ids.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of OOS Execution ids.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * Whether to include sub-execution.
     * 
     */
    @Import(name="includeChildExecution")
    private @Nullable Boolean includeChildExecution;

    /**
     * @return Whether to include sub-execution.
     * 
     */
    public Optional<Boolean> includeChildExecution() {
        return Optional.ofNullable(this.includeChildExecution);
    }

    /**
     * The mode of OOS Execution. Valid: `Automatic`, `Debug`.
     * 
     */
    @Import(name="mode")
    private @Nullable String mode;

    /**
     * @return The mode of OOS Execution. Valid: `Automatic`, `Debug`.
     * 
     */
    public Optional<String> mode() {
        return Optional.ofNullable(this.mode);
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
     * The id of parent OOS Execution.
     * 
     */
    @Import(name="parentExecutionId")
    private @Nullable String parentExecutionId;

    /**
     * @return The id of parent OOS Execution.
     * 
     */
    public Optional<String> parentExecutionId() {
        return Optional.ofNullable(this.parentExecutionId);
    }

    /**
     * The role that executes the current template.
     * 
     */
    @Import(name="ramRole")
    private @Nullable String ramRole;

    /**
     * @return The role that executes the current template.
     * 
     */
    public Optional<String> ramRole() {
        return Optional.ofNullable(this.ramRole);
    }

    /**
     * The sort field.
     * 
     */
    @Import(name="sortField")
    private @Nullable String sortField;

    /**
     * @return The sort field.
     * 
     */
    public Optional<String> sortField() {
        return Optional.ofNullable(this.sortField);
    }

    /**
     * The sort order.
     * 
     */
    @Import(name="sortOrder")
    private @Nullable String sortOrder;

    /**
     * @return The sort order.
     * 
     */
    public Optional<String> sortOrder() {
        return Optional.ofNullable(this.sortOrder);
    }

    /**
     * The execution whose start time is greater than or equal to the specified time.
     * 
     */
    @Import(name="startDateAfter")
    private @Nullable String startDateAfter;

    /**
     * @return The execution whose start time is greater than or equal to the specified time.
     * 
     */
    public Optional<String> startDateAfter() {
        return Optional.ofNullable(this.startDateAfter);
    }

    /**
     * The execution with start time less than or equal to the specified time.
     * 
     */
    @Import(name="startDateBefore")
    private @Nullable String startDateBefore;

    /**
     * @return The execution with start time less than or equal to the specified time.
     * 
     */
    public Optional<String> startDateBefore() {
        return Optional.ofNullable(this.startDateBefore);
    }

    /**
     * The Status of OOS Execution. Valid: `Cancelled`, `Failed`, `Queued`, `Running`, `Started`, `Success`, `Waiting`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The Status of OOS Execution. Valid: `Cancelled`, `Failed`, `Queued`, `Running`, `Started`, `Success`, `Waiting`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
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

    /**
     * The name of execution template.
     * 
     */
    @Import(name="templateName")
    private @Nullable String templateName;

    /**
     * @return The name of execution template.
     * 
     */
    public Optional<String> templateName() {
        return Optional.ofNullable(this.templateName);
    }

    private GetExecutionsPlainArgs() {}

    private GetExecutionsPlainArgs(GetExecutionsPlainArgs $) {
        this.category = $.category;
        this.endDate = $.endDate;
        this.endDateAfter = $.endDateAfter;
        this.executedBy = $.executedBy;
        this.ids = $.ids;
        this.includeChildExecution = $.includeChildExecution;
        this.mode = $.mode;
        this.outputFile = $.outputFile;
        this.parentExecutionId = $.parentExecutionId;
        this.ramRole = $.ramRole;
        this.sortField = $.sortField;
        this.sortOrder = $.sortOrder;
        this.startDateAfter = $.startDateAfter;
        this.startDateBefore = $.startDateBefore;
        this.status = $.status;
        this.tags = $.tags;
        this.templateName = $.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetExecutionsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetExecutionsPlainArgs $;

        public Builder() {
            $ = new GetExecutionsPlainArgs();
        }

        public Builder(GetExecutionsPlainArgs defaults) {
            $ = new GetExecutionsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param category The category of template. Valid: `AlarmTrigger`, `EventTrigger`, `Other` and `TimerTrigger`.
         * 
         * @return builder
         * 
         */
        public Builder category(@Nullable String category) {
            $.category = category;
            return this;
        }

        /**
         * @param endDate The time when the execution was ended.
         * 
         * @return builder
         * 
         */
        public Builder endDate(@Nullable String endDate) {
            $.endDate = endDate;
            return this;
        }

        /**
         * @param endDateAfter Execution whose end time is less than or equal to the specified time.
         * 
         * @return builder
         * 
         */
        public Builder endDateAfter(@Nullable String endDateAfter) {
            $.endDateAfter = endDateAfter;
            return this;
        }

        /**
         * @param executedBy The user who execute the template.
         * 
         * @return builder
         * 
         */
        public Builder executedBy(@Nullable String executedBy) {
            $.executedBy = executedBy;
            return this;
        }

        /**
         * @param ids A list of OOS Execution ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of OOS Execution ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param includeChildExecution Whether to include sub-execution.
         * 
         * @return builder
         * 
         */
        public Builder includeChildExecution(@Nullable Boolean includeChildExecution) {
            $.includeChildExecution = includeChildExecution;
            return this;
        }

        /**
         * @param mode The mode of OOS Execution. Valid: `Automatic`, `Debug`.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable String mode) {
            $.mode = mode;
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
         * @param parentExecutionId The id of parent OOS Execution.
         * 
         * @return builder
         * 
         */
        public Builder parentExecutionId(@Nullable String parentExecutionId) {
            $.parentExecutionId = parentExecutionId;
            return this;
        }

        /**
         * @param ramRole The role that executes the current template.
         * 
         * @return builder
         * 
         */
        public Builder ramRole(@Nullable String ramRole) {
            $.ramRole = ramRole;
            return this;
        }

        /**
         * @param sortField The sort field.
         * 
         * @return builder
         * 
         */
        public Builder sortField(@Nullable String sortField) {
            $.sortField = sortField;
            return this;
        }

        /**
         * @param sortOrder The sort order.
         * 
         * @return builder
         * 
         */
        public Builder sortOrder(@Nullable String sortOrder) {
            $.sortOrder = sortOrder;
            return this;
        }

        /**
         * @param startDateAfter The execution whose start time is greater than or equal to the specified time.
         * 
         * @return builder
         * 
         */
        public Builder startDateAfter(@Nullable String startDateAfter) {
            $.startDateAfter = startDateAfter;
            return this;
        }

        /**
         * @param startDateBefore The execution with start time less than or equal to the specified time.
         * 
         * @return builder
         * 
         */
        public Builder startDateBefore(@Nullable String startDateBefore) {
            $.startDateBefore = startDateBefore;
            return this;
        }

        /**
         * @param status The Status of OOS Execution. Valid: `Cancelled`, `Failed`, `Queued`, `Running`, `Started`, `Success`, `Waiting`.
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
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,String> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param templateName The name of execution template.
         * 
         * @return builder
         * 
         */
        public Builder templateName(@Nullable String templateName) {
            $.templateName = templateName;
            return this;
        }

        public GetExecutionsPlainArgs build() {
            return $;
        }
    }

}
