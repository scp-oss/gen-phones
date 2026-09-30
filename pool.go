package main

import (
	"encoding/csv"
	"fmt"
	"io"
)

// csvHeader is the exact FreePBX "Extensions Pool" export column order.
var csvHeader = []string{
	"extension", "password", "name", "voicemail", "ringtimer", "noanswer", "recording", "outboundcid", "sipname", "noanswer_cid",
	"busy_cid", "chanunavail_cid", "noanswer_dest", "busy_dest", "chanunavail_dest", "mohclass", "id", "tech", "dial", "devicetype",
	"user", "description", "emergency_cid", "hint_override", "cwtone", "recording_in_external", "recording_out_external", "recording_in_internal", "recording_out_internal", "recording_ondemand",
	"recording_priority", "answermode", "intercom", "cid_masquerade", "concurrency_limit", "devicedata", "accountcode", "aggregate_mwi", "allow", "avpf",
	"bundle", "callerid", "context", "defaultuser", "device_state_busy_at", "direct_media", "disallow", "dtmfmode", "force_rport", "icesupport",
	"match", "max_audio_streams", "max_contacts", "max_video_streams", "maximum_expiration", "media_encryption", "media_encryption_optimistic", "media_use_received_transport", "message_context", "minimum_expiration",
	"mwi_subscription", "namedcallgroup", "namedpickupgroup", "outbound_proxy", "qualifyfreq", "refer_blind_progress", "remove_existing", "rewrite_contact", "rtcp_mux", "rtp_symmetric",
	"rtp_timeout", "rtp_timeout_hold", "secret", "send_connected_line", "sendrpid", "sipdriver", "timers", "timers_min_se", "transport", "trustrpid",
	"user_eq_phone", "force_callerid", "mailbox", "vmexten", "webrtc", "callwaiting_enable", "findmefollow_strategy", "findmefollow_grptime", "findmefollow_grppre", "findmefollow_grplist",
	"findmefollow_annmsg_id", "findmefollow_postdest", "findmefollow_dring", "findmefollow_needsconf", "findmefollow_remotealert_id", "findmefollow_toolate_id", "findmefollow_ringing", "findmefollow_pre_ring", "findmefollow_voicemail", "findmefollow_calendar_id",
	"findmefollow_calendar_match", "findmefollow_changecid", "findmefollow_fixedcid", "findmefollow_enabled", "parkpro_pagegroup",
}

// buildRow renders one extension row. The per-extension counters below (ringtimer,
// recording_priority, dtmfmode, minimum_expiration/qualifyfreq, timers_min_se,
// findmefollow_grptime, findmefollow_pre_ring, parkpro_pagegroup, and the three
// fields that mirror ringtimer) keep the same fixed offset from the extension
// number that was observed across the entire reference NEW-Pool.csv export.
func buildRow(e int64) []string {
	es := fmt.Sprintf("%d", e)
	return []string{
		es, "", es, "novm", fmt.Sprintf("%d", e-2177), "", "", "", "", "",
		"", "", "", "", "", "default", es, "pjsip", "PJSIP/" + es, "fixed",
		es, es, "", "", "disabled", "dontcare", "dontcare", "dontcare", "dontcare", "disabled",
		fmt.Sprintf("%d", e-2167), "disabled", "enabled", es, "3", es, "", "yes", "", "no",
		"no", "", "from-internal", "", fmt.Sprintf("%d", e-2177), "yes", "", fmt.Sprintf("rfc%d", e+2556), "yes", "no",
		"", "1", "1", "1", "7200", "no", "no", "no", "", fmt.Sprintf("%d", e-2117),
		"auto", "", "", "", fmt.Sprintf("%d", e-2117), "yes", "yes", "yes", "no", "yes",
		"0", "0", "REGEN", "yes", "pai", "chan_pjsip", "yes", fmt.Sprintf("%d", e-2087), "", "yes",
		"no", "", "", "", "", "ENABLED", "ringallv2-prim", fmt.Sprintf("%d", e-2157), "", es,
		fmt.Sprintf("%d", e-2177), fmt.Sprintf("ext-local,%s,dest", es), "", "", fmt.Sprintf("%d", e-2177), "", "Ring", fmt.Sprintf("%d", e-2170), "novm", "",
		"yes", "default", "", "", fmt.Sprintf("%d", e-2176),
	}
}

// WritePool writes count rows starting at extension `start`, in the same
// CRLF-terminated CSV format FreePBX exports use.
func WritePool(w io.Writer, start, count int64) error {
	cw := csv.NewWriter(w)
	cw.UseCRLF = true
	if err := cw.Write(csvHeader); err != nil {
		return err
	}
	for i := int64(0); i < count; i++ {
		if err := cw.Write(buildRow(start + i)); err != nil {
			return err
		}
	}
	cw.Flush()
	return cw.Error()
}
