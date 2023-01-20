// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetHoneypotImagesImage {
    /**
     * @return The name of the honeypot image display.
     * 
     */
    private String honeypotImageDisplayName;
    /**
     * @return The image ID of the honeypot.
     * 
     */
    private String honeypotImageId;
    /**
     * @return Honeypot mirror name.
     * 
     */
    private String honeypotImageName;
    /**
     * @return Honeypot mirror type.
     * 
     */
    private String honeypotImageType;
    /**
     * @return Honeypot Mirror version.
     * 
     */
    private String honeypotImageVersion;
    /**
     * @return The image ID of the honeypot.The value is the same as `honeypot_image_id`.
     * 
     */
    private String id;
    /**
     * @return Ports supported by honeypots. In JSON format. Contains the following fields:-**log_type**: log type-**proto**: Support Protocol-**description**: description-**ports**: supports Port collection-**port_str**: supports port strings-**type**: type
     * 
     */
    private String multiports;
    /**
     * @return Honeypot-supported protocols.
     * 
     */
    private String proto;
    /**
     * @return Honeypot service port.
     * 
     */
    private String servicePort;
    /**
     * @return Honeypot configuration parameter template.
     * 
     */
    private String template;

    private GetHoneypotImagesImage() {}
    /**
     * @return The name of the honeypot image display.
     * 
     */
    public String honeypotImageDisplayName() {
        return this.honeypotImageDisplayName;
    }
    /**
     * @return The image ID of the honeypot.
     * 
     */
    public String honeypotImageId() {
        return this.honeypotImageId;
    }
    /**
     * @return Honeypot mirror name.
     * 
     */
    public String honeypotImageName() {
        return this.honeypotImageName;
    }
    /**
     * @return Honeypot mirror type.
     * 
     */
    public String honeypotImageType() {
        return this.honeypotImageType;
    }
    /**
     * @return Honeypot Mirror version.
     * 
     */
    public String honeypotImageVersion() {
        return this.honeypotImageVersion;
    }
    /**
     * @return The image ID of the honeypot.The value is the same as `honeypot_image_id`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Ports supported by honeypots. In JSON format. Contains the following fields:-**log_type**: log type-**proto**: Support Protocol-**description**: description-**ports**: supports Port collection-**port_str**: supports port strings-**type**: type
     * 
     */
    public String multiports() {
        return this.multiports;
    }
    /**
     * @return Honeypot-supported protocols.
     * 
     */
    public String proto() {
        return this.proto;
    }
    /**
     * @return Honeypot service port.
     * 
     */
    public String servicePort() {
        return this.servicePort;
    }
    /**
     * @return Honeypot configuration parameter template.
     * 
     */
    public String template() {
        return this.template;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHoneypotImagesImage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String honeypotImageDisplayName;
        private String honeypotImageId;
        private String honeypotImageName;
        private String honeypotImageType;
        private String honeypotImageVersion;
        private String id;
        private String multiports;
        private String proto;
        private String servicePort;
        private String template;
        public Builder() {}
        public Builder(GetHoneypotImagesImage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.honeypotImageDisplayName = defaults.honeypotImageDisplayName;
    	      this.honeypotImageId = defaults.honeypotImageId;
    	      this.honeypotImageName = defaults.honeypotImageName;
    	      this.honeypotImageType = defaults.honeypotImageType;
    	      this.honeypotImageVersion = defaults.honeypotImageVersion;
    	      this.id = defaults.id;
    	      this.multiports = defaults.multiports;
    	      this.proto = defaults.proto;
    	      this.servicePort = defaults.servicePort;
    	      this.template = defaults.template;
        }

        @CustomType.Setter
        public Builder honeypotImageDisplayName(String honeypotImageDisplayName) {
            this.honeypotImageDisplayName = Objects.requireNonNull(honeypotImageDisplayName);
            return this;
        }
        @CustomType.Setter
        public Builder honeypotImageId(String honeypotImageId) {
            this.honeypotImageId = Objects.requireNonNull(honeypotImageId);
            return this;
        }
        @CustomType.Setter
        public Builder honeypotImageName(String honeypotImageName) {
            this.honeypotImageName = Objects.requireNonNull(honeypotImageName);
            return this;
        }
        @CustomType.Setter
        public Builder honeypotImageType(String honeypotImageType) {
            this.honeypotImageType = Objects.requireNonNull(honeypotImageType);
            return this;
        }
        @CustomType.Setter
        public Builder honeypotImageVersion(String honeypotImageVersion) {
            this.honeypotImageVersion = Objects.requireNonNull(honeypotImageVersion);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder multiports(String multiports) {
            this.multiports = Objects.requireNonNull(multiports);
            return this;
        }
        @CustomType.Setter
        public Builder proto(String proto) {
            this.proto = Objects.requireNonNull(proto);
            return this;
        }
        @CustomType.Setter
        public Builder servicePort(String servicePort) {
            this.servicePort = Objects.requireNonNull(servicePort);
            return this;
        }
        @CustomType.Setter
        public Builder template(String template) {
            this.template = Objects.requireNonNull(template);
            return this;
        }
        public GetHoneypotImagesImage build() {
            final var o = new GetHoneypotImagesImage();
            o.honeypotImageDisplayName = honeypotImageDisplayName;
            o.honeypotImageId = honeypotImageId;
            o.honeypotImageName = honeypotImageName;
            o.honeypotImageType = honeypotImageType;
            o.honeypotImageVersion = honeypotImageVersion;
            o.id = id;
            o.multiports = multiports;
            o.proto = proto;
            o.servicePort = servicePort;
            o.template = template;
            return o;
        }
    }
}
