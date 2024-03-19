// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge.inputs;

import com.pulumi.alicloud.eventbridge.inputs.RuleTargetArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleState extends com.pulumi.resources.ResourceArgs {

    public static final RuleState Empty = new RuleState();

    /**
     * The description of the event rule.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the event rule.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the event bus.
     * 
     */
    @Import(name="eventBusName")
    private @Nullable Output<String> eventBusName;

    /**
     * @return The name of the event bus.
     * 
     */
    public Optional<Output<String>> eventBusName() {
        return Optional.ofNullable(this.eventBusName);
    }

    /**
     * The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
     * 
     */
    @Import(name="filterPattern")
    private @Nullable Output<String> filterPattern;

    /**
     * @return The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
     * 
     */
    public Optional<Output<String>> filterPattern() {
        return Optional.ofNullable(this.filterPattern);
    }

    /**
     * The name of the event rule.
     * 
     */
    @Import(name="ruleName")
    private @Nullable Output<String> ruleName;

    /**
     * @return The name of the event rule.
     * 
     */
    public Optional<Output<String>> ruleName() {
        return Optional.ofNullable(this.ruleName);
    }

    /**
     * The status of the event rule. Valid values: `ENABLE`, `DISABLE`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the event rule. Valid values: `ENABLE`, `DISABLE`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The targets of rule. See `targets` below.
     * 
     */
    @Import(name="targets")
    private @Nullable Output<List<RuleTargetArgs>> targets;

    /**
     * @return The targets of rule. See `targets` below.
     * 
     */
    public Optional<Output<List<RuleTargetArgs>>> targets() {
        return Optional.ofNullable(this.targets);
    }

    private RuleState() {}

    private RuleState(RuleState $) {
        this.description = $.description;
        this.eventBusName = $.eventBusName;
        this.filterPattern = $.filterPattern;
        this.ruleName = $.ruleName;
        this.status = $.status;
        this.targets = $.targets;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleState $;

        public Builder() {
            $ = new RuleState();
        }

        public Builder(RuleState defaults) {
            $ = new RuleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the event rule.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the event rule.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param eventBusName The name of the event bus.
         * 
         * @return builder
         * 
         */
        public Builder eventBusName(@Nullable Output<String> eventBusName) {
            $.eventBusName = eventBusName;
            return this;
        }

        /**
         * @param eventBusName The name of the event bus.
         * 
         * @return builder
         * 
         */
        public Builder eventBusName(String eventBusName) {
            return eventBusName(Output.of(eventBusName));
        }

        /**
         * @param filterPattern The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
         * 
         * @return builder
         * 
         */
        public Builder filterPattern(@Nullable Output<String> filterPattern) {
            $.filterPattern = filterPattern;
            return this;
        }

        /**
         * @param filterPattern The pattern to match interested events. Event mode, JSON format. The value description is as follows: `stringEqual` mode. `stringExpression` mode. Each field has up to 5 expressions (map structure).
         * 
         * @return builder
         * 
         */
        public Builder filterPattern(String filterPattern) {
            return filterPattern(Output.of(filterPattern));
        }

        /**
         * @param ruleName The name of the event rule.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(@Nullable Output<String> ruleName) {
            $.ruleName = ruleName;
            return this;
        }

        /**
         * @param ruleName The name of the event rule.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(String ruleName) {
            return ruleName(Output.of(ruleName));
        }

        /**
         * @param status The status of the event rule. Valid values: `ENABLE`, `DISABLE`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the event rule. Valid values: `ENABLE`, `DISABLE`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param targets The targets of rule. See `targets` below.
         * 
         * @return builder
         * 
         */
        public Builder targets(@Nullable Output<List<RuleTargetArgs>> targets) {
            $.targets = targets;
            return this;
        }

        /**
         * @param targets The targets of rule. See `targets` below.
         * 
         * @return builder
         * 
         */
        public Builder targets(List<RuleTargetArgs> targets) {
            return targets(Output.of(targets));
        }

        /**
         * @param targets The targets of rule. See `targets` below.
         * 
         * @return builder
         * 
         */
        public Builder targets(RuleTargetArgs... targets) {
            return targets(List.of(targets));
        }

        public RuleState build() {
            return $;
        }
    }

}
