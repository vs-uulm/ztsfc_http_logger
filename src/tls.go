// The logger is based basically on the logrus implementation
// (see https://github.com/sirupsen/logrus),
// but has extensions, that allow some customization of logging messages.
//
// The logger's tls.go contains:
//   - an internal function getTLSVersionName, that decodes the uint16 const into a TLS version name

package ztsfc_http_logger

// GetTLSVersionName converts uint16 value of http request TLS.Version field into a TSL version name
func getTLSVersionName(input uint16) string {
	switch input {
	case 0x0300:
		return "VersionSSL30"
	case 0x0301:
		return "VersionTLS10"
	case 0x0302:
		return "VersionTLS11"
	case 0x0303:
		return "VersionTLS12"
	case 0x0304:
		return "VersionTLS13"
	default:
		return "unknown"
	}
}
