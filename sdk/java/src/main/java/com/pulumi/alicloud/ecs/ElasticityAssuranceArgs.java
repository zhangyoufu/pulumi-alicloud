// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ElasticityAssuranceArgs extends com.pulumi.resources.ResourceArgs {

    public static final ElasticityAssuranceArgs Empty = new ElasticityAssuranceArgs();

    /**
     * The total number of times that the elasticity assurance can be applied. Set the value to Unlimited. This value indicates that the elasticity assurance can be applied an unlimited number of times within its effective duration. Default value: Unlimited.
     * 
     */
    @Import(name="assuranceTimes")
    private @Nullable Output<String> assuranceTimes;

    /**
     * @return The total number of times that the elasticity assurance can be applied. Set the value to Unlimited. This value indicates that the elasticity assurance can be applied an unlimited number of times within its effective duration. Default value: Unlimited.
     * 
     */
    public Optional<Output<String>> assuranceTimes() {
        return Optional.ofNullable(this.assuranceTimes);
    }

    /**
     * Description of flexible guarantee service.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of flexible guarantee service.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The total number of instances for which to reserve the capacity of an instance type. Valid values: 1 to 1000.
     * 
     */
    @Import(name="instanceAmount", required=true)
    private Output<Integer> instanceAmount;

    /**
     * @return The total number of instances for which to reserve the capacity of an instance type. Valid values: 1 to 1000.
     * 
     */
    public Output<Integer> instanceAmount() {
        return this.instanceAmount;
    }

    /**
     * Instance type. Currently, only one instance type is supported.
     * 
     */
    @Import(name="instanceType", required=true)
    private Output<String> instanceType;

    /**
     * @return Instance type. Currently, only one instance type is supported.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }

    /**
     * Length of purchase. The unit of duration is determined by the &#39;period_unit&#39; parameter. Default value: 1.
     * - When the `period_unit` parameter is set to Month, the valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
     * - When the `period_unit` parameter is set to Year, the valid values are 1, 2, 3, 4, and 5.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return Length of purchase. The unit of duration is determined by the &#39;period_unit&#39; parameter. Default value: 1.
     * - When the `period_unit` parameter is set to Month, the valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
     * - When the `period_unit` parameter is set to Year, the valid values are 1, 2, 3, 4, and 5.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * Duration unit. Value range:-Month: Month-Year: YearDefault value: Year
     * 
     */
    @Import(name="periodUnit")
    private @Nullable Output<String> periodUnit;

    /**
     * @return Duration unit. Value range:-Month: Month-Year: YearDefault value: Year
     * 
     */
    public Optional<Output<String>> periodUnit() {
        return Optional.ofNullable(this.periodUnit);
    }

    /**
     * The matching mode of flexible guarantee service. Possible values:-Open: flexible guarantee service for Open mode.-Target: specifies the flexible guarantee service of the mode.
     * 
     */
    @Import(name="privatePoolOptionsMatchCriteria")
    private @Nullable Output<String> privatePoolOptionsMatchCriteria;

    /**
     * @return The matching mode of flexible guarantee service. Possible values:-Open: flexible guarantee service for Open mode.-Target: specifies the flexible guarantee service of the mode.
     * 
     */
    public Optional<Output<String>> privatePoolOptionsMatchCriteria() {
        return Optional.ofNullable(this.privatePoolOptionsMatchCriteria);
    }

    /**
     * The name of the flexible protection service.
     * 
     */
    @Import(name="privatePoolOptionsName")
    private @Nullable Output<String> privatePoolOptionsName;

    /**
     * @return The name of the flexible protection service.
     * 
     */
    public Optional<Output<String>> privatePoolOptionsName() {
        return Optional.ofNullable(this.privatePoolOptionsName);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * Flexible guarantee service effective time.
     * 
     */
    @Import(name="startTime")
    private @Nullable Output<String> startTime;

    /**
     * @return Flexible guarantee service effective time.
     * 
     */
    public Optional<Output<String>> startTime() {
        return Optional.ofNullable(this.startTime);
    }

    /**
     * The tag key-value pair information bound by the elastic guarantee service.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return The tag key-value pair information bound by the elastic guarantee service.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The zone ID of the region to which the elastic Protection Service belongs. Currently, only the creation of flexible protection services in one available area is supported.
     * 
     */
    @Import(name="zoneIds", required=true)
    private Output<List<String>> zoneIds;

    /**
     * @return The zone ID of the region to which the elastic Protection Service belongs. Currently, only the creation of flexible protection services in one available area is supported.
     * 
     */
    public Output<List<String>> zoneIds() {
        return this.zoneIds;
    }

    private ElasticityAssuranceArgs() {}

    private ElasticityAssuranceArgs(ElasticityAssuranceArgs $) {
        this.assuranceTimes = $.assuranceTimes;
        this.description = $.description;
        this.instanceAmount = $.instanceAmount;
        this.instanceType = $.instanceType;
        this.period = $.period;
        this.periodUnit = $.periodUnit;
        this.privatePoolOptionsMatchCriteria = $.privatePoolOptionsMatchCriteria;
        this.privatePoolOptionsName = $.privatePoolOptionsName;
        this.resourceGroupId = $.resourceGroupId;
        this.startTime = $.startTime;
        this.tags = $.tags;
        this.zoneIds = $.zoneIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ElasticityAssuranceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ElasticityAssuranceArgs $;

        public Builder() {
            $ = new ElasticityAssuranceArgs();
        }

        public Builder(ElasticityAssuranceArgs defaults) {
            $ = new ElasticityAssuranceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param assuranceTimes The total number of times that the elasticity assurance can be applied. Set the value to Unlimited. This value indicates that the elasticity assurance can be applied an unlimited number of times within its effective duration. Default value: Unlimited.
         * 
         * @return builder
         * 
         */
        public Builder assuranceTimes(@Nullable Output<String> assuranceTimes) {
            $.assuranceTimes = assuranceTimes;
            return this;
        }

        /**
         * @param assuranceTimes The total number of times that the elasticity assurance can be applied. Set the value to Unlimited. This value indicates that the elasticity assurance can be applied an unlimited number of times within its effective duration. Default value: Unlimited.
         * 
         * @return builder
         * 
         */
        public Builder assuranceTimes(String assuranceTimes) {
            return assuranceTimes(Output.of(assuranceTimes));
        }

        /**
         * @param description Description of flexible guarantee service.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of flexible guarantee service.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param instanceAmount The total number of instances for which to reserve the capacity of an instance type. Valid values: 1 to 1000.
         * 
         * @return builder
         * 
         */
        public Builder instanceAmount(Output<Integer> instanceAmount) {
            $.instanceAmount = instanceAmount;
            return this;
        }

        /**
         * @param instanceAmount The total number of instances for which to reserve the capacity of an instance type. Valid values: 1 to 1000.
         * 
         * @return builder
         * 
         */
        public Builder instanceAmount(Integer instanceAmount) {
            return instanceAmount(Output.of(instanceAmount));
        }

        /**
         * @param instanceType Instance type. Currently, only one instance type is supported.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType Instance type. Currently, only one instance type is supported.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param period Length of purchase. The unit of duration is determined by the &#39;period_unit&#39; parameter. Default value: 1.
         * - When the `period_unit` parameter is set to Month, the valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
         * - When the `period_unit` parameter is set to Year, the valid values are 1, 2, 3, 4, and 5.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period Length of purchase. The unit of duration is determined by the &#39;period_unit&#39; parameter. Default value: 1.
         * - When the `period_unit` parameter is set to Month, the valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
         * - When the `period_unit` parameter is set to Year, the valid values are 1, 2, 3, 4, and 5.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param periodUnit Duration unit. Value range:-Month: Month-Year: YearDefault value: Year
         * 
         * @return builder
         * 
         */
        public Builder periodUnit(@Nullable Output<String> periodUnit) {
            $.periodUnit = periodUnit;
            return this;
        }

        /**
         * @param periodUnit Duration unit. Value range:-Month: Month-Year: YearDefault value: Year
         * 
         * @return builder
         * 
         */
        public Builder periodUnit(String periodUnit) {
            return periodUnit(Output.of(periodUnit));
        }

        /**
         * @param privatePoolOptionsMatchCriteria The matching mode of flexible guarantee service. Possible values:-Open: flexible guarantee service for Open mode.-Target: specifies the flexible guarantee service of the mode.
         * 
         * @return builder
         * 
         */
        public Builder privatePoolOptionsMatchCriteria(@Nullable Output<String> privatePoolOptionsMatchCriteria) {
            $.privatePoolOptionsMatchCriteria = privatePoolOptionsMatchCriteria;
            return this;
        }

        /**
         * @param privatePoolOptionsMatchCriteria The matching mode of flexible guarantee service. Possible values:-Open: flexible guarantee service for Open mode.-Target: specifies the flexible guarantee service of the mode.
         * 
         * @return builder
         * 
         */
        public Builder privatePoolOptionsMatchCriteria(String privatePoolOptionsMatchCriteria) {
            return privatePoolOptionsMatchCriteria(Output.of(privatePoolOptionsMatchCriteria));
        }

        /**
         * @param privatePoolOptionsName The name of the flexible protection service.
         * 
         * @return builder
         * 
         */
        public Builder privatePoolOptionsName(@Nullable Output<String> privatePoolOptionsName) {
            $.privatePoolOptionsName = privatePoolOptionsName;
            return this;
        }

        /**
         * @param privatePoolOptionsName The name of the flexible protection service.
         * 
         * @return builder
         * 
         */
        public Builder privatePoolOptionsName(String privatePoolOptionsName) {
            return privatePoolOptionsName(Output.of(privatePoolOptionsName));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param startTime Flexible guarantee service effective time.
         * 
         * @return builder
         * 
         */
        public Builder startTime(@Nullable Output<String> startTime) {
            $.startTime = startTime;
            return this;
        }

        /**
         * @param startTime Flexible guarantee service effective time.
         * 
         * @return builder
         * 
         */
        public Builder startTime(String startTime) {
            return startTime(Output.of(startTime));
        }

        /**
         * @param tags The tag key-value pair information bound by the elastic guarantee service.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tag key-value pair information bound by the elastic guarantee service.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param zoneIds The zone ID of the region to which the elastic Protection Service belongs. Currently, only the creation of flexible protection services in one available area is supported.
         * 
         * @return builder
         * 
         */
        public Builder zoneIds(Output<List<String>> zoneIds) {
            $.zoneIds = zoneIds;
            return this;
        }

        /**
         * @param zoneIds The zone ID of the region to which the elastic Protection Service belongs. Currently, only the creation of flexible protection services in one available area is supported.
         * 
         * @return builder
         * 
         */
        public Builder zoneIds(List<String> zoneIds) {
            return zoneIds(Output.of(zoneIds));
        }

        /**
         * @param zoneIds The zone ID of the region to which the elastic Protection Service belongs. Currently, only the creation of flexible protection services in one available area is supported.
         * 
         * @return builder
         * 
         */
        public Builder zoneIds(String... zoneIds) {
            return zoneIds(List.of(zoneIds));
        }

        public ElasticityAssuranceArgs build() {
            $.instanceAmount = Objects.requireNonNull($.instanceAmount, "expected parameter 'instanceAmount' to be non-null");
            $.instanceType = Objects.requireNonNull($.instanceType, "expected parameter 'instanceType' to be non-null");
            $.zoneIds = Objects.requireNonNull($.zoneIds, "expected parameter 'zoneIds' to be non-null");
            return $;
        }
    }

}
