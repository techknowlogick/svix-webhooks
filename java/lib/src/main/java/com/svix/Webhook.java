package com.svix;

import java.net.http.HttpHeaders;
import java.nio.charset.StandardCharsets;
import java.security.InvalidKeyException;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Base64;
import java.util.List;
import java.util.Optional;
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;

import com.svix.exceptions.WebhookVerificationException;

public final class Webhook {
	static final String MSG_ID_KEY = "svix-id";
	static final String MSG_SIGNATURE_KEY = "svix-signature";
	static final String MSG_TIMESTAMP_KEY = "svix-timestamp";
	private static final String HMAC_SHA256 = "HmacSHA256";

	private final byte[] key;

	public Webhook(final String secret) {
		this.key = Base64.getDecoder().decode(secret);
	}

	public void verify(final String payload, final HttpHeaders headers) throws WebhookVerificationException {
		Optional<String> msgId = headers.firstValue(MSG_ID_KEY);
		List<String> msgSignature = headers.allValues(MSG_SIGNATURE_KEY);
		Optional<String> msgTimestamp = headers.firstValue(MSG_TIMESTAMP_KEY);

		if (msgId.isEmpty() || msgSignature.isEmpty() || msgTimestamp.isEmpty()) {
			throw new WebhookVerificationException("Missing required headers");
		}

		String toSign = String.format("%s.%s.%s", msgId.get(), msgTimestamp.get(), payload);
		String expectedSignature = sign(toSign);
		for (String versionedSignature : msgSignature) {
			String[] sigParts = versionedSignature.split(",");
			if (sigParts.length < 2) {
				continue;
			}
			String version = sigParts[0];
			if (!version.equals("v1")) {
				continue;
			}
			String signature = sigParts[1];
			if (MessageDigest.isEqual(signature.getBytes(), expectedSignature.getBytes())) {
				return;
			}
		}
		throw new WebhookVerificationException("No matching signature found");
	}

	private String sign(final String toSign) throws WebhookVerificationException {
		try {
			Mac sha512Hmac = Mac.getInstance(HMAC_SHA256);
			SecretKeySpec keySpec = new SecretKeySpec(key, HMAC_SHA256);
			sha512Hmac.init(keySpec);
			byte[] macData = sha512Hmac.doFinal(toSign.getBytes(StandardCharsets.UTF_8));
			return Base64.getEncoder().encodeToString(macData);
		} catch (InvalidKeyException | NoSuchAlgorithmException e) {
			throw new WebhookVerificationException(e.getMessage());
		}
	}
}