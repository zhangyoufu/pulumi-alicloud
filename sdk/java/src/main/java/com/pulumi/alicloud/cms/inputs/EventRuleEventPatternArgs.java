// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EventRuleEventPatternArgs extends com.pulumi.resources.ResourceArgs {

    public static final EventRuleEventPatternArgs Empty = new EventRuleEventPatternArgs();

    /**
     * The type of the event-triggered alert rule. Valid values:
     * 
     */
    @Import(name="eventTypeLists")
    private @Nullable Output<List<String>> eventTypeLists;

    /**
     * @return The type of the event-triggered alert rule. Valid values:
     * 
     */
    public Optional<Output<List<String>>> eventTypeLists() {
        return Optional.ofNullable(this.eventTypeLists);
    }

    /**
     * The level of the event-triggered alert rule. Valid values:
     * 
     */
    @Import(name="levelLists")
    private @Nullable Output<List<String>> levelLists;

    /**
     * @return The level of the event-triggered alert rule. Valid values:
     * 
     */
    public Optional<Output<List<String>>> levelLists() {
        return Optional.ofNullable(this.levelLists);
    }

    /**
     * The name of the event-triggered alert rule.
     * 
     */
    @Import(name="nameLists")
    private @Nullable Output<List<String>> nameLists;

    /**
     * @return The name of the event-triggered alert rule.
     * 
     */
    public Optional<Output<List<String>>> nameLists() {
        return Optional.ofNullable(this.nameLists);
    }

    /**
     * The type of the cloud service.
     * 
     */
    @Import(name="product", required=true)
    private Output<String> product;

    /**
     * @return The type of the cloud service.
     * 
     */
    public Output<String> product() {
        return this.product;
    }

    /**
     * The SQL condition that is used to filter events. If the content of an event meets the specified SQL condition, an alert is automatically triggered.
     * 
     */
    @Import(name="sqlFilter")
    private @Nullable Output<String> sqlFilter;

    /**
     * @return The SQL condition that is used to filter events. If the content of an event meets the specified SQL condition, an alert is automatically triggered.
     * 
     */
    public Optional<Output<String>> sqlFilter() {
        return Optional.ofNullable(this.sqlFilter);
    }

    private EventRuleEventPatternArgs() {}

    private EventRuleEventPatternArgs(EventRuleEventPatternArgs $) {
        this.eventTypeLists = $.eventTypeLists;
        this.levelLists = $.levelLists;
        this.nameLists = $.nameLists;
        this.product = $.product;
        this.sqlFilter = $.sqlFilter;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EventRuleEventPatternArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EventRuleEventPatternArgs $;

        public Builder() {
            $ = new EventRuleEventPatternArgs();
        }

        public Builder(EventRuleEventPatternArgs defaults) {
            $ = new EventRuleEventPatternArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param eventTypeLists The type of the event-triggered alert rule. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder eventTypeLists(@Nullable Output<List<String>> eventTypeLists) {
            $.eventTypeLists = eventTypeLists;
            return this;
        }

        /**
         * @param eventTypeLists The type of the event-triggered alert rule. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder eventTypeLists(List<String> eventTypeLists) {
            return eventTypeLists(Output.of(eventTypeLists));
        }

        /**
         * @param eventTypeLists The type of the event-triggered alert rule. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder eventTypeLists(String... eventTypeLists) {
            return eventTypeLists(List.of(eventTypeLists));
        }

        /**
         * @param levelLists The level of the event-triggered alert rule. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder levelLists(@Nullable Output<List<String>> levelLists) {
            $.levelLists = levelLists;
            return this;
        }

        /**
         * @param levelLists The level of the event-triggered alert rule. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder levelLists(List<String> levelLists) {
            return levelLists(Output.of(levelLists));
        }

        /**
         * @param levelLists The level of the event-triggered alert rule. Valid values:
         * 
         * @return builder
         * 
         */
        public Builder levelLists(String... levelLists) {
            return levelLists(List.of(levelLists));
        }

        /**
         * @param nameLists The name of the event-triggered alert rule.
         * 
         * @return builder
         * 
         */
        public Builder nameLists(@Nullable Output<List<String>> nameLists) {
            $.nameLists = nameLists;
            return this;
        }

        /**
         * @param nameLists The name of the event-triggered alert rule.
         * 
         * @return builder
         * 
         */
        public Builder nameLists(List<String> nameLists) {
            return nameLists(Output.of(nameLists));
        }

        /**
         * @param nameLists The name of the event-triggered alert rule.
         * 
         * @return builder
         * 
         */
        public Builder nameLists(String... nameLists) {
            return nameLists(List.of(nameLists));
        }

        /**
         * @param product The type of the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder product(Output<String> product) {
            $.product = product;
            return this;
        }

        /**
         * @param product The type of the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder product(String product) {
            return product(Output.of(product));
        }

        /**
         * @param sqlFilter The SQL condition that is used to filter events. If the content of an event meets the specified SQL condition, an alert is automatically triggered.
         * 
         * @return builder
         * 
         */
        public Builder sqlFilter(@Nullable Output<String> sqlFilter) {
            $.sqlFilter = sqlFilter;
            return this;
        }

        /**
         * @param sqlFilter The SQL condition that is used to filter events. If the content of an event meets the specified SQL condition, an alert is automatically triggered.
         * 
         * @return builder
         * 
         */
        public Builder sqlFilter(String sqlFilter) {
            return sqlFilter(Output.of(sqlFilter));
        }

        public EventRuleEventPatternArgs build() {
            $.product = Objects.requireNonNull($.product, "expected parameter 'product' to be non-null");
            return $;
        }
    }

}
