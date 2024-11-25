package pingdom

import (
	"fmt"
	"sort"
	"strconv"
)

// HttpCheck represents a Pingdom HTTP check.
type HttpCheck struct {
	CustomMessage            string            `json:"custom_message,omitempty"`
	Encryption               bool              `json:"encryption,omitempty"`
	Hostname                 string            `json:"hostname,omitempty"`
	IPV6                     bool              `json:"ipv6,omitempty"`
	IntegrationIds           []int             `json:"integrationids,omitempty"`
	Name                     string            `json:"name"`
	NotifyAgainEvery         int               `json:"notifyagainevery,omitempty"`
	NotifyWhenBackup         bool              `json:"notifywhenbackup,omitempty"`
	Password                 string            `json:"password,omitempty"`
	Paused                   bool              `json:"paused,omitempty"`
	Port                     int               `json:"port,omitempty"`
	PostData                 string            `json:"postdata,omitempty"`
	ProbeFilters             string            `json:"probe_filters,omitempty"`
	RequestHeaders           map[string]string `json:"requestheaders,omitempty"`
	Resolution               int               `json:"resolution,omitempty"`
	ResponseTimeThreshold    int               `json:"responsetime_threshold,omitempty"`
	SSLDownDaysBefore        *int              `json:"ssl_down_days_before,omitempty"`
	SendNotificationWhenDown int               `json:"sendnotificationwhendown,omitempty"`
	ShouldContain            string            `json:"shouldcontain,omitempty"`
	ShouldNotContain         string            `json:"shouldnotcontain,omitempty"`
	Tags                     string            `json:"tags,omitempty"`
	TeamIds                  []int             `json:"teamids,omitempty"`
	Url                      string            `json:"url,omitempty"`
	UserIds                  []int             `json:"userids,omitempty"`
	Username                 string            `json:"username,omitempty"`
	VerifyCertificate        *bool             `json:"verify_certificate,omitempty"`
}

// PingCheck represents a Pingdom ping check.
type PingCheck struct {
	Hostname                 string `json:"hostname,omitempty"`
	IntegrationIds           []int  `json:"integrationids,omitempty"`
	Name                     string `json:"name"`
	NotifyAgainEvery         int    `json:"notifyagainevery,omitempty"`
	NotifyWhenBackup         bool   `json:"notifywhenbackup,omitempty"`
	Paused                   bool   `json:"paused,omitempty"`
	ProbeFilters             string `json:"probe_filters,omitempty"`
	Resolution               int    `json:"resolution,omitempty"`
	ResponseTimeThreshold    int    `json:"responsetime_threshold,omitempty"`
	SendNotificationWhenDown int    `json:"sendnotificationwhendown,omitempty"`
	Tags                     string `json:"tags,omitempty"`
	TeamIds                  []int  `json:"teamids,omitempty"`
	UserIds                  []int  `json:"userids,omitempty"`
}

// TCPCheck represents a Pingdom TCP check.
type TCPCheck struct {
	CustomMessage            string `json:"custom_message,omitempty"`
	Hostname                 string `json:"hostname,omitempty"`
	IPV6                     bool   `json:"ipv6,omitempty"`
	IntegrationIds           []int  `json:"integrationids,omitempty"`
	Name                     string `json:"name"`
	NotifyAgainEvery         int    `json:"notifyagainevery,omitempty"`
	NotifyWhenBackup         bool   `json:"notifywhenbackup,omitempty"`
	Paused                   bool   `json:"paused,omitempty"`
	Port                     int    `json:"port"`
	ProbeFilters             string `json:"probe_filters,omitempty"`
	Resolution               int    `json:"resolution,omitempty"`
	ResponseTimeThreshold    int    `json:"responsetime_threshold,omitempty"`
	SendNotificationWhenDown int    `json:"sendnotificationwhendown,omitempty"`
	StringToExpect           string `json:"stringtoexpect,omitempty"`
	StringToSend             string `json:"stringtosend,omitempty"`
	Tags                     string `json:"tags,omitempty"`
	TeamIds                  []int  `json:"teamids,omitempty"`
	UserIds                  []int  `json:"userids,omitempty"`
}

