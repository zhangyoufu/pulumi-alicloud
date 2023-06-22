// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds.inputs;

import com.pulumi.alicloud.eds.inputs.EcdPolicyGroupAuthorizeAccessPolicyRuleArgs;
import com.pulumi.alicloud.eds.inputs.EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EcdPolicyGroupState extends com.pulumi.resources.ResourceArgs {

    public static final EcdPolicyGroupState Empty = new EcdPolicyGroupState();

    /**
     * The rule of authorize access rule. See `authorize_access_policy_rules` below.
     * 
     */
    @Import(name="authorizeAccessPolicyRules")
    private @Nullable Output<List<EcdPolicyGroupAuthorizeAccessPolicyRuleArgs>> authorizeAccessPolicyRules;

    /**
     * @return The rule of authorize access rule. See `authorize_access_policy_rules` below.
     * 
     */
    public Optional<Output<List<EcdPolicyGroupAuthorizeAccessPolicyRuleArgs>>> authorizeAccessPolicyRules() {
        return Optional.ofNullable(this.authorizeAccessPolicyRules);
    }

    /**
     * The policy rule. See `authorize_security_policy_rules` below.
     * 
     */
    @Import(name="authorizeSecurityPolicyRules")
    private @Nullable Output<List<EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs>> authorizeSecurityPolicyRules;

    /**
     * @return The policy rule. See `authorize_security_policy_rules` below.
     * 
     */
    public Optional<Output<List<EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs>>> authorizeSecurityPolicyRules() {
        return Optional.ofNullable(this.authorizeSecurityPolicyRules);
    }

    /**
     * Whether to enable local camera redirection. Valid values: `on`, `off`.
     * 
     */
    @Import(name="cameraRedirect")
    private @Nullable Output<String> cameraRedirect;

    /**
     * @return Whether to enable local camera redirection. Valid values: `on`, `off`.
     * 
     */
    public Optional<Output<String>> cameraRedirect() {
        return Optional.ofNullable(this.cameraRedirect);
    }

    /**
     * The clipboard policy. Valid values: `off`, `read`, `readwrite`.
     * 
     */
    @Import(name="clipboard")
    private @Nullable Output<String> clipboard;

    /**
     * @return The clipboard policy. Valid values: `off`, `read`, `readwrite`.
     * 
     */
    public Optional<Output<String>> clipboard() {
        return Optional.ofNullable(this.clipboard);
    }

    /**
     * The list of domain.
     * 
     */
    @Import(name="domainList")
    private @Nullable Output<String> domainList;

    /**
     * @return The list of domain.
     * 
     */
    public Optional<Output<String>> domainList() {
        return Optional.ofNullable(this.domainList);
    }

    /**
     * The access of html5. Valid values: `off`, `on`.
     * 
     */
    @Import(name="htmlAccess")
    private @Nullable Output<String> htmlAccess;

    /**
     * @return The access of html5. Valid values: `off`, `on`.
     * 
     */
    public Optional<Output<String>> htmlAccess() {
        return Optional.ofNullable(this.htmlAccess);
    }

    /**
     * The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
     * 
     */
    @Import(name="htmlFileTransfer")
    private @Nullable Output<String> htmlFileTransfer;

    /**
     * @return The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
     * 
     */
    public Optional<Output<String>> htmlFileTransfer() {
        return Optional.ofNullable(this.htmlFileTransfer);
    }

    /**
     * Local drive redirect policy. Valid values: `  readwrite `, `off`, `read`.
     * 
     */
    @Import(name="localDrive")
    private @Nullable Output<String> localDrive;

    /**
     * @return Local drive redirect policy. Valid values: `  readwrite `, `off`, `read`.
     * 
     */
    public Optional<Output<String>> localDrive() {
        return Optional.ofNullable(this.localDrive);
    }

    /**
     * The name of policy group.
     * 
     */
    @Import(name="policyGroupName")
    private @Nullable Output<String> policyGroupName;

    /**
     * @return The name of policy group.
     * 
     */
    public Optional<Output<String>> policyGroupName() {
        return Optional.ofNullable(this.policyGroupName);
    }

    /**
     * Whether to enable screen recording. Valid values: `off`, `all-time`, `period`.
     * 
     */
    @Import(name="recording")
    private @Nullable Output<String> recording;

    /**
     * @return Whether to enable screen recording. Valid values: `off`, `all-time`, `period`.
     * 
     */
    public Optional<Output<String>> recording() {
        return Optional.ofNullable(this.recording);
    }

    /**
     * The end time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     * 
     */
    @Import(name="recordingEndTime")
    private @Nullable Output<String> recordingEndTime;

    /**
     * @return The end time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     * 
     */
    public Optional<Output<String>> recordingEndTime() {
        return Optional.ofNullable(this.recordingEndTime);
    }

    /**
     * The screen recording video retention. Valid values between 30 and 180. This return value is meaningful only when the value of `recording` is `period` or `all-time`.
     * 
     */
    @Import(name="recordingExpires")
    private @Nullable Output<Integer> recordingExpires;

    /**
     * @return The screen recording video retention. Valid values between 30 and 180. This return value is meaningful only when the value of `recording` is `period` or `all-time`.
     * 
     */
    public Optional<Output<Integer>> recordingExpires() {
        return Optional.ofNullable(this.recordingExpires);
    }

    /**
     * The fps of recording. Valid values: `2`, `5`, `10`, `15`.
     * 
     */
    @Import(name="recordingFps")
    private @Nullable Output<Integer> recordingFps;

    /**
     * @return The fps of recording. Valid values: `2`, `5`, `10`, `15`.
     * 
     */
    public Optional<Output<Integer>> recordingFps() {
        return Optional.ofNullable(this.recordingFps);
    }

    /**
     * The start time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     * 
     */
    @Import(name="recordingStartTime")
    private @Nullable Output<String> recordingStartTime;

    /**
     * @return The start time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     * 
     */
    public Optional<Output<String>> recordingStartTime() {
        return Optional.ofNullable(this.recordingStartTime);
    }

    /**
     * The status of policy.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of policy.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The usb redirect policy. Valid values: `off`, `on`.
     * 
     */
    @Import(name="usbRedirect")
    private @Nullable Output<String> usbRedirect;

    /**
     * @return The usb redirect policy. Valid values: `off`, `on`.
     * 
     */
    public Optional<Output<String>> usbRedirect() {
        return Optional.ofNullable(this.usbRedirect);
    }

    /**
     * The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
     * 
     */
    @Import(name="visualQuality")
    private @Nullable Output<String> visualQuality;

    /**
     * @return The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
     * 
     */
    public Optional<Output<String>> visualQuality() {
        return Optional.ofNullable(this.visualQuality);
    }

    /**
     * The watermark policy. Valid values: `off`, `on`.
     * 
     */
    @Import(name="watermark")
    private @Nullable Output<String> watermark;

    /**
     * @return The watermark policy. Valid values: `off`, `on`.
     * 
     */
    public Optional<Output<String>> watermark() {
        return Optional.ofNullable(this.watermark);
    }

    /**
     * The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
     * 
     */
    @Import(name="watermarkTransparency")
    private @Nullable Output<String> watermarkTransparency;

    /**
     * @return The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
     * 
     */
    public Optional<Output<String>> watermarkTransparency() {
        return Optional.ofNullable(this.watermarkTransparency);
    }

    /**
     * The type of watemark. Valid values: `EndUserId`, `HostName`.
     * 
     */
    @Import(name="watermarkType")
    private @Nullable Output<String> watermarkType;

    /**
     * @return The type of watemark. Valid values: `EndUserId`, `HostName`.
     * 
     */
    public Optional<Output<String>> watermarkType() {
        return Optional.ofNullable(this.watermarkType);
    }

    private EcdPolicyGroupState() {}

    private EcdPolicyGroupState(EcdPolicyGroupState $) {
        this.authorizeAccessPolicyRules = $.authorizeAccessPolicyRules;
        this.authorizeSecurityPolicyRules = $.authorizeSecurityPolicyRules;
        this.cameraRedirect = $.cameraRedirect;
        this.clipboard = $.clipboard;
        this.domainList = $.domainList;
        this.htmlAccess = $.htmlAccess;
        this.htmlFileTransfer = $.htmlFileTransfer;
        this.localDrive = $.localDrive;
        this.policyGroupName = $.policyGroupName;
        this.recording = $.recording;
        this.recordingEndTime = $.recordingEndTime;
        this.recordingExpires = $.recordingExpires;
        this.recordingFps = $.recordingFps;
        this.recordingStartTime = $.recordingStartTime;
        this.status = $.status;
        this.usbRedirect = $.usbRedirect;
        this.visualQuality = $.visualQuality;
        this.watermark = $.watermark;
        this.watermarkTransparency = $.watermarkTransparency;
        this.watermarkType = $.watermarkType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EcdPolicyGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EcdPolicyGroupState $;

        public Builder() {
            $ = new EcdPolicyGroupState();
        }

        public Builder(EcdPolicyGroupState defaults) {
            $ = new EcdPolicyGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorizeAccessPolicyRules The rule of authorize access rule. See `authorize_access_policy_rules` below.
         * 
         * @return builder
         * 
         */
        public Builder authorizeAccessPolicyRules(@Nullable Output<List<EcdPolicyGroupAuthorizeAccessPolicyRuleArgs>> authorizeAccessPolicyRules) {
            $.authorizeAccessPolicyRules = authorizeAccessPolicyRules;
            return this;
        }

        /**
         * @param authorizeAccessPolicyRules The rule of authorize access rule. See `authorize_access_policy_rules` below.
         * 
         * @return builder
         * 
         */
        public Builder authorizeAccessPolicyRules(List<EcdPolicyGroupAuthorizeAccessPolicyRuleArgs> authorizeAccessPolicyRules) {
            return authorizeAccessPolicyRules(Output.of(authorizeAccessPolicyRules));
        }

        /**
         * @param authorizeAccessPolicyRules The rule of authorize access rule. See `authorize_access_policy_rules` below.
         * 
         * @return builder
         * 
         */
        public Builder authorizeAccessPolicyRules(EcdPolicyGroupAuthorizeAccessPolicyRuleArgs... authorizeAccessPolicyRules) {
            return authorizeAccessPolicyRules(List.of(authorizeAccessPolicyRules));
        }

        /**
         * @param authorizeSecurityPolicyRules The policy rule. See `authorize_security_policy_rules` below.
         * 
         * @return builder
         * 
         */
        public Builder authorizeSecurityPolicyRules(@Nullable Output<List<EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs>> authorizeSecurityPolicyRules) {
            $.authorizeSecurityPolicyRules = authorizeSecurityPolicyRules;
            return this;
        }

        /**
         * @param authorizeSecurityPolicyRules The policy rule. See `authorize_security_policy_rules` below.
         * 
         * @return builder
         * 
         */
        public Builder authorizeSecurityPolicyRules(List<EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs> authorizeSecurityPolicyRules) {
            return authorizeSecurityPolicyRules(Output.of(authorizeSecurityPolicyRules));
        }

        /**
         * @param authorizeSecurityPolicyRules The policy rule. See `authorize_security_policy_rules` below.
         * 
         * @return builder
         * 
         */
        public Builder authorizeSecurityPolicyRules(EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs... authorizeSecurityPolicyRules) {
            return authorizeSecurityPolicyRules(List.of(authorizeSecurityPolicyRules));
        }

        /**
         * @param cameraRedirect Whether to enable local camera redirection. Valid values: `on`, `off`.
         * 
         * @return builder
         * 
         */
        public Builder cameraRedirect(@Nullable Output<String> cameraRedirect) {
            $.cameraRedirect = cameraRedirect;
            return this;
        }

        /**
         * @param cameraRedirect Whether to enable local camera redirection. Valid values: `on`, `off`.
         * 
         * @return builder
         * 
         */
        public Builder cameraRedirect(String cameraRedirect) {
            return cameraRedirect(Output.of(cameraRedirect));
        }

        /**
         * @param clipboard The clipboard policy. Valid values: `off`, `read`, `readwrite`.
         * 
         * @return builder
         * 
         */
        public Builder clipboard(@Nullable Output<String> clipboard) {
            $.clipboard = clipboard;
            return this;
        }

        /**
         * @param clipboard The clipboard policy. Valid values: `off`, `read`, `readwrite`.
         * 
         * @return builder
         * 
         */
        public Builder clipboard(String clipboard) {
            return clipboard(Output.of(clipboard));
        }

        /**
         * @param domainList The list of domain.
         * 
         * @return builder
         * 
         */
        public Builder domainList(@Nullable Output<String> domainList) {
            $.domainList = domainList;
            return this;
        }

        /**
         * @param domainList The list of domain.
         * 
         * @return builder
         * 
         */
        public Builder domainList(String domainList) {
            return domainList(Output.of(domainList));
        }

        /**
         * @param htmlAccess The access of html5. Valid values: `off`, `on`.
         * 
         * @return builder
         * 
         */
        public Builder htmlAccess(@Nullable Output<String> htmlAccess) {
            $.htmlAccess = htmlAccess;
            return this;
        }

        /**
         * @param htmlAccess The access of html5. Valid values: `off`, `on`.
         * 
         * @return builder
         * 
         */
        public Builder htmlAccess(String htmlAccess) {
            return htmlAccess(Output.of(htmlAccess));
        }

        /**
         * @param htmlFileTransfer The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
         * 
         * @return builder
         * 
         */
        public Builder htmlFileTransfer(@Nullable Output<String> htmlFileTransfer) {
            $.htmlFileTransfer = htmlFileTransfer;
            return this;
        }

        /**
         * @param htmlFileTransfer The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
         * 
         * @return builder
         * 
         */
        public Builder htmlFileTransfer(String htmlFileTransfer) {
            return htmlFileTransfer(Output.of(htmlFileTransfer));
        }

        /**
         * @param localDrive Local drive redirect policy. Valid values: `  readwrite `, `off`, `read`.
         * 
         * @return builder
         * 
         */
        public Builder localDrive(@Nullable Output<String> localDrive) {
            $.localDrive = localDrive;
            return this;
        }

        /**
         * @param localDrive Local drive redirect policy. Valid values: `  readwrite `, `off`, `read`.
         * 
         * @return builder
         * 
         */
        public Builder localDrive(String localDrive) {
            return localDrive(Output.of(localDrive));
        }

        /**
         * @param policyGroupName The name of policy group.
         * 
         * @return builder
         * 
         */
        public Builder policyGroupName(@Nullable Output<String> policyGroupName) {
            $.policyGroupName = policyGroupName;
            return this;
        }

        /**
         * @param policyGroupName The name of policy group.
         * 
         * @return builder
         * 
         */
        public Builder policyGroupName(String policyGroupName) {
            return policyGroupName(Output.of(policyGroupName));
        }

        /**
         * @param recording Whether to enable screen recording. Valid values: `off`, `all-time`, `period`.
         * 
         * @return builder
         * 
         */
        public Builder recording(@Nullable Output<String> recording) {
            $.recording = recording;
            return this;
        }

        /**
         * @param recording Whether to enable screen recording. Valid values: `off`, `all-time`, `period`.
         * 
         * @return builder
         * 
         */
        public Builder recording(String recording) {
            return recording(Output.of(recording));
        }

        /**
         * @param recordingEndTime The end time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
         * 
         * @return builder
         * 
         */
        public Builder recordingEndTime(@Nullable Output<String> recordingEndTime) {
            $.recordingEndTime = recordingEndTime;
            return this;
        }

        /**
         * @param recordingEndTime The end time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
         * 
         * @return builder
         * 
         */
        public Builder recordingEndTime(String recordingEndTime) {
            return recordingEndTime(Output.of(recordingEndTime));
        }

        /**
         * @param recordingExpires The screen recording video retention. Valid values between 30 and 180. This return value is meaningful only when the value of `recording` is `period` or `all-time`.
         * 
         * @return builder
         * 
         */
        public Builder recordingExpires(@Nullable Output<Integer> recordingExpires) {
            $.recordingExpires = recordingExpires;
            return this;
        }

        /**
         * @param recordingExpires The screen recording video retention. Valid values between 30 and 180. This return value is meaningful only when the value of `recording` is `period` or `all-time`.
         * 
         * @return builder
         * 
         */
        public Builder recordingExpires(Integer recordingExpires) {
            return recordingExpires(Output.of(recordingExpires));
        }

        /**
         * @param recordingFps The fps of recording. Valid values: `2`, `5`, `10`, `15`.
         * 
         * @return builder
         * 
         */
        public Builder recordingFps(@Nullable Output<Integer> recordingFps) {
            $.recordingFps = recordingFps;
            return this;
        }

        /**
         * @param recordingFps The fps of recording. Valid values: `2`, `5`, `10`, `15`.
         * 
         * @return builder
         * 
         */
        public Builder recordingFps(Integer recordingFps) {
            return recordingFps(Output.of(recordingFps));
        }

        /**
         * @param recordingStartTime The start time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
         * 
         * @return builder
         * 
         */
        public Builder recordingStartTime(@Nullable Output<String> recordingStartTime) {
            $.recordingStartTime = recordingStartTime;
            return this;
        }

        /**
         * @param recordingStartTime The start time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
         * 
         * @return builder
         * 
         */
        public Builder recordingStartTime(String recordingStartTime) {
            return recordingStartTime(Output.of(recordingStartTime));
        }

        /**
         * @param status The status of policy.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of policy.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param usbRedirect The usb redirect policy. Valid values: `off`, `on`.
         * 
         * @return builder
         * 
         */
        public Builder usbRedirect(@Nullable Output<String> usbRedirect) {
            $.usbRedirect = usbRedirect;
            return this;
        }

        /**
         * @param usbRedirect The usb redirect policy. Valid values: `off`, `on`.
         * 
         * @return builder
         * 
         */
        public Builder usbRedirect(String usbRedirect) {
            return usbRedirect(Output.of(usbRedirect));
        }

        /**
         * @param visualQuality The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
         * 
         * @return builder
         * 
         */
        public Builder visualQuality(@Nullable Output<String> visualQuality) {
            $.visualQuality = visualQuality;
            return this;
        }

        /**
         * @param visualQuality The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
         * 
         * @return builder
         * 
         */
        public Builder visualQuality(String visualQuality) {
            return visualQuality(Output.of(visualQuality));
        }

        /**
         * @param watermark The watermark policy. Valid values: `off`, `on`.
         * 
         * @return builder
         * 
         */
        public Builder watermark(@Nullable Output<String> watermark) {
            $.watermark = watermark;
            return this;
        }

        /**
         * @param watermark The watermark policy. Valid values: `off`, `on`.
         * 
         * @return builder
         * 
         */
        public Builder watermark(String watermark) {
            return watermark(Output.of(watermark));
        }

        /**
         * @param watermarkTransparency The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
         * 
         * @return builder
         * 
         */
        public Builder watermarkTransparency(@Nullable Output<String> watermarkTransparency) {
            $.watermarkTransparency = watermarkTransparency;
            return this;
        }

        /**
         * @param watermarkTransparency The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
         * 
         * @return builder
         * 
         */
        public Builder watermarkTransparency(String watermarkTransparency) {
            return watermarkTransparency(Output.of(watermarkTransparency));
        }

        /**
         * @param watermarkType The type of watemark. Valid values: `EndUserId`, `HostName`.
         * 
         * @return builder
         * 
         */
        public Builder watermarkType(@Nullable Output<String> watermarkType) {
            $.watermarkType = watermarkType;
            return this;
        }

        /**
         * @param watermarkType The type of watemark. Valid values: `EndUserId`, `HostName`.
         * 
         * @return builder
         * 
         */
        public Builder watermarkType(String watermarkType) {
            return watermarkType(Output.of(watermarkType));
        }

        public EcdPolicyGroupState build() {
            return $;
        }
    }

}
