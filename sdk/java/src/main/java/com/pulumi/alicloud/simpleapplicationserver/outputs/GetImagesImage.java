// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.simpleapplicationserver.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetImagesImage {
    /**
     * @return The description of the image.
     * 
     */
    private String description;
    /**
     * @return The ID of the Instance Image.
     * 
     */
    private String id;
    /**
     * @return The ID of the image.
     * 
     */
    private String imageId;
    /**
     * @return The name of the resource.
     * 
     */
    private String imageName;
    /**
     * @return The type of the image. Valid values: `app`, `custom`, `system`.
     * 
     */
    private String imageType;
    /**
     * @return The platform of Plan supported.
     * 
     */
    private String platform;

    private GetImagesImage() {}
    /**
     * @return The description of the image.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The ID of the Instance Image.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the image.
     * 
     */
    public String imageId() {
        return this.imageId;
    }
    /**
     * @return The name of the resource.
     * 
     */
    public String imageName() {
        return this.imageName;
    }
    /**
     * @return The type of the image. Valid values: `app`, `custom`, `system`.
     * 
     */
    public String imageType() {
        return this.imageType;
    }
    /**
     * @return The platform of Plan supported.
     * 
     */
    public String platform() {
        return this.platform;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetImagesImage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private String imageId;
        private String imageName;
        private String imageType;
        private String platform;
        public Builder() {}
        public Builder(GetImagesImage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.imageId = defaults.imageId;
    	      this.imageName = defaults.imageName;
    	      this.imageType = defaults.imageType;
    	      this.platform = defaults.platform;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder imageId(String imageId) {
            this.imageId = Objects.requireNonNull(imageId);
            return this;
        }
        @CustomType.Setter
        public Builder imageName(String imageName) {
            this.imageName = Objects.requireNonNull(imageName);
            return this;
        }
        @CustomType.Setter
        public Builder imageType(String imageType) {
            this.imageType = Objects.requireNonNull(imageType);
            return this;
        }
        @CustomType.Setter
        public Builder platform(String platform) {
            this.platform = Objects.requireNonNull(platform);
            return this;
        }
        public GetImagesImage build() {
            final var o = new GetImagesImage();
            o.description = description;
            o.id = id;
            o.imageId = imageId;
            o.imageName = imageName;
            o.imageType = imageType;
            o.platform = platform;
            return o;
        }
    }
}
