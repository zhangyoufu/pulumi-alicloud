// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.quotas.outputs;

import com.pulumi.alicloud.quotas.outputs.TemplateApplicationsQuotaApplicationDetailPeriod;
import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TemplateApplicationsQuotaApplicationDetail {
    /**
     * @return Alibaba Cloud account (primary account).
     * 
     */
    private @Nullable String aliyunUid;
    /**
     * @return The ID of the quota promotion request.
     * 
     */
    private @Nullable String applicationId;
    /**
     * @return The approved quota value of the quota increase request.
     * 
     */
    private @Nullable Double approveValue;
    /**
     * @return Approval comments on quota increase applications.
     * 
     */
    private @Nullable String auditReason;
    /**
     * @return Quota dimension. See `dimensions` below.
     * 
     */
    private @Nullable Map<String,String> dimensions;
    /**
     * @return The language of the quota application result notification. Value:
     * - zh (default): Chinese.
     * - en: English.
     * 
     */
    private @Nullable String envLanguage;
    /**
     * @return Whether to send notification of quota application result. Value:
     * - 0 (default): No.
     * - 3: Yes.
     * 
     */
    private @Nullable Integer noticeType;
    /**
     * @return Quota calculation period.
     * 
     */
    private @Nullable TemplateApplicationsQuotaApplicationDetailPeriod period;
    /**
     * @return Quota ARN.
     * 
     */
    private @Nullable String quotaArn;
    /**
     * @return The quota description.
     * 
     */
    private @Nullable String quotaDescription;
    /**
     * @return The quota name.
     * 
     */
    private @Nullable String quotaName;
    /**
     * @return Quota unit.
     * 
     */
    private @Nullable String quotaUnit;
    /**
     * @return Reason for quota application.
     * &gt; **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
     * 
     */
    private @Nullable String reason;
    /**
     * @return The approval status of the quota promotion application. Value:
     * - Disagree: reject.
     * - Approve: approved.
     * - Process: under review.
     * - Cancel: Closed.
     * 
     */
    private @Nullable String status;

    private TemplateApplicationsQuotaApplicationDetail() {}
    /**
     * @return Alibaba Cloud account (primary account).
     * 
     */
    public Optional<String> aliyunUid() {
        return Optional.ofNullable(this.aliyunUid);
    }
    /**
     * @return The ID of the quota promotion request.
     * 
     */
    public Optional<String> applicationId() {
        return Optional.ofNullable(this.applicationId);
    }
    /**
     * @return The approved quota value of the quota increase request.
     * 
     */
    public Optional<Double> approveValue() {
        return Optional.ofNullable(this.approveValue);
    }
    /**
     * @return Approval comments on quota increase applications.
     * 
     */
    public Optional<String> auditReason() {
        return Optional.ofNullable(this.auditReason);
    }
    /**
     * @return Quota dimension. See `dimensions` below.
     * 
     */
    public Map<String,String> dimensions() {
        return this.dimensions == null ? Map.of() : this.dimensions;
    }
    /**
     * @return The language of the quota application result notification. Value:
     * - zh (default): Chinese.
     * - en: English.
     * 
     */
    public Optional<String> envLanguage() {
        return Optional.ofNullable(this.envLanguage);
    }
    /**
     * @return Whether to send notification of quota application result. Value:
     * - 0 (default): No.
     * - 3: Yes.
     * 
     */
    public Optional<Integer> noticeType() {
        return Optional.ofNullable(this.noticeType);
    }
    /**
     * @return Quota calculation period.
     * 
     */
    public Optional<TemplateApplicationsQuotaApplicationDetailPeriod> period() {
        return Optional.ofNullable(this.period);
    }
    /**
     * @return Quota ARN.
     * 
     */
    public Optional<String> quotaArn() {
        return Optional.ofNullable(this.quotaArn);
    }
    /**
     * @return The quota description.
     * 
     */
    public Optional<String> quotaDescription() {
        return Optional.ofNullable(this.quotaDescription);
    }
    /**
     * @return The quota name.
     * 
     */
    public Optional<String> quotaName() {
        return Optional.ofNullable(this.quotaName);
    }
    /**
     * @return Quota unit.
     * 
     */
    public Optional<String> quotaUnit() {
        return Optional.ofNullable(this.quotaUnit);
    }
    /**
     * @return Reason for quota application.
     * &gt; **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
     * 
     */
    public Optional<String> reason() {
        return Optional.ofNullable(this.reason);
    }
    /**
     * @return The approval status of the quota promotion application. Value:
     * - Disagree: reject.
     * - Approve: approved.
     * - Process: under review.
     * - Cancel: Closed.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TemplateApplicationsQuotaApplicationDetail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String aliyunUid;
        private @Nullable String applicationId;
        private @Nullable Double approveValue;
        private @Nullable String auditReason;
        private @Nullable Map<String,String> dimensions;
        private @Nullable String envLanguage;
        private @Nullable Integer noticeType;
        private @Nullable TemplateApplicationsQuotaApplicationDetailPeriod period;
        private @Nullable String quotaArn;
        private @Nullable String quotaDescription;
        private @Nullable String quotaName;
        private @Nullable String quotaUnit;
        private @Nullable String reason;
        private @Nullable String status;
        public Builder() {}
        public Builder(TemplateApplicationsQuotaApplicationDetail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aliyunUid = defaults.aliyunUid;
    	      this.applicationId = defaults.applicationId;
    	      this.approveValue = defaults.approveValue;
    	      this.auditReason = defaults.auditReason;
    	      this.dimensions = defaults.dimensions;
    	      this.envLanguage = defaults.envLanguage;
    	      this.noticeType = defaults.noticeType;
    	      this.period = defaults.period;
    	      this.quotaArn = defaults.quotaArn;
    	      this.quotaDescription = defaults.quotaDescription;
    	      this.quotaName = defaults.quotaName;
    	      this.quotaUnit = defaults.quotaUnit;
    	      this.reason = defaults.reason;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder aliyunUid(@Nullable String aliyunUid) {

            this.aliyunUid = aliyunUid;
            return this;
        }
        @CustomType.Setter
        public Builder applicationId(@Nullable String applicationId) {

            this.applicationId = applicationId;
            return this;
        }
        @CustomType.Setter
        public Builder approveValue(@Nullable Double approveValue) {

            this.approveValue = approveValue;
            return this;
        }
        @CustomType.Setter
        public Builder auditReason(@Nullable String auditReason) {

            this.auditReason = auditReason;
            return this;
        }
        @CustomType.Setter
        public Builder dimensions(@Nullable Map<String,String> dimensions) {

            this.dimensions = dimensions;
            return this;
        }
        @CustomType.Setter
        public Builder envLanguage(@Nullable String envLanguage) {

            this.envLanguage = envLanguage;
            return this;
        }
        @CustomType.Setter
        public Builder noticeType(@Nullable Integer noticeType) {

            this.noticeType = noticeType;
            return this;
        }
        @CustomType.Setter
        public Builder period(@Nullable TemplateApplicationsQuotaApplicationDetailPeriod period) {

            this.period = period;
            return this;
        }
        @CustomType.Setter
        public Builder quotaArn(@Nullable String quotaArn) {

            this.quotaArn = quotaArn;
            return this;
        }
        @CustomType.Setter
        public Builder quotaDescription(@Nullable String quotaDescription) {

            this.quotaDescription = quotaDescription;
            return this;
        }
        @CustomType.Setter
        public Builder quotaName(@Nullable String quotaName) {

            this.quotaName = quotaName;
            return this;
        }
        @CustomType.Setter
        public Builder quotaUnit(@Nullable String quotaUnit) {

            this.quotaUnit = quotaUnit;
            return this;
        }
        @CustomType.Setter
        public Builder reason(@Nullable String reason) {

            this.reason = reason;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {

            this.status = status;
            return this;
        }
        public TemplateApplicationsQuotaApplicationDetail build() {
            final var _resultValue = new TemplateApplicationsQuotaApplicationDetail();
            _resultValue.aliyunUid = aliyunUid;
            _resultValue.applicationId = applicationId;
            _resultValue.approveValue = approveValue;
            _resultValue.auditReason = auditReason;
            _resultValue.dimensions = dimensions;
            _resultValue.envLanguage = envLanguage;
            _resultValue.noticeType = noticeType;
            _resultValue.period = period;
            _resultValue.quotaArn = quotaArn;
            _resultValue.quotaDescription = quotaDescription;
            _resultValue.quotaName = quotaName;
            _resultValue.quotaUnit = quotaUnit;
            _resultValue.reason = reason;
            _resultValue.status = status;
            return _resultValue;
        }
    }
}
