// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package flex

import (
	"crypto/hmac"
	"encoding/hex"
	"regexp"
	"strings"

	"golang.org/x/crypto/sha3"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SuppressHashedRawSecret(k, old, new string, d *schema.ResourceData) bool {
	if len(d.Id()) == 0 {
		return false
	}
	isSecretRef, _ := regexp.MatchString("[{]{1}(\\b(vault)\\b[:]{2}([ a-zA-Z0-9_-]*)[.]{0,1}(.*))[}]{1}", new)
	if isSecretRef {
		return false
	}
	parts, _ := SepIdParts(d.Id(), "/")
	secret := parts[1]
	mac := hmac.New(sha3.New512, []byte(secret))
	mac.Write([]byte(new))
	secureHmac := hex.EncodeToString(mac.Sum(nil))
	return cmp.Equal(strings.Join([]string{"hash", "SHA3-512", secureHmac}, ":"), old)
}