// DNSCheck represents a Pingdom DNS check.
type DNSCheck struct {
	ExpectedIP               string `json:"expectedip,omitempty"`
	Hostname                 string `json:"hostname,omitempty"`
	IPV6                     bool   `json:"ipv6,omitempty"`
	IntegrationIds           []int  `json:"integrationids,omitempty"`
	Name                     string `json:"name"`
	NameServer               string `json:"nameserver,omitempty"`
	NotifyAgainEvery         int    `json:"notifyagainevery,omitempty"`
	NotifyWhenBackup         bool   `json:"notifywhenbackup,omitempty"`
	Paused                   bool   `json:"paused,omitempty"`
	ProbeFilters             string `json:"probe_filters,omitempty"`
	Resolution               int    `json:"resolution,omitempty"`
	SendNotificationWhenDown int    `json:"sendnotificationwhendown,omitempty"`
	Tags                     string `json:"tags,omitempty"`
	TeamIds                  []int  `json:"teamids,omitempty"`
	UserIds                  []int  `json:"userids,omitempty"`
}

// SummaryPerformanceRequest is the API request to Pingdom for a SummaryPerformance.
type SummaryPerformanceRequest struct {
	From          int
	Id            int
	IncludeUptime bool
	Order         string
	Probes        string
	Resolution    string
	To            int
}

// PutParams returns a map of parameters for an HttpCheck that can be sent along
// with an HTTP PUT request.
func (ck *HttpCheck) PutParams() map[string]string {
	m := map[string]string{
		"custom_message":   ck.CustomMessage,
		"encryption":       strconv.FormatBool(ck.Encryption),
		"host":             ck.Hostname,
		"integrationids":   intListToCDString(ck.IntegrationIds),
		"ipv6":             strconv.FormatBool(ck.IPV6),
		"name":             ck.Name,
		"notifyagainevery": strconv.Itoa(ck.NotifyAgainEvery),
		"notifywhenbackup": strconv.FormatBool(ck.NotifyWhenBackup),
		"paused":           strconv.FormatBool(ck.Paused),
		"postdata":         ck.PostData,
		"probe_filters":    ck.ProbeFilters,
		"tags":             ck.Tags,
		"teamids":          intListToCDString(ck.TeamIds),
		"url":              ck.Url,
		"userids":          intListToCDString(ck.UserIds),
	}

	if ck.Resolution != 0 {
		m["resolution"] = strconv.Itoa(ck.Resolution)
	}

	if ck.SendNotificationWhenDown != 0 {
		m["sendnotificationwhendown"] = strconv.Itoa(ck.SendNotificationWhenDown)
	}

	// Ignore zero values
	if ck.Port != 0 {
		m["port"] = strconv.Itoa(ck.Port)
	}

	if ck.SendNotificationWhenDown != 0 {
		m["sendnotificationwhendown"] = strconv.Itoa(ck.SendNotificationWhenDown)
	}

	if ck.ResponseTimeThreshold != 0 {
		m["responsetime_threshold"] = strconv.Itoa(ck.ResponseTimeThreshold)
	}

	if ck.VerifyCertificate != nil {
		m["verify_certificate"] = strconv.FormatBool(*ck.VerifyCertificate)
	}

	if ck.SSLDownDaysBefore != nil {
		m["ssl_down_days_before"] = strconv.Itoa(*ck.SSLDownDaysBefore)
	}

	// ShouldContain and ShouldNotContain are mutually exclusive.
	// But we must define one so they can be emptied if required.
	if ck.ShouldContain != "" {
		m["shouldcontain"] = ck.ShouldContain
	} else {
		m["shouldnotcontain"] = ck.ShouldNotContain
	}

	// Convert auth
	if ck.Username != "" {
		m["auth"] = fmt.Sprintf("%s:%s", ck.Username, ck.Password)
	}

	// Convert headers
	var headers []string
	for k := range ck.RequestHeaders {
		headers = append(headers, k)
	}
	sort.Strings(headers)
	for i, k := range headers {
		m[fmt.Sprintf("requestheader%d", i)] = fmt.Sprintf("%s:%s", k, ck.RequestHeaders[k])
	}

	return m
}

// PostParams returns a map of parameters for an HttpCheck that can be sent along
// with an HTTP POST request. They are the same than the Put params, but
// empty strings cleared out, to avoid Pingdom API reject the request.
func (ck *HttpCheck) PostParams() map[string]string {
	params := ck.PutParams()

	for k, v := range params {
		if v == "" {
			delete(params, k)
		}
	}
	params["type"] = "http"

	return params
}

