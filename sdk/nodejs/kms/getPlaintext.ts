// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getPlaintext(args: GetPlaintextArgs, opts?: pulumi.InvokeOptions): Promise<GetPlaintextResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:kms/getPlaintext:getPlaintext", {
        "ciphertextBlob": args.ciphertextBlob,
        "encryptionContext": args.encryptionContext,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlaintext.
 */
export interface GetPlaintextArgs {
    /**
     * The ciphertext to be decrypted.
     */
    readonly ciphertextBlob: string;
    readonly encryptionContext?: {[key: string]: string};
}

/**
 * A collection of values returned by getPlaintext.
 */
export interface GetPlaintextResult {
    readonly ciphertextBlob: string;
    readonly encryptionContext?: {[key: string]: string};
    /**
     * The globally unique ID of the CMK. It is the ID of the CMK used to decrypt ciphertext.
     */
    readonly keyId: string;
    /**
     * The decrypted plaintext.
     */
    readonly plaintext: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
