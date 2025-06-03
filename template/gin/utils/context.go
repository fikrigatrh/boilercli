package utils

import "context"

func TraceHeaderWrapToLog(c context.Context) map[string]string {
	// check if trace header is exist
	switch c.Value("trace_header").(type) {
	case map[string]string:
		mapTraceHeader := c.Value("trace_header").(map[string]string)
		return map[string]string{
			"trace_srvc_id": mapTraceHeader["trace_srvc_id"],
			"trace_gw_id":   mapTraceHeader["trace_gw_id"],
			"trace_ui_id":   mapTraceHeader["trace_ui_id"],
		}
	default:
		return map[string]string{
			"trace_srvc_id": "",
			"trace_gw_id":   "",
			"trace_ui_id":   "",
		}
	}
}