// Valid determines whether the HttpCheck contains valid fields.  This can be
// used to guard against sending illegal values to the Pingdom API.
func (ck *HttpCheck) Valid() error {
	if err := validCommonParameters(ck.Name, ck.Hostname, ck.Resolution); err != nil {
		return err
	}

	if ck.ShouldContain != "" && ck.ShouldNotContain != "" {
		return fmt.Errorf("`ShouldContain` and `ShouldNotContain` must not be declared at the same time")
	}

	return nil
}

// PutParams returns a map of parameters for a PingCheck that can be sent along
// with an HTTP PUT request.
func (ck *PingCheck) PutParams() map[string]string {
	m := map[string]string{
		"host":             ck.Hostname,
		"integrationids":   intListToCDString(ck.IntegrationIds),
		"name":             ck.Name,
		"notifyagainevery": strconv.Itoa(ck.NotifyAgainEvery),
		"notifywhenbackup": strconv.FormatBool(ck.NotifyWhenBackup),
		"paused":           strconv.FormatBool(ck.Paused),
		"tags":             ck.Tags,
		"probe_filters":    ck.ProbeFilters,
		"teamids":          intListToCDString(ck.TeamIds),
		"userids":          intListToCDString(ck.UserIds),
	}

	if ck.Resolution != 0 {
		m["resolution"] = strconv.Itoa(ck.Resolution)
	}

	if ck.SendNotificationWhenDown != 0 {
		m["sendnotificationwhendown"] = strconv.Itoa(ck.SendNotificationWhenDown)
	}

	if ck.ResponseTimeThreshold != 0 {
		m["responsetime_threshold"] = strconv.Itoa(ck.ResponseTimeThreshold)
	}

	return m
}

// PostParams returns a map of parameters for a PingCheck that can be sent along
// with an HTTP POST request. Same as PUT.
func (ck *PingCheck) PostParams() map[string]string {
	params := ck.PutParams()

	for k, v := range params {
		if v == "" {
			delete(params, k)
		}
	}

	params["type"] = "ping"
	return params
}

// Valid determines whether the PingCheck contains valid fields.  This can be
// used to guard against sending illegal values to the Pingdom API.
func (ck *PingCheck) Valid() error {
	if err := validCommonParameters(ck.Name, ck.Hostname, ck.Resolution); err != nil {
		return err
	}

	return nil
}

// PutParams returns a map of parameters for a TCPCheck that can be sent along
// with an HTTP PUT request.
func (ck *TCPCheck) PutParams() map[string]string {
	m := map[string]string{
		"custom_message":   ck.CustomMessage,
		"host":             ck.Hostname,
		"integrationids":   intListToCDString(ck.IntegrationIds),
		"ipv6":             strconv.FormatBool(ck.IPV6),
		"name":             ck.Name,
		"notifyagainevery": strconv.Itoa(ck.NotifyAgainEvery),
		"notifywhenbackup": strconv.FormatBool(ck.NotifyWhenBackup),
		"paused":           strconv.FormatBool(ck.Paused),
		"port":             strconv.Itoa(ck.Port),
		"probe_filters":    ck.ProbeFilters,
		"tags":             ck.Tags,
		"teamids":          intListToCDString(ck.TeamIds),
		"userids":          intListToCDString(ck.UserIds),
	}

	if ck.Resolution != 0 {
		m["resolution"] = strconv.Itoa(ck.Resolution)
	}

	if ck.ResponseTimeThreshold != 0 {
		m["responsetime_threshold"] = strconv.Itoa(ck.ResponseTimeThreshold)
	}

	if ck.SendNotificationWhenDown != 0 {
		m["sendnotificationwhendown"] = strconv.Itoa(ck.SendNotificationWhenDown)
	}

	if ck.StringToSend != "" {
		m["stringtosend"] = ck.StringToSend
	}

	if ck.StringToExpect != "" {
		m["stringtoexpect"] = ck.StringToExpect
	}

	return m
}

// PostParams returns a map of parameters for a TCPCheck that can be sent along
// with an HTTP POST request. Same as PUT.
func (ck *TCPCheck) PostParams() map[string]string {
	params := ck.PutParams()

	for k, v := range params {
		if v == "" {
			delete(params, k)
		}
	}

	params["type"] = "tcp"
	return params
}

// Valid determines whether the TCPCheck contains valid fields.  This can be
// used to guard against sending illegal values to the Pingdom API.
func (ck *TCPCheck) Valid() error {
	if err := validCommonParameters(ck.Name, ck.Hostname, ck.Resolution); err != nil {
		return err
	}

	if ck.Port < 1 || ck.Port > 65535 {
		return fmt.Errorf("Invalid value for `Port`.  Must contain an integer >= 1 and <= 65535")
	}

	return nil
}

// PutParams returns a map of parameters for a DNSCheck that can be sent along
// with an HTTP PUT request.
func (ck *DNSCheck) PutParams() map[string]string {
	m := map[string]string{
		"expectedip":       ck.ExpectedIP,
		"host":             ck.Hostname,
		"integrationids":   intListToCDString(ck.IntegrationIds),
		"ipv6":             strconv.FormatBool(ck.IPV6),
		"name":             ck.Name,
		"nameserver":       ck.NameServer,
		"notifyagainevery": strconv.Itoa(ck.NotifyAgainEvery),
		"notifywhenbackup": strconv.FormatBool(ck.NotifyWhenBackup),
		"paused":           strconv.FormatBool(ck.Paused),
		"probe_filters":    ck.ProbeFilters,
		"tags":             ck.Tags,
		"teamids":          intListToCDString(ck.TeamIds),
		"userids":          intListToCDString(ck.UserIds),
	}

	if ck.Resolution != 0 {
		m["resolution"] = strconv.Itoa(ck.Resolution)
	}

	if ck.SendNotificationWhenDown != 0 {
		m["sendnotificationwhendown"] = strconv.Itoa(ck.SendNotificationWhenDown)
	}

	return m
}

// PostParams returns a map of parameters for a DNSCheck that can be sent along
// with an HTTP POST request. Same as PUT.
func (ck *DNSCheck) PostParams() map[string]string {
	params := ck.PutParams()

	for k, v := range params {
		if v == "" {
			delete(params, k)
		}
	}

	params["type"] = "dns"
	return params
}

// Valid determines whether the DNSCheck contains valid fields.  This can be
// used to guard against sending illegal values to the Pingdom API.
func (ck *DNSCheck) Valid() error {
	if err := validCommonParameters(ck.Name, ck.Hostname, ck.Resolution); err != nil {
		return err
	}

	if ck.ExpectedIP == "" {
		return fmt.Errorf("invalid value for `ExpectedIP`, must contain non-empty string")
	}

	if ck.NameServer == "" {
		return fmt.Errorf("invalid value for `NameServer`, must contain non-empty string")
	}

	return nil
}

func intListToCDString(integers []int) string {
	var CDString string
	for i, item := range integers {
		if i == 0 {
			CDString = strconv.Itoa(item)
		} else {
			CDString = fmt.Sprintf("%v,%d", CDString, item)
		}
	}
	return CDString
}

func validCommonParameters(name string, hostname string, resolution int) error {
	if name == "" {
		return fmt.Errorf("invalid value for `Name`, must contain non-empty string")
	}

	if hostname == "" {
		return fmt.Errorf("invalid value for `Hostname`, must contain non-empty string")
	}

	// if resolution value is 0, it will be set to default value which is 5.
	if resolution != 0 && resolution != 1 && resolution != 5 && resolution != 15 &&
		resolution != 30 && resolution != 60 {
		return fmt.Errorf("invalid value %v for `Resolution`, allowed values are [1,5,15,30,60]", resolution)
	}

	return nil
}

// Valid determines whether a SummaryPerformanceRequest contains valid fields for the Pingdom API.
func (csr SummaryPerformanceRequest) Valid() error {
	if csr.Id == 0 {
		return ErrMissingId
	}

	if csr.Resolution != "" && csr.Resolution != "hour" && csr.Resolution != "day" && csr.Resolution != "week" {
		return ErrBadResolution
	}
	return nil
}

// GetParams returns a map of params for a Pingdom SummaryPerformanceRequest.
func (csr SummaryPerformanceRequest) GetParams() (params map[string]string) {
	params = make(map[string]string)

	if csr.Resolution != "" {
		params["resolution"] = csr.Resolution
	}

	if csr.IncludeUptime {
		params["includeuptime"] = "true"
	}

	return
}
